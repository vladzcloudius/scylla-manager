// Copyright (C) 2017 ScyllaDB

package backup

import (
	"path"

	"github.com/scylladb/mermaid/uuid"
)

const dataDir = "/var/lib/scylla/data"

func keyspaceDir(keyspace string) string {
	return path.Join(dataDir, keyspace)
}

const manifest = "manifest.json"

func remoteManifestFile(clusterID, taskID, runID uuid.UUID, dc, nodeID, keyspace, table, version string) string {
	return path.Join(
		remoteBaseDir(clusterID, dc, nodeID, keyspace, table),
		"task",
		taskID.String(),
		"run",
		runID.String(),
		version,
		manifest,
	)
}

func remoteRunDir(clusterID, taskID, runID uuid.UUID, dc, nodeID, keyspace, table string) string {
	return path.Join(
		remoteBaseDir(clusterID, dc, nodeID, keyspace, table),
		"task",
		taskID.String(),
		"run",
		runID.String(),
	)
}

func remoteRunsDir(clusterID, taskID uuid.UUID, dc, nodeID, keyspace, table string) string {
	return path.Join(
		remoteBaseDir(clusterID, dc, nodeID, keyspace, table),
		"task",
		taskID.String(),
		"run",
	)
}

func remoteTasksDir(clusterID uuid.UUID, dc, nodeID, keyspace, table string) string {
	return path.Join(
		remoteBaseDir(clusterID, dc, nodeID, keyspace, table),
		"task",
	)
}

func remoteSSTableVersionDir(clusterID uuid.UUID, dc, nodeID, keyspace, table, version string) string {
	return path.Join(
		remoteBaseDir(clusterID, dc, nodeID, keyspace, table),
		"sst",
		version,
	)
}

func remoteSSTableDir(clusterID uuid.UUID, dc, nodeID, keyspace, table string) string {
	return path.Join(
		remoteBaseDir(clusterID, dc, nodeID, keyspace, table),
		"sst",
	)
}

func remoteBaseDir(clusterID uuid.UUID, dc, nodeID, keyspace, table string) string {
	return path.Join(
		"backup",
		"cluster",
		clusterID.String(),
		"dc",
		dc,
		"node",
		nodeID,
		"keyspace",
		keyspace,
		"table",
		table,
	)
}
