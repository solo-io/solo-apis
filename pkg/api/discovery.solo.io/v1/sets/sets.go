// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	discovery_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/discovery.solo.io/v1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type DestinationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*discovery_solo_io_v1.Destination) bool) []*discovery_solo_io_v1.Destination
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*discovery_solo_io_v1.Destination) bool) []*discovery_solo_io_v1.Destination
	// Return the Set as a map of key to resource.
	Map() map[string]*discovery_solo_io_v1.Destination
	// Insert a resource into the set.
	Insert(destination ...*discovery_solo_io_v1.Destination)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(destinationSet DestinationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(destination ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(destination ezkube.ResourceId)
	// Return the union with the provided set
	Union(set DestinationSet) DestinationSet
	// Return the difference with the provided set
	Difference(set DestinationSet) DestinationSet
	// Return the intersection with the provided set
	Intersection(set DestinationSet) DestinationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*discovery_solo_io_v1.Destination, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another DestinationSet
	Delta(newSet DestinationSet) sksets.ResourceDelta
	// Create a deep copy of the current DestinationSet
	Clone() DestinationSet
}

func makeGenericDestinationSet(destinationList []*discovery_solo_io_v1.Destination) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range destinationList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type destinationSet struct {
	set sksets.ResourceSet
}

func NewDestinationSet(destinationList ...*discovery_solo_io_v1.Destination) DestinationSet {
	return &destinationSet{set: makeGenericDestinationSet(destinationList)}
}

func NewDestinationSetFromList(destinationList *discovery_solo_io_v1.DestinationList) DestinationSet {
	list := make([]*discovery_solo_io_v1.Destination, 0, len(destinationList.Items))
	for idx := range destinationList.Items {
		list = append(list, &destinationList.Items[idx])
	}
	return &destinationSet{set: makeGenericDestinationSet(list)}
}

func (s *destinationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *destinationSet) List(filterResource ...func(*discovery_solo_io_v1.Destination) bool) []*discovery_solo_io_v1.Destination {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*discovery_solo_io_v1.Destination))
		})
	}

	objs := s.Generic().List(genericFilters...)
	destinationList := make([]*discovery_solo_io_v1.Destination, 0, len(objs))
	for _, obj := range objs {
		destinationList = append(destinationList, obj.(*discovery_solo_io_v1.Destination))
	}
	return destinationList
}

func (s *destinationSet) UnsortedList(filterResource ...func(*discovery_solo_io_v1.Destination) bool) []*discovery_solo_io_v1.Destination {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*discovery_solo_io_v1.Destination))
		})
	}

	var destinationList []*discovery_solo_io_v1.Destination
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		destinationList = append(destinationList, obj.(*discovery_solo_io_v1.Destination))
	}
	return destinationList
}

func (s *destinationSet) Map() map[string]*discovery_solo_io_v1.Destination {
	if s == nil {
		return nil
	}

	newMap := map[string]*discovery_solo_io_v1.Destination{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*discovery_solo_io_v1.Destination)
	}
	return newMap
}

func (s *destinationSet) Insert(
	destinationList ...*discovery_solo_io_v1.Destination,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range destinationList {
		s.Generic().Insert(obj)
	}
}

func (s *destinationSet) Has(destination ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(destination)
}

func (s *destinationSet) Equal(
	destinationSet DestinationSet,
) bool {
	if s == nil {
		return destinationSet == nil
	}
	return s.Generic().Equal(destinationSet.Generic())
}

func (s *destinationSet) Delete(Destination ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(Destination)
}

func (s *destinationSet) Union(set DestinationSet) DestinationSet {
	if s == nil {
		return set
	}
	return NewDestinationSet(append(s.List(), set.List()...)...)
}

func (s *destinationSet) Difference(set DestinationSet) DestinationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &destinationSet{set: newSet}
}

func (s *destinationSet) Intersection(set DestinationSet) DestinationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var destinationList []*discovery_solo_io_v1.Destination
	for _, obj := range newSet.List() {
		destinationList = append(destinationList, obj.(*discovery_solo_io_v1.Destination))
	}
	return NewDestinationSet(destinationList...)
}

func (s *destinationSet) Find(id ezkube.ResourceId) (*discovery_solo_io_v1.Destination, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Destination %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&discovery_solo_io_v1.Destination{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*discovery_solo_io_v1.Destination), nil
}

func (s *destinationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *destinationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *destinationSet) Delta(newSet DestinationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *destinationSet) Clone() DestinationSet {
	if s == nil {
		return nil
	}
	return &destinationSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type WorkloadSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*discovery_solo_io_v1.Workload) bool) []*discovery_solo_io_v1.Workload
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*discovery_solo_io_v1.Workload) bool) []*discovery_solo_io_v1.Workload
	// Return the Set as a map of key to resource.
	Map() map[string]*discovery_solo_io_v1.Workload
	// Insert a resource into the set.
	Insert(workload ...*discovery_solo_io_v1.Workload)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(workloadSet WorkloadSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(workload ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(workload ezkube.ResourceId)
	// Return the union with the provided set
	Union(set WorkloadSet) WorkloadSet
	// Return the difference with the provided set
	Difference(set WorkloadSet) WorkloadSet
	// Return the intersection with the provided set
	Intersection(set WorkloadSet) WorkloadSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*discovery_solo_io_v1.Workload, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another WorkloadSet
	Delta(newSet WorkloadSet) sksets.ResourceDelta
	// Create a deep copy of the current WorkloadSet
	Clone() WorkloadSet
}

func makeGenericWorkloadSet(workloadList []*discovery_solo_io_v1.Workload) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range workloadList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type workloadSet struct {
	set sksets.ResourceSet
}

func NewWorkloadSet(workloadList ...*discovery_solo_io_v1.Workload) WorkloadSet {
	return &workloadSet{set: makeGenericWorkloadSet(workloadList)}
}

func NewWorkloadSetFromList(workloadList *discovery_solo_io_v1.WorkloadList) WorkloadSet {
	list := make([]*discovery_solo_io_v1.Workload, 0, len(workloadList.Items))
	for idx := range workloadList.Items {
		list = append(list, &workloadList.Items[idx])
	}
	return &workloadSet{set: makeGenericWorkloadSet(list)}
}

func (s *workloadSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *workloadSet) List(filterResource ...func(*discovery_solo_io_v1.Workload) bool) []*discovery_solo_io_v1.Workload {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*discovery_solo_io_v1.Workload))
		})
	}

	objs := s.Generic().List(genericFilters...)
	workloadList := make([]*discovery_solo_io_v1.Workload, 0, len(objs))
	for _, obj := range objs {
		workloadList = append(workloadList, obj.(*discovery_solo_io_v1.Workload))
	}
	return workloadList
}

func (s *workloadSet) UnsortedList(filterResource ...func(*discovery_solo_io_v1.Workload) bool) []*discovery_solo_io_v1.Workload {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*discovery_solo_io_v1.Workload))
		})
	}

	var workloadList []*discovery_solo_io_v1.Workload
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		workloadList = append(workloadList, obj.(*discovery_solo_io_v1.Workload))
	}
	return workloadList
}

func (s *workloadSet) Map() map[string]*discovery_solo_io_v1.Workload {
	if s == nil {
		return nil
	}

	newMap := map[string]*discovery_solo_io_v1.Workload{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*discovery_solo_io_v1.Workload)
	}
	return newMap
}

func (s *workloadSet) Insert(
	workloadList ...*discovery_solo_io_v1.Workload,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range workloadList {
		s.Generic().Insert(obj)
	}
}

func (s *workloadSet) Has(workload ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(workload)
}

func (s *workloadSet) Equal(
	workloadSet WorkloadSet,
) bool {
	if s == nil {
		return workloadSet == nil
	}
	return s.Generic().Equal(workloadSet.Generic())
}

func (s *workloadSet) Delete(Workload ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(Workload)
}

func (s *workloadSet) Union(set WorkloadSet) WorkloadSet {
	if s == nil {
		return set
	}
	return NewWorkloadSet(append(s.List(), set.List()...)...)
}

func (s *workloadSet) Difference(set WorkloadSet) WorkloadSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &workloadSet{set: newSet}
}

func (s *workloadSet) Intersection(set WorkloadSet) WorkloadSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var workloadList []*discovery_solo_io_v1.Workload
	for _, obj := range newSet.List() {
		workloadList = append(workloadList, obj.(*discovery_solo_io_v1.Workload))
	}
	return NewWorkloadSet(workloadList...)
}

func (s *workloadSet) Find(id ezkube.ResourceId) (*discovery_solo_io_v1.Workload, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Workload %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&discovery_solo_io_v1.Workload{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*discovery_solo_io_v1.Workload), nil
}

func (s *workloadSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *workloadSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *workloadSet) Delta(newSet WorkloadSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *workloadSet) Clone() WorkloadSet {
	if s == nil {
		return nil
	}
	return &workloadSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type MeshSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*discovery_solo_io_v1.Mesh) bool) []*discovery_solo_io_v1.Mesh
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*discovery_solo_io_v1.Mesh) bool) []*discovery_solo_io_v1.Mesh
	// Return the Set as a map of key to resource.
	Map() map[string]*discovery_solo_io_v1.Mesh
	// Insert a resource into the set.
	Insert(mesh ...*discovery_solo_io_v1.Mesh)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(meshSet MeshSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(mesh ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(mesh ezkube.ResourceId)
	// Return the union with the provided set
	Union(set MeshSet) MeshSet
	// Return the difference with the provided set
	Difference(set MeshSet) MeshSet
	// Return the intersection with the provided set
	Intersection(set MeshSet) MeshSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*discovery_solo_io_v1.Mesh, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another MeshSet
	Delta(newSet MeshSet) sksets.ResourceDelta
	// Create a deep copy of the current MeshSet
	Clone() MeshSet
}

func makeGenericMeshSet(meshList []*discovery_solo_io_v1.Mesh) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range meshList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type meshSet struct {
	set sksets.ResourceSet
}

func NewMeshSet(meshList ...*discovery_solo_io_v1.Mesh) MeshSet {
	return &meshSet{set: makeGenericMeshSet(meshList)}
}

func NewMeshSetFromList(meshList *discovery_solo_io_v1.MeshList) MeshSet {
	list := make([]*discovery_solo_io_v1.Mesh, 0, len(meshList.Items))
	for idx := range meshList.Items {
		list = append(list, &meshList.Items[idx])
	}
	return &meshSet{set: makeGenericMeshSet(list)}
}

func (s *meshSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *meshSet) List(filterResource ...func(*discovery_solo_io_v1.Mesh) bool) []*discovery_solo_io_v1.Mesh {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*discovery_solo_io_v1.Mesh))
		})
	}

	objs := s.Generic().List(genericFilters...)
	meshList := make([]*discovery_solo_io_v1.Mesh, 0, len(objs))
	for _, obj := range objs {
		meshList = append(meshList, obj.(*discovery_solo_io_v1.Mesh))
	}
	return meshList
}

func (s *meshSet) UnsortedList(filterResource ...func(*discovery_solo_io_v1.Mesh) bool) []*discovery_solo_io_v1.Mesh {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*discovery_solo_io_v1.Mesh))
		})
	}

	var meshList []*discovery_solo_io_v1.Mesh
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		meshList = append(meshList, obj.(*discovery_solo_io_v1.Mesh))
	}
	return meshList
}

func (s *meshSet) Map() map[string]*discovery_solo_io_v1.Mesh {
	if s == nil {
		return nil
	}

	newMap := map[string]*discovery_solo_io_v1.Mesh{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*discovery_solo_io_v1.Mesh)
	}
	return newMap
}

func (s *meshSet) Insert(
	meshList ...*discovery_solo_io_v1.Mesh,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range meshList {
		s.Generic().Insert(obj)
	}
}

func (s *meshSet) Has(mesh ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(mesh)
}

func (s *meshSet) Equal(
	meshSet MeshSet,
) bool {
	if s == nil {
		return meshSet == nil
	}
	return s.Generic().Equal(meshSet.Generic())
}

func (s *meshSet) Delete(Mesh ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(Mesh)
}

func (s *meshSet) Union(set MeshSet) MeshSet {
	if s == nil {
		return set
	}
	return NewMeshSet(append(s.List(), set.List()...)...)
}

func (s *meshSet) Difference(set MeshSet) MeshSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &meshSet{set: newSet}
}

func (s *meshSet) Intersection(set MeshSet) MeshSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var meshList []*discovery_solo_io_v1.Mesh
	for _, obj := range newSet.List() {
		meshList = append(meshList, obj.(*discovery_solo_io_v1.Mesh))
	}
	return NewMeshSet(meshList...)
}

func (s *meshSet) Find(id ezkube.ResourceId) (*discovery_solo_io_v1.Mesh, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Mesh %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&discovery_solo_io_v1.Mesh{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*discovery_solo_io_v1.Mesh), nil
}

func (s *meshSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *meshSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *meshSet) Delta(newSet MeshSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *meshSet) Clone() MeshSet {
	if s == nil {
		return nil
	}
	return &meshSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
