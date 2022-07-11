// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2 "github.com/solo-io/solo-apis/pkg/api/resilience.policy.gloo.solo.io/v2"
	controller "github.com/solo-io/solo-apis/pkg/api/resilience.policy.gloo.solo.io/v2/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockFailoverPolicyEventHandler is a mock of FailoverPolicyEventHandler interface
type MockFailoverPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFailoverPolicyEventHandlerMockRecorder
}

// MockFailoverPolicyEventHandlerMockRecorder is the mock recorder for MockFailoverPolicyEventHandler
type MockFailoverPolicyEventHandlerMockRecorder struct {
	mock *MockFailoverPolicyEventHandler
}

// NewMockFailoverPolicyEventHandler creates a new mock instance
func NewMockFailoverPolicyEventHandler(ctrl *gomock.Controller) *MockFailoverPolicyEventHandler {
	mock := &MockFailoverPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockFailoverPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFailoverPolicyEventHandler) EXPECT() *MockFailoverPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateFailoverPolicy mocks base method
func (m *MockFailoverPolicyEventHandler) CreateFailoverPolicy(obj *v2.FailoverPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFailoverPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFailoverPolicy indicates an expected call of CreateFailoverPolicy
func (mr *MockFailoverPolicyEventHandlerMockRecorder) CreateFailoverPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFailoverPolicy", reflect.TypeOf((*MockFailoverPolicyEventHandler)(nil).CreateFailoverPolicy), obj)
}

// UpdateFailoverPolicy mocks base method
func (m *MockFailoverPolicyEventHandler) UpdateFailoverPolicy(old, new *v2.FailoverPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFailoverPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFailoverPolicy indicates an expected call of UpdateFailoverPolicy
func (mr *MockFailoverPolicyEventHandlerMockRecorder) UpdateFailoverPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFailoverPolicy", reflect.TypeOf((*MockFailoverPolicyEventHandler)(nil).UpdateFailoverPolicy), old, new)
}

// DeleteFailoverPolicy mocks base method
func (m *MockFailoverPolicyEventHandler) DeleteFailoverPolicy(obj *v2.FailoverPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFailoverPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFailoverPolicy indicates an expected call of DeleteFailoverPolicy
func (mr *MockFailoverPolicyEventHandlerMockRecorder) DeleteFailoverPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFailoverPolicy", reflect.TypeOf((*MockFailoverPolicyEventHandler)(nil).DeleteFailoverPolicy), obj)
}

// GenericFailoverPolicy mocks base method
func (m *MockFailoverPolicyEventHandler) GenericFailoverPolicy(obj *v2.FailoverPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericFailoverPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericFailoverPolicy indicates an expected call of GenericFailoverPolicy
func (mr *MockFailoverPolicyEventHandlerMockRecorder) GenericFailoverPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericFailoverPolicy", reflect.TypeOf((*MockFailoverPolicyEventHandler)(nil).GenericFailoverPolicy), obj)
}

// MockFailoverPolicyEventWatcher is a mock of FailoverPolicyEventWatcher interface
type MockFailoverPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFailoverPolicyEventWatcherMockRecorder
}

// MockFailoverPolicyEventWatcherMockRecorder is the mock recorder for MockFailoverPolicyEventWatcher
type MockFailoverPolicyEventWatcherMockRecorder struct {
	mock *MockFailoverPolicyEventWatcher
}

// NewMockFailoverPolicyEventWatcher creates a new mock instance
func NewMockFailoverPolicyEventWatcher(ctrl *gomock.Controller) *MockFailoverPolicyEventWatcher {
	mock := &MockFailoverPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockFailoverPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFailoverPolicyEventWatcher) EXPECT() *MockFailoverPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockFailoverPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.FailoverPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockFailoverPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockFailoverPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockOutlierDetectionPolicyEventHandler is a mock of OutlierDetectionPolicyEventHandler interface
type MockOutlierDetectionPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockOutlierDetectionPolicyEventHandlerMockRecorder
}

// MockOutlierDetectionPolicyEventHandlerMockRecorder is the mock recorder for MockOutlierDetectionPolicyEventHandler
type MockOutlierDetectionPolicyEventHandlerMockRecorder struct {
	mock *MockOutlierDetectionPolicyEventHandler
}

// NewMockOutlierDetectionPolicyEventHandler creates a new mock instance
func NewMockOutlierDetectionPolicyEventHandler(ctrl *gomock.Controller) *MockOutlierDetectionPolicyEventHandler {
	mock := &MockOutlierDetectionPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockOutlierDetectionPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOutlierDetectionPolicyEventHandler) EXPECT() *MockOutlierDetectionPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateOutlierDetectionPolicy mocks base method
func (m *MockOutlierDetectionPolicyEventHandler) CreateOutlierDetectionPolicy(obj *v2.OutlierDetectionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOutlierDetectionPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOutlierDetectionPolicy indicates an expected call of CreateOutlierDetectionPolicy
func (mr *MockOutlierDetectionPolicyEventHandlerMockRecorder) CreateOutlierDetectionPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOutlierDetectionPolicy", reflect.TypeOf((*MockOutlierDetectionPolicyEventHandler)(nil).CreateOutlierDetectionPolicy), obj)
}

// UpdateOutlierDetectionPolicy mocks base method
func (m *MockOutlierDetectionPolicyEventHandler) UpdateOutlierDetectionPolicy(old, new *v2.OutlierDetectionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOutlierDetectionPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOutlierDetectionPolicy indicates an expected call of UpdateOutlierDetectionPolicy
func (mr *MockOutlierDetectionPolicyEventHandlerMockRecorder) UpdateOutlierDetectionPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOutlierDetectionPolicy", reflect.TypeOf((*MockOutlierDetectionPolicyEventHandler)(nil).UpdateOutlierDetectionPolicy), old, new)
}

// DeleteOutlierDetectionPolicy mocks base method
func (m *MockOutlierDetectionPolicyEventHandler) DeleteOutlierDetectionPolicy(obj *v2.OutlierDetectionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOutlierDetectionPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOutlierDetectionPolicy indicates an expected call of DeleteOutlierDetectionPolicy
func (mr *MockOutlierDetectionPolicyEventHandlerMockRecorder) DeleteOutlierDetectionPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOutlierDetectionPolicy", reflect.TypeOf((*MockOutlierDetectionPolicyEventHandler)(nil).DeleteOutlierDetectionPolicy), obj)
}

// GenericOutlierDetectionPolicy mocks base method
func (m *MockOutlierDetectionPolicyEventHandler) GenericOutlierDetectionPolicy(obj *v2.OutlierDetectionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericOutlierDetectionPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericOutlierDetectionPolicy indicates an expected call of GenericOutlierDetectionPolicy
func (mr *MockOutlierDetectionPolicyEventHandlerMockRecorder) GenericOutlierDetectionPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericOutlierDetectionPolicy", reflect.TypeOf((*MockOutlierDetectionPolicyEventHandler)(nil).GenericOutlierDetectionPolicy), obj)
}

// MockOutlierDetectionPolicyEventWatcher is a mock of OutlierDetectionPolicyEventWatcher interface
type MockOutlierDetectionPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockOutlierDetectionPolicyEventWatcherMockRecorder
}

// MockOutlierDetectionPolicyEventWatcherMockRecorder is the mock recorder for MockOutlierDetectionPolicyEventWatcher
type MockOutlierDetectionPolicyEventWatcherMockRecorder struct {
	mock *MockOutlierDetectionPolicyEventWatcher
}

// NewMockOutlierDetectionPolicyEventWatcher creates a new mock instance
func NewMockOutlierDetectionPolicyEventWatcher(ctrl *gomock.Controller) *MockOutlierDetectionPolicyEventWatcher {
	mock := &MockOutlierDetectionPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockOutlierDetectionPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOutlierDetectionPolicyEventWatcher) EXPECT() *MockOutlierDetectionPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockOutlierDetectionPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.OutlierDetectionPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockOutlierDetectionPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockOutlierDetectionPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockFaultInjectionPolicyEventHandler is a mock of FaultInjectionPolicyEventHandler interface
type MockFaultInjectionPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFaultInjectionPolicyEventHandlerMockRecorder
}

// MockFaultInjectionPolicyEventHandlerMockRecorder is the mock recorder for MockFaultInjectionPolicyEventHandler
type MockFaultInjectionPolicyEventHandlerMockRecorder struct {
	mock *MockFaultInjectionPolicyEventHandler
}

// NewMockFaultInjectionPolicyEventHandler creates a new mock instance
func NewMockFaultInjectionPolicyEventHandler(ctrl *gomock.Controller) *MockFaultInjectionPolicyEventHandler {
	mock := &MockFaultInjectionPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockFaultInjectionPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFaultInjectionPolicyEventHandler) EXPECT() *MockFaultInjectionPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateFaultInjectionPolicy mocks base method
func (m *MockFaultInjectionPolicyEventHandler) CreateFaultInjectionPolicy(obj *v2.FaultInjectionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFaultInjectionPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFaultInjectionPolicy indicates an expected call of CreateFaultInjectionPolicy
func (mr *MockFaultInjectionPolicyEventHandlerMockRecorder) CreateFaultInjectionPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFaultInjectionPolicy", reflect.TypeOf((*MockFaultInjectionPolicyEventHandler)(nil).CreateFaultInjectionPolicy), obj)
}

// UpdateFaultInjectionPolicy mocks base method
func (m *MockFaultInjectionPolicyEventHandler) UpdateFaultInjectionPolicy(old, new *v2.FaultInjectionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFaultInjectionPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFaultInjectionPolicy indicates an expected call of UpdateFaultInjectionPolicy
func (mr *MockFaultInjectionPolicyEventHandlerMockRecorder) UpdateFaultInjectionPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFaultInjectionPolicy", reflect.TypeOf((*MockFaultInjectionPolicyEventHandler)(nil).UpdateFaultInjectionPolicy), old, new)
}

// DeleteFaultInjectionPolicy mocks base method
func (m *MockFaultInjectionPolicyEventHandler) DeleteFaultInjectionPolicy(obj *v2.FaultInjectionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFaultInjectionPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFaultInjectionPolicy indicates an expected call of DeleteFaultInjectionPolicy
func (mr *MockFaultInjectionPolicyEventHandlerMockRecorder) DeleteFaultInjectionPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFaultInjectionPolicy", reflect.TypeOf((*MockFaultInjectionPolicyEventHandler)(nil).DeleteFaultInjectionPolicy), obj)
}

// GenericFaultInjectionPolicy mocks base method
func (m *MockFaultInjectionPolicyEventHandler) GenericFaultInjectionPolicy(obj *v2.FaultInjectionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericFaultInjectionPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericFaultInjectionPolicy indicates an expected call of GenericFaultInjectionPolicy
func (mr *MockFaultInjectionPolicyEventHandlerMockRecorder) GenericFaultInjectionPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericFaultInjectionPolicy", reflect.TypeOf((*MockFaultInjectionPolicyEventHandler)(nil).GenericFaultInjectionPolicy), obj)
}

// MockFaultInjectionPolicyEventWatcher is a mock of FaultInjectionPolicyEventWatcher interface
type MockFaultInjectionPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFaultInjectionPolicyEventWatcherMockRecorder
}

// MockFaultInjectionPolicyEventWatcherMockRecorder is the mock recorder for MockFaultInjectionPolicyEventWatcher
type MockFaultInjectionPolicyEventWatcherMockRecorder struct {
	mock *MockFaultInjectionPolicyEventWatcher
}

// NewMockFaultInjectionPolicyEventWatcher creates a new mock instance
func NewMockFaultInjectionPolicyEventWatcher(ctrl *gomock.Controller) *MockFaultInjectionPolicyEventWatcher {
	mock := &MockFaultInjectionPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockFaultInjectionPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFaultInjectionPolicyEventWatcher) EXPECT() *MockFaultInjectionPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockFaultInjectionPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.FaultInjectionPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockFaultInjectionPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockFaultInjectionPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockRetryTimeoutPolicyEventHandler is a mock of RetryTimeoutPolicyEventHandler interface
type MockRetryTimeoutPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRetryTimeoutPolicyEventHandlerMockRecorder
}

// MockRetryTimeoutPolicyEventHandlerMockRecorder is the mock recorder for MockRetryTimeoutPolicyEventHandler
type MockRetryTimeoutPolicyEventHandlerMockRecorder struct {
	mock *MockRetryTimeoutPolicyEventHandler
}

// NewMockRetryTimeoutPolicyEventHandler creates a new mock instance
func NewMockRetryTimeoutPolicyEventHandler(ctrl *gomock.Controller) *MockRetryTimeoutPolicyEventHandler {
	mock := &MockRetryTimeoutPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockRetryTimeoutPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRetryTimeoutPolicyEventHandler) EXPECT() *MockRetryTimeoutPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateRetryTimeoutPolicy mocks base method
func (m *MockRetryTimeoutPolicyEventHandler) CreateRetryTimeoutPolicy(obj *v2.RetryTimeoutPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRetryTimeoutPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRetryTimeoutPolicy indicates an expected call of CreateRetryTimeoutPolicy
func (mr *MockRetryTimeoutPolicyEventHandlerMockRecorder) CreateRetryTimeoutPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRetryTimeoutPolicy", reflect.TypeOf((*MockRetryTimeoutPolicyEventHandler)(nil).CreateRetryTimeoutPolicy), obj)
}

// UpdateRetryTimeoutPolicy mocks base method
func (m *MockRetryTimeoutPolicyEventHandler) UpdateRetryTimeoutPolicy(old, new *v2.RetryTimeoutPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRetryTimeoutPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRetryTimeoutPolicy indicates an expected call of UpdateRetryTimeoutPolicy
func (mr *MockRetryTimeoutPolicyEventHandlerMockRecorder) UpdateRetryTimeoutPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRetryTimeoutPolicy", reflect.TypeOf((*MockRetryTimeoutPolicyEventHandler)(nil).UpdateRetryTimeoutPolicy), old, new)
}

// DeleteRetryTimeoutPolicy mocks base method
func (m *MockRetryTimeoutPolicyEventHandler) DeleteRetryTimeoutPolicy(obj *v2.RetryTimeoutPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRetryTimeoutPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRetryTimeoutPolicy indicates an expected call of DeleteRetryTimeoutPolicy
func (mr *MockRetryTimeoutPolicyEventHandlerMockRecorder) DeleteRetryTimeoutPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRetryTimeoutPolicy", reflect.TypeOf((*MockRetryTimeoutPolicyEventHandler)(nil).DeleteRetryTimeoutPolicy), obj)
}

// GenericRetryTimeoutPolicy mocks base method
func (m *MockRetryTimeoutPolicyEventHandler) GenericRetryTimeoutPolicy(obj *v2.RetryTimeoutPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericRetryTimeoutPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericRetryTimeoutPolicy indicates an expected call of GenericRetryTimeoutPolicy
func (mr *MockRetryTimeoutPolicyEventHandlerMockRecorder) GenericRetryTimeoutPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericRetryTimeoutPolicy", reflect.TypeOf((*MockRetryTimeoutPolicyEventHandler)(nil).GenericRetryTimeoutPolicy), obj)
}

// MockRetryTimeoutPolicyEventWatcher is a mock of RetryTimeoutPolicyEventWatcher interface
type MockRetryTimeoutPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockRetryTimeoutPolicyEventWatcherMockRecorder
}

// MockRetryTimeoutPolicyEventWatcherMockRecorder is the mock recorder for MockRetryTimeoutPolicyEventWatcher
type MockRetryTimeoutPolicyEventWatcherMockRecorder struct {
	mock *MockRetryTimeoutPolicyEventWatcher
}

// NewMockRetryTimeoutPolicyEventWatcher creates a new mock instance
func NewMockRetryTimeoutPolicyEventWatcher(ctrl *gomock.Controller) *MockRetryTimeoutPolicyEventWatcher {
	mock := &MockRetryTimeoutPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockRetryTimeoutPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRetryTimeoutPolicyEventWatcher) EXPECT() *MockRetryTimeoutPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockRetryTimeoutPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.RetryTimeoutPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockRetryTimeoutPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockRetryTimeoutPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}