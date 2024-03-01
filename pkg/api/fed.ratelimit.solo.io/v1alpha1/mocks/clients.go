// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1alpha1 is a generated GoMock package.
package mock_v1alpha1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/fed.ratelimit.solo.io/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockMulticlusterClientset is a mock of MulticlusterClientset interface
type MockMulticlusterClientset struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClientsetMockRecorder
}

// MockMulticlusterClientsetMockRecorder is the mock recorder for MockMulticlusterClientset
type MockMulticlusterClientsetMockRecorder struct {
	mock *MockMulticlusterClientset
}

// NewMockMulticlusterClientset creates a new mock instance
func NewMockMulticlusterClientset(ctrl *gomock.Controller) *MockMulticlusterClientset {
	mock := &MockMulticlusterClientset{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClientset) EXPECT() *MockMulticlusterClientsetMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1alpha1.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha1.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterClientsetMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterClientset)(nil).Cluster), cluster)
}

// MockClientset is a mock of Clientset interface
type MockClientset struct {
	ctrl     *gomock.Controller
	recorder *MockClientsetMockRecorder
}

// MockClientsetMockRecorder is the mock recorder for MockClientset
type MockClientsetMockRecorder struct {
	mock *MockClientset
}

// NewMockClientset creates a new mock instance
func NewMockClientset(ctrl *gomock.Controller) *MockClientset {
	mock := &MockClientset{ctrl: ctrl}
	mock.recorder = &MockClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientset) EXPECT() *MockClientsetMockRecorder {
	return m.recorder
}

// FederatedRateLimitConfigs mocks base method
func (m *MockClientset) FederatedRateLimitConfigs() v1alpha1.FederatedRateLimitConfigClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedRateLimitConfigs")
	ret0, _ := ret[0].(v1alpha1.FederatedRateLimitConfigClient)
	return ret0
}

// FederatedRateLimitConfigs indicates an expected call of FederatedRateLimitConfigs
func (mr *MockClientsetMockRecorder) FederatedRateLimitConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedRateLimitConfigs", reflect.TypeOf((*MockClientset)(nil).FederatedRateLimitConfigs))
}

// MockFederatedRateLimitConfigReader is a mock of FederatedRateLimitConfigReader interface
type MockFederatedRateLimitConfigReader struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRateLimitConfigReaderMockRecorder
}

// MockFederatedRateLimitConfigReaderMockRecorder is the mock recorder for MockFederatedRateLimitConfigReader
type MockFederatedRateLimitConfigReaderMockRecorder struct {
	mock *MockFederatedRateLimitConfigReader
}

// NewMockFederatedRateLimitConfigReader creates a new mock instance
func NewMockFederatedRateLimitConfigReader(ctrl *gomock.Controller) *MockFederatedRateLimitConfigReader {
	mock := &MockFederatedRateLimitConfigReader{ctrl: ctrl}
	mock.recorder = &MockFederatedRateLimitConfigReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedRateLimitConfigReader) EXPECT() *MockFederatedRateLimitConfigReaderMockRecorder {
	return m.recorder
}

// GetFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigReader) GetFederatedRateLimitConfig(ctx context.Context, key client.ObjectKey) (*v1alpha1.FederatedRateLimitConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederatedRateLimitConfig", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.FederatedRateLimitConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederatedRateLimitConfig indicates an expected call of GetFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigReaderMockRecorder) GetFederatedRateLimitConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigReader)(nil).GetFederatedRateLimitConfig), ctx, key)
}

// ListFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigReader) ListFederatedRateLimitConfig(ctx context.Context, opts ...client.ListOption) (*v1alpha1.FederatedRateLimitConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(*v1alpha1.FederatedRateLimitConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFederatedRateLimitConfig indicates an expected call of ListFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigReaderMockRecorder) ListFederatedRateLimitConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigReader)(nil).ListFederatedRateLimitConfig), varargs...)
}

// MockFederatedRateLimitConfigWriter is a mock of FederatedRateLimitConfigWriter interface
type MockFederatedRateLimitConfigWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRateLimitConfigWriterMockRecorder
}

// MockFederatedRateLimitConfigWriterMockRecorder is the mock recorder for MockFederatedRateLimitConfigWriter
type MockFederatedRateLimitConfigWriterMockRecorder struct {
	mock *MockFederatedRateLimitConfigWriter
}

// NewMockFederatedRateLimitConfigWriter creates a new mock instance
func NewMockFederatedRateLimitConfigWriter(ctrl *gomock.Controller) *MockFederatedRateLimitConfigWriter {
	mock := &MockFederatedRateLimitConfigWriter{ctrl: ctrl}
	mock.recorder = &MockFederatedRateLimitConfigWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedRateLimitConfigWriter) EXPECT() *MockFederatedRateLimitConfigWriterMockRecorder {
	return m.recorder
}

// CreateFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigWriter) CreateFederatedRateLimitConfig(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedRateLimitConfig indicates an expected call of CreateFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigWriterMockRecorder) CreateFederatedRateLimitConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigWriter)(nil).CreateFederatedRateLimitConfig), varargs...)
}

// DeleteFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigWriter) DeleteFederatedRateLimitConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedRateLimitConfig indicates an expected call of DeleteFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigWriterMockRecorder) DeleteFederatedRateLimitConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigWriter)(nil).DeleteFederatedRateLimitConfig), varargs...)
}

// UpdateFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigWriter) UpdateFederatedRateLimitConfig(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedRateLimitConfig indicates an expected call of UpdateFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigWriterMockRecorder) UpdateFederatedRateLimitConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigWriter)(nil).UpdateFederatedRateLimitConfig), varargs...)
}

// PatchFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigWriter) PatchFederatedRateLimitConfig(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchFederatedRateLimitConfig indicates an expected call of PatchFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigWriterMockRecorder) PatchFederatedRateLimitConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigWriter)(nil).PatchFederatedRateLimitConfig), varargs...)
}

// DeleteAllOfFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigWriter) DeleteAllOfFederatedRateLimitConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfFederatedRateLimitConfig indicates an expected call of DeleteAllOfFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigWriterMockRecorder) DeleteAllOfFederatedRateLimitConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigWriter)(nil).DeleteAllOfFederatedRateLimitConfig), varargs...)
}

// UpsertFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigWriter) UpsertFederatedRateLimitConfig(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, transitionFuncs ...v1alpha1.FederatedRateLimitConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertFederatedRateLimitConfig indicates an expected call of UpsertFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigWriterMockRecorder) UpsertFederatedRateLimitConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigWriter)(nil).UpsertFederatedRateLimitConfig), varargs...)
}

// MockFederatedRateLimitConfigStatusWriter is a mock of FederatedRateLimitConfigStatusWriter interface
type MockFederatedRateLimitConfigStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRateLimitConfigStatusWriterMockRecorder
}

// MockFederatedRateLimitConfigStatusWriterMockRecorder is the mock recorder for MockFederatedRateLimitConfigStatusWriter
type MockFederatedRateLimitConfigStatusWriterMockRecorder struct {
	mock *MockFederatedRateLimitConfigStatusWriter
}

// NewMockFederatedRateLimitConfigStatusWriter creates a new mock instance
func NewMockFederatedRateLimitConfigStatusWriter(ctrl *gomock.Controller) *MockFederatedRateLimitConfigStatusWriter {
	mock := &MockFederatedRateLimitConfigStatusWriter{ctrl: ctrl}
	mock.recorder = &MockFederatedRateLimitConfigStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedRateLimitConfigStatusWriter) EXPECT() *MockFederatedRateLimitConfigStatusWriterMockRecorder {
	return m.recorder
}

// UpdateFederatedRateLimitConfigStatus mocks base method
func (m *MockFederatedRateLimitConfigStatusWriter) UpdateFederatedRateLimitConfigStatus(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedRateLimitConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedRateLimitConfigStatus indicates an expected call of UpdateFederatedRateLimitConfigStatus
func (mr *MockFederatedRateLimitConfigStatusWriterMockRecorder) UpdateFederatedRateLimitConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedRateLimitConfigStatus", reflect.TypeOf((*MockFederatedRateLimitConfigStatusWriter)(nil).UpdateFederatedRateLimitConfigStatus), varargs...)
}

// PatchFederatedRateLimitConfigStatus mocks base method
func (m *MockFederatedRateLimitConfigStatusWriter) PatchFederatedRateLimitConfigStatus(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchFederatedRateLimitConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchFederatedRateLimitConfigStatus indicates an expected call of PatchFederatedRateLimitConfigStatus
func (mr *MockFederatedRateLimitConfigStatusWriterMockRecorder) PatchFederatedRateLimitConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchFederatedRateLimitConfigStatus", reflect.TypeOf((*MockFederatedRateLimitConfigStatusWriter)(nil).PatchFederatedRateLimitConfigStatus), varargs...)
}

// MockFederatedRateLimitConfigClient is a mock of FederatedRateLimitConfigClient interface
type MockFederatedRateLimitConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRateLimitConfigClientMockRecorder
}

// MockFederatedRateLimitConfigClientMockRecorder is the mock recorder for MockFederatedRateLimitConfigClient
type MockFederatedRateLimitConfigClientMockRecorder struct {
	mock *MockFederatedRateLimitConfigClient
}

// NewMockFederatedRateLimitConfigClient creates a new mock instance
func NewMockFederatedRateLimitConfigClient(ctrl *gomock.Controller) *MockFederatedRateLimitConfigClient {
	mock := &MockFederatedRateLimitConfigClient{ctrl: ctrl}
	mock.recorder = &MockFederatedRateLimitConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedRateLimitConfigClient) EXPECT() *MockFederatedRateLimitConfigClientMockRecorder {
	return m.recorder
}

// GetFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigClient) GetFederatedRateLimitConfig(ctx context.Context, key client.ObjectKey) (*v1alpha1.FederatedRateLimitConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederatedRateLimitConfig", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.FederatedRateLimitConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederatedRateLimitConfig indicates an expected call of GetFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigClientMockRecorder) GetFederatedRateLimitConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).GetFederatedRateLimitConfig), ctx, key)
}

// ListFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigClient) ListFederatedRateLimitConfig(ctx context.Context, opts ...client.ListOption) (*v1alpha1.FederatedRateLimitConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(*v1alpha1.FederatedRateLimitConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFederatedRateLimitConfig indicates an expected call of ListFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigClientMockRecorder) ListFederatedRateLimitConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).ListFederatedRateLimitConfig), varargs...)
}

// CreateFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigClient) CreateFederatedRateLimitConfig(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedRateLimitConfig indicates an expected call of CreateFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigClientMockRecorder) CreateFederatedRateLimitConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).CreateFederatedRateLimitConfig), varargs...)
}

// DeleteFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigClient) DeleteFederatedRateLimitConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedRateLimitConfig indicates an expected call of DeleteFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigClientMockRecorder) DeleteFederatedRateLimitConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).DeleteFederatedRateLimitConfig), varargs...)
}

// UpdateFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigClient) UpdateFederatedRateLimitConfig(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedRateLimitConfig indicates an expected call of UpdateFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigClientMockRecorder) UpdateFederatedRateLimitConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).UpdateFederatedRateLimitConfig), varargs...)
}

// PatchFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigClient) PatchFederatedRateLimitConfig(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchFederatedRateLimitConfig indicates an expected call of PatchFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigClientMockRecorder) PatchFederatedRateLimitConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).PatchFederatedRateLimitConfig), varargs...)
}

// DeleteAllOfFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigClient) DeleteAllOfFederatedRateLimitConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfFederatedRateLimitConfig indicates an expected call of DeleteAllOfFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigClientMockRecorder) DeleteAllOfFederatedRateLimitConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).DeleteAllOfFederatedRateLimitConfig), varargs...)
}

// UpsertFederatedRateLimitConfig mocks base method
func (m *MockFederatedRateLimitConfigClient) UpsertFederatedRateLimitConfig(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, transitionFuncs ...v1alpha1.FederatedRateLimitConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertFederatedRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertFederatedRateLimitConfig indicates an expected call of UpsertFederatedRateLimitConfig
func (mr *MockFederatedRateLimitConfigClientMockRecorder) UpsertFederatedRateLimitConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertFederatedRateLimitConfig", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).UpsertFederatedRateLimitConfig), varargs...)
}

// UpdateFederatedRateLimitConfigStatus mocks base method
func (m *MockFederatedRateLimitConfigClient) UpdateFederatedRateLimitConfigStatus(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedRateLimitConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedRateLimitConfigStatus indicates an expected call of UpdateFederatedRateLimitConfigStatus
func (mr *MockFederatedRateLimitConfigClientMockRecorder) UpdateFederatedRateLimitConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedRateLimitConfigStatus", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).UpdateFederatedRateLimitConfigStatus), varargs...)
}

// PatchFederatedRateLimitConfigStatus mocks base method
func (m *MockFederatedRateLimitConfigClient) PatchFederatedRateLimitConfigStatus(ctx context.Context, obj *v1alpha1.FederatedRateLimitConfig, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchFederatedRateLimitConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchFederatedRateLimitConfigStatus indicates an expected call of PatchFederatedRateLimitConfigStatus
func (mr *MockFederatedRateLimitConfigClientMockRecorder) PatchFederatedRateLimitConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchFederatedRateLimitConfigStatus", reflect.TypeOf((*MockFederatedRateLimitConfigClient)(nil).PatchFederatedRateLimitConfigStatus), varargs...)
}

// MockMulticlusterFederatedRateLimitConfigClient is a mock of MulticlusterFederatedRateLimitConfigClient interface
type MockMulticlusterFederatedRateLimitConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedRateLimitConfigClientMockRecorder
}

// MockMulticlusterFederatedRateLimitConfigClientMockRecorder is the mock recorder for MockMulticlusterFederatedRateLimitConfigClient
type MockMulticlusterFederatedRateLimitConfigClientMockRecorder struct {
	mock *MockMulticlusterFederatedRateLimitConfigClient
}

// NewMockMulticlusterFederatedRateLimitConfigClient creates a new mock instance
func NewMockMulticlusterFederatedRateLimitConfigClient(ctrl *gomock.Controller) *MockMulticlusterFederatedRateLimitConfigClient {
	mock := &MockMulticlusterFederatedRateLimitConfigClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedRateLimitConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFederatedRateLimitConfigClient) EXPECT() *MockMulticlusterFederatedRateLimitConfigClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterFederatedRateLimitConfigClient) Cluster(cluster string) (v1alpha1.FederatedRateLimitConfigClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha1.FederatedRateLimitConfigClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterFederatedRateLimitConfigClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterFederatedRateLimitConfigClient)(nil).Cluster), cluster)
}
