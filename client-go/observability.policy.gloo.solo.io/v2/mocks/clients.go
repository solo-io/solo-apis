// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v2 is a generated GoMock package.
package mock_v2

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	v2 "github.com/solo-io/solo-apis/client-go/observability.policy.gloo.solo.io/v2"
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
func (m *MockMulticlusterClientset) Cluster(cluster string) (v2.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v2.Clientset)
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

// AccessLogPolicies mocks base method.
func (m *MockClientset) AccessLogPolicies() v2.AccessLogPolicyClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessLogPolicies")
	ret0, _ := ret[0].(v2.AccessLogPolicyClient)
	return ret0
}

// AccessLogPolicies indicates an expected call of AccessLogPolicies.
func (mr *MockClientsetMockRecorder) AccessLogPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessLogPolicies", reflect.TypeOf((*MockClientset)(nil).AccessLogPolicies))
}

// MockAccessLogPolicyReader is a mock of AccessLogPolicyReader interface.
type MockAccessLogPolicyReader struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogPolicyReaderMockRecorder
}

// MockAccessLogPolicyReaderMockRecorder is the mock recorder for MockAccessLogPolicyReader.
type MockAccessLogPolicyReaderMockRecorder struct {
	mock *MockAccessLogPolicyReader
}

// NewMockAccessLogPolicyReader creates a new mock instance.
func NewMockAccessLogPolicyReader(ctrl *gomock.Controller) *MockAccessLogPolicyReader {
	mock := &MockAccessLogPolicyReader{ctrl: ctrl}
	mock.recorder = &MockAccessLogPolicyReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogPolicyReader) EXPECT() *MockAccessLogPolicyReaderMockRecorder {
	return m.recorder
}

// GetAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyReader) GetAccessLogPolicy(ctx context.Context, key client.ObjectKey) (*v2.AccessLogPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessLogPolicy", ctx, key)
	ret0, _ := ret[0].(*v2.AccessLogPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessLogPolicy indicates an expected call of GetAccessLogPolicy.
func (mr *MockAccessLogPolicyReaderMockRecorder) GetAccessLogPolicy(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyReader)(nil).GetAccessLogPolicy), ctx, key)
}

// ListAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyReader) ListAccessLogPolicy(ctx context.Context, opts ...client.ListOption) (*v2.AccessLogPolicyList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(*v2.AccessLogPolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessLogPolicy indicates an expected call of ListAccessLogPolicy.
func (mr *MockAccessLogPolicyReaderMockRecorder) ListAccessLogPolicy(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyReader)(nil).ListAccessLogPolicy), varargs...)
}

// MockAccessLogPolicyWriter is a mock of AccessLogPolicyWriter interface.
type MockAccessLogPolicyWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogPolicyWriterMockRecorder
}

// MockAccessLogPolicyWriterMockRecorder is the mock recorder for MockAccessLogPolicyWriter.
type MockAccessLogPolicyWriterMockRecorder struct {
	mock *MockAccessLogPolicyWriter
}

// NewMockAccessLogPolicyWriter creates a new mock instance.
func NewMockAccessLogPolicyWriter(ctrl *gomock.Controller) *MockAccessLogPolicyWriter {
	mock := &MockAccessLogPolicyWriter{ctrl: ctrl}
	mock.recorder = &MockAccessLogPolicyWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogPolicyWriter) EXPECT() *MockAccessLogPolicyWriterMockRecorder {
	return m.recorder
}

// CreateAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyWriter) CreateAccessLogPolicy(ctx context.Context, obj *v2.AccessLogPolicy, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessLogPolicy indicates an expected call of CreateAccessLogPolicy.
func (mr *MockAccessLogPolicyWriterMockRecorder) CreateAccessLogPolicy(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyWriter)(nil).CreateAccessLogPolicy), varargs...)
}

// DeleteAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyWriter) DeleteAccessLogPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessLogPolicy indicates an expected call of DeleteAccessLogPolicy.
func (mr *MockAccessLogPolicyWriterMockRecorder) DeleteAccessLogPolicy(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyWriter)(nil).DeleteAccessLogPolicy), varargs...)
}

// DeleteAllOfAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyWriter) DeleteAllOfAccessLogPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfAccessLogPolicy indicates an expected call of DeleteAllOfAccessLogPolicy.
func (mr *MockAccessLogPolicyWriterMockRecorder) DeleteAllOfAccessLogPolicy(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyWriter)(nil).DeleteAllOfAccessLogPolicy), varargs...)
}

// PatchAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyWriter) PatchAccessLogPolicy(ctx context.Context, obj *v2.AccessLogPolicy, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchAccessLogPolicy indicates an expected call of PatchAccessLogPolicy.
func (mr *MockAccessLogPolicyWriterMockRecorder) PatchAccessLogPolicy(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyWriter)(nil).PatchAccessLogPolicy), varargs...)
}

// UpdateAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyWriter) UpdateAccessLogPolicy(ctx context.Context, obj *v2.AccessLogPolicy, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccessLogPolicy indicates an expected call of UpdateAccessLogPolicy.
func (mr *MockAccessLogPolicyWriterMockRecorder) UpdateAccessLogPolicy(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyWriter)(nil).UpdateAccessLogPolicy), varargs...)
}

// UpsertAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyWriter) UpsertAccessLogPolicy(ctx context.Context, obj *v2.AccessLogPolicy, transitionFuncs ...v2.AccessLogPolicyTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertAccessLogPolicy indicates an expected call of UpsertAccessLogPolicy.
func (mr *MockAccessLogPolicyWriterMockRecorder) UpsertAccessLogPolicy(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyWriter)(nil).UpsertAccessLogPolicy), varargs...)
}

// MockAccessLogPolicyStatusWriter is a mock of AccessLogPolicyStatusWriter interface.
type MockAccessLogPolicyStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogPolicyStatusWriterMockRecorder
}

// MockAccessLogPolicyStatusWriterMockRecorder is the mock recorder for MockAccessLogPolicyStatusWriter.
type MockAccessLogPolicyStatusWriterMockRecorder struct {
	mock *MockAccessLogPolicyStatusWriter
}

// NewMockAccessLogPolicyStatusWriter creates a new mock instance.
func NewMockAccessLogPolicyStatusWriter(ctrl *gomock.Controller) *MockAccessLogPolicyStatusWriter {
	mock := &MockAccessLogPolicyStatusWriter{ctrl: ctrl}
	mock.recorder = &MockAccessLogPolicyStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogPolicyStatusWriter) EXPECT() *MockAccessLogPolicyStatusWriterMockRecorder {
	return m.recorder
}

// PatchAccessLogPolicyStatus mocks base method.
func (m *MockAccessLogPolicyStatusWriter) PatchAccessLogPolicyStatus(ctx context.Context, obj *v2.AccessLogPolicy, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchAccessLogPolicyStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchAccessLogPolicyStatus indicates an expected call of PatchAccessLogPolicyStatus.
func (mr *MockAccessLogPolicyStatusWriterMockRecorder) PatchAccessLogPolicyStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchAccessLogPolicyStatus", reflect.TypeOf((*MockAccessLogPolicyStatusWriter)(nil).PatchAccessLogPolicyStatus), varargs...)
}

// UpdateAccessLogPolicyStatus mocks base method.
func (m *MockAccessLogPolicyStatusWriter) UpdateAccessLogPolicyStatus(ctx context.Context, obj *v2.AccessLogPolicy, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAccessLogPolicyStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccessLogPolicyStatus indicates an expected call of UpdateAccessLogPolicyStatus.
func (mr *MockAccessLogPolicyStatusWriterMockRecorder) UpdateAccessLogPolicyStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessLogPolicyStatus", reflect.TypeOf((*MockAccessLogPolicyStatusWriter)(nil).UpdateAccessLogPolicyStatus), varargs...)
}

// MockAccessLogPolicyClient is a mock of AccessLogPolicyClient interface.
type MockAccessLogPolicyClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogPolicyClientMockRecorder
}

// MockAccessLogPolicyClientMockRecorder is the mock recorder for MockAccessLogPolicyClient.
type MockAccessLogPolicyClientMockRecorder struct {
	mock *MockAccessLogPolicyClient
}

// NewMockAccessLogPolicyClient creates a new mock instance.
func NewMockAccessLogPolicyClient(ctrl *gomock.Controller) *MockAccessLogPolicyClient {
	mock := &MockAccessLogPolicyClient{ctrl: ctrl}
	mock.recorder = &MockAccessLogPolicyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogPolicyClient) EXPECT() *MockAccessLogPolicyClientMockRecorder {
	return m.recorder
}

// CreateAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyClient) CreateAccessLogPolicy(ctx context.Context, obj *v2.AccessLogPolicy, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessLogPolicy indicates an expected call of CreateAccessLogPolicy.
func (mr *MockAccessLogPolicyClientMockRecorder) CreateAccessLogPolicy(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).CreateAccessLogPolicy), varargs...)
}

// DeleteAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyClient) DeleteAccessLogPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessLogPolicy indicates an expected call of DeleteAccessLogPolicy.
func (mr *MockAccessLogPolicyClientMockRecorder) DeleteAccessLogPolicy(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).DeleteAccessLogPolicy), varargs...)
}

// DeleteAllOfAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyClient) DeleteAllOfAccessLogPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfAccessLogPolicy indicates an expected call of DeleteAllOfAccessLogPolicy.
func (mr *MockAccessLogPolicyClientMockRecorder) DeleteAllOfAccessLogPolicy(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).DeleteAllOfAccessLogPolicy), varargs...)
}

// GetAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyClient) GetAccessLogPolicy(ctx context.Context, key client.ObjectKey) (*v2.AccessLogPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessLogPolicy", ctx, key)
	ret0, _ := ret[0].(*v2.AccessLogPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessLogPolicy indicates an expected call of GetAccessLogPolicy.
func (mr *MockAccessLogPolicyClientMockRecorder) GetAccessLogPolicy(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).GetAccessLogPolicy), ctx, key)
}

// ListAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyClient) ListAccessLogPolicy(ctx context.Context, opts ...client.ListOption) (*v2.AccessLogPolicyList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(*v2.AccessLogPolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessLogPolicy indicates an expected call of ListAccessLogPolicy.
func (mr *MockAccessLogPolicyClientMockRecorder) ListAccessLogPolicy(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).ListAccessLogPolicy), varargs...)
}

// PatchAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyClient) PatchAccessLogPolicy(ctx context.Context, obj *v2.AccessLogPolicy, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchAccessLogPolicy indicates an expected call of PatchAccessLogPolicy.
func (mr *MockAccessLogPolicyClientMockRecorder) PatchAccessLogPolicy(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).PatchAccessLogPolicy), varargs...)
}

// PatchAccessLogPolicyStatus mocks base method.
func (m *MockAccessLogPolicyClient) PatchAccessLogPolicyStatus(ctx context.Context, obj *v2.AccessLogPolicy, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchAccessLogPolicyStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchAccessLogPolicyStatus indicates an expected call of PatchAccessLogPolicyStatus.
func (mr *MockAccessLogPolicyClientMockRecorder) PatchAccessLogPolicyStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchAccessLogPolicyStatus", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).PatchAccessLogPolicyStatus), varargs...)
}

// UpdateAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyClient) UpdateAccessLogPolicy(ctx context.Context, obj *v2.AccessLogPolicy, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccessLogPolicy indicates an expected call of UpdateAccessLogPolicy.
func (mr *MockAccessLogPolicyClientMockRecorder) UpdateAccessLogPolicy(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).UpdateAccessLogPolicy), varargs...)
}

// UpdateAccessLogPolicyStatus mocks base method.
func (m *MockAccessLogPolicyClient) UpdateAccessLogPolicyStatus(ctx context.Context, obj *v2.AccessLogPolicy, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAccessLogPolicyStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccessLogPolicyStatus indicates an expected call of UpdateAccessLogPolicyStatus.
func (mr *MockAccessLogPolicyClientMockRecorder) UpdateAccessLogPolicyStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessLogPolicyStatus", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).UpdateAccessLogPolicyStatus), varargs...)
}

// UpsertAccessLogPolicy mocks base method.
func (m *MockAccessLogPolicyClient) UpsertAccessLogPolicy(ctx context.Context, obj *v2.AccessLogPolicy, transitionFuncs ...v2.AccessLogPolicyTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertAccessLogPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertAccessLogPolicy indicates an expected call of UpsertAccessLogPolicy.
func (mr *MockAccessLogPolicyClientMockRecorder) UpsertAccessLogPolicy(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertAccessLogPolicy", reflect.TypeOf((*MockAccessLogPolicyClient)(nil).UpsertAccessLogPolicy), varargs...)
}

// MockMulticlusterAccessLogPolicyClient is a mock of MulticlusterAccessLogPolicyClient interface.
type MockMulticlusterAccessLogPolicyClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessLogPolicyClientMockRecorder
}

// MockMulticlusterAccessLogPolicyClientMockRecorder is the mock recorder for MockMulticlusterAccessLogPolicyClient.
type MockMulticlusterAccessLogPolicyClientMockRecorder struct {
	mock *MockMulticlusterAccessLogPolicyClient
}

// NewMockMulticlusterAccessLogPolicyClient creates a new mock instance.
func NewMockMulticlusterAccessLogPolicyClient(ctrl *gomock.Controller) *MockMulticlusterAccessLogPolicyClient {
	mock := &MockMulticlusterAccessLogPolicyClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessLogPolicyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAccessLogPolicyClient) EXPECT() *MockMulticlusterAccessLogPolicyClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterAccessLogPolicyClient) Cluster(cluster string) (v2.AccessLogPolicyClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v2.AccessLogPolicyClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterAccessLogPolicyClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterAccessLogPolicyClient)(nil).Cluster), cluster)
}
