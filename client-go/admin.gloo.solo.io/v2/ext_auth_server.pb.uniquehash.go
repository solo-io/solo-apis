// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/ext_auth_server.proto

package v2

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
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ExtAuthServerSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.ExtAuthServerSpec")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDestinationServer()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DestinationServer")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDestinationServer(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DestinationServer")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHttpService()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HttpService")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHttpService(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HttpService")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("FailureModeAllow")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetFailureModeAllow())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRequestBody()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestBody")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestBody(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestBody")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("ClearRouteCache")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetClearRouteCache())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("StatusOnError")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetStatusOnError())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("TransportApiVersion")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetTransportApiVersion())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("StatPrefix")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetStatPrefix())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ExtAuthServerStatus) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.ExtAuthServerStatus")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ObservedGeneration")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetObservedGeneration())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("State")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetState())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ExtAuthServerSpec_HttpService) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.ExtAuthServerSpec_HttpService")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("PathPrefix")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetPathPrefix())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRequest()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Request")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequest(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Request")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResponse()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Response")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponse(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Response")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ExtAuthServerSpec_BufferSettings) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.ExtAuthServerSpec_BufferSettings")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("MaxRequestBytes")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetMaxRequestBytes())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AllowPartialMessage")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAllowPartialMessage())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("PackAsBytes")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetPackAsBytes())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ExtAuthServerSpec_HttpService_Request) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.ExtAuthServerSpec_HttpService_Request")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AllowedHeaders")); err != nil {
		return 0, err
	}
	for i, v := range m.GetAllowedHeaders() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetHeadersToAdd() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ExtAuthServerSpec_HttpService_Response) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("admin.gloo.solo.io.github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2.ExtAuthServerSpec_HttpService_Response")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AllowedUpstreamHeaders")); err != nil {
		return 0, err
	}
	for i, v := range m.GetAllowedUpstreamHeaders() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte("AllowedClientHeaders")); err != nil {
		return 0, err
	}
	for i, v := range m.GetAllowedClientHeaders() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}