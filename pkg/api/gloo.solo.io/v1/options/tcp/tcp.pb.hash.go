// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/tcp/tcp.proto

package tcp

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
func (m *TcpProxySettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("tcp.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/tcp.TcpProxySettings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMaxConnectAttempts()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxConnectAttempts")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxConnectAttempts(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxConnectAttempts")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("IdleTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIdleTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("IdleTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTunnelingConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TunnelingConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTunnelingConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TunnelingConfig")); err != nil {
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
func (m *TcpProxySettings_TunnelingConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("tcp.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/tcp.TcpProxySettings_TunnelingConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHostname())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
