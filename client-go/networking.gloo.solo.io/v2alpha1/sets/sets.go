// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v2alpha1sets

import (
	networking_gloo_solo_io_v2alpha1 "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2alpha1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type ExternalWorkloadSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*networking_gloo_solo_io_v2alpha1.ExternalWorkload) bool) []*networking_gloo_solo_io_v2alpha1.ExternalWorkload
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2alpha1.ExternalWorkload) bool) []*networking_gloo_solo_io_v2alpha1.ExternalWorkload
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_gloo_solo_io_v2alpha1.ExternalWorkload
	// Insert a resource into the set.
	Insert(externalWorkload ...*networking_gloo_solo_io_v2alpha1.ExternalWorkload)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(externalWorkloadSet ExternalWorkloadSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(externalWorkload ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(externalWorkload ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ExternalWorkloadSet) ExternalWorkloadSet
	// Return the difference with the provided set
	Difference(set ExternalWorkloadSet) ExternalWorkloadSet
	// Return the intersection with the provided set
	Intersection(set ExternalWorkloadSet) ExternalWorkloadSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2alpha1.ExternalWorkload, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another ExternalWorkloadSet
	Delta(newSet ExternalWorkloadSet) sksets.ResourceDelta
	// Create a deep copy of the current ExternalWorkloadSet
	Clone() ExternalWorkloadSet
}

func makeGenericExternalWorkloadSet(externalWorkloadList []*networking_gloo_solo_io_v2alpha1.ExternalWorkload) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range externalWorkloadList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type externalWorkloadSet struct {
	set sksets.ResourceSet
}

func NewExternalWorkloadSet(externalWorkloadList ...*networking_gloo_solo_io_v2alpha1.ExternalWorkload) ExternalWorkloadSet {
	return &externalWorkloadSet{set: makeGenericExternalWorkloadSet(externalWorkloadList)}
}

func NewExternalWorkloadSetFromList(externalWorkloadList *networking_gloo_solo_io_v2alpha1.ExternalWorkloadList) ExternalWorkloadSet {
	list := make([]*networking_gloo_solo_io_v2alpha1.ExternalWorkload, 0, len(externalWorkloadList.Items))
	for idx := range externalWorkloadList.Items {
		list = append(list, &externalWorkloadList.Items[idx])
	}
	return &externalWorkloadSet{set: makeGenericExternalWorkloadSet(list)}
}

func (s *externalWorkloadSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *externalWorkloadSet) List(filterResource ...func(*networking_gloo_solo_io_v2alpha1.ExternalWorkload) bool) []*networking_gloo_solo_io_v2alpha1.ExternalWorkload {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2alpha1.ExternalWorkload))
		})
	}

	objs := s.Generic().List(genericFilters...)
	externalWorkloadList := make([]*networking_gloo_solo_io_v2alpha1.ExternalWorkload, 0, len(objs))
	for _, obj := range objs {
		externalWorkloadList = append(externalWorkloadList, obj.(*networking_gloo_solo_io_v2alpha1.ExternalWorkload))
	}
	return externalWorkloadList
}

func (s *externalWorkloadSet) UnsortedList(filterResource ...func(*networking_gloo_solo_io_v2alpha1.ExternalWorkload) bool) []*networking_gloo_solo_io_v2alpha1.ExternalWorkload {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_gloo_solo_io_v2alpha1.ExternalWorkload))
		})
	}

	var externalWorkloadList []*networking_gloo_solo_io_v2alpha1.ExternalWorkload
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		externalWorkloadList = append(externalWorkloadList, obj.(*networking_gloo_solo_io_v2alpha1.ExternalWorkload))
	}
	return externalWorkloadList
}

func (s *externalWorkloadSet) Map() map[string]*networking_gloo_solo_io_v2alpha1.ExternalWorkload {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_gloo_solo_io_v2alpha1.ExternalWorkload{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_gloo_solo_io_v2alpha1.ExternalWorkload)
	}
	return newMap
}

func (s *externalWorkloadSet) Insert(
	externalWorkloadList ...*networking_gloo_solo_io_v2alpha1.ExternalWorkload,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range externalWorkloadList {
		s.Generic().Insert(obj)
	}
}

func (s *externalWorkloadSet) Has(externalWorkload ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(externalWorkload)
}

func (s *externalWorkloadSet) Equal(
	externalWorkloadSet ExternalWorkloadSet,
) bool {
	if s == nil {
		return externalWorkloadSet == nil
	}
	return s.Generic().Equal(externalWorkloadSet.Generic())
}

func (s *externalWorkloadSet) Delete(ExternalWorkload ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(ExternalWorkload)
}

func (s *externalWorkloadSet) Union(set ExternalWorkloadSet) ExternalWorkloadSet {
	if s == nil {
		return set
	}
	return NewExternalWorkloadSet(append(s.List(), set.List()...)...)
}

func (s *externalWorkloadSet) Difference(set ExternalWorkloadSet) ExternalWorkloadSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &externalWorkloadSet{set: newSet}
}

func (s *externalWorkloadSet) Intersection(set ExternalWorkloadSet) ExternalWorkloadSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var externalWorkloadList []*networking_gloo_solo_io_v2alpha1.ExternalWorkload
	for _, obj := range newSet.List() {
		externalWorkloadList = append(externalWorkloadList, obj.(*networking_gloo_solo_io_v2alpha1.ExternalWorkload))
	}
	return NewExternalWorkloadSet(externalWorkloadList...)
}

func (s *externalWorkloadSet) Find(id ezkube.ResourceId) (*networking_gloo_solo_io_v2alpha1.ExternalWorkload, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ExternalWorkload %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_gloo_solo_io_v2alpha1.ExternalWorkload{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_gloo_solo_io_v2alpha1.ExternalWorkload), nil
}

func (s *externalWorkloadSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *externalWorkloadSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *externalWorkloadSet) Delta(newSet ExternalWorkloadSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *externalWorkloadSet) Clone() ExternalWorkloadSet {
	if s == nil {
		return nil
	}
	return &externalWorkloadSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}