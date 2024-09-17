// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/aws/aws.proto

package aws

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
func (m *UpstreamSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("aws.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/aws.UpstreamSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Region")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

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

	if _, err = hasher.Write([]byte("LambdaFunctions")); err != nil {
		return 0, err
	}
	for i, v := range m.GetLambdaFunctions() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if _, err = hasher.Write([]byte("RoleArn")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRoleArn())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AwsAccountId")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetAwsAccountId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DisableRoleChaining")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetDisableRoleChaining())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDestinationOverrides()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DestinationOverrides")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDestinationOverrides(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DestinationOverrides")); err != nil {
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
func (m *LambdaFunctionSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("aws.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/aws.LambdaFunctionSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("LogicalName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetLogicalName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("LambdaFunctionName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetLambdaFunctionName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Qualifier")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetQualifier())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *DestinationSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("aws.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/aws.DestinationSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("LogicalName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetLogicalName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("InvocationStyle")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetInvocationStyle())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RequestTransformation")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetRequestTransformation())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetResponseTransformation())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("UnwrapAsAlb")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetUnwrapAsAlb())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("UnwrapAsApiGateway")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetUnwrapAsApiGateway())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("WrapAsApiGateway")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetWrapAsApiGateway())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
