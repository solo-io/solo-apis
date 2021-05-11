// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	v1sets "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1/sets"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockGatewaySet is a mock of GatewaySet interface
type MockGatewaySet struct {
	ctrl     *gomock.Controller
	recorder *MockGatewaySetMockRecorder
}

// MockGatewaySetMockRecorder is the mock recorder for MockGatewaySet
type MockGatewaySetMockRecorder struct {
	mock *MockGatewaySet
}

// NewMockGatewaySet creates a new mock instance
func NewMockGatewaySet(ctrl *gomock.Controller) *MockGatewaySet {
	mock := &MockGatewaySet{ctrl: ctrl}
	mock.recorder = &MockGatewaySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGatewaySet) EXPECT() *MockGatewaySetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockGatewaySet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockGatewaySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockGatewaySet)(nil).Keys))
}

// List mocks base method
func (m *MockGatewaySet) List(filterResource ...func(*v1.Gateway) bool) []*v1.Gateway {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.Gateway)
	return ret0
}

// List indicates an expected call of List
func (mr *MockGatewaySetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGatewaySet)(nil).List), filterResource...)
}

// UnsortedList mocks base method
func (m *MockGatewaySet) UnsortedList(filterResource ...func(*v1.Gateway) bool) []*v1.Gateway {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.Gateway)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList
func (mr *MockGatewaySetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockGatewaySet)(nil).UnsortedList), filterResource...)
}

// Map mocks base method
func (m *MockGatewaySet) Map() map[string]*v1.Gateway {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Gateway)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockGatewaySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockGatewaySet)(nil).Map))
}

// Insert mocks base method
func (m *MockGatewaySet) Insert(gateway ...*v1.Gateway) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range gateway {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockGatewaySetMockRecorder) Insert(gateway ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockGatewaySet)(nil).Insert), gateway...)
}

// Equal mocks base method
func (m *MockGatewaySet) Equal(gatewaySet v1sets.GatewaySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", gatewaySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockGatewaySetMockRecorder) Equal(gatewaySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockGatewaySet)(nil).Equal), gatewaySet)
}

// Has mocks base method
func (m *MockGatewaySet) Has(gateway ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", gateway)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockGatewaySetMockRecorder) Has(gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockGatewaySet)(nil).Has), gateway)
}

// Delete mocks base method
func (m *MockGatewaySet) Delete(gateway ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", gateway)
}

// Delete indicates an expected call of Delete
func (mr *MockGatewaySetMockRecorder) Delete(gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGatewaySet)(nil).Delete), gateway)
}

// Union mocks base method
func (m *MockGatewaySet) Union(set v1sets.GatewaySet) v1sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.GatewaySet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockGatewaySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockGatewaySet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockGatewaySet) Difference(set v1sets.GatewaySet) v1sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.GatewaySet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockGatewaySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockGatewaySet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockGatewaySet) Intersection(set v1sets.GatewaySet) v1sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.GatewaySet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockGatewaySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockGatewaySet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockGatewaySet) Find(id ezkube.ResourceId) (*v1.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockGatewaySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockGatewaySet)(nil).Find), id)
}

// Length mocks base method
func (m *MockGatewaySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockGatewaySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockGatewaySet)(nil).Length))
}

// Generic mocks base method
func (m *MockGatewaySet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockGatewaySetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockGatewaySet)(nil).Generic))
}

// Delta mocks base method
func (m *MockGatewaySet) Delta(newSet v1sets.GatewaySet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockGatewaySetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockGatewaySet)(nil).Delta), newSet)
}

// MockRouteTableSet is a mock of RouteTableSet interface
type MockRouteTableSet struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTableSetMockRecorder
}

// MockRouteTableSetMockRecorder is the mock recorder for MockRouteTableSet
type MockRouteTableSetMockRecorder struct {
	mock *MockRouteTableSet
}

// NewMockRouteTableSet creates a new mock instance
func NewMockRouteTableSet(ctrl *gomock.Controller) *MockRouteTableSet {
	mock := &MockRouteTableSet{ctrl: ctrl}
	mock.recorder = &MockRouteTableSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteTableSet) EXPECT() *MockRouteTableSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockRouteTableSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockRouteTableSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRouteTableSet)(nil).Keys))
}

// List mocks base method
func (m *MockRouteTableSet) List(filterResource ...func(*v1.RouteTable) bool) []*v1.RouteTable {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.RouteTable)
	return ret0
}

// List indicates an expected call of List
func (mr *MockRouteTableSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRouteTableSet)(nil).List), filterResource...)
}

// UnsortedList mocks base method
func (m *MockRouteTableSet) UnsortedList(filterResource ...func(*v1.RouteTable) bool) []*v1.RouteTable {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.RouteTable)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList
func (mr *MockRouteTableSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockRouteTableSet)(nil).UnsortedList), filterResource...)
}

// Map mocks base method
func (m *MockRouteTableSet) Map() map[string]*v1.RouteTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.RouteTable)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockRouteTableSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRouteTableSet)(nil).Map))
}

// Insert mocks base method
func (m *MockRouteTableSet) Insert(routeTable ...*v1.RouteTable) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range routeTable {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockRouteTableSetMockRecorder) Insert(routeTable ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRouteTableSet)(nil).Insert), routeTable...)
}

// Equal mocks base method
func (m *MockRouteTableSet) Equal(routeTableSet v1sets.RouteTableSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", routeTableSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockRouteTableSetMockRecorder) Equal(routeTableSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockRouteTableSet)(nil).Equal), routeTableSet)
}

// Has mocks base method
func (m *MockRouteTableSet) Has(routeTable ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", routeTable)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockRouteTableSetMockRecorder) Has(routeTable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockRouteTableSet)(nil).Has), routeTable)
}

// Delete mocks base method
func (m *MockRouteTableSet) Delete(routeTable ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", routeTable)
}

// Delete indicates an expected call of Delete
func (mr *MockRouteTableSetMockRecorder) Delete(routeTable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRouteTableSet)(nil).Delete), routeTable)
}

// Union mocks base method
func (m *MockRouteTableSet) Union(set v1sets.RouteTableSet) v1sets.RouteTableSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.RouteTableSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockRouteTableSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockRouteTableSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockRouteTableSet) Difference(set v1sets.RouteTableSet) v1sets.RouteTableSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.RouteTableSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockRouteTableSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockRouteTableSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockRouteTableSet) Intersection(set v1sets.RouteTableSet) v1sets.RouteTableSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.RouteTableSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockRouteTableSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockRouteTableSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockRouteTableSet) Find(id ezkube.ResourceId) (*v1.RouteTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.RouteTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockRouteTableSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRouteTableSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockRouteTableSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockRouteTableSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockRouteTableSet)(nil).Length))
}

// Generic mocks base method
func (m *MockRouteTableSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockRouteTableSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockRouteTableSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockRouteTableSet) Delta(newSet v1sets.RouteTableSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockRouteTableSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockRouteTableSet)(nil).Delta), newSet)
}

// MockVirtualServiceSet is a mock of VirtualServiceSet interface
type MockVirtualServiceSet struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceSetMockRecorder
}

// MockVirtualServiceSetMockRecorder is the mock recorder for MockVirtualServiceSet
type MockVirtualServiceSetMockRecorder struct {
	mock *MockVirtualServiceSet
}

// NewMockVirtualServiceSet creates a new mock instance
func NewMockVirtualServiceSet(ctrl *gomock.Controller) *MockVirtualServiceSet {
	mock := &MockVirtualServiceSet{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualServiceSet) EXPECT() *MockVirtualServiceSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockVirtualServiceSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockVirtualServiceSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockVirtualServiceSet)(nil).Keys))
}

// List mocks base method
func (m *MockVirtualServiceSet) List(filterResource ...func(*v1.VirtualService) bool) []*v1.VirtualService {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.VirtualService)
	return ret0
}

// List indicates an expected call of List
func (mr *MockVirtualServiceSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualServiceSet)(nil).List), filterResource...)
}

// UnsortedList mocks base method
func (m *MockVirtualServiceSet) UnsortedList(filterResource ...func(*v1.VirtualService) bool) []*v1.VirtualService {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.VirtualService)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList
func (mr *MockVirtualServiceSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockVirtualServiceSet)(nil).UnsortedList), filterResource...)
}

// Map mocks base method
func (m *MockVirtualServiceSet) Map() map[string]*v1.VirtualService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.VirtualService)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockVirtualServiceSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockVirtualServiceSet)(nil).Map))
}

// Insert mocks base method
func (m *MockVirtualServiceSet) Insert(virtualService ...*v1.VirtualService) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range virtualService {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockVirtualServiceSetMockRecorder) Insert(virtualService ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockVirtualServiceSet)(nil).Insert), virtualService...)
}

// Equal mocks base method
func (m *MockVirtualServiceSet) Equal(virtualServiceSet v1sets.VirtualServiceSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", virtualServiceSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockVirtualServiceSetMockRecorder) Equal(virtualServiceSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockVirtualServiceSet)(nil).Equal), virtualServiceSet)
}

// Has mocks base method
func (m *MockVirtualServiceSet) Has(virtualService ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", virtualService)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockVirtualServiceSetMockRecorder) Has(virtualService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockVirtualServiceSet)(nil).Has), virtualService)
}

// Delete mocks base method
func (m *MockVirtualServiceSet) Delete(virtualService ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", virtualService)
}

// Delete indicates an expected call of Delete
func (mr *MockVirtualServiceSetMockRecorder) Delete(virtualService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualServiceSet)(nil).Delete), virtualService)
}

// Union mocks base method
func (m *MockVirtualServiceSet) Union(set v1sets.VirtualServiceSet) v1sets.VirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.VirtualServiceSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockVirtualServiceSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockVirtualServiceSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockVirtualServiceSet) Difference(set v1sets.VirtualServiceSet) v1sets.VirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.VirtualServiceSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockVirtualServiceSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockVirtualServiceSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockVirtualServiceSet) Intersection(set v1sets.VirtualServiceSet) v1sets.VirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.VirtualServiceSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockVirtualServiceSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockVirtualServiceSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockVirtualServiceSet) Find(id ezkube.ResourceId) (*v1.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockVirtualServiceSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockVirtualServiceSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockVirtualServiceSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockVirtualServiceSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockVirtualServiceSet)(nil).Length))
}

// Generic mocks base method
func (m *MockVirtualServiceSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockVirtualServiceSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockVirtualServiceSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockVirtualServiceSet) Delta(newSet v1sets.VirtualServiceSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockVirtualServiceSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockVirtualServiceSet)(nil).Delta), newSet)
}
