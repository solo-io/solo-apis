// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/aws/filter.proto

package aws

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
func (m *AWSLambdaPerRoute) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/extensions/aws.AWSLambdaPerRoute")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetQualifier())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetAsync())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEmptyBodyOverride()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetEmptyBodyOverride(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetUnwrapAsAlb())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTransformerConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetTransformerConfig(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRequestTransformerConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetRequestTransformerConfig(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AWSLambdaProtocolExtension) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/extensions/aws.AWSLambdaProtocolExtension")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHost())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAccessKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSecretKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSessionToken())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRoleArn())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDisableRoleChaining())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AWSLambdaConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/extensions/aws.AWSLambdaConfig")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPropagateOriginalRouting())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetCredentialRefreshDelay()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetCredentialRefreshDelay(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	switch m.CredentialsFetcher.(type) {

	case *AWSLambdaConfig_UseDefaultCredentials:

		if h, ok := interface{}(m.GetUseDefaultCredentials()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetUseDefaultCredentials(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *AWSLambdaConfig_ServiceAccountCredentials_:

		if h, ok := interface{}(m.GetServiceAccountCredentials()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetServiceAccountCredentials(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ApiGatewayTransformation) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/extensions/aws.ApiGatewayTransformation")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AWSLambdaConfig_ServiceAccountCredentials) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.aws_lambda.v2.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/extensions/aws.AWSLambdaConfig_ServiceAccountCredentials")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCluster())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetUri())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetTimeout(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
