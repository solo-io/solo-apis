// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/failover_scheme.proto

package v1

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
func (m *FailoverScheme) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.FailoverScheme")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Metadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Spec")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSpec(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Spec")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Status")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Status")); err != nil {
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
func (m *GetFailoverSchemeRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFailoverSchemeRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetUpstreamRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpstreamRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
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
func (m *GetFailoverSchemeResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFailoverSchemeResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFailoverScheme()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FailoverScheme")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFailoverScheme(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FailoverScheme")); err != nil {
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
func (m *GetFailoverSchemeYamlRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFailoverSchemeYamlRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFailoverSchemeRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FailoverSchemeRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFailoverSchemeRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FailoverSchemeRef")); err != nil {
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
func (m *GetFailoverSchemeYamlResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFailoverSchemeYamlResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetYamlData()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("YamlData")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetYamlData(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("YamlData")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
