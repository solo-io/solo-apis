// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v2sets is a generated GoMock package.
package mock_v2sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v2 "github.com/solo-io/solo-apis/pkg/api/resilience.policy.gloo.solo.io/v2"
	v2sets "github.com/solo-io/solo-apis/pkg/api/resilience.policy.gloo.solo.io/v2/sets"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockFailoverPolicySet is a mock of FailoverPolicySet interface.
type MockFailoverPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockFailoverPolicySetMockRecorder
}

// MockFailoverPolicySetMockRecorder is the mock recorder for MockFailoverPolicySet.
type MockFailoverPolicySetMockRecorder struct {
	mock *MockFailoverPolicySet
}

// NewMockFailoverPolicySet creates a new mock instance.
func NewMockFailoverPolicySet(ctrl *gomock.Controller) *MockFailoverPolicySet {
	mock := &MockFailoverPolicySet{ctrl: ctrl}
	mock.recorder = &MockFailoverPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFailoverPolicySet) EXPECT() *MockFailoverPolicySetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockFailoverPolicySet) Clone() v2sets.FailoverPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v2sets.FailoverPolicySet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockFailoverPolicySetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockFailoverPolicySet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockFailoverPolicySet) Delete(failoverPolicy ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", failoverPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockFailoverPolicySetMockRecorder) Delete(failoverPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFailoverPolicySet)(nil).Delete), failoverPolicy)
}

// Delta mocks base method.
func (m *MockFailoverPolicySet) Delta(newSet v2sets.FailoverPolicySet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockFailoverPolicySetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockFailoverPolicySet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockFailoverPolicySet) Difference(set v2sets.FailoverPolicySet) v2sets.FailoverPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v2sets.FailoverPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockFailoverPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockFailoverPolicySet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockFailoverPolicySet) Equal(failoverPolicySet v2sets.FailoverPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", failoverPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockFailoverPolicySetMockRecorder) Equal(failoverPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockFailoverPolicySet)(nil).Equal), failoverPolicySet)
}

// Find mocks base method.
func (m *MockFailoverPolicySet) Find(id ezkube.ResourceId) (*v2.FailoverPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v2.FailoverPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockFailoverPolicySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFailoverPolicySet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockFailoverPolicySet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockFailoverPolicySetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockFailoverPolicySet)(nil).Generic))
}

// Has mocks base method.
func (m *MockFailoverPolicySet) Has(failoverPolicy ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", failoverPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockFailoverPolicySetMockRecorder) Has(failoverPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockFailoverPolicySet)(nil).Has), failoverPolicy)
}

// Insert mocks base method.
func (m *MockFailoverPolicySet) Insert(failoverPolicy ...*v2.FailoverPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range failoverPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockFailoverPolicySetMockRecorder) Insert(failoverPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockFailoverPolicySet)(nil).Insert), failoverPolicy...)
}

// Intersection mocks base method.
func (m *MockFailoverPolicySet) Intersection(set v2sets.FailoverPolicySet) v2sets.FailoverPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v2sets.FailoverPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockFailoverPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockFailoverPolicySet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockFailoverPolicySet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockFailoverPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockFailoverPolicySet)(nil).Keys))
}

// Length mocks base method.
func (m *MockFailoverPolicySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockFailoverPolicySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockFailoverPolicySet)(nil).Length))
}

// List mocks base method.
func (m *MockFailoverPolicySet) List(filterResource ...func(*v2.FailoverPolicy) bool) []*v2.FailoverPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v2.FailoverPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockFailoverPolicySetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFailoverPolicySet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockFailoverPolicySet) Map() map[string]*v2.FailoverPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v2.FailoverPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockFailoverPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockFailoverPolicySet)(nil).Map))
}

// Union mocks base method.
func (m *MockFailoverPolicySet) Union(set v2sets.FailoverPolicySet) v2sets.FailoverPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v2sets.FailoverPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockFailoverPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockFailoverPolicySet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockFailoverPolicySet) UnsortedList(filterResource ...func(*v2.FailoverPolicy) bool) []*v2.FailoverPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v2.FailoverPolicy)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockFailoverPolicySetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockFailoverPolicySet)(nil).UnsortedList), filterResource...)
}

// MockOutlierDetectionPolicySet is a mock of OutlierDetectionPolicySet interface.
type MockOutlierDetectionPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockOutlierDetectionPolicySetMockRecorder
}

// MockOutlierDetectionPolicySetMockRecorder is the mock recorder for MockOutlierDetectionPolicySet.
type MockOutlierDetectionPolicySetMockRecorder struct {
	mock *MockOutlierDetectionPolicySet
}

// NewMockOutlierDetectionPolicySet creates a new mock instance.
func NewMockOutlierDetectionPolicySet(ctrl *gomock.Controller) *MockOutlierDetectionPolicySet {
	mock := &MockOutlierDetectionPolicySet{ctrl: ctrl}
	mock.recorder = &MockOutlierDetectionPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutlierDetectionPolicySet) EXPECT() *MockOutlierDetectionPolicySetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockOutlierDetectionPolicySet) Clone() v2sets.OutlierDetectionPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v2sets.OutlierDetectionPolicySet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockOutlierDetectionPolicySet) Delete(outlierDetectionPolicy ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", outlierDetectionPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Delete(outlierDetectionPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Delete), outlierDetectionPolicy)
}

// Delta mocks base method.
func (m *MockOutlierDetectionPolicySet) Delta(newSet v2sets.OutlierDetectionPolicySet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockOutlierDetectionPolicySet) Difference(set v2sets.OutlierDetectionPolicySet) v2sets.OutlierDetectionPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v2sets.OutlierDetectionPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockOutlierDetectionPolicySet) Equal(outlierDetectionPolicySet v2sets.OutlierDetectionPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", outlierDetectionPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Equal(outlierDetectionPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Equal), outlierDetectionPolicySet)
}

// Find mocks base method.
func (m *MockOutlierDetectionPolicySet) Find(id ezkube.ResourceId) (*v2.OutlierDetectionPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v2.OutlierDetectionPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockOutlierDetectionPolicySet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Generic))
}

// Has mocks base method.
func (m *MockOutlierDetectionPolicySet) Has(outlierDetectionPolicy ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", outlierDetectionPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Has(outlierDetectionPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Has), outlierDetectionPolicy)
}

// Insert mocks base method.
func (m *MockOutlierDetectionPolicySet) Insert(outlierDetectionPolicy ...*v2.OutlierDetectionPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range outlierDetectionPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Insert(outlierDetectionPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Insert), outlierDetectionPolicy...)
}

// Intersection mocks base method.
func (m *MockOutlierDetectionPolicySet) Intersection(set v2sets.OutlierDetectionPolicySet) v2sets.OutlierDetectionPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v2sets.OutlierDetectionPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockOutlierDetectionPolicySet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Keys))
}

// Length mocks base method.
func (m *MockOutlierDetectionPolicySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Length))
}

// List mocks base method.
func (m *MockOutlierDetectionPolicySet) List(filterResource ...func(*v2.OutlierDetectionPolicy) bool) []*v2.OutlierDetectionPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v2.OutlierDetectionPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockOutlierDetectionPolicySetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockOutlierDetectionPolicySet) Map() map[string]*v2.OutlierDetectionPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v2.OutlierDetectionPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Map))
}

// Union mocks base method.
func (m *MockOutlierDetectionPolicySet) Union(set v2sets.OutlierDetectionPolicySet) v2sets.OutlierDetectionPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v2sets.OutlierDetectionPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockOutlierDetectionPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockOutlierDetectionPolicySet) UnsortedList(filterResource ...func(*v2.OutlierDetectionPolicy) bool) []*v2.OutlierDetectionPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v2.OutlierDetectionPolicy)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockOutlierDetectionPolicySetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockOutlierDetectionPolicySet)(nil).UnsortedList), filterResource...)
}

// MockFaultInjectionPolicySet is a mock of FaultInjectionPolicySet interface.
type MockFaultInjectionPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockFaultInjectionPolicySetMockRecorder
}

// MockFaultInjectionPolicySetMockRecorder is the mock recorder for MockFaultInjectionPolicySet.
type MockFaultInjectionPolicySetMockRecorder struct {
	mock *MockFaultInjectionPolicySet
}

// NewMockFaultInjectionPolicySet creates a new mock instance.
func NewMockFaultInjectionPolicySet(ctrl *gomock.Controller) *MockFaultInjectionPolicySet {
	mock := &MockFaultInjectionPolicySet{ctrl: ctrl}
	mock.recorder = &MockFaultInjectionPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFaultInjectionPolicySet) EXPECT() *MockFaultInjectionPolicySetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockFaultInjectionPolicySet) Clone() v2sets.FaultInjectionPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v2sets.FaultInjectionPolicySet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockFaultInjectionPolicySetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockFaultInjectionPolicySet) Delete(faultInjectionPolicy ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", faultInjectionPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockFaultInjectionPolicySetMockRecorder) Delete(faultInjectionPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Delete), faultInjectionPolicy)
}

// Delta mocks base method.
func (m *MockFaultInjectionPolicySet) Delta(newSet v2sets.FaultInjectionPolicySet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockFaultInjectionPolicySetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockFaultInjectionPolicySet) Difference(set v2sets.FaultInjectionPolicySet) v2sets.FaultInjectionPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v2sets.FaultInjectionPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockFaultInjectionPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockFaultInjectionPolicySet) Equal(faultInjectionPolicySet v2sets.FaultInjectionPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", faultInjectionPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockFaultInjectionPolicySetMockRecorder) Equal(faultInjectionPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Equal), faultInjectionPolicySet)
}

// Find mocks base method.
func (m *MockFaultInjectionPolicySet) Find(id ezkube.ResourceId) (*v2.FaultInjectionPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v2.FaultInjectionPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockFaultInjectionPolicySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockFaultInjectionPolicySet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockFaultInjectionPolicySetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Generic))
}

// Has mocks base method.
func (m *MockFaultInjectionPolicySet) Has(faultInjectionPolicy ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", faultInjectionPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockFaultInjectionPolicySetMockRecorder) Has(faultInjectionPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Has), faultInjectionPolicy)
}

// Insert mocks base method.
func (m *MockFaultInjectionPolicySet) Insert(faultInjectionPolicy ...*v2.FaultInjectionPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range faultInjectionPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockFaultInjectionPolicySetMockRecorder) Insert(faultInjectionPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Insert), faultInjectionPolicy...)
}

// Intersection mocks base method.
func (m *MockFaultInjectionPolicySet) Intersection(set v2sets.FaultInjectionPolicySet) v2sets.FaultInjectionPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v2sets.FaultInjectionPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockFaultInjectionPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockFaultInjectionPolicySet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockFaultInjectionPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Keys))
}

// Length mocks base method.
func (m *MockFaultInjectionPolicySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockFaultInjectionPolicySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Length))
}

// List mocks base method.
func (m *MockFaultInjectionPolicySet) List(filterResource ...func(*v2.FaultInjectionPolicy) bool) []*v2.FaultInjectionPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v2.FaultInjectionPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockFaultInjectionPolicySetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockFaultInjectionPolicySet) Map() map[string]*v2.FaultInjectionPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v2.FaultInjectionPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockFaultInjectionPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Map))
}

// Union mocks base method.
func (m *MockFaultInjectionPolicySet) Union(set v2sets.FaultInjectionPolicySet) v2sets.FaultInjectionPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v2sets.FaultInjectionPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockFaultInjectionPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockFaultInjectionPolicySet) UnsortedList(filterResource ...func(*v2.FaultInjectionPolicy) bool) []*v2.FaultInjectionPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v2.FaultInjectionPolicy)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockFaultInjectionPolicySetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockFaultInjectionPolicySet)(nil).UnsortedList), filterResource...)
}

// MockRetryTimeoutPolicySet is a mock of RetryTimeoutPolicySet interface.
type MockRetryTimeoutPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockRetryTimeoutPolicySetMockRecorder
}

// MockRetryTimeoutPolicySetMockRecorder is the mock recorder for MockRetryTimeoutPolicySet.
type MockRetryTimeoutPolicySetMockRecorder struct {
	mock *MockRetryTimeoutPolicySet
}

// NewMockRetryTimeoutPolicySet creates a new mock instance.
func NewMockRetryTimeoutPolicySet(ctrl *gomock.Controller) *MockRetryTimeoutPolicySet {
	mock := &MockRetryTimeoutPolicySet{ctrl: ctrl}
	mock.recorder = &MockRetryTimeoutPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRetryTimeoutPolicySet) EXPECT() *MockRetryTimeoutPolicySetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockRetryTimeoutPolicySet) Clone() v2sets.RetryTimeoutPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v2sets.RetryTimeoutPolicySet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockRetryTimeoutPolicySet) Delete(retryTimeoutPolicy ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", retryTimeoutPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Delete(retryTimeoutPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Delete), retryTimeoutPolicy)
}

// Delta mocks base method.
func (m *MockRetryTimeoutPolicySet) Delta(newSet v2sets.RetryTimeoutPolicySet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockRetryTimeoutPolicySet) Difference(set v2sets.RetryTimeoutPolicySet) v2sets.RetryTimeoutPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v2sets.RetryTimeoutPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockRetryTimeoutPolicySet) Equal(retryTimeoutPolicySet v2sets.RetryTimeoutPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", retryTimeoutPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Equal(retryTimeoutPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Equal), retryTimeoutPolicySet)
}

// Find mocks base method.
func (m *MockRetryTimeoutPolicySet) Find(id ezkube.ResourceId) (*v2.RetryTimeoutPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v2.RetryTimeoutPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockRetryTimeoutPolicySet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Generic))
}

// Has mocks base method.
func (m *MockRetryTimeoutPolicySet) Has(retryTimeoutPolicy ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", retryTimeoutPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Has(retryTimeoutPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Has), retryTimeoutPolicy)
}

// Insert mocks base method.
func (m *MockRetryTimeoutPolicySet) Insert(retryTimeoutPolicy ...*v2.RetryTimeoutPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range retryTimeoutPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Insert(retryTimeoutPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Insert), retryTimeoutPolicy...)
}

// Intersection mocks base method.
func (m *MockRetryTimeoutPolicySet) Intersection(set v2sets.RetryTimeoutPolicySet) v2sets.RetryTimeoutPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v2sets.RetryTimeoutPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockRetryTimeoutPolicySet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Keys))
}

// Length mocks base method.
func (m *MockRetryTimeoutPolicySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Length))
}

// List mocks base method.
func (m *MockRetryTimeoutPolicySet) List(filterResource ...func(*v2.RetryTimeoutPolicy) bool) []*v2.RetryTimeoutPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v2.RetryTimeoutPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRetryTimeoutPolicySetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockRetryTimeoutPolicySet) Map() map[string]*v2.RetryTimeoutPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v2.RetryTimeoutPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Map))
}

// Union mocks base method.
func (m *MockRetryTimeoutPolicySet) Union(set v2sets.RetryTimeoutPolicySet) v2sets.RetryTimeoutPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v2sets.RetryTimeoutPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockRetryTimeoutPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockRetryTimeoutPolicySet) UnsortedList(filterResource ...func(*v2.RetryTimeoutPolicy) bool) []*v2.RetryTimeoutPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v2.RetryTimeoutPolicy)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockRetryTimeoutPolicySetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockRetryTimeoutPolicySet)(nil).UnsortedList), filterResource...)
}
