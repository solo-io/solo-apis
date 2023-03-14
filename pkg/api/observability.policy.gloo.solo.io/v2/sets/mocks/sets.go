// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v2sets is a generated GoMock package.
package mock_v2sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v2 "github.com/solo-io/solo-apis/pkg/api/observability.policy.gloo.solo.io/v2"
	v2sets "github.com/solo-io/solo-apis/pkg/api/observability.policy.gloo.solo.io/v2/sets"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockAccessLogPolicySet is a mock of AccessLogPolicySet interface.
type MockAccessLogPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogPolicySetMockRecorder
}

// MockAccessLogPolicySetMockRecorder is the mock recorder for MockAccessLogPolicySet.
type MockAccessLogPolicySetMockRecorder struct {
	mock *MockAccessLogPolicySet
}

// NewMockAccessLogPolicySet creates a new mock instance.
func NewMockAccessLogPolicySet(ctrl *gomock.Controller) *MockAccessLogPolicySet {
	mock := &MockAccessLogPolicySet{ctrl: ctrl}
	mock.recorder = &MockAccessLogPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogPolicySet) EXPECT() *MockAccessLogPolicySetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockAccessLogPolicySet) Clone() v2sets.AccessLogPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v2sets.AccessLogPolicySet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockAccessLogPolicySetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockAccessLogPolicySet) Delete(accessLogPolicy ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", accessLogPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockAccessLogPolicySetMockRecorder) Delete(accessLogPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Delete), accessLogPolicy)
}

// Delta mocks base method.
func (m *MockAccessLogPolicySet) Delta(newSet v2sets.AccessLogPolicySet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockAccessLogPolicySetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockAccessLogPolicySet) Difference(set v2sets.AccessLogPolicySet) v2sets.AccessLogPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v2sets.AccessLogPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockAccessLogPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockAccessLogPolicySet) Equal(accessLogPolicySet v2sets.AccessLogPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", accessLogPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockAccessLogPolicySetMockRecorder) Equal(accessLogPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Equal), accessLogPolicySet)
}

// Find mocks base method.
func (m *MockAccessLogPolicySet) Find(id ezkube.ResourceId) (*v2.AccessLogPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v2.AccessLogPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockAccessLogPolicySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockAccessLogPolicySet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockAccessLogPolicySetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Generic))
}

// Has mocks base method.
func (m *MockAccessLogPolicySet) Has(accessLogPolicy ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", accessLogPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockAccessLogPolicySetMockRecorder) Has(accessLogPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Has), accessLogPolicy)
}

// Insert mocks base method.
func (m *MockAccessLogPolicySet) Insert(accessLogPolicy ...*v2.AccessLogPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range accessLogPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockAccessLogPolicySetMockRecorder) Insert(accessLogPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Insert), accessLogPolicy...)
}

// Intersection mocks base method.
func (m *MockAccessLogPolicySet) Intersection(set v2sets.AccessLogPolicySet) v2sets.AccessLogPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v2sets.AccessLogPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockAccessLogPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockAccessLogPolicySet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockAccessLogPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Keys))
}

// Length mocks base method.
func (m *MockAccessLogPolicySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockAccessLogPolicySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Length))
}

// List mocks base method.
func (m *MockAccessLogPolicySet) List(filterResource ...func(*v2.AccessLogPolicy) bool) []*v2.AccessLogPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v2.AccessLogPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockAccessLogPolicySetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAccessLogPolicySet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockAccessLogPolicySet) Map() map[string]*v2.AccessLogPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v2.AccessLogPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockAccessLogPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Map))
}

// Union mocks base method.
func (m *MockAccessLogPolicySet) Union(set v2sets.AccessLogPolicySet) v2sets.AccessLogPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v2sets.AccessLogPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockAccessLogPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockAccessLogPolicySet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockAccessLogPolicySet) UnsortedList(filterResource ...func(*v2.AccessLogPolicy) bool) []*v2.AccessLogPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v2.AccessLogPolicy)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockAccessLogPolicySetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockAccessLogPolicySet)(nil).UnsortedList), filterResource...)
}
