// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"

	v2alpha1 "github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2alpha1"
	controller "github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2alpha1/controller"
)

// MockWaypointLifecycleManagerEventHandler is a mock of WaypointLifecycleManagerEventHandler interface.
type MockWaypointLifecycleManagerEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerEventHandlerMockRecorder
}

// MockWaypointLifecycleManagerEventHandlerMockRecorder is the mock recorder for MockWaypointLifecycleManagerEventHandler.
type MockWaypointLifecycleManagerEventHandlerMockRecorder struct {
	mock *MockWaypointLifecycleManagerEventHandler
}

// NewMockWaypointLifecycleManagerEventHandler creates a new mock instance.
func NewMockWaypointLifecycleManagerEventHandler(ctrl *gomock.Controller) *MockWaypointLifecycleManagerEventHandler {
	mock := &MockWaypointLifecycleManagerEventHandler{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerEventHandler) EXPECT() *MockWaypointLifecycleManagerEventHandlerMockRecorder {
	return m.recorder
}

// CreateWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerEventHandler) CreateWaypointLifecycleManager(obj *v2alpha1.WaypointLifecycleManager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWaypointLifecycleManager", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWaypointLifecycleManager indicates an expected call of CreateWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerEventHandlerMockRecorder) CreateWaypointLifecycleManager(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerEventHandler)(nil).CreateWaypointLifecycleManager), obj)
}

// DeleteWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerEventHandler) DeleteWaypointLifecycleManager(obj *v2alpha1.WaypointLifecycleManager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWaypointLifecycleManager", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWaypointLifecycleManager indicates an expected call of DeleteWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerEventHandlerMockRecorder) DeleteWaypointLifecycleManager(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerEventHandler)(nil).DeleteWaypointLifecycleManager), obj)
}

// GenericWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerEventHandler) GenericWaypointLifecycleManager(obj *v2alpha1.WaypointLifecycleManager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericWaypointLifecycleManager", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericWaypointLifecycleManager indicates an expected call of GenericWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerEventHandlerMockRecorder) GenericWaypointLifecycleManager(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerEventHandler)(nil).GenericWaypointLifecycleManager), obj)
}

// UpdateWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerEventHandler) UpdateWaypointLifecycleManager(old, new *v2alpha1.WaypointLifecycleManager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWaypointLifecycleManager", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWaypointLifecycleManager indicates an expected call of UpdateWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerEventHandlerMockRecorder) UpdateWaypointLifecycleManager(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerEventHandler)(nil).UpdateWaypointLifecycleManager), old, new)
}

// MockWaypointLifecycleManagerEventWatcher is a mock of WaypointLifecycleManagerEventWatcher interface.
type MockWaypointLifecycleManagerEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerEventWatcherMockRecorder
}

// MockWaypointLifecycleManagerEventWatcherMockRecorder is the mock recorder for MockWaypointLifecycleManagerEventWatcher.
type MockWaypointLifecycleManagerEventWatcherMockRecorder struct {
	mock *MockWaypointLifecycleManagerEventWatcher
}

// NewMockWaypointLifecycleManagerEventWatcher creates a new mock instance.
func NewMockWaypointLifecycleManagerEventWatcher(ctrl *gomock.Controller) *MockWaypointLifecycleManagerEventWatcher {
	mock := &MockWaypointLifecycleManagerEventWatcher{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerEventWatcher) EXPECT() *MockWaypointLifecycleManagerEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockWaypointLifecycleManagerEventWatcher) AddEventHandler(ctx context.Context, h controller.WaypointLifecycleManagerEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockWaypointLifecycleManagerEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockWaypointLifecycleManagerEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockInsightsConfigEventHandler is a mock of InsightsConfigEventHandler interface.
type MockInsightsConfigEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockInsightsConfigEventHandlerMockRecorder
}

// MockInsightsConfigEventHandlerMockRecorder is the mock recorder for MockInsightsConfigEventHandler.
type MockInsightsConfigEventHandlerMockRecorder struct {
	mock *MockInsightsConfigEventHandler
}

// NewMockInsightsConfigEventHandler creates a new mock instance.
func NewMockInsightsConfigEventHandler(ctrl *gomock.Controller) *MockInsightsConfigEventHandler {
	mock := &MockInsightsConfigEventHandler{ctrl: ctrl}
	mock.recorder = &MockInsightsConfigEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInsightsConfigEventHandler) EXPECT() *MockInsightsConfigEventHandlerMockRecorder {
	return m.recorder
}

// CreateInsightsConfig mocks base method.
func (m *MockInsightsConfigEventHandler) CreateInsightsConfig(obj *v2alpha1.InsightsConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInsightsConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInsightsConfig indicates an expected call of CreateInsightsConfig.
func (mr *MockInsightsConfigEventHandlerMockRecorder) CreateInsightsConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInsightsConfig", reflect.TypeOf((*MockInsightsConfigEventHandler)(nil).CreateInsightsConfig), obj)
}

// DeleteInsightsConfig mocks base method.
func (m *MockInsightsConfigEventHandler) DeleteInsightsConfig(obj *v2alpha1.InsightsConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInsightsConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInsightsConfig indicates an expected call of DeleteInsightsConfig.
func (mr *MockInsightsConfigEventHandlerMockRecorder) DeleteInsightsConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInsightsConfig", reflect.TypeOf((*MockInsightsConfigEventHandler)(nil).DeleteInsightsConfig), obj)
}

// GenericInsightsConfig mocks base method.
func (m *MockInsightsConfigEventHandler) GenericInsightsConfig(obj *v2alpha1.InsightsConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericInsightsConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericInsightsConfig indicates an expected call of GenericInsightsConfig.
func (mr *MockInsightsConfigEventHandlerMockRecorder) GenericInsightsConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericInsightsConfig", reflect.TypeOf((*MockInsightsConfigEventHandler)(nil).GenericInsightsConfig), obj)
}

// UpdateInsightsConfig mocks base method.
func (m *MockInsightsConfigEventHandler) UpdateInsightsConfig(old, new *v2alpha1.InsightsConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInsightsConfig", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInsightsConfig indicates an expected call of UpdateInsightsConfig.
func (mr *MockInsightsConfigEventHandlerMockRecorder) UpdateInsightsConfig(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInsightsConfig", reflect.TypeOf((*MockInsightsConfigEventHandler)(nil).UpdateInsightsConfig), old, new)
}

// MockInsightsConfigEventWatcher is a mock of InsightsConfigEventWatcher interface.
type MockInsightsConfigEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockInsightsConfigEventWatcherMockRecorder
}

// MockInsightsConfigEventWatcherMockRecorder is the mock recorder for MockInsightsConfigEventWatcher.
type MockInsightsConfigEventWatcherMockRecorder struct {
	mock *MockInsightsConfigEventWatcher
}

// NewMockInsightsConfigEventWatcher creates a new mock instance.
func NewMockInsightsConfigEventWatcher(ctrl *gomock.Controller) *MockInsightsConfigEventWatcher {
	mock := &MockInsightsConfigEventWatcher{ctrl: ctrl}
	mock.recorder = &MockInsightsConfigEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInsightsConfigEventWatcher) EXPECT() *MockInsightsConfigEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockInsightsConfigEventWatcher) AddEventHandler(ctx context.Context, h controller.InsightsConfigEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockInsightsConfigEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockInsightsConfigEventWatcher)(nil).AddEventHandler), varargs...)
}
