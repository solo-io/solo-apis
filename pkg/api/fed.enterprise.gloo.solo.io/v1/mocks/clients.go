// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.enterprise.gloo.solo.io/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
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
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1.Clientset)
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

// FederatedAuthConfigs mocks base method.
func (m *MockClientset) FederatedAuthConfigs() v1.FederatedAuthConfigClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederatedAuthConfigs")
	ret0, _ := ret[0].(v1.FederatedAuthConfigClient)
	return ret0
}

// FederatedAuthConfigs indicates an expected call of FederatedAuthConfigs.
func (mr *MockClientsetMockRecorder) FederatedAuthConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederatedAuthConfigs", reflect.TypeOf((*MockClientset)(nil).FederatedAuthConfigs))
}

// MockFederatedAuthConfigReader is a mock of FederatedAuthConfigReader interface.
type MockFederatedAuthConfigReader struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigReaderMockRecorder
}

// MockFederatedAuthConfigReaderMockRecorder is the mock recorder for MockFederatedAuthConfigReader.
type MockFederatedAuthConfigReaderMockRecorder struct {
	mock *MockFederatedAuthConfigReader
}

// NewMockFederatedAuthConfigReader creates a new mock instance.
func NewMockFederatedAuthConfigReader(ctrl *gomock.Controller) *MockFederatedAuthConfigReader {
	mock := &MockFederatedAuthConfigReader{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedAuthConfigReader) EXPECT() *MockFederatedAuthConfigReaderMockRecorder {
	return m.recorder
}

// GetFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigReader) GetFederatedAuthConfig(ctx context.Context, key client.ObjectKey) (*v1.FederatedAuthConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederatedAuthConfig", ctx, key)
	ret0, _ := ret[0].(*v1.FederatedAuthConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederatedAuthConfig indicates an expected call of GetFederatedAuthConfig.
func (mr *MockFederatedAuthConfigReaderMockRecorder) GetFederatedAuthConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigReader)(nil).GetFederatedAuthConfig), ctx, key)
}

// ListFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigReader) ListFederatedAuthConfig(ctx context.Context, opts ...client.ListOption) (*v1.FederatedAuthConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(*v1.FederatedAuthConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFederatedAuthConfig indicates an expected call of ListFederatedAuthConfig.
func (mr *MockFederatedAuthConfigReaderMockRecorder) ListFederatedAuthConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigReader)(nil).ListFederatedAuthConfig), varargs...)
}

// MockFederatedAuthConfigWriter is a mock of FederatedAuthConfigWriter interface.
type MockFederatedAuthConfigWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigWriterMockRecorder
}

// MockFederatedAuthConfigWriterMockRecorder is the mock recorder for MockFederatedAuthConfigWriter.
type MockFederatedAuthConfigWriterMockRecorder struct {
	mock *MockFederatedAuthConfigWriter
}

// NewMockFederatedAuthConfigWriter creates a new mock instance.
func NewMockFederatedAuthConfigWriter(ctrl *gomock.Controller) *MockFederatedAuthConfigWriter {
	mock := &MockFederatedAuthConfigWriter{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedAuthConfigWriter) EXPECT() *MockFederatedAuthConfigWriterMockRecorder {
	return m.recorder
}

// CreateFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigWriter) CreateFederatedAuthConfig(ctx context.Context, obj *v1.FederatedAuthConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedAuthConfig indicates an expected call of CreateFederatedAuthConfig.
func (mr *MockFederatedAuthConfigWriterMockRecorder) CreateFederatedAuthConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigWriter)(nil).CreateFederatedAuthConfig), varargs...)
}

// DeleteAllOfFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigWriter) DeleteAllOfFederatedAuthConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfFederatedAuthConfig indicates an expected call of DeleteAllOfFederatedAuthConfig.
func (mr *MockFederatedAuthConfigWriterMockRecorder) DeleteAllOfFederatedAuthConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigWriter)(nil).DeleteAllOfFederatedAuthConfig), varargs...)
}

// DeleteFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigWriter) DeleteFederatedAuthConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedAuthConfig indicates an expected call of DeleteFederatedAuthConfig.
func (mr *MockFederatedAuthConfigWriterMockRecorder) DeleteFederatedAuthConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigWriter)(nil).DeleteFederatedAuthConfig), varargs...)
}

// PatchFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigWriter) PatchFederatedAuthConfig(ctx context.Context, obj *v1.FederatedAuthConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchFederatedAuthConfig indicates an expected call of PatchFederatedAuthConfig.
func (mr *MockFederatedAuthConfigWriterMockRecorder) PatchFederatedAuthConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigWriter)(nil).PatchFederatedAuthConfig), varargs...)
}

// UpdateFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigWriter) UpdateFederatedAuthConfig(ctx context.Context, obj *v1.FederatedAuthConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedAuthConfig indicates an expected call of UpdateFederatedAuthConfig.
func (mr *MockFederatedAuthConfigWriterMockRecorder) UpdateFederatedAuthConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigWriter)(nil).UpdateFederatedAuthConfig), varargs...)
}

// UpsertFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigWriter) UpsertFederatedAuthConfig(ctx context.Context, obj *v1.FederatedAuthConfig, transitionFuncs ...v1.FederatedAuthConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertFederatedAuthConfig indicates an expected call of UpsertFederatedAuthConfig.
func (mr *MockFederatedAuthConfigWriterMockRecorder) UpsertFederatedAuthConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigWriter)(nil).UpsertFederatedAuthConfig), varargs...)
}

// MockFederatedAuthConfigStatusWriter is a mock of FederatedAuthConfigStatusWriter interface.
type MockFederatedAuthConfigStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigStatusWriterMockRecorder
}

// MockFederatedAuthConfigStatusWriterMockRecorder is the mock recorder for MockFederatedAuthConfigStatusWriter.
type MockFederatedAuthConfigStatusWriterMockRecorder struct {
	mock *MockFederatedAuthConfigStatusWriter
}

// NewMockFederatedAuthConfigStatusWriter creates a new mock instance.
func NewMockFederatedAuthConfigStatusWriter(ctrl *gomock.Controller) *MockFederatedAuthConfigStatusWriter {
	mock := &MockFederatedAuthConfigStatusWriter{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedAuthConfigStatusWriter) EXPECT() *MockFederatedAuthConfigStatusWriterMockRecorder {
	return m.recorder
}

// PatchFederatedAuthConfigStatus mocks base method.
func (m *MockFederatedAuthConfigStatusWriter) PatchFederatedAuthConfigStatus(ctx context.Context, obj *v1.FederatedAuthConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchFederatedAuthConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchFederatedAuthConfigStatus indicates an expected call of PatchFederatedAuthConfigStatus.
func (mr *MockFederatedAuthConfigStatusWriterMockRecorder) PatchFederatedAuthConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchFederatedAuthConfigStatus", reflect.TypeOf((*MockFederatedAuthConfigStatusWriter)(nil).PatchFederatedAuthConfigStatus), varargs...)
}

// UpdateFederatedAuthConfigStatus mocks base method.
func (m *MockFederatedAuthConfigStatusWriter) UpdateFederatedAuthConfigStatus(ctx context.Context, obj *v1.FederatedAuthConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedAuthConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedAuthConfigStatus indicates an expected call of UpdateFederatedAuthConfigStatus.
func (mr *MockFederatedAuthConfigStatusWriterMockRecorder) UpdateFederatedAuthConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedAuthConfigStatus", reflect.TypeOf((*MockFederatedAuthConfigStatusWriter)(nil).UpdateFederatedAuthConfigStatus), varargs...)
}

// MockFederatedAuthConfigClient is a mock of FederatedAuthConfigClient interface.
type MockFederatedAuthConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedAuthConfigClientMockRecorder
}

// MockFederatedAuthConfigClientMockRecorder is the mock recorder for MockFederatedAuthConfigClient.
type MockFederatedAuthConfigClientMockRecorder struct {
	mock *MockFederatedAuthConfigClient
}

// NewMockFederatedAuthConfigClient creates a new mock instance.
func NewMockFederatedAuthConfigClient(ctrl *gomock.Controller) *MockFederatedAuthConfigClient {
	mock := &MockFederatedAuthConfigClient{ctrl: ctrl}
	mock.recorder = &MockFederatedAuthConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedAuthConfigClient) EXPECT() *MockFederatedAuthConfigClientMockRecorder {
	return m.recorder
}

// CreateFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigClient) CreateFederatedAuthConfig(ctx context.Context, obj *v1.FederatedAuthConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedAuthConfig indicates an expected call of CreateFederatedAuthConfig.
func (mr *MockFederatedAuthConfigClientMockRecorder) CreateFederatedAuthConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).CreateFederatedAuthConfig), varargs...)
}

// DeleteAllOfFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigClient) DeleteAllOfFederatedAuthConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfFederatedAuthConfig indicates an expected call of DeleteAllOfFederatedAuthConfig.
func (mr *MockFederatedAuthConfigClientMockRecorder) DeleteAllOfFederatedAuthConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).DeleteAllOfFederatedAuthConfig), varargs...)
}

// DeleteFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigClient) DeleteFederatedAuthConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedAuthConfig indicates an expected call of DeleteFederatedAuthConfig.
func (mr *MockFederatedAuthConfigClientMockRecorder) DeleteFederatedAuthConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).DeleteFederatedAuthConfig), varargs...)
}

// GetFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigClient) GetFederatedAuthConfig(ctx context.Context, key client.ObjectKey) (*v1.FederatedAuthConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederatedAuthConfig", ctx, key)
	ret0, _ := ret[0].(*v1.FederatedAuthConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederatedAuthConfig indicates an expected call of GetFederatedAuthConfig.
func (mr *MockFederatedAuthConfigClientMockRecorder) GetFederatedAuthConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).GetFederatedAuthConfig), ctx, key)
}

// ListFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigClient) ListFederatedAuthConfig(ctx context.Context, opts ...client.ListOption) (*v1.FederatedAuthConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(*v1.FederatedAuthConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFederatedAuthConfig indicates an expected call of ListFederatedAuthConfig.
func (mr *MockFederatedAuthConfigClientMockRecorder) ListFederatedAuthConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).ListFederatedAuthConfig), varargs...)
}

// PatchFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigClient) PatchFederatedAuthConfig(ctx context.Context, obj *v1.FederatedAuthConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchFederatedAuthConfig indicates an expected call of PatchFederatedAuthConfig.
func (mr *MockFederatedAuthConfigClientMockRecorder) PatchFederatedAuthConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).PatchFederatedAuthConfig), varargs...)
}

// PatchFederatedAuthConfigStatus mocks base method.
func (m *MockFederatedAuthConfigClient) PatchFederatedAuthConfigStatus(ctx context.Context, obj *v1.FederatedAuthConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchFederatedAuthConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchFederatedAuthConfigStatus indicates an expected call of PatchFederatedAuthConfigStatus.
func (mr *MockFederatedAuthConfigClientMockRecorder) PatchFederatedAuthConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchFederatedAuthConfigStatus", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).PatchFederatedAuthConfigStatus), varargs...)
}

// UpdateFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigClient) UpdateFederatedAuthConfig(ctx context.Context, obj *v1.FederatedAuthConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedAuthConfig indicates an expected call of UpdateFederatedAuthConfig.
func (mr *MockFederatedAuthConfigClientMockRecorder) UpdateFederatedAuthConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).UpdateFederatedAuthConfig), varargs...)
}

// UpdateFederatedAuthConfigStatus mocks base method.
func (m *MockFederatedAuthConfigClient) UpdateFederatedAuthConfigStatus(ctx context.Context, obj *v1.FederatedAuthConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedAuthConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedAuthConfigStatus indicates an expected call of UpdateFederatedAuthConfigStatus.
func (mr *MockFederatedAuthConfigClientMockRecorder) UpdateFederatedAuthConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedAuthConfigStatus", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).UpdateFederatedAuthConfigStatus), varargs...)
}

// UpsertFederatedAuthConfig mocks base method.
func (m *MockFederatedAuthConfigClient) UpsertFederatedAuthConfig(ctx context.Context, obj *v1.FederatedAuthConfig, transitionFuncs ...v1.FederatedAuthConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertFederatedAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertFederatedAuthConfig indicates an expected call of UpsertFederatedAuthConfig.
func (mr *MockFederatedAuthConfigClientMockRecorder) UpsertFederatedAuthConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertFederatedAuthConfig", reflect.TypeOf((*MockFederatedAuthConfigClient)(nil).UpsertFederatedAuthConfig), varargs...)
}

// MockMulticlusterFederatedAuthConfigClient is a mock of MulticlusterFederatedAuthConfigClient interface.
type MockMulticlusterFederatedAuthConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedAuthConfigClientMockRecorder
}

// MockMulticlusterFederatedAuthConfigClientMockRecorder is the mock recorder for MockMulticlusterFederatedAuthConfigClient.
type MockMulticlusterFederatedAuthConfigClientMockRecorder struct {
	mock *MockMulticlusterFederatedAuthConfigClient
}

// NewMockMulticlusterFederatedAuthConfigClient creates a new mock instance.
func NewMockMulticlusterFederatedAuthConfigClient(ctrl *gomock.Controller) *MockMulticlusterFederatedAuthConfigClient {
	mock := &MockMulticlusterFederatedAuthConfigClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedAuthConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFederatedAuthConfigClient) EXPECT() *MockMulticlusterFederatedAuthConfigClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterFederatedAuthConfigClient) Cluster(cluster string) (v1.FederatedAuthConfigClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1.FederatedAuthConfigClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterFederatedAuthConfigClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterFederatedAuthConfigClient)(nil).Cluster), cluster)
}
