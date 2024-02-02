// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/fed.ratelimit.solo.io/v1alpha1"
	controller "github.com/solo-io/solo-apis/pkg/api/fed.ratelimit.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockFederatedRateLimitConfigReconciler is a mock of FederatedRateLimitConfigReconciler interface.
type MockFederatedRateLimitConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRateLimitConfigReconcilerMockRecorder
}

// MockFederatedRateLimitConfigReconcilerMockRecorder is the mock recorder for MockFederatedRateLimitConfigReconciler.
type MockFederatedRateLimitConfigReconcilerMockRecorder struct {
	mock *MockFederatedRateLimitConfigReconciler
}

// NewMockFederatedRateLimitConfigReconciler creates a new mock instance.
func NewMockFederatedRateLimitConfigReconciler(ctrl *gomock.Controller) *MockFederatedRateLimitConfigReconciler {
	mock := &MockFederatedRateLimitConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedRateLimitConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedRateLimitConfigReconciler) EXPECT() *MockFederatedRateLimitConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedRateLimitConfig mocks base method.
func (m *MockFederatedRateLimitConfigReconciler) ReconcileFederatedRateLimitConfig(obj *v1alpha1.FederatedRateLimitConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedRateLimitConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedRateLimitConfig indicates an expected call of ReconcileFederatedRateLimitConfig.
func (mr *MockFederatedRateLimitConfigReconcilerMockRecorder) ReconcileFederatedRateLimitConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigReconciler)(nil).ReconcileFederatedRateLimitConfig), obj)
}

// MockFederatedRateLimitConfigDeletionReconciler is a mock of FederatedRateLimitConfigDeletionReconciler interface.
type MockFederatedRateLimitConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRateLimitConfigDeletionReconcilerMockRecorder
}

// MockFederatedRateLimitConfigDeletionReconcilerMockRecorder is the mock recorder for MockFederatedRateLimitConfigDeletionReconciler.
type MockFederatedRateLimitConfigDeletionReconcilerMockRecorder struct {
	mock *MockFederatedRateLimitConfigDeletionReconciler
}

// NewMockFederatedRateLimitConfigDeletionReconciler creates a new mock instance.
func NewMockFederatedRateLimitConfigDeletionReconciler(ctrl *gomock.Controller) *MockFederatedRateLimitConfigDeletionReconciler {
	mock := &MockFederatedRateLimitConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedRateLimitConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedRateLimitConfigDeletionReconciler) EXPECT() *MockFederatedRateLimitConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedRateLimitConfigDeletion mocks base method.
func (m *MockFederatedRateLimitConfigDeletionReconciler) ReconcileFederatedRateLimitConfigDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedRateLimitConfigDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedRateLimitConfigDeletion indicates an expected call of ReconcileFederatedRateLimitConfigDeletion.
func (mr *MockFederatedRateLimitConfigDeletionReconcilerMockRecorder) ReconcileFederatedRateLimitConfigDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedRateLimitConfigDeletion", reflect.TypeOf((*MockFederatedRateLimitConfigDeletionReconciler)(nil).ReconcileFederatedRateLimitConfigDeletion), req)
}

// MockFederatedRateLimitConfigFinalizer is a mock of FederatedRateLimitConfigFinalizer interface.
type MockFederatedRateLimitConfigFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRateLimitConfigFinalizerMockRecorder
}

// MockFederatedRateLimitConfigFinalizerMockRecorder is the mock recorder for MockFederatedRateLimitConfigFinalizer.
type MockFederatedRateLimitConfigFinalizerMockRecorder struct {
	mock *MockFederatedRateLimitConfigFinalizer
}

// NewMockFederatedRateLimitConfigFinalizer creates a new mock instance.
func NewMockFederatedRateLimitConfigFinalizer(ctrl *gomock.Controller) *MockFederatedRateLimitConfigFinalizer {
	mock := &MockFederatedRateLimitConfigFinalizer{ctrl: ctrl}
	mock.recorder = &MockFederatedRateLimitConfigFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedRateLimitConfigFinalizer) EXPECT() *MockFederatedRateLimitConfigFinalizerMockRecorder {
	return m.recorder
}

// FederatedRateLimitConfigFinalizerName mocks base method.
func (m *MockFederatedRateLimitConfigFinalizer) FederatedRateLimitConfigFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedRateLimitConfigFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FederatedRateLimitConfigFinalizerName indicates an expected call of FederatedRateLimitConfigFinalizerName.
func (mr *MockFederatedRateLimitConfigFinalizerMockRecorder) FederatedRateLimitConfigFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedRateLimitConfigFinalizerName", reflect.TypeOf((*MockFederatedRateLimitConfigFinalizer)(nil).FederatedRateLimitConfigFinalizerName))
}

// FinalizeFederatedRateLimitConfig mocks base method.
func (m *MockFederatedRateLimitConfigFinalizer) FinalizeFederatedRateLimitConfig(obj *v1alpha1.FederatedRateLimitConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeFederatedRateLimitConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeFederatedRateLimitConfig indicates an expected call of FinalizeFederatedRateLimitConfig.
func (mr *MockFederatedRateLimitConfigFinalizerMockRecorder) FinalizeFederatedRateLimitConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigFinalizer)(nil).FinalizeFederatedRateLimitConfig), obj)
}

// ReconcileFederatedRateLimitConfig mocks base method.
func (m *MockFederatedRateLimitConfigFinalizer) ReconcileFederatedRateLimitConfig(obj *v1alpha1.FederatedRateLimitConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedRateLimitConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedRateLimitConfig indicates an expected call of ReconcileFederatedRateLimitConfig.
func (mr *MockFederatedRateLimitConfigFinalizerMockRecorder) ReconcileFederatedRateLimitConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigFinalizer)(nil).ReconcileFederatedRateLimitConfig), obj)
}

// MockFederatedRateLimitConfigReconcileLoop is a mock of FederatedRateLimitConfigReconcileLoop interface.
type MockFederatedRateLimitConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRateLimitConfigReconcileLoopMockRecorder
}

// MockFederatedRateLimitConfigReconcileLoopMockRecorder is the mock recorder for MockFederatedRateLimitConfigReconcileLoop.
type MockFederatedRateLimitConfigReconcileLoopMockRecorder struct {
	mock *MockFederatedRateLimitConfigReconcileLoop
}

// NewMockFederatedRateLimitConfigReconcileLoop creates a new mock instance.
func NewMockFederatedRateLimitConfigReconcileLoop(ctrl *gomock.Controller) *MockFederatedRateLimitConfigReconcileLoop {
	mock := &MockFederatedRateLimitConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockFederatedRateLimitConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedRateLimitConfigReconcileLoop) EXPECT() *MockFederatedRateLimitConfigReconcileLoopMockRecorder {
	return m.recorder
}

// RunFederatedRateLimitConfigReconciler mocks base method.
func (m *MockFederatedRateLimitConfigReconcileLoop) RunFederatedRateLimitConfigReconciler(ctx context.Context, rec controller.FederatedRateLimitConfigReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunFederatedRateLimitConfigReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFederatedRateLimitConfigReconciler indicates an expected call of RunFederatedRateLimitConfigReconciler.
func (mr *MockFederatedRateLimitConfigReconcileLoopMockRecorder) RunFederatedRateLimitConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFederatedRateLimitConfigReconciler", reflect.TypeOf((*MockFederatedRateLimitConfigReconcileLoop)(nil).RunFederatedRateLimitConfigReconciler), varargs...)
}
