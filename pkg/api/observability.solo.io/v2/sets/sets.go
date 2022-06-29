// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v2sets

import (
	observability_solo_io_v2 "github.com/solo-io/solo-apis/pkg/api/observability.solo.io/v2"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type AccessLogPolicySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*observability_solo_io_v2.AccessLogPolicy) bool) []*observability_solo_io_v2.AccessLogPolicy
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*observability_solo_io_v2.AccessLogPolicy) bool) []*observability_solo_io_v2.AccessLogPolicy
	// Return the Set as a map of key to resource.
	Map() map[string]*observability_solo_io_v2.AccessLogPolicy
	// Insert a resource into the set.
	Insert(accessLogPolicy ...*observability_solo_io_v2.AccessLogPolicy)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(accessLogPolicySet AccessLogPolicySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(accessLogPolicy ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(accessLogPolicy ezkube.ResourceId)
	// Return the union with the provided set
	Union(set AccessLogPolicySet) AccessLogPolicySet
	// Return the difference with the provided set
	Difference(set AccessLogPolicySet) AccessLogPolicySet
	// Return the intersection with the provided set
	Intersection(set AccessLogPolicySet) AccessLogPolicySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*observability_solo_io_v2.AccessLogPolicy, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another AccessLogPolicySet
	Delta(newSet AccessLogPolicySet) sksets.ResourceDelta
	// Create a deep copy of the current AccessLogPolicySet
	Clone() AccessLogPolicySet
}

func makeGenericAccessLogPolicySet(accessLogPolicyList []*observability_solo_io_v2.AccessLogPolicy) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range accessLogPolicyList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type accessLogPolicySet struct {
	set sksets.ResourceSet
}

func NewAccessLogPolicySet(accessLogPolicyList ...*observability_solo_io_v2.AccessLogPolicy) AccessLogPolicySet {
	return &accessLogPolicySet{set: makeGenericAccessLogPolicySet(accessLogPolicyList)}
}

func NewAccessLogPolicySetFromList(accessLogPolicyList *observability_solo_io_v2.AccessLogPolicyList) AccessLogPolicySet {
	list := make([]*observability_solo_io_v2.AccessLogPolicy, 0, len(accessLogPolicyList.Items))
	for idx := range accessLogPolicyList.Items {
		list = append(list, &accessLogPolicyList.Items[idx])
	}
	return &accessLogPolicySet{set: makeGenericAccessLogPolicySet(list)}
}

func (s *accessLogPolicySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *accessLogPolicySet) List(filterResource ...func(*observability_solo_io_v2.AccessLogPolicy) bool) []*observability_solo_io_v2.AccessLogPolicy {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*observability_solo_io_v2.AccessLogPolicy))
		})
	}

	objs := s.Generic().List(genericFilters...)
	accessLogPolicyList := make([]*observability_solo_io_v2.AccessLogPolicy, 0, len(objs))
	for _, obj := range objs {
		accessLogPolicyList = append(accessLogPolicyList, obj.(*observability_solo_io_v2.AccessLogPolicy))
	}
	return accessLogPolicyList
}

func (s *accessLogPolicySet) UnsortedList(filterResource ...func(*observability_solo_io_v2.AccessLogPolicy) bool) []*observability_solo_io_v2.AccessLogPolicy {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*observability_solo_io_v2.AccessLogPolicy))
		})
	}

	var accessLogPolicyList []*observability_solo_io_v2.AccessLogPolicy
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		accessLogPolicyList = append(accessLogPolicyList, obj.(*observability_solo_io_v2.AccessLogPolicy))
	}
	return accessLogPolicyList
}

func (s *accessLogPolicySet) Map() map[string]*observability_solo_io_v2.AccessLogPolicy {
	if s == nil {
		return nil
	}

	newMap := map[string]*observability_solo_io_v2.AccessLogPolicy{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*observability_solo_io_v2.AccessLogPolicy)
	}
	return newMap
}

func (s *accessLogPolicySet) Insert(
	accessLogPolicyList ...*observability_solo_io_v2.AccessLogPolicy,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range accessLogPolicyList {
		s.Generic().Insert(obj)
	}
}

func (s *accessLogPolicySet) Has(accessLogPolicy ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(accessLogPolicy)
}

func (s *accessLogPolicySet) Equal(
	accessLogPolicySet AccessLogPolicySet,
) bool {
	if s == nil {
		return accessLogPolicySet == nil
	}
	return s.Generic().Equal(accessLogPolicySet.Generic())
}

func (s *accessLogPolicySet) Delete(AccessLogPolicy ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(AccessLogPolicy)
}

func (s *accessLogPolicySet) Union(set AccessLogPolicySet) AccessLogPolicySet {
	if s == nil {
		return set
	}
	return NewAccessLogPolicySet(append(s.List(), set.List()...)...)
}

func (s *accessLogPolicySet) Difference(set AccessLogPolicySet) AccessLogPolicySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &accessLogPolicySet{set: newSet}
}

func (s *accessLogPolicySet) Intersection(set AccessLogPolicySet) AccessLogPolicySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var accessLogPolicyList []*observability_solo_io_v2.AccessLogPolicy
	for _, obj := range newSet.List() {
		accessLogPolicyList = append(accessLogPolicyList, obj.(*observability_solo_io_v2.AccessLogPolicy))
	}
	return NewAccessLogPolicySet(accessLogPolicyList...)
}

func (s *accessLogPolicySet) Find(id ezkube.ResourceId) (*observability_solo_io_v2.AccessLogPolicy, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find AccessLogPolicy %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&observability_solo_io_v2.AccessLogPolicy{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*observability_solo_io_v2.AccessLogPolicy), nil
}

func (s *accessLogPolicySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *accessLogPolicySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *accessLogPolicySet) Delta(newSet AccessLogPolicySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *accessLogPolicySet) Clone() AccessLogPolicySet {
	if s == nil {
		return nil
	}
	return &accessLogPolicySet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
