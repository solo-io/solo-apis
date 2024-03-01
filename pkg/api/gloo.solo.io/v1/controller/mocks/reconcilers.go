// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockSettingsReconciler is a mock of SettingsReconciler interface
type MockSettingsReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsReconcilerMockRecorder
}

// MockSettingsReconcilerMockRecorder is the mock recorder for MockSettingsReconciler
type MockSettingsReconcilerMockRecorder struct {
	mock *MockSettingsReconciler
}

// NewMockSettingsReconciler creates a new mock instance
func NewMockSettingsReconciler(ctrl *gomock.Controller) *MockSettingsReconciler {
	mock := &MockSettingsReconciler{ctrl: ctrl}
	mock.recorder = &MockSettingsReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsReconciler) EXPECT() *MockSettingsReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettings mocks base method
func (m *MockSettingsReconciler) ReconcileSettings(obj *v1.Settings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettings", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSettings indicates an expected call of ReconcileSettings
func (mr *MockSettingsReconcilerMockRecorder) ReconcileSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettings", reflect.TypeOf((*MockSettingsReconciler)(nil).ReconcileSettings), obj)
}

// MockSettingsDeletionReconciler is a mock of SettingsDeletionReconciler interface
type MockSettingsDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsDeletionReconcilerMockRecorder
}

// MockSettingsDeletionReconcilerMockRecorder is the mock recorder for MockSettingsDeletionReconciler
type MockSettingsDeletionReconcilerMockRecorder struct {
	mock *MockSettingsDeletionReconciler
}

// NewMockSettingsDeletionReconciler creates a new mock instance
func NewMockSettingsDeletionReconciler(ctrl *gomock.Controller) *MockSettingsDeletionReconciler {
	mock := &MockSettingsDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockSettingsDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsDeletionReconciler) EXPECT() *MockSettingsDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettingsDeletion mocks base method
func (m *MockSettingsDeletionReconciler) ReconcileSettingsDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettingsDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileSettingsDeletion indicates an expected call of ReconcileSettingsDeletion
func (mr *MockSettingsDeletionReconcilerMockRecorder) ReconcileSettingsDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettingsDeletion", reflect.TypeOf((*MockSettingsDeletionReconciler)(nil).ReconcileSettingsDeletion), req)
}

// MockSettingsFinalizer is a mock of SettingsFinalizer interface
type MockSettingsFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsFinalizerMockRecorder
}

// MockSettingsFinalizerMockRecorder is the mock recorder for MockSettingsFinalizer
type MockSettingsFinalizerMockRecorder struct {
	mock *MockSettingsFinalizer
}

// NewMockSettingsFinalizer creates a new mock instance
func NewMockSettingsFinalizer(ctrl *gomock.Controller) *MockSettingsFinalizer {
	mock := &MockSettingsFinalizer{ctrl: ctrl}
	mock.recorder = &MockSettingsFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsFinalizer) EXPECT() *MockSettingsFinalizerMockRecorder {
	return m.recorder
}

// ReconcileSettings mocks base method
func (m *MockSettingsFinalizer) ReconcileSettings(obj *v1.Settings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettings", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSettings indicates an expected call of ReconcileSettings
func (mr *MockSettingsFinalizerMockRecorder) ReconcileSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettings", reflect.TypeOf((*MockSettingsFinalizer)(nil).ReconcileSettings), obj)
}

// SettingsFinalizerName mocks base method
func (m *MockSettingsFinalizer) SettingsFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SettingsFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// SettingsFinalizerName indicates an expected call of SettingsFinalizerName
func (mr *MockSettingsFinalizerMockRecorder) SettingsFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SettingsFinalizerName", reflect.TypeOf((*MockSettingsFinalizer)(nil).SettingsFinalizerName))
}

// FinalizeSettings mocks base method
func (m *MockSettingsFinalizer) FinalizeSettings(obj *v1.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeSettings", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeSettings indicates an expected call of FinalizeSettings
func (mr *MockSettingsFinalizerMockRecorder) FinalizeSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeSettings", reflect.TypeOf((*MockSettingsFinalizer)(nil).FinalizeSettings), obj)
}

// MockSettingsReconcileLoop is a mock of SettingsReconcileLoop interface
type MockSettingsReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsReconcileLoopMockRecorder
}

// MockSettingsReconcileLoopMockRecorder is the mock recorder for MockSettingsReconcileLoop
type MockSettingsReconcileLoopMockRecorder struct {
	mock *MockSettingsReconcileLoop
}

// NewMockSettingsReconcileLoop creates a new mock instance
func NewMockSettingsReconcileLoop(ctrl *gomock.Controller) *MockSettingsReconcileLoop {
	mock := &MockSettingsReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockSettingsReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsReconcileLoop) EXPECT() *MockSettingsReconcileLoopMockRecorder {
	return m.recorder
}

// RunSettingsReconciler mocks base method
func (m *MockSettingsReconcileLoop) RunSettingsReconciler(ctx context.Context, rec controller.SettingsReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunSettingsReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunSettingsReconciler indicates an expected call of RunSettingsReconciler
func (mr *MockSettingsReconcileLoopMockRecorder) RunSettingsReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSettingsReconciler", reflect.TypeOf((*MockSettingsReconcileLoop)(nil).RunSettingsReconciler), varargs...)
}

// MockUpstreamReconciler is a mock of UpstreamReconciler interface
type MockUpstreamReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamReconcilerMockRecorder
}

// MockUpstreamReconcilerMockRecorder is the mock recorder for MockUpstreamReconciler
type MockUpstreamReconcilerMockRecorder struct {
	mock *MockUpstreamReconciler
}

// NewMockUpstreamReconciler creates a new mock instance
func NewMockUpstreamReconciler(ctrl *gomock.Controller) *MockUpstreamReconciler {
	mock := &MockUpstreamReconciler{ctrl: ctrl}
	mock.recorder = &MockUpstreamReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamReconciler) EXPECT() *MockUpstreamReconcilerMockRecorder {
	return m.recorder
}

// ReconcileUpstream mocks base method
func (m *MockUpstreamReconciler) ReconcileUpstream(obj *v1.Upstream) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileUpstream", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileUpstream indicates an expected call of ReconcileUpstream
func (mr *MockUpstreamReconcilerMockRecorder) ReconcileUpstream(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileUpstream", reflect.TypeOf((*MockUpstreamReconciler)(nil).ReconcileUpstream), obj)
}

// MockUpstreamDeletionReconciler is a mock of UpstreamDeletionReconciler interface
type MockUpstreamDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamDeletionReconcilerMockRecorder
}

// MockUpstreamDeletionReconcilerMockRecorder is the mock recorder for MockUpstreamDeletionReconciler
type MockUpstreamDeletionReconcilerMockRecorder struct {
	mock *MockUpstreamDeletionReconciler
}

// NewMockUpstreamDeletionReconciler creates a new mock instance
func NewMockUpstreamDeletionReconciler(ctrl *gomock.Controller) *MockUpstreamDeletionReconciler {
	mock := &MockUpstreamDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockUpstreamDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamDeletionReconciler) EXPECT() *MockUpstreamDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileUpstreamDeletion mocks base method
func (m *MockUpstreamDeletionReconciler) ReconcileUpstreamDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileUpstreamDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileUpstreamDeletion indicates an expected call of ReconcileUpstreamDeletion
func (mr *MockUpstreamDeletionReconcilerMockRecorder) ReconcileUpstreamDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileUpstreamDeletion", reflect.TypeOf((*MockUpstreamDeletionReconciler)(nil).ReconcileUpstreamDeletion), req)
}

// MockUpstreamFinalizer is a mock of UpstreamFinalizer interface
type MockUpstreamFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamFinalizerMockRecorder
}

// MockUpstreamFinalizerMockRecorder is the mock recorder for MockUpstreamFinalizer
type MockUpstreamFinalizerMockRecorder struct {
	mock *MockUpstreamFinalizer
}

// NewMockUpstreamFinalizer creates a new mock instance
func NewMockUpstreamFinalizer(ctrl *gomock.Controller) *MockUpstreamFinalizer {
	mock := &MockUpstreamFinalizer{ctrl: ctrl}
	mock.recorder = &MockUpstreamFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamFinalizer) EXPECT() *MockUpstreamFinalizerMockRecorder {
	return m.recorder
}

// ReconcileUpstream mocks base method
func (m *MockUpstreamFinalizer) ReconcileUpstream(obj *v1.Upstream) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileUpstream", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileUpstream indicates an expected call of ReconcileUpstream
func (mr *MockUpstreamFinalizerMockRecorder) ReconcileUpstream(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileUpstream", reflect.TypeOf((*MockUpstreamFinalizer)(nil).ReconcileUpstream), obj)
}

// UpstreamFinalizerName mocks base method
func (m *MockUpstreamFinalizer) UpstreamFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpstreamFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// UpstreamFinalizerName indicates an expected call of UpstreamFinalizerName
func (mr *MockUpstreamFinalizerMockRecorder) UpstreamFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpstreamFinalizerName", reflect.TypeOf((*MockUpstreamFinalizer)(nil).UpstreamFinalizerName))
}

// FinalizeUpstream mocks base method
func (m *MockUpstreamFinalizer) FinalizeUpstream(obj *v1.Upstream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeUpstream", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeUpstream indicates an expected call of FinalizeUpstream
func (mr *MockUpstreamFinalizerMockRecorder) FinalizeUpstream(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeUpstream", reflect.TypeOf((*MockUpstreamFinalizer)(nil).FinalizeUpstream), obj)
}

// MockUpstreamReconcileLoop is a mock of UpstreamReconcileLoop interface
type MockUpstreamReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamReconcileLoopMockRecorder
}

// MockUpstreamReconcileLoopMockRecorder is the mock recorder for MockUpstreamReconcileLoop
type MockUpstreamReconcileLoopMockRecorder struct {
	mock *MockUpstreamReconcileLoop
}

// NewMockUpstreamReconcileLoop creates a new mock instance
func NewMockUpstreamReconcileLoop(ctrl *gomock.Controller) *MockUpstreamReconcileLoop {
	mock := &MockUpstreamReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockUpstreamReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamReconcileLoop) EXPECT() *MockUpstreamReconcileLoopMockRecorder {
	return m.recorder
}

// RunUpstreamReconciler mocks base method
func (m *MockUpstreamReconcileLoop) RunUpstreamReconciler(ctx context.Context, rec controller.UpstreamReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunUpstreamReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunUpstreamReconciler indicates an expected call of RunUpstreamReconciler
func (mr *MockUpstreamReconcileLoopMockRecorder) RunUpstreamReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunUpstreamReconciler", reflect.TypeOf((*MockUpstreamReconcileLoop)(nil).RunUpstreamReconciler), varargs...)
}

// MockUpstreamGroupReconciler is a mock of UpstreamGroupReconciler interface
type MockUpstreamGroupReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamGroupReconcilerMockRecorder
}

// MockUpstreamGroupReconcilerMockRecorder is the mock recorder for MockUpstreamGroupReconciler
type MockUpstreamGroupReconcilerMockRecorder struct {
	mock *MockUpstreamGroupReconciler
}

// NewMockUpstreamGroupReconciler creates a new mock instance
func NewMockUpstreamGroupReconciler(ctrl *gomock.Controller) *MockUpstreamGroupReconciler {
	mock := &MockUpstreamGroupReconciler{ctrl: ctrl}
	mock.recorder = &MockUpstreamGroupReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamGroupReconciler) EXPECT() *MockUpstreamGroupReconcilerMockRecorder {
	return m.recorder
}

// ReconcileUpstreamGroup mocks base method
func (m *MockUpstreamGroupReconciler) ReconcileUpstreamGroup(obj *v1.UpstreamGroup) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileUpstreamGroup", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileUpstreamGroup indicates an expected call of ReconcileUpstreamGroup
func (mr *MockUpstreamGroupReconcilerMockRecorder) ReconcileUpstreamGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileUpstreamGroup", reflect.TypeOf((*MockUpstreamGroupReconciler)(nil).ReconcileUpstreamGroup), obj)
}

// MockUpstreamGroupDeletionReconciler is a mock of UpstreamGroupDeletionReconciler interface
type MockUpstreamGroupDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamGroupDeletionReconcilerMockRecorder
}

// MockUpstreamGroupDeletionReconcilerMockRecorder is the mock recorder for MockUpstreamGroupDeletionReconciler
type MockUpstreamGroupDeletionReconcilerMockRecorder struct {
	mock *MockUpstreamGroupDeletionReconciler
}

// NewMockUpstreamGroupDeletionReconciler creates a new mock instance
func NewMockUpstreamGroupDeletionReconciler(ctrl *gomock.Controller) *MockUpstreamGroupDeletionReconciler {
	mock := &MockUpstreamGroupDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockUpstreamGroupDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamGroupDeletionReconciler) EXPECT() *MockUpstreamGroupDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileUpstreamGroupDeletion mocks base method
func (m *MockUpstreamGroupDeletionReconciler) ReconcileUpstreamGroupDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileUpstreamGroupDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileUpstreamGroupDeletion indicates an expected call of ReconcileUpstreamGroupDeletion
func (mr *MockUpstreamGroupDeletionReconcilerMockRecorder) ReconcileUpstreamGroupDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileUpstreamGroupDeletion", reflect.TypeOf((*MockUpstreamGroupDeletionReconciler)(nil).ReconcileUpstreamGroupDeletion), req)
}

// MockUpstreamGroupFinalizer is a mock of UpstreamGroupFinalizer interface
type MockUpstreamGroupFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamGroupFinalizerMockRecorder
}

// MockUpstreamGroupFinalizerMockRecorder is the mock recorder for MockUpstreamGroupFinalizer
type MockUpstreamGroupFinalizerMockRecorder struct {
	mock *MockUpstreamGroupFinalizer
}

// NewMockUpstreamGroupFinalizer creates a new mock instance
func NewMockUpstreamGroupFinalizer(ctrl *gomock.Controller) *MockUpstreamGroupFinalizer {
	mock := &MockUpstreamGroupFinalizer{ctrl: ctrl}
	mock.recorder = &MockUpstreamGroupFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamGroupFinalizer) EXPECT() *MockUpstreamGroupFinalizerMockRecorder {
	return m.recorder
}

// ReconcileUpstreamGroup mocks base method
func (m *MockUpstreamGroupFinalizer) ReconcileUpstreamGroup(obj *v1.UpstreamGroup) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileUpstreamGroup", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileUpstreamGroup indicates an expected call of ReconcileUpstreamGroup
func (mr *MockUpstreamGroupFinalizerMockRecorder) ReconcileUpstreamGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileUpstreamGroup", reflect.TypeOf((*MockUpstreamGroupFinalizer)(nil).ReconcileUpstreamGroup), obj)
}

// UpstreamGroupFinalizerName mocks base method
func (m *MockUpstreamGroupFinalizer) UpstreamGroupFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpstreamGroupFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// UpstreamGroupFinalizerName indicates an expected call of UpstreamGroupFinalizerName
func (mr *MockUpstreamGroupFinalizerMockRecorder) UpstreamGroupFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpstreamGroupFinalizerName", reflect.TypeOf((*MockUpstreamGroupFinalizer)(nil).UpstreamGroupFinalizerName))
}

// FinalizeUpstreamGroup mocks base method
func (m *MockUpstreamGroupFinalizer) FinalizeUpstreamGroup(obj *v1.UpstreamGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeUpstreamGroup", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeUpstreamGroup indicates an expected call of FinalizeUpstreamGroup
func (mr *MockUpstreamGroupFinalizerMockRecorder) FinalizeUpstreamGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeUpstreamGroup", reflect.TypeOf((*MockUpstreamGroupFinalizer)(nil).FinalizeUpstreamGroup), obj)
}

// MockUpstreamGroupReconcileLoop is a mock of UpstreamGroupReconcileLoop interface
type MockUpstreamGroupReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamGroupReconcileLoopMockRecorder
}

// MockUpstreamGroupReconcileLoopMockRecorder is the mock recorder for MockUpstreamGroupReconcileLoop
type MockUpstreamGroupReconcileLoopMockRecorder struct {
	mock *MockUpstreamGroupReconcileLoop
}

// NewMockUpstreamGroupReconcileLoop creates a new mock instance
func NewMockUpstreamGroupReconcileLoop(ctrl *gomock.Controller) *MockUpstreamGroupReconcileLoop {
	mock := &MockUpstreamGroupReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockUpstreamGroupReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamGroupReconcileLoop) EXPECT() *MockUpstreamGroupReconcileLoopMockRecorder {
	return m.recorder
}

// RunUpstreamGroupReconciler mocks base method
func (m *MockUpstreamGroupReconcileLoop) RunUpstreamGroupReconciler(ctx context.Context, rec controller.UpstreamGroupReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunUpstreamGroupReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunUpstreamGroupReconciler indicates an expected call of RunUpstreamGroupReconciler
func (mr *MockUpstreamGroupReconcileLoopMockRecorder) RunUpstreamGroupReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunUpstreamGroupReconciler", reflect.TypeOf((*MockUpstreamGroupReconcileLoop)(nil).RunUpstreamGroupReconciler), varargs...)
}

// MockProxyReconciler is a mock of ProxyReconciler interface
type MockProxyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockProxyReconcilerMockRecorder
}

// MockProxyReconcilerMockRecorder is the mock recorder for MockProxyReconciler
type MockProxyReconcilerMockRecorder struct {
	mock *MockProxyReconciler
}

// NewMockProxyReconciler creates a new mock instance
func NewMockProxyReconciler(ctrl *gomock.Controller) *MockProxyReconciler {
	mock := &MockProxyReconciler{ctrl: ctrl}
	mock.recorder = &MockProxyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProxyReconciler) EXPECT() *MockProxyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileProxy mocks base method
func (m *MockProxyReconciler) ReconcileProxy(obj *v1.Proxy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileProxy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileProxy indicates an expected call of ReconcileProxy
func (mr *MockProxyReconcilerMockRecorder) ReconcileProxy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileProxy", reflect.TypeOf((*MockProxyReconciler)(nil).ReconcileProxy), obj)
}

// MockProxyDeletionReconciler is a mock of ProxyDeletionReconciler interface
type MockProxyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockProxyDeletionReconcilerMockRecorder
}

// MockProxyDeletionReconcilerMockRecorder is the mock recorder for MockProxyDeletionReconciler
type MockProxyDeletionReconcilerMockRecorder struct {
	mock *MockProxyDeletionReconciler
}

// NewMockProxyDeletionReconciler creates a new mock instance
func NewMockProxyDeletionReconciler(ctrl *gomock.Controller) *MockProxyDeletionReconciler {
	mock := &MockProxyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockProxyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProxyDeletionReconciler) EXPECT() *MockProxyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileProxyDeletion mocks base method
func (m *MockProxyDeletionReconciler) ReconcileProxyDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileProxyDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileProxyDeletion indicates an expected call of ReconcileProxyDeletion
func (mr *MockProxyDeletionReconcilerMockRecorder) ReconcileProxyDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileProxyDeletion", reflect.TypeOf((*MockProxyDeletionReconciler)(nil).ReconcileProxyDeletion), req)
}

// MockProxyFinalizer is a mock of ProxyFinalizer interface
type MockProxyFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockProxyFinalizerMockRecorder
}

// MockProxyFinalizerMockRecorder is the mock recorder for MockProxyFinalizer
type MockProxyFinalizerMockRecorder struct {
	mock *MockProxyFinalizer
}

// NewMockProxyFinalizer creates a new mock instance
func NewMockProxyFinalizer(ctrl *gomock.Controller) *MockProxyFinalizer {
	mock := &MockProxyFinalizer{ctrl: ctrl}
	mock.recorder = &MockProxyFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProxyFinalizer) EXPECT() *MockProxyFinalizerMockRecorder {
	return m.recorder
}

// ReconcileProxy mocks base method
func (m *MockProxyFinalizer) ReconcileProxy(obj *v1.Proxy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileProxy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileProxy indicates an expected call of ReconcileProxy
func (mr *MockProxyFinalizerMockRecorder) ReconcileProxy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileProxy", reflect.TypeOf((*MockProxyFinalizer)(nil).ReconcileProxy), obj)
}

// ProxyFinalizerName mocks base method
func (m *MockProxyFinalizer) ProxyFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProxyFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ProxyFinalizerName indicates an expected call of ProxyFinalizerName
func (mr *MockProxyFinalizerMockRecorder) ProxyFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProxyFinalizerName", reflect.TypeOf((*MockProxyFinalizer)(nil).ProxyFinalizerName))
}

// FinalizeProxy mocks base method
func (m *MockProxyFinalizer) FinalizeProxy(obj *v1.Proxy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeProxy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeProxy indicates an expected call of FinalizeProxy
func (mr *MockProxyFinalizerMockRecorder) FinalizeProxy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeProxy", reflect.TypeOf((*MockProxyFinalizer)(nil).FinalizeProxy), obj)
}

// MockProxyReconcileLoop is a mock of ProxyReconcileLoop interface
type MockProxyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockProxyReconcileLoopMockRecorder
}

// MockProxyReconcileLoopMockRecorder is the mock recorder for MockProxyReconcileLoop
type MockProxyReconcileLoopMockRecorder struct {
	mock *MockProxyReconcileLoop
}

// NewMockProxyReconcileLoop creates a new mock instance
func NewMockProxyReconcileLoop(ctrl *gomock.Controller) *MockProxyReconcileLoop {
	mock := &MockProxyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockProxyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProxyReconcileLoop) EXPECT() *MockProxyReconcileLoopMockRecorder {
	return m.recorder
}

// RunProxyReconciler mocks base method
func (m *MockProxyReconcileLoop) RunProxyReconciler(ctx context.Context, rec controller.ProxyReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunProxyReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunProxyReconciler indicates an expected call of RunProxyReconciler
func (mr *MockProxyReconcileLoopMockRecorder) RunProxyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunProxyReconciler", reflect.TypeOf((*MockProxyReconcileLoop)(nil).RunProxyReconciler), varargs...)
}
