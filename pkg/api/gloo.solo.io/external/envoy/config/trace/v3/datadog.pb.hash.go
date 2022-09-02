// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/trace/v3/datadog.proto

package v3

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
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
func (m *DatadogConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.trace.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/trace/v3.DatadogConfig")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetServiceName()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ServiceName")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetServiceName(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ServiceName")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.CollectorCluster.(type) {

	case *DatadogConfig_CollectorUpstreamRef:

		if h, ok := interface{}(m.GetCollectorUpstreamRef()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("CollectorUpstreamRef")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetCollectorUpstreamRef(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("CollectorUpstreamRef")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *DatadogConfig_ClusterName:

		if _, err = hasher.Write([]byte(m.GetClusterName())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
