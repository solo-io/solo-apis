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

	v2alpha1 "github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2alpha1"
	controller "github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2alpha1/controller"
)

// MockWaypointLifecycleManagerReconciler is a mock of WaypointLifecycleManagerReconciler interface.
type MockWaypointLifecycleManagerReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerReconcilerMockRecorder
}

// MockWaypointLifecycleManagerReconcilerMockRecorder is the mock recorder for MockWaypointLifecycleManagerReconciler.
type MockWaypointLifecycleManagerReconcilerMockRecorder struct {
	mock *MockWaypointLifecycleManagerReconciler
}

// NewMockWaypointLifecycleManagerReconciler creates a new mock instance.
func NewMockWaypointLifecycleManagerReconciler(ctrl *gomock.Controller) *MockWaypointLifecycleManagerReconciler {
	mock := &MockWaypointLifecycleManagerReconciler{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerReconciler) EXPECT() *MockWaypointLifecycleManagerReconcilerMockRecorder {
	return m.recorder
}

// ReconcileWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerReconciler) ReconcileWaypointLifecycleManager(obj *v2alpha1.WaypointLifecycleManager) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWaypointLifecycleManager", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileWaypointLifecycleManager indicates an expected call of ReconcileWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerReconcilerMockRecorder) ReconcileWaypointLifecycleManager(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerReconciler)(nil).ReconcileWaypointLifecycleManager), obj)
}

// MockWaypointLifecycleManagerDeletionReconciler is a mock of WaypointLifecycleManagerDeletionReconciler interface.
type MockWaypointLifecycleManagerDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerDeletionReconcilerMockRecorder
}

// MockWaypointLifecycleManagerDeletionReconcilerMockRecorder is the mock recorder for MockWaypointLifecycleManagerDeletionReconciler.
type MockWaypointLifecycleManagerDeletionReconcilerMockRecorder struct {
	mock *MockWaypointLifecycleManagerDeletionReconciler
}

// NewMockWaypointLifecycleManagerDeletionReconciler creates a new mock instance.
func NewMockWaypointLifecycleManagerDeletionReconciler(ctrl *gomock.Controller) *MockWaypointLifecycleManagerDeletionReconciler {
	mock := &MockWaypointLifecycleManagerDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerDeletionReconciler) EXPECT() *MockWaypointLifecycleManagerDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileWaypointLifecycleManagerDeletion mocks base method.
func (m *MockWaypointLifecycleManagerDeletionReconciler) ReconcileWaypointLifecycleManagerDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWaypointLifecycleManagerDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileWaypointLifecycleManagerDeletion indicates an expected call of ReconcileWaypointLifecycleManagerDeletion.
func (mr *MockWaypointLifecycleManagerDeletionReconcilerMockRecorder) ReconcileWaypointLifecycleManagerDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWaypointLifecycleManagerDeletion", reflect.TypeOf((*MockWaypointLifecycleManagerDeletionReconciler)(nil).ReconcileWaypointLifecycleManagerDeletion), req)
}

// MockWaypointLifecycleManagerFinalizer is a mock of WaypointLifecycleManagerFinalizer interface.
type MockWaypointLifecycleManagerFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerFinalizerMockRecorder
}

// MockWaypointLifecycleManagerFinalizerMockRecorder is the mock recorder for MockWaypointLifecycleManagerFinalizer.
type MockWaypointLifecycleManagerFinalizerMockRecorder struct {
	mock *MockWaypointLifecycleManagerFinalizer
}

// NewMockWaypointLifecycleManagerFinalizer creates a new mock instance.
func NewMockWaypointLifecycleManagerFinalizer(ctrl *gomock.Controller) *MockWaypointLifecycleManagerFinalizer {
	mock := &MockWaypointLifecycleManagerFinalizer{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerFinalizer) EXPECT() *MockWaypointLifecycleManagerFinalizerMockRecorder {
	return m.recorder
}

// FinalizeWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerFinalizer) FinalizeWaypointLifecycleManager(obj *v2alpha1.WaypointLifecycleManager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeWaypointLifecycleManager", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeWaypointLifecycleManager indicates an expected call of FinalizeWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerFinalizerMockRecorder) FinalizeWaypointLifecycleManager(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerFinalizer)(nil).FinalizeWaypointLifecycleManager), obj)
}

// ReconcileWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerFinalizer) ReconcileWaypointLifecycleManager(obj *v2alpha1.WaypointLifecycleManager) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWaypointLifecycleManager", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileWaypointLifecycleManager indicates an expected call of ReconcileWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerFinalizerMockRecorder) ReconcileWaypointLifecycleManager(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerFinalizer)(nil).ReconcileWaypointLifecycleManager), obj)
}

// WaypointLifecycleManagerFinalizerName mocks base method.
func (m *MockWaypointLifecycleManagerFinalizer) WaypointLifecycleManagerFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaypointLifecycleManagerFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// WaypointLifecycleManagerFinalizerName indicates an expected call of WaypointLifecycleManagerFinalizerName.
func (mr *MockWaypointLifecycleManagerFinalizerMockRecorder) WaypointLifecycleManagerFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaypointLifecycleManagerFinalizerName", reflect.TypeOf((*MockWaypointLifecycleManagerFinalizer)(nil).WaypointLifecycleManagerFinalizerName))
}

// MockWaypointLifecycleManagerReconcileLoop is a mock of WaypointLifecycleManagerReconcileLoop interface.
type MockWaypointLifecycleManagerReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerReconcileLoopMockRecorder
}

// MockWaypointLifecycleManagerReconcileLoopMockRecorder is the mock recorder for MockWaypointLifecycleManagerReconcileLoop.
type MockWaypointLifecycleManagerReconcileLoopMockRecorder struct {
	mock *MockWaypointLifecycleManagerReconcileLoop
}

// NewMockWaypointLifecycleManagerReconcileLoop creates a new mock instance.
func NewMockWaypointLifecycleManagerReconcileLoop(ctrl *gomock.Controller) *MockWaypointLifecycleManagerReconcileLoop {
	mock := &MockWaypointLifecycleManagerReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerReconcileLoop) EXPECT() *MockWaypointLifecycleManagerReconcileLoopMockRecorder {
	return m.recorder
}

// RunWaypointLifecycleManagerReconciler mocks base method.
func (m *MockWaypointLifecycleManagerReconcileLoop) RunWaypointLifecycleManagerReconciler(ctx context.Context, rec controller.WaypointLifecycleManagerReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunWaypointLifecycleManagerReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunWaypointLifecycleManagerReconciler indicates an expected call of RunWaypointLifecycleManagerReconciler.
func (mr *MockWaypointLifecycleManagerReconcileLoopMockRecorder) RunWaypointLifecycleManagerReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWaypointLifecycleManagerReconciler", reflect.TypeOf((*MockWaypointLifecycleManagerReconcileLoop)(nil).RunWaypointLifecycleManagerReconciler), varargs...)
}

// MockInsightsConfigReconciler is a mock of InsightsConfigReconciler interface.
type MockInsightsConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockInsightsConfigReconcilerMockRecorder
}

// MockInsightsConfigReconcilerMockRecorder is the mock recorder for MockInsightsConfigReconciler.
type MockInsightsConfigReconcilerMockRecorder struct {
	mock *MockInsightsConfigReconciler
}

// NewMockInsightsConfigReconciler creates a new mock instance.
func NewMockInsightsConfigReconciler(ctrl *gomock.Controller) *MockInsightsConfigReconciler {
	mock := &MockInsightsConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockInsightsConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInsightsConfigReconciler) EXPECT() *MockInsightsConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileInsightsConfig mocks base method.
func (m *MockInsightsConfigReconciler) ReconcileInsightsConfig(obj *v2alpha1.InsightsConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileInsightsConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileInsightsConfig indicates an expected call of ReconcileInsightsConfig.
func (mr *MockInsightsConfigReconcilerMockRecorder) ReconcileInsightsConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileInsightsConfig", reflect.TypeOf((*MockInsightsConfigReconciler)(nil).ReconcileInsightsConfig), obj)
}

// MockInsightsConfigDeletionReconciler is a mock of InsightsConfigDeletionReconciler interface.
type MockInsightsConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockInsightsConfigDeletionReconcilerMockRecorder
}

// MockInsightsConfigDeletionReconcilerMockRecorder is the mock recorder for MockInsightsConfigDeletionReconciler.
type MockInsightsConfigDeletionReconcilerMockRecorder struct {
	mock *MockInsightsConfigDeletionReconciler
}

// NewMockInsightsConfigDeletionReconciler creates a new mock instance.
func NewMockInsightsConfigDeletionReconciler(ctrl *gomock.Controller) *MockInsightsConfigDeletionReconciler {
	mock := &MockInsightsConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockInsightsConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInsightsConfigDeletionReconciler) EXPECT() *MockInsightsConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileInsightsConfigDeletion mocks base method.
func (m *MockInsightsConfigDeletionReconciler) ReconcileInsightsConfigDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileInsightsConfigDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileInsightsConfigDeletion indicates an expected call of ReconcileInsightsConfigDeletion.
func (mr *MockInsightsConfigDeletionReconcilerMockRecorder) ReconcileInsightsConfigDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileInsightsConfigDeletion", reflect.TypeOf((*MockInsightsConfigDeletionReconciler)(nil).ReconcileInsightsConfigDeletion), req)
}

// MockInsightsConfigFinalizer is a mock of InsightsConfigFinalizer interface.
type MockInsightsConfigFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockInsightsConfigFinalizerMockRecorder
}

// MockInsightsConfigFinalizerMockRecorder is the mock recorder for MockInsightsConfigFinalizer.
type MockInsightsConfigFinalizerMockRecorder struct {
	mock *MockInsightsConfigFinalizer
}

// NewMockInsightsConfigFinalizer creates a new mock instance.
func NewMockInsightsConfigFinalizer(ctrl *gomock.Controller) *MockInsightsConfigFinalizer {
	mock := &MockInsightsConfigFinalizer{ctrl: ctrl}
	mock.recorder = &MockInsightsConfigFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInsightsConfigFinalizer) EXPECT() *MockInsightsConfigFinalizerMockRecorder {
	return m.recorder
}

// FinalizeInsightsConfig mocks base method.
func (m *MockInsightsConfigFinalizer) FinalizeInsightsConfig(obj *v2alpha1.InsightsConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeInsightsConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeInsightsConfig indicates an expected call of FinalizeInsightsConfig.
func (mr *MockInsightsConfigFinalizerMockRecorder) FinalizeInsightsConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeInsightsConfig", reflect.TypeOf((*MockInsightsConfigFinalizer)(nil).FinalizeInsightsConfig), obj)
}

// InsightsConfigFinalizerName mocks base method.
func (m *MockInsightsConfigFinalizer) InsightsConfigFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsightsConfigFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// InsightsConfigFinalizerName indicates an expected call of InsightsConfigFinalizerName.
func (mr *MockInsightsConfigFinalizerMockRecorder) InsightsConfigFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsightsConfigFinalizerName", reflect.TypeOf((*MockInsightsConfigFinalizer)(nil).InsightsConfigFinalizerName))
}

// ReconcileInsightsConfig mocks base method.
func (m *MockInsightsConfigFinalizer) ReconcileInsightsConfig(obj *v2alpha1.InsightsConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileInsightsConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileInsightsConfig indicates an expected call of ReconcileInsightsConfig.
func (mr *MockInsightsConfigFinalizerMockRecorder) ReconcileInsightsConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileInsightsConfig", reflect.TypeOf((*MockInsightsConfigFinalizer)(nil).ReconcileInsightsConfig), obj)
}

// MockInsightsConfigReconcileLoop is a mock of InsightsConfigReconcileLoop interface.
type MockInsightsConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockInsightsConfigReconcileLoopMockRecorder
}

// MockInsightsConfigReconcileLoopMockRecorder is the mock recorder for MockInsightsConfigReconcileLoop.
type MockInsightsConfigReconcileLoopMockRecorder struct {
	mock *MockInsightsConfigReconcileLoop
}

// NewMockInsightsConfigReconcileLoop creates a new mock instance.
func NewMockInsightsConfigReconcileLoop(ctrl *gomock.Controller) *MockInsightsConfigReconcileLoop {
	mock := &MockInsightsConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockInsightsConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInsightsConfigReconcileLoop) EXPECT() *MockInsightsConfigReconcileLoopMockRecorder {
	return m.recorder
}

// RunInsightsConfigReconciler mocks base method.
func (m *MockInsightsConfigReconcileLoop) RunInsightsConfigReconciler(ctx context.Context, rec controller.InsightsConfigReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunInsightsConfigReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunInsightsConfigReconciler indicates an expected call of RunInsightsConfigReconciler.
func (mr *MockInsightsConfigReconcileLoopMockRecorder) RunInsightsConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInsightsConfigReconciler", reflect.TypeOf((*MockInsightsConfigReconcileLoop)(nil).RunInsightsConfigReconciler), varargs...)
}