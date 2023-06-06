// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v2sets is a generated GoMock package.
package mock_v2sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets0 "k8s.io/apimachinery/pkg/util/sets"

	v2 "github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2"
	v2sets "github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2/sets"
)

// MockWasmDeploymentPolicySet is a mock of WasmDeploymentPolicySet interface.
type MockWasmDeploymentPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockWasmDeploymentPolicySetMockRecorder
}

// MockWasmDeploymentPolicySetMockRecorder is the mock recorder for MockWasmDeploymentPolicySet.
type MockWasmDeploymentPolicySetMockRecorder struct {
	mock *MockWasmDeploymentPolicySet
}

// NewMockWasmDeploymentPolicySet creates a new mock instance.
func NewMockWasmDeploymentPolicySet(ctrl *gomock.Controller) *MockWasmDeploymentPolicySet {
	mock := &MockWasmDeploymentPolicySet{ctrl: ctrl}
	mock.recorder = &MockWasmDeploymentPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmDeploymentPolicySet) EXPECT() *MockWasmDeploymentPolicySetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockWasmDeploymentPolicySet) Clone() v2sets.WasmDeploymentPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v2sets.WasmDeploymentPolicySet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockWasmDeploymentPolicySet) Delete(wasmDeploymentPolicy ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", wasmDeploymentPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Delete(wasmDeploymentPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Delete), wasmDeploymentPolicy)
}

// Delta mocks base method.
func (m *MockWasmDeploymentPolicySet) Delta(newSet v2sets.WasmDeploymentPolicySet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockWasmDeploymentPolicySet) Difference(set v2sets.WasmDeploymentPolicySet) v2sets.WasmDeploymentPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v2sets.WasmDeploymentPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockWasmDeploymentPolicySet) Equal(wasmDeploymentPolicySet v2sets.WasmDeploymentPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", wasmDeploymentPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Equal(wasmDeploymentPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Equal), wasmDeploymentPolicySet)
}

// Find mocks base method.
func (m *MockWasmDeploymentPolicySet) Find(id ezkube.ResourceId) (*v2.WasmDeploymentPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v2.WasmDeploymentPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockWasmDeploymentPolicySet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Generic))
}

// Has mocks base method.
func (m *MockWasmDeploymentPolicySet) Has(wasmDeploymentPolicy ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", wasmDeploymentPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Has(wasmDeploymentPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Has), wasmDeploymentPolicy)
}

// Insert mocks base method.
func (m *MockWasmDeploymentPolicySet) Insert(wasmDeploymentPolicy ...*v2.WasmDeploymentPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range wasmDeploymentPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Insert(wasmDeploymentPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Insert), wasmDeploymentPolicy...)
}

// Intersection mocks base method.
func (m *MockWasmDeploymentPolicySet) Intersection(set v2sets.WasmDeploymentPolicySet) v2sets.WasmDeploymentPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v2sets.WasmDeploymentPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockWasmDeploymentPolicySet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Keys))
}

// Length mocks base method.
func (m *MockWasmDeploymentPolicySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Length))
}

// List mocks base method.
func (m *MockWasmDeploymentPolicySet) List(filterResource ...func(*v2.WasmDeploymentPolicy) bool) []*v2.WasmDeploymentPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v2.WasmDeploymentPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockWasmDeploymentPolicySetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockWasmDeploymentPolicySet) Map() map[string]*v2.WasmDeploymentPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v2.WasmDeploymentPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Map))
}

// Union mocks base method.
func (m *MockWasmDeploymentPolicySet) Union(set v2sets.WasmDeploymentPolicySet) v2sets.WasmDeploymentPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v2sets.WasmDeploymentPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockWasmDeploymentPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockWasmDeploymentPolicySet) UnsortedList(filterResource ...func(*v2.WasmDeploymentPolicy) bool) []*v2.WasmDeploymentPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v2.WasmDeploymentPolicy)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockWasmDeploymentPolicySetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockWasmDeploymentPolicySet)(nil).UnsortedList), filterResource...)
}
