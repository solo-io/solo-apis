// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/protocol_upgrade/protocol_upgrade.proto

package protocol_upgrade

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
func (m *ProtocolUpgradeConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("protocol_upgrade.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/protocol_upgrade.ProtocolUpgradeConfig")); err != nil {
		return 0, err
	}

	switch m.UpgradeType.(type) {

	case *ProtocolUpgradeConfig_Websocket:

		if h, ok := interface{}(m.GetWebsocket()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetWebsocket(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("protocol_upgrade.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/protocol_upgrade.ProtocolUpgradeConfig_ProtocolUpgradeSpec")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEnabled()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetEnabled(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
