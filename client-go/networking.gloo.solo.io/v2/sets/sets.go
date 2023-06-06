// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v2sets

import (
	networking_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type ExternalServiceSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_gloo_solo_io_v2.ExternalService) bool) []*networking_gloo_solo_io_v2.ExternalService
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.ExternalService) bool) []*networking_gloo_solo_io_v2.ExternalService
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_gloo_solo_io_v2.ExternalService
	// Insert a resource into the set.
	Insert(externalService ...*networking_gloo_solo_io_v2.ExternalService)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(externalServiceSet ExternalServiceSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(externalService ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(externalService ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ExternalServiceSet) ExternalServiceSet
	// Return the difference with the provided set
	Difference(set ExternalServiceSet) ExternalServiceSet
	// Return the intersection with the provided set
	Intersection(set ExternalServiceSet) ExternalServiceSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.ExternalService, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another ExternalServiceSet
	Delta(newSet ExternalServiceSet) sksets.ResourceDelta
	// Create a deep copy of the current ExternalServiceSet
	Clone() ExternalServiceSet
}

func makeGenericExternalServiceSet(externalServiceList []*networking_gloo_solo_io_v2.ExternalService) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range externalServiceList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type externalServiceSet struct {
	set sksets.ResourceSet
}

func NewExternalServiceSet(externalServiceList ...*networking_gloo_solo_io_v2.ExternalService) ExternalServiceSet {
	return &externalServiceSet{set: makeGenericExternalServiceSet(externalServiceList)}
}

func NewExternalServiceSetFromList(externalServiceList *networking_gloo_solo_io_v2.ExternalServiceList) ExternalServiceSet {
	list := make([]*networking_gloo_solo_io_v2.ExternalService, 0, len(externalServiceList.Items))
	for idx := range externalServiceList.Items {
		list = append(list, &externalServiceList.Items[idx])
	}
	return &externalServiceSet{set: makeGenericExternalServiceSet(list)}
}

func (s *externalServiceSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *externalServiceSet) List(filterResource ...func(*networking_gloo_solo_io_v2.ExternalService) bool) []*networking_gloo_solo_io_v2.ExternalService {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.ExternalService))
		})
	}

	objs := s.Generic().List(genericFilters...)
	externalServiceList := make([]*networking_gloo_solo_io_v2.ExternalService, 0, len(objs))
	for _, obj := range objs {
		externalServiceList = append(externalServiceList, obj.(*networking_gloo_solo_io_v2.ExternalService))
	}
	return externalServiceList
}

func (s *externalServiceSet) UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.ExternalService) bool) []*networking_gloo_solo_io_v2.ExternalService {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.ExternalService))
		})
	}

	var externalServiceList []*networking_gloo_solo_io_v2.ExternalService
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		externalServiceList = append(externalServiceList, obj.(*networking_gloo_solo_io_v2.ExternalService))
	}
	return externalServiceList
}

func (s *externalServiceSet) Map() map[string]*networking_gloo_solo_io_v2.ExternalService {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_gloo_solo_io_v2.ExternalService{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_gloo_solo_io_v2.ExternalService)
	}
	return newMap
}

func (s *externalServiceSet) Insert(
	externalServiceList ...*networking_gloo_solo_io_v2.ExternalService,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range externalServiceList {
		s.Generic().Insert(obj)
	}
}

func (s *externalServiceSet) Has(externalService ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(externalService)
}

func (s *externalServiceSet) Equal(
	externalServiceSet ExternalServiceSet,
) bool {
	if s == nil {
		return externalServiceSet == nil
	}
	return s.Generic().Equal(externalServiceSet.Generic())
}

func (s *externalServiceSet) Delete(ExternalService ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(ExternalService)
}

func (s *externalServiceSet) Union(set ExternalServiceSet) ExternalServiceSet {
	if s == nil {
		return set
	}
	return NewExternalServiceSet(append(s.List(), set.List()...)...)
}

func (s *externalServiceSet) Difference(set ExternalServiceSet) ExternalServiceSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &externalServiceSet{set: newSet}
}

func (s *externalServiceSet) Intersection(set ExternalServiceSet) ExternalServiceSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var externalServiceList []*networking_gloo_solo_io_v2.ExternalService
	for _, obj := range newSet.List() {
		externalServiceList = append(externalServiceList, obj.(*networking_gloo_solo_io_v2.ExternalService))
	}
	return NewExternalServiceSet(externalServiceList...)
}

func (s *externalServiceSet) Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.ExternalService, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ExternalService %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_gloo_solo_io_v2.ExternalService{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_gloo_solo_io_v2.ExternalService), nil
}

func (s *externalServiceSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *externalServiceSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *externalServiceSet) Delta(newSet ExternalServiceSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *externalServiceSet) Clone() ExternalServiceSet {
	if s == nil {
		return nil
	}
	return &externalServiceSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type ExternalEndpointSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_gloo_solo_io_v2.ExternalEndpoint) bool) []*networking_gloo_solo_io_v2.ExternalEndpoint
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.ExternalEndpoint) bool) []*networking_gloo_solo_io_v2.ExternalEndpoint
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_gloo_solo_io_v2.ExternalEndpoint
	// Insert a resource into the set.
	Insert(externalEndpoint ...*networking_gloo_solo_io_v2.ExternalEndpoint)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(externalEndpointSet ExternalEndpointSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(externalEndpoint ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(externalEndpoint ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ExternalEndpointSet) ExternalEndpointSet
	// Return the difference with the provided set
	Difference(set ExternalEndpointSet) ExternalEndpointSet
	// Return the intersection with the provided set
	Intersection(set ExternalEndpointSet) ExternalEndpointSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.ExternalEndpoint, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another ExternalEndpointSet
	Delta(newSet ExternalEndpointSet) sksets.ResourceDelta
	// Create a deep copy of the current ExternalEndpointSet
	Clone() ExternalEndpointSet
}

func makeGenericExternalEndpointSet(externalEndpointList []*networking_gloo_solo_io_v2.ExternalEndpoint) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range externalEndpointList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type externalEndpointSet struct {
	set sksets.ResourceSet
}

func NewExternalEndpointSet(externalEndpointList ...*networking_gloo_solo_io_v2.ExternalEndpoint) ExternalEndpointSet {
	return &externalEndpointSet{set: makeGenericExternalEndpointSet(externalEndpointList)}
}

func NewExternalEndpointSetFromList(externalEndpointList *networking_gloo_solo_io_v2.ExternalEndpointList) ExternalEndpointSet {
	list := make([]*networking_gloo_solo_io_v2.ExternalEndpoint, 0, len(externalEndpointList.Items))
	for idx := range externalEndpointList.Items {
		list = append(list, &externalEndpointList.Items[idx])
	}
	return &externalEndpointSet{set: makeGenericExternalEndpointSet(list)}
}

func (s *externalEndpointSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *externalEndpointSet) List(filterResource ...func(*networking_gloo_solo_io_v2.ExternalEndpoint) bool) []*networking_gloo_solo_io_v2.ExternalEndpoint {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.ExternalEndpoint))
		})
	}

	objs := s.Generic().List(genericFilters...)
	externalEndpointList := make([]*networking_gloo_solo_io_v2.ExternalEndpoint, 0, len(objs))
	for _, obj := range objs {
		externalEndpointList = append(externalEndpointList, obj.(*networking_gloo_solo_io_v2.ExternalEndpoint))
	}
	return externalEndpointList
}

func (s *externalEndpointSet) UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.ExternalEndpoint) bool) []*networking_gloo_solo_io_v2.ExternalEndpoint {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.ExternalEndpoint))
		})
	}

	var externalEndpointList []*networking_gloo_solo_io_v2.ExternalEndpoint
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		externalEndpointList = append(externalEndpointList, obj.(*networking_gloo_solo_io_v2.ExternalEndpoint))
	}
	return externalEndpointList
}

func (s *externalEndpointSet) Map() map[string]*networking_gloo_solo_io_v2.ExternalEndpoint {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_gloo_solo_io_v2.ExternalEndpoint{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_gloo_solo_io_v2.ExternalEndpoint)
	}
	return newMap
}

func (s *externalEndpointSet) Insert(
	externalEndpointList ...*networking_gloo_solo_io_v2.ExternalEndpoint,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range externalEndpointList {
		s.Generic().Insert(obj)
	}
}

func (s *externalEndpointSet) Has(externalEndpoint ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(externalEndpoint)
}

func (s *externalEndpointSet) Equal(
	externalEndpointSet ExternalEndpointSet,
) bool {
	if s == nil {
		return externalEndpointSet == nil
	}
	return s.Generic().Equal(externalEndpointSet.Generic())
}

func (s *externalEndpointSet) Delete(ExternalEndpoint ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(ExternalEndpoint)
}

func (s *externalEndpointSet) Union(set ExternalEndpointSet) ExternalEndpointSet {
	if s == nil {
		return set
	}
	return NewExternalEndpointSet(append(s.List(), set.List()...)...)
}

func (s *externalEndpointSet) Difference(set ExternalEndpointSet) ExternalEndpointSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &externalEndpointSet{set: newSet}
}

func (s *externalEndpointSet) Intersection(set ExternalEndpointSet) ExternalEndpointSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var externalEndpointList []*networking_gloo_solo_io_v2.ExternalEndpoint
	for _, obj := range newSet.List() {
		externalEndpointList = append(externalEndpointList, obj.(*networking_gloo_solo_io_v2.ExternalEndpoint))
	}
	return NewExternalEndpointSet(externalEndpointList...)
}

func (s *externalEndpointSet) Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.ExternalEndpoint, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ExternalEndpoint %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_gloo_solo_io_v2.ExternalEndpoint{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_gloo_solo_io_v2.ExternalEndpoint), nil
}

func (s *externalEndpointSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *externalEndpointSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *externalEndpointSet) Delta(newSet ExternalEndpointSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *externalEndpointSet) Clone() ExternalEndpointSet {
	if s == nil {
		return nil
	}
	return &externalEndpointSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type RouteTableSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_gloo_solo_io_v2.RouteTable) bool) []*networking_gloo_solo_io_v2.RouteTable
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.RouteTable) bool) []*networking_gloo_solo_io_v2.RouteTable
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_gloo_solo_io_v2.RouteTable
	// Insert a resource into the set.
	Insert(routeTable ...*networking_gloo_solo_io_v2.RouteTable)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(routeTableSet RouteTableSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(routeTable ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(routeTable ezkube.ResourceId)
	// Return the union with the provided set
	Union(set RouteTableSet) RouteTableSet
	// Return the difference with the provided set
	Difference(set RouteTableSet) RouteTableSet
	// Return the intersection with the provided set
	Intersection(set RouteTableSet) RouteTableSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.RouteTable, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another RouteTableSet
	Delta(newSet RouteTableSet) sksets.ResourceDelta
	// Create a deep copy of the current RouteTableSet
	Clone() RouteTableSet
}

func makeGenericRouteTableSet(routeTableList []*networking_gloo_solo_io_v2.RouteTable) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range routeTableList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type routeTableSet struct {
	set sksets.ResourceSet
}

func NewRouteTableSet(routeTableList ...*networking_gloo_solo_io_v2.RouteTable) RouteTableSet {
	return &routeTableSet{set: makeGenericRouteTableSet(routeTableList)}
}

func NewRouteTableSetFromList(routeTableList *networking_gloo_solo_io_v2.RouteTableList) RouteTableSet {
	list := make([]*networking_gloo_solo_io_v2.RouteTable, 0, len(routeTableList.Items))
	for idx := range routeTableList.Items {
		list = append(list, &routeTableList.Items[idx])
	}
	return &routeTableSet{set: makeGenericRouteTableSet(list)}
}

func (s *routeTableSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *routeTableSet) List(filterResource ...func(*networking_gloo_solo_io_v2.RouteTable) bool) []*networking_gloo_solo_io_v2.RouteTable {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.RouteTable))
		})
	}

	objs := s.Generic().List(genericFilters...)
	routeTableList := make([]*networking_gloo_solo_io_v2.RouteTable, 0, len(objs))
	for _, obj := range objs {
		routeTableList = append(routeTableList, obj.(*networking_gloo_solo_io_v2.RouteTable))
	}
	return routeTableList
}

func (s *routeTableSet) UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.RouteTable) bool) []*networking_gloo_solo_io_v2.RouteTable {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.RouteTable))
		})
	}

	var routeTableList []*networking_gloo_solo_io_v2.RouteTable
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		routeTableList = append(routeTableList, obj.(*networking_gloo_solo_io_v2.RouteTable))
	}
	return routeTableList
}

func (s *routeTableSet) Map() map[string]*networking_gloo_solo_io_v2.RouteTable {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_gloo_solo_io_v2.RouteTable{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_gloo_solo_io_v2.RouteTable)
	}
	return newMap
}

func (s *routeTableSet) Insert(
	routeTableList ...*networking_gloo_solo_io_v2.RouteTable,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range routeTableList {
		s.Generic().Insert(obj)
	}
}

func (s *routeTableSet) Has(routeTable ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(routeTable)
}

func (s *routeTableSet) Equal(
	routeTableSet RouteTableSet,
) bool {
	if s == nil {
		return routeTableSet == nil
	}
	return s.Generic().Equal(routeTableSet.Generic())
}

func (s *routeTableSet) Delete(RouteTable ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(RouteTable)
}

func (s *routeTableSet) Union(set RouteTableSet) RouteTableSet {
	if s == nil {
		return set
	}
	return NewRouteTableSet(append(s.List(), set.List()...)...)
}

func (s *routeTableSet) Difference(set RouteTableSet) RouteTableSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &routeTableSet{set: newSet}
}

func (s *routeTableSet) Intersection(set RouteTableSet) RouteTableSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var routeTableList []*networking_gloo_solo_io_v2.RouteTable
	for _, obj := range newSet.List() {
		routeTableList = append(routeTableList, obj.(*networking_gloo_solo_io_v2.RouteTable))
	}
	return NewRouteTableSet(routeTableList...)
}

func (s *routeTableSet) Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.RouteTable, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find RouteTable %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_gloo_solo_io_v2.RouteTable{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_gloo_solo_io_v2.RouteTable), nil
}

func (s *routeTableSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *routeTableSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *routeTableSet) Delta(newSet RouteTableSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *routeTableSet) Clone() RouteTableSet {
	if s == nil {
		return nil
	}
	return &routeTableSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type VirtualDestinationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_gloo_solo_io_v2.VirtualDestination) bool) []*networking_gloo_solo_io_v2.VirtualDestination
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.VirtualDestination) bool) []*networking_gloo_solo_io_v2.VirtualDestination
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_gloo_solo_io_v2.VirtualDestination
	// Insert a resource into the set.
	Insert(virtualDestination ...*networking_gloo_solo_io_v2.VirtualDestination)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualDestinationSet VirtualDestinationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualDestination ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualDestination ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualDestinationSet) VirtualDestinationSet
	// Return the difference with the provided set
	Difference(set VirtualDestinationSet) VirtualDestinationSet
	// Return the intersection with the provided set
	Intersection(set VirtualDestinationSet) VirtualDestinationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.VirtualDestination, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another VirtualDestinationSet
	Delta(newSet VirtualDestinationSet) sksets.ResourceDelta
	// Create a deep copy of the current VirtualDestinationSet
	Clone() VirtualDestinationSet
}

func makeGenericVirtualDestinationSet(virtualDestinationList []*networking_gloo_solo_io_v2.VirtualDestination) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualDestinationList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualDestinationSet struct {
	set sksets.ResourceSet
}

func NewVirtualDestinationSet(virtualDestinationList ...*networking_gloo_solo_io_v2.VirtualDestination) VirtualDestinationSet {
	return &virtualDestinationSet{set: makeGenericVirtualDestinationSet(virtualDestinationList)}
}

func NewVirtualDestinationSetFromList(virtualDestinationList *networking_gloo_solo_io_v2.VirtualDestinationList) VirtualDestinationSet {
	list := make([]*networking_gloo_solo_io_v2.VirtualDestination, 0, len(virtualDestinationList.Items))
	for idx := range virtualDestinationList.Items {
		list = append(list, &virtualDestinationList.Items[idx])
	}
	return &virtualDestinationSet{set: makeGenericVirtualDestinationSet(list)}
}

func (s *virtualDestinationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *virtualDestinationSet) List(filterResource ...func(*networking_gloo_solo_io_v2.VirtualDestination) bool) []*networking_gloo_solo_io_v2.VirtualDestination {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.VirtualDestination))
		})
	}

	objs := s.Generic().List(genericFilters...)
	virtualDestinationList := make([]*networking_gloo_solo_io_v2.VirtualDestination, 0, len(objs))
	for _, obj := range objs {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_gloo_solo_io_v2.VirtualDestination))
	}
	return virtualDestinationList
}

func (s *virtualDestinationSet) UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.VirtualDestination) bool) []*networking_gloo_solo_io_v2.VirtualDestination {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.VirtualDestination))
		})
	}

	var virtualDestinationList []*networking_gloo_solo_io_v2.VirtualDestination
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_gloo_solo_io_v2.VirtualDestination))
	}
	return virtualDestinationList
}

func (s *virtualDestinationSet) Map() map[string]*networking_gloo_solo_io_v2.VirtualDestination {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_gloo_solo_io_v2.VirtualDestination{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_gloo_solo_io_v2.VirtualDestination)
	}
	return newMap
}

func (s *virtualDestinationSet) Insert(
	virtualDestinationList ...*networking_gloo_solo_io_v2.VirtualDestination,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualDestinationList {
		s.Generic().Insert(obj)
	}
}

func (s *virtualDestinationSet) Has(virtualDestination ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(virtualDestination)
}

func (s *virtualDestinationSet) Equal(
	virtualDestinationSet VirtualDestinationSet,
) bool {
	if s == nil {
		return virtualDestinationSet == nil
	}
	return s.Generic().Equal(virtualDestinationSet.Generic())
}

func (s *virtualDestinationSet) Delete(VirtualDestination ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(VirtualDestination)
}

func (s *virtualDestinationSet) Union(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return set
	}
	return NewVirtualDestinationSet(append(s.List(), set.List()...)...)
}

func (s *virtualDestinationSet) Difference(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &virtualDestinationSet{set: newSet}
}

func (s *virtualDestinationSet) Intersection(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var virtualDestinationList []*networking_gloo_solo_io_v2.VirtualDestination
	for _, obj := range newSet.List() {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_gloo_solo_io_v2.VirtualDestination))
	}
	return NewVirtualDestinationSet(virtualDestinationList...)
}

func (s *virtualDestinationSet) Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.VirtualDestination, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualDestination %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_gloo_solo_io_v2.VirtualDestination{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_gloo_solo_io_v2.VirtualDestination), nil
}

func (s *virtualDestinationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *virtualDestinationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *virtualDestinationSet) Delta(newSet VirtualDestinationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *virtualDestinationSet) Clone() VirtualDestinationSet {
	if s == nil {
		return nil
	}
	return &virtualDestinationSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type VirtualGatewaySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_gloo_solo_io_v2.VirtualGateway) bool) []*networking_gloo_solo_io_v2.VirtualGateway
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.VirtualGateway) bool) []*networking_gloo_solo_io_v2.VirtualGateway
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_gloo_solo_io_v2.VirtualGateway
	// Insert a resource into the set.
	Insert(virtualGateway ...*networking_gloo_solo_io_v2.VirtualGateway)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualGatewaySet VirtualGatewaySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualGateway ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualGateway ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualGatewaySet) VirtualGatewaySet
	// Return the difference with the provided set
	Difference(set VirtualGatewaySet) VirtualGatewaySet
	// Return the intersection with the provided set
	Intersection(set VirtualGatewaySet) VirtualGatewaySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.VirtualGateway, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another VirtualGatewaySet
	Delta(newSet VirtualGatewaySet) sksets.ResourceDelta
	// Create a deep copy of the current VirtualGatewaySet
	Clone() VirtualGatewaySet
}

func makeGenericVirtualGatewaySet(virtualGatewayList []*networking_gloo_solo_io_v2.VirtualGateway) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualGatewayList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualGatewaySet struct {
	set sksets.ResourceSet
}

func NewVirtualGatewaySet(virtualGatewayList ...*networking_gloo_solo_io_v2.VirtualGateway) VirtualGatewaySet {
	return &virtualGatewaySet{set: makeGenericVirtualGatewaySet(virtualGatewayList)}
}

func NewVirtualGatewaySetFromList(virtualGatewayList *networking_gloo_solo_io_v2.VirtualGatewayList) VirtualGatewaySet {
	list := make([]*networking_gloo_solo_io_v2.VirtualGateway, 0, len(virtualGatewayList.Items))
	for idx := range virtualGatewayList.Items {
		list = append(list, &virtualGatewayList.Items[idx])
	}
	return &virtualGatewaySet{set: makeGenericVirtualGatewaySet(list)}
}

func (s *virtualGatewaySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *virtualGatewaySet) List(filterResource ...func(*networking_gloo_solo_io_v2.VirtualGateway) bool) []*networking_gloo_solo_io_v2.VirtualGateway {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.VirtualGateway))
		})
	}

	objs := s.Generic().List(genericFilters...)
	virtualGatewayList := make([]*networking_gloo_solo_io_v2.VirtualGateway, 0, len(objs))
	for _, obj := range objs {
		virtualGatewayList = append(virtualGatewayList, obj.(*networking_gloo_solo_io_v2.VirtualGateway))
	}
	return virtualGatewayList
}

func (s *virtualGatewaySet) UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2.VirtualGateway) bool) []*networking_gloo_solo_io_v2.VirtualGateway {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2.VirtualGateway))
		})
	}

	var virtualGatewayList []*networking_gloo_solo_io_v2.VirtualGateway
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		virtualGatewayList = append(virtualGatewayList, obj.(*networking_gloo_solo_io_v2.VirtualGateway))
	}
	return virtualGatewayList
}

func (s *virtualGatewaySet) Map() map[string]*networking_gloo_solo_io_v2.VirtualGateway {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_gloo_solo_io_v2.VirtualGateway{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_gloo_solo_io_v2.VirtualGateway)
	}
	return newMap
}

func (s *virtualGatewaySet) Insert(
	virtualGatewayList ...*networking_gloo_solo_io_v2.VirtualGateway,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualGatewayList {
		s.Generic().Insert(obj)
	}
}

func (s *virtualGatewaySet) Has(virtualGateway ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(virtualGateway)
}

func (s *virtualGatewaySet) Equal(
	virtualGatewaySet VirtualGatewaySet,
) bool {
	if s == nil {
		return virtualGatewaySet == nil
	}
	return s.Generic().Equal(virtualGatewaySet.Generic())
}

func (s *virtualGatewaySet) Delete(VirtualGateway ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(VirtualGateway)
}

func (s *virtualGatewaySet) Union(set VirtualGatewaySet) VirtualGatewaySet {
	if s == nil {
		return set
	}
	return NewVirtualGatewaySet(append(s.List(), set.List()...)...)
}

func (s *virtualGatewaySet) Difference(set VirtualGatewaySet) VirtualGatewaySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &virtualGatewaySet{set: newSet}
}

func (s *virtualGatewaySet) Intersection(set VirtualGatewaySet) VirtualGatewaySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var virtualGatewayList []*networking_gloo_solo_io_v2.VirtualGateway
	for _, obj := range newSet.List() {
		virtualGatewayList = append(virtualGatewayList, obj.(*networking_gloo_solo_io_v2.VirtualGateway))
	}
	return NewVirtualGatewaySet(virtualGatewayList...)
}

func (s *virtualGatewaySet) Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2.VirtualGateway, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualGateway %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_gloo_solo_io_v2.VirtualGateway{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_gloo_solo_io_v2.VirtualGateway), nil
}

func (s *virtualGatewaySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *virtualGatewaySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *virtualGatewaySet) Delta(newSet VirtualGatewaySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *virtualGatewaySet) Clone() VirtualGatewaySet {
	if s == nil {
		return nil
	}
	return &virtualGatewaySet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
