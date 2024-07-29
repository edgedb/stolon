// Code generated by MockGen. DO NOT EDIT.
// Source: postgresql.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	postgresql "github.com/sorintlab/stolon/internal/postgresql"
)

// MockPGManager is a mock of PGManager interface.
type MockPGManager struct {
	ctrl     *gomock.Controller
	recorder *MockPGManagerMockRecorder
}

// MockPGManagerMockRecorder is the mock recorder for MockPGManager.
type MockPGManagerMockRecorder struct {
	mock *MockPGManager
}

// NewMockPGManager creates a new mock instance.
func NewMockPGManager(ctrl *gomock.Controller) *MockPGManager {
	mock := &MockPGManager{ctrl: ctrl}
	mock.recorder = &MockPGManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPGManager) EXPECT() *MockPGManagerMockRecorder {
	return m.recorder
}

// GetTimelinesHistory mocks base method.
func (m *MockPGManager) GetTimelinesHistory(timeline uint64) ([]*postgresql.TimelineHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimelinesHistory", timeline)
	ret0, _ := ret[0].([]*postgresql.TimelineHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimelinesHistory indicates an expected call of GetTimelinesHistory.
func (mr *MockPGManagerMockRecorder) GetTimelinesHistory(timeline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimelinesHistory", reflect.TypeOf((*MockPGManager)(nil).GetTimelinesHistory), timeline)
}
