// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hrvadl/converter/sub/internal/service/sender (interfaces: RateGetter)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_rategetter.go -package=mocks . RateGetter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRateGetter is a mock of RateGetter interface.
type MockRateGetter struct {
	ctrl     *gomock.Controller
	recorder *MockRateGetterMockRecorder
}

// MockRateGetterMockRecorder is the mock recorder for MockRateGetter.
type MockRateGetterMockRecorder struct {
	mock *MockRateGetter
}

// NewMockRateGetter creates a new mock instance.
func NewMockRateGetter(ctrl *gomock.Controller) *MockRateGetter {
	mock := &MockRateGetter{ctrl: ctrl}
	mock.recorder = &MockRateGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateGetter) EXPECT() *MockRateGetterMockRecorder {
	return m.recorder
}

// GetRate mocks base method.
func (m *MockRateGetter) GetRate(arg0 context.Context) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRate", arg0)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRate indicates an expected call of GetRate.
func (mr *MockRateGetterMockRecorder) GetRate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRate", reflect.TypeOf((*MockRateGetter)(nil).GetRate), arg0)
}
