// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	cluster "github.com/sorintlab/stolon/internal/cluster"
	store "github.com/sorintlab/stolon/internal/store"
	reflect "reflect"
	time "time"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AtomicPutClusterData mocks base method
func (m *MockStore) AtomicPutClusterData(ctx context.Context, cd *cluster.ClusterData, previous *store.KVPair) (*store.KVPair, error) {
	ret := m.ctrl.Call(m, "AtomicPutClusterData", ctx, cd, previous)
	ret0, _ := ret[0].(*store.KVPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtomicPutClusterData indicates an expected call of AtomicPutClusterData
func (mr *MockStoreMockRecorder) AtomicPutClusterData(ctx, cd, previous interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtomicPutClusterData", reflect.TypeOf((*MockStore)(nil).AtomicPutClusterData), ctx, cd, previous)
}

// PutClusterData mocks base method
func (m *MockStore) PutClusterData(ctx context.Context, cd *cluster.ClusterData) error {
	ret := m.ctrl.Call(m, "PutClusterData", ctx, cd)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutClusterData indicates an expected call of PutClusterData
func (mr *MockStoreMockRecorder) PutClusterData(ctx, cd interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutClusterData", reflect.TypeOf((*MockStore)(nil).PutClusterData), ctx, cd)
}

// GetClusterData mocks base method
func (m *MockStore) GetClusterData(ctx context.Context) (*cluster.ClusterData, *store.KVPair, error) {
	ret := m.ctrl.Call(m, "GetClusterData", ctx)
	ret0, _ := ret[0].(*cluster.ClusterData)
	ret1, _ := ret[1].(*store.KVPair)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetClusterData indicates an expected call of GetClusterData
func (mr *MockStoreMockRecorder) GetClusterData(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterData", reflect.TypeOf((*MockStore)(nil).GetClusterData), ctx)
}

// SetKeeperInfo mocks base method
func (m *MockStore) SetKeeperInfo(ctx context.Context, id string, ms *cluster.KeeperInfo, ttl time.Duration) error {
	ret := m.ctrl.Call(m, "SetKeeperInfo", ctx, id, ms, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetKeeperInfo indicates an expected call of SetKeeperInfo
func (mr *MockStoreMockRecorder) SetKeeperInfo(ctx, id, ms, ttl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKeeperInfo", reflect.TypeOf((*MockStore)(nil).SetKeeperInfo), ctx, id, ms, ttl)
}

// GetKeepersInfo mocks base method
func (m *MockStore) GetKeepersInfo(ctx context.Context, blocking bool) (cluster.KeepersInfo, error) {
	ret := m.ctrl.Call(m, "GetKeepersInfo", ctx, blocking)
	ret0, _ := ret[0].(cluster.KeepersInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeepersInfo indicates an expected call of GetKeepersInfo
func (mr *MockStoreMockRecorder) GetKeepersInfo(ctx interface{}, blocking bool) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeepersInfo", reflect.TypeOf((*MockStore)(nil).GetKeepersInfo), ctx, blocking)
}

// SetSentinelInfo mocks base method
func (m *MockStore) SetSentinelInfo(ctx context.Context, si *cluster.SentinelInfo, ttl time.Duration) error {
	ret := m.ctrl.Call(m, "SetSentinelInfo", ctx, si, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSentinelInfo indicates an expected call of SetSentinelInfo
func (mr *MockStoreMockRecorder) SetSentinelInfo(ctx, si, ttl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSentinelInfo", reflect.TypeOf((*MockStore)(nil).SetSentinelInfo), ctx, si, ttl)
}

// GetSentinelsInfo mocks base method
func (m *MockStore) GetSentinelsInfo(ctx context.Context) (cluster.SentinelsInfo, error) {
	ret := m.ctrl.Call(m, "GetSentinelsInfo", ctx)
	ret0, _ := ret[0].(cluster.SentinelsInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSentinelsInfo indicates an expected call of GetSentinelsInfo
func (mr *MockStoreMockRecorder) GetSentinelsInfo(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSentinelsInfo", reflect.TypeOf((*MockStore)(nil).GetSentinelsInfo), ctx)
}

// SetProxyInfo mocks base method
func (m *MockStore) SetProxyInfo(ctx context.Context, pi *cluster.ProxyInfo, ttl time.Duration) error {
	ret := m.ctrl.Call(m, "SetProxyInfo", ctx, pi, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProxyInfo indicates an expected call of SetProxyInfo
func (mr *MockStoreMockRecorder) SetProxyInfo(ctx, pi, ttl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProxyInfo", reflect.TypeOf((*MockStore)(nil).SetProxyInfo), ctx, pi, ttl)
}

// GetProxiesInfo mocks base method
func (m *MockStore) GetProxiesInfo(ctx context.Context) (cluster.ProxiesInfo, error) {
	ret := m.ctrl.Call(m, "GetProxiesInfo", ctx)
	ret0, _ := ret[0].(cluster.ProxiesInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProxiesInfo indicates an expected call of GetProxiesInfo
func (mr *MockStoreMockRecorder) GetProxiesInfo(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProxiesInfo", reflect.TypeOf((*MockStore)(nil).GetProxiesInfo), ctx)
}

// MockElection is a mock of Election interface
type MockElection struct {
	ctrl     *gomock.Controller
	recorder *MockElectionMockRecorder
}

// MockElectionMockRecorder is the mock recorder for MockElection
type MockElectionMockRecorder struct {
	mock *MockElection
}

// NewMockElection creates a new mock instance
func NewMockElection(ctrl *gomock.Controller) *MockElection {
	mock := &MockElection{ctrl: ctrl}
	mock.recorder = &MockElectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockElection) EXPECT() *MockElectionMockRecorder {
	return m.recorder
}

// RunForElection mocks base method
func (m *MockElection) RunForElection() (<-chan bool, <-chan error) {
	ret := m.ctrl.Call(m, "RunForElection")
	ret0, _ := ret[0].(<-chan bool)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// RunForElection indicates an expected call of RunForElection
func (mr *MockElectionMockRecorder) RunForElection() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunForElection", reflect.TypeOf((*MockElection)(nil).RunForElection))
}

// Leader mocks base method
func (m *MockElection) Leader() (string, error) {
	ret := m.ctrl.Call(m, "Leader")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leader indicates an expected call of Leader
func (mr *MockElectionMockRecorder) Leader() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leader", reflect.TypeOf((*MockElection)(nil).Leader))
}

// Stop mocks base method
func (m *MockElection) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockElectionMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockElection)(nil).Stop))
}
