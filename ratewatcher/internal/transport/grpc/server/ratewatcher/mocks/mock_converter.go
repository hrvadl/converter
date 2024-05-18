// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hrvadl/converter/ratewatcher/internal/transport/grpc/server/ratewatcher (interfaces: Converter)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_converter.go -package=mocks . Converter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockConverter is a mock of Converter interface.
type MockConverter struct {
	ctrl     *gomock.Controller
	recorder *MockConverterMockRecorder
}

// MockConverterMockRecorder is the mock recorder for MockConverter.
type MockConverterMockRecorder struct {
	mock *MockConverter
}

// NewMockConverter creates a new mock instance.
func NewMockConverter(ctrl *gomock.Controller) *MockConverter {
	mock := &MockConverter{ctrl: ctrl}
	mock.recorder = &MockConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConverter) EXPECT() *MockConverterMockRecorder {
	return m.recorder
}

// Convert mocks base method.
func (m *MockConverter) Convert(arg0 context.Context) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Convert", arg0)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Convert indicates an expected call of Convert.
func (mr *MockConverterMockRecorder) Convert(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Convert", reflect.TypeOf((*MockConverter)(nil).Convert), arg0)
}