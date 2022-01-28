// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

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

// MockMulticlusterGraphQLSchemaReconciler is a mock of MulticlusterGraphQLSchemaReconciler interface.
type MockMulticlusterGraphQLSchemaReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGraphQLSchemaReconcilerMockRecorder
}

// MockMulticlusterGraphQLSchemaReconcilerMockRecorder is the mock recorder for MockMulticlusterGraphQLSchemaReconciler.
type MockMulticlusterGraphQLSchemaReconcilerMockRecorder struct {
	mock *MockMulticlusterGraphQLSchemaReconciler
}

// NewMockMulticlusterGraphQLSchemaReconciler creates a new mock instance.
func NewMockMulticlusterGraphQLSchemaReconciler(ctrl *gomock.Controller) *MockMulticlusterGraphQLSchemaReconciler {
	mock := &MockMulticlusterGraphQLSchemaReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGraphQLSchemaReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGraphQLSchemaReconciler) EXPECT() *MockMulticlusterGraphQLSchemaReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGraphQLSchema mocks base method.
func (m *MockMulticlusterGraphQLSchemaReconciler) ReconcileGraphQLSchema(clusterName string, obj *v1alpha1.GraphQLSchema) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGraphQLSchema", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGraphQLSchema indicates an expected call of ReconcileGraphQLSchema.
func (mr *MockMulticlusterGraphQLSchemaReconcilerMockRecorder) ReconcileGraphQLSchema(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGraphQLSchema", reflect.TypeOf((*MockMulticlusterGraphQLSchemaReconciler)(nil).ReconcileGraphQLSchema), clusterName, obj)
}

// MockMulticlusterGraphQLSchemaDeletionReconciler is a mock of MulticlusterGraphQLSchemaDeletionReconciler interface.
type MockMulticlusterGraphQLSchemaDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGraphQLSchemaDeletionReconcilerMockRecorder
}

// MockMulticlusterGraphQLSchemaDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterGraphQLSchemaDeletionReconciler.
type MockMulticlusterGraphQLSchemaDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterGraphQLSchemaDeletionReconciler
}

// NewMockMulticlusterGraphQLSchemaDeletionReconciler creates a new mock instance.
func NewMockMulticlusterGraphQLSchemaDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterGraphQLSchemaDeletionReconciler {
	mock := &MockMulticlusterGraphQLSchemaDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGraphQLSchemaDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGraphQLSchemaDeletionReconciler) EXPECT() *MockMulticlusterGraphQLSchemaDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGraphQLSchemaDeletion mocks base method.
func (m *MockMulticlusterGraphQLSchemaDeletionReconciler) ReconcileGraphQLSchemaDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGraphQLSchemaDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileGraphQLSchemaDeletion indicates an expected call of ReconcileGraphQLSchemaDeletion.
func (mr *MockMulticlusterGraphQLSchemaDeletionReconcilerMockRecorder) ReconcileGraphQLSchemaDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGraphQLSchemaDeletion", reflect.TypeOf((*MockMulticlusterGraphQLSchemaDeletionReconciler)(nil).ReconcileGraphQLSchemaDeletion), clusterName, req)
}

// MockMulticlusterGraphQLSchemaReconcileLoop is a mock of MulticlusterGraphQLSchemaReconcileLoop interface.
type MockMulticlusterGraphQLSchemaReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGraphQLSchemaReconcileLoopMockRecorder
}

// MockMulticlusterGraphQLSchemaReconcileLoopMockRecorder is the mock recorder for MockMulticlusterGraphQLSchemaReconcileLoop.
type MockMulticlusterGraphQLSchemaReconcileLoopMockRecorder struct {
	mock *MockMulticlusterGraphQLSchemaReconcileLoop
}

// NewMockMulticlusterGraphQLSchemaReconcileLoop creates a new mock instance.
func NewMockMulticlusterGraphQLSchemaReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterGraphQLSchemaReconcileLoop {
	mock := &MockMulticlusterGraphQLSchemaReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGraphQLSchemaReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGraphQLSchemaReconcileLoop) EXPECT() *MockMulticlusterGraphQLSchemaReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterGraphQLSchemaReconciler mocks base method.
func (m *MockMulticlusterGraphQLSchemaReconcileLoop) AddMulticlusterGraphQLSchemaReconciler(ctx context.Context, rec controller.MulticlusterGraphQLSchemaReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterGraphQLSchemaReconciler", varargs...)
}

// AddMulticlusterGraphQLSchemaReconciler indicates an expected call of AddMulticlusterGraphQLSchemaReconciler.
func (mr *MockMulticlusterGraphQLSchemaReconcileLoopMockRecorder) AddMulticlusterGraphQLSchemaReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterGraphQLSchemaReconciler", reflect.TypeOf((*MockMulticlusterGraphQLSchemaReconcileLoop)(nil).AddMulticlusterGraphQLSchemaReconciler), varargs...)
}
