// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go
//
// Generated by this command:
//
//	mockgen -source ./reconcilers.go -destination mocks/reconcilers.go
//
// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.gloo.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/fed.gloo.solo.io/v1/controller"
	gomock "go.uber.org/mock/gomock"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockFederatedSettingsReconciler is a mock of FederatedSettingsReconciler interface.
type MockFederatedSettingsReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedSettingsReconcilerMockRecorder
}

// MockFederatedSettingsReconcilerMockRecorder is the mock recorder for MockFederatedSettingsReconciler.
type MockFederatedSettingsReconcilerMockRecorder struct {
	mock *MockFederatedSettingsReconciler
}

// NewMockFederatedSettingsReconciler creates a new mock instance.
func NewMockFederatedSettingsReconciler(ctrl *gomock.Controller) *MockFederatedSettingsReconciler {
	mock := &MockFederatedSettingsReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedSettingsReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedSettingsReconciler) EXPECT() *MockFederatedSettingsReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedSettings mocks base method.
func (m *MockFederatedSettingsReconciler) ReconcileFederatedSettings(obj *v1.FederatedSettings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedSettings", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedSettings indicates an expected call of ReconcileFederatedSettings.
func (mr *MockFederatedSettingsReconcilerMockRecorder) ReconcileFederatedSettings(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedSettings", reflect.TypeOf((*MockFederatedSettingsReconciler)(nil).ReconcileFederatedSettings), obj)
}

// MockFederatedSettingsDeletionReconciler is a mock of FederatedSettingsDeletionReconciler interface.
type MockFederatedSettingsDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedSettingsDeletionReconcilerMockRecorder
}

// MockFederatedSettingsDeletionReconcilerMockRecorder is the mock recorder for MockFederatedSettingsDeletionReconciler.
type MockFederatedSettingsDeletionReconcilerMockRecorder struct {
	mock *MockFederatedSettingsDeletionReconciler
}

// NewMockFederatedSettingsDeletionReconciler creates a new mock instance.
func NewMockFederatedSettingsDeletionReconciler(ctrl *gomock.Controller) *MockFederatedSettingsDeletionReconciler {
	mock := &MockFederatedSettingsDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedSettingsDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedSettingsDeletionReconciler) EXPECT() *MockFederatedSettingsDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedSettingsDeletion mocks base method.
func (m *MockFederatedSettingsDeletionReconciler) ReconcileFederatedSettingsDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedSettingsDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedSettingsDeletion indicates an expected call of ReconcileFederatedSettingsDeletion.
func (mr *MockFederatedSettingsDeletionReconcilerMockRecorder) ReconcileFederatedSettingsDeletion(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedSettingsDeletion", reflect.TypeOf((*MockFederatedSettingsDeletionReconciler)(nil).ReconcileFederatedSettingsDeletion), req)
}

// MockFederatedSettingsFinalizer is a mock of FederatedSettingsFinalizer interface.
type MockFederatedSettingsFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedSettingsFinalizerMockRecorder
}

// MockFederatedSettingsFinalizerMockRecorder is the mock recorder for MockFederatedSettingsFinalizer.
type MockFederatedSettingsFinalizerMockRecorder struct {
	mock *MockFederatedSettingsFinalizer
}

// NewMockFederatedSettingsFinalizer creates a new mock instance.
func NewMockFederatedSettingsFinalizer(ctrl *gomock.Controller) *MockFederatedSettingsFinalizer {
	mock := &MockFederatedSettingsFinalizer{ctrl: ctrl}
	mock.recorder = &MockFederatedSettingsFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedSettingsFinalizer) EXPECT() *MockFederatedSettingsFinalizerMockRecorder {
	return m.recorder
}

// FederatedSettingsFinalizerName mocks base method.
func (m *MockFederatedSettingsFinalizer) FederatedSettingsFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedSettingsFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FederatedSettingsFinalizerName indicates an expected call of FederatedSettingsFinalizerName.
func (mr *MockFederatedSettingsFinalizerMockRecorder) FederatedSettingsFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedSettingsFinalizerName", reflect.TypeOf((*MockFederatedSettingsFinalizer)(nil).FederatedSettingsFinalizerName))
}

// FinalizeFederatedSettings mocks base method.
func (m *MockFederatedSettingsFinalizer) FinalizeFederatedSettings(obj *v1.FederatedSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeFederatedSettings", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeFederatedSettings indicates an expected call of FinalizeFederatedSettings.
func (mr *MockFederatedSettingsFinalizerMockRecorder) FinalizeFederatedSettings(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeFederatedSettings", reflect.TypeOf((*MockFederatedSettingsFinalizer)(nil).FinalizeFederatedSettings), obj)
}

// ReconcileFederatedSettings mocks base method.
func (m *MockFederatedSettingsFinalizer) ReconcileFederatedSettings(obj *v1.FederatedSettings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedSettings", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedSettings indicates an expected call of ReconcileFederatedSettings.
func (mr *MockFederatedSettingsFinalizerMockRecorder) ReconcileFederatedSettings(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedSettings", reflect.TypeOf((*MockFederatedSettingsFinalizer)(nil).ReconcileFederatedSettings), obj)
}

// MockFederatedSettingsReconcileLoop is a mock of FederatedSettingsReconcileLoop interface.
type MockFederatedSettingsReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedSettingsReconcileLoopMockRecorder
}

// MockFederatedSettingsReconcileLoopMockRecorder is the mock recorder for MockFederatedSettingsReconcileLoop.
type MockFederatedSettingsReconcileLoopMockRecorder struct {
	mock *MockFederatedSettingsReconcileLoop
}

// NewMockFederatedSettingsReconcileLoop creates a new mock instance.
func NewMockFederatedSettingsReconcileLoop(ctrl *gomock.Controller) *MockFederatedSettingsReconcileLoop {
	mock := &MockFederatedSettingsReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockFederatedSettingsReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedSettingsReconcileLoop) EXPECT() *MockFederatedSettingsReconcileLoopMockRecorder {
	return m.recorder
}

// RunFederatedSettingsReconciler mocks base method.
func (m *MockFederatedSettingsReconcileLoop) RunFederatedSettingsReconciler(ctx context.Context, rec controller.FederatedSettingsReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunFederatedSettingsReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFederatedSettingsReconciler indicates an expected call of RunFederatedSettingsReconciler.
func (mr *MockFederatedSettingsReconcileLoopMockRecorder) RunFederatedSettingsReconciler(ctx, rec any, predicates ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFederatedSettingsReconciler", reflect.TypeOf((*MockFederatedSettingsReconcileLoop)(nil).RunFederatedSettingsReconciler), varargs...)
}

// MockFederatedUpstreamReconciler is a mock of FederatedUpstreamReconciler interface.
type MockFederatedUpstreamReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamReconcilerMockRecorder
}

// MockFederatedUpstreamReconcilerMockRecorder is the mock recorder for MockFederatedUpstreamReconciler.
type MockFederatedUpstreamReconcilerMockRecorder struct {
	mock *MockFederatedUpstreamReconciler
}

// NewMockFederatedUpstreamReconciler creates a new mock instance.
func NewMockFederatedUpstreamReconciler(ctrl *gomock.Controller) *MockFederatedUpstreamReconciler {
	mock := &MockFederatedUpstreamReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedUpstreamReconciler) EXPECT() *MockFederatedUpstreamReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedUpstream mocks base method.
func (m *MockFederatedUpstreamReconciler) ReconcileFederatedUpstream(obj *v1.FederatedUpstream) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstream", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedUpstream indicates an expected call of ReconcileFederatedUpstream.
func (mr *MockFederatedUpstreamReconcilerMockRecorder) ReconcileFederatedUpstream(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstream", reflect.TypeOf((*MockFederatedUpstreamReconciler)(nil).ReconcileFederatedUpstream), obj)
}

// MockFederatedUpstreamDeletionReconciler is a mock of FederatedUpstreamDeletionReconciler interface.
type MockFederatedUpstreamDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamDeletionReconcilerMockRecorder
}

// MockFederatedUpstreamDeletionReconcilerMockRecorder is the mock recorder for MockFederatedUpstreamDeletionReconciler.
type MockFederatedUpstreamDeletionReconcilerMockRecorder struct {
	mock *MockFederatedUpstreamDeletionReconciler
}

// NewMockFederatedUpstreamDeletionReconciler creates a new mock instance.
func NewMockFederatedUpstreamDeletionReconciler(ctrl *gomock.Controller) *MockFederatedUpstreamDeletionReconciler {
	mock := &MockFederatedUpstreamDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedUpstreamDeletionReconciler) EXPECT() *MockFederatedUpstreamDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedUpstreamDeletion mocks base method.
func (m *MockFederatedUpstreamDeletionReconciler) ReconcileFederatedUpstreamDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstreamDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedUpstreamDeletion indicates an expected call of ReconcileFederatedUpstreamDeletion.
func (mr *MockFederatedUpstreamDeletionReconcilerMockRecorder) ReconcileFederatedUpstreamDeletion(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstreamDeletion", reflect.TypeOf((*MockFederatedUpstreamDeletionReconciler)(nil).ReconcileFederatedUpstreamDeletion), req)
}

// MockFederatedUpstreamFinalizer is a mock of FederatedUpstreamFinalizer interface.
type MockFederatedUpstreamFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamFinalizerMockRecorder
}

// MockFederatedUpstreamFinalizerMockRecorder is the mock recorder for MockFederatedUpstreamFinalizer.
type MockFederatedUpstreamFinalizerMockRecorder struct {
	mock *MockFederatedUpstreamFinalizer
}

// NewMockFederatedUpstreamFinalizer creates a new mock instance.
func NewMockFederatedUpstreamFinalizer(ctrl *gomock.Controller) *MockFederatedUpstreamFinalizer {
	mock := &MockFederatedUpstreamFinalizer{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedUpstreamFinalizer) EXPECT() *MockFederatedUpstreamFinalizerMockRecorder {
	return m.recorder
}

// FederatedUpstreamFinalizerName mocks base method.
func (m *MockFederatedUpstreamFinalizer) FederatedUpstreamFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedUpstreamFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FederatedUpstreamFinalizerName indicates an expected call of FederatedUpstreamFinalizerName.
func (mr *MockFederatedUpstreamFinalizerMockRecorder) FederatedUpstreamFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedUpstreamFinalizerName", reflect.TypeOf((*MockFederatedUpstreamFinalizer)(nil).FederatedUpstreamFinalizerName))
}

// FinalizeFederatedUpstream mocks base method.
func (m *MockFederatedUpstreamFinalizer) FinalizeFederatedUpstream(obj *v1.FederatedUpstream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeFederatedUpstream", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeFederatedUpstream indicates an expected call of FinalizeFederatedUpstream.
func (mr *MockFederatedUpstreamFinalizerMockRecorder) FinalizeFederatedUpstream(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeFederatedUpstream", reflect.TypeOf((*MockFederatedUpstreamFinalizer)(nil).FinalizeFederatedUpstream), obj)
}

// ReconcileFederatedUpstream mocks base method.
func (m *MockFederatedUpstreamFinalizer) ReconcileFederatedUpstream(obj *v1.FederatedUpstream) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstream", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedUpstream indicates an expected call of ReconcileFederatedUpstream.
func (mr *MockFederatedUpstreamFinalizerMockRecorder) ReconcileFederatedUpstream(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstream", reflect.TypeOf((*MockFederatedUpstreamFinalizer)(nil).ReconcileFederatedUpstream), obj)
}

// MockFederatedUpstreamReconcileLoop is a mock of FederatedUpstreamReconcileLoop interface.
type MockFederatedUpstreamReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamReconcileLoopMockRecorder
}

// MockFederatedUpstreamReconcileLoopMockRecorder is the mock recorder for MockFederatedUpstreamReconcileLoop.
type MockFederatedUpstreamReconcileLoopMockRecorder struct {
	mock *MockFederatedUpstreamReconcileLoop
}

// NewMockFederatedUpstreamReconcileLoop creates a new mock instance.
func NewMockFederatedUpstreamReconcileLoop(ctrl *gomock.Controller) *MockFederatedUpstreamReconcileLoop {
	mock := &MockFederatedUpstreamReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedUpstreamReconcileLoop) EXPECT() *MockFederatedUpstreamReconcileLoopMockRecorder {
	return m.recorder
}

// RunFederatedUpstreamReconciler mocks base method.
func (m *MockFederatedUpstreamReconcileLoop) RunFederatedUpstreamReconciler(ctx context.Context, rec controller.FederatedUpstreamReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunFederatedUpstreamReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFederatedUpstreamReconciler indicates an expected call of RunFederatedUpstreamReconciler.
func (mr *MockFederatedUpstreamReconcileLoopMockRecorder) RunFederatedUpstreamReconciler(ctx, rec any, predicates ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFederatedUpstreamReconciler", reflect.TypeOf((*MockFederatedUpstreamReconcileLoop)(nil).RunFederatedUpstreamReconciler), varargs...)
}

// MockFederatedUpstreamGroupReconciler is a mock of FederatedUpstreamGroupReconciler interface.
type MockFederatedUpstreamGroupReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamGroupReconcilerMockRecorder
}

// MockFederatedUpstreamGroupReconcilerMockRecorder is the mock recorder for MockFederatedUpstreamGroupReconciler.
type MockFederatedUpstreamGroupReconcilerMockRecorder struct {
	mock *MockFederatedUpstreamGroupReconciler
}

// NewMockFederatedUpstreamGroupReconciler creates a new mock instance.
func NewMockFederatedUpstreamGroupReconciler(ctrl *gomock.Controller) *MockFederatedUpstreamGroupReconciler {
	mock := &MockFederatedUpstreamGroupReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamGroupReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedUpstreamGroupReconciler) EXPECT() *MockFederatedUpstreamGroupReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedUpstreamGroup mocks base method.
func (m *MockFederatedUpstreamGroupReconciler) ReconcileFederatedUpstreamGroup(obj *v1.FederatedUpstreamGroup) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstreamGroup", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedUpstreamGroup indicates an expected call of ReconcileFederatedUpstreamGroup.
func (mr *MockFederatedUpstreamGroupReconcilerMockRecorder) ReconcileFederatedUpstreamGroup(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstreamGroup", reflect.TypeOf((*MockFederatedUpstreamGroupReconciler)(nil).ReconcileFederatedUpstreamGroup), obj)
}

// MockFederatedUpstreamGroupDeletionReconciler is a mock of FederatedUpstreamGroupDeletionReconciler interface.
type MockFederatedUpstreamGroupDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamGroupDeletionReconcilerMockRecorder
}

// MockFederatedUpstreamGroupDeletionReconcilerMockRecorder is the mock recorder for MockFederatedUpstreamGroupDeletionReconciler.
type MockFederatedUpstreamGroupDeletionReconcilerMockRecorder struct {
	mock *MockFederatedUpstreamGroupDeletionReconciler
}

// NewMockFederatedUpstreamGroupDeletionReconciler creates a new mock instance.
func NewMockFederatedUpstreamGroupDeletionReconciler(ctrl *gomock.Controller) *MockFederatedUpstreamGroupDeletionReconciler {
	mock := &MockFederatedUpstreamGroupDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamGroupDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedUpstreamGroupDeletionReconciler) EXPECT() *MockFederatedUpstreamGroupDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedUpstreamGroupDeletion mocks base method.
func (m *MockFederatedUpstreamGroupDeletionReconciler) ReconcileFederatedUpstreamGroupDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstreamGroupDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedUpstreamGroupDeletion indicates an expected call of ReconcileFederatedUpstreamGroupDeletion.
func (mr *MockFederatedUpstreamGroupDeletionReconcilerMockRecorder) ReconcileFederatedUpstreamGroupDeletion(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstreamGroupDeletion", reflect.TypeOf((*MockFederatedUpstreamGroupDeletionReconciler)(nil).ReconcileFederatedUpstreamGroupDeletion), req)
}

// MockFederatedUpstreamGroupFinalizer is a mock of FederatedUpstreamGroupFinalizer interface.
type MockFederatedUpstreamGroupFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamGroupFinalizerMockRecorder
}

// MockFederatedUpstreamGroupFinalizerMockRecorder is the mock recorder for MockFederatedUpstreamGroupFinalizer.
type MockFederatedUpstreamGroupFinalizerMockRecorder struct {
	mock *MockFederatedUpstreamGroupFinalizer
}

// NewMockFederatedUpstreamGroupFinalizer creates a new mock instance.
func NewMockFederatedUpstreamGroupFinalizer(ctrl *gomock.Controller) *MockFederatedUpstreamGroupFinalizer {
	mock := &MockFederatedUpstreamGroupFinalizer{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamGroupFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedUpstreamGroupFinalizer) EXPECT() *MockFederatedUpstreamGroupFinalizerMockRecorder {
	return m.recorder
}

// FederatedUpstreamGroupFinalizerName mocks base method.
func (m *MockFederatedUpstreamGroupFinalizer) FederatedUpstreamGroupFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedUpstreamGroupFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FederatedUpstreamGroupFinalizerName indicates an expected call of FederatedUpstreamGroupFinalizerName.
func (mr *MockFederatedUpstreamGroupFinalizerMockRecorder) FederatedUpstreamGroupFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedUpstreamGroupFinalizerName", reflect.TypeOf((*MockFederatedUpstreamGroupFinalizer)(nil).FederatedUpstreamGroupFinalizerName))
}

// FinalizeFederatedUpstreamGroup mocks base method.
func (m *MockFederatedUpstreamGroupFinalizer) FinalizeFederatedUpstreamGroup(obj *v1.FederatedUpstreamGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeFederatedUpstreamGroup", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeFederatedUpstreamGroup indicates an expected call of FinalizeFederatedUpstreamGroup.
func (mr *MockFederatedUpstreamGroupFinalizerMockRecorder) FinalizeFederatedUpstreamGroup(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeFederatedUpstreamGroup", reflect.TypeOf((*MockFederatedUpstreamGroupFinalizer)(nil).FinalizeFederatedUpstreamGroup), obj)
}

// ReconcileFederatedUpstreamGroup mocks base method.
func (m *MockFederatedUpstreamGroupFinalizer) ReconcileFederatedUpstreamGroup(obj *v1.FederatedUpstreamGroup) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstreamGroup", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedUpstreamGroup indicates an expected call of ReconcileFederatedUpstreamGroup.
func (mr *MockFederatedUpstreamGroupFinalizerMockRecorder) ReconcileFederatedUpstreamGroup(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstreamGroup", reflect.TypeOf((*MockFederatedUpstreamGroupFinalizer)(nil).ReconcileFederatedUpstreamGroup), obj)
}

// MockFederatedUpstreamGroupReconcileLoop is a mock of FederatedUpstreamGroupReconcileLoop interface.
type MockFederatedUpstreamGroupReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamGroupReconcileLoopMockRecorder
}

// MockFederatedUpstreamGroupReconcileLoopMockRecorder is the mock recorder for MockFederatedUpstreamGroupReconcileLoop.
type MockFederatedUpstreamGroupReconcileLoopMockRecorder struct {
	mock *MockFederatedUpstreamGroupReconcileLoop
}

// NewMockFederatedUpstreamGroupReconcileLoop creates a new mock instance.
func NewMockFederatedUpstreamGroupReconcileLoop(ctrl *gomock.Controller) *MockFederatedUpstreamGroupReconcileLoop {
	mock := &MockFederatedUpstreamGroupReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamGroupReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedUpstreamGroupReconcileLoop) EXPECT() *MockFederatedUpstreamGroupReconcileLoopMockRecorder {
	return m.recorder
}

// RunFederatedUpstreamGroupReconciler mocks base method.
func (m *MockFederatedUpstreamGroupReconcileLoop) RunFederatedUpstreamGroupReconciler(ctx context.Context, rec controller.FederatedUpstreamGroupReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunFederatedUpstreamGroupReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFederatedUpstreamGroupReconciler indicates an expected call of RunFederatedUpstreamGroupReconciler.
func (mr *MockFederatedUpstreamGroupReconcileLoopMockRecorder) RunFederatedUpstreamGroupReconciler(ctx, rec any, predicates ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFederatedUpstreamGroupReconciler", reflect.TypeOf((*MockFederatedUpstreamGroupReconcileLoop)(nil).RunFederatedUpstreamGroupReconciler), varargs...)
}
