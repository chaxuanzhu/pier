// Code generated by MockGen. DO NOT EDIT.
// Source: peermgr.go

// Package mock_peermgr is a generated GoMock package.
package mock_peermgr

import (
	gomock "github.com/golang/mock/gomock"
	network "github.com/libp2p/go-libp2p-core/network"
	peer "github.com/libp2p/go-libp2p-core/peer"
	peermgr "github.com/meshplus/pier/internal/peermgr"
	peermgr0 "github.com/meshplus/pier/internal/peermgr/proto"
	reflect "reflect"
)

// MockPeerManager is a mock of PeerManager interface
type MockPeerManager struct {
	ctrl     *gomock.Controller
	recorder *MockPeerManagerMockRecorder
}

// MockPeerManagerMockRecorder is the mock recorder for MockPeerManager
type MockPeerManagerMockRecorder struct {
	mock *MockPeerManager
}

// NewMockPeerManager creates a new mock instance
func NewMockPeerManager(ctrl *gomock.Controller) *MockPeerManager {
	mock := &MockPeerManager{ctrl: ctrl}
	mock.recorder = &MockPeerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeerManager) EXPECT() *MockPeerManagerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockPeerManager) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockPeerManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockPeerManager)(nil).Start))
}

// Stop mocks base method
func (m *MockPeerManager) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockPeerManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockPeerManager)(nil).Stop))
}

// AsyncSend mocks base method
func (m *MockPeerManager) AsyncSend(arg0 string, arg1 *peermgr0.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsyncSend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AsyncSend indicates an expected call of AsyncSend
func (mr *MockPeerManagerMockRecorder) AsyncSend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsyncSend", reflect.TypeOf((*MockPeerManager)(nil).AsyncSend), arg0, arg1)
}

// SendWithStream mocks base method
func (m *MockPeerManager) SendWithStream(arg0 network.Stream, arg1 *peermgr0.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendWithStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendWithStream indicates an expected call of SendWithStream
func (mr *MockPeerManagerMockRecorder) SendWithStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendWithStream", reflect.TypeOf((*MockPeerManager)(nil).SendWithStream), arg0, arg1)
}

// Send mocks base method
func (m *MockPeerManager) Send(arg0 string, arg1 *peermgr0.Message) (*peermgr0.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*peermgr0.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockPeerManagerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockPeerManager)(nil).Send), arg0, arg1)
}

// Peers mocks base method
func (m *MockPeerManager) Peers() map[string]*peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peers")
	ret0, _ := ret[0].(map[string]*peer.AddrInfo)
	return ret0
}

// Peers indicates an expected call of Peers
func (mr *MockPeerManagerMockRecorder) Peers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockPeerManager)(nil).Peers))
}

// RegisterMsgHandler mocks base method
func (m *MockPeerManager) RegisterMsgHandler(arg0 peermgr0.Message_Type, arg1 peermgr.MessageHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterMsgHandler", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterMsgHandler indicates an expected call of RegisterMsgHandler
func (mr *MockPeerManagerMockRecorder) RegisterMsgHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMsgHandler", reflect.TypeOf((*MockPeerManager)(nil).RegisterMsgHandler), arg0, arg1)
}

// RegisterConnectHandler mocks base method
func (m *MockPeerManager) RegisterConnectHandler(arg0 peermgr.ConnectHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterConnectHandler", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterConnectHandler indicates an expected call of RegisterConnectHandler
func (mr *MockPeerManagerMockRecorder) RegisterConnectHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterConnectHandler", reflect.TypeOf((*MockPeerManager)(nil).RegisterConnectHandler), arg0)
}
