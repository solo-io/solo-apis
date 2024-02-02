// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1beta1 is a generated GoMock package.
package mock_v1beta1

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1beta1 "github.com/solo-io/solo-apis/pkg/api/graphql.gloo.solo.io/v1beta1"
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
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1beta1.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1beta1.Clientset)
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

// GraphQLApis mocks base method.
func (m *MockClientset) GraphQLApis() v1beta1.GraphQLApiClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GraphQLApis")
	ret0, _ := ret[0].(v1beta1.GraphQLApiClient)
	return ret0
}

// GraphQLApis indicates an expected call of GraphQLApis.
func (mr *MockClientsetMockRecorder) GraphQLApis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GraphQLApis", reflect.TypeOf((*MockClientset)(nil).GraphQLApis))
}

// MockGraphQLApiReader is a mock of GraphQLApiReader interface.
type MockGraphQLApiReader struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLApiReaderMockRecorder
}

// MockGraphQLApiReaderMockRecorder is the mock recorder for MockGraphQLApiReader.
type MockGraphQLApiReaderMockRecorder struct {
	mock *MockGraphQLApiReader
}

// NewMockGraphQLApiReader creates a new mock instance.
func NewMockGraphQLApiReader(ctrl *gomock.Controller) *MockGraphQLApiReader {
	mock := &MockGraphQLApiReader{ctrl: ctrl}
	mock.recorder = &MockGraphQLApiReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQLApiReader) EXPECT() *MockGraphQLApiReaderMockRecorder {
	return m.recorder
}

// GetGraphQLApi mocks base method.
func (m *MockGraphQLApiReader) GetGraphQLApi(ctx context.Context, key client.ObjectKey) (*v1beta1.GraphQLApi, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGraphQLApi", ctx, key)
	ret0, _ := ret[0].(*v1beta1.GraphQLApi)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGraphQLApi indicates an expected call of GetGraphQLApi.
func (mr *MockGraphQLApiReaderMockRecorder) GetGraphQLApi(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGraphQLApi", reflect.TypeOf((*MockGraphQLApiReader)(nil).GetGraphQLApi), ctx, key)
}

// ListGraphQLApi mocks base method.
func (m *MockGraphQLApiReader) ListGraphQLApi(ctx context.Context, opts ...client.ListOption) (*v1beta1.GraphQLApiList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGraphQLApi", varargs...)
	ret0, _ := ret[0].(*v1beta1.GraphQLApiList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGraphQLApi indicates an expected call of ListGraphQLApi.
func (mr *MockGraphQLApiReaderMockRecorder) ListGraphQLApi(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGraphQLApi", reflect.TypeOf((*MockGraphQLApiReader)(nil).ListGraphQLApi), varargs...)
}

// MockGraphQLApiWriter is a mock of GraphQLApiWriter interface.
type MockGraphQLApiWriter struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLApiWriterMockRecorder
}

// MockGraphQLApiWriterMockRecorder is the mock recorder for MockGraphQLApiWriter.
type MockGraphQLApiWriterMockRecorder struct {
	mock *MockGraphQLApiWriter
}

// NewMockGraphQLApiWriter creates a new mock instance.
func NewMockGraphQLApiWriter(ctrl *gomock.Controller) *MockGraphQLApiWriter {
	mock := &MockGraphQLApiWriter{ctrl: ctrl}
	mock.recorder = &MockGraphQLApiWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQLApiWriter) EXPECT() *MockGraphQLApiWriterMockRecorder {
	return m.recorder
}

// CreateGraphQLApi mocks base method.
func (m *MockGraphQLApiWriter) CreateGraphQLApi(ctx context.Context, obj *v1beta1.GraphQLApi, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGraphQLApi indicates an expected call of CreateGraphQLApi.
func (mr *MockGraphQLApiWriterMockRecorder) CreateGraphQLApi(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGraphQLApi", reflect.TypeOf((*MockGraphQLApiWriter)(nil).CreateGraphQLApi), varargs...)
}

// DeleteAllOfGraphQLApi mocks base method.
func (m *MockGraphQLApiWriter) DeleteAllOfGraphQLApi(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfGraphQLApi indicates an expected call of DeleteAllOfGraphQLApi.
func (mr *MockGraphQLApiWriterMockRecorder) DeleteAllOfGraphQLApi(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfGraphQLApi", reflect.TypeOf((*MockGraphQLApiWriter)(nil).DeleteAllOfGraphQLApi), varargs...)
}

// DeleteGraphQLApi mocks base method.
func (m *MockGraphQLApiWriter) DeleteGraphQLApi(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGraphQLApi indicates an expected call of DeleteGraphQLApi.
func (mr *MockGraphQLApiWriterMockRecorder) DeleteGraphQLApi(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGraphQLApi", reflect.TypeOf((*MockGraphQLApiWriter)(nil).DeleteGraphQLApi), varargs...)
}

// PatchGraphQLApi mocks base method.
func (m *MockGraphQLApiWriter) PatchGraphQLApi(ctx context.Context, obj *v1beta1.GraphQLApi, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchGraphQLApi indicates an expected call of PatchGraphQLApi.
func (mr *MockGraphQLApiWriterMockRecorder) PatchGraphQLApi(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchGraphQLApi", reflect.TypeOf((*MockGraphQLApiWriter)(nil).PatchGraphQLApi), varargs...)
}

// UpdateGraphQLApi mocks base method.
func (m *MockGraphQLApiWriter) UpdateGraphQLApi(ctx context.Context, obj *v1beta1.GraphQLApi, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGraphQLApi indicates an expected call of UpdateGraphQLApi.
func (mr *MockGraphQLApiWriterMockRecorder) UpdateGraphQLApi(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGraphQLApi", reflect.TypeOf((*MockGraphQLApiWriter)(nil).UpdateGraphQLApi), varargs...)
}

// UpsertGraphQLApi mocks base method.
func (m *MockGraphQLApiWriter) UpsertGraphQLApi(ctx context.Context, obj *v1beta1.GraphQLApi, transitionFuncs ...v1beta1.GraphQLApiTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertGraphQLApi indicates an expected call of UpsertGraphQLApi.
func (mr *MockGraphQLApiWriterMockRecorder) UpsertGraphQLApi(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertGraphQLApi", reflect.TypeOf((*MockGraphQLApiWriter)(nil).UpsertGraphQLApi), varargs...)
}

// MockGraphQLApiStatusWriter is a mock of GraphQLApiStatusWriter interface.
type MockGraphQLApiStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLApiStatusWriterMockRecorder
}

// MockGraphQLApiStatusWriterMockRecorder is the mock recorder for MockGraphQLApiStatusWriter.
type MockGraphQLApiStatusWriterMockRecorder struct {
	mock *MockGraphQLApiStatusWriter
}

// NewMockGraphQLApiStatusWriter creates a new mock instance.
func NewMockGraphQLApiStatusWriter(ctrl *gomock.Controller) *MockGraphQLApiStatusWriter {
	mock := &MockGraphQLApiStatusWriter{ctrl: ctrl}
	mock.recorder = &MockGraphQLApiStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQLApiStatusWriter) EXPECT() *MockGraphQLApiStatusWriterMockRecorder {
	return m.recorder
}

// PatchGraphQLApiStatus mocks base method.
func (m *MockGraphQLApiStatusWriter) PatchGraphQLApiStatus(ctx context.Context, obj *v1beta1.GraphQLApi, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchGraphQLApiStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchGraphQLApiStatus indicates an expected call of PatchGraphQLApiStatus.
func (mr *MockGraphQLApiStatusWriterMockRecorder) PatchGraphQLApiStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchGraphQLApiStatus", reflect.TypeOf((*MockGraphQLApiStatusWriter)(nil).PatchGraphQLApiStatus), varargs...)
}

// UpdateGraphQLApiStatus mocks base method.
func (m *MockGraphQLApiStatusWriter) UpdateGraphQLApiStatus(ctx context.Context, obj *v1beta1.GraphQLApi, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateGraphQLApiStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGraphQLApiStatus indicates an expected call of UpdateGraphQLApiStatus.
func (mr *MockGraphQLApiStatusWriterMockRecorder) UpdateGraphQLApiStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGraphQLApiStatus", reflect.TypeOf((*MockGraphQLApiStatusWriter)(nil).UpdateGraphQLApiStatus), varargs...)
}

// MockGraphQLApiClient is a mock of GraphQLApiClient interface.
type MockGraphQLApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLApiClientMockRecorder
}

// MockGraphQLApiClientMockRecorder is the mock recorder for MockGraphQLApiClient.
type MockGraphQLApiClientMockRecorder struct {
	mock *MockGraphQLApiClient
}

// NewMockGraphQLApiClient creates a new mock instance.
func NewMockGraphQLApiClient(ctrl *gomock.Controller) *MockGraphQLApiClient {
	mock := &MockGraphQLApiClient{ctrl: ctrl}
	mock.recorder = &MockGraphQLApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQLApiClient) EXPECT() *MockGraphQLApiClientMockRecorder {
	return m.recorder
}

// CreateGraphQLApi mocks base method.
func (m *MockGraphQLApiClient) CreateGraphQLApi(ctx context.Context, obj *v1beta1.GraphQLApi, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGraphQLApi indicates an expected call of CreateGraphQLApi.
func (mr *MockGraphQLApiClientMockRecorder) CreateGraphQLApi(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGraphQLApi", reflect.TypeOf((*MockGraphQLApiClient)(nil).CreateGraphQLApi), varargs...)
}

// DeleteAllOfGraphQLApi mocks base method.
func (m *MockGraphQLApiClient) DeleteAllOfGraphQLApi(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfGraphQLApi indicates an expected call of DeleteAllOfGraphQLApi.
func (mr *MockGraphQLApiClientMockRecorder) DeleteAllOfGraphQLApi(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfGraphQLApi", reflect.TypeOf((*MockGraphQLApiClient)(nil).DeleteAllOfGraphQLApi), varargs...)
}

// DeleteGraphQLApi mocks base method.
func (m *MockGraphQLApiClient) DeleteGraphQLApi(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGraphQLApi indicates an expected call of DeleteGraphQLApi.
func (mr *MockGraphQLApiClientMockRecorder) DeleteGraphQLApi(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGraphQLApi", reflect.TypeOf((*MockGraphQLApiClient)(nil).DeleteGraphQLApi), varargs...)
}

// GetGraphQLApi mocks base method.
func (m *MockGraphQLApiClient) GetGraphQLApi(ctx context.Context, key client.ObjectKey) (*v1beta1.GraphQLApi, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGraphQLApi", ctx, key)
	ret0, _ := ret[0].(*v1beta1.GraphQLApi)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGraphQLApi indicates an expected call of GetGraphQLApi.
func (mr *MockGraphQLApiClientMockRecorder) GetGraphQLApi(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGraphQLApi", reflect.TypeOf((*MockGraphQLApiClient)(nil).GetGraphQLApi), ctx, key)
}

// ListGraphQLApi mocks base method.
func (m *MockGraphQLApiClient) ListGraphQLApi(ctx context.Context, opts ...client.ListOption) (*v1beta1.GraphQLApiList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGraphQLApi", varargs...)
	ret0, _ := ret[0].(*v1beta1.GraphQLApiList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGraphQLApi indicates an expected call of ListGraphQLApi.
func (mr *MockGraphQLApiClientMockRecorder) ListGraphQLApi(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGraphQLApi", reflect.TypeOf((*MockGraphQLApiClient)(nil).ListGraphQLApi), varargs...)
}

// PatchGraphQLApi mocks base method.
func (m *MockGraphQLApiClient) PatchGraphQLApi(ctx context.Context, obj *v1beta1.GraphQLApi, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchGraphQLApi indicates an expected call of PatchGraphQLApi.
func (mr *MockGraphQLApiClientMockRecorder) PatchGraphQLApi(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchGraphQLApi", reflect.TypeOf((*MockGraphQLApiClient)(nil).PatchGraphQLApi), varargs...)
}

// PatchGraphQLApiStatus mocks base method.
func (m *MockGraphQLApiClient) PatchGraphQLApiStatus(ctx context.Context, obj *v1beta1.GraphQLApi, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchGraphQLApiStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchGraphQLApiStatus indicates an expected call of PatchGraphQLApiStatus.
func (mr *MockGraphQLApiClientMockRecorder) PatchGraphQLApiStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchGraphQLApiStatus", reflect.TypeOf((*MockGraphQLApiClient)(nil).PatchGraphQLApiStatus), varargs...)
}

// UpdateGraphQLApi mocks base method.
func (m *MockGraphQLApiClient) UpdateGraphQLApi(ctx context.Context, obj *v1beta1.GraphQLApi, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGraphQLApi indicates an expected call of UpdateGraphQLApi.
func (mr *MockGraphQLApiClientMockRecorder) UpdateGraphQLApi(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGraphQLApi", reflect.TypeOf((*MockGraphQLApiClient)(nil).UpdateGraphQLApi), varargs...)
}

// UpdateGraphQLApiStatus mocks base method.
func (m *MockGraphQLApiClient) UpdateGraphQLApiStatus(ctx context.Context, obj *v1beta1.GraphQLApi, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateGraphQLApiStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGraphQLApiStatus indicates an expected call of UpdateGraphQLApiStatus.
func (mr *MockGraphQLApiClientMockRecorder) UpdateGraphQLApiStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGraphQLApiStatus", reflect.TypeOf((*MockGraphQLApiClient)(nil).UpdateGraphQLApiStatus), varargs...)
}

// UpsertGraphQLApi mocks base method.
func (m *MockGraphQLApiClient) UpsertGraphQLApi(ctx context.Context, obj *v1beta1.GraphQLApi, transitionFuncs ...v1beta1.GraphQLApiTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertGraphQLApi", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertGraphQLApi indicates an expected call of UpsertGraphQLApi.
func (mr *MockGraphQLApiClientMockRecorder) UpsertGraphQLApi(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertGraphQLApi", reflect.TypeOf((*MockGraphQLApiClient)(nil).UpsertGraphQLApi), varargs...)
}

// MockMulticlusterGraphQLApiClient is a mock of MulticlusterGraphQLApiClient interface.
type MockMulticlusterGraphQLApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGraphQLApiClientMockRecorder
}

// MockMulticlusterGraphQLApiClientMockRecorder is the mock recorder for MockMulticlusterGraphQLApiClient.
type MockMulticlusterGraphQLApiClientMockRecorder struct {
	mock *MockMulticlusterGraphQLApiClient
}

// NewMockMulticlusterGraphQLApiClient creates a new mock instance.
func NewMockMulticlusterGraphQLApiClient(ctrl *gomock.Controller) *MockMulticlusterGraphQLApiClient {
	mock := &MockMulticlusterGraphQLApiClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGraphQLApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGraphQLApiClient) EXPECT() *MockMulticlusterGraphQLApiClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterGraphQLApiClient) Cluster(cluster string) (v1beta1.GraphQLApiClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1beta1.GraphQLApiClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterGraphQLApiClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterGraphQLApiClient)(nil).Cluster), cluster)
}
