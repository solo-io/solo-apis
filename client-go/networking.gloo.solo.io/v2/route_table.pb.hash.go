// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/networking/v2/route_table.proto

package v2

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *RouteTableSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.RouteTableSpec")); err != nil {
		return 0, err
	}

	for _, v := range m.GetHosts() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetVirtualGateways() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetWorkloadSelectors() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetApplyToDestinations() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetDefaultDestination()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DefaultDestination")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDefaultDestination(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DefaultDestination")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetHttp() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetTcp() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetTls() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetWeight())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetPortalMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PortalMetadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPortalMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PortalMetadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HTTPRoute) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.HTTPRoute")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetLabels() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetMatchers() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.ActionType.(type) {

	case *HTTPRoute_ForwardTo:

		if h, ok := interface{}(m.GetForwardTo()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ForwardTo")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetForwardTo(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ForwardTo")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HTTPRoute_Delegate:

		if h, ok := interface{}(m.GetDelegate()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Delegate")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDelegate(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Delegate")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HTTPRoute_Redirect:

		if h, ok := interface{}(m.GetRedirect()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Redirect")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRedirect(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Redirect")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HTTPRoute_DirectResponse:

		if h, ok := interface{}(m.GetDirectResponse()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DirectResponse")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDirectResponse(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DirectResponse")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HTTPRoute_Graphql:

		if h, ok := interface{}(m.GetGraphql()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Graphql")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGraphql(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Graphql")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TCPRoute) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.TCPRoute")); err != nil {
		return 0, err
	}

	for _, v := range m.GetMatchers() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.ActionType.(type) {

	case *TCPRoute_ForwardTo:

		if h, ok := interface{}(m.GetForwardTo()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ForwardTo")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetForwardTo(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ForwardTo")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TLSRoute) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.TLSRoute")); err != nil {
		return 0, err
	}

	for _, v := range m.GetMatchers() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.ActionType.(type) {

	case *TLSRoute_ForwardTo:

		if h, ok := interface{}(m.GetForwardTo()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ForwardTo")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetForwardTo(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ForwardTo")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GraphQLAction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.GraphQLAction")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Options")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Options")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.GraphqlSchema.(type) {

	case *GraphQLAction_Schema:

		if h, ok := interface{}(m.GetSchema()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Schema")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSchema(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Schema")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GraphQLAction_StitchedSchema:

		if h, ok := interface{}(m.GetStitchedSchema()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("StitchedSchema")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetStitchedSchema(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("StitchedSchema")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ForwardToAction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.ForwardToAction")); err != nil {
		return 0, err
	}

	for _, v := range m.GetDestinations() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.PathRewriteSpecifier.(type) {

	case *ForwardToAction_PathRewrite:

		if _, err = hasher.Write([]byte(m.GetPathRewrite())); err != nil {
			return 0, err
		}

	case *ForwardToAction_RegexRewrite:

		if h, ok := interface{}(m.GetRegexRewrite()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RegexRewrite")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRegexRewrite(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RegexRewrite")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.HostRewriteSpecifier.(type) {

	case *ForwardToAction_HostRewrite:

		if _, err = hasher.Write([]byte(m.GetHostRewrite())); err != nil {
			return 0, err
		}

	case *ForwardToAction_AutoHostRewrite:

		err = binary.Write(hasher, binary.LittleEndian, m.GetAutoHostRewrite())
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RedirectAction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.RedirectAction")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHostRedirect())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetResponseCode())
	if err != nil {
		return 0, err
	}

	switch m.PathRewriteSpecifier.(type) {

	case *RedirectAction_PathRedirect:

		if _, err = hasher.Write([]byte(m.GetPathRedirect())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DirectResponseAction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.DirectResponseAction")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetStatus())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetBody())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DelegateAction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.DelegateAction")); err != nil {
		return 0, err
	}

	for _, v := range m.GetRouteTables() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetAllowedRoutes() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSortMethod())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *PortalMetadata) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.PortalMetadata")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetApiProductId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetApiProductDisplayName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetApiVersion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTitle())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetDescription())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTermsOfService())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetContact())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetLicense())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetLifecycle())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetCustomMetadata() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RouteTableStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.RouteTableStatus")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetCommon()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Common")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCommon(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Common")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetNumAppliedRoutePolicies() {
			innerHash.Reset()

			err = binary.Write(innerHash, binary.LittleEndian, v)
			if err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetNumParentRouteTables())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetOwnedByWorkspace())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetNumAllowedVirtualGateways())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RouteTableReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.RouteTableReport")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetWorkspaces() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetAppliedRoutePolicies() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetParentRouteTables() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if _, err = hasher.Write([]byte(m.GetOwnerWorkspace())); err != nil {
		return 0, err
	}

	for _, v := range m.GetAllowedVirtualGateways() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetDelegatedToRouteTables() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TLSRoute_TLSForwardToAction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.TLSRoute_TLSForwardToAction")); err != nil {
		return 0, err
	}

	for _, v := range m.GetDestinations() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GraphQLAction_Options) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.GraphQLAction_Options")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetLogSensitiveInfo()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LogSensitiveInfo")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLogSensitiveInfo(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LogSensitiveInfo")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RouteTableReport_DelegatedRouteTableReference) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.gloo.solo.io.github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2.RouteTableReport_DelegatedRouteTableReference")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRouteIndex())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRouteTable()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RouteTable")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRouteTable(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RouteTable")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
