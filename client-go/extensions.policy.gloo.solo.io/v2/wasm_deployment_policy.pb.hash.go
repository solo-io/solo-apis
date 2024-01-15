// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/extensions/wasm_deployment_policy.proto

package v2

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"

	v1alpha3 "istio.io/api/networking/v1alpha3"
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

	_ = v1alpha3.EnvoyFilter_PatchContext(0)
)

// Hash function
func (m *WasmDeploymentPolicySpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extensions.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2.WasmDeploymentPolicySpec")); err != nil {
		return 0, err
	}

	for _, v := range m.GetApplyToWorkloads() {

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
func (m *WasmDeploymentPolicyStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extensions.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2.WasmDeploymentPolicyStatus")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetCommon()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Common")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCommon(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Common")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetNumSelectedWorkloads())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *WasmDeploymentPolicyReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extensions.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2.WasmDeploymentPolicyReport")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetWorkspaces() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	for _, v := range m.GetSelectedWorkloads() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *WasmDeploymentPolicySpec_Config) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extensions.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2.WasmDeploymentPolicySpec_Config")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFilters() {

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

	err = binary.Write(hasher, binary.LittleEndian, m.GetWeight())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *WasmDeploymentPolicySpec_Config_WasmFilter) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extensions.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2.WasmDeploymentPolicySpec_Config_WasmFilter")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRootId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetVmId())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFilterContext())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetInsertBeforeFilter())); err != nil {
		return 0, err
	}

	switch m.FilterSource.(type) {

	case *WasmDeploymentPolicySpec_Config_WasmFilter_LocalPathSource:

		if _, err = hasher.Write([]byte(m.GetLocalPathSource())); err != nil {
			return 0, err
		}

	case *WasmDeploymentPolicySpec_Config_WasmFilter_HttpUriSource:

		if h, ok := interface{}(m.GetHttpUriSource()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("HttpUriSource")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHttpUriSource(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("HttpUriSource")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *WasmDeploymentPolicySpec_Config_WasmFilter_WasmImageSource_:

		if h, ok := interface{}(m.GetWasmImageSource()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("WasmImageSource")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetWasmImageSource(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("WasmImageSource")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.FilterConfigSource.(type) {

	case *WasmDeploymentPolicySpec_Config_WasmFilter_StaticFilterConfig:

		if h, ok := interface{}(m.GetStaticFilterConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("StaticFilterConfig")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetStaticFilterConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("StaticFilterConfig")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *WasmDeploymentPolicySpec_Config_WasmFilter_DynamicFilterConfig:

		if _, err = hasher.Write([]byte(m.GetDynamicFilterConfig())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *WasmDeploymentPolicySpec_Config_WasmFilter_UriSource) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extensions.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2.WasmDeploymentPolicySpec_Config_WasmFilter_UriSource")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetUri())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSha())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *WasmDeploymentPolicySpec_Config_WasmFilter_WasmImageSource) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extensions.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/extensions.policy.gloo.solo.io/v2.WasmDeploymentPolicySpec_Config_WasmFilter_WasmImageSource")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetWasmImageTag())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
