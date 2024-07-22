// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/internal/v2/mesh.proto

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
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *MeshSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.MeshSpec")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetInstallation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Installation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetInstallation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Installation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetTrustDomain())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetIstiodServiceAccount())); err != nil {
		return 0, err
	}

	for _, v := range m.GetDiscoveryNamespaces() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSmartDnsProxyingEnabled())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRootNamespace())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAgentInfo()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AgentInfo")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAgentInfo(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AgentInfo")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetHub())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTag())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetIpFamily())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetAmbientCapable())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *MeshStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.MeshStatus")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetObservedGeneration())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *MeshSpec_Installation) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.MeshSpec_Installation")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetNamespace())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCluster())); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetPodLabels() {
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

	if _, err = hasher.Write([]byte(m.GetVersion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRevision())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *MeshSpec_AgentInfo) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.MeshSpec_AgentInfo")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetNamespace())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetVersion())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRelayRootTlsSecret()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RelayRootTlsSecret")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRelayRootTlsSecret(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RelayRootTlsSecret")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
