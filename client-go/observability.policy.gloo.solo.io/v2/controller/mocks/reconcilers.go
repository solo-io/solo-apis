// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"

	v2 "github.com/solo-io/solo-apis/client-go/observability.policy.gloo.solo.io/v2"
	controller "github.com/solo-io/solo-apis/client-go/observability.policy.gloo.solo.io/v2/controller"
)

// MockAccessLogPolicyReconciler is a mock of AccessLogPolicyReconciler interface.
type MockAccessLogPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogPolicyReconcilerMockRecorder
}

// MockAccessLogPolicyReconcilerMockRecorder is the mock recorder for MockAccessLogPolicyReconciler.
type MockAccessLogPolicyReconcilerMockRecorder struct {
	mock *MockAccessLogPolicyReconciler
}

// NewMockAccessLogPolicyReconciler creates a new mock instance.
func NewMockAccessLogPolicyReconciler(ctrl *gomock.Controller) *MockAccessLogPolicyReconciler {
	mock := &MockAccessLogPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockAccessLogPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogPolicyReconciler) EXPECT() *MockAccessLogPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyReconciler) ReconcileAccessLogPolicy(obj *v2.AccessLogPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessLogPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessLogPolicy indicates an expected call of ReconcileAccessLogPolicy.
func (mr *MockAccessLogPolicyReconcilerMockRecorder) ReconcileAccessLogPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyReconciler)(nil).ReconcileAccessLogPolicy), obj)
}

// MockAccessLogPolicyDeletionReconciler is a mock of AccessLogPolicyDeletionReconciler interface.
type MockAccessLogPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogPolicyDeletionReconcilerMockRecorder
}

// MockAccessLogPolicyDeletionReconcilerMockRecorder is the mock recorder for MockAccessLogPolicyDeletionReconciler.
type MockAccessLogPolicyDeletionReconcilerMockRecorder struct {
	mock *MockAccessLogPolicyDeletionReconciler
}

// NewMockAccessLogPolicyDeletionReconciler creates a new mock instance.
func NewMockAccessLogPolicyDeletionReconciler(ctrl *gomock.Controller) *MockAccessLogPolicyDeletionReconciler {
	mock := &MockAccessLogPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockAccessLogPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogPolicyDeletionReconciler) EXPECT() *MockAccessLogPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessLogPolicyDeletion mocks base method.
func (m *MockAccessLogPolicyDeletionReconciler) ReconcileAccessLogPolicyDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessLogPolicyDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAccessLogPolicyDeletion indicates an expected call of ReconcileAccessLogPolicyDeletion.
func (mr *MockAccessLogPolicyDeletionReconcilerMockRecorder) ReconcileAccessLogPolicyDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessLogPolicyDeletion", reflect.TypeOf((*MockAccessLogPolicyDeletionReconciler)(nil).ReconcileAccessLogPolicyDeletion), req)
}

// MockAccessLogPolicyFinalizer is a mock of AccessLogPolicyFinalizer interface.
type MockAccessLogPolicyFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogPolicyFinalizerMockRecorder
}

// MockAccessLogPolicyFinalizerMockRecorder is the mock recorder for MockAccessLogPolicyFinalizer.
type MockAccessLogPolicyFinalizerMockRecorder struct {
	mock *MockAccessLogPolicyFinalizer
}

// NewMockAccessLogPolicyFinalizer creates a new mock instance.
func NewMockAccessLogPolicyFinalizer(ctrl *gomock.Controller) *MockAccessLogPolicyFinalizer {
	mock := &MockAccessLogPolicyFinalizer{ctrl: ctrl}
	mock.recorder = &MockAccessLogPolicyFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogPolicyFinalizer) EXPECT() *MockAccessLogPolicyFinalizerMockRecorder {
	return m.recorder
}

// AccessLogPolicyFinalizerName mocks base method.
func (m *MockAccessLogPolicyFinalizer) AccessLogPolicyFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessLogPolicyFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// AccessLogPolicyFinalizerName indicates an expected call of AccessLogPolicyFinalizerName.
func (mr *MockAccessLogPolicyFinalizerMockRecorder) AccessLogPolicyFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessLogPolicyFinalizerName", reflect.TypeOf((*MockAccessLogPolicyFinalizer)(nil).AccessLogPolicyFinalizerName))
}

// FinalizeAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyFinalizer) FinalizeAccessLogPolicy(obj *v2.AccessLogPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeAccessLogPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeAccessLogPolicy indicates an expected call of FinalizeAccessLogPolicy.
func (mr *MockAccessLogPolicyFinalizerMockRecorder) FinalizeAccessLogPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyFinalizer)(nil).FinalizeAccessLogPolicy), obj)
}

// ReconcileAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyFinalizer) ReconcileAccessLogPolicy(obj *v2.AccessLogPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessLogPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessLogPolicy indicates an expected call of ReconcileAccessLogPolicy.
func (mr *MockAccessLogPolicyFinalizerMockRecorder) ReconcileAccessLogPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyFinalizer)(nil).ReconcileAccessLogPolicy), obj)
}

// MockAccessLogPolicyReconcileLoop is a mock of AccessLogPolicyReconcileLoop interface.
type MockAccessLogPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogPolicyReconcileLoopMockRecorder
}

// MockAccessLogPolicyReconcileLoopMockRecorder is the mock recorder for MockAccessLogPolicyReconcileLoop.
type MockAccessLogPolicyReconcileLoopMockRecorder struct {
	mock *MockAccessLogPolicyReconcileLoop
}

// NewMockAccessLogPolicyReconcileLoop creates a new mock instance.
func NewMockAccessLogPolicyReconcileLoop(ctrl *gomock.Controller) *MockAccessLogPolicyReconcileLoop {
	mock := &MockAccessLogPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockAccessLogPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogPolicyReconcileLoop) EXPECT() *MockAccessLogPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// RunAccessLogPolicyReconciler mocks base method.
func (m *MockAccessLogPolicyReconcileLoop) RunAccessLogPolicyReconciler(ctx context.Context, rec controller.AccessLogPolicyReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunAccessLogPolicyReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunAccessLogPolicyReconciler indicates an expected call of RunAccessLogPolicyReconciler.
func (mr *MockAccessLogPolicyReconcileLoopMockRecorder) RunAccessLogPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAccessLogPolicyReconciler", reflect.TypeOf((*MockAccessLogPolicyReconcileLoop)(nil).RunAccessLogPolicyReconciler), varargs...)
}