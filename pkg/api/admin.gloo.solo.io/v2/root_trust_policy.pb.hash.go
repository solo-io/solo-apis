// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/admin/v2/root_trust_policy.proto

package v2

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
func (m *RootTrustPolicySpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/admin.gloo.solo.io/v2.RootTrustPolicySpec")); err != nil {
		return 0, err
	}

	for _, v := range m.GetApplyToMeshes() {

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

	if h, ok := interface{}(m.GetConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Config")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Config")); err != nil {
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
func (m *RootTrustPolicyStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/admin.gloo.solo.io/v2.RootTrustPolicyStatus")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetObservedGeneration())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetState())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RootTrustPolicySpec_Config) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/admin.gloo.solo.io/v2.RootTrustPolicySpec_Config")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetIntermediateCertOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("IntermediateCertOptions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIntermediateCertOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("IntermediateCertOptions")); err != nil {
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

	switch m.CertificateAuthorityType.(type) {

	case *RootTrustPolicySpec_Config_MgmtServerCa:

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

	case *RootTrustPolicySpec_Config_AgentCa:

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
func (m *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/admin.gloo.solo.io/v2.RootTrustPolicySpec_Config_MgmtServerCertificateAuthority")); err != nil {
		return 0, err
	}

	switch m.CaSource.(type) {

	case *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_Generated:

		if h, ok := interface{}(m.GetGenerated()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Generated")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGenerated(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Generated")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *RootTrustPolicySpec_Config_MgmtServerCertificateAuthority_SecretRef:

		if h, ok := interface{}(m.GetSecretRef()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("SecretRef")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSecretRef(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("SecretRef")); err != nil {
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
