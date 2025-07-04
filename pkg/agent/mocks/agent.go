// Code generated by MockGen. DO NOT EDIT.
// Source: ricebean/pkg/agent (interfaces: Agent,AgentFactory)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	agent "ricebean/pkg/agent"
	protos "ricebean/pkg/protos"
	session "ricebean/pkg/session"
)

// MockAgent is a mock of Agent interface.
type MockAgent struct {
	ctrl     *gomock.Controller
	recorder *MockAgentMockRecorder
}

// MockAgentMockRecorder is the mock recorder for MockAgent.
type MockAgentMockRecorder struct {
	mock *MockAgent
}

// NewMockAgent creates a new mock instance.
func NewMockAgent(ctrl *gomock.Controller) *MockAgent {
	mock := &MockAgent{ctrl: ctrl}
	mock.recorder = &MockAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgent) EXPECT() *MockAgentMockRecorder {
	return m.recorder
}

// AnswerWithError mocks base method.
func (m *MockAgent) AnswerWithError(arg0 context.Context, arg1 uint, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AnswerWithError", arg0, arg1, arg2)
}

// AnswerWithError indicates an expected call of AnswerWithError.
func (mr *MockAgentMockRecorder) AnswerWithError(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnswerWithError", reflect.TypeOf((*MockAgent)(nil).AnswerWithError), arg0, arg1, arg2)
}

// Close mocks base method.
func (m *MockAgent) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAgentMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAgent)(nil).Close))
}

// GetSession mocks base method.
func (m *MockAgent) GetSession() session.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession")
	ret0, _ := ret[0].(session.Session)
	return ret0
}

// GetSession indicates an expected call of GetSession.
func (mr *MockAgentMockRecorder) GetSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockAgent)(nil).GetSession))
}

// GetStatus mocks base method.
func (m *MockAgent) GetStatus() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockAgentMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockAgent)(nil).GetStatus))
}

// Handle mocks base method.
func (m *MockAgent) Handle() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle")
}

// Handle indicates an expected call of Handle.
func (mr *MockAgentMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockAgent)(nil).Handle))
}

// IPVersion mocks base method.
func (m *MockAgent) IPVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IPVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// IPVersion indicates an expected call of IPVersion.
func (mr *MockAgentMockRecorder) IPVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IPVersion", reflect.TypeOf((*MockAgent)(nil).IPVersion))
}

// Kick mocks base method.
func (m *MockAgent) Kick(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kick", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Kick indicates an expected call of Kick.
func (mr *MockAgentMockRecorder) Kick(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kick", reflect.TypeOf((*MockAgent)(nil).Kick), arg0)
}

// Push mocks base method.
func (m *MockAgent) Push(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockAgentMockRecorder) Push(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockAgent)(nil).Push), arg0, arg1)
}

// RemoteAddr mocks base method.
func (m *MockAgent) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockAgentMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockAgent)(nil).RemoteAddr))
}

// ResponseMID mocks base method.
func (m *MockAgent) ResponseMID(arg0 context.Context, arg1 uint, arg2 interface{}, arg3 ...bool) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResponseMID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResponseMID indicates an expected call of ResponseMID.
func (mr *MockAgentMockRecorder) ResponseMID(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseMID", reflect.TypeOf((*MockAgent)(nil).ResponseMID), varargs...)
}

// SendHandshakeErrorResponse mocks base method.
func (m *MockAgent) SendHandshakeErrorResponse() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHandshakeErrorResponse")
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHandshakeErrorResponse indicates an expected call of SendHandshakeErrorResponse.
func (mr *MockAgentMockRecorder) SendHandshakeErrorResponse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHandshakeErrorResponse", reflect.TypeOf((*MockAgent)(nil).SendHandshakeErrorResponse))
}

// SendHandshakeResponse mocks base method.
func (m *MockAgent) SendHandshakeResponse() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHandshakeResponse")
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHandshakeResponse indicates an expected call of SendHandshakeResponse.
func (mr *MockAgentMockRecorder) SendHandshakeResponse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHandshakeResponse", reflect.TypeOf((*MockAgent)(nil).SendHandshakeResponse))
}

// SendRequest mocks base method.
func (m *MockAgent) SendRequest(arg0 context.Context, arg1, arg2 string, arg3 interface{}) (*protos.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRequest", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRequest indicates an expected call of SendRequest.
func (mr *MockAgentMockRecorder) SendRequest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRequest", reflect.TypeOf((*MockAgent)(nil).SendRequest), arg0, arg1, arg2, arg3)
}

// SetLastAt mocks base method.
func (m *MockAgent) SetLastAt() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastAt")
}

// SetLastAt indicates an expected call of SetLastAt.
func (mr *MockAgentMockRecorder) SetLastAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastAt", reflect.TypeOf((*MockAgent)(nil).SetLastAt))
}

// SetStatus mocks base method.
func (m *MockAgent) SetStatus(arg0 int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStatus", arg0)
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockAgentMockRecorder) SetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockAgent)(nil).SetStatus), arg0)
}

// String mocks base method.
func (m *MockAgent) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockAgentMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockAgent)(nil).String))
}

// MockAgentFactory is a mock of AgentFactory interface.
type MockAgentFactory struct {
	ctrl     *gomock.Controller
	recorder *MockAgentFactoryMockRecorder
}

// MockAgentFactoryMockRecorder is the mock recorder for MockAgentFactory.
type MockAgentFactoryMockRecorder struct {
	mock *MockAgentFactory
}

// NewMockAgentFactory creates a new mock instance.
func NewMockAgentFactory(ctrl *gomock.Controller) *MockAgentFactory {
	mock := &MockAgentFactory{ctrl: ctrl}
	mock.recorder = &MockAgentFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentFactory) EXPECT() *MockAgentFactoryMockRecorder {
	return m.recorder
}

// CreateAgent mocks base method.
func (m *MockAgentFactory) CreateAgent(arg0 net.Conn) agent.Agent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAgent", arg0)
	ret0, _ := ret[0].(agent.Agent)
	return ret0
}

// CreateAgent indicates an expected call of CreateAgent.
func (mr *MockAgentFactoryMockRecorder) CreateAgent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAgent", reflect.TypeOf((*MockAgentFactory)(nil).CreateAgent), arg0)
}
