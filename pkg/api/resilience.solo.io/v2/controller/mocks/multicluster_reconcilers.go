// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v2 "github.com/solo-io/solo-apis/pkg/api/resilience.solo.io/v2"
	controller "github.com/solo-io/solo-apis/pkg/api/resilience.solo.io/v2/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterFailoverPolicyReconciler is a mock of MulticlusterFailoverPolicyReconciler interface
type MockMulticlusterFailoverPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFailoverPolicyReconcilerMockRecorder
}

// MockMulticlusterFailoverPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterFailoverPolicyReconciler
type MockMulticlusterFailoverPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterFailoverPolicyReconciler
}

// NewMockMulticlusterFailoverPolicyReconciler creates a new mock instance
func NewMockMulticlusterFailoverPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterFailoverPolicyReconciler {
	mock := &MockMulticlusterFailoverPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFailoverPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFailoverPolicyReconciler) EXPECT() *MockMulticlusterFailoverPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFailoverPolicy mocks base method
func (m *MockMulticlusterFailoverPolicyReconciler) ReconcileFailoverPolicy(clusterName string, obj *v2.FailoverPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFailoverPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFailoverPolicy indicates an expected call of ReconcileFailoverPolicy
func (mr *MockMulticlusterFailoverPolicyReconcilerMockRecorder) ReconcileFailoverPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFailoverPolicy", reflect.TypeOf((*MockMulticlusterFailoverPolicyReconciler)(nil).ReconcileFailoverPolicy), clusterName, obj)
}

// MockMulticlusterFailoverPolicyDeletionReconciler is a mock of MulticlusterFailoverPolicyDeletionReconciler interface
type MockMulticlusterFailoverPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFailoverPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterFailoverPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterFailoverPolicyDeletionReconciler
type MockMulticlusterFailoverPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterFailoverPolicyDeletionReconciler
}

// NewMockMulticlusterFailoverPolicyDeletionReconciler creates a new mock instance
func NewMockMulticlusterFailoverPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterFailoverPolicyDeletionReconciler {
	mock := &MockMulticlusterFailoverPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFailoverPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFailoverPolicyDeletionReconciler) EXPECT() *MockMulticlusterFailoverPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFailoverPolicyDeletion mocks base method
func (m *MockMulticlusterFailoverPolicyDeletionReconciler) ReconcileFailoverPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFailoverPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFailoverPolicyDeletion indicates an expected call of ReconcileFailoverPolicyDeletion
func (mr *MockMulticlusterFailoverPolicyDeletionReconcilerMockRecorder) ReconcileFailoverPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFailoverPolicyDeletion", reflect.TypeOf((*MockMulticlusterFailoverPolicyDeletionReconciler)(nil).ReconcileFailoverPolicyDeletion), clusterName, req)
}

// MockMulticlusterFailoverPolicyReconcileLoop is a mock of MulticlusterFailoverPolicyReconcileLoop interface
type MockMulticlusterFailoverPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFailoverPolicyReconcileLoopMockRecorder
}

// MockMulticlusterFailoverPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterFailoverPolicyReconcileLoop
type MockMulticlusterFailoverPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterFailoverPolicyReconcileLoop
}

// NewMockMulticlusterFailoverPolicyReconcileLoop creates a new mock instance
func NewMockMulticlusterFailoverPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterFailoverPolicyReconcileLoop {
	mock := &MockMulticlusterFailoverPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFailoverPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFailoverPolicyReconcileLoop) EXPECT() *MockMulticlusterFailoverPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterFailoverPolicyReconciler mocks base method
func (m *MockMulticlusterFailoverPolicyReconcileLoop) AddMulticlusterFailoverPolicyReconciler(ctx context.Context, rec controller.MulticlusterFailoverPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterFailoverPolicyReconciler", varargs...)
}

// AddMulticlusterFailoverPolicyReconciler indicates an expected call of AddMulticlusterFailoverPolicyReconciler
func (mr *MockMulticlusterFailoverPolicyReconcileLoopMockRecorder) AddMulticlusterFailoverPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterFailoverPolicyReconciler", reflect.TypeOf((*MockMulticlusterFailoverPolicyReconcileLoop)(nil).AddMulticlusterFailoverPolicyReconciler), varargs...)
}

// MockMulticlusterOutlierDetectionPolicyReconciler is a mock of MulticlusterOutlierDetectionPolicyReconciler interface
type MockMulticlusterOutlierDetectionPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterOutlierDetectionPolicyReconcilerMockRecorder
}

// MockMulticlusterOutlierDetectionPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterOutlierDetectionPolicyReconciler
type MockMulticlusterOutlierDetectionPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterOutlierDetectionPolicyReconciler
}

// NewMockMulticlusterOutlierDetectionPolicyReconciler creates a new mock instance
func NewMockMulticlusterOutlierDetectionPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterOutlierDetectionPolicyReconciler {
	mock := &MockMulticlusterOutlierDetectionPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterOutlierDetectionPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterOutlierDetectionPolicyReconciler) EXPECT() *MockMulticlusterOutlierDetectionPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileOutlierDetectionPolicy mocks base method
func (m *MockMulticlusterOutlierDetectionPolicyReconciler) ReconcileOutlierDetectionPolicy(clusterName string, obj *v2.OutlierDetectionPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileOutlierDetectionPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileOutlierDetectionPolicy indicates an expected call of ReconcileOutlierDetectionPolicy
func (mr *MockMulticlusterOutlierDetectionPolicyReconcilerMockRecorder) ReconcileOutlierDetectionPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileOutlierDetectionPolicy", reflect.TypeOf((*MockMulticlusterOutlierDetectionPolicyReconciler)(nil).ReconcileOutlierDetectionPolicy), clusterName, obj)
}

// MockMulticlusterOutlierDetectionPolicyDeletionReconciler is a mock of MulticlusterOutlierDetectionPolicyDeletionReconciler interface
type MockMulticlusterOutlierDetectionPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterOutlierDetectionPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterOutlierDetectionPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterOutlierDetectionPolicyDeletionReconciler
type MockMulticlusterOutlierDetectionPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterOutlierDetectionPolicyDeletionReconciler
}

// NewMockMulticlusterOutlierDetectionPolicyDeletionReconciler creates a new mock instance
func NewMockMulticlusterOutlierDetectionPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterOutlierDetectionPolicyDeletionReconciler {
	mock := &MockMulticlusterOutlierDetectionPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterOutlierDetectionPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterOutlierDetectionPolicyDeletionReconciler) EXPECT() *MockMulticlusterOutlierDetectionPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileOutlierDetectionPolicyDeletion mocks base method
func (m *MockMulticlusterOutlierDetectionPolicyDeletionReconciler) ReconcileOutlierDetectionPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileOutlierDetectionPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileOutlierDetectionPolicyDeletion indicates an expected call of ReconcileOutlierDetectionPolicyDeletion
func (mr *MockMulticlusterOutlierDetectionPolicyDeletionReconcilerMockRecorder) ReconcileOutlierDetectionPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileOutlierDetectionPolicyDeletion", reflect.TypeOf((*MockMulticlusterOutlierDetectionPolicyDeletionReconciler)(nil).ReconcileOutlierDetectionPolicyDeletion), clusterName, req)
}

// MockMulticlusterOutlierDetectionPolicyReconcileLoop is a mock of MulticlusterOutlierDetectionPolicyReconcileLoop interface
type MockMulticlusterOutlierDetectionPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterOutlierDetectionPolicyReconcileLoopMockRecorder
}

// MockMulticlusterOutlierDetectionPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterOutlierDetectionPolicyReconcileLoop
type MockMulticlusterOutlierDetectionPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterOutlierDetectionPolicyReconcileLoop
}

// NewMockMulticlusterOutlierDetectionPolicyReconcileLoop creates a new mock instance
func NewMockMulticlusterOutlierDetectionPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterOutlierDetectionPolicyReconcileLoop {
	mock := &MockMulticlusterOutlierDetectionPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterOutlierDetectionPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterOutlierDetectionPolicyReconcileLoop) EXPECT() *MockMulticlusterOutlierDetectionPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterOutlierDetectionPolicyReconciler mocks base method
func (m *MockMulticlusterOutlierDetectionPolicyReconcileLoop) AddMulticlusterOutlierDetectionPolicyReconciler(ctx context.Context, rec controller.MulticlusterOutlierDetectionPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterOutlierDetectionPolicyReconciler", varargs...)
}

// AddMulticlusterOutlierDetectionPolicyReconciler indicates an expected call of AddMulticlusterOutlierDetectionPolicyReconciler
func (mr *MockMulticlusterOutlierDetectionPolicyReconcileLoopMockRecorder) AddMulticlusterOutlierDetectionPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterOutlierDetectionPolicyReconciler", reflect.TypeOf((*MockMulticlusterOutlierDetectionPolicyReconcileLoop)(nil).AddMulticlusterOutlierDetectionPolicyReconciler), varargs...)
}

// MockMulticlusterFaultInjectionPolicyReconciler is a mock of MulticlusterFaultInjectionPolicyReconciler interface
type MockMulticlusterFaultInjectionPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFaultInjectionPolicyReconcilerMockRecorder
}

// MockMulticlusterFaultInjectionPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterFaultInjectionPolicyReconciler
type MockMulticlusterFaultInjectionPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterFaultInjectionPolicyReconciler
}

// NewMockMulticlusterFaultInjectionPolicyReconciler creates a new mock instance
func NewMockMulticlusterFaultInjectionPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterFaultInjectionPolicyReconciler {
	mock := &MockMulticlusterFaultInjectionPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFaultInjectionPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFaultInjectionPolicyReconciler) EXPECT() *MockMulticlusterFaultInjectionPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFaultInjectionPolicy mocks base method
func (m *MockMulticlusterFaultInjectionPolicyReconciler) ReconcileFaultInjectionPolicy(clusterName string, obj *v2.FaultInjectionPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFaultInjectionPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFaultInjectionPolicy indicates an expected call of ReconcileFaultInjectionPolicy
func (mr *MockMulticlusterFaultInjectionPolicyReconcilerMockRecorder) ReconcileFaultInjectionPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFaultInjectionPolicy", reflect.TypeOf((*MockMulticlusterFaultInjectionPolicyReconciler)(nil).ReconcileFaultInjectionPolicy), clusterName, obj)
}

// MockMulticlusterFaultInjectionPolicyDeletionReconciler is a mock of MulticlusterFaultInjectionPolicyDeletionReconciler interface
type MockMulticlusterFaultInjectionPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFaultInjectionPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterFaultInjectionPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterFaultInjectionPolicyDeletionReconciler
type MockMulticlusterFaultInjectionPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterFaultInjectionPolicyDeletionReconciler
}

// NewMockMulticlusterFaultInjectionPolicyDeletionReconciler creates a new mock instance
func NewMockMulticlusterFaultInjectionPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterFaultInjectionPolicyDeletionReconciler {
	mock := &MockMulticlusterFaultInjectionPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFaultInjectionPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFaultInjectionPolicyDeletionReconciler) EXPECT() *MockMulticlusterFaultInjectionPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFaultInjectionPolicyDeletion mocks base method
func (m *MockMulticlusterFaultInjectionPolicyDeletionReconciler) ReconcileFaultInjectionPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFaultInjectionPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFaultInjectionPolicyDeletion indicates an expected call of ReconcileFaultInjectionPolicyDeletion
func (mr *MockMulticlusterFaultInjectionPolicyDeletionReconcilerMockRecorder) ReconcileFaultInjectionPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFaultInjectionPolicyDeletion", reflect.TypeOf((*MockMulticlusterFaultInjectionPolicyDeletionReconciler)(nil).ReconcileFaultInjectionPolicyDeletion), clusterName, req)
}

// MockMulticlusterFaultInjectionPolicyReconcileLoop is a mock of MulticlusterFaultInjectionPolicyReconcileLoop interface
type MockMulticlusterFaultInjectionPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFaultInjectionPolicyReconcileLoopMockRecorder
}

// MockMulticlusterFaultInjectionPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterFaultInjectionPolicyReconcileLoop
type MockMulticlusterFaultInjectionPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterFaultInjectionPolicyReconcileLoop
}

// NewMockMulticlusterFaultInjectionPolicyReconcileLoop creates a new mock instance
func NewMockMulticlusterFaultInjectionPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterFaultInjectionPolicyReconcileLoop {
	mock := &MockMulticlusterFaultInjectionPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFaultInjectionPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFaultInjectionPolicyReconcileLoop) EXPECT() *MockMulticlusterFaultInjectionPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterFaultInjectionPolicyReconciler mocks base method
func (m *MockMulticlusterFaultInjectionPolicyReconcileLoop) AddMulticlusterFaultInjectionPolicyReconciler(ctx context.Context, rec controller.MulticlusterFaultInjectionPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterFaultInjectionPolicyReconciler", varargs...)
}

// AddMulticlusterFaultInjectionPolicyReconciler indicates an expected call of AddMulticlusterFaultInjectionPolicyReconciler
func (mr *MockMulticlusterFaultInjectionPolicyReconcileLoopMockRecorder) AddMulticlusterFaultInjectionPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterFaultInjectionPolicyReconciler", reflect.TypeOf((*MockMulticlusterFaultInjectionPolicyReconcileLoop)(nil).AddMulticlusterFaultInjectionPolicyReconciler), varargs...)
}

// MockMulticlusterRetryTimeoutPolicyReconciler is a mock of MulticlusterRetryTimeoutPolicyReconciler interface
type MockMulticlusterRetryTimeoutPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRetryTimeoutPolicyReconcilerMockRecorder
}

// MockMulticlusterRetryTimeoutPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterRetryTimeoutPolicyReconciler
type MockMulticlusterRetryTimeoutPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterRetryTimeoutPolicyReconciler
}

// NewMockMulticlusterRetryTimeoutPolicyReconciler creates a new mock instance
func NewMockMulticlusterRetryTimeoutPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterRetryTimeoutPolicyReconciler {
	mock := &MockMulticlusterRetryTimeoutPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRetryTimeoutPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRetryTimeoutPolicyReconciler) EXPECT() *MockMulticlusterRetryTimeoutPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRetryTimeoutPolicy mocks base method
func (m *MockMulticlusterRetryTimeoutPolicyReconciler) ReconcileRetryTimeoutPolicy(clusterName string, obj *v2.RetryTimeoutPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRetryTimeoutPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRetryTimeoutPolicy indicates an expected call of ReconcileRetryTimeoutPolicy
func (mr *MockMulticlusterRetryTimeoutPolicyReconcilerMockRecorder) ReconcileRetryTimeoutPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRetryTimeoutPolicy", reflect.TypeOf((*MockMulticlusterRetryTimeoutPolicyReconciler)(nil).ReconcileRetryTimeoutPolicy), clusterName, obj)
}

// MockMulticlusterRetryTimeoutPolicyDeletionReconciler is a mock of MulticlusterRetryTimeoutPolicyDeletionReconciler interface
type MockMulticlusterRetryTimeoutPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRetryTimeoutPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterRetryTimeoutPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterRetryTimeoutPolicyDeletionReconciler
type MockMulticlusterRetryTimeoutPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterRetryTimeoutPolicyDeletionReconciler
}

// NewMockMulticlusterRetryTimeoutPolicyDeletionReconciler creates a new mock instance
func NewMockMulticlusterRetryTimeoutPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterRetryTimeoutPolicyDeletionReconciler {
	mock := &MockMulticlusterRetryTimeoutPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRetryTimeoutPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRetryTimeoutPolicyDeletionReconciler) EXPECT() *MockMulticlusterRetryTimeoutPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRetryTimeoutPolicyDeletion mocks base method
func (m *MockMulticlusterRetryTimeoutPolicyDeletionReconciler) ReconcileRetryTimeoutPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRetryTimeoutPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRetryTimeoutPolicyDeletion indicates an expected call of ReconcileRetryTimeoutPolicyDeletion
func (mr *MockMulticlusterRetryTimeoutPolicyDeletionReconcilerMockRecorder) ReconcileRetryTimeoutPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRetryTimeoutPolicyDeletion", reflect.TypeOf((*MockMulticlusterRetryTimeoutPolicyDeletionReconciler)(nil).ReconcileRetryTimeoutPolicyDeletion), clusterName, req)
}

// MockMulticlusterRetryTimeoutPolicyReconcileLoop is a mock of MulticlusterRetryTimeoutPolicyReconcileLoop interface
type MockMulticlusterRetryTimeoutPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRetryTimeoutPolicyReconcileLoopMockRecorder
}

// MockMulticlusterRetryTimeoutPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterRetryTimeoutPolicyReconcileLoop
type MockMulticlusterRetryTimeoutPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterRetryTimeoutPolicyReconcileLoop
}

// NewMockMulticlusterRetryTimeoutPolicyReconcileLoop creates a new mock instance
func NewMockMulticlusterRetryTimeoutPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterRetryTimeoutPolicyReconcileLoop {
	mock := &MockMulticlusterRetryTimeoutPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRetryTimeoutPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRetryTimeoutPolicyReconcileLoop) EXPECT() *MockMulticlusterRetryTimeoutPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterRetryTimeoutPolicyReconciler mocks base method
func (m *MockMulticlusterRetryTimeoutPolicyReconcileLoop) AddMulticlusterRetryTimeoutPolicyReconciler(ctx context.Context, rec controller.MulticlusterRetryTimeoutPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterRetryTimeoutPolicyReconciler", varargs...)
}

// AddMulticlusterRetryTimeoutPolicyReconciler indicates an expected call of AddMulticlusterRetryTimeoutPolicyReconciler
func (mr *MockMulticlusterRetryTimeoutPolicyReconcileLoopMockRecorder) AddMulticlusterRetryTimeoutPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterRetryTimeoutPolicyReconciler", reflect.TypeOf((*MockMulticlusterRetryTimeoutPolicyReconcileLoop)(nil).AddMulticlusterRetryTimeoutPolicyReconciler), varargs...)
}
