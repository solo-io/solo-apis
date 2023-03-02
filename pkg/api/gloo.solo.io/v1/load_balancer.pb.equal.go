// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/load_balancer.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *LoadBalancerConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LoadBalancerConfig)
	if !ok {
		that2, ok := that.(LoadBalancerConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetHealthyPanicThreshold()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHealthyPanicThreshold()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHealthyPanicThreshold(), target.GetHealthyPanicThreshold()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUpdateMergeWindow()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpdateMergeWindow()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpdateMergeWindow(), target.GetUpdateMergeWindow()) {
			return false
		}
	}

	switch m.Type.(type) {

	case *LoadBalancerConfig_RoundRobin_:
		if _, ok := target.Type.(*LoadBalancerConfig_RoundRobin_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRoundRobin()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRoundRobin()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRoundRobin(), target.GetRoundRobin()) {
				return false
			}
		}

	case *LoadBalancerConfig_LeastRequest_:
		if _, ok := target.Type.(*LoadBalancerConfig_LeastRequest_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLeastRequest()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLeastRequest()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLeastRequest(), target.GetLeastRequest()) {
				return false
			}
		}

	case *LoadBalancerConfig_Random_:
		if _, ok := target.Type.(*LoadBalancerConfig_Random_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRandom()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRandom()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRandom(), target.GetRandom()) {
				return false
			}
		}

	case *LoadBalancerConfig_RingHash_:
		if _, ok := target.Type.(*LoadBalancerConfig_RingHash_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRingHash()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRingHash()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRingHash(), target.GetRingHash()) {
				return false
			}
		}

	case *LoadBalancerConfig_Maglev_:
		if _, ok := target.Type.(*LoadBalancerConfig_Maglev_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMaglev()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMaglev()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMaglev(), target.GetMaglev()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Type != target.Type {
			return false
		}
	}

	switch m.LocalityConfig.(type) {

	case *LoadBalancerConfig_LocalityWeightedLbConfig:
		if _, ok := target.LocalityConfig.(*LoadBalancerConfig_LocalityWeightedLbConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLocalityWeightedLbConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocalityWeightedLbConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLocalityWeightedLbConfig(), target.GetLocalityWeightedLbConfig()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.LocalityConfig != target.LocalityConfig {
			return false
		}
	}

	return true
}

// Equal function
func (m *LoadBalancerConfig_RoundRobin) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LoadBalancerConfig_RoundRobin)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RoundRobin)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSlowStartConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSlowStartConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSlowStartConfig(), target.GetSlowStartConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *LoadBalancerConfig_LeastRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LoadBalancerConfig_LeastRequest)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_LeastRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetChoiceCount() != target.GetChoiceCount() {
		return false
	}

	if h, ok := interface{}(m.GetSlowStartConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSlowStartConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSlowStartConfig(), target.GetSlowStartConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *LoadBalancerConfig_Random) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LoadBalancerConfig_Random)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Random)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *LoadBalancerConfig_RingHashConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LoadBalancerConfig_RingHashConfig)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHashConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetMinimumRingSize() != target.GetMinimumRingSize() {
		return false
	}

	if m.GetMaximumRingSize() != target.GetMaximumRingSize() {
		return false
	}

	return true
}

// Equal function
func (m *LoadBalancerConfig_RingHash) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LoadBalancerConfig_RingHash)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHash)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRingHashConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRingHashConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRingHashConfig(), target.GetRingHashConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *LoadBalancerConfig_Maglev) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LoadBalancerConfig_Maglev)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Maglev)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *LoadBalancerConfig_SlowStartConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LoadBalancerConfig_SlowStartConfig)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_SlowStartConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSlowStartWindow()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSlowStartWindow()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSlowStartWindow(), target.GetSlowStartWindow()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAggression()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAggression()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAggression(), target.GetAggression()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMinWeightPercent()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMinWeightPercent()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMinWeightPercent(), target.GetMinWeightPercent()) {
			return false
		}
	}

	return true
}
