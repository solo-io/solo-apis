// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/graphql/graphql.proto

package graphql

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
func (m *ServiceSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/graphql.ServiceSpec")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEndpoint()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetEndpoint(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ServiceSpec_Endpoint) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("graphql.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/graphql.ServiceSpec_Endpoint")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetUrl())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
