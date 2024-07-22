// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/internal/v2/portal_config.proto

package v2

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetPortalCustomMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PortalCustomMetadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPortalCustomMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PortalCustomMetadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Apis")); err != nil {
		return 0, err
	}
	for i, v := range m.GetApis() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if _, err = hasher.Write([]byte("UsagePlans")); err != nil {
		return 0, err
	}
	for i, v := range m.GetUsagePlans() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetPortalRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PortalRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPortalRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PortalRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Domains")); err != nil {
		return 0, err
	}
	for i, v := range m.GetDomains() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte("Groups")); err != nil {
		return 0, err
	}
	for i, v := range m.GetGroups() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if _, err = hasher.Write([]byte("Public")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetPublic())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigStatus) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigStatus")); err != nil {
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

	if _, err = hasher.Write([]byte("OwnedByWorkspace")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetOwnedByWorkspace())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigReport) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigReport")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetWorkspaces() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
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

	if _, err = hasher.Write([]byte("OwnerWorkspace")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetOwnerWorkspace())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec_Group) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec_Group")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Apis")); err != nil {
		return 0, err
	}
	for i, v := range m.GetApis() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if _, err = hasher.Write([]byte("UsagePlans")); err != nil {
		return 0, err
	}
	for i, v := range m.GetUsagePlans() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte("MembershipClaims")); err != nil {
		return 0, err
	}
	for i, v := range m.GetMembershipClaims() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec_API) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec_API")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ApiProductId")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetApiProductId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ApiProductDisplayName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetApiProductDisplayName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ApiId")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetApiId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ApiVersion")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetApiVersion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Title")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetTitle())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Description")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetDescription())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("TermsOfService")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetTermsOfService())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Contact")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetContact())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("License")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetLicense())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Lifecycle")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetLifecycle())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetApiSchema()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ApiSchema")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetApiSchema(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ApiSchema")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
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

	if _, err = hasher.Write([]byte("UsagePlans")); err != nil {
		return 0, err
	}
	for i, v := range m.GetUsagePlans() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if _, err = hasher.Write([]byte("IsPrivate")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetIsPrivate())
	if err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetCustomMetadata() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec_UsagePlan) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec_UsagePlan")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DisplayName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetDisplayName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Description")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetDescription())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ExtAuthPolicies")); err != nil {
		return 0, err
	}
	for i, v := range m.GetExtAuthPolicies() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetRateLimitPolicy()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RateLimitPolicy")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRateLimitPolicy(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RateLimitPolicy")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec_UsagePlanRef) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec_UsagePlanRef")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec_ExtAuthPolicy) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec_ExtAuthPolicy")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetExtAuthPolicyRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ExtAuthPolicyRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExtAuthPolicyRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ExtAuthPolicyRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("AuthConfigId")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetAuthConfigId())); err != nil {
		return 0, err
	}

	switch m.AuthCfg.(type) {

	case *PortalConfigSpec_ExtAuthPolicy_ApiKeyAuth:

		if h, ok := interface{}(m.GetApiKeyAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ApiKeyAuth")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetApiKeyAuth(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ApiKeyAuth")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *PortalConfigSpec_ExtAuthPolicy_OidcAuth:

		if h, ok := interface{}(m.GetOidcAuth()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("OidcAuth")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOidcAuth(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("OidcAuth")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *PortalConfigSpec_ExtAuthPolicy_AccessTokenValidation:

		if h, ok := interface{}(m.GetAccessTokenValidation()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AccessTokenValidation")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAccessTokenValidation(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AccessTokenValidation")); err != nil {
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec_ApiKeyAuth) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec_ApiKeyAuth")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetExtAuthLabelSelector() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec_OidcAuth) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec_OidcAuth")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("WellKnownOpenidConfig")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetWellKnownOpenidConfig())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec_AccessTokenValidation) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec_AccessTokenValidation")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("BearerFormat")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetBearerFormat())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *PortalConfigSpec_RateLimitPolicy) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.PortalConfigSpec_RateLimitPolicy")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Unit")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetUnit())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RequestsPerUnit")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetRequestsPerUnit())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRateLimitPolicyRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RateLimitPolicyRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRateLimitPolicyRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RateLimitPolicyRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
