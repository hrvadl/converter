// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GenesisEducationKyiv/software-engineering-school-4-0-hrvadl/mailer/internal/platform/mail/gomail (interfaces: ChainedSender)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_sender.go -package=mocks . ChainedSender
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	mailer "github.com/GenesisEducationKyiv/software-engineering-school-4-0-hrvadl/protos/gen/go/v1/mailer"
	gomock "go.uber.org/mock/gomock"
)

// MockChainedSender is a mock of ChainedSender interface.
type MockChainedSender struct {
	ctrl     *gomock.Controller
	recorder *MockChainedSenderMockRecorder
}

// MockChainedSenderMockRecorder is the mock recorder for MockChainedSender.
type MockChainedSenderMockRecorder struct {
	mock *MockChainedSender
}

// NewMockChainedSender creates a new mock instance.
func NewMockChainedSender(ctrl *gomock.Controller) *MockChainedSender {
	mock := &MockChainedSender{ctrl: ctrl}
	mock.recorder = &MockChainedSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainedSender) EXPECT() *MockChainedSenderMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockChainedSender) Send(arg0 context.Context, arg1 *mailer.Mail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockChainedSenderMockRecorder) Send(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChainedSender)(nil).Send), arg0, arg1)
}