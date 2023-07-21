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

	v2alpha1 "github.com/solo-io/solo-apis/client-go/internal.gloo.solo.io/v2alpha1"
	controller "github.com/solo-io/solo-apis/client-go/internal.gloo.solo.io/v2alpha1/controller"
)

// MockSpireRegistrationEntryReconciler is a mock of SpireRegistrationEntryReconciler interface.
type MockSpireRegistrationEntryReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSpireRegistrationEntryReconcilerMockRecorder
}

// MockSpireRegistrationEntryReconcilerMockRecorder is the mock recorder for MockSpireRegistrationEntryReconciler.
type MockSpireRegistrationEntryReconcilerMockRecorder struct {
	mock *MockSpireRegistrationEntryReconciler
}

// NewMockSpireRegistrationEntryReconciler creates a new mock instance.
func NewMockSpireRegistrationEntryReconciler(ctrl *gomock.Controller) *MockSpireRegistrationEntryReconciler {
	mock := &MockSpireRegistrationEntryReconciler{ctrl: ctrl}
	mock.recorder = &MockSpireRegistrationEntryReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpireRegistrationEntryReconciler) EXPECT() *MockSpireRegistrationEntryReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSpireRegistrationEntry mocks base method.
func (m *MockSpireRegistrationEntryReconciler) ReconcileSpireRegistrationEntry(obj *v2alpha1.SpireRegistrationEntry) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSpireRegistrationEntry", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSpireRegistrationEntry indicates an expected call of ReconcileSpireRegistrationEntry.
func (mr *MockSpireRegistrationEntryReconcilerMockRecorder) ReconcileSpireRegistrationEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSpireRegistrationEntry", reflect.TypeOf((*MockSpireRegistrationEntryReconciler)(nil).ReconcileSpireRegistrationEntry), obj)
}

// MockSpireRegistrationEntryDeletionReconciler is a mock of SpireRegistrationEntryDeletionReconciler interface.
type MockSpireRegistrationEntryDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSpireRegistrationEntryDeletionReconcilerMockRecorder
}

// MockSpireRegistrationEntryDeletionReconcilerMockRecorder is the mock recorder for MockSpireRegistrationEntryDeletionReconciler.
type MockSpireRegistrationEntryDeletionReconcilerMockRecorder struct {
	mock *MockSpireRegistrationEntryDeletionReconciler
}

// NewMockSpireRegistrationEntryDeletionReconciler creates a new mock instance.
func NewMockSpireRegistrationEntryDeletionReconciler(ctrl *gomock.Controller) *MockSpireRegistrationEntryDeletionReconciler {
	mock := &MockSpireRegistrationEntryDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockSpireRegistrationEntryDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpireRegistrationEntryDeletionReconciler) EXPECT() *MockSpireRegistrationEntryDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSpireRegistrationEntryDeletion mocks base method.
func (m *MockSpireRegistrationEntryDeletionReconciler) ReconcileSpireRegistrationEntryDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSpireRegistrationEntryDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileSpireRegistrationEntryDeletion indicates an expected call of ReconcileSpireRegistrationEntryDeletion.
func (mr *MockSpireRegistrationEntryDeletionReconcilerMockRecorder) ReconcileSpireRegistrationEntryDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSpireRegistrationEntryDeletion", reflect.TypeOf((*MockSpireRegistrationEntryDeletionReconciler)(nil).ReconcileSpireRegistrationEntryDeletion), req)
}

// MockSpireRegistrationEntryFinalizer is a mock of SpireRegistrationEntryFinalizer interface.
type MockSpireRegistrationEntryFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockSpireRegistrationEntryFinalizerMockRecorder
}

// MockSpireRegistrationEntryFinalizerMockRecorder is the mock recorder for MockSpireRegistrationEntryFinalizer.
type MockSpireRegistrationEntryFinalizerMockRecorder struct {
	mock *MockSpireRegistrationEntryFinalizer
}

// NewMockSpireRegistrationEntryFinalizer creates a new mock instance.
func NewMockSpireRegistrationEntryFinalizer(ctrl *gomock.Controller) *MockSpireRegistrationEntryFinalizer {
	mock := &MockSpireRegistrationEntryFinalizer{ctrl: ctrl}
	mock.recorder = &MockSpireRegistrationEntryFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpireRegistrationEntryFinalizer) EXPECT() *MockSpireRegistrationEntryFinalizerMockRecorder {
	return m.recorder
}

// FinalizeSpireRegistrationEntry mocks base method.
func (m *MockSpireRegistrationEntryFinalizer) FinalizeSpireRegistrationEntry(obj *v2alpha1.SpireRegistrationEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeSpireRegistrationEntry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeSpireRegistrationEntry indicates an expected call of FinalizeSpireRegistrationEntry.
func (mr *MockSpireRegistrationEntryFinalizerMockRecorder) FinalizeSpireRegistrationEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeSpireRegistrationEntry", reflect.TypeOf((*MockSpireRegistrationEntryFinalizer)(nil).FinalizeSpireRegistrationEntry), obj)
}

// ReconcileSpireRegistrationEntry mocks base method.
func (m *MockSpireRegistrationEntryFinalizer) ReconcileSpireRegistrationEntry(obj *v2alpha1.SpireRegistrationEntry) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSpireRegistrationEntry", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSpireRegistrationEntry indicates an expected call of ReconcileSpireRegistrationEntry.
func (mr *MockSpireRegistrationEntryFinalizerMockRecorder) ReconcileSpireRegistrationEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSpireRegistrationEntry", reflect.TypeOf((*MockSpireRegistrationEntryFinalizer)(nil).ReconcileSpireRegistrationEntry), obj)
}

// SpireRegistrationEntryFinalizerName mocks base method.
func (m *MockSpireRegistrationEntryFinalizer) SpireRegistrationEntryFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpireRegistrationEntryFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// SpireRegistrationEntryFinalizerName indicates an expected call of SpireRegistrationEntryFinalizerName.
func (mr *MockSpireRegistrationEntryFinalizerMockRecorder) SpireRegistrationEntryFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpireRegistrationEntryFinalizerName", reflect.TypeOf((*MockSpireRegistrationEntryFinalizer)(nil).SpireRegistrationEntryFinalizerName))
}

// MockSpireRegistrationEntryReconcileLoop is a mock of SpireRegistrationEntryReconcileLoop interface.
type MockSpireRegistrationEntryReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockSpireRegistrationEntryReconcileLoopMockRecorder
}

// MockSpireRegistrationEntryReconcileLoopMockRecorder is the mock recorder for MockSpireRegistrationEntryReconcileLoop.
type MockSpireRegistrationEntryReconcileLoopMockRecorder struct {
	mock *MockSpireRegistrationEntryReconcileLoop
}

// NewMockSpireRegistrationEntryReconcileLoop creates a new mock instance.
func NewMockSpireRegistrationEntryReconcileLoop(ctrl *gomock.Controller) *MockSpireRegistrationEntryReconcileLoop {
	mock := &MockSpireRegistrationEntryReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockSpireRegistrationEntryReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpireRegistrationEntryReconcileLoop) EXPECT() *MockSpireRegistrationEntryReconcileLoopMockRecorder {
	return m.recorder
}

// RunSpireRegistrationEntryReconciler mocks base method.
func (m *MockSpireRegistrationEntryReconcileLoop) RunSpireRegistrationEntryReconciler(ctx context.Context, rec controller.SpireRegistrationEntryReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunSpireRegistrationEntryReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunSpireRegistrationEntryReconciler indicates an expected call of RunSpireRegistrationEntryReconciler.
func (mr *MockSpireRegistrationEntryReconcileLoopMockRecorder) RunSpireRegistrationEntryReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSpireRegistrationEntryReconciler", reflect.TypeOf((*MockSpireRegistrationEntryReconcileLoop)(nil).RunSpireRegistrationEntryReconciler), varargs...)
}
