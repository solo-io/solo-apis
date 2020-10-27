// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockSettingsEventHandler is a mock of SettingsEventHandler interface
type MockSettingsEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsEventHandlerMockRecorder
}

// MockSettingsEventHandlerMockRecorder is the mock recorder for MockSettingsEventHandler
type MockSettingsEventHandlerMockRecorder struct {
	mock *MockSettingsEventHandler
}

// NewMockSettingsEventHandler creates a new mock instance
func NewMockSettingsEventHandler(ctrl *gomock.Controller) *MockSettingsEventHandler {
	mock := &MockSettingsEventHandler{ctrl: ctrl}
	mock.recorder = &MockSettingsEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsEventHandler) EXPECT() *MockSettingsEventHandlerMockRecorder {
	return m.recorder
}

// CreateSettings mocks base method
func (m *MockSettingsEventHandler) CreateSettings(obj *v1.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSettings", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSettings indicates an expected call of CreateSettings
func (mr *MockSettingsEventHandlerMockRecorder) CreateSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSettings", reflect.TypeOf((*MockSettingsEventHandler)(nil).CreateSettings), obj)
}

// UpdateSettings mocks base method
func (m *MockSettingsEventHandler) UpdateSettings(old, new *v1.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSettings", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettings indicates an expected call of UpdateSettings
func (mr *MockSettingsEventHandlerMockRecorder) UpdateSettings(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettings", reflect.TypeOf((*MockSettingsEventHandler)(nil).UpdateSettings), old, new)
}

// DeleteSettings mocks base method
func (m *MockSettingsEventHandler) DeleteSettings(obj *v1.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSettings", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSettings indicates an expected call of DeleteSettings
func (mr *MockSettingsEventHandlerMockRecorder) DeleteSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSettings", reflect.TypeOf((*MockSettingsEventHandler)(nil).DeleteSettings), obj)
}

// GenericSettings mocks base method
func (m *MockSettingsEventHandler) GenericSettings(obj *v1.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericSettings", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericSettings indicates an expected call of GenericSettings
func (mr *MockSettingsEventHandlerMockRecorder) GenericSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericSettings", reflect.TypeOf((*MockSettingsEventHandler)(nil).GenericSettings), obj)
}

// MockSettingsEventWatcher is a mock of SettingsEventWatcher interface
type MockSettingsEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsEventWatcherMockRecorder
}

// MockSettingsEventWatcherMockRecorder is the mock recorder for MockSettingsEventWatcher
type MockSettingsEventWatcherMockRecorder struct {
	mock *MockSettingsEventWatcher
}

// NewMockSettingsEventWatcher creates a new mock instance
func NewMockSettingsEventWatcher(ctrl *gomock.Controller) *MockSettingsEventWatcher {
	mock := &MockSettingsEventWatcher{ctrl: ctrl}
	mock.recorder = &MockSettingsEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsEventWatcher) EXPECT() *MockSettingsEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockSettingsEventWatcher) AddEventHandler(ctx context.Context, h controller.SettingsEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockSettingsEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockSettingsEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockUpstreamEventHandler is a mock of UpstreamEventHandler interface
type MockUpstreamEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamEventHandlerMockRecorder
}

// MockUpstreamEventHandlerMockRecorder is the mock recorder for MockUpstreamEventHandler
type MockUpstreamEventHandlerMockRecorder struct {
	mock *MockUpstreamEventHandler
}

// NewMockUpstreamEventHandler creates a new mock instance
func NewMockUpstreamEventHandler(ctrl *gomock.Controller) *MockUpstreamEventHandler {
	mock := &MockUpstreamEventHandler{ctrl: ctrl}
	mock.recorder = &MockUpstreamEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamEventHandler) EXPECT() *MockUpstreamEventHandlerMockRecorder {
	return m.recorder
}

// CreateUpstream mocks base method
func (m *MockUpstreamEventHandler) CreateUpstream(obj *v1.Upstream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpstream", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUpstream indicates an expected call of CreateUpstream
func (mr *MockUpstreamEventHandlerMockRecorder) CreateUpstream(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpstream", reflect.TypeOf((*MockUpstreamEventHandler)(nil).CreateUpstream), obj)
}

// UpdateUpstream mocks base method
func (m *MockUpstreamEventHandler) UpdateUpstream(old, new *v1.Upstream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUpstream", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUpstream indicates an expected call of UpdateUpstream
func (mr *MockUpstreamEventHandlerMockRecorder) UpdateUpstream(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUpstream", reflect.TypeOf((*MockUpstreamEventHandler)(nil).UpdateUpstream), old, new)
}

// DeleteUpstream mocks base method
func (m *MockUpstreamEventHandler) DeleteUpstream(obj *v1.Upstream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUpstream", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUpstream indicates an expected call of DeleteUpstream
func (mr *MockUpstreamEventHandlerMockRecorder) DeleteUpstream(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUpstream", reflect.TypeOf((*MockUpstreamEventHandler)(nil).DeleteUpstream), obj)
}

// GenericUpstream mocks base method
func (m *MockUpstreamEventHandler) GenericUpstream(obj *v1.Upstream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericUpstream", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericUpstream indicates an expected call of GenericUpstream
func (mr *MockUpstreamEventHandlerMockRecorder) GenericUpstream(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericUpstream", reflect.TypeOf((*MockUpstreamEventHandler)(nil).GenericUpstream), obj)
}

// MockUpstreamEventWatcher is a mock of UpstreamEventWatcher interface
type MockUpstreamEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamEventWatcherMockRecorder
}

// MockUpstreamEventWatcherMockRecorder is the mock recorder for MockUpstreamEventWatcher
type MockUpstreamEventWatcherMockRecorder struct {
	mock *MockUpstreamEventWatcher
}

// NewMockUpstreamEventWatcher creates a new mock instance
func NewMockUpstreamEventWatcher(ctrl *gomock.Controller) *MockUpstreamEventWatcher {
	mock := &MockUpstreamEventWatcher{ctrl: ctrl}
	mock.recorder = &MockUpstreamEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamEventWatcher) EXPECT() *MockUpstreamEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockUpstreamEventWatcher) AddEventHandler(ctx context.Context, h controller.UpstreamEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockUpstreamEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockUpstreamEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockUpstreamGroupEventHandler is a mock of UpstreamGroupEventHandler interface
type MockUpstreamGroupEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamGroupEventHandlerMockRecorder
}

// MockUpstreamGroupEventHandlerMockRecorder is the mock recorder for MockUpstreamGroupEventHandler
type MockUpstreamGroupEventHandlerMockRecorder struct {
	mock *MockUpstreamGroupEventHandler
}

// NewMockUpstreamGroupEventHandler creates a new mock instance
func NewMockUpstreamGroupEventHandler(ctrl *gomock.Controller) *MockUpstreamGroupEventHandler {
	mock := &MockUpstreamGroupEventHandler{ctrl: ctrl}
	mock.recorder = &MockUpstreamGroupEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamGroupEventHandler) EXPECT() *MockUpstreamGroupEventHandlerMockRecorder {
	return m.recorder
}

// CreateUpstreamGroup mocks base method
func (m *MockUpstreamGroupEventHandler) CreateUpstreamGroup(obj *v1.UpstreamGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpstreamGroup", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUpstreamGroup indicates an expected call of CreateUpstreamGroup
func (mr *MockUpstreamGroupEventHandlerMockRecorder) CreateUpstreamGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpstreamGroup", reflect.TypeOf((*MockUpstreamGroupEventHandler)(nil).CreateUpstreamGroup), obj)
}

// UpdateUpstreamGroup mocks base method
func (m *MockUpstreamGroupEventHandler) UpdateUpstreamGroup(old, new *v1.UpstreamGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUpstreamGroup", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUpstreamGroup indicates an expected call of UpdateUpstreamGroup
func (mr *MockUpstreamGroupEventHandlerMockRecorder) UpdateUpstreamGroup(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUpstreamGroup", reflect.TypeOf((*MockUpstreamGroupEventHandler)(nil).UpdateUpstreamGroup), old, new)
}

// DeleteUpstreamGroup mocks base method
func (m *MockUpstreamGroupEventHandler) DeleteUpstreamGroup(obj *v1.UpstreamGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUpstreamGroup", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUpstreamGroup indicates an expected call of DeleteUpstreamGroup
func (mr *MockUpstreamGroupEventHandlerMockRecorder) DeleteUpstreamGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUpstreamGroup", reflect.TypeOf((*MockUpstreamGroupEventHandler)(nil).DeleteUpstreamGroup), obj)
}

// GenericUpstreamGroup mocks base method
func (m *MockUpstreamGroupEventHandler) GenericUpstreamGroup(obj *v1.UpstreamGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericUpstreamGroup", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericUpstreamGroup indicates an expected call of GenericUpstreamGroup
func (mr *MockUpstreamGroupEventHandlerMockRecorder) GenericUpstreamGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericUpstreamGroup", reflect.TypeOf((*MockUpstreamGroupEventHandler)(nil).GenericUpstreamGroup), obj)
}

// MockUpstreamGroupEventWatcher is a mock of UpstreamGroupEventWatcher interface
type MockUpstreamGroupEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamGroupEventWatcherMockRecorder
}

// MockUpstreamGroupEventWatcherMockRecorder is the mock recorder for MockUpstreamGroupEventWatcher
type MockUpstreamGroupEventWatcherMockRecorder struct {
	mock *MockUpstreamGroupEventWatcher
}

// NewMockUpstreamGroupEventWatcher creates a new mock instance
func NewMockUpstreamGroupEventWatcher(ctrl *gomock.Controller) *MockUpstreamGroupEventWatcher {
	mock := &MockUpstreamGroupEventWatcher{ctrl: ctrl}
	mock.recorder = &MockUpstreamGroupEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamGroupEventWatcher) EXPECT() *MockUpstreamGroupEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockUpstreamGroupEventWatcher) AddEventHandler(ctx context.Context, h controller.UpstreamGroupEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockUpstreamGroupEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockUpstreamGroupEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockProxyEventHandler is a mock of ProxyEventHandler interface
type MockProxyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockProxyEventHandlerMockRecorder
}

// MockProxyEventHandlerMockRecorder is the mock recorder for MockProxyEventHandler
type MockProxyEventHandlerMockRecorder struct {
	mock *MockProxyEventHandler
}

// NewMockProxyEventHandler creates a new mock instance
func NewMockProxyEventHandler(ctrl *gomock.Controller) *MockProxyEventHandler {
	mock := &MockProxyEventHandler{ctrl: ctrl}
	mock.recorder = &MockProxyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProxyEventHandler) EXPECT() *MockProxyEventHandlerMockRecorder {
	return m.recorder
}

// CreateProxy mocks base method
func (m *MockProxyEventHandler) CreateProxy(obj *v1.Proxy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProxy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProxy indicates an expected call of CreateProxy
func (mr *MockProxyEventHandlerMockRecorder) CreateProxy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProxy", reflect.TypeOf((*MockProxyEventHandler)(nil).CreateProxy), obj)
}

// UpdateProxy mocks base method
func (m *MockProxyEventHandler) UpdateProxy(old, new *v1.Proxy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProxy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProxy indicates an expected call of UpdateProxy
func (mr *MockProxyEventHandlerMockRecorder) UpdateProxy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProxy", reflect.TypeOf((*MockProxyEventHandler)(nil).UpdateProxy), old, new)
}

// DeleteProxy mocks base method
func (m *MockProxyEventHandler) DeleteProxy(obj *v1.Proxy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProxy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProxy indicates an expected call of DeleteProxy
func (mr *MockProxyEventHandlerMockRecorder) DeleteProxy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProxy", reflect.TypeOf((*MockProxyEventHandler)(nil).DeleteProxy), obj)
}

// GenericProxy mocks base method
func (m *MockProxyEventHandler) GenericProxy(obj *v1.Proxy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericProxy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericProxy indicates an expected call of GenericProxy
func (mr *MockProxyEventHandlerMockRecorder) GenericProxy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericProxy", reflect.TypeOf((*MockProxyEventHandler)(nil).GenericProxy), obj)
}

// MockProxyEventWatcher is a mock of ProxyEventWatcher interface
type MockProxyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockProxyEventWatcherMockRecorder
}

// MockProxyEventWatcherMockRecorder is the mock recorder for MockProxyEventWatcher
type MockProxyEventWatcherMockRecorder struct {
	mock *MockProxyEventWatcher
}

// NewMockProxyEventWatcher creates a new mock instance
func NewMockProxyEventWatcher(ctrl *gomock.Controller) *MockProxyEventWatcher {
	mock := &MockProxyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockProxyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProxyEventWatcher) EXPECT() *MockProxyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockProxyEventWatcher) AddEventHandler(ctx context.Context, h controller.ProxyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockProxyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockProxyEventWatcher)(nil).AddEventHandler), varargs...)
}
