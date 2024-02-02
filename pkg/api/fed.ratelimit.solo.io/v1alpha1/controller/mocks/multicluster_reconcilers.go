// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/fed.ratelimit.solo.io/v1alpha1"
	controller "github.com/solo-io/solo-apis/pkg/api/fed.ratelimit.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterFederatedRateLimitConfigReconciler is a mock of MulticlusterFederatedRateLimitConfigReconciler interface.
type MockMulticlusterFederatedRateLimitConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedRateLimitConfigReconcilerMockRecorder
}

// MockMulticlusterFederatedRateLimitConfigReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedRateLimitConfigReconciler.
type MockMulticlusterFederatedRateLimitConfigReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedRateLimitConfigReconciler
}

// NewMockMulticlusterFederatedRateLimitConfigReconciler creates a new mock instance.
func NewMockMulticlusterFederatedRateLimitConfigReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedRateLimitConfigReconciler {
	mock := &MockMulticlusterFederatedRateLimitConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedRateLimitConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedRateLimitConfigReconciler) EXPECT() *MockMulticlusterFederatedRateLimitConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedRateLimitConfig mocks base method.
func (m *MockMulticlusterFederatedRateLimitConfigReconciler) ReconcileFederatedRateLimitConfig(clusterName string, obj *v1alpha1.FederatedRateLimitConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedRateLimitConfig", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedRateLimitConfig indicates an expected call of ReconcileFederatedRateLimitConfig.
func (mr *MockMulticlusterFederatedRateLimitConfigReconcilerMockRecorder) ReconcileFederatedRateLimitConfig(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedRateLimitConfig", reflect.TypeOf((*MockMulticlusterFederatedRateLimitConfigReconciler)(nil).ReconcileFederatedRateLimitConfig), clusterName, obj)
}

// MockMulticlusterFederatedRateLimitConfigDeletionReconciler is a mock of MulticlusterFederatedRateLimitConfigDeletionReconciler interface.
type MockMulticlusterFederatedRateLimitConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedRateLimitConfigDeletionReconcilerMockRecorder
}

// MockMulticlusterFederatedRateLimitConfigDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedRateLimitConfigDeletionReconciler.
type MockMulticlusterFederatedRateLimitConfigDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedRateLimitConfigDeletionReconciler
}

// NewMockMulticlusterFederatedRateLimitConfigDeletionReconciler creates a new mock instance.
func NewMockMulticlusterFederatedRateLimitConfigDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedRateLimitConfigDeletionReconciler {
	mock := &MockMulticlusterFederatedRateLimitConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedRateLimitConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedRateLimitConfigDeletionReconciler) EXPECT() *MockMulticlusterFederatedRateLimitConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedRateLimitConfigDeletion mocks base method.
func (m *MockMulticlusterFederatedRateLimitConfigDeletionReconciler) ReconcileFederatedRateLimitConfigDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedRateLimitConfigDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedRateLimitConfigDeletion indicates an expected call of ReconcileFederatedRateLimitConfigDeletion.
func (mr *MockMulticlusterFederatedRateLimitConfigDeletionReconcilerMockRecorder) ReconcileFederatedRateLimitConfigDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedRateLimitConfigDeletion", reflect.TypeOf((*MockMulticlusterFederatedRateLimitConfigDeletionReconciler)(nil).ReconcileFederatedRateLimitConfigDeletion), clusterName, req)
}

// MockMulticlusterFederatedRateLimitConfigReconcileLoop is a mock of MulticlusterFederatedRateLimitConfigReconcileLoop interface.
type MockMulticlusterFederatedRateLimitConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedRateLimitConfigReconcileLoopMockRecorder
}

// MockMulticlusterFederatedRateLimitConfigReconcileLoopMockRecorder is the mock recorder for MockMulticlusterFederatedRateLimitConfigReconcileLoop.
type MockMulticlusterFederatedRateLimitConfigReconcileLoopMockRecorder struct {
	mock *MockMulticlusterFederatedRateLimitConfigReconcileLoop
}

// NewMockMulticlusterFederatedRateLimitConfigReconcileLoop creates a new mock instance.
func NewMockMulticlusterFederatedRateLimitConfigReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterFederatedRateLimitConfigReconcileLoop {
	mock := &MockMulticlusterFederatedRateLimitConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedRateLimitConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedRateLimitConfigReconcileLoop) EXPECT() *MockMulticlusterFederatedRateLimitConfigReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterFederatedRateLimitConfigReconciler mocks base method.
func (m *MockMulticlusterFederatedRateLimitConfigReconcileLoop) AddMulticlusterFederatedRateLimitConfigReconciler(ctx context.Context, rec controller.MulticlusterFederatedRateLimitConfigReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterFederatedRateLimitConfigReconciler", varargs...)
}

// AddMulticlusterFederatedRateLimitConfigReconciler indicates an expected call of AddMulticlusterFederatedRateLimitConfigReconciler.
func (mr *MockMulticlusterFederatedRateLimitConfigReconcileLoopMockRecorder) AddMulticlusterFederatedRateLimitConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterFederatedRateLimitConfigReconciler", reflect.TypeOf((*MockMulticlusterFederatedRateLimitConfigReconcileLoop)(nil).AddMulticlusterFederatedRateLimitConfigReconciler), varargs...)
}
