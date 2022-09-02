// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/hcm/hcm.proto

package hcm

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
func (m *HttpConnectionManagerSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("hcm.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/hcm.HttpConnectionManagerSettings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSkipXffAppend()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SkipXffAppend")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSkipXffAppend(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SkipXffAppend")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetVia()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Via")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetVia(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Via")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetXffNumTrustedHops()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("XffNumTrustedHops")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetXffNumTrustedHops(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("XffNumTrustedHops")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUseRemoteAddress()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UseRemoteAddress")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUseRemoteAddress(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UseRemoteAddress")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetGenerateRequestId()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GenerateRequestId")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGenerateRequestId(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GenerateRequestId")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetProxy_100Continue()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Proxy_100Continue")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetProxy_100Continue(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Proxy_100Continue")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStreamIdleTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StreamIdleTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStreamIdleTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StreamIdleTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
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

	if h, ok := interface{}(m.GetMaxRequestHeadersKb()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxRequestHeadersKb")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxRequestHeadersKb(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxRequestHeadersKb")); err != nil {
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

	if h, ok := interface{}(m.GetRequestHeadersTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RequestHeadersTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequestHeadersTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RequestHeadersTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDrainTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DrainTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDrainTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DrainTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDelayedCloseTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DelayedCloseTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDelayedCloseTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DelayedCloseTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetServerName()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ServerName")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetServerName(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ServerName")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStripAnyHostPort()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StripAnyHostPort")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStripAnyHostPort(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StripAnyHostPort")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAcceptHttp_10()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AcceptHttp_10")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAcceptHttp_10(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AcceptHttp_10")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDefaultHostForHttp_10()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DefaultHostForHttp_10")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDefaultHostForHttp_10(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DefaultHostForHttp_10")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAllowChunkedLength()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AllowChunkedLength")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAllowChunkedLength(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AllowChunkedLength")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEnableTrailers()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EnableTrailers")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnableTrailers(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EnableTrailers")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTracing()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Tracing")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTracing(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Tracing")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetForwardClientCertDetails())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSetCurrentClientCertDetails()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SetCurrentClientCertDetails")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSetCurrentClientCertDetails(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SetCurrentClientCertDetails")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPreserveExternalRequestId()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PreserveExternalRequestId")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPreserveExternalRequestId(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PreserveExternalRequestId")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetUpgrades() {

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

	if h, ok := interface{}(m.GetMaxConnectionDuration()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxConnectionDuration")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxConnectionDuration(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxConnectionDuration")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
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

	if h, ok := interface{}(m.GetMaxHeadersCount()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxHeadersCount")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxHeadersCount(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxHeadersCount")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetHeadersWithUnderscoresAction())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMaxRequestsPerConnection()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxRequestsPerConnection")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxRequestsPerConnection(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxRequestsPerConnection")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetServerHeaderTransformation())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPathWithEscapedSlashesAction())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetCodecType())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMergeSlashes()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MergeSlashes")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMergeSlashes(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MergeSlashes")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetNormalizePath()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("NormalizePath")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetNormalizePath(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("NormalizePath")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUuidRequestIdConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UuidRequestIdConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUuidRequestIdConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UuidRequestIdConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHttp2ProtocolOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Http2ProtocolOptions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHttp2ProtocolOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Http2ProtocolOptions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.HeaderFormat.(type) {

	case *HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat:

		if h, ok := interface{}(m.GetProperCaseHeaderKeyFormat()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ProperCaseHeaderKeyFormat")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetProperCaseHeaderKeyFormat(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ProperCaseHeaderKeyFormat")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat:

		if h, ok := interface{}(m.GetPreserveCaseHeaderKeyFormat()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("PreserveCaseHeaderKeyFormat")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetPreserveCaseHeaderKeyFormat(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("PreserveCaseHeaderKeyFormat")); err != nil {
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
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("hcm.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/hcm.HttpConnectionManagerSettings_SetCurrentClientCertDetails")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSubject()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Subject")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSubject(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Subject")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCert()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Cert")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCert(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Cert")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetChain()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Chain")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetChain(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Chain")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDns()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Dns")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDns(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Dns")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUri()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Uri")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUri(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Uri")); err != nil {
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
func (m *HttpConnectionManagerSettings_UuidRequestIdConfigSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("hcm.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/hcm.HttpConnectionManagerSettings_UuidRequestIdConfigSettings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetPackTraceReason()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PackTraceReason")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPackTraceReason(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PackTraceReason")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUseRequestIdForTraceSampling()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UseRequestIdForTraceSampling")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUseRequestIdForTraceSampling(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UseRequestIdForTraceSampling")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
