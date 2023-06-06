// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2 "github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2"
	controller "github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockWasmDeploymentPolicyReconciler is a mock of WasmDeploymentPolicyReconciler interface.
type MockWasmDeploymentPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockWasmDeploymentPolicyReconcilerMockRecorder
}

// MockWasmDeploymentPolicyReconcilerMockRecorder is the mock recorder for MockWasmDeploymentPolicyReconciler.
type MockWasmDeploymentPolicyReconcilerMockRecorder struct {
	mock *MockWasmDeploymentPolicyReconciler
}

// NewMockWasmDeploymentPolicyReconciler creates a new mock instance.
func NewMockWasmDeploymentPolicyReconciler(ctrl *gomock.Controller) *MockWasmDeploymentPolicyReconciler {
	mock := &MockWasmDeploymentPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockWasmDeploymentPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmDeploymentPolicyReconciler) EXPECT() *MockWasmDeploymentPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileWasmDeploymentPolicy mocks base method.
func (m *MockWasmDeploymentPolicyReconciler) ReconcileWasmDeploymentPolicy(obj *v2.WasmDeploymentPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWasmDeploymentPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileWasmDeploymentPolicy indicates an expected call of ReconcileWasmDeploymentPolicy.
func (mr *MockWasmDeploymentPolicyReconcilerMockRecorder) ReconcileWasmDeploymentPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWasmDeploymentPolicy", reflect.TypeOf((*MockWasmDeploymentPolicyReconciler)(nil).ReconcileWasmDeploymentPolicy), obj)
}

// MockWasmDeploymentPolicyDeletionReconciler is a mock of WasmDeploymentPolicyDeletionReconciler interface.
type MockWasmDeploymentPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockWasmDeploymentPolicyDeletionReconcilerMockRecorder
}

// MockWasmDeploymentPolicyDeletionReconcilerMockRecorder is the mock recorder for MockWasmDeploymentPolicyDeletionReconciler.
type MockWasmDeploymentPolicyDeletionReconcilerMockRecorder struct {
	mock *MockWasmDeploymentPolicyDeletionReconciler
}

// NewMockWasmDeploymentPolicyDeletionReconciler creates a new mock instance.
func NewMockWasmDeploymentPolicyDeletionReconciler(ctrl *gomock.Controller) *MockWasmDeploymentPolicyDeletionReconciler {
	mock := &MockWasmDeploymentPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockWasmDeploymentPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmDeploymentPolicyDeletionReconciler) EXPECT() *MockWasmDeploymentPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileWasmDeploymentPolicyDeletion mocks base method.
func (m *MockWasmDeploymentPolicyDeletionReconciler) ReconcileWasmDeploymentPolicyDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWasmDeploymentPolicyDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileWasmDeploymentPolicyDeletion indicates an expected call of ReconcileWasmDeploymentPolicyDeletion.
func (mr *MockWasmDeploymentPolicyDeletionReconcilerMockRecorder) ReconcileWasmDeploymentPolicyDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWasmDeploymentPolicyDeletion", reflect.TypeOf((*MockWasmDeploymentPolicyDeletionReconciler)(nil).ReconcileWasmDeploymentPolicyDeletion), req)
}

// MockWasmDeploymentPolicyFinalizer is a mock of WasmDeploymentPolicyFinalizer interface.
type MockWasmDeploymentPolicyFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockWasmDeploymentPolicyFinalizerMockRecorder
}

// MockWasmDeploymentPolicyFinalizerMockRecorder is the mock recorder for MockWasmDeploymentPolicyFinalizer.
type MockWasmDeploymentPolicyFinalizerMockRecorder struct {
	mock *MockWasmDeploymentPolicyFinalizer
}

// NewMockWasmDeploymentPolicyFinalizer creates a new mock instance.
func NewMockWasmDeploymentPolicyFinalizer(ctrl *gomock.Controller) *MockWasmDeploymentPolicyFinalizer {
	mock := &MockWasmDeploymentPolicyFinalizer{ctrl: ctrl}
	mock.recorder = &MockWasmDeploymentPolicyFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmDeploymentPolicyFinalizer) EXPECT() *MockWasmDeploymentPolicyFinalizerMockRecorder {
	return m.recorder
}

// FinalizeWasmDeploymentPolicy mocks base method.
func (m *MockWasmDeploymentPolicyFinalizer) FinalizeWasmDeploymentPolicy(obj *v2.WasmDeploymentPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeWasmDeploymentPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeWasmDeploymentPolicy indicates an expected call of FinalizeWasmDeploymentPolicy.
func (mr *MockWasmDeploymentPolicyFinalizerMockRecorder) FinalizeWasmDeploymentPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeWasmDeploymentPolicy", reflect.TypeOf((*MockWasmDeploymentPolicyFinalizer)(nil).FinalizeWasmDeploymentPolicy), obj)
}

// ReconcileWasmDeploymentPolicy mocks base method.
func (m *MockWasmDeploymentPolicyFinalizer) ReconcileWasmDeploymentPolicy(obj *v2.WasmDeploymentPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWasmDeploymentPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileWasmDeploymentPolicy indicates an expected call of ReconcileWasmDeploymentPolicy.
func (mr *MockWasmDeploymentPolicyFinalizerMockRecorder) ReconcileWasmDeploymentPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWasmDeploymentPolicy", reflect.TypeOf((*MockWasmDeploymentPolicyFinalizer)(nil).ReconcileWasmDeploymentPolicy), obj)
}

// WasmDeploymentPolicyFinalizerName mocks base method.
func (m *MockWasmDeploymentPolicyFinalizer) WasmDeploymentPolicyFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WasmDeploymentPolicyFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// WasmDeploymentPolicyFinalizerName indicates an expected call of WasmDeploymentPolicyFinalizerName.
func (mr *MockWasmDeploymentPolicyFinalizerMockRecorder) WasmDeploymentPolicyFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WasmDeploymentPolicyFinalizerName", reflect.TypeOf((*MockWasmDeploymentPolicyFinalizer)(nil).WasmDeploymentPolicyFinalizerName))
}

// MockWasmDeploymentPolicyReconcileLoop is a mock of WasmDeploymentPolicyReconcileLoop interface.
type MockWasmDeploymentPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockWasmDeploymentPolicyReconcileLoopMockRecorder
}

// MockWasmDeploymentPolicyReconcileLoopMockRecorder is the mock recorder for MockWasmDeploymentPolicyReconcileLoop.
type MockWasmDeploymentPolicyReconcileLoopMockRecorder struct {
	mock *MockWasmDeploymentPolicyReconcileLoop
}

// NewMockWasmDeploymentPolicyReconcileLoop creates a new mock instance.
func NewMockWasmDeploymentPolicyReconcileLoop(ctrl *gomock.Controller) *MockWasmDeploymentPolicyReconcileLoop {
	mock := &MockWasmDeploymentPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockWasmDeploymentPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmDeploymentPolicyReconcileLoop) EXPECT() *MockWasmDeploymentPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// RunWasmDeploymentPolicyReconciler mocks base method.
func (m *MockWasmDeploymentPolicyReconcileLoop) RunWasmDeploymentPolicyReconciler(ctx context.Context, rec controller.WasmDeploymentPolicyReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunWasmDeploymentPolicyReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunWasmDeploymentPolicyReconciler indicates an expected call of RunWasmDeploymentPolicyReconciler.
func (mr *MockWasmDeploymentPolicyReconcileLoopMockRecorder) RunWasmDeploymentPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWasmDeploymentPolicyReconciler", reflect.TypeOf((*MockWasmDeploymentPolicyReconcileLoop)(nil).RunWasmDeploymentPolicyReconciler), varargs...)
}
