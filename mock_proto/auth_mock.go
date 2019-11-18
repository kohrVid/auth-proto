// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kohrVid/auth/proto (interfaces: AuthenticationServiceClient)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/kohrVid/auth/proto"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAuthenticationServiceClient is a mock of AuthenticationServiceClient interface
type MockAuthenticationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationServiceClientMockRecorder
}

// MockAuthenticationServiceClientMockRecorder is the mock recorder for MockAuthenticationServiceClient
type MockAuthenticationServiceClientMockRecorder struct {
	mock *MockAuthenticationServiceClient
}

// NewMockAuthenticationServiceClient creates a new mock instance
func NewMockAuthenticationServiceClient(ctrl *gomock.Controller) *MockAuthenticationServiceClient {
	mock := &MockAuthenticationServiceClient{ctrl: ctrl}
	mock.recorder = &MockAuthenticationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationServiceClient) EXPECT() *MockAuthenticationServiceClientMockRecorder {
	return m.recorder
}

// CredentialCheck mocks base method
func (m *MockAuthenticationServiceClient) CredentialCheck(arg0 context.Context, arg1 *proto.AuthenticationRequest, arg2 ...grpc.CallOption) (*proto.AuthenticationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CredentialCheck", varargs...)
	ret0, _ := ret[0].(*proto.AuthenticationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredentialCheck indicates an expected call of CredentialCheck
func (mr *MockAuthenticationServiceClientMockRecorder) CredentialCheck(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialCheck", reflect.TypeOf((*MockAuthenticationServiceClient)(nil).CredentialCheck), varargs...)
}