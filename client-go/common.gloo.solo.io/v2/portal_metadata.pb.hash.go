// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/portal_metadata.proto

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
func (m *PortalMetadata) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("common.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/common.gloo.solo.io/v2.PortalMetadata")); err != nil {
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
