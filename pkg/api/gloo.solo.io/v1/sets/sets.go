// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	gloo_solo_io_v1 "github.com/solo-io/gloo-fed/pkg/api/gloo.solo.io/v1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type AuthConfigSet interface {
	Keys() sets.String
	List() []*gloo_solo_io_v1.AuthConfig
	Map() map[string]*gloo_solo_io_v1.AuthConfig
	Insert(authConfig ...*gloo_solo_io_v1.AuthConfig)
	Equal(authConfigSet AuthConfigSet) bool
	Has(authConfig *gloo_solo_io_v1.AuthConfig) bool
	Delete(authConfig *gloo_solo_io_v1.AuthConfig)
	Union(set AuthConfigSet) AuthConfigSet
	Difference(set AuthConfigSet) AuthConfigSet
	Intersection(set AuthConfigSet) AuthConfigSet
	Find(id ezkube.ResourceId) (*gloo_solo_io_v1.AuthConfig, error)
	Length() int
}

func makeGenericAuthConfigSet(authConfigList []*gloo_solo_io_v1.AuthConfig) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range authConfigList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type authConfigSet struct {
	set sksets.ResourceSet
}

func NewAuthConfigSet(authConfigList ...*gloo_solo_io_v1.AuthConfig) AuthConfigSet {
	return &authConfigSet{set: makeGenericAuthConfigSet(authConfigList)}
}

func NewAuthConfigSetFromList(authConfigList *gloo_solo_io_v1.AuthConfigList) AuthConfigSet {
	list := make([]*gloo_solo_io_v1.AuthConfig, 0, len(authConfigList.Items))
	for idx := range authConfigList.Items {
		list = append(list, &authConfigList.Items[idx])
	}
	return &authConfigSet{set: makeGenericAuthConfigSet(list)}
}

func (s *authConfigSet) Keys() sets.String {
	return s.set.Keys()
}

func (s *authConfigSet) List() []*gloo_solo_io_v1.AuthConfig {
	var authConfigList []*gloo_solo_io_v1.AuthConfig
	for _, obj := range s.set.List() {
		authConfigList = append(authConfigList, obj.(*gloo_solo_io_v1.AuthConfig))
	}
	return authConfigList
}

func (s *authConfigSet) Map() map[string]*gloo_solo_io_v1.AuthConfig {
	newMap := map[string]*gloo_solo_io_v1.AuthConfig{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*gloo_solo_io_v1.AuthConfig)
	}
	return newMap
}

func (s *authConfigSet) Insert(
	authConfigList ...*gloo_solo_io_v1.AuthConfig,
) {
	for _, obj := range authConfigList {
		s.set.Insert(obj)
	}
}

func (s *authConfigSet) Has(authConfig *gloo_solo_io_v1.AuthConfig) bool {
	return s.set.Has(authConfig)
}

func (s *authConfigSet) Equal(
	authConfigSet AuthConfigSet,
) bool {
	return s.set.Equal(makeGenericAuthConfigSet(authConfigSet.List()))
}

func (s *authConfigSet) Delete(AuthConfig *gloo_solo_io_v1.AuthConfig) {
	s.set.Delete(AuthConfig)
}

func (s *authConfigSet) Union(set AuthConfigSet) AuthConfigSet {
	return NewAuthConfigSet(append(s.List(), set.List()...)...)
}

func (s *authConfigSet) Difference(set AuthConfigSet) AuthConfigSet {
	newSet := s.set.Difference(makeGenericAuthConfigSet(set.List()))
	return &authConfigSet{set: newSet}
}

func (s *authConfigSet) Intersection(set AuthConfigSet) AuthConfigSet {
	newSet := s.set.Intersection(makeGenericAuthConfigSet(set.List()))
	var authConfigList []*gloo_solo_io_v1.AuthConfig
	for _, obj := range newSet.List() {
		authConfigList = append(authConfigList, obj.(*gloo_solo_io_v1.AuthConfig))
	}
	return NewAuthConfigSet(authConfigList...)
}

func (s *authConfigSet) Find(id ezkube.ResourceId) (*gloo_solo_io_v1.AuthConfig, error) {
	obj, err := s.set.Find(&gloo_solo_io_v1.AuthConfig{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*gloo_solo_io_v1.AuthConfig), nil
}

func (s *authConfigSet) Length() int {
	return s.set.Length()
}

type SettingsSet interface {
	Keys() sets.String
	List() []*gloo_solo_io_v1.Settings
	Map() map[string]*gloo_solo_io_v1.Settings
	Insert(settings ...*gloo_solo_io_v1.Settings)
	Equal(settingsSet SettingsSet) bool
	Has(settings *gloo_solo_io_v1.Settings) bool
	Delete(settings *gloo_solo_io_v1.Settings)
	Union(set SettingsSet) SettingsSet
	Difference(set SettingsSet) SettingsSet
	Intersection(set SettingsSet) SettingsSet
	Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Settings, error)
	Length() int
}

func makeGenericSettingsSet(settingsList []*gloo_solo_io_v1.Settings) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range settingsList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type settingsSet struct {
	set sksets.ResourceSet
}

func NewSettingsSet(settingsList ...*gloo_solo_io_v1.Settings) SettingsSet {
	return &settingsSet{set: makeGenericSettingsSet(settingsList)}
}

func NewSettingsSetFromList(settingsList *gloo_solo_io_v1.SettingsList) SettingsSet {
	list := make([]*gloo_solo_io_v1.Settings, 0, len(settingsList.Items))
	for idx := range settingsList.Items {
		list = append(list, &settingsList.Items[idx])
	}
	return &settingsSet{set: makeGenericSettingsSet(list)}
}

func (s *settingsSet) Keys() sets.String {
	return s.set.Keys()
}

func (s *settingsSet) List() []*gloo_solo_io_v1.Settings {
	var settingsList []*gloo_solo_io_v1.Settings
	for _, obj := range s.set.List() {
		settingsList = append(settingsList, obj.(*gloo_solo_io_v1.Settings))
	}
	return settingsList
}

func (s *settingsSet) Map() map[string]*gloo_solo_io_v1.Settings {
	newMap := map[string]*gloo_solo_io_v1.Settings{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*gloo_solo_io_v1.Settings)
	}
	return newMap
}

func (s *settingsSet) Insert(
	settingsList ...*gloo_solo_io_v1.Settings,
) {
	for _, obj := range settingsList {
		s.set.Insert(obj)
	}
}

func (s *settingsSet) Has(settings *gloo_solo_io_v1.Settings) bool {
	return s.set.Has(settings)
}

func (s *settingsSet) Equal(
	settingsSet SettingsSet,
) bool {
	return s.set.Equal(makeGenericSettingsSet(settingsSet.List()))
}

func (s *settingsSet) Delete(Settings *gloo_solo_io_v1.Settings) {
	s.set.Delete(Settings)
}

func (s *settingsSet) Union(set SettingsSet) SettingsSet {
	return NewSettingsSet(append(s.List(), set.List()...)...)
}

func (s *settingsSet) Difference(set SettingsSet) SettingsSet {
	newSet := s.set.Difference(makeGenericSettingsSet(set.List()))
	return &settingsSet{set: newSet}
}

func (s *settingsSet) Intersection(set SettingsSet) SettingsSet {
	newSet := s.set.Intersection(makeGenericSettingsSet(set.List()))
	var settingsList []*gloo_solo_io_v1.Settings
	for _, obj := range newSet.List() {
		settingsList = append(settingsList, obj.(*gloo_solo_io_v1.Settings))
	}
	return NewSettingsSet(settingsList...)
}

func (s *settingsSet) Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Settings, error) {
	obj, err := s.set.Find(&gloo_solo_io_v1.Settings{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*gloo_solo_io_v1.Settings), nil
}

func (s *settingsSet) Length() int {
	return s.set.Length()
}

type UpstreamSet interface {
	Keys() sets.String
	List() []*gloo_solo_io_v1.Upstream
	Map() map[string]*gloo_solo_io_v1.Upstream
	Insert(upstream ...*gloo_solo_io_v1.Upstream)
	Equal(upstreamSet UpstreamSet) bool
	Has(upstream *gloo_solo_io_v1.Upstream) bool
	Delete(upstream *gloo_solo_io_v1.Upstream)
	Union(set UpstreamSet) UpstreamSet
	Difference(set UpstreamSet) UpstreamSet
	Intersection(set UpstreamSet) UpstreamSet
	Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Upstream, error)
	Length() int
}

func makeGenericUpstreamSet(upstreamList []*gloo_solo_io_v1.Upstream) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range upstreamList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type upstreamSet struct {
	set sksets.ResourceSet
}

func NewUpstreamSet(upstreamList ...*gloo_solo_io_v1.Upstream) UpstreamSet {
	return &upstreamSet{set: makeGenericUpstreamSet(upstreamList)}
}

func NewUpstreamSetFromList(upstreamList *gloo_solo_io_v1.UpstreamList) UpstreamSet {
	list := make([]*gloo_solo_io_v1.Upstream, 0, len(upstreamList.Items))
	for idx := range upstreamList.Items {
		list = append(list, &upstreamList.Items[idx])
	}
	return &upstreamSet{set: makeGenericUpstreamSet(list)}
}

func (s *upstreamSet) Keys() sets.String {
	return s.set.Keys()
}

func (s *upstreamSet) List() []*gloo_solo_io_v1.Upstream {
	var upstreamList []*gloo_solo_io_v1.Upstream
	for _, obj := range s.set.List() {
		upstreamList = append(upstreamList, obj.(*gloo_solo_io_v1.Upstream))
	}
	return upstreamList
}

func (s *upstreamSet) Map() map[string]*gloo_solo_io_v1.Upstream {
	newMap := map[string]*gloo_solo_io_v1.Upstream{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*gloo_solo_io_v1.Upstream)
	}
	return newMap
}

func (s *upstreamSet) Insert(
	upstreamList ...*gloo_solo_io_v1.Upstream,
) {
	for _, obj := range upstreamList {
		s.set.Insert(obj)
	}
}

func (s *upstreamSet) Has(upstream *gloo_solo_io_v1.Upstream) bool {
	return s.set.Has(upstream)
}

func (s *upstreamSet) Equal(
	upstreamSet UpstreamSet,
) bool {
	return s.set.Equal(makeGenericUpstreamSet(upstreamSet.List()))
}

func (s *upstreamSet) Delete(Upstream *gloo_solo_io_v1.Upstream) {
	s.set.Delete(Upstream)
}

func (s *upstreamSet) Union(set UpstreamSet) UpstreamSet {
	return NewUpstreamSet(append(s.List(), set.List()...)...)
}

func (s *upstreamSet) Difference(set UpstreamSet) UpstreamSet {
	newSet := s.set.Difference(makeGenericUpstreamSet(set.List()))
	return &upstreamSet{set: newSet}
}

func (s *upstreamSet) Intersection(set UpstreamSet) UpstreamSet {
	newSet := s.set.Intersection(makeGenericUpstreamSet(set.List()))
	var upstreamList []*gloo_solo_io_v1.Upstream
	for _, obj := range newSet.List() {
		upstreamList = append(upstreamList, obj.(*gloo_solo_io_v1.Upstream))
	}
	return NewUpstreamSet(upstreamList...)
}

func (s *upstreamSet) Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Upstream, error) {
	obj, err := s.set.Find(&gloo_solo_io_v1.Upstream{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*gloo_solo_io_v1.Upstream), nil
}

func (s *upstreamSet) Length() int {
	return s.set.Length()
}

type UpstreamGroupSet interface {
	Keys() sets.String
	List() []*gloo_solo_io_v1.UpstreamGroup
	Map() map[string]*gloo_solo_io_v1.UpstreamGroup
	Insert(upstreamGroup ...*gloo_solo_io_v1.UpstreamGroup)
	Equal(upstreamGroupSet UpstreamGroupSet) bool
	Has(upstreamGroup *gloo_solo_io_v1.UpstreamGroup) bool
	Delete(upstreamGroup *gloo_solo_io_v1.UpstreamGroup)
	Union(set UpstreamGroupSet) UpstreamGroupSet
	Difference(set UpstreamGroupSet) UpstreamGroupSet
	Intersection(set UpstreamGroupSet) UpstreamGroupSet
	Find(id ezkube.ResourceId) (*gloo_solo_io_v1.UpstreamGroup, error)
	Length() int
}

func makeGenericUpstreamGroupSet(upstreamGroupList []*gloo_solo_io_v1.UpstreamGroup) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range upstreamGroupList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type upstreamGroupSet struct {
	set sksets.ResourceSet
}

func NewUpstreamGroupSet(upstreamGroupList ...*gloo_solo_io_v1.UpstreamGroup) UpstreamGroupSet {
	return &upstreamGroupSet{set: makeGenericUpstreamGroupSet(upstreamGroupList)}
}

func NewUpstreamGroupSetFromList(upstreamGroupList *gloo_solo_io_v1.UpstreamGroupList) UpstreamGroupSet {
	list := make([]*gloo_solo_io_v1.UpstreamGroup, 0, len(upstreamGroupList.Items))
	for idx := range upstreamGroupList.Items {
		list = append(list, &upstreamGroupList.Items[idx])
	}
	return &upstreamGroupSet{set: makeGenericUpstreamGroupSet(list)}
}

func (s *upstreamGroupSet) Keys() sets.String {
	return s.set.Keys()
}

func (s *upstreamGroupSet) List() []*gloo_solo_io_v1.UpstreamGroup {
	var upstreamGroupList []*gloo_solo_io_v1.UpstreamGroup
	for _, obj := range s.set.List() {
		upstreamGroupList = append(upstreamGroupList, obj.(*gloo_solo_io_v1.UpstreamGroup))
	}
	return upstreamGroupList
}

func (s *upstreamGroupSet) Map() map[string]*gloo_solo_io_v1.UpstreamGroup {
	newMap := map[string]*gloo_solo_io_v1.UpstreamGroup{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*gloo_solo_io_v1.UpstreamGroup)
	}
	return newMap
}

func (s *upstreamGroupSet) Insert(
	upstreamGroupList ...*gloo_solo_io_v1.UpstreamGroup,
) {
	for _, obj := range upstreamGroupList {
		s.set.Insert(obj)
	}
}

func (s *upstreamGroupSet) Has(upstreamGroup *gloo_solo_io_v1.UpstreamGroup) bool {
	return s.set.Has(upstreamGroup)
}

func (s *upstreamGroupSet) Equal(
	upstreamGroupSet UpstreamGroupSet,
) bool {
	return s.set.Equal(makeGenericUpstreamGroupSet(upstreamGroupSet.List()))
}

func (s *upstreamGroupSet) Delete(UpstreamGroup *gloo_solo_io_v1.UpstreamGroup) {
	s.set.Delete(UpstreamGroup)
}

func (s *upstreamGroupSet) Union(set UpstreamGroupSet) UpstreamGroupSet {
	return NewUpstreamGroupSet(append(s.List(), set.List()...)...)
}

func (s *upstreamGroupSet) Difference(set UpstreamGroupSet) UpstreamGroupSet {
	newSet := s.set.Difference(makeGenericUpstreamGroupSet(set.List()))
	return &upstreamGroupSet{set: newSet}
}

func (s *upstreamGroupSet) Intersection(set UpstreamGroupSet) UpstreamGroupSet {
	newSet := s.set.Intersection(makeGenericUpstreamGroupSet(set.List()))
	var upstreamGroupList []*gloo_solo_io_v1.UpstreamGroup
	for _, obj := range newSet.List() {
		upstreamGroupList = append(upstreamGroupList, obj.(*gloo_solo_io_v1.UpstreamGroup))
	}
	return NewUpstreamGroupSet(upstreamGroupList...)
}

func (s *upstreamGroupSet) Find(id ezkube.ResourceId) (*gloo_solo_io_v1.UpstreamGroup, error) {
	obj, err := s.set.Find(&gloo_solo_io_v1.UpstreamGroup{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*gloo_solo_io_v1.UpstreamGroup), nil
}

func (s *upstreamGroupSet) Length() int {
	return s.set.Length()
}

type ProxySet interface {
	Keys() sets.String
	List() []*gloo_solo_io_v1.Proxy
	Map() map[string]*gloo_solo_io_v1.Proxy
	Insert(proxy ...*gloo_solo_io_v1.Proxy)
	Equal(proxySet ProxySet) bool
	Has(proxy *gloo_solo_io_v1.Proxy) bool
	Delete(proxy *gloo_solo_io_v1.Proxy)
	Union(set ProxySet) ProxySet
	Difference(set ProxySet) ProxySet
	Intersection(set ProxySet) ProxySet
	Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Proxy, error)
	Length() int
}

func makeGenericProxySet(proxyList []*gloo_solo_io_v1.Proxy) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range proxyList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type proxySet struct {
	set sksets.ResourceSet
}

func NewProxySet(proxyList ...*gloo_solo_io_v1.Proxy) ProxySet {
	return &proxySet{set: makeGenericProxySet(proxyList)}
}

func NewProxySetFromList(proxyList *gloo_solo_io_v1.ProxyList) ProxySet {
	list := make([]*gloo_solo_io_v1.Proxy, 0, len(proxyList.Items))
	for idx := range proxyList.Items {
		list = append(list, &proxyList.Items[idx])
	}
	return &proxySet{set: makeGenericProxySet(list)}
}

func (s *proxySet) Keys() sets.String {
	return s.set.Keys()
}

func (s *proxySet) List() []*gloo_solo_io_v1.Proxy {
	var proxyList []*gloo_solo_io_v1.Proxy
	for _, obj := range s.set.List() {
		proxyList = append(proxyList, obj.(*gloo_solo_io_v1.Proxy))
	}
	return proxyList
}

func (s *proxySet) Map() map[string]*gloo_solo_io_v1.Proxy {
	newMap := map[string]*gloo_solo_io_v1.Proxy{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*gloo_solo_io_v1.Proxy)
	}
	return newMap
}

func (s *proxySet) Insert(
	proxyList ...*gloo_solo_io_v1.Proxy,
) {
	for _, obj := range proxyList {
		s.set.Insert(obj)
	}
}

func (s *proxySet) Has(proxy *gloo_solo_io_v1.Proxy) bool {
	return s.set.Has(proxy)
}

func (s *proxySet) Equal(
	proxySet ProxySet,
) bool {
	return s.set.Equal(makeGenericProxySet(proxySet.List()))
}

func (s *proxySet) Delete(Proxy *gloo_solo_io_v1.Proxy) {
	s.set.Delete(Proxy)
}

func (s *proxySet) Union(set ProxySet) ProxySet {
	return NewProxySet(append(s.List(), set.List()...)...)
}

func (s *proxySet) Difference(set ProxySet) ProxySet {
	newSet := s.set.Difference(makeGenericProxySet(set.List()))
	return &proxySet{set: newSet}
}

func (s *proxySet) Intersection(set ProxySet) ProxySet {
	newSet := s.set.Intersection(makeGenericProxySet(set.List()))
	var proxyList []*gloo_solo_io_v1.Proxy
	for _, obj := range newSet.List() {
		proxyList = append(proxyList, obj.(*gloo_solo_io_v1.Proxy))
	}
	return NewProxySet(proxyList...)
}

func (s *proxySet) Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Proxy, error) {
	obj, err := s.set.Find(&gloo_solo_io_v1.Proxy{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*gloo_solo_io_v1.Proxy), nil
}

func (s *proxySet) Length() int {
	return s.set.Length()
}
