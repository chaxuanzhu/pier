// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/meshplus/bitxhub-model/pb"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CommitCallback mocks base method.
func (m *MockClient) CommitCallback(ibtp *pb.IBTP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitCallback", ibtp)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitCallback indicates an expected call of CommitCallback.
func (mr *MockClientMockRecorder) CommitCallback(ibtp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitCallback", reflect.TypeOf((*MockClient)(nil).CommitCallback), ibtp)
}

// GetCallbackMeta mocks base method.
func (m *MockClient) GetCallbackMeta() (map[string]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallbackMeta")
	ret0, _ := ret[0].(map[string]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallbackMeta indicates an expected call of GetCallbackMeta.
func (mr *MockClientMockRecorder) GetCallbackMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallbackMeta", reflect.TypeOf((*MockClient)(nil).GetCallbackMeta))
}

// GetIBTP mocks base method.
func (m *MockClient) GetIBTP() chan *pb.IBTP {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIBTP")
	ret0, _ := ret[0].(chan *pb.IBTP)
	return ret0
}

// GetIBTP indicates an expected call of GetIBTP.
func (mr *MockClientMockRecorder) GetIBTP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIBTP", reflect.TypeOf((*MockClient)(nil).GetIBTP))
}

// GetInMessage mocks base method.
func (m *MockClient) GetInMessage(from string, idx uint64) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInMessage", from, idx)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInMessage indicates an expected call of GetInMessage.
func (mr *MockClientMockRecorder) GetInMessage(from, idx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInMessage", reflect.TypeOf((*MockClient)(nil).GetInMessage), from, idx)
}

// GetInMeta mocks base method.
func (m *MockClient) GetInMeta() (map[string]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInMeta")
	ret0, _ := ret[0].(map[string]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInMeta indicates an expected call of GetInMeta.
func (mr *MockClientMockRecorder) GetInMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInMeta", reflect.TypeOf((*MockClient)(nil).GetInMeta))
}

// GetLockEvent mocks base method.
func (m *MockClient) GetLockEvent() <-chan *pb.LockEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLockEvent")
	ret0, _ := ret[0].(<-chan *pb.LockEvent)
	return ret0
}

// GetLockEvent indicates an expected call of GetLockEvent.
func (mr *MockClientMockRecorder) GetLockEvent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLockEvent", reflect.TypeOf((*MockClient)(nil).GetLockEvent))
}

// GetOutMessage mocks base method.
func (m *MockClient) GetOutMessage(to string, idx uint64) (*pb.IBTP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutMessage", to, idx)
	ret0, _ := ret[0].(*pb.IBTP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutMessage indicates an expected call of GetOutMessage.
func (mr *MockClientMockRecorder) GetOutMessage(to, idx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutMessage", reflect.TypeOf((*MockClient)(nil).GetOutMessage), to, idx)
}

// GetOutMeta mocks base method.
func (m *MockClient) GetOutMeta() (map[string]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutMeta")
	ret0, _ := ret[0].(map[string]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutMeta indicates an expected call of GetOutMeta.
func (mr *MockClientMockRecorder) GetOutMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutMeta", reflect.TypeOf((*MockClient)(nil).GetOutMeta))
}

// GetReceipt mocks base method.
func (m *MockClient) GetReceipt(ibtp *pb.IBTP) (*pb.IBTP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipt", ibtp)
	ret0, _ := ret[0].(*pb.IBTP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipt indicates an expected call of GetReceipt.
func (mr *MockClientMockRecorder) GetReceipt(ibtp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockClient)(nil).GetReceipt), ibtp)
}

// GetUpdateMeta mocks base method.
func (m *MockClient) GetUpdateMeta() <-chan *pb.UpdateMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateMeta")
	ret0, _ := ret[0].(<-chan *pb.UpdateMeta)
	return ret0
}

// GetUpdateMeta indicates an expected call of GetUpdateMeta.
func (mr *MockClientMockRecorder) GetUpdateMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateMeta", reflect.TypeOf((*MockClient)(nil).GetUpdateMeta))
}

// IncreaseInMeta mocks base method.
func (m *MockClient) IncreaseInMeta(ibtp *pb.IBTP) (*pb.IBTP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseInMeta", ibtp)
	ret0, _ := ret[0].(*pb.IBTP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncreaseInMeta indicates an expected call of IncreaseInMeta.
func (mr *MockClientMockRecorder) IncreaseInMeta(ibtp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseInMeta", reflect.TypeOf((*MockClient)(nil).IncreaseInMeta), ibtp)
}

// Initialize mocks base method.
func (m *MockClient) Initialize(configPath, pierID string, extra []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", configPath, pierID, extra)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockClientMockRecorder) Initialize(configPath, pierID, extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockClient)(nil).Initialize), configPath, pierID, extra)
}

// Name mocks base method.
func (m *MockClient) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockClientMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockClient)(nil).Name))
}

// QueryAppchainIndex mocks base method.
func (m *MockClient) QueryAppchainIndex() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAppchainIndex")
	ret0, _ := ret[0].(int64)
	return ret0
}

// QueryAppchainIndex indicates an expected call of QueryAppchainIndex.
func (mr *MockClientMockRecorder) QueryAppchainIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAppchainIndex", reflect.TypeOf((*MockClient)(nil).QueryAppchainIndex))
}

// QueryFilterLockStart mocks base method.
func (m *MockClient) QueryFilterLockStart(appchainIndex int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFilterLockStart", appchainIndex)
	ret0, _ := ret[0].(int64)
	return ret0
}

// QueryFilterLockStart indicates an expected call of QueryFilterLockStart.
func (mr *MockClientMockRecorder) QueryFilterLockStart(appchainIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFilterLockStart", reflect.TypeOf((*MockClient)(nil).QueryFilterLockStart), appchainIndex)
}

// QueryLockEventByIndex mocks base method.
func (m *MockClient) QueryLockEventByIndex(index int64) *pb.LockEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLockEventByIndex", index)
	ret0, _ := ret[0].(*pb.LockEvent)
	return ret0
}

// QueryLockEventByIndex indicates an expected call of QueryLockEventByIndex.
func (mr *MockClientMockRecorder) QueryLockEventByIndex(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLockEventByIndex", reflect.TypeOf((*MockClient)(nil).QueryLockEventByIndex), index)
}

// QueryRelayIndex mocks base method.
func (m *MockClient) QueryRelayIndex() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRelayIndex")
	ret0, _ := ret[0].(int64)
	return ret0
}

// QueryRelayIndex indicates an expected call of QueryRelayIndex.
func (mr *MockClientMockRecorder) QueryRelayIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRelayIndex", reflect.TypeOf((*MockClient)(nil).QueryRelayIndex))
}

// RollbackIBTP mocks base method.
func (m *MockClient) RollbackIBTP(ibtp *pb.IBTP, isSrcChain bool) (*pb.RollbackIBTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackIBTP", ibtp, isSrcChain)
	ret0, _ := ret[0].(*pb.RollbackIBTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RollbackIBTP indicates an expected call of RollbackIBTP.
func (mr *MockClientMockRecorder) RollbackIBTP(ibtp, isSrcChain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackIBTP", reflect.TypeOf((*MockClient)(nil).RollbackIBTP), ibtp, isSrcChain)
}

// Start mocks base method.
func (m *MockClient) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockClientMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockClient)(nil).Start))
}

// Stop mocks base method.
func (m *MockClient) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockClientMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockClient)(nil).Stop))
}

// SubmitIBTP mocks base method.
func (m *MockClient) SubmitIBTP(arg0 *pb.IBTP) (*pb.SubmitIBTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitIBTP", arg0)
	ret0, _ := ret[0].(*pb.SubmitIBTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitIBTP indicates an expected call of SubmitIBTP.
func (mr *MockClientMockRecorder) SubmitIBTP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitIBTP", reflect.TypeOf((*MockClient)(nil).SubmitIBTP), arg0)
}

// Type mocks base method.
func (m *MockClient) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockClientMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockClient)(nil).Type))
}

// Unescrow mocks base method.
func (m *MockClient) Unescrow(unlock *pb.UnLock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unescrow", unlock)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unescrow indicates an expected call of Unescrow.
func (mr *MockClientMockRecorder) Unescrow(unlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unescrow", reflect.TypeOf((*MockClient)(nil).Unescrow), unlock)
}
