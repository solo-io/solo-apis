// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/rate-limiter/v1alpha1/ratelimit.proto

package v1alpha1

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
func (m *RateLimitConfigSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.RateLimitConfigSpec")); err != nil {
		return 0, err
	}

	switch m.ConfigType.(type) {

	case *RateLimitConfigSpec_Raw_:

		if h, ok := interface{}(m.GetRaw()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Raw")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRaw(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Raw")); err != nil {
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
func (m *RateLimitConfigStatus) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.RateLimitConfigStatus")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("State")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetState())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Message")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetMessage())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ObservedGeneration")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetObservedGeneration())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RateLimitConfigNamespacedStatuses) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.RateLimitConfigNamespacedStatuses")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetStatuses() {
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

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Descriptor) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Descriptor")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Key")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Value")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetValue())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRateLimit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RateLimit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRateLimit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RateLimit")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Descriptors")); err != nil {
		return 0, err
	}
	for i, v := range m.GetDescriptors() {
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

	if _, err = hasher.Write([]byte("Weight")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetWeight())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AlwaysApply")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAlwaysApply())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *SetDescriptor) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.SetDescriptor")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("SimpleDescriptors")); err != nil {
		return 0, err
	}
	for i, v := range m.GetSimpleDescriptors() {
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

	if h, ok := interface{}(m.GetRateLimit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RateLimit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRateLimit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RateLimit")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("AlwaysApply")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAlwaysApply())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *SimpleDescriptor) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.SimpleDescriptor")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Key")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Value")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetValue())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RateLimitActions) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.RateLimitActions")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Actions")); err != nil {
		return 0, err
	}
	for i, v := range m.GetActions() {
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

	if _, err = hasher.Write([]byte("SetActions")); err != nil {
		return 0, err
	}
	for i, v := range m.GetSetActions() {
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

	if h, ok := interface{}(m.GetLimit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Limit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLimit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Limit")); err != nil {
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
func (m *RateLimit) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.RateLimit")); err != nil {
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

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Action) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Action")); err != nil {
		return 0, err
	}

	switch m.ActionSpecifier.(type) {

	case *Action_SourceCluster_:

		if h, ok := interface{}(m.GetSourceCluster()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("SourceCluster")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSourceCluster(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("SourceCluster")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Action_DestinationCluster_:

		if h, ok := interface{}(m.GetDestinationCluster()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DestinationCluster")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDestinationCluster(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DestinationCluster")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Action_RequestHeaders_:

		if h, ok := interface{}(m.GetRequestHeaders()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RequestHeaders")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRequestHeaders(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RequestHeaders")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Action_RemoteAddress_:

		if h, ok := interface{}(m.GetRemoteAddress()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RemoteAddress")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRemoteAddress(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RemoteAddress")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Action_GenericKey_:

		if h, ok := interface{}(m.GetGenericKey()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GenericKey")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGenericKey(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GenericKey")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Action_HeaderValueMatch_:

		if h, ok := interface{}(m.GetHeaderValueMatch()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("HeaderValueMatch")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHeaderValueMatch(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("HeaderValueMatch")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Action_Metadata:

		if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Metadata")); err != nil {
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
func (m *MetaData) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.MetaData")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DescriptorKey")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetDescriptorKey())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadataKey()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MetadataKey")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadataKey(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MetadataKey")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("DefaultValue")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetDefaultValue())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Source")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetSource())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Override) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Override")); err != nil {
		return 0, err
	}

	switch m.OverrideSpecifier.(type) {

	case *Override_DynamicMetadata_:

		if h, ok := interface{}(m.GetDynamicMetadata()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DynamicMetadata")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDynamicMetadata(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DynamicMetadata")); err != nil {
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
func (m *RateLimitConfigSpec_Raw) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.RateLimitConfigSpec_Raw")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Descriptors")); err != nil {
		return 0, err
	}
	for i, v := range m.GetDescriptors() {
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

	if _, err = hasher.Write([]byte("RateLimits")); err != nil {
		return 0, err
	}
	for i, v := range m.GetRateLimits() {
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

	if _, err = hasher.Write([]byte("SetDescriptors")); err != nil {
		return 0, err
	}
	for i, v := range m.GetSetDescriptors() {
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
func (m *Action_SourceCluster) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Action_SourceCluster")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Action_DestinationCluster) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Action_DestinationCluster")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Action_RequestHeaders) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Action_RequestHeaders")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("HeaderName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetHeaderName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DescriptorKey")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetDescriptorKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Action_RemoteAddress) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Action_RemoteAddress")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Action_GenericKey) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Action_GenericKey")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DescriptorValue")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetDescriptorValue())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Action_HeaderValueMatch) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Action_HeaderValueMatch")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DescriptorValue")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetDescriptorValue())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetExpectMatch()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ExpectMatch")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExpectMatch(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ExpectMatch")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Headers")); err != nil {
		return 0, err
	}
	for i, v := range m.GetHeaders() {
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
func (m *Action_HeaderValueMatch_HeaderMatcher) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Action_HeaderValueMatch_HeaderMatcher")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("InvertMatch")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetInvertMatch())
	if err != nil {
		return 0, err
	}

	switch m.HeaderMatchSpecifier.(type) {

	case *Action_HeaderValueMatch_HeaderMatcher_ExactMatch:

		if _, err = hasher.Write([]byte("ExactMatch")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetExactMatch())); err != nil {
			return 0, err
		}

	case *Action_HeaderValueMatch_HeaderMatcher_RegexMatch:

		if _, err = hasher.Write([]byte("RegexMatch")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetRegexMatch())); err != nil {
			return 0, err
		}

	case *Action_HeaderValueMatch_HeaderMatcher_RangeMatch:

		if h, ok := interface{}(m.GetRangeMatch()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RangeMatch")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRangeMatch(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RangeMatch")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Action_HeaderValueMatch_HeaderMatcher_PresentMatch:

		if _, err = hasher.Write([]byte("PresentMatch")); err != nil {
			return 0, err
		}
		err = binary.Write(hasher, binary.LittleEndian, m.GetPresentMatch())
		if err != nil {
			return 0, err
		}

	case *Action_HeaderValueMatch_HeaderMatcher_PrefixMatch:

		if _, err = hasher.Write([]byte("PrefixMatch")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetPrefixMatch())); err != nil {
			return 0, err
		}

	case *Action_HeaderValueMatch_HeaderMatcher_SuffixMatch:

		if _, err = hasher.Write([]byte("SuffixMatch")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetSuffixMatch())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Action_HeaderValueMatch_HeaderMatcher_Int64Range) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Action_HeaderValueMatch_HeaderMatcher_Int64Range")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Start")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetStart())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("End")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetEnd())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *MetaData_MetadataKey) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.MetaData_MetadataKey")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Key")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Path")); err != nil {
		return 0, err
	}
	for i, v := range m.GetPath() {
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
func (m *MetaData_MetadataKey_PathSegment) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.MetaData_MetadataKey_PathSegment")); err != nil {
		return 0, err
	}

	switch m.Segment.(type) {

	case *MetaData_MetadataKey_PathSegment_Key:

		if _, err = hasher.Write([]byte("Key")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Override_DynamicMetadata) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ratelimit.api.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/ratelimit.solo.io/v1alpha1.Override_DynamicMetadata")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadataKey()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MetadataKey")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadataKey(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MetadataKey")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
