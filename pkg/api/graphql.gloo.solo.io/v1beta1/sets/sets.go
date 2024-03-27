// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1beta1sets



import (
    graphql_gloo_solo_io_v1beta1 "github.com/solo-io/solo-apis/pkg/api/graphql.gloo.solo.io/v1beta1"

    "github.com/rotisserie/eris"
    sksets "github.com/solo-io/skv2/contrib/pkg/sets"
    "github.com/solo-io/skv2/pkg/ezkube"
    "k8s.io/apimachinery/pkg/util/sets"
)

type GraphQLApiSet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    List(filterResource ... func(*graphql_gloo_solo_io_v1beta1.GraphQLApi) bool) []*graphql_gloo_solo_io_v1beta1.GraphQLApi
    // Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    UnsortedList(filterResource ... func(*graphql_gloo_solo_io_v1beta1.GraphQLApi) bool) []*graphql_gloo_solo_io_v1beta1.GraphQLApi
    // Return the Set as a map of key to resource.
    Map() map[string]*graphql_gloo_solo_io_v1beta1.GraphQLApi
    // Insert a resource into the set.
    Insert(graphQLApi ...*graphql_gloo_solo_io_v1beta1.GraphQLApi)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(graphQLApiSet GraphQLApiSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(graphQLApi ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(graphQLApi  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set GraphQLApiSet) GraphQLApiSet
    // Return the difference with the provided set
    Difference(set GraphQLApiSet) GraphQLApiSet
    // Return the intersection with the provided set
    Intersection(set GraphQLApiSet) GraphQLApiSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*graphql_gloo_solo_io_v1beta1.GraphQLApi, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another GraphQLApiSet
    Delta(newSet GraphQLApiSet) sksets.ResourceDelta
    // Create a deep copy of the current GraphQLApiSet
    Clone() GraphQLApiSet
}

func makeGenericGraphQLApiSet(graphQLApiList []*graphql_gloo_solo_io_v1beta1.GraphQLApi) sksets.ResourceSet {
    var genericResources []ezkube.ResourceId
    for _, obj := range graphQLApiList {
        genericResources = append(genericResources, obj)
    }
    return sksets.NewResourceSet(genericResources...)
}

type graphQLApiSet struct {
    set sksets.ResourceSet
}

func NewGraphQLApiSet(graphQLApiList ...*graphql_gloo_solo_io_v1beta1.GraphQLApi) GraphQLApiSet {
    return &graphQLApiSet{set: makeGenericGraphQLApiSet(graphQLApiList)}
}

func NewGraphQLApiSetFromList(graphQLApiList *graphql_gloo_solo_io_v1beta1.GraphQLApiList) GraphQLApiSet {
    list := make([]*graphql_gloo_solo_io_v1beta1.GraphQLApi, 0, len(graphQLApiList.Items))
    for idx := range graphQLApiList.Items {
        list = append(list, &graphQLApiList.Items[idx])
    }
    return &graphQLApiSet{set: makeGenericGraphQLApiSet(list)}
}

func (s *graphQLApiSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *graphQLApiSet) List(filterResource ... func(*graphql_gloo_solo_io_v1beta1.GraphQLApi) bool) []*graphql_gloo_solo_io_v1beta1.GraphQLApi {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi))
        })
    }

    objs := s.Generic().List(genericFilters...)
    graphQLApiList := make([]*graphql_gloo_solo_io_v1beta1.GraphQLApi, 0, len(objs))
    for _, obj := range objs {
        graphQLApiList = append(graphQLApiList, obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi))
    }
    return graphQLApiList
}

func (s *graphQLApiSet) UnsortedList(filterResource ... func(*graphql_gloo_solo_io_v1beta1.GraphQLApi) bool) []*graphql_gloo_solo_io_v1beta1.GraphQLApi {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi))
        })
    }

    var graphQLApiList []*graphql_gloo_solo_io_v1beta1.GraphQLApi
    for _, obj := range s.Generic().UnsortedList(genericFilters...) {
        graphQLApiList = append(graphQLApiList, obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi))
    }
    return graphQLApiList
}

func (s *graphQLApiSet) Map() map[string]*graphql_gloo_solo_io_v1beta1.GraphQLApi {
    if s == nil {
        return nil
    }

    newMap := map[string]*graphql_gloo_solo_io_v1beta1.GraphQLApi{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*graphql_gloo_solo_io_v1beta1.GraphQLApi)
    }
    return newMap
}

func (s *graphQLApiSet) Insert(
        graphQLApiList ...*graphql_gloo_solo_io_v1beta1.GraphQLApi,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range graphQLApiList {
        s.Generic().Insert(obj)
    }
}

func (s *graphQLApiSet) Has(graphQLApi ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(graphQLApi)
}

func (s *graphQLApiSet) Equal(
        graphQLApiSet GraphQLApiSet,
) bool {
    if s == nil {
        return graphQLApiSet == nil
    }
    return s.Generic().Equal(graphQLApiSet.Generic())
}

func (s *graphQLApiSet) Delete(GraphQLApi ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(GraphQLApi)
}

func (s *graphQLApiSet) Union(set GraphQLApiSet) GraphQLApiSet {
    if s == nil {
        return set
    }
    return &graphQLApiMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *graphQLApiSet) Difference(set GraphQLApiSet) GraphQLApiSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &graphQLApiSet{set: newSet}
}

func (s *graphQLApiSet) Intersection(set GraphQLApiSet) GraphQLApiSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var graphQLApiList []*graphql_gloo_solo_io_v1beta1.GraphQLApi
    for _, obj := range newSet.List() {
        graphQLApiList = append(graphQLApiList, obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi))
    }
    return NewGraphQLApiSet(graphQLApiList...)
}


func (s *graphQLApiSet) Find(id ezkube.ResourceId) (*graphql_gloo_solo_io_v1beta1.GraphQLApi, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find GraphQLApi %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&graphql_gloo_solo_io_v1beta1.GraphQLApi{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi), nil
}

func (s *graphQLApiSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *graphQLApiSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *graphQLApiSet) Delta(newSet GraphQLApiSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

func (s *graphQLApiSet) Clone() GraphQLApiSet {
	if s == nil {
		return nil
	}
	return &graphQLApiMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type graphQLApiMergedSet struct {
    sets []sksets.ResourceSet
}

func NewGraphQLApiMergedSet(graphQLApiList ...*graphql_gloo_solo_io_v1beta1.GraphQLApi) GraphQLApiSet {
    return &graphQLApiMergedSet{sets: []sksets.ResourceSet{makeGenericGraphQLApiSet(graphQLApiList)}}
}

func NewGraphQLApiMergedSetFromList(graphQLApiList *graphql_gloo_solo_io_v1beta1.GraphQLApiList) GraphQLApiSet {
    list := make([]*graphql_gloo_solo_io_v1beta1.GraphQLApi, 0, len(graphQLApiList.Items))
    for idx := range graphQLApiList.Items {
        list = append(list, &graphQLApiList.Items[idx])
    }
    return &graphQLApiMergedSet{sets: []sksets.ResourceSet{makeGenericGraphQLApiSet(list)}}
}

func (s *graphQLApiMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
    }
    toRet := sets.String{}
	for _ , set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *graphQLApiMergedSet) List(filterResource ... func(*graphql_gloo_solo_io_v1beta1.GraphQLApi) bool) []*graphql_gloo_solo_io_v1beta1.GraphQLApi {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi))
        })
    }
   graphQLApiList := []*graphql_gloo_solo_io_v1beta1.GraphQLApi{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.List(genericFilters...) {
            if tracker[obj] {
				continue
			}
			tracker[obj] = true
			graphQLApiList = append(graphQLApiList, obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi))
		}
	}
    return graphQLApiList
}

func (s *graphQLApiMergedSet) UnsortedList(filterResource ... func(*graphql_gloo_solo_io_v1beta1.GraphQLApi) bool) []*graphql_gloo_solo_io_v1beta1.GraphQLApi {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi))
        })
    }
   graphQLApiList := []*graphql_gloo_solo_io_v1beta1.GraphQLApi{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.UnsortedList(genericFilters...) {
            if tracker[obj] {
				continue
			}
			tracker[obj] = true
			graphQLApiList = append(graphQLApiList, obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi))
		}
	}
    return graphQLApiList
}

func (s *graphQLApiMergedSet) Map() map[string]*graphql_gloo_solo_io_v1beta1.GraphQLApi {
    if s == nil {
        return nil
    }

    newMap := map[string]*graphql_gloo_solo_io_v1beta1.GraphQLApi{}
    for _, set := range s.sets {
		for k, v := range set.Map() {
            newMap[k] = v.(*graphql_gloo_solo_io_v1beta1.GraphQLApi)
        }
    }
    return newMap
}

func (s *graphQLApiMergedSet) Insert(
        graphQLApiList ...*graphql_gloo_solo_io_v1beta1.GraphQLApi,
) {
    if s == nil {
    }
    if len(s.sets) == 0 {
        s.sets = append(s.sets, makeGenericGraphQLApiSet(graphQLApiList))
    }
    for _, obj := range graphQLApiList {
        inserted := false
        for _, set := range s.sets {
            if set.Has(obj) {
                set.Insert(obj)
                inserted = true
                break
            }
        }
        if !inserted {
            s.sets[0].Insert(obj)
        }
    }
}

func (s *graphQLApiMergedSet) Has(graphQLApi ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    for _, set := range s.sets {
		if set.Has(graphQLApi) {
			return true
		}
	}
    return false
}

func (s *graphQLApiMergedSet) Equal(
        graphQLApiSet GraphQLApiSet,
) bool {
    panic("unimplemented")
}

func (s *graphQLApiMergedSet) Delete(GraphQLApi ezkube.ResourceId) {
    for _, set := range s.sets {
        set.Delete(GraphQLApi)
    }    
}

func (s *graphQLApiMergedSet) Union(set GraphQLApiSet) GraphQLApiSet {
    if s == nil {
        return set
    }
    return &graphQLApiMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *graphQLApiMergedSet) Difference(set GraphQLApiSet) GraphQLApiSet {
    panic("unimplemented")
}

func (s *graphQLApiMergedSet) Intersection(set GraphQLApiSet) GraphQLApiSet {
    panic("unimplemented")
}

func (s *graphQLApiMergedSet) Find(id ezkube.ResourceId) (*graphql_gloo_solo_io_v1beta1.GraphQLApi, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find GraphQLApi %v", sksets.Key(id))
    }

    var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&graphql_gloo_solo_io_v1beta1.GraphQLApi{}, id)
		if err == nil {
			return obj.(*graphql_gloo_solo_io_v1beta1.GraphQLApi), nil
		}
	}

    return nil, err
}

func (s *graphQLApiMergedSet) Length() int {
    if s == nil {
        return 0
    }
    totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
    return totalLen
}

func (s *graphQLApiMergedSet) Generic() sksets.ResourceSet {
    panic("unimplemented")
}

func (s *graphQLApiMergedSet) Delta(newSet GraphQLApiSet) sksets.ResourceDelta {
    panic("unimplemented")
}

func (s *graphQLApiMergedSet) Clone() GraphQLApiSet {
	if s == nil {
		return nil
	}
	return &graphQLApiMergedSet{sets: s.sets[:]}
}
