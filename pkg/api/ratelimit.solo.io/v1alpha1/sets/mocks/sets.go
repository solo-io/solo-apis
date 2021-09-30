// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha1sets is a generated GoMock package.
package mock_v1alpha1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
	v1alpha1sets "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1/sets"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockRateLimitConfigSet is a mock of RateLimitConfigSet interface.
type MockRateLimitConfigSet struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimitConfigSetMockRecorder
}

// MockRateLimitConfigSetMockRecorder is the mock recorder for MockRateLimitConfigSet.
type MockRateLimitConfigSetMockRecorder struct {
	mock *MockRateLimitConfigSet
}

// NewMockRateLimitConfigSet creates a new mock instance.
func NewMockRateLimitConfigSet(ctrl *gomock.Controller) *MockRateLimitConfigSet {
	mock := &MockRateLimitConfigSet{ctrl: ctrl}
	mock.recorder = &MockRateLimitConfigSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateLimitConfigSet) EXPECT() *MockRateLimitConfigSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockRateLimitConfigSet) Clone() v1alpha1sets.RateLimitConfigSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1alpha1sets.RateLimitConfigSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockRateLimitConfigSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockRateLimitConfigSet) Delete(rateLimitConfig ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", rateLimitConfig)
}

// Delete indicates an expected call of Delete.
func (mr *MockRateLimitConfigSetMockRecorder) Delete(rateLimitConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Delete), rateLimitConfig)
}

// Delta mocks base method.
func (m *MockRateLimitConfigSet) Delta(newSet v1alpha1sets.RateLimitConfigSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockRateLimitConfigSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockRateLimitConfigSet) Difference(set v1alpha1sets.RateLimitConfigSet) v1alpha1sets.RateLimitConfigSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.RateLimitConfigSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockRateLimitConfigSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockRateLimitConfigSet) Equal(rateLimitConfigSet v1alpha1sets.RateLimitConfigSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", rateLimitConfigSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockRateLimitConfigSetMockRecorder) Equal(rateLimitConfigSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Equal), rateLimitConfigSet)
}

// Find mocks base method.
func (m *MockRateLimitConfigSet) Find(id ezkube.ResourceId) (*v1alpha1.RateLimitConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.RateLimitConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRateLimitConfigSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockRateLimitConfigSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockRateLimitConfigSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockRateLimitConfigSet) Has(rateLimitConfig ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", rateLimitConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockRateLimitConfigSetMockRecorder) Has(rateLimitConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Has), rateLimitConfig)
}

// Insert mocks base method.
func (m *MockRateLimitConfigSet) Insert(rateLimitConfig ...*v1alpha1.RateLimitConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range rateLimitConfig {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockRateLimitConfigSetMockRecorder) Insert(rateLimitConfig ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Insert), rateLimitConfig...)
}

// Intersection mocks base method.
func (m *MockRateLimitConfigSet) Intersection(set v1alpha1sets.RateLimitConfigSet) v1alpha1sets.RateLimitConfigSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.RateLimitConfigSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockRateLimitConfigSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockRateLimitConfigSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockRateLimitConfigSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockRateLimitConfigSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockRateLimitConfigSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Length))
}

// List mocks base method.
func (m *MockRateLimitConfigSet) List(filterResource ...func(*v1alpha1.RateLimitConfig) bool) []*v1alpha1.RateLimitConfig {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1alpha1.RateLimitConfig)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRateLimitConfigSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRateLimitConfigSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockRateLimitConfigSet) Map() map[string]*v1alpha1.RateLimitConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.RateLimitConfig)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockRateLimitConfigSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Map))
}

// Union mocks base method.
func (m *MockRateLimitConfigSet) Union(set v1alpha1sets.RateLimitConfigSet) v1alpha1sets.RateLimitConfigSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.RateLimitConfigSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockRateLimitConfigSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockRateLimitConfigSet) UnsortedList(filterResource ...func(*v1alpha1.RateLimitConfig) bool) []*v1alpha1.RateLimitConfig {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1alpha1.RateLimitConfig)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockRateLimitConfigSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockRateLimitConfigSet)(nil).UnsortedList), filterResource...)
}
