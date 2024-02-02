// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go
//
// Generated by this command:
//
//	mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go
//
// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1beta1 "github.com/solo-io/solo-apis/pkg/api/graphql.gloo.solo.io/v1beta1"
	controller "github.com/solo-io/solo-apis/pkg/api/graphql.gloo.solo.io/v1beta1/controller"
	gomock "go.uber.org/mock/gomock"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterGraphQLApiReconciler is a mock of MulticlusterGraphQLApiReconciler interface.
type MockMulticlusterGraphQLApiReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGraphQLApiReconcilerMockRecorder
}

// MockMulticlusterGraphQLApiReconcilerMockRecorder is the mock recorder for MockMulticlusterGraphQLApiReconciler.
type MockMulticlusterGraphQLApiReconcilerMockRecorder struct {
	mock *MockMulticlusterGraphQLApiReconciler
}

// NewMockMulticlusterGraphQLApiReconciler creates a new mock instance.
func NewMockMulticlusterGraphQLApiReconciler(ctrl *gomock.Controller) *MockMulticlusterGraphQLApiReconciler {
	mock := &MockMulticlusterGraphQLApiReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGraphQLApiReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGraphQLApiReconciler) EXPECT() *MockMulticlusterGraphQLApiReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGraphQLApi mocks base method.
func (m *MockMulticlusterGraphQLApiReconciler) ReconcileGraphQLApi(clusterName string, obj *v1beta1.GraphQLApi) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGraphQLApi", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGraphQLApi indicates an expected call of ReconcileGraphQLApi.
func (mr *MockMulticlusterGraphQLApiReconcilerMockRecorder) ReconcileGraphQLApi(clusterName, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGraphQLApi", reflect.TypeOf((*MockMulticlusterGraphQLApiReconciler)(nil).ReconcileGraphQLApi), clusterName, obj)
}

// MockMulticlusterGraphQLApiDeletionReconciler is a mock of MulticlusterGraphQLApiDeletionReconciler interface.
type MockMulticlusterGraphQLApiDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGraphQLApiDeletionReconcilerMockRecorder
}

// MockMulticlusterGraphQLApiDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterGraphQLApiDeletionReconciler.
type MockMulticlusterGraphQLApiDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterGraphQLApiDeletionReconciler
}

// NewMockMulticlusterGraphQLApiDeletionReconciler creates a new mock instance.
func NewMockMulticlusterGraphQLApiDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterGraphQLApiDeletionReconciler {
	mock := &MockMulticlusterGraphQLApiDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGraphQLApiDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGraphQLApiDeletionReconciler) EXPECT() *MockMulticlusterGraphQLApiDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGraphQLApiDeletion mocks base method.
func (m *MockMulticlusterGraphQLApiDeletionReconciler) ReconcileGraphQLApiDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGraphQLApiDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileGraphQLApiDeletion indicates an expected call of ReconcileGraphQLApiDeletion.
func (mr *MockMulticlusterGraphQLApiDeletionReconcilerMockRecorder) ReconcileGraphQLApiDeletion(clusterName, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGraphQLApiDeletion", reflect.TypeOf((*MockMulticlusterGraphQLApiDeletionReconciler)(nil).ReconcileGraphQLApiDeletion), clusterName, req)
}

// MockMulticlusterGraphQLApiReconcileLoop is a mock of MulticlusterGraphQLApiReconcileLoop interface.
type MockMulticlusterGraphQLApiReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGraphQLApiReconcileLoopMockRecorder
}

// MockMulticlusterGraphQLApiReconcileLoopMockRecorder is the mock recorder for MockMulticlusterGraphQLApiReconcileLoop.
type MockMulticlusterGraphQLApiReconcileLoopMockRecorder struct {
	mock *MockMulticlusterGraphQLApiReconcileLoop
}

// NewMockMulticlusterGraphQLApiReconcileLoop creates a new mock instance.
func NewMockMulticlusterGraphQLApiReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterGraphQLApiReconcileLoop {
	mock := &MockMulticlusterGraphQLApiReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGraphQLApiReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGraphQLApiReconcileLoop) EXPECT() *MockMulticlusterGraphQLApiReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterGraphQLApiReconciler mocks base method.
func (m *MockMulticlusterGraphQLApiReconcileLoop) AddMulticlusterGraphQLApiReconciler(ctx context.Context, rec controller.MulticlusterGraphQLApiReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterGraphQLApiReconciler", varargs...)
}

// AddMulticlusterGraphQLApiReconciler indicates an expected call of AddMulticlusterGraphQLApiReconciler.
func (mr *MockMulticlusterGraphQLApiReconcileLoopMockRecorder) AddMulticlusterGraphQLApiReconciler(ctx, rec any, predicates ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterGraphQLApiReconciler", reflect.TypeOf((*MockMulticlusterGraphQLApiReconcileLoop)(nil).AddMulticlusterGraphQLApiReconciler), varargs...)
}
