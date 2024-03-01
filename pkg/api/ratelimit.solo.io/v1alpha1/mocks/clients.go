// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1alpha1 is a generated GoMock package.
package mock_v1alpha1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
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

// RateLimitConfigs mocks base method
func (m *MockClientset) RateLimitConfigs() v1alpha1.RateLimitConfigClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RateLimitConfigs")
	ret0, _ := ret[0].(v1alpha1.RateLimitConfigClient)
	return ret0
}

// RateLimitConfigs indicates an expected call of RateLimitConfigs
func (mr *MockClientsetMockRecorder) RateLimitConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RateLimitConfigs", reflect.TypeOf((*MockClientset)(nil).RateLimitConfigs))
}

// MockRateLimitConfigReader is a mock of RateLimitConfigReader interface
type MockRateLimitConfigReader struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimitConfigReaderMockRecorder
}

// MockRateLimitConfigReaderMockRecorder is the mock recorder for MockRateLimitConfigReader
type MockRateLimitConfigReaderMockRecorder struct {
	mock *MockRateLimitConfigReader
}

// NewMockRateLimitConfigReader creates a new mock instance
func NewMockRateLimitConfigReader(ctrl *gomock.Controller) *MockRateLimitConfigReader {
	mock := &MockRateLimitConfigReader{ctrl: ctrl}
	mock.recorder = &MockRateLimitConfigReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRateLimitConfigReader) EXPECT() *MockRateLimitConfigReaderMockRecorder {
	return m.recorder
}

// GetRateLimitConfig mocks base method
func (m *MockRateLimitConfigReader) GetRateLimitConfig(ctx context.Context, key client.ObjectKey) (*v1alpha1.RateLimitConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRateLimitConfig", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.RateLimitConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateLimitConfig indicates an expected call of GetRateLimitConfig
func (mr *MockRateLimitConfigReaderMockRecorder) GetRateLimitConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigReader)(nil).GetRateLimitConfig), ctx, key)
}

// ListRateLimitConfig mocks base method
func (m *MockRateLimitConfigReader) ListRateLimitConfig(ctx context.Context, opts ...client.ListOption) (*v1alpha1.RateLimitConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRateLimitConfig", varargs...)
	ret0, _ := ret[0].(*v1alpha1.RateLimitConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRateLimitConfig indicates an expected call of ListRateLimitConfig
func (mr *MockRateLimitConfigReaderMockRecorder) ListRateLimitConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigReader)(nil).ListRateLimitConfig), varargs...)
}

// MockRateLimitConfigWriter is a mock of RateLimitConfigWriter interface
type MockRateLimitConfigWriter struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimitConfigWriterMockRecorder
}

// MockRateLimitConfigWriterMockRecorder is the mock recorder for MockRateLimitConfigWriter
type MockRateLimitConfigWriterMockRecorder struct {
	mock *MockRateLimitConfigWriter
}

// NewMockRateLimitConfigWriter creates a new mock instance
func NewMockRateLimitConfigWriter(ctrl *gomock.Controller) *MockRateLimitConfigWriter {
	mock := &MockRateLimitConfigWriter{ctrl: ctrl}
	mock.recorder = &MockRateLimitConfigWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRateLimitConfigWriter) EXPECT() *MockRateLimitConfigWriterMockRecorder {
	return m.recorder
}

// CreateRateLimitConfig mocks base method
func (m *MockRateLimitConfigWriter) CreateRateLimitConfig(ctx context.Context, obj *v1alpha1.RateLimitConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRateLimitConfig indicates an expected call of CreateRateLimitConfig
func (mr *MockRateLimitConfigWriterMockRecorder) CreateRateLimitConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigWriter)(nil).CreateRateLimitConfig), varargs...)
}

// DeleteRateLimitConfig mocks base method
func (m *MockRateLimitConfigWriter) DeleteRateLimitConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRateLimitConfig indicates an expected call of DeleteRateLimitConfig
func (mr *MockRateLimitConfigWriterMockRecorder) DeleteRateLimitConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigWriter)(nil).DeleteRateLimitConfig), varargs...)
}

// UpdateRateLimitConfig mocks base method
func (m *MockRateLimitConfigWriter) UpdateRateLimitConfig(ctx context.Context, obj *v1alpha1.RateLimitConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRateLimitConfig indicates an expected call of UpdateRateLimitConfig
func (mr *MockRateLimitConfigWriterMockRecorder) UpdateRateLimitConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigWriter)(nil).UpdateRateLimitConfig), varargs...)
}

// PatchRateLimitConfig mocks base method
func (m *MockRateLimitConfigWriter) PatchRateLimitConfig(ctx context.Context, obj *v1alpha1.RateLimitConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchRateLimitConfig indicates an expected call of PatchRateLimitConfig
func (mr *MockRateLimitConfigWriterMockRecorder) PatchRateLimitConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigWriter)(nil).PatchRateLimitConfig), varargs...)
}

// DeleteAllOfRateLimitConfig mocks base method
func (m *MockRateLimitConfigWriter) DeleteAllOfRateLimitConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfRateLimitConfig indicates an expected call of DeleteAllOfRateLimitConfig
func (mr *MockRateLimitConfigWriterMockRecorder) DeleteAllOfRateLimitConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigWriter)(nil).DeleteAllOfRateLimitConfig), varargs...)
}

// UpsertRateLimitConfig mocks base method
func (m *MockRateLimitConfigWriter) UpsertRateLimitConfig(ctx context.Context, obj *v1alpha1.RateLimitConfig, transitionFuncs ...v1alpha1.RateLimitConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRateLimitConfig indicates an expected call of UpsertRateLimitConfig
func (mr *MockRateLimitConfigWriterMockRecorder) UpsertRateLimitConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigWriter)(nil).UpsertRateLimitConfig), varargs...)
}

// MockRateLimitConfigStatusWriter is a mock of RateLimitConfigStatusWriter interface
type MockRateLimitConfigStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimitConfigStatusWriterMockRecorder
}

// MockRateLimitConfigStatusWriterMockRecorder is the mock recorder for MockRateLimitConfigStatusWriter
type MockRateLimitConfigStatusWriterMockRecorder struct {
	mock *MockRateLimitConfigStatusWriter
}

// NewMockRateLimitConfigStatusWriter creates a new mock instance
func NewMockRateLimitConfigStatusWriter(ctrl *gomock.Controller) *MockRateLimitConfigStatusWriter {
	mock := &MockRateLimitConfigStatusWriter{ctrl: ctrl}
	mock.recorder = &MockRateLimitConfigStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRateLimitConfigStatusWriter) EXPECT() *MockRateLimitConfigStatusWriterMockRecorder {
	return m.recorder
}

// UpdateRateLimitConfigStatus mocks base method
func (m *MockRateLimitConfigStatusWriter) UpdateRateLimitConfigStatus(ctx context.Context, obj *v1alpha1.RateLimitConfig, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRateLimitConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRateLimitConfigStatus indicates an expected call of UpdateRateLimitConfigStatus
func (mr *MockRateLimitConfigStatusWriterMockRecorder) UpdateRateLimitConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRateLimitConfigStatus", reflect.TypeOf((*MockRateLimitConfigStatusWriter)(nil).UpdateRateLimitConfigStatus), varargs...)
}

// PatchRateLimitConfigStatus mocks base method
func (m *MockRateLimitConfigStatusWriter) PatchRateLimitConfigStatus(ctx context.Context, obj *v1alpha1.RateLimitConfig, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchRateLimitConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchRateLimitConfigStatus indicates an expected call of PatchRateLimitConfigStatus
func (mr *MockRateLimitConfigStatusWriterMockRecorder) PatchRateLimitConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchRateLimitConfigStatus", reflect.TypeOf((*MockRateLimitConfigStatusWriter)(nil).PatchRateLimitConfigStatus), varargs...)
}

// MockRateLimitConfigClient is a mock of RateLimitConfigClient interface
type MockRateLimitConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimitConfigClientMockRecorder
}

// MockRateLimitConfigClientMockRecorder is the mock recorder for MockRateLimitConfigClient
type MockRateLimitConfigClientMockRecorder struct {
	mock *MockRateLimitConfigClient
}

// NewMockRateLimitConfigClient creates a new mock instance
func NewMockRateLimitConfigClient(ctrl *gomock.Controller) *MockRateLimitConfigClient {
	mock := &MockRateLimitConfigClient{ctrl: ctrl}
	mock.recorder = &MockRateLimitConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRateLimitConfigClient) EXPECT() *MockRateLimitConfigClientMockRecorder {
	return m.recorder
}

// GetRateLimitConfig mocks base method
func (m *MockRateLimitConfigClient) GetRateLimitConfig(ctx context.Context, key client.ObjectKey) (*v1alpha1.RateLimitConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRateLimitConfig", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.RateLimitConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateLimitConfig indicates an expected call of GetRateLimitConfig
func (mr *MockRateLimitConfigClientMockRecorder) GetRateLimitConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigClient)(nil).GetRateLimitConfig), ctx, key)
}

// ListRateLimitConfig mocks base method
func (m *MockRateLimitConfigClient) ListRateLimitConfig(ctx context.Context, opts ...client.ListOption) (*v1alpha1.RateLimitConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRateLimitConfig", varargs...)
	ret0, _ := ret[0].(*v1alpha1.RateLimitConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRateLimitConfig indicates an expected call of ListRateLimitConfig
func (mr *MockRateLimitConfigClientMockRecorder) ListRateLimitConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigClient)(nil).ListRateLimitConfig), varargs...)
}

// CreateRateLimitConfig mocks base method
func (m *MockRateLimitConfigClient) CreateRateLimitConfig(ctx context.Context, obj *v1alpha1.RateLimitConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRateLimitConfig indicates an expected call of CreateRateLimitConfig
func (mr *MockRateLimitConfigClientMockRecorder) CreateRateLimitConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigClient)(nil).CreateRateLimitConfig), varargs...)
}

// DeleteRateLimitConfig mocks base method
func (m *MockRateLimitConfigClient) DeleteRateLimitConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRateLimitConfig indicates an expected call of DeleteRateLimitConfig
func (mr *MockRateLimitConfigClientMockRecorder) DeleteRateLimitConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigClient)(nil).DeleteRateLimitConfig), varargs...)
}

// UpdateRateLimitConfig mocks base method
func (m *MockRateLimitConfigClient) UpdateRateLimitConfig(ctx context.Context, obj *v1alpha1.RateLimitConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRateLimitConfig indicates an expected call of UpdateRateLimitConfig
func (mr *MockRateLimitConfigClientMockRecorder) UpdateRateLimitConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigClient)(nil).UpdateRateLimitConfig), varargs...)
}

// PatchRateLimitConfig mocks base method
func (m *MockRateLimitConfigClient) PatchRateLimitConfig(ctx context.Context, obj *v1alpha1.RateLimitConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchRateLimitConfig indicates an expected call of PatchRateLimitConfig
func (mr *MockRateLimitConfigClientMockRecorder) PatchRateLimitConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigClient)(nil).PatchRateLimitConfig), varargs...)
}

// DeleteAllOfRateLimitConfig mocks base method
func (m *MockRateLimitConfigClient) DeleteAllOfRateLimitConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfRateLimitConfig indicates an expected call of DeleteAllOfRateLimitConfig
func (mr *MockRateLimitConfigClientMockRecorder) DeleteAllOfRateLimitConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigClient)(nil).DeleteAllOfRateLimitConfig), varargs...)
}

// UpsertRateLimitConfig mocks base method
func (m *MockRateLimitConfigClient) UpsertRateLimitConfig(ctx context.Context, obj *v1alpha1.RateLimitConfig, transitionFuncs ...v1alpha1.RateLimitConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertRateLimitConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRateLimitConfig indicates an expected call of UpsertRateLimitConfig
func (mr *MockRateLimitConfigClientMockRecorder) UpsertRateLimitConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRateLimitConfig", reflect.TypeOf((*MockRateLimitConfigClient)(nil).UpsertRateLimitConfig), varargs...)
}

// UpdateRateLimitConfigStatus mocks base method
func (m *MockRateLimitConfigClient) UpdateRateLimitConfigStatus(ctx context.Context, obj *v1alpha1.RateLimitConfig, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRateLimitConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRateLimitConfigStatus indicates an expected call of UpdateRateLimitConfigStatus
func (mr *MockRateLimitConfigClientMockRecorder) UpdateRateLimitConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRateLimitConfigStatus", reflect.TypeOf((*MockRateLimitConfigClient)(nil).UpdateRateLimitConfigStatus), varargs...)
}

// PatchRateLimitConfigStatus mocks base method
func (m *MockRateLimitConfigClient) PatchRateLimitConfigStatus(ctx context.Context, obj *v1alpha1.RateLimitConfig, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchRateLimitConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchRateLimitConfigStatus indicates an expected call of PatchRateLimitConfigStatus
func (mr *MockRateLimitConfigClientMockRecorder) PatchRateLimitConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchRateLimitConfigStatus", reflect.TypeOf((*MockRateLimitConfigClient)(nil).PatchRateLimitConfigStatus), varargs...)
}

// MockMulticlusterRateLimitConfigClient is a mock of MulticlusterRateLimitConfigClient interface
type MockMulticlusterRateLimitConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRateLimitConfigClientMockRecorder
}

// MockMulticlusterRateLimitConfigClientMockRecorder is the mock recorder for MockMulticlusterRateLimitConfigClient
type MockMulticlusterRateLimitConfigClientMockRecorder struct {
	mock *MockMulticlusterRateLimitConfigClient
}

// NewMockMulticlusterRateLimitConfigClient creates a new mock instance
func NewMockMulticlusterRateLimitConfigClient(ctrl *gomock.Controller) *MockMulticlusterRateLimitConfigClient {
	mock := &MockMulticlusterRateLimitConfigClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRateLimitConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRateLimitConfigClient) EXPECT() *MockMulticlusterRateLimitConfigClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterRateLimitConfigClient) Cluster(cluster string) (v1alpha1.RateLimitConfigClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha1.RateLimitConfigClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterRateLimitConfigClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterRateLimitConfigClient)(nil).Cluster), cluster)
}
