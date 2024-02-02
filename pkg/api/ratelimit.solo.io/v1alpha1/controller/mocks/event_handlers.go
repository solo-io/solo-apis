// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
	controller "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockRateLimitConfigEventHandler is a mock of RateLimitConfigEventHandler interface.
type MockRateLimitConfigEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimitConfigEventHandlerMockRecorder
}

// MockRateLimitConfigEventHandlerMockRecorder is the mock recorder for MockRateLimitConfigEventHandler.
type MockRateLimitConfigEventHandlerMockRecorder struct {
	mock *MockRateLimitConfigEventHandler
}

// NewMockRateLimitConfigEventHandler creates a new mock instance.
func NewMockRateLimitConfigEventHandler(ctrl *gomock.Controller) *MockRateLimitConfigEventHandler {
	mock := &MockRateLimitConfigEventHandler{ctrl: ctrl}
	mock.recorder = &MockRateLimitConfigEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateLimitConfigEventHandler) EXPECT() *MockRateLimitConfigEventHandlerMockRecorder {
	return m.recorder
}

// CreateRateLimitConfig mocks base method.
func (m *MockRateLimitConfigEventHandler) CreateRateLimitConfig(obj *v1alpha1.RateLimitConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRateLimitConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRateLimitConfig indicates an expected call of CreateRateLimitConfig.
func (mr *MockRateLimitConfigEventHandlerMockRecorder) CreateRateLimitConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigEventHandler)(nil).CreateRateLimitConfig), obj)
}

// DeleteRateLimitConfig mocks base method.
func (m *MockRateLimitConfigEventHandler) DeleteRateLimitConfig(obj *v1alpha1.RateLimitConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRateLimitConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRateLimitConfig indicates an expected call of DeleteRateLimitConfig.
func (mr *MockRateLimitConfigEventHandlerMockRecorder) DeleteRateLimitConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigEventHandler)(nil).DeleteRateLimitConfig), obj)
}

// GenericRateLimitConfig mocks base method.
func (m *MockRateLimitConfigEventHandler) GenericRateLimitConfig(obj *v1alpha1.RateLimitConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericRateLimitConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericRateLimitConfig indicates an expected call of GenericRateLimitConfig.
func (mr *MockRateLimitConfigEventHandlerMockRecorder) GenericRateLimitConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigEventHandler)(nil).GenericRateLimitConfig), obj)
}

// UpdateRateLimitConfig mocks base method.
func (m *MockRateLimitConfigEventHandler) UpdateRateLimitConfig(old, new *v1alpha1.RateLimitConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRateLimitConfig", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRateLimitConfig indicates an expected call of UpdateRateLimitConfig.
func (mr *MockRateLimitConfigEventHandlerMockRecorder) UpdateRateLimitConfig(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigEventHandler)(nil).UpdateRateLimitConfig), old, new)
}

// MockRateLimitConfigEventWatcher is a mock of RateLimitConfigEventWatcher interface.
type MockRateLimitConfigEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimitConfigEventWatcherMockRecorder
}

// MockRateLimitConfigEventWatcherMockRecorder is the mock recorder for MockRateLimitConfigEventWatcher.
type MockRateLimitConfigEventWatcherMockRecorder struct {
	mock *MockRateLimitConfigEventWatcher
}

// NewMockRateLimitConfigEventWatcher creates a new mock instance.
func NewMockRateLimitConfigEventWatcher(ctrl *gomock.Controller) *MockRateLimitConfigEventWatcher {
	mock := &MockRateLimitConfigEventWatcher{ctrl: ctrl}
	mock.recorder = &MockRateLimitConfigEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateLimitConfigEventWatcher) EXPECT() *MockRateLimitConfigEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockRateLimitConfigEventWatcher) AddEventHandler(ctx context.Context, h controller.RateLimitConfigEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockRateLimitConfigEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockRateLimitConfigEventWatcher)(nil).AddEventHandler), varargs...)
}
