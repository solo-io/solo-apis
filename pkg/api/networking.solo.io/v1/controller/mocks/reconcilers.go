// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/networking.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/networking.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockTrafficPolicyReconciler is a mock of TrafficPolicyReconciler interface.
type MockTrafficPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyReconcilerMockRecorder
}

// MockTrafficPolicyReconcilerMockRecorder is the mock recorder for MockTrafficPolicyReconciler.
type MockTrafficPolicyReconcilerMockRecorder struct {
	mock *MockTrafficPolicyReconciler
}

// NewMockTrafficPolicyReconciler creates a new mock instance.
func NewMockTrafficPolicyReconciler(ctrl *gomock.Controller) *MockTrafficPolicyReconciler {
	mock := &MockTrafficPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyReconciler) EXPECT() *MockTrafficPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficPolicy mocks base method.
func (m *MockTrafficPolicyReconciler) ReconcileTrafficPolicy(obj *v1.TrafficPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficPolicy indicates an expected call of ReconcileTrafficPolicy.
func (mr *MockTrafficPolicyReconcilerMockRecorder) ReconcileTrafficPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicy", reflect.TypeOf((*MockTrafficPolicyReconciler)(nil).ReconcileTrafficPolicy), obj)
}

// MockTrafficPolicyDeletionReconciler is a mock of TrafficPolicyDeletionReconciler interface.
type MockTrafficPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyDeletionReconcilerMockRecorder
}

// MockTrafficPolicyDeletionReconcilerMockRecorder is the mock recorder for MockTrafficPolicyDeletionReconciler.
type MockTrafficPolicyDeletionReconcilerMockRecorder struct {
	mock *MockTrafficPolicyDeletionReconciler
}

// NewMockTrafficPolicyDeletionReconciler creates a new mock instance.
func NewMockTrafficPolicyDeletionReconciler(ctrl *gomock.Controller) *MockTrafficPolicyDeletionReconciler {
	mock := &MockTrafficPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyDeletionReconciler) EXPECT() *MockTrafficPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficPolicyDeletion mocks base method.
func (m *MockTrafficPolicyDeletionReconciler) ReconcileTrafficPolicyDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficPolicyDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileTrafficPolicyDeletion indicates an expected call of ReconcileTrafficPolicyDeletion.
func (mr *MockTrafficPolicyDeletionReconcilerMockRecorder) ReconcileTrafficPolicyDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicyDeletion", reflect.TypeOf((*MockTrafficPolicyDeletionReconciler)(nil).ReconcileTrafficPolicyDeletion), req)
}

// MockTrafficPolicyFinalizer is a mock of TrafficPolicyFinalizer interface.
type MockTrafficPolicyFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyFinalizerMockRecorder
}

// MockTrafficPolicyFinalizerMockRecorder is the mock recorder for MockTrafficPolicyFinalizer.
type MockTrafficPolicyFinalizerMockRecorder struct {
	mock *MockTrafficPolicyFinalizer
}

// NewMockTrafficPolicyFinalizer creates a new mock instance.
func NewMockTrafficPolicyFinalizer(ctrl *gomock.Controller) *MockTrafficPolicyFinalizer {
	mock := &MockTrafficPolicyFinalizer{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyFinalizer) EXPECT() *MockTrafficPolicyFinalizerMockRecorder {
	return m.recorder
}

// FinalizeTrafficPolicy mocks base method.
func (m *MockTrafficPolicyFinalizer) FinalizeTrafficPolicy(obj *v1.TrafficPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeTrafficPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeTrafficPolicy indicates an expected call of FinalizeTrafficPolicy.
func (mr *MockTrafficPolicyFinalizerMockRecorder) FinalizeTrafficPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeTrafficPolicy", reflect.TypeOf((*MockTrafficPolicyFinalizer)(nil).FinalizeTrafficPolicy), obj)
}

// ReconcileTrafficPolicy mocks base method.
func (m *MockTrafficPolicyFinalizer) ReconcileTrafficPolicy(obj *v1.TrafficPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficPolicy indicates an expected call of ReconcileTrafficPolicy.
func (mr *MockTrafficPolicyFinalizerMockRecorder) ReconcileTrafficPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicy", reflect.TypeOf((*MockTrafficPolicyFinalizer)(nil).ReconcileTrafficPolicy), obj)
}

// TrafficPolicyFinalizerName mocks base method.
func (m *MockTrafficPolicyFinalizer) TrafficPolicyFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficPolicyFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TrafficPolicyFinalizerName indicates an expected call of TrafficPolicyFinalizerName.
func (mr *MockTrafficPolicyFinalizerMockRecorder) TrafficPolicyFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficPolicyFinalizerName", reflect.TypeOf((*MockTrafficPolicyFinalizer)(nil).TrafficPolicyFinalizerName))
}

// MockTrafficPolicyReconcileLoop is a mock of TrafficPolicyReconcileLoop interface.
type MockTrafficPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyReconcileLoopMockRecorder
}

// MockTrafficPolicyReconcileLoopMockRecorder is the mock recorder for MockTrafficPolicyReconcileLoop.
type MockTrafficPolicyReconcileLoopMockRecorder struct {
	mock *MockTrafficPolicyReconcileLoop
}

// NewMockTrafficPolicyReconcileLoop creates a new mock instance.
func NewMockTrafficPolicyReconcileLoop(ctrl *gomock.Controller) *MockTrafficPolicyReconcileLoop {
	mock := &MockTrafficPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyReconcileLoop) EXPECT() *MockTrafficPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// RunTrafficPolicyReconciler mocks base method.
func (m *MockTrafficPolicyReconcileLoop) RunTrafficPolicyReconciler(ctx context.Context, rec controller.TrafficPolicyReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunTrafficPolicyReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTrafficPolicyReconciler indicates an expected call of RunTrafficPolicyReconciler.
func (mr *MockTrafficPolicyReconcileLoopMockRecorder) RunTrafficPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTrafficPolicyReconciler", reflect.TypeOf((*MockTrafficPolicyReconcileLoop)(nil).RunTrafficPolicyReconciler), varargs...)
}

// MockAccessPolicyReconciler is a mock of AccessPolicyReconciler interface.
type MockAccessPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAccessPolicyReconcilerMockRecorder
}

// MockAccessPolicyReconcilerMockRecorder is the mock recorder for MockAccessPolicyReconciler.
type MockAccessPolicyReconcilerMockRecorder struct {
	mock *MockAccessPolicyReconciler
}

// NewMockAccessPolicyReconciler creates a new mock instance.
func NewMockAccessPolicyReconciler(ctrl *gomock.Controller) *MockAccessPolicyReconciler {
	mock := &MockAccessPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockAccessPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessPolicyReconciler) EXPECT() *MockAccessPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessPolicy mocks base method.
func (m *MockAccessPolicyReconciler) ReconcileAccessPolicy(obj *v1.AccessPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessPolicy indicates an expected call of ReconcileAccessPolicy.
func (mr *MockAccessPolicyReconcilerMockRecorder) ReconcileAccessPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessPolicy", reflect.TypeOf((*MockAccessPolicyReconciler)(nil).ReconcileAccessPolicy), obj)
}

// MockAccessPolicyDeletionReconciler is a mock of AccessPolicyDeletionReconciler interface.
type MockAccessPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAccessPolicyDeletionReconcilerMockRecorder
}

// MockAccessPolicyDeletionReconcilerMockRecorder is the mock recorder for MockAccessPolicyDeletionReconciler.
type MockAccessPolicyDeletionReconcilerMockRecorder struct {
	mock *MockAccessPolicyDeletionReconciler
}

// NewMockAccessPolicyDeletionReconciler creates a new mock instance.
func NewMockAccessPolicyDeletionReconciler(ctrl *gomock.Controller) *MockAccessPolicyDeletionReconciler {
	mock := &MockAccessPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockAccessPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessPolicyDeletionReconciler) EXPECT() *MockAccessPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessPolicyDeletion mocks base method.
func (m *MockAccessPolicyDeletionReconciler) ReconcileAccessPolicyDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessPolicyDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAccessPolicyDeletion indicates an expected call of ReconcileAccessPolicyDeletion.
func (mr *MockAccessPolicyDeletionReconcilerMockRecorder) ReconcileAccessPolicyDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessPolicyDeletion", reflect.TypeOf((*MockAccessPolicyDeletionReconciler)(nil).ReconcileAccessPolicyDeletion), req)
}

// MockAccessPolicyFinalizer is a mock of AccessPolicyFinalizer interface.
type MockAccessPolicyFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockAccessPolicyFinalizerMockRecorder
}

// MockAccessPolicyFinalizerMockRecorder is the mock recorder for MockAccessPolicyFinalizer.
type MockAccessPolicyFinalizerMockRecorder struct {
	mock *MockAccessPolicyFinalizer
}

// NewMockAccessPolicyFinalizer creates a new mock instance.
func NewMockAccessPolicyFinalizer(ctrl *gomock.Controller) *MockAccessPolicyFinalizer {
	mock := &MockAccessPolicyFinalizer{ctrl: ctrl}
	mock.recorder = &MockAccessPolicyFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessPolicyFinalizer) EXPECT() *MockAccessPolicyFinalizerMockRecorder {
	return m.recorder
}

// AccessPolicyFinalizerName mocks base method.
func (m *MockAccessPolicyFinalizer) AccessPolicyFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessPolicyFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// AccessPolicyFinalizerName indicates an expected call of AccessPolicyFinalizerName.
func (mr *MockAccessPolicyFinalizerMockRecorder) AccessPolicyFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessPolicyFinalizerName", reflect.TypeOf((*MockAccessPolicyFinalizer)(nil).AccessPolicyFinalizerName))
}

// FinalizeAccessPolicy mocks base method.
func (m *MockAccessPolicyFinalizer) FinalizeAccessPolicy(obj *v1.AccessPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeAccessPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeAccessPolicy indicates an expected call of FinalizeAccessPolicy.
func (mr *MockAccessPolicyFinalizerMockRecorder) FinalizeAccessPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeAccessPolicy", reflect.TypeOf((*MockAccessPolicyFinalizer)(nil).FinalizeAccessPolicy), obj)
}

// ReconcileAccessPolicy mocks base method.
func (m *MockAccessPolicyFinalizer) ReconcileAccessPolicy(obj *v1.AccessPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessPolicy indicates an expected call of ReconcileAccessPolicy.
func (mr *MockAccessPolicyFinalizerMockRecorder) ReconcileAccessPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessPolicy", reflect.TypeOf((*MockAccessPolicyFinalizer)(nil).ReconcileAccessPolicy), obj)
}

// MockAccessPolicyReconcileLoop is a mock of AccessPolicyReconcileLoop interface.
type MockAccessPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockAccessPolicyReconcileLoopMockRecorder
}

// MockAccessPolicyReconcileLoopMockRecorder is the mock recorder for MockAccessPolicyReconcileLoop.
type MockAccessPolicyReconcileLoopMockRecorder struct {
	mock *MockAccessPolicyReconcileLoop
}

// NewMockAccessPolicyReconcileLoop creates a new mock instance.
func NewMockAccessPolicyReconcileLoop(ctrl *gomock.Controller) *MockAccessPolicyReconcileLoop {
	mock := &MockAccessPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockAccessPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessPolicyReconcileLoop) EXPECT() *MockAccessPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// RunAccessPolicyReconciler mocks base method.
func (m *MockAccessPolicyReconcileLoop) RunAccessPolicyReconciler(ctx context.Context, rec controller.AccessPolicyReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunAccessPolicyReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunAccessPolicyReconciler indicates an expected call of RunAccessPolicyReconciler.
func (mr *MockAccessPolicyReconcileLoopMockRecorder) RunAccessPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAccessPolicyReconciler", reflect.TypeOf((*MockAccessPolicyReconcileLoop)(nil).RunAccessPolicyReconciler), varargs...)
}

// MockVirtualMeshReconciler is a mock of VirtualMeshReconciler interface.
type MockVirtualMeshReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshReconcilerMockRecorder
}

// MockVirtualMeshReconcilerMockRecorder is the mock recorder for MockVirtualMeshReconciler.
type MockVirtualMeshReconcilerMockRecorder struct {
	mock *MockVirtualMeshReconciler
}

// NewMockVirtualMeshReconciler creates a new mock instance.
func NewMockVirtualMeshReconciler(ctrl *gomock.Controller) *MockVirtualMeshReconciler {
	mock := &MockVirtualMeshReconciler{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshReconciler) EXPECT() *MockVirtualMeshReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualMesh mocks base method.
func (m *MockVirtualMeshReconciler) ReconcileVirtualMesh(obj *v1.VirtualMesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMesh", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualMesh indicates an expected call of ReconcileVirtualMesh.
func (mr *MockVirtualMeshReconcilerMockRecorder) ReconcileVirtualMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMesh", reflect.TypeOf((*MockVirtualMeshReconciler)(nil).ReconcileVirtualMesh), obj)
}

// MockVirtualMeshDeletionReconciler is a mock of VirtualMeshDeletionReconciler interface.
type MockVirtualMeshDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshDeletionReconcilerMockRecorder
}

// MockVirtualMeshDeletionReconcilerMockRecorder is the mock recorder for MockVirtualMeshDeletionReconciler.
type MockVirtualMeshDeletionReconcilerMockRecorder struct {
	mock *MockVirtualMeshDeletionReconciler
}

// NewMockVirtualMeshDeletionReconciler creates a new mock instance.
func NewMockVirtualMeshDeletionReconciler(ctrl *gomock.Controller) *MockVirtualMeshDeletionReconciler {
	mock := &MockVirtualMeshDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshDeletionReconciler) EXPECT() *MockVirtualMeshDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualMeshDeletion mocks base method.
func (m *MockVirtualMeshDeletionReconciler) ReconcileVirtualMeshDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMeshDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualMeshDeletion indicates an expected call of ReconcileVirtualMeshDeletion.
func (mr *MockVirtualMeshDeletionReconcilerMockRecorder) ReconcileVirtualMeshDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMeshDeletion", reflect.TypeOf((*MockVirtualMeshDeletionReconciler)(nil).ReconcileVirtualMeshDeletion), req)
}

// MockVirtualMeshFinalizer is a mock of VirtualMeshFinalizer interface.
type MockVirtualMeshFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshFinalizerMockRecorder
}

// MockVirtualMeshFinalizerMockRecorder is the mock recorder for MockVirtualMeshFinalizer.
type MockVirtualMeshFinalizerMockRecorder struct {
	mock *MockVirtualMeshFinalizer
}

// NewMockVirtualMeshFinalizer creates a new mock instance.
func NewMockVirtualMeshFinalizer(ctrl *gomock.Controller) *MockVirtualMeshFinalizer {
	mock := &MockVirtualMeshFinalizer{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshFinalizer) EXPECT() *MockVirtualMeshFinalizerMockRecorder {
	return m.recorder
}

// FinalizeVirtualMesh mocks base method.
func (m *MockVirtualMeshFinalizer) FinalizeVirtualMesh(obj *v1.VirtualMesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeVirtualMesh", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeVirtualMesh indicates an expected call of FinalizeVirtualMesh.
func (mr *MockVirtualMeshFinalizerMockRecorder) FinalizeVirtualMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeVirtualMesh", reflect.TypeOf((*MockVirtualMeshFinalizer)(nil).FinalizeVirtualMesh), obj)
}

// ReconcileVirtualMesh mocks base method.
func (m *MockVirtualMeshFinalizer) ReconcileVirtualMesh(obj *v1.VirtualMesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMesh", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualMesh indicates an expected call of ReconcileVirtualMesh.
func (mr *MockVirtualMeshFinalizerMockRecorder) ReconcileVirtualMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMesh", reflect.TypeOf((*MockVirtualMeshFinalizer)(nil).ReconcileVirtualMesh), obj)
}

// VirtualMeshFinalizerName mocks base method.
func (m *MockVirtualMeshFinalizer) VirtualMeshFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualMeshFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// VirtualMeshFinalizerName indicates an expected call of VirtualMeshFinalizerName.
func (mr *MockVirtualMeshFinalizerMockRecorder) VirtualMeshFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualMeshFinalizerName", reflect.TypeOf((*MockVirtualMeshFinalizer)(nil).VirtualMeshFinalizerName))
}

// MockVirtualMeshReconcileLoop is a mock of VirtualMeshReconcileLoop interface.
type MockVirtualMeshReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshReconcileLoopMockRecorder
}

// MockVirtualMeshReconcileLoopMockRecorder is the mock recorder for MockVirtualMeshReconcileLoop.
type MockVirtualMeshReconcileLoopMockRecorder struct {
	mock *MockVirtualMeshReconcileLoop
}

// NewMockVirtualMeshReconcileLoop creates a new mock instance.
func NewMockVirtualMeshReconcileLoop(ctrl *gomock.Controller) *MockVirtualMeshReconcileLoop {
	mock := &MockVirtualMeshReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshReconcileLoop) EXPECT() *MockVirtualMeshReconcileLoopMockRecorder {
	return m.recorder
}

// RunVirtualMeshReconciler mocks base method.
func (m *MockVirtualMeshReconcileLoop) RunVirtualMeshReconciler(ctx context.Context, rec controller.VirtualMeshReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunVirtualMeshReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunVirtualMeshReconciler indicates an expected call of RunVirtualMeshReconciler.
func (mr *MockVirtualMeshReconcileLoopMockRecorder) RunVirtualMeshReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunVirtualMeshReconciler", reflect.TypeOf((*MockVirtualMeshReconcileLoop)(nil).RunVirtualMeshReconciler), varargs...)
}
