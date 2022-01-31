// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/certificates.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/certificates.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterIssuedCertificateReconciler is a mock of MulticlusterIssuedCertificateReconciler interface
type MockMulticlusterIssuedCertificateReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterIssuedCertificateReconcilerMockRecorder
}

// MockMulticlusterIssuedCertificateReconcilerMockRecorder is the mock recorder for MockMulticlusterIssuedCertificateReconciler
type MockMulticlusterIssuedCertificateReconcilerMockRecorder struct {
	mock *MockMulticlusterIssuedCertificateReconciler
}

// NewMockMulticlusterIssuedCertificateReconciler creates a new mock instance
func NewMockMulticlusterIssuedCertificateReconciler(ctrl *gomock.Controller) *MockMulticlusterIssuedCertificateReconciler {
	mock := &MockMulticlusterIssuedCertificateReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterIssuedCertificateReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterIssuedCertificateReconciler) EXPECT() *MockMulticlusterIssuedCertificateReconcilerMockRecorder {
	return m.recorder
}

// ReconcileIssuedCertificate mocks base method
func (m *MockMulticlusterIssuedCertificateReconciler) ReconcileIssuedCertificate(clusterName string, obj *v1.IssuedCertificate) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIssuedCertificate", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileIssuedCertificate indicates an expected call of ReconcileIssuedCertificate
func (mr *MockMulticlusterIssuedCertificateReconcilerMockRecorder) ReconcileIssuedCertificate(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIssuedCertificate", reflect.TypeOf((*MockMulticlusterIssuedCertificateReconciler)(nil).ReconcileIssuedCertificate), clusterName, obj)
}

// MockMulticlusterIssuedCertificateDeletionReconciler is a mock of MulticlusterIssuedCertificateDeletionReconciler interface
type MockMulticlusterIssuedCertificateDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterIssuedCertificateDeletionReconcilerMockRecorder
}

// MockMulticlusterIssuedCertificateDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterIssuedCertificateDeletionReconciler
type MockMulticlusterIssuedCertificateDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterIssuedCertificateDeletionReconciler
}

// NewMockMulticlusterIssuedCertificateDeletionReconciler creates a new mock instance
func NewMockMulticlusterIssuedCertificateDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterIssuedCertificateDeletionReconciler {
	mock := &MockMulticlusterIssuedCertificateDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterIssuedCertificateDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterIssuedCertificateDeletionReconciler) EXPECT() *MockMulticlusterIssuedCertificateDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileIssuedCertificateDeletion mocks base method
func (m *MockMulticlusterIssuedCertificateDeletionReconciler) ReconcileIssuedCertificateDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIssuedCertificateDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileIssuedCertificateDeletion indicates an expected call of ReconcileIssuedCertificateDeletion
func (mr *MockMulticlusterIssuedCertificateDeletionReconcilerMockRecorder) ReconcileIssuedCertificateDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIssuedCertificateDeletion", reflect.TypeOf((*MockMulticlusterIssuedCertificateDeletionReconciler)(nil).ReconcileIssuedCertificateDeletion), clusterName, req)
}

// MockMulticlusterIssuedCertificateReconcileLoop is a mock of MulticlusterIssuedCertificateReconcileLoop interface
type MockMulticlusterIssuedCertificateReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterIssuedCertificateReconcileLoopMockRecorder
}

// MockMulticlusterIssuedCertificateReconcileLoopMockRecorder is the mock recorder for MockMulticlusterIssuedCertificateReconcileLoop
type MockMulticlusterIssuedCertificateReconcileLoopMockRecorder struct {
	mock *MockMulticlusterIssuedCertificateReconcileLoop
}

// NewMockMulticlusterIssuedCertificateReconcileLoop creates a new mock instance
func NewMockMulticlusterIssuedCertificateReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterIssuedCertificateReconcileLoop {
	mock := &MockMulticlusterIssuedCertificateReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterIssuedCertificateReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterIssuedCertificateReconcileLoop) EXPECT() *MockMulticlusterIssuedCertificateReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterIssuedCertificateReconciler mocks base method
func (m *MockMulticlusterIssuedCertificateReconcileLoop) AddMulticlusterIssuedCertificateReconciler(ctx context.Context, rec controller.MulticlusterIssuedCertificateReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterIssuedCertificateReconciler", varargs...)
}

// AddMulticlusterIssuedCertificateReconciler indicates an expected call of AddMulticlusterIssuedCertificateReconciler
func (mr *MockMulticlusterIssuedCertificateReconcileLoopMockRecorder) AddMulticlusterIssuedCertificateReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterIssuedCertificateReconciler", reflect.TypeOf((*MockMulticlusterIssuedCertificateReconcileLoop)(nil).AddMulticlusterIssuedCertificateReconciler), varargs...)
}

// MockMulticlusterCertificateRequestReconciler is a mock of MulticlusterCertificateRequestReconciler interface
type MockMulticlusterCertificateRequestReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCertificateRequestReconcilerMockRecorder
}

// MockMulticlusterCertificateRequestReconcilerMockRecorder is the mock recorder for MockMulticlusterCertificateRequestReconciler
type MockMulticlusterCertificateRequestReconcilerMockRecorder struct {
	mock *MockMulticlusterCertificateRequestReconciler
}

// NewMockMulticlusterCertificateRequestReconciler creates a new mock instance
func NewMockMulticlusterCertificateRequestReconciler(ctrl *gomock.Controller) *MockMulticlusterCertificateRequestReconciler {
	mock := &MockMulticlusterCertificateRequestReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCertificateRequestReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCertificateRequestReconciler) EXPECT() *MockMulticlusterCertificateRequestReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCertificateRequest mocks base method
func (m *MockMulticlusterCertificateRequestReconciler) ReconcileCertificateRequest(clusterName string, obj *v1.CertificateRequest) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCertificateRequest", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCertificateRequest indicates an expected call of ReconcileCertificateRequest
func (mr *MockMulticlusterCertificateRequestReconcilerMockRecorder) ReconcileCertificateRequest(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCertificateRequest", reflect.TypeOf((*MockMulticlusterCertificateRequestReconciler)(nil).ReconcileCertificateRequest), clusterName, obj)
}

// MockMulticlusterCertificateRequestDeletionReconciler is a mock of MulticlusterCertificateRequestDeletionReconciler interface
type MockMulticlusterCertificateRequestDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCertificateRequestDeletionReconcilerMockRecorder
}

// MockMulticlusterCertificateRequestDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterCertificateRequestDeletionReconciler
type MockMulticlusterCertificateRequestDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterCertificateRequestDeletionReconciler
}

// NewMockMulticlusterCertificateRequestDeletionReconciler creates a new mock instance
func NewMockMulticlusterCertificateRequestDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterCertificateRequestDeletionReconciler {
	mock := &MockMulticlusterCertificateRequestDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCertificateRequestDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCertificateRequestDeletionReconciler) EXPECT() *MockMulticlusterCertificateRequestDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCertificateRequestDeletion mocks base method
func (m *MockMulticlusterCertificateRequestDeletionReconciler) ReconcileCertificateRequestDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCertificateRequestDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileCertificateRequestDeletion indicates an expected call of ReconcileCertificateRequestDeletion
func (mr *MockMulticlusterCertificateRequestDeletionReconcilerMockRecorder) ReconcileCertificateRequestDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCertificateRequestDeletion", reflect.TypeOf((*MockMulticlusterCertificateRequestDeletionReconciler)(nil).ReconcileCertificateRequestDeletion), clusterName, req)
}

// MockMulticlusterCertificateRequestReconcileLoop is a mock of MulticlusterCertificateRequestReconcileLoop interface
type MockMulticlusterCertificateRequestReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCertificateRequestReconcileLoopMockRecorder
}

// MockMulticlusterCertificateRequestReconcileLoopMockRecorder is the mock recorder for MockMulticlusterCertificateRequestReconcileLoop
type MockMulticlusterCertificateRequestReconcileLoopMockRecorder struct {
	mock *MockMulticlusterCertificateRequestReconcileLoop
}

// NewMockMulticlusterCertificateRequestReconcileLoop creates a new mock instance
func NewMockMulticlusterCertificateRequestReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterCertificateRequestReconcileLoop {
	mock := &MockMulticlusterCertificateRequestReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCertificateRequestReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCertificateRequestReconcileLoop) EXPECT() *MockMulticlusterCertificateRequestReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterCertificateRequestReconciler mocks base method
func (m *MockMulticlusterCertificateRequestReconcileLoop) AddMulticlusterCertificateRequestReconciler(ctx context.Context, rec controller.MulticlusterCertificateRequestReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterCertificateRequestReconciler", varargs...)
}

// AddMulticlusterCertificateRequestReconciler indicates an expected call of AddMulticlusterCertificateRequestReconciler
func (mr *MockMulticlusterCertificateRequestReconcileLoopMockRecorder) AddMulticlusterCertificateRequestReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterCertificateRequestReconciler", reflect.TypeOf((*MockMulticlusterCertificateRequestReconcileLoop)(nil).AddMulticlusterCertificateRequestReconciler), varargs...)
}

// MockMulticlusterPodBounceDirectiveReconciler is a mock of MulticlusterPodBounceDirectiveReconciler interface
type MockMulticlusterPodBounceDirectiveReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPodBounceDirectiveReconcilerMockRecorder
}

// MockMulticlusterPodBounceDirectiveReconcilerMockRecorder is the mock recorder for MockMulticlusterPodBounceDirectiveReconciler
type MockMulticlusterPodBounceDirectiveReconcilerMockRecorder struct {
	mock *MockMulticlusterPodBounceDirectiveReconciler
}

// NewMockMulticlusterPodBounceDirectiveReconciler creates a new mock instance
func NewMockMulticlusterPodBounceDirectiveReconciler(ctrl *gomock.Controller) *MockMulticlusterPodBounceDirectiveReconciler {
	mock := &MockMulticlusterPodBounceDirectiveReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPodBounceDirectiveReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterPodBounceDirectiveReconciler) EXPECT() *MockMulticlusterPodBounceDirectiveReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePodBounceDirective mocks base method
func (m *MockMulticlusterPodBounceDirectiveReconciler) ReconcilePodBounceDirective(clusterName string, obj *v1.PodBounceDirective) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePodBounceDirective", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePodBounceDirective indicates an expected call of ReconcilePodBounceDirective
func (mr *MockMulticlusterPodBounceDirectiveReconcilerMockRecorder) ReconcilePodBounceDirective(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePodBounceDirective", reflect.TypeOf((*MockMulticlusterPodBounceDirectiveReconciler)(nil).ReconcilePodBounceDirective), clusterName, obj)
}

// MockMulticlusterPodBounceDirectiveDeletionReconciler is a mock of MulticlusterPodBounceDirectiveDeletionReconciler interface
type MockMulticlusterPodBounceDirectiveDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPodBounceDirectiveDeletionReconcilerMockRecorder
}

// MockMulticlusterPodBounceDirectiveDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterPodBounceDirectiveDeletionReconciler
type MockMulticlusterPodBounceDirectiveDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterPodBounceDirectiveDeletionReconciler
}

// NewMockMulticlusterPodBounceDirectiveDeletionReconciler creates a new mock instance
func NewMockMulticlusterPodBounceDirectiveDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterPodBounceDirectiveDeletionReconciler {
	mock := &MockMulticlusterPodBounceDirectiveDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPodBounceDirectiveDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterPodBounceDirectiveDeletionReconciler) EXPECT() *MockMulticlusterPodBounceDirectiveDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePodBounceDirectiveDeletion mocks base method
func (m *MockMulticlusterPodBounceDirectiveDeletionReconciler) ReconcilePodBounceDirectiveDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePodBounceDirectiveDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcilePodBounceDirectiveDeletion indicates an expected call of ReconcilePodBounceDirectiveDeletion
func (mr *MockMulticlusterPodBounceDirectiveDeletionReconcilerMockRecorder) ReconcilePodBounceDirectiveDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePodBounceDirectiveDeletion", reflect.TypeOf((*MockMulticlusterPodBounceDirectiveDeletionReconciler)(nil).ReconcilePodBounceDirectiveDeletion), clusterName, req)
}

// MockMulticlusterPodBounceDirectiveReconcileLoop is a mock of MulticlusterPodBounceDirectiveReconcileLoop interface
type MockMulticlusterPodBounceDirectiveReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPodBounceDirectiveReconcileLoopMockRecorder
}

// MockMulticlusterPodBounceDirectiveReconcileLoopMockRecorder is the mock recorder for MockMulticlusterPodBounceDirectiveReconcileLoop
type MockMulticlusterPodBounceDirectiveReconcileLoopMockRecorder struct {
	mock *MockMulticlusterPodBounceDirectiveReconcileLoop
}

// NewMockMulticlusterPodBounceDirectiveReconcileLoop creates a new mock instance
func NewMockMulticlusterPodBounceDirectiveReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterPodBounceDirectiveReconcileLoop {
	mock := &MockMulticlusterPodBounceDirectiveReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPodBounceDirectiveReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterPodBounceDirectiveReconcileLoop) EXPECT() *MockMulticlusterPodBounceDirectiveReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterPodBounceDirectiveReconciler mocks base method
func (m *MockMulticlusterPodBounceDirectiveReconcileLoop) AddMulticlusterPodBounceDirectiveReconciler(ctx context.Context, rec controller.MulticlusterPodBounceDirectiveReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterPodBounceDirectiveReconciler", varargs...)
}

// AddMulticlusterPodBounceDirectiveReconciler indicates an expected call of AddMulticlusterPodBounceDirectiveReconciler
func (mr *MockMulticlusterPodBounceDirectiveReconcileLoopMockRecorder) AddMulticlusterPodBounceDirectiveReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterPodBounceDirectiveReconciler", reflect.TypeOf((*MockMulticlusterPodBounceDirectiveReconcileLoop)(nil).AddMulticlusterPodBounceDirectiveReconciler), varargs...)
}
