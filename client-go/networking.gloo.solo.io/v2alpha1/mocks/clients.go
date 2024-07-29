// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v2alpha1 is a generated GoMock package.
package mock_v2alpha1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	v2alpha1 "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2alpha1"
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

// ExternalWorkloads mocks base method.
func (m *MockClientset) ExternalWorkloads() v2alpha1.ExternalWorkloadClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalWorkloads")
	ret0, _ := ret[0].(v2alpha1.ExternalWorkloadClient)
	return ret0
}

// ExternalWorkloads indicates an expected call of ExternalWorkloads.
func (mr *MockClientsetMockRecorder) ExternalWorkloads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalWorkloads", reflect.TypeOf((*MockClientset)(nil).ExternalWorkloads))
}

// MockExternalWorkloadReader is a mock of ExternalWorkloadReader interface.
type MockExternalWorkloadReader struct {
	ctrl     *gomock.Controller
	recorder *MockExternalWorkloadReaderMockRecorder
}

// MockExternalWorkloadReaderMockRecorder is the mock recorder for MockExternalWorkloadReader.
type MockExternalWorkloadReaderMockRecorder struct {
	mock *MockExternalWorkloadReader
}

// NewMockExternalWorkloadReader creates a new mock instance.
func NewMockExternalWorkloadReader(ctrl *gomock.Controller) *MockExternalWorkloadReader {
	mock := &MockExternalWorkloadReader{ctrl: ctrl}
	mock.recorder = &MockExternalWorkloadReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalWorkloadReader) EXPECT() *MockExternalWorkloadReaderMockRecorder {
	return m.recorder
}

// GetExternalWorkload mocks base method.
func (m *MockExternalWorkloadReader) GetExternalWorkload(ctx context.Context, key client.ObjectKey) (*v2alpha1.ExternalWorkload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalWorkload", ctx, key)
	ret0, _ := ret[0].(*v2alpha1.ExternalWorkload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalWorkload indicates an expected call of GetExternalWorkload.
func (mr *MockExternalWorkloadReaderMockRecorder) GetExternalWorkload(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalWorkload", reflect.TypeOf((*MockExternalWorkloadReader)(nil).GetExternalWorkload), ctx, key)
}

// ListExternalWorkload mocks base method.
func (m *MockExternalWorkloadReader) ListExternalWorkload(ctx context.Context, opts ...client.ListOption) (*v2alpha1.ExternalWorkloadList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListExternalWorkload", varargs...)
	ret0, _ := ret[0].(*v2alpha1.ExternalWorkloadList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExternalWorkload indicates an expected call of ListExternalWorkload.
func (mr *MockExternalWorkloadReaderMockRecorder) ListExternalWorkload(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExternalWorkload", reflect.TypeOf((*MockExternalWorkloadReader)(nil).ListExternalWorkload), varargs...)
}

// MockExternalWorkloadWriter is a mock of ExternalWorkloadWriter interface.
type MockExternalWorkloadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockExternalWorkloadWriterMockRecorder
}

// MockExternalWorkloadWriterMockRecorder is the mock recorder for MockExternalWorkloadWriter.
type MockExternalWorkloadWriterMockRecorder struct {
	mock *MockExternalWorkloadWriter
}

// NewMockExternalWorkloadWriter creates a new mock instance.
func NewMockExternalWorkloadWriter(ctrl *gomock.Controller) *MockExternalWorkloadWriter {
	mock := &MockExternalWorkloadWriter{ctrl: ctrl}
	mock.recorder = &MockExternalWorkloadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalWorkloadWriter) EXPECT() *MockExternalWorkloadWriterMockRecorder {
	return m.recorder
}

// CreateExternalWorkload mocks base method.
func (m *MockExternalWorkloadWriter) CreateExternalWorkload(ctx context.Context, obj *v2alpha1.ExternalWorkload, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExternalWorkload indicates an expected call of CreateExternalWorkload.
func (mr *MockExternalWorkloadWriterMockRecorder) CreateExternalWorkload(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalWorkload", reflect.TypeOf((*MockExternalWorkloadWriter)(nil).CreateExternalWorkload), varargs...)
}

// DeleteAllOfExternalWorkload mocks base method.
func (m *MockExternalWorkloadWriter) DeleteAllOfExternalWorkload(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfExternalWorkload indicates an expected call of DeleteAllOfExternalWorkload.
func (mr *MockExternalWorkloadWriterMockRecorder) DeleteAllOfExternalWorkload(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfExternalWorkload", reflect.TypeOf((*MockExternalWorkloadWriter)(nil).DeleteAllOfExternalWorkload), varargs...)
}

// DeleteExternalWorkload mocks base method.
func (m *MockExternalWorkloadWriter) DeleteExternalWorkload(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalWorkload indicates an expected call of DeleteExternalWorkload.
func (mr *MockExternalWorkloadWriterMockRecorder) DeleteExternalWorkload(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalWorkload", reflect.TypeOf((*MockExternalWorkloadWriter)(nil).DeleteExternalWorkload), varargs...)
}

// PatchExternalWorkload mocks base method.
func (m *MockExternalWorkloadWriter) PatchExternalWorkload(ctx context.Context, obj *v2alpha1.ExternalWorkload, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchExternalWorkload indicates an expected call of PatchExternalWorkload.
func (mr *MockExternalWorkloadWriterMockRecorder) PatchExternalWorkload(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchExternalWorkload", reflect.TypeOf((*MockExternalWorkloadWriter)(nil).PatchExternalWorkload), varargs...)
}

// UpdateExternalWorkload mocks base method.
func (m *MockExternalWorkloadWriter) UpdateExternalWorkload(ctx context.Context, obj *v2alpha1.ExternalWorkload, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalWorkload indicates an expected call of UpdateExternalWorkload.
func (mr *MockExternalWorkloadWriterMockRecorder) UpdateExternalWorkload(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalWorkload", reflect.TypeOf((*MockExternalWorkloadWriter)(nil).UpdateExternalWorkload), varargs...)
}

// UpsertExternalWorkload mocks base method.
func (m *MockExternalWorkloadWriter) UpsertExternalWorkload(ctx context.Context, obj *v2alpha1.ExternalWorkload, transitionFuncs ...v2alpha1.ExternalWorkloadTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertExternalWorkload indicates an expected call of UpsertExternalWorkload.
func (mr *MockExternalWorkloadWriterMockRecorder) UpsertExternalWorkload(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertExternalWorkload", reflect.TypeOf((*MockExternalWorkloadWriter)(nil).UpsertExternalWorkload), varargs...)
}

// MockExternalWorkloadStatusWriter is a mock of ExternalWorkloadStatusWriter interface.
type MockExternalWorkloadStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockExternalWorkloadStatusWriterMockRecorder
}

// MockExternalWorkloadStatusWriterMockRecorder is the mock recorder for MockExternalWorkloadStatusWriter.
type MockExternalWorkloadStatusWriterMockRecorder struct {
	mock *MockExternalWorkloadStatusWriter
}

// NewMockExternalWorkloadStatusWriter creates a new mock instance.
func NewMockExternalWorkloadStatusWriter(ctrl *gomock.Controller) *MockExternalWorkloadStatusWriter {
	mock := &MockExternalWorkloadStatusWriter{ctrl: ctrl}
	mock.recorder = &MockExternalWorkloadStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalWorkloadStatusWriter) EXPECT() *MockExternalWorkloadStatusWriterMockRecorder {
	return m.recorder
}

// PatchExternalWorkloadStatus mocks base method.
func (m *MockExternalWorkloadStatusWriter) PatchExternalWorkloadStatus(ctx context.Context, obj *v2alpha1.ExternalWorkload, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchExternalWorkloadStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchExternalWorkloadStatus indicates an expected call of PatchExternalWorkloadStatus.
func (mr *MockExternalWorkloadStatusWriterMockRecorder) PatchExternalWorkloadStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchExternalWorkloadStatus", reflect.TypeOf((*MockExternalWorkloadStatusWriter)(nil).PatchExternalWorkloadStatus), varargs...)
}

// UpdateExternalWorkloadStatus mocks base method.
func (m *MockExternalWorkloadStatusWriter) UpdateExternalWorkloadStatus(ctx context.Context, obj *v2alpha1.ExternalWorkload, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateExternalWorkloadStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalWorkloadStatus indicates an expected call of UpdateExternalWorkloadStatus.
func (mr *MockExternalWorkloadStatusWriterMockRecorder) UpdateExternalWorkloadStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalWorkloadStatus", reflect.TypeOf((*MockExternalWorkloadStatusWriter)(nil).UpdateExternalWorkloadStatus), varargs...)
}

// MockExternalWorkloadClient is a mock of ExternalWorkloadClient interface.
type MockExternalWorkloadClient struct {
	ctrl     *gomock.Controller
	recorder *MockExternalWorkloadClientMockRecorder
}

// MockExternalWorkloadClientMockRecorder is the mock recorder for MockExternalWorkloadClient.
type MockExternalWorkloadClientMockRecorder struct {
	mock *MockExternalWorkloadClient
}

// NewMockExternalWorkloadClient creates a new mock instance.
func NewMockExternalWorkloadClient(ctrl *gomock.Controller) *MockExternalWorkloadClient {
	mock := &MockExternalWorkloadClient{ctrl: ctrl}
	mock.recorder = &MockExternalWorkloadClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalWorkloadClient) EXPECT() *MockExternalWorkloadClientMockRecorder {
	return m.recorder
}

// CreateExternalWorkload mocks base method.
func (m *MockExternalWorkloadClient) CreateExternalWorkload(ctx context.Context, obj *v2alpha1.ExternalWorkload, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExternalWorkload indicates an expected call of CreateExternalWorkload.
func (mr *MockExternalWorkloadClientMockRecorder) CreateExternalWorkload(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalWorkload", reflect.TypeOf((*MockExternalWorkloadClient)(nil).CreateExternalWorkload), varargs...)
}

// DeleteAllOfExternalWorkload mocks base method.
func (m *MockExternalWorkloadClient) DeleteAllOfExternalWorkload(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfExternalWorkload indicates an expected call of DeleteAllOfExternalWorkload.
func (mr *MockExternalWorkloadClientMockRecorder) DeleteAllOfExternalWorkload(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfExternalWorkload", reflect.TypeOf((*MockExternalWorkloadClient)(nil).DeleteAllOfExternalWorkload), varargs...)
}

// DeleteExternalWorkload mocks base method.
func (m *MockExternalWorkloadClient) DeleteExternalWorkload(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalWorkload indicates an expected call of DeleteExternalWorkload.
func (mr *MockExternalWorkloadClientMockRecorder) DeleteExternalWorkload(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalWorkload", reflect.TypeOf((*MockExternalWorkloadClient)(nil).DeleteExternalWorkload), varargs...)
}

// GetExternalWorkload mocks base method.
func (m *MockExternalWorkloadClient) GetExternalWorkload(ctx context.Context, key client.ObjectKey) (*v2alpha1.ExternalWorkload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalWorkload", ctx, key)
	ret0, _ := ret[0].(*v2alpha1.ExternalWorkload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalWorkload indicates an expected call of GetExternalWorkload.
func (mr *MockExternalWorkloadClientMockRecorder) GetExternalWorkload(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalWorkload", reflect.TypeOf((*MockExternalWorkloadClient)(nil).GetExternalWorkload), ctx, key)
}

// ListExternalWorkload mocks base method.
func (m *MockExternalWorkloadClient) ListExternalWorkload(ctx context.Context, opts ...client.ListOption) (*v2alpha1.ExternalWorkloadList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListExternalWorkload", varargs...)
	ret0, _ := ret[0].(*v2alpha1.ExternalWorkloadList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExternalWorkload indicates an expected call of ListExternalWorkload.
func (mr *MockExternalWorkloadClientMockRecorder) ListExternalWorkload(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExternalWorkload", reflect.TypeOf((*MockExternalWorkloadClient)(nil).ListExternalWorkload), varargs...)
}

// PatchExternalWorkload mocks base method.
func (m *MockExternalWorkloadClient) PatchExternalWorkload(ctx context.Context, obj *v2alpha1.ExternalWorkload, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchExternalWorkload indicates an expected call of PatchExternalWorkload.
func (mr *MockExternalWorkloadClientMockRecorder) PatchExternalWorkload(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchExternalWorkload", reflect.TypeOf((*MockExternalWorkloadClient)(nil).PatchExternalWorkload), varargs...)
}

// PatchExternalWorkloadStatus mocks base method.
func (m *MockExternalWorkloadClient) PatchExternalWorkloadStatus(ctx context.Context, obj *v2alpha1.ExternalWorkload, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchExternalWorkloadStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchExternalWorkloadStatus indicates an expected call of PatchExternalWorkloadStatus.
func (mr *MockExternalWorkloadClientMockRecorder) PatchExternalWorkloadStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchExternalWorkloadStatus", reflect.TypeOf((*MockExternalWorkloadClient)(nil).PatchExternalWorkloadStatus), varargs...)
}

// UpdateExternalWorkload mocks base method.
func (m *MockExternalWorkloadClient) UpdateExternalWorkload(ctx context.Context, obj *v2alpha1.ExternalWorkload, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalWorkload indicates an expected call of UpdateExternalWorkload.
func (mr *MockExternalWorkloadClientMockRecorder) UpdateExternalWorkload(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalWorkload", reflect.TypeOf((*MockExternalWorkloadClient)(nil).UpdateExternalWorkload), varargs...)
}

// UpdateExternalWorkloadStatus mocks base method.
func (m *MockExternalWorkloadClient) UpdateExternalWorkloadStatus(ctx context.Context, obj *v2alpha1.ExternalWorkload, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateExternalWorkloadStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalWorkloadStatus indicates an expected call of UpdateExternalWorkloadStatus.
func (mr *MockExternalWorkloadClientMockRecorder) UpdateExternalWorkloadStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalWorkloadStatus", reflect.TypeOf((*MockExternalWorkloadClient)(nil).UpdateExternalWorkloadStatus), varargs...)
}

// UpsertExternalWorkload mocks base method.
func (m *MockExternalWorkloadClient) UpsertExternalWorkload(ctx context.Context, obj *v2alpha1.ExternalWorkload, transitionFuncs ...v2alpha1.ExternalWorkloadTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertExternalWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertExternalWorkload indicates an expected call of UpsertExternalWorkload.
func (mr *MockExternalWorkloadClientMockRecorder) UpsertExternalWorkload(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertExternalWorkload", reflect.TypeOf((*MockExternalWorkloadClient)(nil).UpsertExternalWorkload), varargs...)
}

// MockMulticlusterExternalWorkloadClient is a mock of MulticlusterExternalWorkloadClient interface.
type MockMulticlusterExternalWorkloadClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterExternalWorkloadClientMockRecorder
}

// MockMulticlusterExternalWorkloadClientMockRecorder is the mock recorder for MockMulticlusterExternalWorkloadClient.
type MockMulticlusterExternalWorkloadClientMockRecorder struct {
	mock *MockMulticlusterExternalWorkloadClient
}

// NewMockMulticlusterExternalWorkloadClient creates a new mock instance.
func NewMockMulticlusterExternalWorkloadClient(ctrl *gomock.Controller) *MockMulticlusterExternalWorkloadClient {
	mock := &MockMulticlusterExternalWorkloadClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterExternalWorkloadClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterExternalWorkloadClient) EXPECT() *MockMulticlusterExternalWorkloadClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterExternalWorkloadClient) Cluster(cluster string) (v2alpha1.ExternalWorkloadClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v2alpha1.ExternalWorkloadClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterExternalWorkloadClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterExternalWorkloadClient)(nil).Cluster), cluster)
}