// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2 "github.com/solo-io/solo-apis/pkg/api/security.solo.io/v2"
	controller "github.com/solo-io/solo-apis/pkg/api/security.solo.io/v2/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockAccessPolicyEventHandler is a mock of AccessPolicyEventHandler interface
type MockAccessPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAccessPolicyEventHandlerMockRecorder
}

// MockAccessPolicyEventHandlerMockRecorder is the mock recorder for MockAccessPolicyEventHandler
type MockAccessPolicyEventHandlerMockRecorder struct {
	mock *MockAccessPolicyEventHandler
}

// NewMockAccessPolicyEventHandler creates a new mock instance
func NewMockAccessPolicyEventHandler(ctrl *gomock.Controller) *MockAccessPolicyEventHandler {
	mock := &MockAccessPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockAccessPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccessPolicyEventHandler) EXPECT() *MockAccessPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateAccessPolicy mocks base method
func (m *MockAccessPolicyEventHandler) CreateAccessPolicy(obj *v2.AccessPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessPolicy indicates an expected call of CreateAccessPolicy
func (mr *MockAccessPolicyEventHandlerMockRecorder) CreateAccessPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessPolicy", reflect.TypeOf((*MockAccessPolicyEventHandler)(nil).CreateAccessPolicy), obj)
}

// UpdateAccessPolicy mocks base method
func (m *MockAccessPolicyEventHandler) UpdateAccessPolicy(old, new *v2.AccessPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccessPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccessPolicy indicates an expected call of UpdateAccessPolicy
func (mr *MockAccessPolicyEventHandlerMockRecorder) UpdateAccessPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessPolicy", reflect.TypeOf((*MockAccessPolicyEventHandler)(nil).UpdateAccessPolicy), old, new)
}

// DeleteAccessPolicy mocks base method
func (m *MockAccessPolicyEventHandler) DeleteAccessPolicy(obj *v2.AccessPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessPolicy indicates an expected call of DeleteAccessPolicy
func (mr *MockAccessPolicyEventHandlerMockRecorder) DeleteAccessPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessPolicy", reflect.TypeOf((*MockAccessPolicyEventHandler)(nil).DeleteAccessPolicy), obj)
}

// GenericAccessPolicy mocks base method
func (m *MockAccessPolicyEventHandler) GenericAccessPolicy(obj *v2.AccessPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericAccessPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericAccessPolicy indicates an expected call of GenericAccessPolicy
func (mr *MockAccessPolicyEventHandlerMockRecorder) GenericAccessPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericAccessPolicy", reflect.TypeOf((*MockAccessPolicyEventHandler)(nil).GenericAccessPolicy), obj)
}

// MockAccessPolicyEventWatcher is a mock of AccessPolicyEventWatcher interface
type MockAccessPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockAccessPolicyEventWatcherMockRecorder
}

// MockAccessPolicyEventWatcherMockRecorder is the mock recorder for MockAccessPolicyEventWatcher
type MockAccessPolicyEventWatcherMockRecorder struct {
	mock *MockAccessPolicyEventWatcher
}

// NewMockAccessPolicyEventWatcher creates a new mock instance
func NewMockAccessPolicyEventWatcher(ctrl *gomock.Controller) *MockAccessPolicyEventWatcher {
	mock := &MockAccessPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockAccessPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccessPolicyEventWatcher) EXPECT() *MockAccessPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockAccessPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.AccessPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockAccessPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockAccessPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockCORSPolicyEventHandler is a mock of CORSPolicyEventHandler interface
type MockCORSPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCORSPolicyEventHandlerMockRecorder
}

// MockCORSPolicyEventHandlerMockRecorder is the mock recorder for MockCORSPolicyEventHandler
type MockCORSPolicyEventHandlerMockRecorder struct {
	mock *MockCORSPolicyEventHandler
}

// NewMockCORSPolicyEventHandler creates a new mock instance
func NewMockCORSPolicyEventHandler(ctrl *gomock.Controller) *MockCORSPolicyEventHandler {
	mock := &MockCORSPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockCORSPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCORSPolicyEventHandler) EXPECT() *MockCORSPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateCORSPolicy mocks base method
func (m *MockCORSPolicyEventHandler) CreateCORSPolicy(obj *v2.CORSPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCORSPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCORSPolicy indicates an expected call of CreateCORSPolicy
func (mr *MockCORSPolicyEventHandlerMockRecorder) CreateCORSPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCORSPolicy", reflect.TypeOf((*MockCORSPolicyEventHandler)(nil).CreateCORSPolicy), obj)
}

// UpdateCORSPolicy mocks base method
func (m *MockCORSPolicyEventHandler) UpdateCORSPolicy(old, new *v2.CORSPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCORSPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCORSPolicy indicates an expected call of UpdateCORSPolicy
func (mr *MockCORSPolicyEventHandlerMockRecorder) UpdateCORSPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCORSPolicy", reflect.TypeOf((*MockCORSPolicyEventHandler)(nil).UpdateCORSPolicy), old, new)
}

// DeleteCORSPolicy mocks base method
func (m *MockCORSPolicyEventHandler) DeleteCORSPolicy(obj *v2.CORSPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCORSPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCORSPolicy indicates an expected call of DeleteCORSPolicy
func (mr *MockCORSPolicyEventHandlerMockRecorder) DeleteCORSPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCORSPolicy", reflect.TypeOf((*MockCORSPolicyEventHandler)(nil).DeleteCORSPolicy), obj)
}

// GenericCORSPolicy mocks base method
func (m *MockCORSPolicyEventHandler) GenericCORSPolicy(obj *v2.CORSPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericCORSPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericCORSPolicy indicates an expected call of GenericCORSPolicy
func (mr *MockCORSPolicyEventHandlerMockRecorder) GenericCORSPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericCORSPolicy", reflect.TypeOf((*MockCORSPolicyEventHandler)(nil).GenericCORSPolicy), obj)
}

// MockCORSPolicyEventWatcher is a mock of CORSPolicyEventWatcher interface
type MockCORSPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockCORSPolicyEventWatcherMockRecorder
}

// MockCORSPolicyEventWatcherMockRecorder is the mock recorder for MockCORSPolicyEventWatcher
type MockCORSPolicyEventWatcherMockRecorder struct {
	mock *MockCORSPolicyEventWatcher
}

// NewMockCORSPolicyEventWatcher creates a new mock instance
func NewMockCORSPolicyEventWatcher(ctrl *gomock.Controller) *MockCORSPolicyEventWatcher {
	mock := &MockCORSPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockCORSPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCORSPolicyEventWatcher) EXPECT() *MockCORSPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockCORSPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.CORSPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockCORSPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockCORSPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockCSRFPolicyEventHandler is a mock of CSRFPolicyEventHandler interface
type MockCSRFPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCSRFPolicyEventHandlerMockRecorder
}

// MockCSRFPolicyEventHandlerMockRecorder is the mock recorder for MockCSRFPolicyEventHandler
type MockCSRFPolicyEventHandlerMockRecorder struct {
	mock *MockCSRFPolicyEventHandler
}

// NewMockCSRFPolicyEventHandler creates a new mock instance
func NewMockCSRFPolicyEventHandler(ctrl *gomock.Controller) *MockCSRFPolicyEventHandler {
	mock := &MockCSRFPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockCSRFPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCSRFPolicyEventHandler) EXPECT() *MockCSRFPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateCSRFPolicy mocks base method
func (m *MockCSRFPolicyEventHandler) CreateCSRFPolicy(obj *v2.CSRFPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCSRFPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCSRFPolicy indicates an expected call of CreateCSRFPolicy
func (mr *MockCSRFPolicyEventHandlerMockRecorder) CreateCSRFPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCSRFPolicy", reflect.TypeOf((*MockCSRFPolicyEventHandler)(nil).CreateCSRFPolicy), obj)
}

// UpdateCSRFPolicy mocks base method
func (m *MockCSRFPolicyEventHandler) UpdateCSRFPolicy(old, new *v2.CSRFPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCSRFPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCSRFPolicy indicates an expected call of UpdateCSRFPolicy
func (mr *MockCSRFPolicyEventHandlerMockRecorder) UpdateCSRFPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCSRFPolicy", reflect.TypeOf((*MockCSRFPolicyEventHandler)(nil).UpdateCSRFPolicy), old, new)
}

// DeleteCSRFPolicy mocks base method
func (m *MockCSRFPolicyEventHandler) DeleteCSRFPolicy(obj *v2.CSRFPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCSRFPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCSRFPolicy indicates an expected call of DeleteCSRFPolicy
func (mr *MockCSRFPolicyEventHandlerMockRecorder) DeleteCSRFPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCSRFPolicy", reflect.TypeOf((*MockCSRFPolicyEventHandler)(nil).DeleteCSRFPolicy), obj)
}

// GenericCSRFPolicy mocks base method
func (m *MockCSRFPolicyEventHandler) GenericCSRFPolicy(obj *v2.CSRFPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericCSRFPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericCSRFPolicy indicates an expected call of GenericCSRFPolicy
func (mr *MockCSRFPolicyEventHandlerMockRecorder) GenericCSRFPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericCSRFPolicy", reflect.TypeOf((*MockCSRFPolicyEventHandler)(nil).GenericCSRFPolicy), obj)
}

// MockCSRFPolicyEventWatcher is a mock of CSRFPolicyEventWatcher interface
type MockCSRFPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockCSRFPolicyEventWatcherMockRecorder
}

// MockCSRFPolicyEventWatcherMockRecorder is the mock recorder for MockCSRFPolicyEventWatcher
type MockCSRFPolicyEventWatcherMockRecorder struct {
	mock *MockCSRFPolicyEventWatcher
}

// NewMockCSRFPolicyEventWatcher creates a new mock instance
func NewMockCSRFPolicyEventWatcher(ctrl *gomock.Controller) *MockCSRFPolicyEventWatcher {
	mock := &MockCSRFPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockCSRFPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCSRFPolicyEventWatcher) EXPECT() *MockCSRFPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockCSRFPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.CSRFPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockCSRFPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockCSRFPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockExtAuthPolicyEventHandler is a mock of ExtAuthPolicyEventHandler interface
type MockExtAuthPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockExtAuthPolicyEventHandlerMockRecorder
}

// MockExtAuthPolicyEventHandlerMockRecorder is the mock recorder for MockExtAuthPolicyEventHandler
type MockExtAuthPolicyEventHandlerMockRecorder struct {
	mock *MockExtAuthPolicyEventHandler
}

// NewMockExtAuthPolicyEventHandler creates a new mock instance
func NewMockExtAuthPolicyEventHandler(ctrl *gomock.Controller) *MockExtAuthPolicyEventHandler {
	mock := &MockExtAuthPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockExtAuthPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExtAuthPolicyEventHandler) EXPECT() *MockExtAuthPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateExtAuthPolicy mocks base method
func (m *MockExtAuthPolicyEventHandler) CreateExtAuthPolicy(obj *v2.ExtAuthPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExtAuthPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExtAuthPolicy indicates an expected call of CreateExtAuthPolicy
func (mr *MockExtAuthPolicyEventHandlerMockRecorder) CreateExtAuthPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExtAuthPolicy", reflect.TypeOf((*MockExtAuthPolicyEventHandler)(nil).CreateExtAuthPolicy), obj)
}

// UpdateExtAuthPolicy mocks base method
func (m *MockExtAuthPolicyEventHandler) UpdateExtAuthPolicy(old, new *v2.ExtAuthPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExtAuthPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExtAuthPolicy indicates an expected call of UpdateExtAuthPolicy
func (mr *MockExtAuthPolicyEventHandlerMockRecorder) UpdateExtAuthPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExtAuthPolicy", reflect.TypeOf((*MockExtAuthPolicyEventHandler)(nil).UpdateExtAuthPolicy), old, new)
}

// DeleteExtAuthPolicy mocks base method
func (m *MockExtAuthPolicyEventHandler) DeleteExtAuthPolicy(obj *v2.ExtAuthPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExtAuthPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExtAuthPolicy indicates an expected call of DeleteExtAuthPolicy
func (mr *MockExtAuthPolicyEventHandlerMockRecorder) DeleteExtAuthPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExtAuthPolicy", reflect.TypeOf((*MockExtAuthPolicyEventHandler)(nil).DeleteExtAuthPolicy), obj)
}

// GenericExtAuthPolicy mocks base method
func (m *MockExtAuthPolicyEventHandler) GenericExtAuthPolicy(obj *v2.ExtAuthPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericExtAuthPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericExtAuthPolicy indicates an expected call of GenericExtAuthPolicy
func (mr *MockExtAuthPolicyEventHandlerMockRecorder) GenericExtAuthPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericExtAuthPolicy", reflect.TypeOf((*MockExtAuthPolicyEventHandler)(nil).GenericExtAuthPolicy), obj)
}

// MockExtAuthPolicyEventWatcher is a mock of ExtAuthPolicyEventWatcher interface
type MockExtAuthPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockExtAuthPolicyEventWatcherMockRecorder
}

// MockExtAuthPolicyEventWatcherMockRecorder is the mock recorder for MockExtAuthPolicyEventWatcher
type MockExtAuthPolicyEventWatcherMockRecorder struct {
	mock *MockExtAuthPolicyEventWatcher
}

// NewMockExtAuthPolicyEventWatcher creates a new mock instance
func NewMockExtAuthPolicyEventWatcher(ctrl *gomock.Controller) *MockExtAuthPolicyEventWatcher {
	mock := &MockExtAuthPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockExtAuthPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExtAuthPolicyEventWatcher) EXPECT() *MockExtAuthPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockExtAuthPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.ExtAuthPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockExtAuthPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockExtAuthPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockWAFPolicyEventHandler is a mock of WAFPolicyEventHandler interface
type MockWAFPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWAFPolicyEventHandlerMockRecorder
}

// MockWAFPolicyEventHandlerMockRecorder is the mock recorder for MockWAFPolicyEventHandler
type MockWAFPolicyEventHandlerMockRecorder struct {
	mock *MockWAFPolicyEventHandler
}

// NewMockWAFPolicyEventHandler creates a new mock instance
func NewMockWAFPolicyEventHandler(ctrl *gomock.Controller) *MockWAFPolicyEventHandler {
	mock := &MockWAFPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockWAFPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWAFPolicyEventHandler) EXPECT() *MockWAFPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateWAFPolicy mocks base method
func (m *MockWAFPolicyEventHandler) CreateWAFPolicy(obj *v2.WAFPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWAFPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWAFPolicy indicates an expected call of CreateWAFPolicy
func (mr *MockWAFPolicyEventHandlerMockRecorder) CreateWAFPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWAFPolicy", reflect.TypeOf((*MockWAFPolicyEventHandler)(nil).CreateWAFPolicy), obj)
}

// UpdateWAFPolicy mocks base method
func (m *MockWAFPolicyEventHandler) UpdateWAFPolicy(old, new *v2.WAFPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWAFPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWAFPolicy indicates an expected call of UpdateWAFPolicy
func (mr *MockWAFPolicyEventHandlerMockRecorder) UpdateWAFPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWAFPolicy", reflect.TypeOf((*MockWAFPolicyEventHandler)(nil).UpdateWAFPolicy), old, new)
}

// DeleteWAFPolicy mocks base method
func (m *MockWAFPolicyEventHandler) DeleteWAFPolicy(obj *v2.WAFPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWAFPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWAFPolicy indicates an expected call of DeleteWAFPolicy
func (mr *MockWAFPolicyEventHandlerMockRecorder) DeleteWAFPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWAFPolicy", reflect.TypeOf((*MockWAFPolicyEventHandler)(nil).DeleteWAFPolicy), obj)
}

// GenericWAFPolicy mocks base method
func (m *MockWAFPolicyEventHandler) GenericWAFPolicy(obj *v2.WAFPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericWAFPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericWAFPolicy indicates an expected call of GenericWAFPolicy
func (mr *MockWAFPolicyEventHandlerMockRecorder) GenericWAFPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericWAFPolicy", reflect.TypeOf((*MockWAFPolicyEventHandler)(nil).GenericWAFPolicy), obj)
}

// MockWAFPolicyEventWatcher is a mock of WAFPolicyEventWatcher interface
type MockWAFPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockWAFPolicyEventWatcherMockRecorder
}

// MockWAFPolicyEventWatcherMockRecorder is the mock recorder for MockWAFPolicyEventWatcher
type MockWAFPolicyEventWatcherMockRecorder struct {
	mock *MockWAFPolicyEventWatcher
}

// NewMockWAFPolicyEventWatcher creates a new mock instance
func NewMockWAFPolicyEventWatcher(ctrl *gomock.Controller) *MockWAFPolicyEventWatcher {
	mock := &MockWAFPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockWAFPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWAFPolicyEventWatcher) EXPECT() *MockWAFPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockWAFPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.WAFPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockWAFPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockWAFPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockJWTPolicyEventHandler is a mock of JWTPolicyEventHandler interface
type MockJWTPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockJWTPolicyEventHandlerMockRecorder
}

// MockJWTPolicyEventHandlerMockRecorder is the mock recorder for MockJWTPolicyEventHandler
type MockJWTPolicyEventHandlerMockRecorder struct {
	mock *MockJWTPolicyEventHandler
}

// NewMockJWTPolicyEventHandler creates a new mock instance
func NewMockJWTPolicyEventHandler(ctrl *gomock.Controller) *MockJWTPolicyEventHandler {
	mock := &MockJWTPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockJWTPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJWTPolicyEventHandler) EXPECT() *MockJWTPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateJWTPolicy mocks base method
func (m *MockJWTPolicyEventHandler) CreateJWTPolicy(obj *v2.JWTPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJWTPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateJWTPolicy indicates an expected call of CreateJWTPolicy
func (mr *MockJWTPolicyEventHandlerMockRecorder) CreateJWTPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJWTPolicy", reflect.TypeOf((*MockJWTPolicyEventHandler)(nil).CreateJWTPolicy), obj)
}

// UpdateJWTPolicy mocks base method
func (m *MockJWTPolicyEventHandler) UpdateJWTPolicy(old, new *v2.JWTPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJWTPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateJWTPolicy indicates an expected call of UpdateJWTPolicy
func (mr *MockJWTPolicyEventHandlerMockRecorder) UpdateJWTPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJWTPolicy", reflect.TypeOf((*MockJWTPolicyEventHandler)(nil).UpdateJWTPolicy), old, new)
}

// DeleteJWTPolicy mocks base method
func (m *MockJWTPolicyEventHandler) DeleteJWTPolicy(obj *v2.JWTPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJWTPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJWTPolicy indicates an expected call of DeleteJWTPolicy
func (mr *MockJWTPolicyEventHandlerMockRecorder) DeleteJWTPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJWTPolicy", reflect.TypeOf((*MockJWTPolicyEventHandler)(nil).DeleteJWTPolicy), obj)
}

// GenericJWTPolicy mocks base method
func (m *MockJWTPolicyEventHandler) GenericJWTPolicy(obj *v2.JWTPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericJWTPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericJWTPolicy indicates an expected call of GenericJWTPolicy
func (mr *MockJWTPolicyEventHandlerMockRecorder) GenericJWTPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericJWTPolicy", reflect.TypeOf((*MockJWTPolicyEventHandler)(nil).GenericJWTPolicy), obj)
}

// MockJWTPolicyEventWatcher is a mock of JWTPolicyEventWatcher interface
type MockJWTPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockJWTPolicyEventWatcherMockRecorder
}

// MockJWTPolicyEventWatcherMockRecorder is the mock recorder for MockJWTPolicyEventWatcher
type MockJWTPolicyEventWatcherMockRecorder struct {
	mock *MockJWTPolicyEventWatcher
}

// NewMockJWTPolicyEventWatcher creates a new mock instance
func NewMockJWTPolicyEventWatcher(ctrl *gomock.Controller) *MockJWTPolicyEventWatcher {
	mock := &MockJWTPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockJWTPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJWTPolicyEventWatcher) EXPECT() *MockJWTPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockJWTPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.JWTPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockJWTPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockJWTPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}
