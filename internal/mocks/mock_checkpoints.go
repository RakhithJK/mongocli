// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/checkpoints.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	reflect "reflect"
)

// MockCheckpointsLister is a mock of CheckpointsLister interface
type MockCheckpointsLister struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointsListerMockRecorder
}

// MockCheckpointsListerMockRecorder is the mock recorder for MockCheckpointsLister
type MockCheckpointsListerMockRecorder struct {
	mock *MockCheckpointsLister
}

// NewMockCheckpointsLister creates a new mock instance
func NewMockCheckpointsLister(ctrl *gomock.Controller) *MockCheckpointsLister {
	mock := &MockCheckpointsLister{ctrl: ctrl}
	mock.recorder = &MockCheckpointsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCheckpointsLister) EXPECT() *MockCheckpointsListerMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockCheckpointsLister) List(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*mongodbatlas.Checkpoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Checkpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockCheckpointsListerMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCheckpointsLister)(nil).List), arg0, arg1, arg2)
}
