// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1alpha1"
	controller "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockGraphQLSchemaReconciler is a mock of GraphQLSchemaReconciler interface.
type MockGraphQLSchemaReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLSchemaReconcilerMockRecorder
}

// MockGraphQLSchemaReconcilerMockRecorder is the mock recorder for MockGraphQLSchemaReconciler.
type MockGraphQLSchemaReconcilerMockRecorder struct {
	mock *MockGraphQLSchemaReconciler
}

// NewMockGraphQLSchemaReconciler creates a new mock instance.
func NewMockGraphQLSchemaReconciler(ctrl *gomock.Controller) *MockGraphQLSchemaReconciler {
	mock := &MockGraphQLSchemaReconciler{ctrl: ctrl}
	mock.recorder = &MockGraphQLSchemaReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQLSchemaReconciler) EXPECT() *MockGraphQLSchemaReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGraphQLSchema mocks base method.
func (m *MockGraphQLSchemaReconciler) ReconcileGraphQLSchema(obj *v1alpha1.GraphQLSchema) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGraphQLSchema", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGraphQLSchema indicates an expected call of ReconcileGraphQLSchema.
func (mr *MockGraphQLSchemaReconcilerMockRecorder) ReconcileGraphQLSchema(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGraphQLSchema", reflect.TypeOf((*MockGraphQLSchemaReconciler)(nil).ReconcileGraphQLSchema), obj)
}

// MockGraphQLSchemaDeletionReconciler is a mock of GraphQLSchemaDeletionReconciler interface.
type MockGraphQLSchemaDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLSchemaDeletionReconcilerMockRecorder
}

// MockGraphQLSchemaDeletionReconcilerMockRecorder is the mock recorder for MockGraphQLSchemaDeletionReconciler.
type MockGraphQLSchemaDeletionReconcilerMockRecorder struct {
	mock *MockGraphQLSchemaDeletionReconciler
}

// NewMockGraphQLSchemaDeletionReconciler creates a new mock instance.
func NewMockGraphQLSchemaDeletionReconciler(ctrl *gomock.Controller) *MockGraphQLSchemaDeletionReconciler {
	mock := &MockGraphQLSchemaDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockGraphQLSchemaDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQLSchemaDeletionReconciler) EXPECT() *MockGraphQLSchemaDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGraphQLSchemaDeletion mocks base method.
func (m *MockGraphQLSchemaDeletionReconciler) ReconcileGraphQLSchemaDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGraphQLSchemaDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileGraphQLSchemaDeletion indicates an expected call of ReconcileGraphQLSchemaDeletion.
func (mr *MockGraphQLSchemaDeletionReconcilerMockRecorder) ReconcileGraphQLSchemaDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGraphQLSchemaDeletion", reflect.TypeOf((*MockGraphQLSchemaDeletionReconciler)(nil).ReconcileGraphQLSchemaDeletion), req)
}

// MockGraphQLSchemaFinalizer is a mock of GraphQLSchemaFinalizer interface.
type MockGraphQLSchemaFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLSchemaFinalizerMockRecorder
}

// MockGraphQLSchemaFinalizerMockRecorder is the mock recorder for MockGraphQLSchemaFinalizer.
type MockGraphQLSchemaFinalizerMockRecorder struct {
	mock *MockGraphQLSchemaFinalizer
}

// NewMockGraphQLSchemaFinalizer creates a new mock instance.
func NewMockGraphQLSchemaFinalizer(ctrl *gomock.Controller) *MockGraphQLSchemaFinalizer {
	mock := &MockGraphQLSchemaFinalizer{ctrl: ctrl}
	mock.recorder = &MockGraphQLSchemaFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQLSchemaFinalizer) EXPECT() *MockGraphQLSchemaFinalizerMockRecorder {
	return m.recorder
}

// FinalizeGraphQLSchema mocks base method.
func (m *MockGraphQLSchemaFinalizer) FinalizeGraphQLSchema(obj *v1alpha1.GraphQLSchema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeGraphQLSchema", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeGraphQLSchema indicates an expected call of FinalizeGraphQLSchema.
func (mr *MockGraphQLSchemaFinalizerMockRecorder) FinalizeGraphQLSchema(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeGraphQLSchema", reflect.TypeOf((*MockGraphQLSchemaFinalizer)(nil).FinalizeGraphQLSchema), obj)
}

// GraphQLSchemaFinalizerName mocks base method.
func (m *MockGraphQLSchemaFinalizer) GraphQLSchemaFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GraphQLSchemaFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GraphQLSchemaFinalizerName indicates an expected call of GraphQLSchemaFinalizerName.
func (mr *MockGraphQLSchemaFinalizerMockRecorder) GraphQLSchemaFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GraphQLSchemaFinalizerName", reflect.TypeOf((*MockGraphQLSchemaFinalizer)(nil).GraphQLSchemaFinalizerName))
}

// ReconcileGraphQLSchema mocks base method.
func (m *MockGraphQLSchemaFinalizer) ReconcileGraphQLSchema(obj *v1alpha1.GraphQLSchema) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGraphQLSchema", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGraphQLSchema indicates an expected call of ReconcileGraphQLSchema.
func (mr *MockGraphQLSchemaFinalizerMockRecorder) ReconcileGraphQLSchema(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGraphQLSchema", reflect.TypeOf((*MockGraphQLSchemaFinalizer)(nil).ReconcileGraphQLSchema), obj)
}

// MockGraphQLSchemaReconcileLoop is a mock of GraphQLSchemaReconcileLoop interface.
type MockGraphQLSchemaReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLSchemaReconcileLoopMockRecorder
}

// MockGraphQLSchemaReconcileLoopMockRecorder is the mock recorder for MockGraphQLSchemaReconcileLoop.
type MockGraphQLSchemaReconcileLoopMockRecorder struct {
	mock *MockGraphQLSchemaReconcileLoop
}

// NewMockGraphQLSchemaReconcileLoop creates a new mock instance.
func NewMockGraphQLSchemaReconcileLoop(ctrl *gomock.Controller) *MockGraphQLSchemaReconcileLoop {
	mock := &MockGraphQLSchemaReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockGraphQLSchemaReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQLSchemaReconcileLoop) EXPECT() *MockGraphQLSchemaReconcileLoopMockRecorder {
	return m.recorder
}

// RunGraphQLSchemaReconciler mocks base method.
func (m *MockGraphQLSchemaReconcileLoop) RunGraphQLSchemaReconciler(ctx context.Context, rec controller.GraphQLSchemaReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunGraphQLSchemaReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunGraphQLSchemaReconciler indicates an expected call of RunGraphQLSchemaReconciler.
func (mr *MockGraphQLSchemaReconcileLoopMockRecorder) RunGraphQLSchemaReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunGraphQLSchemaReconciler", reflect.TypeOf((*MockGraphQLSchemaReconcileLoop)(nil).RunGraphQLSchemaReconciler), varargs...)
}
