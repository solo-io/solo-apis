// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v2 "github.com/solo-io/solo-apis/pkg/api/extensions.solo.io/v2"
	controller "github.com/solo-io/solo-apis/pkg/api/extensions.solo.io/v2/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterWasmDeploymentPolicyReconciler is a mock of MulticlusterWasmDeploymentPolicyReconciler interface
type MockMulticlusterWasmDeploymentPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterWasmDeploymentPolicyReconcilerMockRecorder
}

// MockMulticlusterWasmDeploymentPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterWasmDeploymentPolicyReconciler
type MockMulticlusterWasmDeploymentPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterWasmDeploymentPolicyReconciler
}

// NewMockMulticlusterWasmDeploymentPolicyReconciler creates a new mock instance
func NewMockMulticlusterWasmDeploymentPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterWasmDeploymentPolicyReconciler {
	mock := &MockMulticlusterWasmDeploymentPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterWasmDeploymentPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterWasmDeploymentPolicyReconciler) EXPECT() *MockMulticlusterWasmDeploymentPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileWasmDeploymentPolicy mocks base method
func (m *MockMulticlusterWasmDeploymentPolicyReconciler) ReconcileWasmDeploymentPolicy(clusterName string, obj *v2.WasmDeploymentPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWasmDeploymentPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileWasmDeploymentPolicy indicates an expected call of ReconcileWasmDeploymentPolicy
func (mr *MockMulticlusterWasmDeploymentPolicyReconcilerMockRecorder) ReconcileWasmDeploymentPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWasmDeploymentPolicy", reflect.TypeOf((*MockMulticlusterWasmDeploymentPolicyReconciler)(nil).ReconcileWasmDeploymentPolicy), clusterName, obj)
}

// MockMulticlusterWasmDeploymentPolicyDeletionReconciler is a mock of MulticlusterWasmDeploymentPolicyDeletionReconciler interface
type MockMulticlusterWasmDeploymentPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterWasmDeploymentPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterWasmDeploymentPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterWasmDeploymentPolicyDeletionReconciler
type MockMulticlusterWasmDeploymentPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterWasmDeploymentPolicyDeletionReconciler
}

// NewMockMulticlusterWasmDeploymentPolicyDeletionReconciler creates a new mock instance
func NewMockMulticlusterWasmDeploymentPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterWasmDeploymentPolicyDeletionReconciler {
	mock := &MockMulticlusterWasmDeploymentPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterWasmDeploymentPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterWasmDeploymentPolicyDeletionReconciler) EXPECT() *MockMulticlusterWasmDeploymentPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileWasmDeploymentPolicyDeletion mocks base method
func (m *MockMulticlusterWasmDeploymentPolicyDeletionReconciler) ReconcileWasmDeploymentPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWasmDeploymentPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileWasmDeploymentPolicyDeletion indicates an expected call of ReconcileWasmDeploymentPolicyDeletion
func (mr *MockMulticlusterWasmDeploymentPolicyDeletionReconcilerMockRecorder) ReconcileWasmDeploymentPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWasmDeploymentPolicyDeletion", reflect.TypeOf((*MockMulticlusterWasmDeploymentPolicyDeletionReconciler)(nil).ReconcileWasmDeploymentPolicyDeletion), clusterName, req)
}

// MockMulticlusterWasmDeploymentPolicyReconcileLoop is a mock of MulticlusterWasmDeploymentPolicyReconcileLoop interface
type MockMulticlusterWasmDeploymentPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterWasmDeploymentPolicyReconcileLoopMockRecorder
}

// MockMulticlusterWasmDeploymentPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterWasmDeploymentPolicyReconcileLoop
type MockMulticlusterWasmDeploymentPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterWasmDeploymentPolicyReconcileLoop
}

// NewMockMulticlusterWasmDeploymentPolicyReconcileLoop creates a new mock instance
func NewMockMulticlusterWasmDeploymentPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterWasmDeploymentPolicyReconcileLoop {
	mock := &MockMulticlusterWasmDeploymentPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterWasmDeploymentPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterWasmDeploymentPolicyReconcileLoop) EXPECT() *MockMulticlusterWasmDeploymentPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterWasmDeploymentPolicyReconciler mocks base method
func (m *MockMulticlusterWasmDeploymentPolicyReconcileLoop) AddMulticlusterWasmDeploymentPolicyReconciler(ctx context.Context, rec controller.MulticlusterWasmDeploymentPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterWasmDeploymentPolicyReconciler", varargs...)
}

// AddMulticlusterWasmDeploymentPolicyReconciler indicates an expected call of AddMulticlusterWasmDeploymentPolicyReconciler
func (mr *MockMulticlusterWasmDeploymentPolicyReconcileLoopMockRecorder) AddMulticlusterWasmDeploymentPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterWasmDeploymentPolicyReconciler", reflect.TypeOf((*MockMulticlusterWasmDeploymentPolicyReconcileLoop)(nil).AddMulticlusterWasmDeploymentPolicyReconciler), varargs...)
}
