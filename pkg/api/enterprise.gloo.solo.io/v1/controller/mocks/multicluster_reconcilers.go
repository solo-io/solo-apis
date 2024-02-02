// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterAuthConfigReconciler is a mock of MulticlusterAuthConfigReconciler interface.
type MockMulticlusterAuthConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAuthConfigReconcilerMockRecorder
}

// MockMulticlusterAuthConfigReconcilerMockRecorder is the mock recorder for MockMulticlusterAuthConfigReconciler.
type MockMulticlusterAuthConfigReconcilerMockRecorder struct {
	mock *MockMulticlusterAuthConfigReconciler
}

// NewMockMulticlusterAuthConfigReconciler creates a new mock instance.
func NewMockMulticlusterAuthConfigReconciler(ctrl *gomock.Controller) *MockMulticlusterAuthConfigReconciler {
	mock := &MockMulticlusterAuthConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAuthConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAuthConfigReconciler) EXPECT() *MockMulticlusterAuthConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAuthConfig mocks base method.
func (m *MockMulticlusterAuthConfigReconciler) ReconcileAuthConfig(clusterName string, obj *v1.AuthConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthConfig", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAuthConfig indicates an expected call of ReconcileAuthConfig.
func (mr *MockMulticlusterAuthConfigReconcilerMockRecorder) ReconcileAuthConfig(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthConfig", reflect.TypeOf((*MockMulticlusterAuthConfigReconciler)(nil).ReconcileAuthConfig), clusterName, obj)
}

// MockMulticlusterAuthConfigDeletionReconciler is a mock of MulticlusterAuthConfigDeletionReconciler interface.
type MockMulticlusterAuthConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAuthConfigDeletionReconcilerMockRecorder
}

// MockMulticlusterAuthConfigDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterAuthConfigDeletionReconciler.
type MockMulticlusterAuthConfigDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterAuthConfigDeletionReconciler
}

// NewMockMulticlusterAuthConfigDeletionReconciler creates a new mock instance.
func NewMockMulticlusterAuthConfigDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterAuthConfigDeletionReconciler {
	mock := &MockMulticlusterAuthConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAuthConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAuthConfigDeletionReconciler) EXPECT() *MockMulticlusterAuthConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAuthConfigDeletion mocks base method.
func (m *MockMulticlusterAuthConfigDeletionReconciler) ReconcileAuthConfigDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthConfigDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAuthConfigDeletion indicates an expected call of ReconcileAuthConfigDeletion.
func (mr *MockMulticlusterAuthConfigDeletionReconcilerMockRecorder) ReconcileAuthConfigDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthConfigDeletion", reflect.TypeOf((*MockMulticlusterAuthConfigDeletionReconciler)(nil).ReconcileAuthConfigDeletion), clusterName, req)
}

// MockMulticlusterAuthConfigReconcileLoop is a mock of MulticlusterAuthConfigReconcileLoop interface.
type MockMulticlusterAuthConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAuthConfigReconcileLoopMockRecorder
}

// MockMulticlusterAuthConfigReconcileLoopMockRecorder is the mock recorder for MockMulticlusterAuthConfigReconcileLoop.
type MockMulticlusterAuthConfigReconcileLoopMockRecorder struct {
	mock *MockMulticlusterAuthConfigReconcileLoop
}

// NewMockMulticlusterAuthConfigReconcileLoop creates a new mock instance.
func NewMockMulticlusterAuthConfigReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterAuthConfigReconcileLoop {
	mock := &MockMulticlusterAuthConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAuthConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAuthConfigReconcileLoop) EXPECT() *MockMulticlusterAuthConfigReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterAuthConfigReconciler mocks base method.
func (m *MockMulticlusterAuthConfigReconcileLoop) AddMulticlusterAuthConfigReconciler(ctx context.Context, rec controller.MulticlusterAuthConfigReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterAuthConfigReconciler", varargs...)
}

// AddMulticlusterAuthConfigReconciler indicates an expected call of AddMulticlusterAuthConfigReconciler.
func (mr *MockMulticlusterAuthConfigReconcileLoopMockRecorder) AddMulticlusterAuthConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterAuthConfigReconciler", reflect.TypeOf((*MockMulticlusterAuthConfigReconcileLoop)(nil).AddMulticlusterAuthConfigReconciler), varargs...)
}
