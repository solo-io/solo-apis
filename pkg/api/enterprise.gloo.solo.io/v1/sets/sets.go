// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets



import (
    enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"

    "github.com/rotisserie/eris"
    sksets "github.com/solo-io/skv2/contrib/pkg/sets"
    "github.com/solo-io/skv2/pkg/ezkube"
    "k8s.io/apimachinery/pkg/util/sets"
)

type AuthConfigSet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    List(filterResource ... func(*enterprise_gloo_solo_io_v1.AuthConfig) bool) []*enterprise_gloo_solo_io_v1.AuthConfig
    // Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    UnsortedList(filterResource ... func(*enterprise_gloo_solo_io_v1.AuthConfig) bool) []*enterprise_gloo_solo_io_v1.AuthConfig
    // Return the Set as a map of key to resource.
    Map() map[string]*enterprise_gloo_solo_io_v1.AuthConfig
    // Insert a resource into the set.
    Insert(authConfig ...*enterprise_gloo_solo_io_v1.AuthConfig)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(authConfigSet AuthConfigSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(authConfig ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(authConfig  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set AuthConfigSet) AuthConfigSet
    // Return the difference with the provided set
    Difference(set AuthConfigSet) AuthConfigSet
    // Return the intersection with the provided set
    Intersection(set AuthConfigSet) AuthConfigSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*enterprise_gloo_solo_io_v1.AuthConfig, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another AuthConfigSet
    Delta(newSet AuthConfigSet) sksets.ResourceDelta
    // Create a deep copy of the current AuthConfigSet
    Clone() AuthConfigSet
}

func makeGenericAuthConfigSet(authConfigList []*enterprise_gloo_solo_io_v1.AuthConfig) sksets.ResourceSet {
    var genericResources []ezkube.ResourceId
    for _, obj := range authConfigList {
        genericResources = append(genericResources, obj)
    }
    return sksets.NewResourceSet(genericResources...)
}

type authConfigSet struct {
    set sksets.ResourceSet
}

func NewAuthConfigSet(authConfigList ...*enterprise_gloo_solo_io_v1.AuthConfig) AuthConfigSet {
    return &authConfigSet{set: makeGenericAuthConfigSet(authConfigList)}
}

func NewAuthConfigSetFromList(authConfigList *enterprise_gloo_solo_io_v1.AuthConfigList) AuthConfigSet {
    list := make([]*enterprise_gloo_solo_io_v1.AuthConfig, 0, len(authConfigList.Items))
    for idx := range authConfigList.Items {
        list = append(list, &authConfigList.Items[idx])
    }
    return &authConfigSet{set: makeGenericAuthConfigSet(list)}
}

func (s *authConfigSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *authConfigSet) List(filterResource ... func(*enterprise_gloo_solo_io_v1.AuthConfig) bool) []*enterprise_gloo_solo_io_v1.AuthConfig {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*enterprise_gloo_solo_io_v1.AuthConfig))
        })
    }

    objs := s.Generic().List(genericFilters...)
    authConfigList := make([]*enterprise_gloo_solo_io_v1.AuthConfig, 0, len(objs))
    for _, obj := range objs {
        authConfigList = append(authConfigList, obj.(*enterprise_gloo_solo_io_v1.AuthConfig))
    }
    return authConfigList
}

func (s *authConfigSet) UnsortedList(filterResource ... func(*enterprise_gloo_solo_io_v1.AuthConfig) bool) []*enterprise_gloo_solo_io_v1.AuthConfig {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*enterprise_gloo_solo_io_v1.AuthConfig))
        })
    }

    var authConfigList []*enterprise_gloo_solo_io_v1.AuthConfig
    for _, obj := range s.Generic().UnsortedList(genericFilters...) {
        authConfigList = append(authConfigList, obj.(*enterprise_gloo_solo_io_v1.AuthConfig))
    }
    return authConfigList
}

func (s *authConfigSet) Map() map[string]*enterprise_gloo_solo_io_v1.AuthConfig {
    if s == nil {
        return nil
    }

    newMap := map[string]*enterprise_gloo_solo_io_v1.AuthConfig{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*enterprise_gloo_solo_io_v1.AuthConfig)
    }
    return newMap
}

func (s *authConfigSet) Insert(
        authConfigList ...*enterprise_gloo_solo_io_v1.AuthConfig,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range authConfigList {
        s.Generic().Insert(obj)
    }
}

func (s *authConfigSet) Has(authConfig ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(authConfig)
}

func (s *authConfigSet) Equal(
        authConfigSet AuthConfigSet,
) bool {
    if s == nil {
        return authConfigSet == nil
    }
    return s.Generic().Equal(authConfigSet.Generic())
}

func (s *authConfigSet) Delete(AuthConfig ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(AuthConfig)
}

func (s *authConfigSet) Union(set AuthConfigSet) AuthConfigSet {
    if s == nil {
        return set
    }
    return NewAuthConfigSet(append(s.List(), set.List()...)...)
}

func (s *authConfigSet) Difference(set AuthConfigSet) AuthConfigSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &authConfigSet{set: newSet}
}

func (s *authConfigSet) Intersection(set AuthConfigSet) AuthConfigSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var authConfigList []*enterprise_gloo_solo_io_v1.AuthConfig
    for _, obj := range newSet.List() {
        authConfigList = append(authConfigList, obj.(*enterprise_gloo_solo_io_v1.AuthConfig))
    }
    return NewAuthConfigSet(authConfigList...)
}


func (s *authConfigSet) Find(id ezkube.ResourceId) (*enterprise_gloo_solo_io_v1.AuthConfig, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find AuthConfig %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&enterprise_gloo_solo_io_v1.AuthConfig{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*enterprise_gloo_solo_io_v1.AuthConfig), nil
}

func (s *authConfigSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *authConfigSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *authConfigSet) Delta(newSet AuthConfigSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

func (s *authConfigSet) Clone() AuthConfigSet {
	if s == nil {
		return nil
	}
	return &authConfigSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
