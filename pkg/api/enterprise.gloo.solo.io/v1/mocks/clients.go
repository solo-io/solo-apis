// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
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

// AuthConfigs mocks base method.
func (m *MockClientset) AuthConfigs() v1.AuthConfigClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthConfigs")
	ret0, _ := ret[0].(v1.AuthConfigClient)
	return ret0
}

// AuthConfigs indicates an expected call of AuthConfigs.
func (mr *MockClientsetMockRecorder) AuthConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthConfigs", reflect.TypeOf((*MockClientset)(nil).AuthConfigs))
}

// MockAuthConfigReader is a mock of AuthConfigReader interface.
type MockAuthConfigReader struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigReaderMockRecorder
}

// MockAuthConfigReaderMockRecorder is the mock recorder for MockAuthConfigReader.
type MockAuthConfigReaderMockRecorder struct {
	mock *MockAuthConfigReader
}

// NewMockAuthConfigReader creates a new mock instance.
func NewMockAuthConfigReader(ctrl *gomock.Controller) *MockAuthConfigReader {
	mock := &MockAuthConfigReader{ctrl: ctrl}
	mock.recorder = &MockAuthConfigReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthConfigReader) EXPECT() *MockAuthConfigReaderMockRecorder {
	return m.recorder
}

// GetAuthConfig mocks base method.
func (m *MockAuthConfigReader) GetAuthConfig(ctx context.Context, key client.ObjectKey) (*v1.AuthConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthConfig", ctx, key)
	ret0, _ := ret[0].(*v1.AuthConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthConfig indicates an expected call of GetAuthConfig.
func (mr *MockAuthConfigReaderMockRecorder) GetAuthConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthConfig", reflect.TypeOf((*MockAuthConfigReader)(nil).GetAuthConfig), ctx, key)
}

// ListAuthConfig mocks base method.
func (m *MockAuthConfigReader) ListAuthConfig(ctx context.Context, opts ...client.ListOption) (*v1.AuthConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAuthConfig", varargs...)
	ret0, _ := ret[0].(*v1.AuthConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthConfig indicates an expected call of ListAuthConfig.
func (mr *MockAuthConfigReaderMockRecorder) ListAuthConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthConfig", reflect.TypeOf((*MockAuthConfigReader)(nil).ListAuthConfig), varargs...)
}

// MockAuthConfigWriter is a mock of AuthConfigWriter interface.
type MockAuthConfigWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigWriterMockRecorder
}

// MockAuthConfigWriterMockRecorder is the mock recorder for MockAuthConfigWriter.
type MockAuthConfigWriterMockRecorder struct {
	mock *MockAuthConfigWriter
}

// NewMockAuthConfigWriter creates a new mock instance.
func NewMockAuthConfigWriter(ctrl *gomock.Controller) *MockAuthConfigWriter {
	mock := &MockAuthConfigWriter{ctrl: ctrl}
	mock.recorder = &MockAuthConfigWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthConfigWriter) EXPECT() *MockAuthConfigWriterMockRecorder {
	return m.recorder
}

// CreateAuthConfig mocks base method.
func (m *MockAuthConfigWriter) CreateAuthConfig(ctx context.Context, obj *v1.AuthConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAuthConfig indicates an expected call of CreateAuthConfig.
func (mr *MockAuthConfigWriterMockRecorder) CreateAuthConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthConfig", reflect.TypeOf((*MockAuthConfigWriter)(nil).CreateAuthConfig), varargs...)
}

// DeleteAllOfAuthConfig mocks base method.
func (m *MockAuthConfigWriter) DeleteAllOfAuthConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfAuthConfig indicates an expected call of DeleteAllOfAuthConfig.
func (mr *MockAuthConfigWriterMockRecorder) DeleteAllOfAuthConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfAuthConfig", reflect.TypeOf((*MockAuthConfigWriter)(nil).DeleteAllOfAuthConfig), varargs...)
}

// DeleteAuthConfig mocks base method.
func (m *MockAuthConfigWriter) DeleteAuthConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthConfig indicates an expected call of DeleteAuthConfig.
func (mr *MockAuthConfigWriterMockRecorder) DeleteAuthConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthConfig", reflect.TypeOf((*MockAuthConfigWriter)(nil).DeleteAuthConfig), varargs...)
}

// PatchAuthConfig mocks base method.
func (m *MockAuthConfigWriter) PatchAuthConfig(ctx context.Context, obj *v1.AuthConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchAuthConfig indicates an expected call of PatchAuthConfig.
func (mr *MockAuthConfigWriterMockRecorder) PatchAuthConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchAuthConfig", reflect.TypeOf((*MockAuthConfigWriter)(nil).PatchAuthConfig), varargs...)
}

// UpdateAuthConfig mocks base method.
func (m *MockAuthConfigWriter) UpdateAuthConfig(ctx context.Context, obj *v1.AuthConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthConfig indicates an expected call of UpdateAuthConfig.
func (mr *MockAuthConfigWriterMockRecorder) UpdateAuthConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthConfig", reflect.TypeOf((*MockAuthConfigWriter)(nil).UpdateAuthConfig), varargs...)
}

// UpsertAuthConfig mocks base method.
func (m *MockAuthConfigWriter) UpsertAuthConfig(ctx context.Context, obj *v1.AuthConfig, transitionFuncs ...v1.AuthConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertAuthConfig indicates an expected call of UpsertAuthConfig.
func (mr *MockAuthConfigWriterMockRecorder) UpsertAuthConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertAuthConfig", reflect.TypeOf((*MockAuthConfigWriter)(nil).UpsertAuthConfig), varargs...)
}

// MockAuthConfigStatusWriter is a mock of AuthConfigStatusWriter interface.
type MockAuthConfigStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigStatusWriterMockRecorder
}

// MockAuthConfigStatusWriterMockRecorder is the mock recorder for MockAuthConfigStatusWriter.
type MockAuthConfigStatusWriterMockRecorder struct {
	mock *MockAuthConfigStatusWriter
}

// NewMockAuthConfigStatusWriter creates a new mock instance.
func NewMockAuthConfigStatusWriter(ctrl *gomock.Controller) *MockAuthConfigStatusWriter {
	mock := &MockAuthConfigStatusWriter{ctrl: ctrl}
	mock.recorder = &MockAuthConfigStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthConfigStatusWriter) EXPECT() *MockAuthConfigStatusWriterMockRecorder {
	return m.recorder
}

// PatchAuthConfigStatus mocks base method.
func (m *MockAuthConfigStatusWriter) PatchAuthConfigStatus(ctx context.Context, obj *v1.AuthConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchAuthConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchAuthConfigStatus indicates an expected call of PatchAuthConfigStatus.
func (mr *MockAuthConfigStatusWriterMockRecorder) PatchAuthConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchAuthConfigStatus", reflect.TypeOf((*MockAuthConfigStatusWriter)(nil).PatchAuthConfigStatus), varargs...)
}

// UpdateAuthConfigStatus mocks base method.
func (m *MockAuthConfigStatusWriter) UpdateAuthConfigStatus(ctx context.Context, obj *v1.AuthConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAuthConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthConfigStatus indicates an expected call of UpdateAuthConfigStatus.
func (mr *MockAuthConfigStatusWriterMockRecorder) UpdateAuthConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthConfigStatus", reflect.TypeOf((*MockAuthConfigStatusWriter)(nil).UpdateAuthConfigStatus), varargs...)
}

// MockAuthConfigClient is a mock of AuthConfigClient interface.
type MockAuthConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigClientMockRecorder
}

// MockAuthConfigClientMockRecorder is the mock recorder for MockAuthConfigClient.
type MockAuthConfigClientMockRecorder struct {
	mock *MockAuthConfigClient
}

// NewMockAuthConfigClient creates a new mock instance.
func NewMockAuthConfigClient(ctrl *gomock.Controller) *MockAuthConfigClient {
	mock := &MockAuthConfigClient{ctrl: ctrl}
	mock.recorder = &MockAuthConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthConfigClient) EXPECT() *MockAuthConfigClientMockRecorder {
	return m.recorder
}

// CreateAuthConfig mocks base method.
func (m *MockAuthConfigClient) CreateAuthConfig(ctx context.Context, obj *v1.AuthConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAuthConfig indicates an expected call of CreateAuthConfig.
func (mr *MockAuthConfigClientMockRecorder) CreateAuthConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthConfig", reflect.TypeOf((*MockAuthConfigClient)(nil).CreateAuthConfig), varargs...)
}

// DeleteAllOfAuthConfig mocks base method.
func (m *MockAuthConfigClient) DeleteAllOfAuthConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfAuthConfig indicates an expected call of DeleteAllOfAuthConfig.
func (mr *MockAuthConfigClientMockRecorder) DeleteAllOfAuthConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfAuthConfig", reflect.TypeOf((*MockAuthConfigClient)(nil).DeleteAllOfAuthConfig), varargs...)
}

// DeleteAuthConfig mocks base method.
func (m *MockAuthConfigClient) DeleteAuthConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthConfig indicates an expected call of DeleteAuthConfig.
func (mr *MockAuthConfigClientMockRecorder) DeleteAuthConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthConfig", reflect.TypeOf((*MockAuthConfigClient)(nil).DeleteAuthConfig), varargs...)
}

// GetAuthConfig mocks base method.
func (m *MockAuthConfigClient) GetAuthConfig(ctx context.Context, key client.ObjectKey) (*v1.AuthConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthConfig", ctx, key)
	ret0, _ := ret[0].(*v1.AuthConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthConfig indicates an expected call of GetAuthConfig.
func (mr *MockAuthConfigClientMockRecorder) GetAuthConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthConfig", reflect.TypeOf((*MockAuthConfigClient)(nil).GetAuthConfig), ctx, key)
}

// ListAuthConfig mocks base method.
func (m *MockAuthConfigClient) ListAuthConfig(ctx context.Context, opts ...client.ListOption) (*v1.AuthConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAuthConfig", varargs...)
	ret0, _ := ret[0].(*v1.AuthConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthConfig indicates an expected call of ListAuthConfig.
func (mr *MockAuthConfigClientMockRecorder) ListAuthConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthConfig", reflect.TypeOf((*MockAuthConfigClient)(nil).ListAuthConfig), varargs...)
}

// PatchAuthConfig mocks base method.
func (m *MockAuthConfigClient) PatchAuthConfig(ctx context.Context, obj *v1.AuthConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchAuthConfig indicates an expected call of PatchAuthConfig.
func (mr *MockAuthConfigClientMockRecorder) PatchAuthConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchAuthConfig", reflect.TypeOf((*MockAuthConfigClient)(nil).PatchAuthConfig), varargs...)
}

// PatchAuthConfigStatus mocks base method.
func (m *MockAuthConfigClient) PatchAuthConfigStatus(ctx context.Context, obj *v1.AuthConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchAuthConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchAuthConfigStatus indicates an expected call of PatchAuthConfigStatus.
func (mr *MockAuthConfigClientMockRecorder) PatchAuthConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchAuthConfigStatus", reflect.TypeOf((*MockAuthConfigClient)(nil).PatchAuthConfigStatus), varargs...)
}

// UpdateAuthConfig mocks base method.
func (m *MockAuthConfigClient) UpdateAuthConfig(ctx context.Context, obj *v1.AuthConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthConfig indicates an expected call of UpdateAuthConfig.
func (mr *MockAuthConfigClientMockRecorder) UpdateAuthConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthConfig", reflect.TypeOf((*MockAuthConfigClient)(nil).UpdateAuthConfig), varargs...)
}

// UpdateAuthConfigStatus mocks base method.
func (m *MockAuthConfigClient) UpdateAuthConfigStatus(ctx context.Context, obj *v1.AuthConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAuthConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthConfigStatus indicates an expected call of UpdateAuthConfigStatus.
func (mr *MockAuthConfigClientMockRecorder) UpdateAuthConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthConfigStatus", reflect.TypeOf((*MockAuthConfigClient)(nil).UpdateAuthConfigStatus), varargs...)
}

// UpsertAuthConfig mocks base method.
func (m *MockAuthConfigClient) UpsertAuthConfig(ctx context.Context, obj *v1.AuthConfig, transitionFuncs ...v1.AuthConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertAuthConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertAuthConfig indicates an expected call of UpsertAuthConfig.
func (mr *MockAuthConfigClientMockRecorder) UpsertAuthConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertAuthConfig", reflect.TypeOf((*MockAuthConfigClient)(nil).UpsertAuthConfig), varargs...)
}

// MockMulticlusterAuthConfigClient is a mock of MulticlusterAuthConfigClient interface.
type MockMulticlusterAuthConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAuthConfigClientMockRecorder
}

// MockMulticlusterAuthConfigClientMockRecorder is the mock recorder for MockMulticlusterAuthConfigClient.
type MockMulticlusterAuthConfigClientMockRecorder struct {
	mock *MockMulticlusterAuthConfigClient
}

// NewMockMulticlusterAuthConfigClient creates a new mock instance.
func NewMockMulticlusterAuthConfigClient(ctrl *gomock.Controller) *MockMulticlusterAuthConfigClient {
	mock := &MockMulticlusterAuthConfigClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAuthConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAuthConfigClient) EXPECT() *MockMulticlusterAuthConfigClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterAuthConfigClient) Cluster(cluster string) (v1.AuthConfigClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1.AuthConfigClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterAuthConfigClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterAuthConfigClient)(nil).Cluster), cluster)
}
