// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.enterprise.gloo.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/fed.enterprise.gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockFederatedAuthConfigEventHandler is a mock of FederatedAuthConfigEventHandler interface.
type MockFederatedAuthConfigEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigEventHandlerMockRecorder
}

// MockFederatedAuthConfigEventHandlerMockRecorder is the mock recorder for MockFederatedAuthConfigEventHandler.
type MockFederatedAuthConfigEventHandlerMockRecorder struct {
	mock *MockFederatedAuthConfigEventHandler
}

// NewMockFederatedAuthConfigEventHandler creates a new mock instance.
func NewMockFederatedAuthConfigEventHandler(ctrl *gomock.Controller) *MockFederatedAuthConfigEventHandler {
	mock := &MockFederatedAuthConfigEventHandler{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedAuthConfigEventHandler) EXPECT() *MockFederatedAuthConfigEventHandlerMockRecorder {
	return m.recorder
}

// CreateFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigEventHandler) CreateFederatedAuthConfig(obj *v1.FederatedAuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFederatedAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedAuthConfig indicates an expected call of CreateFederatedAuthConfig.
func (mr *MockFederatedAuthConfigEventHandlerMockRecorder) CreateFederatedAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigEventHandler)(nil).CreateFederatedAuthConfig), obj)
}

// DeleteFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigEventHandler) DeleteFederatedAuthConfig(obj *v1.FederatedAuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFederatedAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedAuthConfig indicates an expected call of DeleteFederatedAuthConfig.
func (mr *MockFederatedAuthConfigEventHandlerMockRecorder) DeleteFederatedAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigEventHandler)(nil).DeleteFederatedAuthConfig), obj)
}

// GenericFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigEventHandler) GenericFederatedAuthConfig(obj *v1.FederatedAuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericFederatedAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericFederatedAuthConfig indicates an expected call of GenericFederatedAuthConfig.
func (mr *MockFederatedAuthConfigEventHandlerMockRecorder) GenericFederatedAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigEventHandler)(nil).GenericFederatedAuthConfig), obj)
}

// UpdateFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigEventHandler) UpdateFederatedAuthConfig(old, new *v1.FederatedAuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFederatedAuthConfig", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedAuthConfig indicates an expected call of UpdateFederatedAuthConfig.
func (mr *MockFederatedAuthConfigEventHandlerMockRecorder) UpdateFederatedAuthConfig(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigEventHandler)(nil).UpdateFederatedAuthConfig), old, new)
}

// MockFederatedAuthConfigEventWatcher is a mock of FederatedAuthConfigEventWatcher interface.
type MockFederatedAuthConfigEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigEventWatcherMockRecorder
}

// MockFederatedAuthConfigEventWatcherMockRecorder is the mock recorder for MockFederatedAuthConfigEventWatcher.
type MockFederatedAuthConfigEventWatcherMockRecorder struct {
	mock *MockFederatedAuthConfigEventWatcher
}

// NewMockFederatedAuthConfigEventWatcher creates a new mock instance.
func NewMockFederatedAuthConfigEventWatcher(ctrl *gomock.Controller) *MockFederatedAuthConfigEventWatcher {
	mock := &MockFederatedAuthConfigEventWatcher{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedAuthConfigEventWatcher) EXPECT() *MockFederatedAuthConfigEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockFederatedAuthConfigEventWatcher) AddEventHandler(ctx context.Context, h controller.FederatedAuthConfigEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockFederatedAuthConfigEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockFederatedAuthConfigEventWatcher)(nil).AddEventHandler), varargs...)
}
