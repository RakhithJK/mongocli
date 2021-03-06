// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/host_disks.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	reflect "reflect"
)

// MockHostDisksLister is a mock of HostDisksLister interface
type MockHostDisksLister struct {
	ctrl     *gomock.Controller
	recorder *MockHostDisksListerMockRecorder
}

// MockHostDisksListerMockRecorder is the mock recorder for MockHostDisksLister
type MockHostDisksListerMockRecorder struct {
	mock *MockHostDisksLister
}

// NewMockHostDisksLister creates a new mock instance
func NewMockHostDisksLister(ctrl *gomock.Controller) *MockHostDisksLister {
	mock := &MockHostDisksLister{ctrl: ctrl}
	mock.recorder = &MockHostDisksListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHostDisksLister) EXPECT() *MockHostDisksListerMockRecorder {
	return m.recorder
}

// HostDisks mocks base method
func (m *MockHostDisksLister) HostDisks(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*mongodbatlas.ProcessDisksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostDisks", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.ProcessDisksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HostDisks indicates an expected call of HostDisks
func (mr *MockHostDisksListerMockRecorder) HostDisks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostDisks", reflect.TypeOf((*MockHostDisksLister)(nil).HostDisks), arg0, arg1, arg2)
}
