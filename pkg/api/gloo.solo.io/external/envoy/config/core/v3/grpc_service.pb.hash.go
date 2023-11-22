// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/core/v3/grpc_service.proto

package v3

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
func (m *GrpcService) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Timeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Timeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetInitialMetadata() {

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

	switch m.TargetSpecifier.(type) {

	case *GrpcService_EnvoyGrpc_:

		if h, ok := interface{}(m.GetEnvoyGrpc()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("EnvoyGrpc")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetEnvoyGrpc(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("EnvoyGrpc")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GrpcService_GoogleGrpc_:

		if h, ok := interface{}(m.GetGoogleGrpc()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GoogleGrpc")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGoogleGrpc(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GoogleGrpc")); err != nil {
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
func (m *GrpcService_EnvoyGrpc) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_EnvoyGrpc")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetClusterName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAuthority())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRetryPolicy()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RetryPolicy")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRetryPolicy(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RetryPolicy")); err != nil {
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
func (m *GrpcService_GoogleGrpc) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTargetUri())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetChannelCredentials()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ChannelCredentials")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetChannelCredentials(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ChannelCredentials")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetCallCredentials() {

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

	if _, err = hasher.Write([]byte(m.GetStatPrefix())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCredentialsFactoryName())); err != nil {
		return 0, err
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

	if h, ok := interface{}(m.GetPerStreamBufferLimitBytes()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PerStreamBufferLimitBytes")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPerStreamBufferLimitBytes(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PerStreamBufferLimitBytes")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetChannelArgs()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ChannelArgs")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetChannelArgs(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ChannelArgs")); err != nil {
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
func (m *GrpcService_GoogleGrpc_SslCredentials) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_SslCredentials")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRootCerts()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RootCerts")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRootCerts(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RootCerts")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPrivateKey()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PrivateKey")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPrivateKey(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PrivateKey")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCertChain()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CertChain")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCertChain(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CertChain")); err != nil {
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
func (m *GrpcService_GoogleGrpc_GoogleLocalCredentials) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_GoogleLocalCredentials")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GrpcService_GoogleGrpc_ChannelCredentials) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_ChannelCredentials")); err != nil {
		return 0, err
	}

	switch m.CredentialSpecifier.(type) {

	case *GrpcService_GoogleGrpc_ChannelCredentials_SslCredentials:

		if h, ok := interface{}(m.GetSslCredentials()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("SslCredentials")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSslCredentials(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("SslCredentials")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GrpcService_GoogleGrpc_ChannelCredentials_GoogleDefault:

		if h, ok := interface{}(m.GetGoogleDefault()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GoogleDefault")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGoogleDefault(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GoogleDefault")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GrpcService_GoogleGrpc_ChannelCredentials_LocalCredentials:

		if h, ok := interface{}(m.GetLocalCredentials()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("LocalCredentials")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetLocalCredentials(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("LocalCredentials")); err != nil {
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
func (m *GrpcService_GoogleGrpc_CallCredentials) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_CallCredentials")); err != nil {
		return 0, err
	}

	switch m.CredentialSpecifier.(type) {

	case *GrpcService_GoogleGrpc_CallCredentials_AccessToken:

		if _, err = hasher.Write([]byte(m.GetAccessToken())); err != nil {
			return 0, err
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleComputeEngine:

		if h, ok := interface{}(m.GetGoogleComputeEngine()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GoogleComputeEngine")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGoogleComputeEngine(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GoogleComputeEngine")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleRefreshToken:

		if _, err = hasher.Write([]byte(m.GetGoogleRefreshToken())); err != nil {
			return 0, err
		}

	case *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJwtAccess:

		if h, ok := interface{}(m.GetServiceAccountJwtAccess()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ServiceAccountJwtAccess")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetServiceAccountJwtAccess(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ServiceAccountJwtAccess")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleIam:

		if h, ok := interface{}(m.GetGoogleIam()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GoogleIam")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGoogleIam(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GoogleIam")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_FromPlugin:

		if h, ok := interface{}(m.GetFromPlugin()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("FromPlugin")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetFromPlugin(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("FromPlugin")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_StsService_:

		if h, ok := interface{}(m.GetStsService()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("StsService")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetStsService(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("StsService")); err != nil {
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
func (m *GrpcService_GoogleGrpc_ChannelArgs) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_ChannelArgs")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetArgs() {
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

	return hasher.Sum64(), nil
}

// Hash function
func (m *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetJsonKey())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTokenLifetimeSeconds())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAuthorizationToken())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAuthoritySelector())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	switch m.ConfigType.(type) {

	case *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin_TypedConfig:

		if h, ok := interface{}(m.GetTypedConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("TypedConfig")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTypedConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("TypedConfig")); err != nil {
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
func (m *GrpcService_GoogleGrpc_CallCredentials_StsService) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_CallCredentials_StsService")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTokenExchangeServiceUri())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetResource())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAudience())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetScope())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRequestedTokenType())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSubjectTokenPath())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetSubjectTokenType())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetActorTokenPath())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetActorTokenType())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GrpcService_GoogleGrpc_ChannelArgs_Value) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.GrpcService_GoogleGrpc_ChannelArgs_Value")); err != nil {
		return 0, err
	}

	switch m.ValueSpecifier.(type) {

	case *GrpcService_GoogleGrpc_ChannelArgs_Value_StringValue:

		if _, err = hasher.Write([]byte(m.GetStringValue())); err != nil {
			return 0, err
		}

	case *GrpcService_GoogleGrpc_ChannelArgs_Value_IntValue:

		err = binary.Write(hasher, binary.LittleEndian, m.GetIntValue())
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
