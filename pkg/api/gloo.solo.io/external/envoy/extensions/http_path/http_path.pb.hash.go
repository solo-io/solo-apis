// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/http_path/http_path.proto

package http_path

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
func (m *HttpPath) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.health_checker.http_path.v2.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/extensions/http_path.HttpPath")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHttpHealthCheck()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HttpHealthCheck")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHttpHealthCheck(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HttpHealthCheck")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
