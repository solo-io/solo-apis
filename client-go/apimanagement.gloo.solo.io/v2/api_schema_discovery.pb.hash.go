// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/apimanagement/v2/api_schema_discovery.proto

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
func (m *ApiSchemaDiscoverySpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/client-go/apimanagement.gloo.solo.io/v2.ApiSchemaDiscoverySpec")); err != nil {
		return 0, err
	}

	for _, v := range m.GetServedBy() {

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

	switch m.FetchSchemaType.(type) {

	case *ApiSchemaDiscoverySpec_Openapi:

		if h, ok := interface{}(m.GetOpenapi()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Openapi")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOpenapi(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Openapi")); err != nil {
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
func (m *ApiSchemaDiscoveryStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/client-go/apimanagement.gloo.solo.io/v2.ApiSchemaDiscoveryStatus")); err != nil {
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

	if _, err = hasher.Write([]byte(m.GetOwnerWorkspace())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ApiSchemaDiscoveryReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/client-go/apimanagement.gloo.solo.io/v2.ApiSchemaDiscoveryReport")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetOwnerWorkspace())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ApiSchemaDiscoverySpec_OpenAPI) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/client-go/apimanagement.gloo.solo.io/v2.ApiSchemaDiscoverySpec_OpenAPI")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFetchEndpoint()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FetchEndpoint")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFetchEndpoint(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FetchEndpoint")); err != nil {
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
func (m *ApiSchemaDiscoverySpec_FetchEndpoint) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("apimanagement.gloo.solo.io.github.com/solo-io/solo-apis/client-go/apimanagement.gloo.solo.io/v2.ApiSchemaDiscoverySpec_FetchEndpoint")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetUrl())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRetryDelay()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RetryDelay")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRetryDelay(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RetryDelay")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPullAttempts())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetUseBackoff())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
