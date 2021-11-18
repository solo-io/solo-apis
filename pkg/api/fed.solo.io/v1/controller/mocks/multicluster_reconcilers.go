// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterGlooInstanceReconciler is a mock of MulticlusterGlooInstanceReconciler interface
type MockMulticlusterGlooInstanceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGlooInstanceReconcilerMockRecorder
}

// MockMulticlusterGlooInstanceReconcilerMockRecorder is the mock recorder for MockMulticlusterGlooInstanceReconciler
type MockMulticlusterGlooInstanceReconcilerMockRecorder struct {
	mock *MockMulticlusterGlooInstanceReconciler
}

// NewMockMulticlusterGlooInstanceReconciler creates a new mock instance
func NewMockMulticlusterGlooInstanceReconciler(ctrl *gomock.Controller) *MockMulticlusterGlooInstanceReconciler {
	mock := &MockMulticlusterGlooInstanceReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGlooInstanceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterGlooInstanceReconciler) EXPECT() *MockMulticlusterGlooInstanceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGlooInstance mocks base method
func (m *MockMulticlusterGlooInstanceReconciler) ReconcileGlooInstance(clusterName string, obj *v1.GlooInstance) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGlooInstance", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGlooInstance indicates an expected call of ReconcileGlooInstance
func (mr *MockMulticlusterGlooInstanceReconcilerMockRecorder) ReconcileGlooInstance(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGlooInstance", reflect.TypeOf((*MockMulticlusterGlooInstanceReconciler)(nil).ReconcileGlooInstance), clusterName, obj)
}

// MockMulticlusterGlooInstanceDeletionReconciler is a mock of MulticlusterGlooInstanceDeletionReconciler interface
type MockMulticlusterGlooInstanceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGlooInstanceDeletionReconcilerMockRecorder
}

// MockMulticlusterGlooInstanceDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterGlooInstanceDeletionReconciler
type MockMulticlusterGlooInstanceDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterGlooInstanceDeletionReconciler
}

// NewMockMulticlusterGlooInstanceDeletionReconciler creates a new mock instance
func NewMockMulticlusterGlooInstanceDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterGlooInstanceDeletionReconciler {
	mock := &MockMulticlusterGlooInstanceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGlooInstanceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterGlooInstanceDeletionReconciler) EXPECT() *MockMulticlusterGlooInstanceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGlooInstanceDeletion mocks base method
func (m *MockMulticlusterGlooInstanceDeletionReconciler) ReconcileGlooInstanceDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGlooInstanceDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileGlooInstanceDeletion indicates an expected call of ReconcileGlooInstanceDeletion
func (mr *MockMulticlusterGlooInstanceDeletionReconcilerMockRecorder) ReconcileGlooInstanceDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGlooInstanceDeletion", reflect.TypeOf((*MockMulticlusterGlooInstanceDeletionReconciler)(nil).ReconcileGlooInstanceDeletion), clusterName, req)
}

// MockMulticlusterGlooInstanceReconcileLoop is a mock of MulticlusterGlooInstanceReconcileLoop interface
type MockMulticlusterGlooInstanceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGlooInstanceReconcileLoopMockRecorder
}

// MockMulticlusterGlooInstanceReconcileLoopMockRecorder is the mock recorder for MockMulticlusterGlooInstanceReconcileLoop
type MockMulticlusterGlooInstanceReconcileLoopMockRecorder struct {
	mock *MockMulticlusterGlooInstanceReconcileLoop
}

// NewMockMulticlusterGlooInstanceReconcileLoop creates a new mock instance
func NewMockMulticlusterGlooInstanceReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterGlooInstanceReconcileLoop {
	mock := &MockMulticlusterGlooInstanceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGlooInstanceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterGlooInstanceReconcileLoop) EXPECT() *MockMulticlusterGlooInstanceReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterGlooInstanceReconciler mocks base method
func (m *MockMulticlusterGlooInstanceReconcileLoop) AddMulticlusterGlooInstanceReconciler(ctx context.Context, rec controller.MulticlusterGlooInstanceReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterGlooInstanceReconciler", varargs...)
}

// AddMulticlusterGlooInstanceReconciler indicates an expected call of AddMulticlusterGlooInstanceReconciler
func (mr *MockMulticlusterGlooInstanceReconcileLoopMockRecorder) AddMulticlusterGlooInstanceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterGlooInstanceReconciler", reflect.TypeOf((*MockMulticlusterGlooInstanceReconcileLoop)(nil).AddMulticlusterGlooInstanceReconciler), varargs...)
}

// MockMulticlusterFailoverSchemeReconciler is a mock of MulticlusterFailoverSchemeReconciler interface
type MockMulticlusterFailoverSchemeReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFailoverSchemeReconcilerMockRecorder
}

// MockMulticlusterFailoverSchemeReconcilerMockRecorder is the mock recorder for MockMulticlusterFailoverSchemeReconciler
type MockMulticlusterFailoverSchemeReconcilerMockRecorder struct {
	mock *MockMulticlusterFailoverSchemeReconciler
}

// NewMockMulticlusterFailoverSchemeReconciler creates a new mock instance
func NewMockMulticlusterFailoverSchemeReconciler(ctrl *gomock.Controller) *MockMulticlusterFailoverSchemeReconciler {
	mock := &MockMulticlusterFailoverSchemeReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFailoverSchemeReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFailoverSchemeReconciler) EXPECT() *MockMulticlusterFailoverSchemeReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFailoverScheme mocks base method
func (m *MockMulticlusterFailoverSchemeReconciler) ReconcileFailoverScheme(clusterName string, obj *v1.FailoverScheme) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFailoverScheme", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFailoverScheme indicates an expected call of ReconcileFailoverScheme
func (mr *MockMulticlusterFailoverSchemeReconcilerMockRecorder) ReconcileFailoverScheme(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFailoverScheme", reflect.TypeOf((*MockMulticlusterFailoverSchemeReconciler)(nil).ReconcileFailoverScheme), clusterName, obj)
}

// MockMulticlusterFailoverSchemeDeletionReconciler is a mock of MulticlusterFailoverSchemeDeletionReconciler interface
type MockMulticlusterFailoverSchemeDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFailoverSchemeDeletionReconcilerMockRecorder
}

// MockMulticlusterFailoverSchemeDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterFailoverSchemeDeletionReconciler
type MockMulticlusterFailoverSchemeDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterFailoverSchemeDeletionReconciler
}

// NewMockMulticlusterFailoverSchemeDeletionReconciler creates a new mock instance
func NewMockMulticlusterFailoverSchemeDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterFailoverSchemeDeletionReconciler {
	mock := &MockMulticlusterFailoverSchemeDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFailoverSchemeDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFailoverSchemeDeletionReconciler) EXPECT() *MockMulticlusterFailoverSchemeDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFailoverSchemeDeletion mocks base method
func (m *MockMulticlusterFailoverSchemeDeletionReconciler) ReconcileFailoverSchemeDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFailoverSchemeDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFailoverSchemeDeletion indicates an expected call of ReconcileFailoverSchemeDeletion
func (mr *MockMulticlusterFailoverSchemeDeletionReconcilerMockRecorder) ReconcileFailoverSchemeDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFailoverSchemeDeletion", reflect.TypeOf((*MockMulticlusterFailoverSchemeDeletionReconciler)(nil).ReconcileFailoverSchemeDeletion), clusterName, req)
}

// MockMulticlusterFailoverSchemeReconcileLoop is a mock of MulticlusterFailoverSchemeReconcileLoop interface
type MockMulticlusterFailoverSchemeReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFailoverSchemeReconcileLoopMockRecorder
}

// MockMulticlusterFailoverSchemeReconcileLoopMockRecorder is the mock recorder for MockMulticlusterFailoverSchemeReconcileLoop
type MockMulticlusterFailoverSchemeReconcileLoopMockRecorder struct {
	mock *MockMulticlusterFailoverSchemeReconcileLoop
}

// NewMockMulticlusterFailoverSchemeReconcileLoop creates a new mock instance
func NewMockMulticlusterFailoverSchemeReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterFailoverSchemeReconcileLoop {
	mock := &MockMulticlusterFailoverSchemeReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFailoverSchemeReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFailoverSchemeReconcileLoop) EXPECT() *MockMulticlusterFailoverSchemeReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterFailoverSchemeReconciler mocks base method
func (m *MockMulticlusterFailoverSchemeReconcileLoop) AddMulticlusterFailoverSchemeReconciler(ctx context.Context, rec controller.MulticlusterFailoverSchemeReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterFailoverSchemeReconciler", varargs...)
}

// AddMulticlusterFailoverSchemeReconciler indicates an expected call of AddMulticlusterFailoverSchemeReconciler
func (mr *MockMulticlusterFailoverSchemeReconcileLoopMockRecorder) AddMulticlusterFailoverSchemeReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterFailoverSchemeReconciler", reflect.TypeOf((*MockMulticlusterFailoverSchemeReconcileLoop)(nil).AddMulticlusterFailoverSchemeReconciler), varargs...)
}