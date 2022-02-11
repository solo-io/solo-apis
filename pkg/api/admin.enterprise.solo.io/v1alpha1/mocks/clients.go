// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1alpha1 is a generated GoMock package.
package mock_v1alpha1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/admin.enterprise.solo.io/v1alpha1"
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

// IstioInstallations mocks base method
func (m *MockClientset) IstioInstallations() v1alpha1.IstioInstallationClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IstioInstallations")
	ret0, _ := ret[0].(v1alpha1.IstioInstallationClient)
	return ret0
}

// IstioInstallations indicates an expected call of IstioInstallations
func (mr *MockClientsetMockRecorder) IstioInstallations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IstioInstallations", reflect.TypeOf((*MockClientset)(nil).IstioInstallations))
}

// MockIstioInstallationReader is a mock of IstioInstallationReader interface
type MockIstioInstallationReader struct {
	ctrl     *gomock.Controller
	recorder *MockIstioInstallationReaderMockRecorder
}

// MockIstioInstallationReaderMockRecorder is the mock recorder for MockIstioInstallationReader
type MockIstioInstallationReaderMockRecorder struct {
	mock *MockIstioInstallationReader
}

// NewMockIstioInstallationReader creates a new mock instance
func NewMockIstioInstallationReader(ctrl *gomock.Controller) *MockIstioInstallationReader {
	mock := &MockIstioInstallationReader{ctrl: ctrl}
	mock.recorder = &MockIstioInstallationReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIstioInstallationReader) EXPECT() *MockIstioInstallationReaderMockRecorder {
	return m.recorder
}

// GetIstioInstallation mocks base method
func (m *MockIstioInstallationReader) GetIstioInstallation(ctx context.Context, key client.ObjectKey) (*v1alpha1.IstioInstallation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIstioInstallation", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.IstioInstallation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIstioInstallation indicates an expected call of GetIstioInstallation
func (mr *MockIstioInstallationReaderMockRecorder) GetIstioInstallation(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIstioInstallation", reflect.TypeOf((*MockIstioInstallationReader)(nil).GetIstioInstallation), ctx, key)
}

// ListIstioInstallation mocks base method
func (m *MockIstioInstallationReader) ListIstioInstallation(ctx context.Context, opts ...client.ListOption) (*v1alpha1.IstioInstallationList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIstioInstallation", varargs...)
	ret0, _ := ret[0].(*v1alpha1.IstioInstallationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIstioInstallation indicates an expected call of ListIstioInstallation
func (mr *MockIstioInstallationReaderMockRecorder) ListIstioInstallation(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIstioInstallation", reflect.TypeOf((*MockIstioInstallationReader)(nil).ListIstioInstallation), varargs...)
}

// MockIstioInstallationWriter is a mock of IstioInstallationWriter interface
type MockIstioInstallationWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIstioInstallationWriterMockRecorder
}

// MockIstioInstallationWriterMockRecorder is the mock recorder for MockIstioInstallationWriter
type MockIstioInstallationWriterMockRecorder struct {
	mock *MockIstioInstallationWriter
}

// NewMockIstioInstallationWriter creates a new mock instance
func NewMockIstioInstallationWriter(ctrl *gomock.Controller) *MockIstioInstallationWriter {
	mock := &MockIstioInstallationWriter{ctrl: ctrl}
	mock.recorder = &MockIstioInstallationWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIstioInstallationWriter) EXPECT() *MockIstioInstallationWriterMockRecorder {
	return m.recorder
}

// CreateIstioInstallation mocks base method
func (m *MockIstioInstallationWriter) CreateIstioInstallation(ctx context.Context, obj *v1alpha1.IstioInstallation, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIstioInstallation indicates an expected call of CreateIstioInstallation
func (mr *MockIstioInstallationWriterMockRecorder) CreateIstioInstallation(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIstioInstallation", reflect.TypeOf((*MockIstioInstallationWriter)(nil).CreateIstioInstallation), varargs...)
}

// DeleteIstioInstallation mocks base method
func (m *MockIstioInstallationWriter) DeleteIstioInstallation(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIstioInstallation indicates an expected call of DeleteIstioInstallation
func (mr *MockIstioInstallationWriterMockRecorder) DeleteIstioInstallation(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIstioInstallation", reflect.TypeOf((*MockIstioInstallationWriter)(nil).DeleteIstioInstallation), varargs...)
}

// UpdateIstioInstallation mocks base method
func (m *MockIstioInstallationWriter) UpdateIstioInstallation(ctx context.Context, obj *v1alpha1.IstioInstallation, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIstioInstallation indicates an expected call of UpdateIstioInstallation
func (mr *MockIstioInstallationWriterMockRecorder) UpdateIstioInstallation(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIstioInstallation", reflect.TypeOf((*MockIstioInstallationWriter)(nil).UpdateIstioInstallation), varargs...)
}

// PatchIstioInstallation mocks base method
func (m *MockIstioInstallationWriter) PatchIstioInstallation(ctx context.Context, obj *v1alpha1.IstioInstallation, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchIstioInstallation indicates an expected call of PatchIstioInstallation
func (mr *MockIstioInstallationWriterMockRecorder) PatchIstioInstallation(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchIstioInstallation", reflect.TypeOf((*MockIstioInstallationWriter)(nil).PatchIstioInstallation), varargs...)
}

// DeleteAllOfIstioInstallation mocks base method
func (m *MockIstioInstallationWriter) DeleteAllOfIstioInstallation(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfIstioInstallation indicates an expected call of DeleteAllOfIstioInstallation
func (mr *MockIstioInstallationWriterMockRecorder) DeleteAllOfIstioInstallation(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfIstioInstallation", reflect.TypeOf((*MockIstioInstallationWriter)(nil).DeleteAllOfIstioInstallation), varargs...)
}

// UpsertIstioInstallation mocks base method
func (m *MockIstioInstallationWriter) UpsertIstioInstallation(ctx context.Context, obj *v1alpha1.IstioInstallation, transitionFuncs ...v1alpha1.IstioInstallationTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertIstioInstallation indicates an expected call of UpsertIstioInstallation
func (mr *MockIstioInstallationWriterMockRecorder) UpsertIstioInstallation(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertIstioInstallation", reflect.TypeOf((*MockIstioInstallationWriter)(nil).UpsertIstioInstallation), varargs...)
}

// MockIstioInstallationStatusWriter is a mock of IstioInstallationStatusWriter interface
type MockIstioInstallationStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIstioInstallationStatusWriterMockRecorder
}

// MockIstioInstallationStatusWriterMockRecorder is the mock recorder for MockIstioInstallationStatusWriter
type MockIstioInstallationStatusWriterMockRecorder struct {
	mock *MockIstioInstallationStatusWriter
}

// NewMockIstioInstallationStatusWriter creates a new mock instance
func NewMockIstioInstallationStatusWriter(ctrl *gomock.Controller) *MockIstioInstallationStatusWriter {
	mock := &MockIstioInstallationStatusWriter{ctrl: ctrl}
	mock.recorder = &MockIstioInstallationStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIstioInstallationStatusWriter) EXPECT() *MockIstioInstallationStatusWriterMockRecorder {
	return m.recorder
}

// UpdateIstioInstallationStatus mocks base method
func (m *MockIstioInstallationStatusWriter) UpdateIstioInstallationStatus(ctx context.Context, obj *v1alpha1.IstioInstallation, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIstioInstallationStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIstioInstallationStatus indicates an expected call of UpdateIstioInstallationStatus
func (mr *MockIstioInstallationStatusWriterMockRecorder) UpdateIstioInstallationStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIstioInstallationStatus", reflect.TypeOf((*MockIstioInstallationStatusWriter)(nil).UpdateIstioInstallationStatus), varargs...)
}

// PatchIstioInstallationStatus mocks base method
func (m *MockIstioInstallationStatusWriter) PatchIstioInstallationStatus(ctx context.Context, obj *v1alpha1.IstioInstallation, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchIstioInstallationStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchIstioInstallationStatus indicates an expected call of PatchIstioInstallationStatus
func (mr *MockIstioInstallationStatusWriterMockRecorder) PatchIstioInstallationStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchIstioInstallationStatus", reflect.TypeOf((*MockIstioInstallationStatusWriter)(nil).PatchIstioInstallationStatus), varargs...)
}

// MockIstioInstallationClient is a mock of IstioInstallationClient interface
type MockIstioInstallationClient struct {
	ctrl     *gomock.Controller
	recorder *MockIstioInstallationClientMockRecorder
}

// MockIstioInstallationClientMockRecorder is the mock recorder for MockIstioInstallationClient
type MockIstioInstallationClientMockRecorder struct {
	mock *MockIstioInstallationClient
}

// NewMockIstioInstallationClient creates a new mock instance
func NewMockIstioInstallationClient(ctrl *gomock.Controller) *MockIstioInstallationClient {
	mock := &MockIstioInstallationClient{ctrl: ctrl}
	mock.recorder = &MockIstioInstallationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIstioInstallationClient) EXPECT() *MockIstioInstallationClientMockRecorder {
	return m.recorder
}

// GetIstioInstallation mocks base method
func (m *MockIstioInstallationClient) GetIstioInstallation(ctx context.Context, key client.ObjectKey) (*v1alpha1.IstioInstallation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIstioInstallation", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.IstioInstallation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIstioInstallation indicates an expected call of GetIstioInstallation
func (mr *MockIstioInstallationClientMockRecorder) GetIstioInstallation(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIstioInstallation", reflect.TypeOf((*MockIstioInstallationClient)(nil).GetIstioInstallation), ctx, key)
}

// ListIstioInstallation mocks base method
func (m *MockIstioInstallationClient) ListIstioInstallation(ctx context.Context, opts ...client.ListOption) (*v1alpha1.IstioInstallationList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIstioInstallation", varargs...)
	ret0, _ := ret[0].(*v1alpha1.IstioInstallationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIstioInstallation indicates an expected call of ListIstioInstallation
func (mr *MockIstioInstallationClientMockRecorder) ListIstioInstallation(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIstioInstallation", reflect.TypeOf((*MockIstioInstallationClient)(nil).ListIstioInstallation), varargs...)
}

// CreateIstioInstallation mocks base method
func (m *MockIstioInstallationClient) CreateIstioInstallation(ctx context.Context, obj *v1alpha1.IstioInstallation, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIstioInstallation indicates an expected call of CreateIstioInstallation
func (mr *MockIstioInstallationClientMockRecorder) CreateIstioInstallation(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIstioInstallation", reflect.TypeOf((*MockIstioInstallationClient)(nil).CreateIstioInstallation), varargs...)
}

// DeleteIstioInstallation mocks base method
func (m *MockIstioInstallationClient) DeleteIstioInstallation(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIstioInstallation indicates an expected call of DeleteIstioInstallation
func (mr *MockIstioInstallationClientMockRecorder) DeleteIstioInstallation(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIstioInstallation", reflect.TypeOf((*MockIstioInstallationClient)(nil).DeleteIstioInstallation), varargs...)
}

// UpdateIstioInstallation mocks base method
func (m *MockIstioInstallationClient) UpdateIstioInstallation(ctx context.Context, obj *v1alpha1.IstioInstallation, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIstioInstallation indicates an expected call of UpdateIstioInstallation
func (mr *MockIstioInstallationClientMockRecorder) UpdateIstioInstallation(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIstioInstallation", reflect.TypeOf((*MockIstioInstallationClient)(nil).UpdateIstioInstallation), varargs...)
}

// PatchIstioInstallation mocks base method
func (m *MockIstioInstallationClient) PatchIstioInstallation(ctx context.Context, obj *v1alpha1.IstioInstallation, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchIstioInstallation indicates an expected call of PatchIstioInstallation
func (mr *MockIstioInstallationClientMockRecorder) PatchIstioInstallation(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchIstioInstallation", reflect.TypeOf((*MockIstioInstallationClient)(nil).PatchIstioInstallation), varargs...)
}

// DeleteAllOfIstioInstallation mocks base method
func (m *MockIstioInstallationClient) DeleteAllOfIstioInstallation(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfIstioInstallation indicates an expected call of DeleteAllOfIstioInstallation
func (mr *MockIstioInstallationClientMockRecorder) DeleteAllOfIstioInstallation(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfIstioInstallation", reflect.TypeOf((*MockIstioInstallationClient)(nil).DeleteAllOfIstioInstallation), varargs...)
}

// UpsertIstioInstallation mocks base method
func (m *MockIstioInstallationClient) UpsertIstioInstallation(ctx context.Context, obj *v1alpha1.IstioInstallation, transitionFuncs ...v1alpha1.IstioInstallationTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertIstioInstallation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertIstioInstallation indicates an expected call of UpsertIstioInstallation
func (mr *MockIstioInstallationClientMockRecorder) UpsertIstioInstallation(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertIstioInstallation", reflect.TypeOf((*MockIstioInstallationClient)(nil).UpsertIstioInstallation), varargs...)
}

// UpdateIstioInstallationStatus mocks base method
func (m *MockIstioInstallationClient) UpdateIstioInstallationStatus(ctx context.Context, obj *v1alpha1.IstioInstallation, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIstioInstallationStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIstioInstallationStatus indicates an expected call of UpdateIstioInstallationStatus
func (mr *MockIstioInstallationClientMockRecorder) UpdateIstioInstallationStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIstioInstallationStatus", reflect.TypeOf((*MockIstioInstallationClient)(nil).UpdateIstioInstallationStatus), varargs...)
}

// PatchIstioInstallationStatus mocks base method
func (m *MockIstioInstallationClient) PatchIstioInstallationStatus(ctx context.Context, obj *v1alpha1.IstioInstallation, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchIstioInstallationStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchIstioInstallationStatus indicates an expected call of PatchIstioInstallationStatus
func (mr *MockIstioInstallationClientMockRecorder) PatchIstioInstallationStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchIstioInstallationStatus", reflect.TypeOf((*MockIstioInstallationClient)(nil).PatchIstioInstallationStatus), varargs...)
}

// MockMulticlusterIstioInstallationClient is a mock of MulticlusterIstioInstallationClient interface
type MockMulticlusterIstioInstallationClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterIstioInstallationClientMockRecorder
}

// MockMulticlusterIstioInstallationClientMockRecorder is the mock recorder for MockMulticlusterIstioInstallationClient
type MockMulticlusterIstioInstallationClientMockRecorder struct {
	mock *MockMulticlusterIstioInstallationClient
}

// NewMockMulticlusterIstioInstallationClient creates a new mock instance
func NewMockMulticlusterIstioInstallationClient(ctrl *gomock.Controller) *MockMulticlusterIstioInstallationClient {
	mock := &MockMulticlusterIstioInstallationClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterIstioInstallationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterIstioInstallationClient) EXPECT() *MockMulticlusterIstioInstallationClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterIstioInstallationClient) Cluster(cluster string) (v1alpha1.IstioInstallationClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha1.IstioInstallationClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterIstioInstallationClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterIstioInstallationClient)(nil).Cluster), cluster)
}
