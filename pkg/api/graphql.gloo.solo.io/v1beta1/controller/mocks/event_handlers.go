// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/solo-io/solo-apis/pkg/api/graphql.gloo.solo.io/v1beta1"
	controller "github.com/solo-io/solo-apis/pkg/api/graphql.gloo.solo.io/v1beta1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockGraphQLApiEventHandler is a mock of GraphQLApiEventHandler interface
type MockGraphQLApiEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLApiEventHandlerMockRecorder
}

// MockGraphQLApiEventHandlerMockRecorder is the mock recorder for MockGraphQLApiEventHandler
type MockGraphQLApiEventHandlerMockRecorder struct {
	mock *MockGraphQLApiEventHandler
}

// NewMockGraphQLApiEventHandler creates a new mock instance
func NewMockGraphQLApiEventHandler(ctrl *gomock.Controller) *MockGraphQLApiEventHandler {
	mock := &MockGraphQLApiEventHandler{ctrl: ctrl}
	mock.recorder = &MockGraphQLApiEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGraphQLApiEventHandler) EXPECT() *MockGraphQLApiEventHandlerMockRecorder {
	return m.recorder
}

// CreateGraphQLApi mocks base method
func (m *MockGraphQLApiEventHandler) CreateGraphQLApi(obj *v1beta1.GraphQLApi) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGraphQLApi", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGraphQLApi indicates an expected call of CreateGraphQLApi
func (mr *MockGraphQLApiEventHandlerMockRecorder) CreateGraphQLApi(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGraphQLApi", reflect.TypeOf((*MockGraphQLApiEventHandler)(nil).CreateGraphQLApi), obj)
}

// UpdateGraphQLApi mocks base method
func (m *MockGraphQLApiEventHandler) UpdateGraphQLApi(old, new *v1beta1.GraphQLApi) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGraphQLApi", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGraphQLApi indicates an expected call of UpdateGraphQLApi
func (mr *MockGraphQLApiEventHandlerMockRecorder) UpdateGraphQLApi(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGraphQLApi", reflect.TypeOf((*MockGraphQLApiEventHandler)(nil).UpdateGraphQLApi), old, new)
}

// DeleteGraphQLApi mocks base method
func (m *MockGraphQLApiEventHandler) DeleteGraphQLApi(obj *v1beta1.GraphQLApi) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGraphQLApi", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGraphQLApi indicates an expected call of DeleteGraphQLApi
func (mr *MockGraphQLApiEventHandlerMockRecorder) DeleteGraphQLApi(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGraphQLApi", reflect.TypeOf((*MockGraphQLApiEventHandler)(nil).DeleteGraphQLApi), obj)
}

// GenericGraphQLApi mocks base method
func (m *MockGraphQLApiEventHandler) GenericGraphQLApi(obj *v1beta1.GraphQLApi) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericGraphQLApi", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericGraphQLApi indicates an expected call of GenericGraphQLApi
func (mr *MockGraphQLApiEventHandlerMockRecorder) GenericGraphQLApi(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericGraphQLApi", reflect.TypeOf((*MockGraphQLApiEventHandler)(nil).GenericGraphQLApi), obj)
}

// MockGraphQLApiEventWatcher is a mock of GraphQLApiEventWatcher interface
type MockGraphQLApiEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLApiEventWatcherMockRecorder
}

// MockGraphQLApiEventWatcherMockRecorder is the mock recorder for MockGraphQLApiEventWatcher
type MockGraphQLApiEventWatcherMockRecorder struct {
	mock *MockGraphQLApiEventWatcher
}

// NewMockGraphQLApiEventWatcher creates a new mock instance
func NewMockGraphQLApiEventWatcher(ctrl *gomock.Controller) *MockGraphQLApiEventWatcher {
	mock := &MockGraphQLApiEventWatcher{ctrl: ctrl}
	mock.recorder = &MockGraphQLApiEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGraphQLApiEventWatcher) EXPECT() *MockGraphQLApiEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockGraphQLApiEventWatcher) AddEventHandler(ctx context.Context, h controller.GraphQLApiEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockGraphQLApiEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockGraphQLApiEventWatcher)(nil).AddEventHandler), varargs...)
}
