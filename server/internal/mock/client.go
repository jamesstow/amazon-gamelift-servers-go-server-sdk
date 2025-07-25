// Code generated by MockGen. DO NOT EDIT.
// Source: aws/amazon-gamelift-go-sdk/server/internal (interfaces: IWebSocketClient)

// Package mock is a generated GoMock package.
package mock

import (
	common "github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5/common"
	message "github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5/model/message"
	internal "github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5/server/internal"
	url "net/url"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIWebSocketClient is a mock of IWebSocketClient interface.
type MockIWebSocketClient struct {
	ctrl     *gomock.Controller
	recorder *MockIWebSocketClientMockRecorder
}

// MockIWebSocketClientMockRecorder is the mock recorder for MockIWebSocketClient.
type MockIWebSocketClientMockRecorder struct {
	mock *MockIWebSocketClient
}

// NewMockIWebSocketClient creates a new mock instance.
func NewMockIWebSocketClient(ctrl *gomock.Controller) *MockIWebSocketClient {
	mock := &MockIWebSocketClient{ctrl: ctrl}
	mock.recorder = &MockIWebSocketClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWebSocketClient) EXPECT() *MockIWebSocketClientMockRecorder {
	return m.recorder
}

// AddHandler mocks base method.
func (m *MockIWebSocketClient) AddHandler(arg0 message.MessageAction, arg1 func([]byte)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddHandler", arg0, arg1)
}

// AddHandler indicates an expected call of AddHandler.
func (mr *MockIWebSocketClientMockRecorder) AddHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandler", reflect.TypeOf((*MockIWebSocketClient)(nil).AddHandler), arg0, arg1)
}

// CancelRequest mocks base method.
func (m *MockIWebSocketClient) CancelRequest(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelRequest", arg0)
}

// CancelRequest indicates an expected call of CancelRequest.
func (mr *MockIWebSocketClientMockRecorder) CancelRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRequest", reflect.TypeOf((*MockIWebSocketClient)(nil).CancelRequest), arg0)
}

// Close mocks base method.
func (m *MockIWebSocketClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIWebSocketClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIWebSocketClient)(nil).Close))
}

// Connect mocks base method.
func (m *MockIWebSocketClient) Connect(arg0 *url.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockIWebSocketClientMockRecorder) Connect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockIWebSocketClient)(nil).Connect), arg0)
}

// SendRequest mocks base method.
func (m *MockIWebSocketClient) SendRequest(arg0 internal.MessageGetter, arg1 chan<- common.Outcome) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendRequest indicates an expected call of SendRequest.
func (mr *MockIWebSocketClientMockRecorder) SendRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRequest", reflect.TypeOf((*MockIWebSocketClient)(nil).SendRequest), arg0, arg1)
}
