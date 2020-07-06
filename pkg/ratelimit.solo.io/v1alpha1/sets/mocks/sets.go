// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha1sets is a generated GoMock package.
package mock_v1alpha1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/ratelimit.solo.io/v1alpha1"
	v1alpha1sets "github.com/solo-io/solo-apis/pkg/ratelimit.solo.io/v1alpha1/sets"
	sets "k8s.io/apimachinery/pkg/util/sets"
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

// Keys mocks base method.
func (m *MockRateLimitConfigSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockRateLimitConfigSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Keys))
}

// List mocks base method.
func (m *MockRateLimitConfigSet) List() []*v1alpha1.RateLimitConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.RateLimitConfig)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRateLimitConfigSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRateLimitConfigSet)(nil).List))
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

// Has mocks base method.
func (m *MockRateLimitConfigSet) Has(rateLimitConfig *v1alpha1.RateLimitConfig) bool {
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

// Delete mocks base method.
func (m *MockRateLimitConfigSet) Delete(rateLimitConfig *v1alpha1.RateLimitConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", rateLimitConfig)
}

// Delete indicates an expected call of Delete.
func (mr *MockRateLimitConfigSetMockRecorder) Delete(rateLimitConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRateLimitConfigSet)(nil).Delete), rateLimitConfig)
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
