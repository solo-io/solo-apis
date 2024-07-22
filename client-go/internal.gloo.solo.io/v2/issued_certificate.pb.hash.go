// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/internal/v2/issued_certificate.proto

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
func (m *IssuedCertificateSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.IssuedCertificateSpec")); err != nil {
		return 0, err
	}

	for _, v := range m.GetHosts() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetIssuedCertificateSecret()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("IssuedCertificateSecret")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIssuedCertificateSecret(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("IssuedCertificateSecret")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCertOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CertOptions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCertOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CertOptions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMeshRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MeshRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMeshRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MeshRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetAutoRestartPods())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetPassiveCertificateAuthorities() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.CertificateAuthority.(type) {

	case *IssuedCertificateSpec_MgmtServerCa:

		if h, ok := interface{}(m.GetMgmtServerCa()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("MgmtServerCa")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMgmtServerCa(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("MgmtServerCa")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *IssuedCertificateSpec_AgentCa:

		if h, ok := interface{}(m.GetAgentCa()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("AgentCa")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAgentCa(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("AgentCa")); err != nil {
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

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *MgmtServerCertificateAuthority) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.MgmtServerCertificateAuthority")); err != nil {
		return 0, err
	}

	switch m.CertificateAuthority.(type) {

	case *MgmtServerCertificateAuthority_SigningCertificateSecret:

		if h, ok := interface{}(m.GetSigningCertificateSecret()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("SigningCertificateSecret")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSigningCertificateSecret(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("SigningCertificateSecret")); err != nil {
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

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *IssuedCertificateStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("internal.gloo.solo.io.github.com/solo-io/gloo-mesh-enterprise/pkg/api/internal.gloo.solo.io/v2.IssuedCertificateStatus")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetObservedGeneration())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetError())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetState())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
