// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scylladb/mermaid/restapi (interfaces: SchedService)

// Package restapi is a generated GoMock package.
package restapi

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	scheduler "github.com/scylladb/mermaid/service/scheduler"
	uuid "github.com/scylladb/mermaid/uuid"
	reflect "reflect"
)

// MockSchedService is a mock of SchedService interface
type MockSchedService struct {
	ctrl     *gomock.Controller
	recorder *MockSchedServiceMockRecorder
}

// MockSchedServiceMockRecorder is the mock recorder for MockSchedService
type MockSchedServiceMockRecorder struct {
	mock *MockSchedService
}

// NewMockSchedService creates a new mock instance
func NewMockSchedService(ctrl *gomock.Controller) *MockSchedService {
	mock := &MockSchedService{ctrl: ctrl}
	mock.recorder = &MockSchedServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSchedService) EXPECT() *MockSchedServiceMockRecorder {
	return m.recorder
}

// DeleteTask mocks base method
func (m *MockSchedService) DeleteTask(arg0 context.Context, arg1 *scheduler.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask
func (mr *MockSchedServiceMockRecorder) DeleteTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockSchedService)(nil).DeleteTask), arg0, arg1)
}

// GetLastRun mocks base method
func (m *MockSchedService) GetLastRun(arg0 context.Context, arg1 *scheduler.Task, arg2 int) ([]*scheduler.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastRun", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*scheduler.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastRun indicates an expected call of GetLastRun
func (mr *MockSchedServiceMockRecorder) GetLastRun(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastRun", reflect.TypeOf((*MockSchedService)(nil).GetLastRun), arg0, arg1, arg2)
}

// GetRun mocks base method
func (m *MockSchedService) GetRun(arg0 context.Context, arg1 *scheduler.Task, arg2 uuid.UUID) (*scheduler.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRun", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scheduler.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRun indicates an expected call of GetRun
func (mr *MockSchedServiceMockRecorder) GetRun(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRun", reflect.TypeOf((*MockSchedService)(nil).GetRun), arg0, arg1, arg2)
}

// GetTask mocks base method
func (m *MockSchedService) GetTask(arg0 context.Context, arg1 uuid.UUID, arg2 scheduler.TaskType, arg3 string) (*scheduler.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scheduler.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask
func (mr *MockSchedServiceMockRecorder) GetTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockSchedService)(nil).GetTask), arg0, arg1, arg2, arg3)
}

// ListTasks mocks base method
func (m *MockSchedService) ListTasks(arg0 context.Context, arg1 uuid.UUID, arg2 scheduler.TaskType) ([]*scheduler.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*scheduler.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks
func (mr *MockSchedServiceMockRecorder) ListTasks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockSchedService)(nil).ListTasks), arg0, arg1, arg2)
}

// PutTask mocks base method
func (m *MockSchedService) PutTask(arg0 context.Context, arg1 *scheduler.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutTask indicates an expected call of PutTask
func (mr *MockSchedServiceMockRecorder) PutTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutTask", reflect.TypeOf((*MockSchedService)(nil).PutTask), arg0, arg1)
}

// PutTaskOnce mocks base method
func (m *MockSchedService) PutTaskOnce(arg0 context.Context, arg1 *scheduler.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutTaskOnce", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutTaskOnce indicates an expected call of PutTaskOnce
func (mr *MockSchedServiceMockRecorder) PutTaskOnce(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutTaskOnce", reflect.TypeOf((*MockSchedService)(nil).PutTaskOnce), arg0, arg1)
}

// StartTask mocks base method
func (m *MockSchedService) StartTask(arg0 context.Context, arg1 *scheduler.Task, arg2 ...scheduler.Opt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartTask indicates an expected call of StartTask
func (mr *MockSchedServiceMockRecorder) StartTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTask", reflect.TypeOf((*MockSchedService)(nil).StartTask), varargs...)
}

// StopTask mocks base method
func (m *MockSchedService) StopTask(arg0 context.Context, arg1 *scheduler.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopTask indicates an expected call of StopTask
func (mr *MockSchedServiceMockRecorder) StopTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopTask", reflect.TypeOf((*MockSchedService)(nil).StopTask), arg0, arg1)
}
