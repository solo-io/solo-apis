// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"

	v2alpha1 "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2alpha1"
	controller "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2alpha1/controller"
)

// MockExternalWorkloadEventHandler is a mock of ExternalWorkloadEventHandler interface.
type MockExternalWorkloadEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockExternalWorkloadEventHandlerMockRecorder
}

// MockExternalWorkloadEventHandlerMockRecorder is the mock recorder for MockExternalWorkloadEventHandler.
type MockExternalWorkloadEventHandlerMockRecorder struct {
	mock *MockExternalWorkloadEventHandler
}

// NewMockExternalWorkloadEventHandler creates a new mock instance.
func NewMockExternalWorkloadEventHandler(ctrl *gomock.Controller) *MockExternalWorkloadEventHandler {
	mock := &MockExternalWorkloadEventHandler{ctrl: ctrl}
	mock.recorder = &MockExternalWorkloadEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalWorkloadEventHandler) EXPECT() *MockExternalWorkloadEventHandlerMockRecorder {
	return m.recorder
}

// CreateExternalWorkload mocks base method.
func (m *MockExternalWorkloadEventHandler) CreateExternalWorkload(obj *v2alpha1.ExternalWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExternalWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExternalWorkload indicates an expected call of CreateExternalWorkload.
func (mr *MockExternalWorkloadEventHandlerMockRecorder) CreateExternalWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalWorkload", reflect.TypeOf((*MockExternalWorkloadEventHandler)(nil).CreateExternalWorkload), obj)
}

// DeleteExternalWorkload mocks base method.
func (m *MockExternalWorkloadEventHandler) DeleteExternalWorkload(obj *v2alpha1.ExternalWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExternalWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalWorkload indicates an expected call of DeleteExternalWorkload.
func (mr *MockExternalWorkloadEventHandlerMockRecorder) DeleteExternalWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalWorkload", reflect.TypeOf((*MockExternalWorkloadEventHandler)(nil).DeleteExternalWorkload), obj)
}

// GenericExternalWorkload mocks base method.
func (m *MockExternalWorkloadEventHandler) GenericExternalWorkload(obj *v2alpha1.ExternalWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericExternalWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericExternalWorkload indicates an expected call of GenericExternalWorkload.
func (mr *MockExternalWorkloadEventHandlerMockRecorder) GenericExternalWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericExternalWorkload", reflect.TypeOf((*MockExternalWorkloadEventHandler)(nil).GenericExternalWorkload), obj)
}

// UpdateExternalWorkload mocks base method.
func (m *MockExternalWorkloadEventHandler) UpdateExternalWorkload(old, new *v2alpha1.ExternalWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExternalWorkload", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalWorkload indicates an expected call of UpdateExternalWorkload.
func (mr *MockExternalWorkloadEventHandlerMockRecorder) UpdateExternalWorkload(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalWorkload", reflect.TypeOf((*MockExternalWorkloadEventHandler)(nil).UpdateExternalWorkload), old, new)
}

// MockExternalWorkloadEventWatcher is a mock of ExternalWorkloadEventWatcher interface.
type MockExternalWorkloadEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockExternalWorkloadEventWatcherMockRecorder
}

// MockExternalWorkloadEventWatcherMockRecorder is the mock recorder for MockExternalWorkloadEventWatcher.
type MockExternalWorkloadEventWatcherMockRecorder struct {
	mock *MockExternalWorkloadEventWatcher
}

// NewMockExternalWorkloadEventWatcher creates a new mock instance.
func NewMockExternalWorkloadEventWatcher(ctrl *gomock.Controller) *MockExternalWorkloadEventWatcher {
	mock := &MockExternalWorkloadEventWatcher{ctrl: ctrl}
	mock.recorder = &MockExternalWorkloadEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalWorkloadEventWatcher) EXPECT() *MockExternalWorkloadEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockExternalWorkloadEventWatcher) AddEventHandler(ctx context.Context, h controller.ExternalWorkloadEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockExternalWorkloadEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockExternalWorkloadEventWatcher)(nil).AddEventHandler), varargs...)
}