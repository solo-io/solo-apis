// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets



import (
    fed_enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/fed.enterprise.gloo.solo.io/v1"

    "github.com/rotisserie/eris"
    sksets "github.com/solo-io/skv2/contrib/pkg/sets"
    "github.com/solo-io/skv2/pkg/ezkube"
    "k8s.io/apimachinery/pkg/util/sets"
)

type FederatedAuthConfigSet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    List(filterResource ... func(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) bool) []*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig
    // Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
    UnsortedList(filterResource ... func(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) bool) []*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig
    // Return the Set as a map of key to resource.
    Map() map[string]*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig
    // Insert a resource into the set.
    Insert(federatedAuthConfig ...*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(federatedAuthConfigSet FederatedAuthConfigSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(federatedAuthConfig ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(federatedAuthConfig  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set FederatedAuthConfigSet) FederatedAuthConfigSet
    // Return the difference with the provided set
    Difference(set FederatedAuthConfigSet) FederatedAuthConfigSet
    // Return the intersection with the provided set
    Intersection(set FederatedAuthConfigSet) FederatedAuthConfigSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another FederatedAuthConfigSet
    Delta(newSet FederatedAuthConfigSet) sksets.ResourceDelta
    // Create a deep copy of the current FederatedAuthConfigSet
    Clone() FederatedAuthConfigSet
}

func makeGenericFederatedAuthConfigSet(federatedAuthConfigList []*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) sksets.ResourceSet {
    var genericResources []ezkube.ResourceId
    for _, obj := range federatedAuthConfigList {
        genericResources = append(genericResources, obj)
    }
    return sksets.NewResourceSet(genericResources...)
}

type federatedAuthConfigSet struct {
    set sksets.ResourceSet
}

func NewFederatedAuthConfigSet(federatedAuthConfigList ...*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) FederatedAuthConfigSet {
    return &federatedAuthConfigSet{set: makeGenericFederatedAuthConfigSet(federatedAuthConfigList)}
}

func NewFederatedAuthConfigSetFromList(federatedAuthConfigList *fed_enterprise_gloo_solo_io_v1.FederatedAuthConfigList) FederatedAuthConfigSet {
    list := make([]*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig, 0, len(federatedAuthConfigList.Items))
    for idx := range federatedAuthConfigList.Items {
        list = append(list, &federatedAuthConfigList.Items[idx])
    }
    return &federatedAuthConfigSet{set: makeGenericFederatedAuthConfigSet(list)}
}

func (s *federatedAuthConfigSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *federatedAuthConfigSet) List(filterResource ... func(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) bool) []*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig))
        })
    }

    objs := s.Generic().List(genericFilters...)
    federatedAuthConfigList := make([]*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig, 0, len(objs))
    for _, obj := range objs {
        federatedAuthConfigList = append(federatedAuthConfigList, obj.(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig))
    }
    return federatedAuthConfigList
}

func (s *federatedAuthConfigSet) UnsortedList(filterResource ... func(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) bool) []*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig))
        })
    }

    var federatedAuthConfigList []*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig
    for _, obj := range s.Generic().UnsortedList(genericFilters...) {
        federatedAuthConfigList = append(federatedAuthConfigList, obj.(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig))
    }
    return federatedAuthConfigList
}

func (s *federatedAuthConfigSet) Map() map[string]*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig {
    if s == nil {
        return nil
    }

    newMap := map[string]*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig)
    }
    return newMap
}

func (s *federatedAuthConfigSet) Insert(
        federatedAuthConfigList ...*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range federatedAuthConfigList {
        s.Generic().Insert(obj)
    }
}

func (s *federatedAuthConfigSet) Has(federatedAuthConfig ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(federatedAuthConfig)
}

func (s *federatedAuthConfigSet) Equal(
        federatedAuthConfigSet FederatedAuthConfigSet,
) bool {
    if s == nil {
        return federatedAuthConfigSet == nil
    }
    return s.Generic().Equal(federatedAuthConfigSet.Generic())
}

func (s *federatedAuthConfigSet) Delete(FederatedAuthConfig ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(FederatedAuthConfig)
}

func (s *federatedAuthConfigSet) Union(set FederatedAuthConfigSet) FederatedAuthConfigSet {
    if s == nil {
        return set
    }
    return NewFederatedAuthConfigSet(append(s.List(), set.List()...)...)
}

func (s *federatedAuthConfigSet) Difference(set FederatedAuthConfigSet) FederatedAuthConfigSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &federatedAuthConfigSet{set: newSet}
}

func (s *federatedAuthConfigSet) Intersection(set FederatedAuthConfigSet) FederatedAuthConfigSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var federatedAuthConfigList []*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig
    for _, obj := range newSet.List() {
        federatedAuthConfigList = append(federatedAuthConfigList, obj.(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig))
    }
    return NewFederatedAuthConfigSet(federatedAuthConfigList...)
}


func (s *federatedAuthConfigSet) Find(id ezkube.ResourceId) (*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find FederatedAuthConfig %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig), nil
}

func (s *federatedAuthConfigSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *federatedAuthConfigSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *federatedAuthConfigSet) Delta(newSet FederatedAuthConfigSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

func (s *federatedAuthConfigSet) Clone() FederatedAuthConfigSet {
	if s == nil {
		return nil
	}
	return &federatedAuthConfigSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
