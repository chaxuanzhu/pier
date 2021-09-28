// Code generated by MockGen. DO NOT EDIT.
// Source: appchain_adapter.go

// Package mock_appchainAdapt is a generated GoMock package.
package mock_appchainAdapt

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/meshplus/bitxhub-model/pb"
)

// MockAppchainAdapter is a mock of AppchainAdapter interface.
type MockAppchainAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockAppchainAdapterMockRecorder
}

// MockAppchainAdapterMockRecorder is the mock recorder for MockAppchainAdapter.
type MockAppchainAdapterMockRecorder struct {
	mock *MockAppchainAdapter
}

// NewMockAppchainAdapter creates a new mock instance.
func NewMockAppchainAdapter(ctrl *gomock.Controller) *MockAppchainAdapter {
	mock := &MockAppchainAdapter{ctrl: ctrl}
	mock.recorder = &MockAppchainAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppchainAdapter) EXPECT() *MockAppchainAdapterMockRecorder {
	return m.recorder
}

// GetAppchainID mocks base method.
func (m *MockAppchainAdapter) GetAppchainID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppchainID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAppchainID indicates an expected call of GetAppchainID.
func (mr *MockAppchainAdapterMockRecorder) GetAppchainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppchainID", reflect.TypeOf((*MockAppchainAdapter)(nil).GetAppchainID))
}

// GetServiceIDList mocks base method.
func (m *MockAppchainAdapter) GetServiceIDList() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceIDList")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceIDList indicates an expected call of GetServiceIDList.
func (mr *MockAppchainAdapterMockRecorder) GetServiceIDList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceIDList", reflect.TypeOf((*MockAppchainAdapter)(nil).GetServiceIDList))
}

// MonitorIBTP mocks base method.
func (m *MockAppchainAdapter) MonitorIBTP() chan *pb.IBTP {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonitorIBTP")
	ret0, _ := ret[0].(chan *pb.IBTP)
	return ret0
}

// MonitorIBTP indicates an expected call of MonitorIBTP.
func (mr *MockAppchainAdapterMockRecorder) MonitorIBTP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitorIBTP", reflect.TypeOf((*MockAppchainAdapter)(nil).MonitorIBTP))
}

// MonitorUpdatedMeta mocks base method.
func (m *MockAppchainAdapter) MonitorUpdatedMeta() chan *[]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonitorUpdatedMeta")
	ret0, _ := ret[0].(chan *[]byte)
	return ret0
}

// MonitorUpdatedMeta indicates an expected call of MonitorUpdatedMeta.
func (mr *MockAppchainAdapterMockRecorder) MonitorUpdatedMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitorUpdatedMeta", reflect.TypeOf((*MockAppchainAdapter)(nil).MonitorUpdatedMeta))
}

// Name mocks base method.
func (m *MockAppchainAdapter) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAppchainAdapterMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppchainAdapter)(nil).Name))
}

// QueryIBTP mocks base method.
func (m *MockAppchainAdapter) QueryIBTP(id string, isReq bool) (*pb.IBTP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryIBTP", id, isReq)
	ret0, _ := ret[0].(*pb.IBTP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryIBTP indicates an expected call of QueryIBTP.
func (mr *MockAppchainAdapterMockRecorder) QueryIBTP(id, isReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryIBTP", reflect.TypeOf((*MockAppchainAdapter)(nil).QueryIBTP), id, isReq)
}

// QueryInterchain mocks base method.
func (m *MockAppchainAdapter) QueryInterchain(serviceID string) (*pb.Interchain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryInterchain", serviceID)
	ret0, _ := ret[0].(*pb.Interchain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryInterchain indicates an expected call of QueryInterchain.
func (mr *MockAppchainAdapterMockRecorder) QueryInterchain(serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryInterchain", reflect.TypeOf((*MockAppchainAdapter)(nil).QueryInterchain), serviceID)
}

// SendIBTP mocks base method.
func (m *MockAppchainAdapter) SendIBTP(ibtp *pb.IBTP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendIBTP", ibtp)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendIBTP indicates an expected call of SendIBTP.
func (mr *MockAppchainAdapterMockRecorder) SendIBTP(ibtp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendIBTP", reflect.TypeOf((*MockAppchainAdapter)(nil).SendIBTP), ibtp)
}

// SendUpdatedMeta mocks base method.
func (m *MockAppchainAdapter) SendUpdatedMeta(byte []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendUpdatedMeta", byte)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendUpdatedMeta indicates an expected call of SendUpdatedMeta.
func (mr *MockAppchainAdapterMockRecorder) SendUpdatedMeta(byte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUpdatedMeta", reflect.TypeOf((*MockAppchainAdapter)(nil).SendUpdatedMeta), byte)
}

// Start mocks base method.
func (m *MockAppchainAdapter) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockAppchainAdapterMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockAppchainAdapter)(nil).Start))
}

// Stop mocks base method.
func (m *MockAppchainAdapter) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockAppchainAdapterMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockAppchainAdapter)(nil).Stop))
}
