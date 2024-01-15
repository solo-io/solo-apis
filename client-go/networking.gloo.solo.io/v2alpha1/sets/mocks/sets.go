// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v2alpha1sets is a generated GoMock package.
package mock_v2alpha1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets0 "k8s.io/apimachinery/pkg/util/sets"

	v2alpha1 "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2alpha1"
	v2alpha1sets "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2alpha1/sets"
)

// MockExternalWorkloadSet is a mock of ExternalWorkloadSet interface.
type MockExternalWorkloadSet struct {
	ctrl     *gomock.Controller
	recorder *MockExternalWorkloadSetMockRecorder
}

// MockExternalWorkloadSetMockRecorder is the mock recorder for MockExternalWorkloadSet.
type MockExternalWorkloadSetMockRecorder struct {
	mock *MockExternalWorkloadSet
}

// NewMockExternalWorkloadSet creates a new mock instance.
func NewMockExternalWorkloadSet(ctrl *gomock.Controller) *MockExternalWorkloadSet {
	mock := &MockExternalWorkloadSet{ctrl: ctrl}
	mock.recorder = &MockExternalWorkloadSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalWorkloadSet) EXPECT() *MockExternalWorkloadSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockExternalWorkloadSet) Clone() v2alpha1sets.ExternalWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v2alpha1sets.ExternalWorkloadSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockExternalWorkloadSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockExternalWorkloadSet) Delete(externalWorkload ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", externalWorkload)
}

// Delete indicates an expected call of Delete.
func (mr *MockExternalWorkloadSetMockRecorder) Delete(externalWorkload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Delete), externalWorkload)
}

// Delta mocks base method.
func (m *MockExternalWorkloadSet) Delta(newSet v2alpha1sets.ExternalWorkloadSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockExternalWorkloadSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockExternalWorkloadSet) Difference(set v2alpha1sets.ExternalWorkloadSet) v2alpha1sets.ExternalWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v2alpha1sets.ExternalWorkloadSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockExternalWorkloadSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockExternalWorkloadSet) Equal(externalWorkloadSet v2alpha1sets.ExternalWorkloadSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", externalWorkloadSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockExternalWorkloadSetMockRecorder) Equal(externalWorkloadSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Equal), externalWorkloadSet)
}

// Find mocks base method.
func (m *MockExternalWorkloadSet) Find(id ezkube.ResourceId) (*v2alpha1.ExternalWorkload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v2alpha1.ExternalWorkload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockExternalWorkloadSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockExternalWorkloadSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockExternalWorkloadSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockExternalWorkloadSet) Has(externalWorkload ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", externalWorkload)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockExternalWorkloadSetMockRecorder) Has(externalWorkload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Has), externalWorkload)
}

// Insert mocks base method.
func (m *MockExternalWorkloadSet) Insert(externalWorkload ...*v2alpha1.ExternalWorkload) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range externalWorkload {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockExternalWorkloadSetMockRecorder) Insert(externalWorkload ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Insert), externalWorkload...)
}

// Intersection mocks base method.
func (m *MockExternalWorkloadSet) Intersection(set v2alpha1sets.ExternalWorkloadSet) v2alpha1sets.ExternalWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v2alpha1sets.ExternalWorkloadSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockExternalWorkloadSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockExternalWorkloadSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockExternalWorkloadSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockExternalWorkloadSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockExternalWorkloadSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Length))
}

// List mocks base method.
func (m *MockExternalWorkloadSet) List(filterResource ...func(*v2alpha1.ExternalWorkload) bool) []*v2alpha1.ExternalWorkload {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v2alpha1.ExternalWorkload)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockExternalWorkloadSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockExternalWorkloadSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockExternalWorkloadSet) Map() map[string]*v2alpha1.ExternalWorkload {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v2alpha1.ExternalWorkload)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockExternalWorkloadSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Map))
}

// Union mocks base method.
func (m *MockExternalWorkloadSet) Union(set v2alpha1sets.ExternalWorkloadSet) v2alpha1sets.ExternalWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v2alpha1sets.ExternalWorkloadSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockExternalWorkloadSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockExternalWorkloadSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockExternalWorkloadSet) UnsortedList(filterResource ...func(*v2alpha1.ExternalWorkload) bool) []*v2alpha1.ExternalWorkload {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v2alpha1.ExternalWorkload)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockExternalWorkloadSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockExternalWorkloadSet)(nil).UnsortedList), filterResource...)
}
