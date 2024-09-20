// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/protocol/protocol.proto

package protocol

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

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
	_ = strconv.Itoa
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *HttpProtocolOptions) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("protocol.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/protocol.HttpProtocolOptions")); err != nil {
		return 0, err
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

	if _, err = hasher.Write([]byte("MaxHeadersCount")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetMaxHeadersCount())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxStreamDuration")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxStreamDuration(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxStreamDuration")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("HeadersWithUnderscoresAction")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetHeadersWithUnderscoresAction())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Http1ProtocolOptions) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("protocol.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/protocol.Http1ProtocolOptions")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("EnableTrailers")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetEnableTrailers())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OverrideStreamErrorOnInvalidHttpMessage")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOverrideStreamErrorOnInvalidHttpMessage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OverrideStreamErrorOnInvalidHttpMessage")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.HeaderFormat.(type) {

	case *Http1ProtocolOptions_ProperCaseHeaderKeyFormat:

		if _, err = hasher.Write([]byte("ProperCaseHeaderKeyFormat")); err != nil {
			return 0, err
		}
		err = binary.Write(hasher, binary.LittleEndian, m.GetProperCaseHeaderKeyFormat())
		if err != nil {
			return 0, err
		}

	case *Http1ProtocolOptions_PreserveCaseHeaderKeyFormat:

		if _, err = hasher.Write([]byte("PreserveCaseHeaderKeyFormat")); err != nil {
			return 0, err
		}
		err = binary.Write(hasher, binary.LittleEndian, m.GetPreserveCaseHeaderKeyFormat())
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Http2ProtocolOptions) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("protocol.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/protocol.Http2ProtocolOptions")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMaxConcurrentStreams()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxConcurrentStreams")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxConcurrentStreams(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxConcurrentStreams")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetInitialStreamWindowSize()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("InitialStreamWindowSize")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetInitialStreamWindowSize(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("InitialStreamWindowSize")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetInitialConnectionWindowSize()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("InitialConnectionWindowSize")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetInitialConnectionWindowSize(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("InitialConnectionWindowSize")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OverrideStreamErrorOnInvalidHttpMessage")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOverrideStreamErrorOnInvalidHttpMessage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OverrideStreamErrorOnInvalidHttpMessage")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
