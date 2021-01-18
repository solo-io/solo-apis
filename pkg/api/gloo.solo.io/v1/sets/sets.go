// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets



import (
    gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"

    "github.com/rotisserie/eris"
    sksets "github.com/solo-io/skv2/contrib/pkg/sets"
    "github.com/solo-io/skv2/pkg/ezkube"
    "k8s.io/apimachinery/pkg/util/sets"
)

type SettingsSet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    List(filterResource ... func(*gloo_solo_io_v1.Settings) bool) []*gloo_solo_io_v1.Settings
    // Return the Set as a map of key to resource.
    Map() map[string]*gloo_solo_io_v1.Settings
    // Insert a resource into the set.
    Insert(settings ...*gloo_solo_io_v1.Settings)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(settingsSet SettingsSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(settings ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(settings  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set SettingsSet) SettingsSet
    // Return the difference with the provided set
    Difference(set SettingsSet) SettingsSet
    // Return the intersection with the provided set
    Intersection(set SettingsSet) SettingsSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Settings, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another SettingsSet
    Delta(newSet SettingsSet) sksets.ResourceDelta
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
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *settingsSet) List(filterResource ... func(*gloo_solo_io_v1.Settings) bool) []*gloo_solo_io_v1.Settings {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*gloo_solo_io_v1.Settings))
        })
    }

    var settingsList []*gloo_solo_io_v1.Settings
    for _, obj := range s.Generic().List(genericFilters...) {
        settingsList = append(settingsList, obj.(*gloo_solo_io_v1.Settings))
    }
    return settingsList
}

func (s *settingsSet) Map() map[string]*gloo_solo_io_v1.Settings {
    if s == nil {
        return nil
    }

    newMap := map[string]*gloo_solo_io_v1.Settings{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*gloo_solo_io_v1.Settings)
    }
    return newMap
}

func (s *settingsSet) Insert(
        settingsList ...*gloo_solo_io_v1.Settings,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range settingsList {
        s.Generic().Insert(obj)
    }
}

func (s *settingsSet) Has(settings ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(settings)
}

func (s *settingsSet) Equal(
        settingsSet SettingsSet,
) bool {
    if s == nil {
        return settingsSet == nil
    }
    return s.Generic().Equal(settingsSet.Generic())
}

func (s *settingsSet) Delete(Settings ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(Settings)
}

func (s *settingsSet) Union(set SettingsSet) SettingsSet {
    if s == nil {
        return set
    }
    return NewSettingsSet(append(s.List(), set.List()...)...)
}

func (s *settingsSet) Difference(set SettingsSet) SettingsSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &settingsSet{set: newSet}
}

func (s *settingsSet) Intersection(set SettingsSet) SettingsSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var settingsList []*gloo_solo_io_v1.Settings
    for _, obj := range newSet.List() {
        settingsList = append(settingsList, obj.(*gloo_solo_io_v1.Settings))
    }
    return NewSettingsSet(settingsList...)
}


func (s *settingsSet) Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Settings, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find Settings %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&gloo_solo_io_v1.Settings{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*gloo_solo_io_v1.Settings), nil
}

func (s *settingsSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *settingsSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *settingsSet) Delta(newSet SettingsSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

type UpstreamSet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    List(filterResource ... func(*gloo_solo_io_v1.Upstream) bool) []*gloo_solo_io_v1.Upstream
    // Return the Set as a map of key to resource.
    Map() map[string]*gloo_solo_io_v1.Upstream
    // Insert a resource into the set.
    Insert(upstream ...*gloo_solo_io_v1.Upstream)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(upstreamSet UpstreamSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(upstream ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(upstream  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set UpstreamSet) UpstreamSet
    // Return the difference with the provided set
    Difference(set UpstreamSet) UpstreamSet
    // Return the intersection with the provided set
    Intersection(set UpstreamSet) UpstreamSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Upstream, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another UpstreamSet
    Delta(newSet UpstreamSet) sksets.ResourceDelta
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
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *upstreamSet) List(filterResource ... func(*gloo_solo_io_v1.Upstream) bool) []*gloo_solo_io_v1.Upstream {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*gloo_solo_io_v1.Upstream))
        })
    }

    var upstreamList []*gloo_solo_io_v1.Upstream
    for _, obj := range s.Generic().List(genericFilters...) {
        upstreamList = append(upstreamList, obj.(*gloo_solo_io_v1.Upstream))
    }
    return upstreamList
}

func (s *upstreamSet) Map() map[string]*gloo_solo_io_v1.Upstream {
    if s == nil {
        return nil
    }

    newMap := map[string]*gloo_solo_io_v1.Upstream{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*gloo_solo_io_v1.Upstream)
    }
    return newMap
}

func (s *upstreamSet) Insert(
        upstreamList ...*gloo_solo_io_v1.Upstream,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range upstreamList {
        s.Generic().Insert(obj)
    }
}

func (s *upstreamSet) Has(upstream ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(upstream)
}

func (s *upstreamSet) Equal(
        upstreamSet UpstreamSet,
) bool {
    if s == nil {
        return upstreamSet == nil
    }
    return s.Generic().Equal(upstreamSet.Generic())
}

func (s *upstreamSet) Delete(Upstream ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(Upstream)
}

func (s *upstreamSet) Union(set UpstreamSet) UpstreamSet {
    if s == nil {
        return set
    }
    return NewUpstreamSet(append(s.List(), set.List()...)...)
}

func (s *upstreamSet) Difference(set UpstreamSet) UpstreamSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &upstreamSet{set: newSet}
}

func (s *upstreamSet) Intersection(set UpstreamSet) UpstreamSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var upstreamList []*gloo_solo_io_v1.Upstream
    for _, obj := range newSet.List() {
        upstreamList = append(upstreamList, obj.(*gloo_solo_io_v1.Upstream))
    }
    return NewUpstreamSet(upstreamList...)
}


func (s *upstreamSet) Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Upstream, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find Upstream %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&gloo_solo_io_v1.Upstream{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*gloo_solo_io_v1.Upstream), nil
}

func (s *upstreamSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *upstreamSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *upstreamSet) Delta(newSet UpstreamSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

type UpstreamGroupSet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    List(filterResource ... func(*gloo_solo_io_v1.UpstreamGroup) bool) []*gloo_solo_io_v1.UpstreamGroup
    // Return the Set as a map of key to resource.
    Map() map[string]*gloo_solo_io_v1.UpstreamGroup
    // Insert a resource into the set.
    Insert(upstreamGroup ...*gloo_solo_io_v1.UpstreamGroup)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(upstreamGroupSet UpstreamGroupSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(upstreamGroup ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(upstreamGroup  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set UpstreamGroupSet) UpstreamGroupSet
    // Return the difference with the provided set
    Difference(set UpstreamGroupSet) UpstreamGroupSet
    // Return the intersection with the provided set
    Intersection(set UpstreamGroupSet) UpstreamGroupSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*gloo_solo_io_v1.UpstreamGroup, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another UpstreamGroupSet
    Delta(newSet UpstreamGroupSet) sksets.ResourceDelta
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
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *upstreamGroupSet) List(filterResource ... func(*gloo_solo_io_v1.UpstreamGroup) bool) []*gloo_solo_io_v1.UpstreamGroup {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*gloo_solo_io_v1.UpstreamGroup))
        })
    }

    var upstreamGroupList []*gloo_solo_io_v1.UpstreamGroup
    for _, obj := range s.Generic().List(genericFilters...) {
        upstreamGroupList = append(upstreamGroupList, obj.(*gloo_solo_io_v1.UpstreamGroup))
    }
    return upstreamGroupList
}

func (s *upstreamGroupSet) Map() map[string]*gloo_solo_io_v1.UpstreamGroup {
    if s == nil {
        return nil
    }

    newMap := map[string]*gloo_solo_io_v1.UpstreamGroup{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*gloo_solo_io_v1.UpstreamGroup)
    }
    return newMap
}

func (s *upstreamGroupSet) Insert(
        upstreamGroupList ...*gloo_solo_io_v1.UpstreamGroup,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range upstreamGroupList {
        s.Generic().Insert(obj)
    }
}

func (s *upstreamGroupSet) Has(upstreamGroup ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(upstreamGroup)
}

func (s *upstreamGroupSet) Equal(
        upstreamGroupSet UpstreamGroupSet,
) bool {
    if s == nil {
        return upstreamGroupSet == nil
    }
    return s.Generic().Equal(upstreamGroupSet.Generic())
}

func (s *upstreamGroupSet) Delete(UpstreamGroup ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(UpstreamGroup)
}

func (s *upstreamGroupSet) Union(set UpstreamGroupSet) UpstreamGroupSet {
    if s == nil {
        return set
    }
    return NewUpstreamGroupSet(append(s.List(), set.List()...)...)
}

func (s *upstreamGroupSet) Difference(set UpstreamGroupSet) UpstreamGroupSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &upstreamGroupSet{set: newSet}
}

func (s *upstreamGroupSet) Intersection(set UpstreamGroupSet) UpstreamGroupSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var upstreamGroupList []*gloo_solo_io_v1.UpstreamGroup
    for _, obj := range newSet.List() {
        upstreamGroupList = append(upstreamGroupList, obj.(*gloo_solo_io_v1.UpstreamGroup))
    }
    return NewUpstreamGroupSet(upstreamGroupList...)
}


func (s *upstreamGroupSet) Find(id ezkube.ResourceId) (*gloo_solo_io_v1.UpstreamGroup, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find UpstreamGroup %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&gloo_solo_io_v1.UpstreamGroup{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*gloo_solo_io_v1.UpstreamGroup), nil
}

func (s *upstreamGroupSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *upstreamGroupSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *upstreamGroupSet) Delta(newSet UpstreamGroupSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

type ProxySet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    List(filterResource ... func(*gloo_solo_io_v1.Proxy) bool) []*gloo_solo_io_v1.Proxy
    // Return the Set as a map of key to resource.
    Map() map[string]*gloo_solo_io_v1.Proxy
    // Insert a resource into the set.
    Insert(proxy ...*gloo_solo_io_v1.Proxy)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(proxySet ProxySet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(proxy ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(proxy  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set ProxySet) ProxySet
    // Return the difference with the provided set
    Difference(set ProxySet) ProxySet
    // Return the intersection with the provided set
    Intersection(set ProxySet) ProxySet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Proxy, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another ProxySet
    Delta(newSet ProxySet) sksets.ResourceDelta
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
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *proxySet) List(filterResource ... func(*gloo_solo_io_v1.Proxy) bool) []*gloo_solo_io_v1.Proxy {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*gloo_solo_io_v1.Proxy))
        })
    }

    var proxyList []*gloo_solo_io_v1.Proxy
    for _, obj := range s.Generic().List(genericFilters...) {
        proxyList = append(proxyList, obj.(*gloo_solo_io_v1.Proxy))
    }
    return proxyList
}

func (s *proxySet) Map() map[string]*gloo_solo_io_v1.Proxy {
    if s == nil {
        return nil
    }

    newMap := map[string]*gloo_solo_io_v1.Proxy{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*gloo_solo_io_v1.Proxy)
    }
    return newMap
}

func (s *proxySet) Insert(
        proxyList ...*gloo_solo_io_v1.Proxy,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range proxyList {
        s.Generic().Insert(obj)
    }
}

func (s *proxySet) Has(proxy ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(proxy)
}

func (s *proxySet) Equal(
        proxySet ProxySet,
) bool {
    if s == nil {
        return proxySet == nil
    }
    return s.Generic().Equal(proxySet.Generic())
}

func (s *proxySet) Delete(Proxy ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(Proxy)
}

func (s *proxySet) Union(set ProxySet) ProxySet {
    if s == nil {
        return set
    }
    return NewProxySet(append(s.List(), set.List()...)...)
}

func (s *proxySet) Difference(set ProxySet) ProxySet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &proxySet{set: newSet}
}

func (s *proxySet) Intersection(set ProxySet) ProxySet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var proxyList []*gloo_solo_io_v1.Proxy
    for _, obj := range newSet.List() {
        proxyList = append(proxyList, obj.(*gloo_solo_io_v1.Proxy))
    }
    return NewProxySet(proxyList...)
}


func (s *proxySet) Find(id ezkube.ResourceId) (*gloo_solo_io_v1.Proxy, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find Proxy %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&gloo_solo_io_v1.Proxy{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*gloo_solo_io_v1.Proxy), nil
}

func (s *proxySet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *proxySet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *proxySet) Delta(newSet ProxySet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}
