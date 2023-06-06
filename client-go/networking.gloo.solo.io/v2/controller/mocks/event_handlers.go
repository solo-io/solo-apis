// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2 "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2"
	controller "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockCloudProviderEventHandler is a mock of CloudProviderEventHandler interface.
type MockCloudProviderEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderEventHandlerMockRecorder
}

// MockCloudProviderEventHandlerMockRecorder is the mock recorder for MockCloudProviderEventHandler.
type MockCloudProviderEventHandlerMockRecorder struct {
	mock *MockCloudProviderEventHandler
}

// NewMockCloudProviderEventHandler creates a new mock instance.
func NewMockCloudProviderEventHandler(ctrl *gomock.Controller) *MockCloudProviderEventHandler {
	mock := &MockCloudProviderEventHandler{ctrl: ctrl}
	mock.recorder = &MockCloudProviderEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudProviderEventHandler) EXPECT() *MockCloudProviderEventHandlerMockRecorder {
	return m.recorder
}

// CreateCloudProvider mocks base method.
func (m *MockCloudProviderEventHandler) CreateCloudProvider(obj *v2.CloudProvider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCloudProvider", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCloudProvider indicates an expected call of CreateCloudProvider.
func (mr *MockCloudProviderEventHandlerMockRecorder) CreateCloudProvider(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCloudProvider", reflect.TypeOf((*MockCloudProviderEventHandler)(nil).CreateCloudProvider), obj)
}

// DeleteCloudProvider mocks base method.
func (m *MockCloudProviderEventHandler) DeleteCloudProvider(obj *v2.CloudProvider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCloudProvider", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCloudProvider indicates an expected call of DeleteCloudProvider.
func (mr *MockCloudProviderEventHandlerMockRecorder) DeleteCloudProvider(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCloudProvider", reflect.TypeOf((*MockCloudProviderEventHandler)(nil).DeleteCloudProvider), obj)
}

// GenericCloudProvider mocks base method.
func (m *MockCloudProviderEventHandler) GenericCloudProvider(obj *v2.CloudProvider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericCloudProvider", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericCloudProvider indicates an expected call of GenericCloudProvider.
func (mr *MockCloudProviderEventHandlerMockRecorder) GenericCloudProvider(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericCloudProvider", reflect.TypeOf((*MockCloudProviderEventHandler)(nil).GenericCloudProvider), obj)
}

// UpdateCloudProvider mocks base method.
func (m *MockCloudProviderEventHandler) UpdateCloudProvider(old, new *v2.CloudProvider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCloudProvider", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCloudProvider indicates an expected call of UpdateCloudProvider.
func (mr *MockCloudProviderEventHandlerMockRecorder) UpdateCloudProvider(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCloudProvider", reflect.TypeOf((*MockCloudProviderEventHandler)(nil).UpdateCloudProvider), old, new)
}

// MockCloudProviderEventWatcher is a mock of CloudProviderEventWatcher interface.
type MockCloudProviderEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockCloudProviderEventWatcherMockRecorder
}

// MockCloudProviderEventWatcherMockRecorder is the mock recorder for MockCloudProviderEventWatcher.
type MockCloudProviderEventWatcherMockRecorder struct {
	mock *MockCloudProviderEventWatcher
}

// NewMockCloudProviderEventWatcher creates a new mock instance.
func NewMockCloudProviderEventWatcher(ctrl *gomock.Controller) *MockCloudProviderEventWatcher {
	mock := &MockCloudProviderEventWatcher{ctrl: ctrl}
	mock.recorder = &MockCloudProviderEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudProviderEventWatcher) EXPECT() *MockCloudProviderEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockCloudProviderEventWatcher) AddEventHandler(ctx context.Context, h controller.CloudProviderEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockCloudProviderEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockCloudProviderEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockCloudResourcesEventHandler is a mock of CloudResourcesEventHandler interface.
type MockCloudResourcesEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCloudResourcesEventHandlerMockRecorder
}

// MockCloudResourcesEventHandlerMockRecorder is the mock recorder for MockCloudResourcesEventHandler.
type MockCloudResourcesEventHandlerMockRecorder struct {
	mock *MockCloudResourcesEventHandler
}

// NewMockCloudResourcesEventHandler creates a new mock instance.
func NewMockCloudResourcesEventHandler(ctrl *gomock.Controller) *MockCloudResourcesEventHandler {
	mock := &MockCloudResourcesEventHandler{ctrl: ctrl}
	mock.recorder = &MockCloudResourcesEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudResourcesEventHandler) EXPECT() *MockCloudResourcesEventHandlerMockRecorder {
	return m.recorder
}

// CreateCloudResources mocks base method.
func (m *MockCloudResourcesEventHandler) CreateCloudResources(obj *v2.CloudResources) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCloudResources", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCloudResources indicates an expected call of CreateCloudResources.
func (mr *MockCloudResourcesEventHandlerMockRecorder) CreateCloudResources(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCloudResources", reflect.TypeOf((*MockCloudResourcesEventHandler)(nil).CreateCloudResources), obj)
}

// DeleteCloudResources mocks base method.
func (m *MockCloudResourcesEventHandler) DeleteCloudResources(obj *v2.CloudResources) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCloudResources", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCloudResources indicates an expected call of DeleteCloudResources.
func (mr *MockCloudResourcesEventHandlerMockRecorder) DeleteCloudResources(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCloudResources", reflect.TypeOf((*MockCloudResourcesEventHandler)(nil).DeleteCloudResources), obj)
}

// GenericCloudResources mocks base method.
func (m *MockCloudResourcesEventHandler) GenericCloudResources(obj *v2.CloudResources) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericCloudResources", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericCloudResources indicates an expected call of GenericCloudResources.
func (mr *MockCloudResourcesEventHandlerMockRecorder) GenericCloudResources(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericCloudResources", reflect.TypeOf((*MockCloudResourcesEventHandler)(nil).GenericCloudResources), obj)
}

// UpdateCloudResources mocks base method.
func (m *MockCloudResourcesEventHandler) UpdateCloudResources(old, new *v2.CloudResources) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCloudResources", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCloudResources indicates an expected call of UpdateCloudResources.
func (mr *MockCloudResourcesEventHandlerMockRecorder) UpdateCloudResources(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCloudResources", reflect.TypeOf((*MockCloudResourcesEventHandler)(nil).UpdateCloudResources), old, new)
}

// MockCloudResourcesEventWatcher is a mock of CloudResourcesEventWatcher interface.
type MockCloudResourcesEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockCloudResourcesEventWatcherMockRecorder
}

// MockCloudResourcesEventWatcherMockRecorder is the mock recorder for MockCloudResourcesEventWatcher.
type MockCloudResourcesEventWatcherMockRecorder struct {
	mock *MockCloudResourcesEventWatcher
}

// NewMockCloudResourcesEventWatcher creates a new mock instance.
func NewMockCloudResourcesEventWatcher(ctrl *gomock.Controller) *MockCloudResourcesEventWatcher {
	mock := &MockCloudResourcesEventWatcher{ctrl: ctrl}
	mock.recorder = &MockCloudResourcesEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudResourcesEventWatcher) EXPECT() *MockCloudResourcesEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockCloudResourcesEventWatcher) AddEventHandler(ctx context.Context, h controller.CloudResourcesEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockCloudResourcesEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockCloudResourcesEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockExternalServiceEventHandler is a mock of ExternalServiceEventHandler interface.
type MockExternalServiceEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockExternalServiceEventHandlerMockRecorder
}

// MockExternalServiceEventHandlerMockRecorder is the mock recorder for MockExternalServiceEventHandler.
type MockExternalServiceEventHandlerMockRecorder struct {
	mock *MockExternalServiceEventHandler
}

// NewMockExternalServiceEventHandler creates a new mock instance.
func NewMockExternalServiceEventHandler(ctrl *gomock.Controller) *MockExternalServiceEventHandler {
	mock := &MockExternalServiceEventHandler{ctrl: ctrl}
	mock.recorder = &MockExternalServiceEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalServiceEventHandler) EXPECT() *MockExternalServiceEventHandlerMockRecorder {
	return m.recorder
}

// CreateExternalService mocks base method.
func (m *MockExternalServiceEventHandler) CreateExternalService(obj *v2.ExternalService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExternalService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExternalService indicates an expected call of CreateExternalService.
func (mr *MockExternalServiceEventHandlerMockRecorder) CreateExternalService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalService", reflect.TypeOf((*MockExternalServiceEventHandler)(nil).CreateExternalService), obj)
}

// DeleteExternalService mocks base method.
func (m *MockExternalServiceEventHandler) DeleteExternalService(obj *v2.ExternalService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExternalService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalService indicates an expected call of DeleteExternalService.
func (mr *MockExternalServiceEventHandlerMockRecorder) DeleteExternalService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalService", reflect.TypeOf((*MockExternalServiceEventHandler)(nil).DeleteExternalService), obj)
}

// GenericExternalService mocks base method.
func (m *MockExternalServiceEventHandler) GenericExternalService(obj *v2.ExternalService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericExternalService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericExternalService indicates an expected call of GenericExternalService.
func (mr *MockExternalServiceEventHandlerMockRecorder) GenericExternalService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericExternalService", reflect.TypeOf((*MockExternalServiceEventHandler)(nil).GenericExternalService), obj)
}

// UpdateExternalService mocks base method.
func (m *MockExternalServiceEventHandler) UpdateExternalService(old, new *v2.ExternalService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExternalService", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalService indicates an expected call of UpdateExternalService.
func (mr *MockExternalServiceEventHandlerMockRecorder) UpdateExternalService(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalService", reflect.TypeOf((*MockExternalServiceEventHandler)(nil).UpdateExternalService), old, new)
}

// MockExternalServiceEventWatcher is a mock of ExternalServiceEventWatcher interface.
type MockExternalServiceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockExternalServiceEventWatcherMockRecorder
}

// MockExternalServiceEventWatcherMockRecorder is the mock recorder for MockExternalServiceEventWatcher.
type MockExternalServiceEventWatcherMockRecorder struct {
	mock *MockExternalServiceEventWatcher
}

// NewMockExternalServiceEventWatcher creates a new mock instance.
func NewMockExternalServiceEventWatcher(ctrl *gomock.Controller) *MockExternalServiceEventWatcher {
	mock := &MockExternalServiceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockExternalServiceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalServiceEventWatcher) EXPECT() *MockExternalServiceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockExternalServiceEventWatcher) AddEventHandler(ctx context.Context, h controller.ExternalServiceEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockExternalServiceEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockExternalServiceEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockExternalEndpointEventHandler is a mock of ExternalEndpointEventHandler interface.
type MockExternalEndpointEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockExternalEndpointEventHandlerMockRecorder
}

// MockExternalEndpointEventHandlerMockRecorder is the mock recorder for MockExternalEndpointEventHandler.
type MockExternalEndpointEventHandlerMockRecorder struct {
	mock *MockExternalEndpointEventHandler
}

// NewMockExternalEndpointEventHandler creates a new mock instance.
func NewMockExternalEndpointEventHandler(ctrl *gomock.Controller) *MockExternalEndpointEventHandler {
	mock := &MockExternalEndpointEventHandler{ctrl: ctrl}
	mock.recorder = &MockExternalEndpointEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalEndpointEventHandler) EXPECT() *MockExternalEndpointEventHandlerMockRecorder {
	return m.recorder
}

// CreateExternalEndpoint mocks base method.
func (m *MockExternalEndpointEventHandler) CreateExternalEndpoint(obj *v2.ExternalEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExternalEndpoint", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExternalEndpoint indicates an expected call of CreateExternalEndpoint.
func (mr *MockExternalEndpointEventHandlerMockRecorder) CreateExternalEndpoint(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalEndpoint", reflect.TypeOf((*MockExternalEndpointEventHandler)(nil).CreateExternalEndpoint), obj)
}

// DeleteExternalEndpoint mocks base method.
func (m *MockExternalEndpointEventHandler) DeleteExternalEndpoint(obj *v2.ExternalEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExternalEndpoint", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalEndpoint indicates an expected call of DeleteExternalEndpoint.
func (mr *MockExternalEndpointEventHandlerMockRecorder) DeleteExternalEndpoint(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalEndpoint", reflect.TypeOf((*MockExternalEndpointEventHandler)(nil).DeleteExternalEndpoint), obj)
}

// GenericExternalEndpoint mocks base method.
func (m *MockExternalEndpointEventHandler) GenericExternalEndpoint(obj *v2.ExternalEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericExternalEndpoint", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericExternalEndpoint indicates an expected call of GenericExternalEndpoint.
func (mr *MockExternalEndpointEventHandlerMockRecorder) GenericExternalEndpoint(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericExternalEndpoint", reflect.TypeOf((*MockExternalEndpointEventHandler)(nil).GenericExternalEndpoint), obj)
}

// UpdateExternalEndpoint mocks base method.
func (m *MockExternalEndpointEventHandler) UpdateExternalEndpoint(old, new *v2.ExternalEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExternalEndpoint", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalEndpoint indicates an expected call of UpdateExternalEndpoint.
func (mr *MockExternalEndpointEventHandlerMockRecorder) UpdateExternalEndpoint(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalEndpoint", reflect.TypeOf((*MockExternalEndpointEventHandler)(nil).UpdateExternalEndpoint), old, new)
}

// MockExternalEndpointEventWatcher is a mock of ExternalEndpointEventWatcher interface.
type MockExternalEndpointEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockExternalEndpointEventWatcherMockRecorder
}

// MockExternalEndpointEventWatcherMockRecorder is the mock recorder for MockExternalEndpointEventWatcher.
type MockExternalEndpointEventWatcherMockRecorder struct {
	mock *MockExternalEndpointEventWatcher
}

// NewMockExternalEndpointEventWatcher creates a new mock instance.
func NewMockExternalEndpointEventWatcher(ctrl *gomock.Controller) *MockExternalEndpointEventWatcher {
	mock := &MockExternalEndpointEventWatcher{ctrl: ctrl}
	mock.recorder = &MockExternalEndpointEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalEndpointEventWatcher) EXPECT() *MockExternalEndpointEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockExternalEndpointEventWatcher) AddEventHandler(ctx context.Context, h controller.ExternalEndpointEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockExternalEndpointEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockExternalEndpointEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockRouteTableEventHandler is a mock of RouteTableEventHandler interface.
type MockRouteTableEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTableEventHandlerMockRecorder
}

// MockRouteTableEventHandlerMockRecorder is the mock recorder for MockRouteTableEventHandler.
type MockRouteTableEventHandlerMockRecorder struct {
	mock *MockRouteTableEventHandler
}

// NewMockRouteTableEventHandler creates a new mock instance.
func NewMockRouteTableEventHandler(ctrl *gomock.Controller) *MockRouteTableEventHandler {
	mock := &MockRouteTableEventHandler{ctrl: ctrl}
	mock.recorder = &MockRouteTableEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteTableEventHandler) EXPECT() *MockRouteTableEventHandlerMockRecorder {
	return m.recorder
}

// CreateRouteTable mocks base method.
func (m *MockRouteTableEventHandler) CreateRouteTable(obj *v2.RouteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRouteTable", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRouteTable indicates an expected call of CreateRouteTable.
func (mr *MockRouteTableEventHandlerMockRecorder) CreateRouteTable(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRouteTable", reflect.TypeOf((*MockRouteTableEventHandler)(nil).CreateRouteTable), obj)
}

// DeleteRouteTable mocks base method.
func (m *MockRouteTableEventHandler) DeleteRouteTable(obj *v2.RouteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRouteTable", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRouteTable indicates an expected call of DeleteRouteTable.
func (mr *MockRouteTableEventHandlerMockRecorder) DeleteRouteTable(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRouteTable", reflect.TypeOf((*MockRouteTableEventHandler)(nil).DeleteRouteTable), obj)
}

// GenericRouteTable mocks base method.
func (m *MockRouteTableEventHandler) GenericRouteTable(obj *v2.RouteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericRouteTable", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericRouteTable indicates an expected call of GenericRouteTable.
func (mr *MockRouteTableEventHandlerMockRecorder) GenericRouteTable(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericRouteTable", reflect.TypeOf((*MockRouteTableEventHandler)(nil).GenericRouteTable), obj)
}

// UpdateRouteTable mocks base method.
func (m *MockRouteTableEventHandler) UpdateRouteTable(old, new *v2.RouteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRouteTable", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRouteTable indicates an expected call of UpdateRouteTable.
func (mr *MockRouteTableEventHandlerMockRecorder) UpdateRouteTable(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRouteTable", reflect.TypeOf((*MockRouteTableEventHandler)(nil).UpdateRouteTable), old, new)
}

// MockRouteTableEventWatcher is a mock of RouteTableEventWatcher interface.
type MockRouteTableEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTableEventWatcherMockRecorder
}

// MockRouteTableEventWatcherMockRecorder is the mock recorder for MockRouteTableEventWatcher.
type MockRouteTableEventWatcherMockRecorder struct {
	mock *MockRouteTableEventWatcher
}

// NewMockRouteTableEventWatcher creates a new mock instance.
func NewMockRouteTableEventWatcher(ctrl *gomock.Controller) *MockRouteTableEventWatcher {
	mock := &MockRouteTableEventWatcher{ctrl: ctrl}
	mock.recorder = &MockRouteTableEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteTableEventWatcher) EXPECT() *MockRouteTableEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockRouteTableEventWatcher) AddEventHandler(ctx context.Context, h controller.RouteTableEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockRouteTableEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockRouteTableEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockVirtualDestinationEventHandler is a mock of VirtualDestinationEventHandler interface.
type MockVirtualDestinationEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualDestinationEventHandlerMockRecorder
}

// MockVirtualDestinationEventHandlerMockRecorder is the mock recorder for MockVirtualDestinationEventHandler.
type MockVirtualDestinationEventHandlerMockRecorder struct {
	mock *MockVirtualDestinationEventHandler
}

// NewMockVirtualDestinationEventHandler creates a new mock instance.
func NewMockVirtualDestinationEventHandler(ctrl *gomock.Controller) *MockVirtualDestinationEventHandler {
	mock := &MockVirtualDestinationEventHandler{ctrl: ctrl}
	mock.recorder = &MockVirtualDestinationEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualDestinationEventHandler) EXPECT() *MockVirtualDestinationEventHandlerMockRecorder {
	return m.recorder
}

// CreateVirtualDestination mocks base method.
func (m *MockVirtualDestinationEventHandler) CreateVirtualDestination(obj *v2.VirtualDestination) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualDestination", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVirtualDestination indicates an expected call of CreateVirtualDestination.
func (mr *MockVirtualDestinationEventHandlerMockRecorder) CreateVirtualDestination(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualDestination", reflect.TypeOf((*MockVirtualDestinationEventHandler)(nil).CreateVirtualDestination), obj)
}

// DeleteVirtualDestination mocks base method.
func (m *MockVirtualDestinationEventHandler) DeleteVirtualDestination(obj *v2.VirtualDestination) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVirtualDestination", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVirtualDestination indicates an expected call of DeleteVirtualDestination.
func (mr *MockVirtualDestinationEventHandlerMockRecorder) DeleteVirtualDestination(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualDestination", reflect.TypeOf((*MockVirtualDestinationEventHandler)(nil).DeleteVirtualDestination), obj)
}

// GenericVirtualDestination mocks base method.
func (m *MockVirtualDestinationEventHandler) GenericVirtualDestination(obj *v2.VirtualDestination) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericVirtualDestination", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericVirtualDestination indicates an expected call of GenericVirtualDestination.
func (mr *MockVirtualDestinationEventHandlerMockRecorder) GenericVirtualDestination(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericVirtualDestination", reflect.TypeOf((*MockVirtualDestinationEventHandler)(nil).GenericVirtualDestination), obj)
}

// UpdateVirtualDestination mocks base method.
func (m *MockVirtualDestinationEventHandler) UpdateVirtualDestination(old, new *v2.VirtualDestination) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVirtualDestination", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVirtualDestination indicates an expected call of UpdateVirtualDestination.
func (mr *MockVirtualDestinationEventHandlerMockRecorder) UpdateVirtualDestination(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVirtualDestination", reflect.TypeOf((*MockVirtualDestinationEventHandler)(nil).UpdateVirtualDestination), old, new)
}

// MockVirtualDestinationEventWatcher is a mock of VirtualDestinationEventWatcher interface.
type MockVirtualDestinationEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualDestinationEventWatcherMockRecorder
}

// MockVirtualDestinationEventWatcherMockRecorder is the mock recorder for MockVirtualDestinationEventWatcher.
type MockVirtualDestinationEventWatcherMockRecorder struct {
	mock *MockVirtualDestinationEventWatcher
}

// NewMockVirtualDestinationEventWatcher creates a new mock instance.
func NewMockVirtualDestinationEventWatcher(ctrl *gomock.Controller) *MockVirtualDestinationEventWatcher {
	mock := &MockVirtualDestinationEventWatcher{ctrl: ctrl}
	mock.recorder = &MockVirtualDestinationEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualDestinationEventWatcher) EXPECT() *MockVirtualDestinationEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockVirtualDestinationEventWatcher) AddEventHandler(ctx context.Context, h controller.VirtualDestinationEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockVirtualDestinationEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockVirtualDestinationEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockVirtualGatewayEventHandler is a mock of VirtualGatewayEventHandler interface.
type MockVirtualGatewayEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualGatewayEventHandlerMockRecorder
}

// MockVirtualGatewayEventHandlerMockRecorder is the mock recorder for MockVirtualGatewayEventHandler.
type MockVirtualGatewayEventHandlerMockRecorder struct {
	mock *MockVirtualGatewayEventHandler
}

// NewMockVirtualGatewayEventHandler creates a new mock instance.
func NewMockVirtualGatewayEventHandler(ctrl *gomock.Controller) *MockVirtualGatewayEventHandler {
	mock := &MockVirtualGatewayEventHandler{ctrl: ctrl}
	mock.recorder = &MockVirtualGatewayEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualGatewayEventHandler) EXPECT() *MockVirtualGatewayEventHandlerMockRecorder {
	return m.recorder
}

// CreateVirtualGateway mocks base method.
func (m *MockVirtualGatewayEventHandler) CreateVirtualGateway(obj *v2.VirtualGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVirtualGateway indicates an expected call of CreateVirtualGateway.
func (mr *MockVirtualGatewayEventHandlerMockRecorder) CreateVirtualGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualGateway", reflect.TypeOf((*MockVirtualGatewayEventHandler)(nil).CreateVirtualGateway), obj)
}

// DeleteVirtualGateway mocks base method.
func (m *MockVirtualGatewayEventHandler) DeleteVirtualGateway(obj *v2.VirtualGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVirtualGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVirtualGateway indicates an expected call of DeleteVirtualGateway.
func (mr *MockVirtualGatewayEventHandlerMockRecorder) DeleteVirtualGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualGateway", reflect.TypeOf((*MockVirtualGatewayEventHandler)(nil).DeleteVirtualGateway), obj)
}

// GenericVirtualGateway mocks base method.
func (m *MockVirtualGatewayEventHandler) GenericVirtualGateway(obj *v2.VirtualGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericVirtualGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericVirtualGateway indicates an expected call of GenericVirtualGateway.
func (mr *MockVirtualGatewayEventHandlerMockRecorder) GenericVirtualGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericVirtualGateway", reflect.TypeOf((*MockVirtualGatewayEventHandler)(nil).GenericVirtualGateway), obj)
}

// UpdateVirtualGateway mocks base method.
func (m *MockVirtualGatewayEventHandler) UpdateVirtualGateway(old, new *v2.VirtualGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVirtualGateway", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVirtualGateway indicates an expected call of UpdateVirtualGateway.
func (mr *MockVirtualGatewayEventHandlerMockRecorder) UpdateVirtualGateway(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVirtualGateway", reflect.TypeOf((*MockVirtualGatewayEventHandler)(nil).UpdateVirtualGateway), old, new)
}

// MockVirtualGatewayEventWatcher is a mock of VirtualGatewayEventWatcher interface.
type MockVirtualGatewayEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualGatewayEventWatcherMockRecorder
}

// MockVirtualGatewayEventWatcherMockRecorder is the mock recorder for MockVirtualGatewayEventWatcher.
type MockVirtualGatewayEventWatcherMockRecorder struct {
	mock *MockVirtualGatewayEventWatcher
}

// NewMockVirtualGatewayEventWatcher creates a new mock instance.
func NewMockVirtualGatewayEventWatcher(ctrl *gomock.Controller) *MockVirtualGatewayEventWatcher {
	mock := &MockVirtualGatewayEventWatcher{ctrl: ctrl}
	mock.recorder = &MockVirtualGatewayEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualGatewayEventWatcher) EXPECT() *MockVirtualGatewayEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockVirtualGatewayEventWatcher) AddEventHandler(ctx context.Context, h controller.VirtualGatewayEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockVirtualGatewayEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockVirtualGatewayEventWatcher)(nil).AddEventHandler), varargs...)
}
