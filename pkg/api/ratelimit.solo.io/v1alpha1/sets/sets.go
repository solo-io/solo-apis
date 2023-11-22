// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	ratelimit_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type RateLimitConfigSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*ratelimit_solo_io_v1alpha1.RateLimitConfig) bool) []*ratelimit_solo_io_v1alpha1.RateLimitConfig
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*ratelimit_solo_io_v1alpha1.RateLimitConfig) bool) []*ratelimit_solo_io_v1alpha1.RateLimitConfig
	// Return the Set as a map of key to resource.
	Map() map[string]*ratelimit_solo_io_v1alpha1.RateLimitConfig
	// Insert a resource into the set.
	Insert(rateLimitConfig ...*ratelimit_solo_io_v1alpha1.RateLimitConfig)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(rateLimitConfigSet RateLimitConfigSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(rateLimitConfig ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(rateLimitConfig ezkube.ResourceId)
	// Return the union with the provided set
	Union(set RateLimitConfigSet) RateLimitConfigSet
	// Return the difference with the provided set
	Difference(set RateLimitConfigSet) RateLimitConfigSet
	// Return the intersection with the provided set
	Intersection(set RateLimitConfigSet) RateLimitConfigSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*ratelimit_solo_io_v1alpha1.RateLimitConfig, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another RateLimitConfigSet
	Delta(newSet RateLimitConfigSet) sksets.ResourceDelta
	// Create a deep copy of the current RateLimitConfigSet
	Clone() RateLimitConfigSet
}

func makeGenericRateLimitConfigSet(rateLimitConfigList []*ratelimit_solo_io_v1alpha1.RateLimitConfig) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range rateLimitConfigList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type rateLimitConfigSet struct {
	set sksets.ResourceSet
}

func NewRateLimitConfigSet(rateLimitConfigList ...*ratelimit_solo_io_v1alpha1.RateLimitConfig) RateLimitConfigSet {
	return &rateLimitConfigSet{set: makeGenericRateLimitConfigSet(rateLimitConfigList)}
}

func NewRateLimitConfigSetFromList(rateLimitConfigList *ratelimit_solo_io_v1alpha1.RateLimitConfigList) RateLimitConfigSet {
	list := make([]*ratelimit_solo_io_v1alpha1.RateLimitConfig, 0, len(rateLimitConfigList.Items))
	for idx := range rateLimitConfigList.Items {
		list = append(list, &rateLimitConfigList.Items[idx])
	}
	return &rateLimitConfigSet{set: makeGenericRateLimitConfigSet(list)}
}

func (s *rateLimitConfigSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *rateLimitConfigSet) List(filterResource ...func(*ratelimit_solo_io_v1alpha1.RateLimitConfig) bool) []*ratelimit_solo_io_v1alpha1.RateLimitConfig {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*ratelimit_solo_io_v1alpha1.RateLimitConfig))
		})
	}

	objs := s.Generic().List(genericFilters...)
	rateLimitConfigList := make([]*ratelimit_solo_io_v1alpha1.RateLimitConfig, 0, len(objs))
	for _, obj := range objs {
		rateLimitConfigList = append(rateLimitConfigList, obj.(*ratelimit_solo_io_v1alpha1.RateLimitConfig))
	}
	return rateLimitConfigList
}

func (s *rateLimitConfigSet) UnsortedList(filterResource ...func(*ratelimit_solo_io_v1alpha1.RateLimitConfig) bool) []*ratelimit_solo_io_v1alpha1.RateLimitConfig {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*ratelimit_solo_io_v1alpha1.RateLimitConfig))
		})
	}

	var rateLimitConfigList []*ratelimit_solo_io_v1alpha1.RateLimitConfig
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		rateLimitConfigList = append(rateLimitConfigList, obj.(*ratelimit_solo_io_v1alpha1.RateLimitConfig))
	}
	return rateLimitConfigList
}

func (s *rateLimitConfigSet) Map() map[string]*ratelimit_solo_io_v1alpha1.RateLimitConfig {
	if s == nil {
		return nil
	}

	newMap := map[string]*ratelimit_solo_io_v1alpha1.RateLimitConfig{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*ratelimit_solo_io_v1alpha1.RateLimitConfig)
	}
	return newMap
}

func (s *rateLimitConfigSet) Insert(
	rateLimitConfigList ...*ratelimit_solo_io_v1alpha1.RateLimitConfig,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range rateLimitConfigList {
		s.Generic().Insert(obj)
	}
}

func (s *rateLimitConfigSet) Has(rateLimitConfig ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(rateLimitConfig)
}

func (s *rateLimitConfigSet) Equal(
	rateLimitConfigSet RateLimitConfigSet,
) bool {
	if s == nil {
		return rateLimitConfigSet == nil
	}
	return s.Generic().Equal(rateLimitConfigSet.Generic())
}

func (s *rateLimitConfigSet) Delete(RateLimitConfig ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(RateLimitConfig)
}

func (s *rateLimitConfigSet) Union(set RateLimitConfigSet) RateLimitConfigSet {
	if s == nil {
		return set
	}
	return NewRateLimitConfigSet(append(s.List(), set.List()...)...)
}

func (s *rateLimitConfigSet) Difference(set RateLimitConfigSet) RateLimitConfigSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &rateLimitConfigSet{set: newSet}
}

func (s *rateLimitConfigSet) Intersection(set RateLimitConfigSet) RateLimitConfigSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var rateLimitConfigList []*ratelimit_solo_io_v1alpha1.RateLimitConfig
	for _, obj := range newSet.List() {
		rateLimitConfigList = append(rateLimitConfigList, obj.(*ratelimit_solo_io_v1alpha1.RateLimitConfig))
	}
	return NewRateLimitConfigSet(rateLimitConfigList...)
}

func (s *rateLimitConfigSet) Find(id ezkube.ResourceId) (*ratelimit_solo_io_v1alpha1.RateLimitConfig, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find RateLimitConfig %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&ratelimit_solo_io_v1alpha1.RateLimitConfig{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*ratelimit_solo_io_v1alpha1.RateLimitConfig), nil
}

func (s *rateLimitConfigSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *rateLimitConfigSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *rateLimitConfigSet) Delta(newSet RateLimitConfigSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *rateLimitConfigSet) Clone() RateLimitConfigSet {
	if s == nil {
		return nil
	}
	return &rateLimitConfigSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
