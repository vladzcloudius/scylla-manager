// Copyright (C) 2017 ScyllaDB

package main

import (
	"context"
	"net/http"
	"os/user"
	"path/filepath"
	"sync"
	"time"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	log "github.com/scylladb/golog"
	"github.com/scylladb/mermaid/cluster"
	"github.com/scylladb/mermaid/healthcheck"
	"github.com/scylladb/mermaid/internal/kv"
	"github.com/scylladb/mermaid/repair"
	"github.com/scylladb/mermaid/restapi"
	"github.com/scylladb/mermaid/sched"
)

type server struct {
	config  *serverConfig
	session *gocql.Session
	logger  log.Logger

	clusterSvc *cluster.Service
	repairSvc  *repair.Service
	schedSvc   *sched.Service

	httpServer       *http.Server
	httpsServer      *http.Server
	prometheusServer *http.Server
	errCh            chan error
}

func newServer(config *serverConfig, logger log.Logger) (*server, error) {
	session, err := gocqlConfig(config).CreateSession()
	if err != nil {
		return nil, errors.Wrapf(err, "database")
	}

	s := &server{
		config:  config,
		session: session,
		logger:  logger,

		errCh: make(chan error, 2),
	}

	if err := s.makeServices(); err != nil {
		return nil, err
	}

	s.makeHTTPServers()

	return s, nil
}

func (s *server) makeServices() error {
	var err error

	u, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}
	keyStore, err := kv.NewFsStore(filepath.Join(u.HomeDir, ".certs"))
	if err != nil {
		return errors.Wrap(err, "failed to create key store")
	}

	s.clusterSvc, err = cluster.NewService(s.session, keyStore, s.logger.Named("cluster"))
	if err != nil {
		return errors.Wrapf(err, "cluster service")
	}
	s.clusterSvc.SetOnChangeListener(s.onClusterChange)

	s.repairSvc, err = repair.NewService(
		s.session,
		s.config.Repair,
		s.clusterSvc.GetClusterByID,
		s.clusterSvc.Client,
		s.logger.Named("repair"),
	)
	if err != nil {
		return errors.Wrapf(err, "repair service")
	}

	s.schedSvc, err = sched.NewService(
		s.session,
		s.clusterSvc.GetClusterByID,
		s.logger.Named("scheduler"),
	)
	if err != nil {
		return errors.Wrapf(err, "scheduler service")
	}

	// register runners
	s.schedSvc.SetRunner(sched.HealthCheckTask, healthcheck.NewRunner(s.clusterSvc.GetClusterByID, s.clusterSvc.Client))
	s.schedSvc.SetRunner(sched.RepairTask, repair.Runner{Service: s.repairSvc})

	return nil
}

func (s *server) onClusterChange(ctx context.Context, c cluster.Change) error {
	if c.Type == cluster.Create {
		if err := s.schedSvc.PutTask(ctx, autoRepairTask(c.ID)); err != nil {
			return errors.Wrap(err, "failed to add scheduled tasks")
		}
	}

	return nil
}

func (s *server) makeHTTPServers() {
	h := restapi.New(&restapi.Services{
		Cluster:   s.clusterSvc,
		Repair:    s.repairSvc,
		Scheduler: s.schedSvc,
	}, s.logger.Named("restapi"))

	if s.config.HTTP != "" {
		s.httpServer = &http.Server{Addr: s.config.HTTP, Handler: h}
	}
	if s.config.HTTPS != "" {
		s.httpsServer = &http.Server{Addr: s.config.HTTPS, Handler: h}
	}
	if s.config.Prometheus != "" {
		s.prometheusServer = &http.Server{Addr: s.config.Prometheus, Handler: restapi.NewPrometheus()}
	}
}

func (s *server) initServices(ctx context.Context) error {
	if err := s.repairSvc.Init(ctx); err != nil {
		return errors.Wrapf(err, "repair service")
	}
	if err := s.schedSvc.Init(ctx); err != nil {
		return errors.Wrapf(err, "schedule service")
	}
	return nil
}

func (s *server) startHTTPServers(ctx context.Context) {
	if s.httpServer != nil {
		s.logger.Info(ctx, "Starting HTTP", "address", s.httpServer.Addr)
		go func() {
			s.errCh <- s.httpServer.ListenAndServe()
		}()
	}

	if s.httpsServer != nil {
		s.logger.Info(ctx, "Starting HTTPS", "address", s.httpsServer.Addr)
		go func() {
			s.errCh <- s.httpsServer.ListenAndServeTLS(s.config.TLSCertFile, s.config.TLSKeyFile)
		}()
	}

	if s.prometheusServer != nil {
		s.logger.Info(ctx, "Starting Prometheus HTTP", "address", s.prometheusServer.Addr)
		go func() {
			s.errCh <- s.prometheusServer.ListenAndServe()
		}()
	}
}

func (s *server) shutdownServers(ctx context.Context, timeout time.Duration) {
	tctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(3)
	shutdownHTTPServer(tctx, s.httpServer, &wg, s.logger)
	shutdownHTTPServer(tctx, s.httpsServer, &wg, s.logger)
	shutdownHTTPServer(tctx, s.prometheusServer, &wg, s.logger)
	wg.Wait()
}

func (s *server) close() {
	if s.httpServer != nil {
		s.httpServer.Close()
	}
	if s.httpsServer != nil {
		s.httpsServer.Close()
	}
	if s.prometheusServer != nil {
		s.prometheusServer.Close()
	}

	s.clusterSvc.Close()
	s.repairSvc.Close()
	s.schedSvc.Close() // must be closed last

	s.session.Close()
}

func shutdownHTTPServer(ctx context.Context, s *http.Server, wg *sync.WaitGroup, l log.Logger) {
	if s == nil {
		wg.Done()
		return
	}

	go func() {
		defer wg.Done()
		if err := s.Shutdown(ctx); err != nil {
			l.Info(ctx, "Closing HTTP(S) server failed", "addr", s.Addr, "error", err)
		} else {
			l.Info(ctx, "Closing HTTP(S) server done", "addr", s.Addr)
		}
	}()
}
