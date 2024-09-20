// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/consul/consul.proto

package consul

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
	_ = strconv.Itoa
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *UpstreamSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("consul.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/consul.UpstreamSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ServiceName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetServiceName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ServiceTags")); err != nil {
		return 0, err
	}
	for i, v := range m.GetServiceTags() {
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

	if _, err = hasher.Write([]byte("SubsetTags")); err != nil {
		return 0, err
	}
	for i, v := range m.GetSubsetTags() {
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

	if _, err = hasher.Write([]byte("InstanceTags")); err != nil {
		return 0, err
	}
	for i, v := range m.GetInstanceTags() {
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

	if _, err = hasher.Write([]byte("InstanceBlacklistTags")); err != nil {
		return 0, err
	}
	for i, v := range m.GetInstanceBlacklistTags() {
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

	if h, ok := interface{}(m.GetServiceSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ServiceSpec")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetServiceSpec(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ServiceSpec")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("ConsistencyMode")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetConsistencyMode())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetQueryOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("QueryOptions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetQueryOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("QueryOptions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("ConnectEnabled")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetConnectEnabled())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DataCenters")); err != nil {
		return 0, err
	}
	for i, v := range m.GetDataCenters() {
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

	return hasher.Sum64(), nil
}
