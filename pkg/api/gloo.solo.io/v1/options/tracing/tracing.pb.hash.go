// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/tracing/tracing.proto

package tracing

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
func (m *ListenerTracingSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("tracing.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/tracing.ListenerTracingSettings")); err != nil {
		return 0, err
	}

	for _, v := range m.GetRequestHeadersForTags() {

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

	if h, ok := interface{}(m.GetVerbose()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Verbose")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetVerbose(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Verbose")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTracePercentages()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TracePercentages")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTracePercentages(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TracePercentages")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetEnvironmentVariablesForTags() {

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

	for _, v := range m.GetLiteralsForTags() {

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

	switch m.ProviderConfig.(type) {

	case *ListenerTracingSettings_ZipkinConfig:

		if h, ok := interface{}(m.GetZipkinConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ZipkinConfig")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetZipkinConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ZipkinConfig")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ListenerTracingSettings_DatadogConfig:

		if h, ok := interface{}(m.GetDatadogConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DatadogConfig")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDatadogConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DatadogConfig")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ListenerTracingSettings_OpenTelemetryConfig:

		if h, ok := interface{}(m.GetOpenTelemetryConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("OpenTelemetryConfig")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOpenTelemetryConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("OpenTelemetryConfig")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ListenerTracingSettings_OpenCensusConfig:

		if h, ok := interface{}(m.GetOpenCensusConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("OpenCensusConfig")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOpenCensusConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("OpenCensusConfig")); err != nil {
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
func (m *RouteTracingSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("tracing.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/tracing.RouteTracingSettings")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRouteDescriptor())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTracePercentages()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TracePercentages")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTracePercentages(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TracePercentages")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPropagate()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Propagate")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPropagate(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Propagate")); err != nil {
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
func (m *TracePercentages) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("tracing.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/tracing.TracePercentages")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetClientSamplePercentage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ClientSamplePercentage")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetClientSamplePercentage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ClientSamplePercentage")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRandomSamplePercentage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RandomSamplePercentage")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRandomSamplePercentage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RandomSamplePercentage")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetOverallSamplePercentage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OverallSamplePercentage")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOverallSamplePercentage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OverallSamplePercentage")); err != nil {
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
func (m *TracingTagEnvironmentVariable) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("tracing.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/tracing.TracingTagEnvironmentVariable")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTag()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Tag")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTag(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Tag")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetName()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Name")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetName(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Name")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDefaultValue()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DefaultValue")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDefaultValue(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DefaultValue")); err != nil {
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
func (m *TracingTagLiteral) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("tracing.options.gloo.solo.io.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/tracing.TracingTagLiteral")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTag()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Tag")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTag(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Tag")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetValue()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Value")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetValue(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Value")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
