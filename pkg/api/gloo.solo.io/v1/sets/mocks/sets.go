// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	v1sets "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/sets"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockSettingsSet is a mock of SettingsSet interface
type MockSettingsSet struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsSetMockRecorder
}

// MockSettingsSetMockRecorder is the mock recorder for MockSettingsSet
type MockSettingsSetMockRecorder struct {
	mock *MockSettingsSet
}

// NewMockSettingsSet creates a new mock instance
func NewMockSettingsSet(ctrl *gomock.Controller) *MockSettingsSet {
	mock := &MockSettingsSet{ctrl: ctrl}
	mock.recorder = &MockSettingsSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsSet) EXPECT() *MockSettingsSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockSettingsSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockSettingsSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockSettingsSet)(nil).Keys))
}

// List mocks base method
func (m *MockSettingsSet) List() []*v1.Settings {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1.Settings)
	return ret0
}

// List indicates an expected call of List
func (mr *MockSettingsSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSettingsSet)(nil).List))
}

// Map mocks base method
func (m *MockSettingsSet) Map() map[string]*v1.Settings {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Settings)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockSettingsSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockSettingsSet)(nil).Map))
}

// Insert mocks base method
func (m *MockSettingsSet) Insert(settings ...*v1.Settings) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range settings {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockSettingsSetMockRecorder) Insert(settings ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSettingsSet)(nil).Insert), settings...)
}

// Equal mocks base method
func (m *MockSettingsSet) Equal(settingsSet v1sets.SettingsSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", settingsSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockSettingsSetMockRecorder) Equal(settingsSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockSettingsSet)(nil).Equal), settingsSet)
}

// Has mocks base method
func (m *MockSettingsSet) Has(settings *v1.Settings) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", settings)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockSettingsSetMockRecorder) Has(settings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockSettingsSet)(nil).Has), settings)
}

// Delete mocks base method
func (m *MockSettingsSet) Delete(settings *v1.Settings) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", settings)
}

// Delete indicates an expected call of Delete
func (mr *MockSettingsSetMockRecorder) Delete(settings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSettingsSet)(nil).Delete), settings)
}

// Union mocks base method
func (m *MockSettingsSet) Union(set v1sets.SettingsSet) v1sets.SettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.SettingsSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockSettingsSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockSettingsSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockSettingsSet) Difference(set v1sets.SettingsSet) v1sets.SettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.SettingsSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockSettingsSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockSettingsSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockSettingsSet) Intersection(set v1sets.SettingsSet) v1sets.SettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.SettingsSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockSettingsSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockSettingsSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockSettingsSet) Find(id ezkube.ResourceId) (*v1.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockSettingsSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockSettingsSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockSettingsSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockSettingsSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockSettingsSet)(nil).Length))
}

// MockUpstreamSet is a mock of UpstreamSet interface
type MockUpstreamSet struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamSetMockRecorder
}

// MockUpstreamSetMockRecorder is the mock recorder for MockUpstreamSet
type MockUpstreamSetMockRecorder struct {
	mock *MockUpstreamSet
}

// NewMockUpstreamSet creates a new mock instance
func NewMockUpstreamSet(ctrl *gomock.Controller) *MockUpstreamSet {
	mock := &MockUpstreamSet{ctrl: ctrl}
	mock.recorder = &MockUpstreamSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamSet) EXPECT() *MockUpstreamSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockUpstreamSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockUpstreamSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockUpstreamSet)(nil).Keys))
}

// List mocks base method
func (m *MockUpstreamSet) List() []*v1.Upstream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1.Upstream)
	return ret0
}

// List indicates an expected call of List
func (mr *MockUpstreamSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUpstreamSet)(nil).List))
}

// Map mocks base method
func (m *MockUpstreamSet) Map() map[string]*v1.Upstream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Upstream)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockUpstreamSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockUpstreamSet)(nil).Map))
}

// Insert mocks base method
func (m *MockUpstreamSet) Insert(upstream ...*v1.Upstream) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range upstream {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockUpstreamSetMockRecorder) Insert(upstream ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUpstreamSet)(nil).Insert), upstream...)
}

// Equal mocks base method
func (m *MockUpstreamSet) Equal(upstreamSet v1sets.UpstreamSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", upstreamSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockUpstreamSetMockRecorder) Equal(upstreamSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockUpstreamSet)(nil).Equal), upstreamSet)
}

// Has mocks base method
func (m *MockUpstreamSet) Has(upstream *v1.Upstream) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", upstream)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockUpstreamSetMockRecorder) Has(upstream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockUpstreamSet)(nil).Has), upstream)
}

// Delete mocks base method
func (m *MockUpstreamSet) Delete(upstream *v1.Upstream) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", upstream)
}

// Delete indicates an expected call of Delete
func (mr *MockUpstreamSetMockRecorder) Delete(upstream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUpstreamSet)(nil).Delete), upstream)
}

// Union mocks base method
func (m *MockUpstreamSet) Union(set v1sets.UpstreamSet) v1sets.UpstreamSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.UpstreamSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockUpstreamSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockUpstreamSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockUpstreamSet) Difference(set v1sets.UpstreamSet) v1sets.UpstreamSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.UpstreamSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockUpstreamSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockUpstreamSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockUpstreamSet) Intersection(set v1sets.UpstreamSet) v1sets.UpstreamSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.UpstreamSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockUpstreamSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockUpstreamSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockUpstreamSet) Find(id ezkube.ResourceId) (*v1.Upstream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Upstream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUpstreamSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUpstreamSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockUpstreamSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockUpstreamSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockUpstreamSet)(nil).Length))
}

// MockUpstreamGroupSet is a mock of UpstreamGroupSet interface
type MockUpstreamGroupSet struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamGroupSetMockRecorder
}

// MockUpstreamGroupSetMockRecorder is the mock recorder for MockUpstreamGroupSet
type MockUpstreamGroupSetMockRecorder struct {
	mock *MockUpstreamGroupSet
}

// NewMockUpstreamGroupSet creates a new mock instance
func NewMockUpstreamGroupSet(ctrl *gomock.Controller) *MockUpstreamGroupSet {
	mock := &MockUpstreamGroupSet{ctrl: ctrl}
	mock.recorder = &MockUpstreamGroupSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamGroupSet) EXPECT() *MockUpstreamGroupSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockUpstreamGroupSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockUpstreamGroupSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Keys))
}

// List mocks base method
func (m *MockUpstreamGroupSet) List() []*v1.UpstreamGroup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1.UpstreamGroup)
	return ret0
}

// List indicates an expected call of List
func (mr *MockUpstreamGroupSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUpstreamGroupSet)(nil).List))
}

// Map mocks base method
func (m *MockUpstreamGroupSet) Map() map[string]*v1.UpstreamGroup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.UpstreamGroup)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockUpstreamGroupSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Map))
}

// Insert mocks base method
func (m *MockUpstreamGroupSet) Insert(upstreamGroup ...*v1.UpstreamGroup) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range upstreamGroup {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockUpstreamGroupSetMockRecorder) Insert(upstreamGroup ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Insert), upstreamGroup...)
}

// Equal mocks base method
func (m *MockUpstreamGroupSet) Equal(upstreamGroupSet v1sets.UpstreamGroupSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", upstreamGroupSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockUpstreamGroupSetMockRecorder) Equal(upstreamGroupSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Equal), upstreamGroupSet)
}

// Has mocks base method
func (m *MockUpstreamGroupSet) Has(upstreamGroup *v1.UpstreamGroup) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", upstreamGroup)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockUpstreamGroupSetMockRecorder) Has(upstreamGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Has), upstreamGroup)
}

// Delete mocks base method
func (m *MockUpstreamGroupSet) Delete(upstreamGroup *v1.UpstreamGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", upstreamGroup)
}

// Delete indicates an expected call of Delete
func (mr *MockUpstreamGroupSetMockRecorder) Delete(upstreamGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Delete), upstreamGroup)
}

// Union mocks base method
func (m *MockUpstreamGroupSet) Union(set v1sets.UpstreamGroupSet) v1sets.UpstreamGroupSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.UpstreamGroupSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockUpstreamGroupSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockUpstreamGroupSet) Difference(set v1sets.UpstreamGroupSet) v1sets.UpstreamGroupSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.UpstreamGroupSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockUpstreamGroupSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockUpstreamGroupSet) Intersection(set v1sets.UpstreamGroupSet) v1sets.UpstreamGroupSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.UpstreamGroupSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockUpstreamGroupSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockUpstreamGroupSet) Find(id ezkube.ResourceId) (*v1.UpstreamGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.UpstreamGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUpstreamGroupSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockUpstreamGroupSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockUpstreamGroupSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockUpstreamGroupSet)(nil).Length))
}

// MockProxySet is a mock of ProxySet interface
type MockProxySet struct {
	ctrl     *gomock.Controller
	recorder *MockProxySetMockRecorder
}

// MockProxySetMockRecorder is the mock recorder for MockProxySet
type MockProxySetMockRecorder struct {
	mock *MockProxySet
}

// NewMockProxySet creates a new mock instance
func NewMockProxySet(ctrl *gomock.Controller) *MockProxySet {
	mock := &MockProxySet{ctrl: ctrl}
	mock.recorder = &MockProxySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProxySet) EXPECT() *MockProxySetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockProxySet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockProxySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockProxySet)(nil).Keys))
}

// List mocks base method
func (m *MockProxySet) List() []*v1.Proxy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1.Proxy)
	return ret0
}

// List indicates an expected call of List
func (mr *MockProxySetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProxySet)(nil).List))
}

// Map mocks base method
func (m *MockProxySet) Map() map[string]*v1.Proxy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Proxy)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockProxySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockProxySet)(nil).Map))
}

// Insert mocks base method
func (m *MockProxySet) Insert(proxy ...*v1.Proxy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range proxy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockProxySetMockRecorder) Insert(proxy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockProxySet)(nil).Insert), proxy...)
}

// Equal mocks base method
func (m *MockProxySet) Equal(proxySet v1sets.ProxySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", proxySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockProxySetMockRecorder) Equal(proxySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockProxySet)(nil).Equal), proxySet)
}

// Has mocks base method
func (m *MockProxySet) Has(proxy *v1.Proxy) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", proxy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockProxySetMockRecorder) Has(proxy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockProxySet)(nil).Has), proxy)
}

// Delete mocks base method
func (m *MockProxySet) Delete(proxy *v1.Proxy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", proxy)
}

// Delete indicates an expected call of Delete
func (mr *MockProxySetMockRecorder) Delete(proxy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProxySet)(nil).Delete), proxy)
}

// Union mocks base method
func (m *MockProxySet) Union(set v1sets.ProxySet) v1sets.ProxySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.ProxySet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockProxySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockProxySet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockProxySet) Difference(set v1sets.ProxySet) v1sets.ProxySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.ProxySet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockProxySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockProxySet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockProxySet) Intersection(set v1sets.ProxySet) v1sets.ProxySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.ProxySet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockProxySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockProxySet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockProxySet) Find(id ezkube.ResourceId) (*v1.Proxy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Proxy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockProxySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockProxySet)(nil).Find), id)
}

// Length mocks base method
func (m *MockProxySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockProxySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockProxySet)(nil).Length))
}
