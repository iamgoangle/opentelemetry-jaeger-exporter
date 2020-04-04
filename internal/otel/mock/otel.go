// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel (interfaces: Tracer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	trace "go.opentelemetry.io/otel/api/trace"
	reflect "reflect"
)

// MockTracer is a mock of Tracer interface
type MockTracer struct {
	ctrl     *gomock.Controller
	recorder *MockTracerMockRecorder
}

// MockTracerMockRecorder is the mock recorder for MockTracer
type MockTracerMockRecorder struct {
	mock *MockTracer
}

// NewMockTracer creates a new mock instance
func NewMockTracer(ctrl *gomock.Controller) *MockTracer {
	mock := &MockTracer{ctrl: ctrl}
	mock.recorder = &MockTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTracer) EXPECT() *MockTracerMockRecorder {
	return m.recorder
}

// PrintSpanContext mocks base method
func (m *MockTracer) PrintSpanContext(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrintSpanContext", arg0)
}

// PrintSpanContext indicates an expected call of PrintSpanContext
func (mr *MockTracerMockRecorder) PrintSpanContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintSpanContext", reflect.TypeOf((*MockTracer)(nil).PrintSpanContext), arg0)
}

// SetIntAttribute mocks base method
func (m *MockTracer) SetIntAttribute(arg0 context.Context, arg1 string, arg2 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIntAttribute", arg0, arg1, arg2)
}

// SetIntAttribute indicates an expected call of SetIntAttribute
func (mr *MockTracerMockRecorder) SetIntAttribute(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIntAttribute", reflect.TypeOf((*MockTracer)(nil).SetIntAttribute), arg0, arg1, arg2)
}

// SetJaegerStatusCanceled mocks base method
func (m *MockTracer) SetJaegerStatusCanceled(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetJaegerStatusCanceled", arg0)
}

// SetJaegerStatusCanceled indicates an expected call of SetJaegerStatusCanceled
func (mr *MockTracerMockRecorder) SetJaegerStatusCanceled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJaegerStatusCanceled", reflect.TypeOf((*MockTracer)(nil).SetJaegerStatusCanceled), arg0)
}

// SetJaegerStatusInternal mocks base method
func (m *MockTracer) SetJaegerStatusInternal(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetJaegerStatusInternal", arg0)
}

// SetJaegerStatusInternal indicates an expected call of SetJaegerStatusInternal
func (mr *MockTracerMockRecorder) SetJaegerStatusInternal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJaegerStatusInternal", reflect.TypeOf((*MockTracer)(nil).SetJaegerStatusInternal), arg0)
}

// SetJaegerStatusOK mocks base method
func (m *MockTracer) SetJaegerStatusOK(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetJaegerStatusOK", arg0)
}

// SetJaegerStatusOK indicates an expected call of SetJaegerStatusOK
func (mr *MockTracerMockRecorder) SetJaegerStatusOK(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJaegerStatusOK", reflect.TypeOf((*MockTracer)(nil).SetJaegerStatusOK), arg0)
}

// SetStringAttribute mocks base method
func (m *MockTracer) SetStringAttribute(arg0 context.Context, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStringAttribute", arg0, arg1, arg2)
}

// SetStringAttribute indicates an expected call of SetStringAttribute
func (mr *MockTracerMockRecorder) SetStringAttribute(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStringAttribute", reflect.TypeOf((*MockTracer)(nil).SetStringAttribute), arg0, arg1, arg2)
}

// SpanID mocks base method
func (m *MockTracer) SpanID(arg0 context.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpanID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// SpanID indicates an expected call of SpanID
func (mr *MockTracerMockRecorder) SpanID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpanID", reflect.TypeOf((*MockTracer)(nil).SpanID), arg0)
}

// TraceID mocks base method
func (m *MockTracer) TraceID(arg0 context.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TraceID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// TraceID indicates an expected call of TraceID
func (mr *MockTracerMockRecorder) TraceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceID", reflect.TypeOf((*MockTracer)(nil).TraceID), arg0)
}

// TracerStart mocks base method
func (m *MockTracer) TracerStart(arg0 context.Context, arg1 string) (context.Context, trace.Span) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TracerStart", arg0, arg1)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(trace.Span)
	return ret0, ret1
}

// TracerStart indicates an expected call of TracerStart
func (mr *MockTracerMockRecorder) TracerStart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TracerStart", reflect.TypeOf((*MockTracer)(nil).TracerStart), arg0, arg1)
}
