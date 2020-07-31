// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterGatewayReconciler is a mock of MulticlusterGatewayReconciler interface.
type MockMulticlusterGatewayReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGatewayReconcilerMockRecorder
}

// MockMulticlusterGatewayReconcilerMockRecorder is the mock recorder for MockMulticlusterGatewayReconciler.
type MockMulticlusterGatewayReconcilerMockRecorder struct {
	mock *MockMulticlusterGatewayReconciler
}

// NewMockMulticlusterGatewayReconciler creates a new mock instance.
func NewMockMulticlusterGatewayReconciler(ctrl *gomock.Controller) *MockMulticlusterGatewayReconciler {
	mock := &MockMulticlusterGatewayReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGatewayReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGatewayReconciler) EXPECT() *MockMulticlusterGatewayReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGateway mocks base method.
func (m *MockMulticlusterGatewayReconciler) ReconcileGateway(clusterName string, obj *v1.Gateway) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGateway", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGateway indicates an expected call of ReconcileGateway.
func (mr *MockMulticlusterGatewayReconcilerMockRecorder) ReconcileGateway(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGateway", reflect.TypeOf((*MockMulticlusterGatewayReconciler)(nil).ReconcileGateway), clusterName, obj)
}

// MockMulticlusterGatewayDeletionReconciler is a mock of MulticlusterGatewayDeletionReconciler interface.
type MockMulticlusterGatewayDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGatewayDeletionReconcilerMockRecorder
}

// MockMulticlusterGatewayDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterGatewayDeletionReconciler.
type MockMulticlusterGatewayDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterGatewayDeletionReconciler
}

// NewMockMulticlusterGatewayDeletionReconciler creates a new mock instance.
func NewMockMulticlusterGatewayDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterGatewayDeletionReconciler {
	mock := &MockMulticlusterGatewayDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGatewayDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGatewayDeletionReconciler) EXPECT() *MockMulticlusterGatewayDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGatewayDeletion mocks base method.
func (m *MockMulticlusterGatewayDeletionReconciler) ReconcileGatewayDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileGatewayDeletion indicates an expected call of ReconcileGatewayDeletion.
func (mr *MockMulticlusterGatewayDeletionReconcilerMockRecorder) ReconcileGatewayDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayDeletion", reflect.TypeOf((*MockMulticlusterGatewayDeletionReconciler)(nil).ReconcileGatewayDeletion), clusterName, req)
}

// MockMulticlusterGatewayReconcileLoop is a mock of MulticlusterGatewayReconcileLoop interface.
type MockMulticlusterGatewayReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGatewayReconcileLoopMockRecorder
}

// MockMulticlusterGatewayReconcileLoopMockRecorder is the mock recorder for MockMulticlusterGatewayReconcileLoop.
type MockMulticlusterGatewayReconcileLoopMockRecorder struct {
	mock *MockMulticlusterGatewayReconcileLoop
}

// NewMockMulticlusterGatewayReconcileLoop creates a new mock instance.
func NewMockMulticlusterGatewayReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterGatewayReconcileLoop {
	mock := &MockMulticlusterGatewayReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGatewayReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGatewayReconcileLoop) EXPECT() *MockMulticlusterGatewayReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterGatewayReconciler mocks base method.
func (m *MockMulticlusterGatewayReconcileLoop) AddMulticlusterGatewayReconciler(ctx context.Context, rec controller.MulticlusterGatewayReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterGatewayReconciler", varargs...)
}

// AddMulticlusterGatewayReconciler indicates an expected call of AddMulticlusterGatewayReconciler.
func (mr *MockMulticlusterGatewayReconcileLoopMockRecorder) AddMulticlusterGatewayReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterGatewayReconciler", reflect.TypeOf((*MockMulticlusterGatewayReconcileLoop)(nil).AddMulticlusterGatewayReconciler), varargs...)
}

// MockMulticlusterRouteTableReconciler is a mock of MulticlusterRouteTableReconciler interface.
type MockMulticlusterRouteTableReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRouteTableReconcilerMockRecorder
}

// MockMulticlusterRouteTableReconcilerMockRecorder is the mock recorder for MockMulticlusterRouteTableReconciler.
type MockMulticlusterRouteTableReconcilerMockRecorder struct {
	mock *MockMulticlusterRouteTableReconciler
}

// NewMockMulticlusterRouteTableReconciler creates a new mock instance.
func NewMockMulticlusterRouteTableReconciler(ctrl *gomock.Controller) *MockMulticlusterRouteTableReconciler {
	mock := &MockMulticlusterRouteTableReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRouteTableReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRouteTableReconciler) EXPECT() *MockMulticlusterRouteTableReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRouteTable mocks base method.
func (m *MockMulticlusterRouteTableReconciler) ReconcileRouteTable(clusterName string, obj *v1.RouteTable) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRouteTable", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRouteTable indicates an expected call of ReconcileRouteTable.
func (mr *MockMulticlusterRouteTableReconcilerMockRecorder) ReconcileRouteTable(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRouteTable", reflect.TypeOf((*MockMulticlusterRouteTableReconciler)(nil).ReconcileRouteTable), clusterName, obj)
}

// MockMulticlusterRouteTableDeletionReconciler is a mock of MulticlusterRouteTableDeletionReconciler interface.
type MockMulticlusterRouteTableDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRouteTableDeletionReconcilerMockRecorder
}

// MockMulticlusterRouteTableDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterRouteTableDeletionReconciler.
type MockMulticlusterRouteTableDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterRouteTableDeletionReconciler
}

// NewMockMulticlusterRouteTableDeletionReconciler creates a new mock instance.
func NewMockMulticlusterRouteTableDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterRouteTableDeletionReconciler {
	mock := &MockMulticlusterRouteTableDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRouteTableDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRouteTableDeletionReconciler) EXPECT() *MockMulticlusterRouteTableDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRouteTableDeletion mocks base method.
func (m *MockMulticlusterRouteTableDeletionReconciler) ReconcileRouteTableDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRouteTableDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRouteTableDeletion indicates an expected call of ReconcileRouteTableDeletion.
func (mr *MockMulticlusterRouteTableDeletionReconcilerMockRecorder) ReconcileRouteTableDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRouteTableDeletion", reflect.TypeOf((*MockMulticlusterRouteTableDeletionReconciler)(nil).ReconcileRouteTableDeletion), clusterName, req)
}

// MockMulticlusterRouteTableReconcileLoop is a mock of MulticlusterRouteTableReconcileLoop interface.
type MockMulticlusterRouteTableReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRouteTableReconcileLoopMockRecorder
}

// MockMulticlusterRouteTableReconcileLoopMockRecorder is the mock recorder for MockMulticlusterRouteTableReconcileLoop.
type MockMulticlusterRouteTableReconcileLoopMockRecorder struct {
	mock *MockMulticlusterRouteTableReconcileLoop
}

// NewMockMulticlusterRouteTableReconcileLoop creates a new mock instance.
func NewMockMulticlusterRouteTableReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterRouteTableReconcileLoop {
	mock := &MockMulticlusterRouteTableReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRouteTableReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRouteTableReconcileLoop) EXPECT() *MockMulticlusterRouteTableReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterRouteTableReconciler mocks base method.
func (m *MockMulticlusterRouteTableReconcileLoop) AddMulticlusterRouteTableReconciler(ctx context.Context, rec controller.MulticlusterRouteTableReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterRouteTableReconciler", varargs...)
}

// AddMulticlusterRouteTableReconciler indicates an expected call of AddMulticlusterRouteTableReconciler.
func (mr *MockMulticlusterRouteTableReconcileLoopMockRecorder) AddMulticlusterRouteTableReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterRouteTableReconciler", reflect.TypeOf((*MockMulticlusterRouteTableReconcileLoop)(nil).AddMulticlusterRouteTableReconciler), varargs...)
}

// MockMulticlusterVirtualServiceReconciler is a mock of MulticlusterVirtualServiceReconciler interface.
type MockMulticlusterVirtualServiceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualServiceReconcilerMockRecorder
}

// MockMulticlusterVirtualServiceReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualServiceReconciler.
type MockMulticlusterVirtualServiceReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualServiceReconciler
}

// NewMockMulticlusterVirtualServiceReconciler creates a new mock instance.
func NewMockMulticlusterVirtualServiceReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualServiceReconciler {
	mock := &MockMulticlusterVirtualServiceReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualServiceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualServiceReconciler) EXPECT() *MockMulticlusterVirtualServiceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualService mocks base method.
func (m *MockMulticlusterVirtualServiceReconciler) ReconcileVirtualService(clusterName string, obj *v1.VirtualService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualService", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualService indicates an expected call of ReconcileVirtualService.
func (mr *MockMulticlusterVirtualServiceReconcilerMockRecorder) ReconcileVirtualService(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualService", reflect.TypeOf((*MockMulticlusterVirtualServiceReconciler)(nil).ReconcileVirtualService), clusterName, obj)
}

// MockMulticlusterVirtualServiceDeletionReconciler is a mock of MulticlusterVirtualServiceDeletionReconciler interface.
type MockMulticlusterVirtualServiceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder
}

// MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualServiceDeletionReconciler.
type MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualServiceDeletionReconciler
}

// NewMockMulticlusterVirtualServiceDeletionReconciler creates a new mock instance.
func NewMockMulticlusterVirtualServiceDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualServiceDeletionReconciler {
	mock := &MockMulticlusterVirtualServiceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualServiceDeletionReconciler) EXPECT() *MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualServiceDeletion mocks base method.
func (m *MockMulticlusterVirtualServiceDeletionReconciler) ReconcileVirtualServiceDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualServiceDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualServiceDeletion indicates an expected call of ReconcileVirtualServiceDeletion.
func (mr *MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder) ReconcileVirtualServiceDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualServiceDeletion", reflect.TypeOf((*MockMulticlusterVirtualServiceDeletionReconciler)(nil).ReconcileVirtualServiceDeletion), clusterName, req)
}

// MockMulticlusterVirtualServiceReconcileLoop is a mock of MulticlusterVirtualServiceReconcileLoop interface.
type MockMulticlusterVirtualServiceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualServiceReconcileLoopMockRecorder
}

// MockMulticlusterVirtualServiceReconcileLoopMockRecorder is the mock recorder for MockMulticlusterVirtualServiceReconcileLoop.
type MockMulticlusterVirtualServiceReconcileLoopMockRecorder struct {
	mock *MockMulticlusterVirtualServiceReconcileLoop
}

// NewMockMulticlusterVirtualServiceReconcileLoop creates a new mock instance.
func NewMockMulticlusterVirtualServiceReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterVirtualServiceReconcileLoop {
	mock := &MockMulticlusterVirtualServiceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualServiceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualServiceReconcileLoop) EXPECT() *MockMulticlusterVirtualServiceReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterVirtualServiceReconciler mocks base method.
func (m *MockMulticlusterVirtualServiceReconcileLoop) AddMulticlusterVirtualServiceReconciler(ctx context.Context, rec controller.MulticlusterVirtualServiceReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterVirtualServiceReconciler", varargs...)
}

// AddMulticlusterVirtualServiceReconciler indicates an expected call of AddMulticlusterVirtualServiceReconciler.
func (mr *MockMulticlusterVirtualServiceReconcileLoopMockRecorder) AddMulticlusterVirtualServiceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterVirtualServiceReconciler", reflect.TypeOf((*MockMulticlusterVirtualServiceReconcileLoop)(nil).AddMulticlusterVirtualServiceReconciler), varargs...)
}
