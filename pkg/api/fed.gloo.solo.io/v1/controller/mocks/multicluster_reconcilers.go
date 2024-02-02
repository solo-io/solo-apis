// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.gloo.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/fed.gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterFederatedSettingsReconciler is a mock of MulticlusterFederatedSettingsReconciler interface.
type MockMulticlusterFederatedSettingsReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedSettingsReconcilerMockRecorder
}

// MockMulticlusterFederatedSettingsReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedSettingsReconciler.
type MockMulticlusterFederatedSettingsReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedSettingsReconciler
}

// NewMockMulticlusterFederatedSettingsReconciler creates a new mock instance.
func NewMockMulticlusterFederatedSettingsReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedSettingsReconciler {
	mock := &MockMulticlusterFederatedSettingsReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedSettingsReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedSettingsReconciler) EXPECT() *MockMulticlusterFederatedSettingsReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedSettings mocks base method.
func (m *MockMulticlusterFederatedSettingsReconciler) ReconcileFederatedSettings(clusterName string, obj *v1.FederatedSettings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedSettings", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedSettings indicates an expected call of ReconcileFederatedSettings.
func (mr *MockMulticlusterFederatedSettingsReconcilerMockRecorder) ReconcileFederatedSettings(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedSettings", reflect.TypeOf((*MockMulticlusterFederatedSettingsReconciler)(nil).ReconcileFederatedSettings), clusterName, obj)
}

// MockMulticlusterFederatedSettingsDeletionReconciler is a mock of MulticlusterFederatedSettingsDeletionReconciler interface.
type MockMulticlusterFederatedSettingsDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedSettingsDeletionReconcilerMockRecorder
}

// MockMulticlusterFederatedSettingsDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedSettingsDeletionReconciler.
type MockMulticlusterFederatedSettingsDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedSettingsDeletionReconciler
}

// NewMockMulticlusterFederatedSettingsDeletionReconciler creates a new mock instance.
func NewMockMulticlusterFederatedSettingsDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedSettingsDeletionReconciler {
	mock := &MockMulticlusterFederatedSettingsDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedSettingsDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedSettingsDeletionReconciler) EXPECT() *MockMulticlusterFederatedSettingsDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedSettingsDeletion mocks base method.
func (m *MockMulticlusterFederatedSettingsDeletionReconciler) ReconcileFederatedSettingsDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedSettingsDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedSettingsDeletion indicates an expected call of ReconcileFederatedSettingsDeletion.
func (mr *MockMulticlusterFederatedSettingsDeletionReconcilerMockRecorder) ReconcileFederatedSettingsDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedSettingsDeletion", reflect.TypeOf((*MockMulticlusterFederatedSettingsDeletionReconciler)(nil).ReconcileFederatedSettingsDeletion), clusterName, req)
}

// MockMulticlusterFederatedSettingsReconcileLoop is a mock of MulticlusterFederatedSettingsReconcileLoop interface.
type MockMulticlusterFederatedSettingsReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedSettingsReconcileLoopMockRecorder
}

// MockMulticlusterFederatedSettingsReconcileLoopMockRecorder is the mock recorder for MockMulticlusterFederatedSettingsReconcileLoop.
type MockMulticlusterFederatedSettingsReconcileLoopMockRecorder struct {
	mock *MockMulticlusterFederatedSettingsReconcileLoop
}

// NewMockMulticlusterFederatedSettingsReconcileLoop creates a new mock instance.
func NewMockMulticlusterFederatedSettingsReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterFederatedSettingsReconcileLoop {
	mock := &MockMulticlusterFederatedSettingsReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedSettingsReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedSettingsReconcileLoop) EXPECT() *MockMulticlusterFederatedSettingsReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterFederatedSettingsReconciler mocks base method.
func (m *MockMulticlusterFederatedSettingsReconcileLoop) AddMulticlusterFederatedSettingsReconciler(ctx context.Context, rec controller.MulticlusterFederatedSettingsReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterFederatedSettingsReconciler", varargs...)
}

// AddMulticlusterFederatedSettingsReconciler indicates an expected call of AddMulticlusterFederatedSettingsReconciler.
func (mr *MockMulticlusterFederatedSettingsReconcileLoopMockRecorder) AddMulticlusterFederatedSettingsReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterFederatedSettingsReconciler", reflect.TypeOf((*MockMulticlusterFederatedSettingsReconcileLoop)(nil).AddMulticlusterFederatedSettingsReconciler), varargs...)
}

// MockMulticlusterFederatedUpstreamReconciler is a mock of MulticlusterFederatedUpstreamReconciler interface.
type MockMulticlusterFederatedUpstreamReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedUpstreamReconcilerMockRecorder
}

// MockMulticlusterFederatedUpstreamReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedUpstreamReconciler.
type MockMulticlusterFederatedUpstreamReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedUpstreamReconciler
}

// NewMockMulticlusterFederatedUpstreamReconciler creates a new mock instance.
func NewMockMulticlusterFederatedUpstreamReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedUpstreamReconciler {
	mock := &MockMulticlusterFederatedUpstreamReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedUpstreamReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedUpstreamReconciler) EXPECT() *MockMulticlusterFederatedUpstreamReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedUpstream mocks base method.
func (m *MockMulticlusterFederatedUpstreamReconciler) ReconcileFederatedUpstream(clusterName string, obj *v1.FederatedUpstream) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstream", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedUpstream indicates an expected call of ReconcileFederatedUpstream.
func (mr *MockMulticlusterFederatedUpstreamReconcilerMockRecorder) ReconcileFederatedUpstream(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstream", reflect.TypeOf((*MockMulticlusterFederatedUpstreamReconciler)(nil).ReconcileFederatedUpstream), clusterName, obj)
}

// MockMulticlusterFederatedUpstreamDeletionReconciler is a mock of MulticlusterFederatedUpstreamDeletionReconciler interface.
type MockMulticlusterFederatedUpstreamDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedUpstreamDeletionReconcilerMockRecorder
}

// MockMulticlusterFederatedUpstreamDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedUpstreamDeletionReconciler.
type MockMulticlusterFederatedUpstreamDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedUpstreamDeletionReconciler
}

// NewMockMulticlusterFederatedUpstreamDeletionReconciler creates a new mock instance.
func NewMockMulticlusterFederatedUpstreamDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedUpstreamDeletionReconciler {
	mock := &MockMulticlusterFederatedUpstreamDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedUpstreamDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedUpstreamDeletionReconciler) EXPECT() *MockMulticlusterFederatedUpstreamDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedUpstreamDeletion mocks base method.
func (m *MockMulticlusterFederatedUpstreamDeletionReconciler) ReconcileFederatedUpstreamDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstreamDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedUpstreamDeletion indicates an expected call of ReconcileFederatedUpstreamDeletion.
func (mr *MockMulticlusterFederatedUpstreamDeletionReconcilerMockRecorder) ReconcileFederatedUpstreamDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstreamDeletion", reflect.TypeOf((*MockMulticlusterFederatedUpstreamDeletionReconciler)(nil).ReconcileFederatedUpstreamDeletion), clusterName, req)
}

// MockMulticlusterFederatedUpstreamReconcileLoop is a mock of MulticlusterFederatedUpstreamReconcileLoop interface.
type MockMulticlusterFederatedUpstreamReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedUpstreamReconcileLoopMockRecorder
}

// MockMulticlusterFederatedUpstreamReconcileLoopMockRecorder is the mock recorder for MockMulticlusterFederatedUpstreamReconcileLoop.
type MockMulticlusterFederatedUpstreamReconcileLoopMockRecorder struct {
	mock *MockMulticlusterFederatedUpstreamReconcileLoop
}

// NewMockMulticlusterFederatedUpstreamReconcileLoop creates a new mock instance.
func NewMockMulticlusterFederatedUpstreamReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterFederatedUpstreamReconcileLoop {
	mock := &MockMulticlusterFederatedUpstreamReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedUpstreamReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedUpstreamReconcileLoop) EXPECT() *MockMulticlusterFederatedUpstreamReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterFederatedUpstreamReconciler mocks base method.
func (m *MockMulticlusterFederatedUpstreamReconcileLoop) AddMulticlusterFederatedUpstreamReconciler(ctx context.Context, rec controller.MulticlusterFederatedUpstreamReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterFederatedUpstreamReconciler", varargs...)
}

// AddMulticlusterFederatedUpstreamReconciler indicates an expected call of AddMulticlusterFederatedUpstreamReconciler.
func (mr *MockMulticlusterFederatedUpstreamReconcileLoopMockRecorder) AddMulticlusterFederatedUpstreamReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterFederatedUpstreamReconciler", reflect.TypeOf((*MockMulticlusterFederatedUpstreamReconcileLoop)(nil).AddMulticlusterFederatedUpstreamReconciler), varargs...)
}

// MockMulticlusterFederatedUpstreamGroupReconciler is a mock of MulticlusterFederatedUpstreamGroupReconciler interface.
type MockMulticlusterFederatedUpstreamGroupReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedUpstreamGroupReconcilerMockRecorder
}

// MockMulticlusterFederatedUpstreamGroupReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedUpstreamGroupReconciler.
type MockMulticlusterFederatedUpstreamGroupReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedUpstreamGroupReconciler
}

// NewMockMulticlusterFederatedUpstreamGroupReconciler creates a new mock instance.
func NewMockMulticlusterFederatedUpstreamGroupReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedUpstreamGroupReconciler {
	mock := &MockMulticlusterFederatedUpstreamGroupReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedUpstreamGroupReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedUpstreamGroupReconciler) EXPECT() *MockMulticlusterFederatedUpstreamGroupReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedUpstreamGroup mocks base method.
func (m *MockMulticlusterFederatedUpstreamGroupReconciler) ReconcileFederatedUpstreamGroup(clusterName string, obj *v1.FederatedUpstreamGroup) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstreamGroup", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedUpstreamGroup indicates an expected call of ReconcileFederatedUpstreamGroup.
func (mr *MockMulticlusterFederatedUpstreamGroupReconcilerMockRecorder) ReconcileFederatedUpstreamGroup(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstreamGroup", reflect.TypeOf((*MockMulticlusterFederatedUpstreamGroupReconciler)(nil).ReconcileFederatedUpstreamGroup), clusterName, obj)
}

// MockMulticlusterFederatedUpstreamGroupDeletionReconciler is a mock of MulticlusterFederatedUpstreamGroupDeletionReconciler interface.
type MockMulticlusterFederatedUpstreamGroupDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedUpstreamGroupDeletionReconcilerMockRecorder
}

// MockMulticlusterFederatedUpstreamGroupDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedUpstreamGroupDeletionReconciler.
type MockMulticlusterFederatedUpstreamGroupDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedUpstreamGroupDeletionReconciler
}

// NewMockMulticlusterFederatedUpstreamGroupDeletionReconciler creates a new mock instance.
func NewMockMulticlusterFederatedUpstreamGroupDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedUpstreamGroupDeletionReconciler {
	mock := &MockMulticlusterFederatedUpstreamGroupDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedUpstreamGroupDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedUpstreamGroupDeletionReconciler) EXPECT() *MockMulticlusterFederatedUpstreamGroupDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedUpstreamGroupDeletion mocks base method.
func (m *MockMulticlusterFederatedUpstreamGroupDeletionReconciler) ReconcileFederatedUpstreamGroupDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedUpstreamGroupDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedUpstreamGroupDeletion indicates an expected call of ReconcileFederatedUpstreamGroupDeletion.
func (mr *MockMulticlusterFederatedUpstreamGroupDeletionReconcilerMockRecorder) ReconcileFederatedUpstreamGroupDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedUpstreamGroupDeletion", reflect.TypeOf((*MockMulticlusterFederatedUpstreamGroupDeletionReconciler)(nil).ReconcileFederatedUpstreamGroupDeletion), clusterName, req)
}

// MockMulticlusterFederatedUpstreamGroupReconcileLoop is a mock of MulticlusterFederatedUpstreamGroupReconcileLoop interface.
type MockMulticlusterFederatedUpstreamGroupReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedUpstreamGroupReconcileLoopMockRecorder
}

// MockMulticlusterFederatedUpstreamGroupReconcileLoopMockRecorder is the mock recorder for MockMulticlusterFederatedUpstreamGroupReconcileLoop.
type MockMulticlusterFederatedUpstreamGroupReconcileLoopMockRecorder struct {
	mock *MockMulticlusterFederatedUpstreamGroupReconcileLoop
}

// NewMockMulticlusterFederatedUpstreamGroupReconcileLoop creates a new mock instance.
func NewMockMulticlusterFederatedUpstreamGroupReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterFederatedUpstreamGroupReconcileLoop {
	mock := &MockMulticlusterFederatedUpstreamGroupReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedUpstreamGroupReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedUpstreamGroupReconcileLoop) EXPECT() *MockMulticlusterFederatedUpstreamGroupReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterFederatedUpstreamGroupReconciler mocks base method.
func (m *MockMulticlusterFederatedUpstreamGroupReconcileLoop) AddMulticlusterFederatedUpstreamGroupReconciler(ctx context.Context, rec controller.MulticlusterFederatedUpstreamGroupReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterFederatedUpstreamGroupReconciler", varargs...)
}

// AddMulticlusterFederatedUpstreamGroupReconciler indicates an expected call of AddMulticlusterFederatedUpstreamGroupReconciler.
func (mr *MockMulticlusterFederatedUpstreamGroupReconcileLoopMockRecorder) AddMulticlusterFederatedUpstreamGroupReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterFederatedUpstreamGroupReconciler", reflect.TypeOf((*MockMulticlusterFederatedUpstreamGroupReconcileLoop)(nil).AddMulticlusterFederatedUpstreamGroupReconciler), varargs...)
}
