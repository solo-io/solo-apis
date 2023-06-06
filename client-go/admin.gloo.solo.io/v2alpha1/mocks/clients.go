// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v2alpha1 is a generated GoMock package.
package mock_v2alpha1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	v2alpha1 "github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2alpha1"
)

// MockMulticlusterClientset is a mock of MulticlusterClientset interface.
type MockMulticlusterClientset struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClientsetMockRecorder
}

// MockMulticlusterClientsetMockRecorder is the mock recorder for MockMulticlusterClientset.
type MockMulticlusterClientsetMockRecorder struct {
	mock *MockMulticlusterClientset
}

// NewMockMulticlusterClientset creates a new mock instance.
func NewMockMulticlusterClientset(ctrl *gomock.Controller) *MockMulticlusterClientset {
	mock := &MockMulticlusterClientset{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterClientset) EXPECT() *MockMulticlusterClientsetMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterClientset) Cluster(cluster string) (v2alpha1.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v2alpha1.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterClientsetMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterClientset)(nil).Cluster), cluster)
}

// MockClientset is a mock of Clientset interface.
type MockClientset struct {
	ctrl     *gomock.Controller
	recorder *MockClientsetMockRecorder
}

// MockClientsetMockRecorder is the mock recorder for MockClientset.
type MockClientsetMockRecorder struct {
	mock *MockClientset
}

// NewMockClientset creates a new mock instance.
func NewMockClientset(ctrl *gomock.Controller) *MockClientset {
	mock := &MockClientset{ctrl: ctrl}
	mock.recorder = &MockClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientset) EXPECT() *MockClientsetMockRecorder {
	return m.recorder
}

// WaypointLifecycleManagers mocks base method.
func (m *MockClientset) WaypointLifecycleManagers() v2alpha1.WaypointLifecycleManagerClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaypointLifecycleManagers")
	ret0, _ := ret[0].(v2alpha1.WaypointLifecycleManagerClient)
	return ret0
}

// WaypointLifecycleManagers indicates an expected call of WaypointLifecycleManagers.
func (mr *MockClientsetMockRecorder) WaypointLifecycleManagers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaypointLifecycleManagers", reflect.TypeOf((*MockClientset)(nil).WaypointLifecycleManagers))
}

// MockWaypointLifecycleManagerReader is a mock of WaypointLifecycleManagerReader interface.
type MockWaypointLifecycleManagerReader struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerReaderMockRecorder
}

// MockWaypointLifecycleManagerReaderMockRecorder is the mock recorder for MockWaypointLifecycleManagerReader.
type MockWaypointLifecycleManagerReaderMockRecorder struct {
	mock *MockWaypointLifecycleManagerReader
}

// NewMockWaypointLifecycleManagerReader creates a new mock instance.
func NewMockWaypointLifecycleManagerReader(ctrl *gomock.Controller) *MockWaypointLifecycleManagerReader {
	mock := &MockWaypointLifecycleManagerReader{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerReader) EXPECT() *MockWaypointLifecycleManagerReaderMockRecorder {
	return m.recorder
}

// GetWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerReader) GetWaypointLifecycleManager(ctx context.Context, key client.ObjectKey) (*v2alpha1.WaypointLifecycleManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWaypointLifecycleManager", ctx, key)
	ret0, _ := ret[0].(*v2alpha1.WaypointLifecycleManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWaypointLifecycleManager indicates an expected call of GetWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerReaderMockRecorder) GetWaypointLifecycleManager(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerReader)(nil).GetWaypointLifecycleManager), ctx, key)
}

// ListWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerReader) ListWaypointLifecycleManager(ctx context.Context, opts ...client.ListOption) (*v2alpha1.WaypointLifecycleManagerList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(*v2alpha1.WaypointLifecycleManagerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWaypointLifecycleManager indicates an expected call of ListWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerReaderMockRecorder) ListWaypointLifecycleManager(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerReader)(nil).ListWaypointLifecycleManager), varargs...)
}

// MockWaypointLifecycleManagerWriter is a mock of WaypointLifecycleManagerWriter interface.
type MockWaypointLifecycleManagerWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerWriterMockRecorder
}

// MockWaypointLifecycleManagerWriterMockRecorder is the mock recorder for MockWaypointLifecycleManagerWriter.
type MockWaypointLifecycleManagerWriterMockRecorder struct {
	mock *MockWaypointLifecycleManagerWriter
}

// NewMockWaypointLifecycleManagerWriter creates a new mock instance.
func NewMockWaypointLifecycleManagerWriter(ctrl *gomock.Controller) *MockWaypointLifecycleManagerWriter {
	mock := &MockWaypointLifecycleManagerWriter{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerWriter) EXPECT() *MockWaypointLifecycleManagerWriterMockRecorder {
	return m.recorder
}

// CreateWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerWriter) CreateWaypointLifecycleManager(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWaypointLifecycleManager indicates an expected call of CreateWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerWriterMockRecorder) CreateWaypointLifecycleManager(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerWriter)(nil).CreateWaypointLifecycleManager), varargs...)
}

// DeleteAllOfWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerWriter) DeleteAllOfWaypointLifecycleManager(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfWaypointLifecycleManager indicates an expected call of DeleteAllOfWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerWriterMockRecorder) DeleteAllOfWaypointLifecycleManager(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerWriter)(nil).DeleteAllOfWaypointLifecycleManager), varargs...)
}

// DeleteWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerWriter) DeleteWaypointLifecycleManager(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWaypointLifecycleManager indicates an expected call of DeleteWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerWriterMockRecorder) DeleteWaypointLifecycleManager(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerWriter)(nil).DeleteWaypointLifecycleManager), varargs...)
}

// PatchWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerWriter) PatchWaypointLifecycleManager(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchWaypointLifecycleManager indicates an expected call of PatchWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerWriterMockRecorder) PatchWaypointLifecycleManager(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerWriter)(nil).PatchWaypointLifecycleManager), varargs...)
}

// UpdateWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerWriter) UpdateWaypointLifecycleManager(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWaypointLifecycleManager indicates an expected call of UpdateWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerWriterMockRecorder) UpdateWaypointLifecycleManager(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerWriter)(nil).UpdateWaypointLifecycleManager), varargs...)
}

// UpsertWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerWriter) UpsertWaypointLifecycleManager(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, transitionFuncs ...v2alpha1.WaypointLifecycleManagerTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWaypointLifecycleManager indicates an expected call of UpsertWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerWriterMockRecorder) UpsertWaypointLifecycleManager(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerWriter)(nil).UpsertWaypointLifecycleManager), varargs...)
}

// MockWaypointLifecycleManagerStatusWriter is a mock of WaypointLifecycleManagerStatusWriter interface.
type MockWaypointLifecycleManagerStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerStatusWriterMockRecorder
}

// MockWaypointLifecycleManagerStatusWriterMockRecorder is the mock recorder for MockWaypointLifecycleManagerStatusWriter.
type MockWaypointLifecycleManagerStatusWriterMockRecorder struct {
	mock *MockWaypointLifecycleManagerStatusWriter
}

// NewMockWaypointLifecycleManagerStatusWriter creates a new mock instance.
func NewMockWaypointLifecycleManagerStatusWriter(ctrl *gomock.Controller) *MockWaypointLifecycleManagerStatusWriter {
	mock := &MockWaypointLifecycleManagerStatusWriter{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerStatusWriter) EXPECT() *MockWaypointLifecycleManagerStatusWriterMockRecorder {
	return m.recorder
}

// PatchWaypointLifecycleManagerStatus mocks base method.
func (m *MockWaypointLifecycleManagerStatusWriter) PatchWaypointLifecycleManagerStatus(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchWaypointLifecycleManagerStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchWaypointLifecycleManagerStatus indicates an expected call of PatchWaypointLifecycleManagerStatus.
func (mr *MockWaypointLifecycleManagerStatusWriterMockRecorder) PatchWaypointLifecycleManagerStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchWaypointLifecycleManagerStatus", reflect.TypeOf((*MockWaypointLifecycleManagerStatusWriter)(nil).PatchWaypointLifecycleManagerStatus), varargs...)
}

// UpdateWaypointLifecycleManagerStatus mocks base method.
func (m *MockWaypointLifecycleManagerStatusWriter) UpdateWaypointLifecycleManagerStatus(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateWaypointLifecycleManagerStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWaypointLifecycleManagerStatus indicates an expected call of UpdateWaypointLifecycleManagerStatus.
func (mr *MockWaypointLifecycleManagerStatusWriterMockRecorder) UpdateWaypointLifecycleManagerStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWaypointLifecycleManagerStatus", reflect.TypeOf((*MockWaypointLifecycleManagerStatusWriter)(nil).UpdateWaypointLifecycleManagerStatus), varargs...)
}

// MockWaypointLifecycleManagerClient is a mock of WaypointLifecycleManagerClient interface.
type MockWaypointLifecycleManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockWaypointLifecycleManagerClientMockRecorder
}

// MockWaypointLifecycleManagerClientMockRecorder is the mock recorder for MockWaypointLifecycleManagerClient.
type MockWaypointLifecycleManagerClientMockRecorder struct {
	mock *MockWaypointLifecycleManagerClient
}

// NewMockWaypointLifecycleManagerClient creates a new mock instance.
func NewMockWaypointLifecycleManagerClient(ctrl *gomock.Controller) *MockWaypointLifecycleManagerClient {
	mock := &MockWaypointLifecycleManagerClient{ctrl: ctrl}
	mock.recorder = &MockWaypointLifecycleManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaypointLifecycleManagerClient) EXPECT() *MockWaypointLifecycleManagerClientMockRecorder {
	return m.recorder
}

// CreateWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerClient) CreateWaypointLifecycleManager(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWaypointLifecycleManager indicates an expected call of CreateWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) CreateWaypointLifecycleManager(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).CreateWaypointLifecycleManager), varargs...)
}

// DeleteAllOfWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerClient) DeleteAllOfWaypointLifecycleManager(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfWaypointLifecycleManager indicates an expected call of DeleteAllOfWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) DeleteAllOfWaypointLifecycleManager(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).DeleteAllOfWaypointLifecycleManager), varargs...)
}

// DeleteWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerClient) DeleteWaypointLifecycleManager(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWaypointLifecycleManager indicates an expected call of DeleteWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) DeleteWaypointLifecycleManager(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).DeleteWaypointLifecycleManager), varargs...)
}

// GetWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerClient) GetWaypointLifecycleManager(ctx context.Context, key client.ObjectKey) (*v2alpha1.WaypointLifecycleManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWaypointLifecycleManager", ctx, key)
	ret0, _ := ret[0].(*v2alpha1.WaypointLifecycleManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWaypointLifecycleManager indicates an expected call of GetWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) GetWaypointLifecycleManager(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).GetWaypointLifecycleManager), ctx, key)
}

// ListWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerClient) ListWaypointLifecycleManager(ctx context.Context, opts ...client.ListOption) (*v2alpha1.WaypointLifecycleManagerList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(*v2alpha1.WaypointLifecycleManagerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWaypointLifecycleManager indicates an expected call of ListWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) ListWaypointLifecycleManager(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).ListWaypointLifecycleManager), varargs...)
}

// PatchWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerClient) PatchWaypointLifecycleManager(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchWaypointLifecycleManager indicates an expected call of PatchWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) PatchWaypointLifecycleManager(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).PatchWaypointLifecycleManager), varargs...)
}

// PatchWaypointLifecycleManagerStatus mocks base method.
func (m *MockWaypointLifecycleManagerClient) PatchWaypointLifecycleManagerStatus(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchWaypointLifecycleManagerStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchWaypointLifecycleManagerStatus indicates an expected call of PatchWaypointLifecycleManagerStatus.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) PatchWaypointLifecycleManagerStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchWaypointLifecycleManagerStatus", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).PatchWaypointLifecycleManagerStatus), varargs...)
}

// UpdateWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerClient) UpdateWaypointLifecycleManager(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWaypointLifecycleManager indicates an expected call of UpdateWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) UpdateWaypointLifecycleManager(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).UpdateWaypointLifecycleManager), varargs...)
}

// UpdateWaypointLifecycleManagerStatus mocks base method.
func (m *MockWaypointLifecycleManagerClient) UpdateWaypointLifecycleManagerStatus(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateWaypointLifecycleManagerStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWaypointLifecycleManagerStatus indicates an expected call of UpdateWaypointLifecycleManagerStatus.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) UpdateWaypointLifecycleManagerStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWaypointLifecycleManagerStatus", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).UpdateWaypointLifecycleManagerStatus), varargs...)
}

// UpsertWaypointLifecycleManager mocks base method.
func (m *MockWaypointLifecycleManagerClient) UpsertWaypointLifecycleManager(ctx context.Context, obj *v2alpha1.WaypointLifecycleManager, transitionFuncs ...v2alpha1.WaypointLifecycleManagerTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertWaypointLifecycleManager", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWaypointLifecycleManager indicates an expected call of UpsertWaypointLifecycleManager.
func (mr *MockWaypointLifecycleManagerClientMockRecorder) UpsertWaypointLifecycleManager(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWaypointLifecycleManager", reflect.TypeOf((*MockWaypointLifecycleManagerClient)(nil).UpsertWaypointLifecycleManager), varargs...)
}

// MockMulticlusterWaypointLifecycleManagerClient is a mock of MulticlusterWaypointLifecycleManagerClient interface.
type MockMulticlusterWaypointLifecycleManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterWaypointLifecycleManagerClientMockRecorder
}

// MockMulticlusterWaypointLifecycleManagerClientMockRecorder is the mock recorder for MockMulticlusterWaypointLifecycleManagerClient.
type MockMulticlusterWaypointLifecycleManagerClientMockRecorder struct {
	mock *MockMulticlusterWaypointLifecycleManagerClient
}

// NewMockMulticlusterWaypointLifecycleManagerClient creates a new mock instance.
func NewMockMulticlusterWaypointLifecycleManagerClient(ctrl *gomock.Controller) *MockMulticlusterWaypointLifecycleManagerClient {
	mock := &MockMulticlusterWaypointLifecycleManagerClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterWaypointLifecycleManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterWaypointLifecycleManagerClient) EXPECT() *MockMulticlusterWaypointLifecycleManagerClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterWaypointLifecycleManagerClient) Cluster(cluster string) (v2alpha1.WaypointLifecycleManagerClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v2alpha1.WaypointLifecycleManagerClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterWaypointLifecycleManagerClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterWaypointLifecycleManagerClient)(nil).Cluster), cluster)
}
