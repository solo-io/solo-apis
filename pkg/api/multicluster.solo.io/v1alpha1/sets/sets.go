// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	multicluster_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type MultiClusterRoleSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*multicluster_solo_io_v1alpha1.MultiClusterRole) bool) []*multicluster_solo_io_v1alpha1.MultiClusterRole
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*multicluster_solo_io_v1alpha1.MultiClusterRole) bool) []*multicluster_solo_io_v1alpha1.MultiClusterRole
	// Return the Set as a map of key to resource.
	Map() map[string]*multicluster_solo_io_v1alpha1.MultiClusterRole
	// Insert a resource into the set.
	Insert(multiClusterRole ...*multicluster_solo_io_v1alpha1.MultiClusterRole)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(multiClusterRoleSet MultiClusterRoleSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(multiClusterRole ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(multiClusterRole ezkube.ResourceId)
	// Return the union with the provided set
	Union(set MultiClusterRoleSet) MultiClusterRoleSet
	// Return the difference with the provided set
	Difference(set MultiClusterRoleSet) MultiClusterRoleSet
	// Return the intersection with the provided set
	Intersection(set MultiClusterRoleSet) MultiClusterRoleSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*multicluster_solo_io_v1alpha1.MultiClusterRole, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another MultiClusterRoleSet
	Delta(newSet MultiClusterRoleSet) sksets.ResourceDelta
	// Create a deep copy of the current MultiClusterRoleSet
	Clone() MultiClusterRoleSet
}

func makeGenericMultiClusterRoleSet(multiClusterRoleList []*multicluster_solo_io_v1alpha1.MultiClusterRole) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range multiClusterRoleList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type multiClusterRoleSet struct {
	set sksets.ResourceSet
}

func NewMultiClusterRoleSet(multiClusterRoleList ...*multicluster_solo_io_v1alpha1.MultiClusterRole) MultiClusterRoleSet {
	return &multiClusterRoleSet{set: makeGenericMultiClusterRoleSet(multiClusterRoleList)}
}

func NewMultiClusterRoleSetFromList(multiClusterRoleList *multicluster_solo_io_v1alpha1.MultiClusterRoleList) MultiClusterRoleSet {
	list := make([]*multicluster_solo_io_v1alpha1.MultiClusterRole, 0, len(multiClusterRoleList.Items))
	for idx := range multiClusterRoleList.Items {
		list = append(list, &multiClusterRoleList.Items[idx])
	}
	return &multiClusterRoleSet{set: makeGenericMultiClusterRoleSet(list)}
}

func (s *multiClusterRoleSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *multiClusterRoleSet) List(filterResource ...func(*multicluster_solo_io_v1alpha1.MultiClusterRole) bool) []*multicluster_solo_io_v1alpha1.MultiClusterRole {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*multicluster_solo_io_v1alpha1.MultiClusterRole))
		})
	}

	objs := s.Generic().List(genericFilters...)
	multiClusterRoleList := make([]*multicluster_solo_io_v1alpha1.MultiClusterRole, 0, len(objs))
	for _, obj := range objs {
		multiClusterRoleList = append(multiClusterRoleList, obj.(*multicluster_solo_io_v1alpha1.MultiClusterRole))
	}
	return multiClusterRoleList
}

func (s *multiClusterRoleSet) UnsortedList(filterResource ...func(*multicluster_solo_io_v1alpha1.MultiClusterRole) bool) []*multicluster_solo_io_v1alpha1.MultiClusterRole {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*multicluster_solo_io_v1alpha1.MultiClusterRole))
		})
	}

	var multiClusterRoleList []*multicluster_solo_io_v1alpha1.MultiClusterRole
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		multiClusterRoleList = append(multiClusterRoleList, obj.(*multicluster_solo_io_v1alpha1.MultiClusterRole))
	}
	return multiClusterRoleList
}

func (s *multiClusterRoleSet) Map() map[string]*multicluster_solo_io_v1alpha1.MultiClusterRole {
	if s == nil {
		return nil
	}

	newMap := map[string]*multicluster_solo_io_v1alpha1.MultiClusterRole{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*multicluster_solo_io_v1alpha1.MultiClusterRole)
	}
	return newMap
}

func (s *multiClusterRoleSet) Insert(
	multiClusterRoleList ...*multicluster_solo_io_v1alpha1.MultiClusterRole,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range multiClusterRoleList {
		s.Generic().Insert(obj)
	}
}

func (s *multiClusterRoleSet) Has(multiClusterRole ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(multiClusterRole)
}

func (s *multiClusterRoleSet) Equal(
	multiClusterRoleSet MultiClusterRoleSet,
) bool {
	if s == nil {
		return multiClusterRoleSet == nil
	}
	return s.Generic().Equal(multiClusterRoleSet.Generic())
}

func (s *multiClusterRoleSet) Delete(MultiClusterRole ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(MultiClusterRole)
}

func (s *multiClusterRoleSet) Union(set MultiClusterRoleSet) MultiClusterRoleSet {
	if s == nil {
		return set
	}
	return NewMultiClusterRoleSet(append(s.List(), set.List()...)...)
}

func (s *multiClusterRoleSet) Difference(set MultiClusterRoleSet) MultiClusterRoleSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &multiClusterRoleSet{set: newSet}
}

func (s *multiClusterRoleSet) Intersection(set MultiClusterRoleSet) MultiClusterRoleSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var multiClusterRoleList []*multicluster_solo_io_v1alpha1.MultiClusterRole
	for _, obj := range newSet.List() {
		multiClusterRoleList = append(multiClusterRoleList, obj.(*multicluster_solo_io_v1alpha1.MultiClusterRole))
	}
	return NewMultiClusterRoleSet(multiClusterRoleList...)
}

func (s *multiClusterRoleSet) Find(id ezkube.ResourceId) (*multicluster_solo_io_v1alpha1.MultiClusterRole, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find MultiClusterRole %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&multicluster_solo_io_v1alpha1.MultiClusterRole{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*multicluster_solo_io_v1alpha1.MultiClusterRole), nil
}

func (s *multiClusterRoleSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *multiClusterRoleSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *multiClusterRoleSet) Delta(newSet MultiClusterRoleSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *multiClusterRoleSet) Clone() MultiClusterRoleSet {
	if s == nil {
		return nil
	}
	return &multiClusterRoleSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
