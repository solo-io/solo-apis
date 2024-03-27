// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets



import (
    fed_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1"

    "github.com/rotisserie/eris"
    sksets "github.com/solo-io/skv2/contrib/pkg/sets"
    "github.com/solo-io/skv2/pkg/ezkube"
    "k8s.io/apimachinery/pkg/util/sets"
)

type GlooInstanceSet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    List(filterResource ... func(*fed_solo_io_v1.GlooInstance) bool) []*fed_solo_io_v1.GlooInstance
    // Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    UnsortedList(filterResource ... func(*fed_solo_io_v1.GlooInstance) bool) []*fed_solo_io_v1.GlooInstance
    // Return the Set as a map of key to resource.
    Map() map[string]*fed_solo_io_v1.GlooInstance
    // Insert a resource into the set.
    Insert(glooInstance ...*fed_solo_io_v1.GlooInstance)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(glooInstanceSet GlooInstanceSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(glooInstance ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(glooInstance  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set GlooInstanceSet) GlooInstanceSet
    // Return the difference with the provided set
    Difference(set GlooInstanceSet) GlooInstanceSet
    // Return the intersection with the provided set
    Intersection(set GlooInstanceSet) GlooInstanceSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*fed_solo_io_v1.GlooInstance, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another GlooInstanceSet
    Delta(newSet GlooInstanceSet) sksets.ResourceDelta
    // Create a deep copy of the current GlooInstanceSet
    Clone() GlooInstanceSet
}

func makeGenericGlooInstanceSet(glooInstanceList []*fed_solo_io_v1.GlooInstance) sksets.ResourceSet {
    var genericResources []ezkube.ResourceId
    for _, obj := range glooInstanceList {
        genericResources = append(genericResources, obj)
    }
    return sksets.NewResourceSet(genericResources...)
}

type glooInstanceSet struct {
    set sksets.ResourceSet
}

func NewGlooInstanceSet(glooInstanceList ...*fed_solo_io_v1.GlooInstance) GlooInstanceSet {
    return &glooInstanceSet{set: makeGenericGlooInstanceSet(glooInstanceList)}
}

func NewGlooInstanceSetFromList(glooInstanceList *fed_solo_io_v1.GlooInstanceList) GlooInstanceSet {
    list := make([]*fed_solo_io_v1.GlooInstance, 0, len(glooInstanceList.Items))
    for idx := range glooInstanceList.Items {
        list = append(list, &glooInstanceList.Items[idx])
    }
    return &glooInstanceSet{set: makeGenericGlooInstanceSet(list)}
}

func (s *glooInstanceSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *glooInstanceSet) List(filterResource ... func(*fed_solo_io_v1.GlooInstance) bool) []*fed_solo_io_v1.GlooInstance {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_solo_io_v1.GlooInstance))
        })
    }

    objs := s.Generic().List(genericFilters...)
    glooInstanceList := make([]*fed_solo_io_v1.GlooInstance, 0, len(objs))
    for _, obj := range objs {
        glooInstanceList = append(glooInstanceList, obj.(*fed_solo_io_v1.GlooInstance))
    }
    return glooInstanceList
}

func (s *glooInstanceSet) UnsortedList(filterResource ... func(*fed_solo_io_v1.GlooInstance) bool) []*fed_solo_io_v1.GlooInstance {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_solo_io_v1.GlooInstance))
        })
    }

    var glooInstanceList []*fed_solo_io_v1.GlooInstance
    for _, obj := range s.Generic().UnsortedList(genericFilters...) {
        glooInstanceList = append(glooInstanceList, obj.(*fed_solo_io_v1.GlooInstance))
    }
    return glooInstanceList
}

func (s *glooInstanceSet) Map() map[string]*fed_solo_io_v1.GlooInstance {
    if s == nil {
        return nil
    }

    newMap := map[string]*fed_solo_io_v1.GlooInstance{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*fed_solo_io_v1.GlooInstance)
    }
    return newMap
}

func (s *glooInstanceSet) Insert(
        glooInstanceList ...*fed_solo_io_v1.GlooInstance,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range glooInstanceList {
        s.Generic().Insert(obj)
    }
}

func (s *glooInstanceSet) Has(glooInstance ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(glooInstance)
}

func (s *glooInstanceSet) Equal(
        glooInstanceSet GlooInstanceSet,
) bool {
    if s == nil {
        return glooInstanceSet == nil
    }
    return s.Generic().Equal(glooInstanceSet.Generic())
}

func (s *glooInstanceSet) Delete(GlooInstance ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(GlooInstance)
}

func (s *glooInstanceSet) Union(set GlooInstanceSet) GlooInstanceSet {
    if s == nil {
        return set
    }
    return &glooInstanceMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *glooInstanceSet) Difference(set GlooInstanceSet) GlooInstanceSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &glooInstanceSet{set: newSet}
}

func (s *glooInstanceSet) Intersection(set GlooInstanceSet) GlooInstanceSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var glooInstanceList []*fed_solo_io_v1.GlooInstance
    for _, obj := range newSet.List() {
        glooInstanceList = append(glooInstanceList, obj.(*fed_solo_io_v1.GlooInstance))
    }
    return NewGlooInstanceSet(glooInstanceList...)
}


func (s *glooInstanceSet) Find(id ezkube.ResourceId) (*fed_solo_io_v1.GlooInstance, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find GlooInstance %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&fed_solo_io_v1.GlooInstance{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*fed_solo_io_v1.GlooInstance), nil
}

func (s *glooInstanceSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *glooInstanceSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *glooInstanceSet) Delta(newSet GlooInstanceSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

func (s *glooInstanceSet) Clone() GlooInstanceSet {
	if s == nil {
		return nil
	}
	return &glooInstanceMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type glooInstanceMergedSet struct {
    sets []sksets.ResourceSet
}

func NewGlooInstanceMergedSet(glooInstanceList ...*fed_solo_io_v1.GlooInstance) GlooInstanceSet {
    return &glooInstanceMergedSet{sets: []sksets.ResourceSet{makeGenericGlooInstanceSet(glooInstanceList)}}
}

func NewGlooInstanceMergedSetFromList(glooInstanceList *fed_solo_io_v1.GlooInstanceList) GlooInstanceSet {
    list := make([]*fed_solo_io_v1.GlooInstance, 0, len(glooInstanceList.Items))
    for idx := range glooInstanceList.Items {
        list = append(list, &glooInstanceList.Items[idx])
    }
    return &glooInstanceMergedSet{sets: []sksets.ResourceSet{makeGenericGlooInstanceSet(list)}}
}

func (s *glooInstanceMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
    }
    toRet := sets.String{}
	for _ , set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *glooInstanceMergedSet) List(filterResource ... func(*fed_solo_io_v1.GlooInstance) bool) []*fed_solo_io_v1.GlooInstance {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_solo_io_v1.GlooInstance))
        })
    }
   glooInstanceList := []*fed_solo_io_v1.GlooInstance{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.List(genericFilters...) {
            if tracker[obj] {
				continue
			}
			tracker[obj] = true
			glooInstanceList = append(glooInstanceList, obj.(*fed_solo_io_v1.GlooInstance))
		}
	}
    return glooInstanceList
}

func (s *glooInstanceMergedSet) UnsortedList(filterResource ... func(*fed_solo_io_v1.GlooInstance) bool) []*fed_solo_io_v1.GlooInstance {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_solo_io_v1.GlooInstance))
        })
    }
   glooInstanceList := []*fed_solo_io_v1.GlooInstance{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.UnsortedList(genericFilters...) {
            if tracker[obj] {
				continue
			}
			tracker[obj] = true
			glooInstanceList = append(glooInstanceList, obj.(*fed_solo_io_v1.GlooInstance))
		}
	}
    return glooInstanceList
}

func (s *glooInstanceMergedSet) Map() map[string]*fed_solo_io_v1.GlooInstance {
    if s == nil {
        return nil
    }

    newMap := map[string]*fed_solo_io_v1.GlooInstance{}
    for _, set := range s.sets {
		for k, v := range set.Map() {
            newMap[k] = v.(*fed_solo_io_v1.GlooInstance)
        }
    }
    return newMap
}

func (s *glooInstanceMergedSet) Insert(
        glooInstanceList ...*fed_solo_io_v1.GlooInstance,
) {
    if s == nil {
    }
    if len(s.sets) == 0 {
        s.sets = append(s.sets, makeGenericGlooInstanceSet(glooInstanceList))
    }
    for _, obj := range glooInstanceList {
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

func (s *glooInstanceMergedSet) Has(glooInstance ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    for _, set := range s.sets {
		if set.Has(glooInstance) {
			return true
		}
	}
    return false
}

func (s *glooInstanceMergedSet) Equal(
        glooInstanceSet GlooInstanceSet,
) bool {
    panic("unimplemented")
}

func (s *glooInstanceMergedSet) Delete(GlooInstance ezkube.ResourceId) {
    for _, set := range s.sets {
        set.Delete(GlooInstance)
    }    
}

func (s *glooInstanceMergedSet) Union(set GlooInstanceSet) GlooInstanceSet {
    if s == nil {
        return set
    }
    return &glooInstanceMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *glooInstanceMergedSet) Difference(set GlooInstanceSet) GlooInstanceSet {
    panic("unimplemented")
}

func (s *glooInstanceMergedSet) Intersection(set GlooInstanceSet) GlooInstanceSet {
    panic("unimplemented")
}

func (s *glooInstanceMergedSet) Find(id ezkube.ResourceId) (*fed_solo_io_v1.GlooInstance, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find GlooInstance %v", sksets.Key(id))
    }

    var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&fed_solo_io_v1.GlooInstance{}, id)
		if err == nil {
			return obj.(*fed_solo_io_v1.GlooInstance), nil
		}
	}

    return nil, err
}

func (s *glooInstanceMergedSet) Length() int {
    if s == nil {
        return 0
    }
    totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
    return totalLen
}

func (s *glooInstanceMergedSet) Generic() sksets.ResourceSet {
    panic("unimplemented")
}

func (s *glooInstanceMergedSet) Delta(newSet GlooInstanceSet) sksets.ResourceDelta {
    panic("unimplemented")
}

func (s *glooInstanceMergedSet) Clone() GlooInstanceSet {
	if s == nil {
		return nil
	}
	return &glooInstanceMergedSet{sets: s.sets[:]}
}

type FailoverSchemeSet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    List(filterResource ... func(*fed_solo_io_v1.FailoverScheme) bool) []*fed_solo_io_v1.FailoverScheme
    // Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    UnsortedList(filterResource ... func(*fed_solo_io_v1.FailoverScheme) bool) []*fed_solo_io_v1.FailoverScheme
    // Return the Set as a map of key to resource.
    Map() map[string]*fed_solo_io_v1.FailoverScheme
    // Insert a resource into the set.
    Insert(failoverScheme ...*fed_solo_io_v1.FailoverScheme)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(failoverSchemeSet FailoverSchemeSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(failoverScheme ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(failoverScheme  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set FailoverSchemeSet) FailoverSchemeSet
    // Return the difference with the provided set
    Difference(set FailoverSchemeSet) FailoverSchemeSet
    // Return the intersection with the provided set
    Intersection(set FailoverSchemeSet) FailoverSchemeSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*fed_solo_io_v1.FailoverScheme, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another FailoverSchemeSet
    Delta(newSet FailoverSchemeSet) sksets.ResourceDelta
    // Create a deep copy of the current FailoverSchemeSet
    Clone() FailoverSchemeSet
}

func makeGenericFailoverSchemeSet(failoverSchemeList []*fed_solo_io_v1.FailoverScheme) sksets.ResourceSet {
    var genericResources []ezkube.ResourceId
    for _, obj := range failoverSchemeList {
        genericResources = append(genericResources, obj)
    }
    return sksets.NewResourceSet(genericResources...)
}

type failoverSchemeSet struct {
    set sksets.ResourceSet
}

func NewFailoverSchemeSet(failoverSchemeList ...*fed_solo_io_v1.FailoverScheme) FailoverSchemeSet {
    return &failoverSchemeSet{set: makeGenericFailoverSchemeSet(failoverSchemeList)}
}

func NewFailoverSchemeSetFromList(failoverSchemeList *fed_solo_io_v1.FailoverSchemeList) FailoverSchemeSet {
    list := make([]*fed_solo_io_v1.FailoverScheme, 0, len(failoverSchemeList.Items))
    for idx := range failoverSchemeList.Items {
        list = append(list, &failoverSchemeList.Items[idx])
    }
    return &failoverSchemeSet{set: makeGenericFailoverSchemeSet(list)}
}

func (s *failoverSchemeSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *failoverSchemeSet) List(filterResource ... func(*fed_solo_io_v1.FailoverScheme) bool) []*fed_solo_io_v1.FailoverScheme {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_solo_io_v1.FailoverScheme))
        })
    }

    objs := s.Generic().List(genericFilters...)
    failoverSchemeList := make([]*fed_solo_io_v1.FailoverScheme, 0, len(objs))
    for _, obj := range objs {
        failoverSchemeList = append(failoverSchemeList, obj.(*fed_solo_io_v1.FailoverScheme))
    }
    return failoverSchemeList
}

func (s *failoverSchemeSet) UnsortedList(filterResource ... func(*fed_solo_io_v1.FailoverScheme) bool) []*fed_solo_io_v1.FailoverScheme {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_solo_io_v1.FailoverScheme))
        })
    }

    var failoverSchemeList []*fed_solo_io_v1.FailoverScheme
    for _, obj := range s.Generic().UnsortedList(genericFilters...) {
        failoverSchemeList = append(failoverSchemeList, obj.(*fed_solo_io_v1.FailoverScheme))
    }
    return failoverSchemeList
}

func (s *failoverSchemeSet) Map() map[string]*fed_solo_io_v1.FailoverScheme {
    if s == nil {
        return nil
    }

    newMap := map[string]*fed_solo_io_v1.FailoverScheme{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*fed_solo_io_v1.FailoverScheme)
    }
    return newMap
}

func (s *failoverSchemeSet) Insert(
        failoverSchemeList ...*fed_solo_io_v1.FailoverScheme,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range failoverSchemeList {
        s.Generic().Insert(obj)
    }
}

func (s *failoverSchemeSet) Has(failoverScheme ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(failoverScheme)
}

func (s *failoverSchemeSet) Equal(
        failoverSchemeSet FailoverSchemeSet,
) bool {
    if s == nil {
        return failoverSchemeSet == nil
    }
    return s.Generic().Equal(failoverSchemeSet.Generic())
}

func (s *failoverSchemeSet) Delete(FailoverScheme ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(FailoverScheme)
}

func (s *failoverSchemeSet) Union(set FailoverSchemeSet) FailoverSchemeSet {
    if s == nil {
        return set
    }
    return &failoverSchemeMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *failoverSchemeSet) Difference(set FailoverSchemeSet) FailoverSchemeSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &failoverSchemeSet{set: newSet}
}

func (s *failoverSchemeSet) Intersection(set FailoverSchemeSet) FailoverSchemeSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var failoverSchemeList []*fed_solo_io_v1.FailoverScheme
    for _, obj := range newSet.List() {
        failoverSchemeList = append(failoverSchemeList, obj.(*fed_solo_io_v1.FailoverScheme))
    }
    return NewFailoverSchemeSet(failoverSchemeList...)
}


func (s *failoverSchemeSet) Find(id ezkube.ResourceId) (*fed_solo_io_v1.FailoverScheme, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find FailoverScheme %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&fed_solo_io_v1.FailoverScheme{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*fed_solo_io_v1.FailoverScheme), nil
}

func (s *failoverSchemeSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *failoverSchemeSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *failoverSchemeSet) Delta(newSet FailoverSchemeSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

func (s *failoverSchemeSet) Clone() FailoverSchemeSet {
	if s == nil {
		return nil
	}
	return &failoverSchemeMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type failoverSchemeMergedSet struct {
    sets []sksets.ResourceSet
}

func NewFailoverSchemeMergedSet(failoverSchemeList ...*fed_solo_io_v1.FailoverScheme) FailoverSchemeSet {
    return &failoverSchemeMergedSet{sets: []sksets.ResourceSet{makeGenericFailoverSchemeSet(failoverSchemeList)}}
}

func NewFailoverSchemeMergedSetFromList(failoverSchemeList *fed_solo_io_v1.FailoverSchemeList) FailoverSchemeSet {
    list := make([]*fed_solo_io_v1.FailoverScheme, 0, len(failoverSchemeList.Items))
    for idx := range failoverSchemeList.Items {
        list = append(list, &failoverSchemeList.Items[idx])
    }
    return &failoverSchemeMergedSet{sets: []sksets.ResourceSet{makeGenericFailoverSchemeSet(list)}}
}

func (s *failoverSchemeMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
    }
    toRet := sets.String{}
	for _ , set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *failoverSchemeMergedSet) List(filterResource ... func(*fed_solo_io_v1.FailoverScheme) bool) []*fed_solo_io_v1.FailoverScheme {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_solo_io_v1.FailoverScheme))
        })
    }
   failoverSchemeList := []*fed_solo_io_v1.FailoverScheme{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.List(genericFilters...) {
            if tracker[obj] {
				continue
			}
			tracker[obj] = true
			failoverSchemeList = append(failoverSchemeList, obj.(*fed_solo_io_v1.FailoverScheme))
		}
	}
    return failoverSchemeList
}

func (s *failoverSchemeMergedSet) UnsortedList(filterResource ... func(*fed_solo_io_v1.FailoverScheme) bool) []*fed_solo_io_v1.FailoverScheme {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_solo_io_v1.FailoverScheme))
        })
    }
   failoverSchemeList := []*fed_solo_io_v1.FailoverScheme{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.UnsortedList(genericFilters...) {
            if tracker[obj] {
				continue
			}
			tracker[obj] = true
			failoverSchemeList = append(failoverSchemeList, obj.(*fed_solo_io_v1.FailoverScheme))
		}
	}
    return failoverSchemeList
}

func (s *failoverSchemeMergedSet) Map() map[string]*fed_solo_io_v1.FailoverScheme {
    if s == nil {
        return nil
    }

    newMap := map[string]*fed_solo_io_v1.FailoverScheme{}
    for _, set := range s.sets {
		for k, v := range set.Map() {
            newMap[k] = v.(*fed_solo_io_v1.FailoverScheme)
        }
    }
    return newMap
}

func (s *failoverSchemeMergedSet) Insert(
        failoverSchemeList ...*fed_solo_io_v1.FailoverScheme,
) {
    if s == nil {
    }
    if len(s.sets) == 0 {
        s.sets = append(s.sets, makeGenericFailoverSchemeSet(failoverSchemeList))
    }
    for _, obj := range failoverSchemeList {
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

func (s *failoverSchemeMergedSet) Has(failoverScheme ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    for _, set := range s.sets {
		if set.Has(failoverScheme) {
			return true
		}
	}
    return false
}

func (s *failoverSchemeMergedSet) Equal(
        failoverSchemeSet FailoverSchemeSet,
) bool {
    panic("unimplemented")
}

func (s *failoverSchemeMergedSet) Delete(FailoverScheme ezkube.ResourceId) {
    for _, set := range s.sets {
        set.Delete(FailoverScheme)
    }    
}

func (s *failoverSchemeMergedSet) Union(set FailoverSchemeSet) FailoverSchemeSet {
    if s == nil {
        return set
    }
    return &failoverSchemeMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *failoverSchemeMergedSet) Difference(set FailoverSchemeSet) FailoverSchemeSet {
    panic("unimplemented")
}

func (s *failoverSchemeMergedSet) Intersection(set FailoverSchemeSet) FailoverSchemeSet {
    panic("unimplemented")
}

func (s *failoverSchemeMergedSet) Find(id ezkube.ResourceId) (*fed_solo_io_v1.FailoverScheme, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find FailoverScheme %v", sksets.Key(id))
    }

    var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&fed_solo_io_v1.FailoverScheme{}, id)
		if err == nil {
			return obj.(*fed_solo_io_v1.FailoverScheme), nil
		}
	}

    return nil, err
}

func (s *failoverSchemeMergedSet) Length() int {
    if s == nil {
        return 0
    }
    totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
    return totalLen
}

func (s *failoverSchemeMergedSet) Generic() sksets.ResourceSet {
    panic("unimplemented")
}

func (s *failoverSchemeMergedSet) Delta(newSet FailoverSchemeSet) sksets.ResourceDelta {
    panic("unimplemented")
}

func (s *failoverSchemeMergedSet) Clone() FailoverSchemeSet {
	if s == nil {
		return nil
	}
	return &failoverSchemeMergedSet{sets: s.sets[:]}
}
