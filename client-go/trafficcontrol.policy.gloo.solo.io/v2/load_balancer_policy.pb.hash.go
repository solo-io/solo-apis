// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/trafficcontrol/load_balancer_policy.proto

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
func (m *LoadBalancerPolicySpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("trafficcontrol.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/trafficcontrol.policy.gloo.solo.io/v2.LoadBalancerPolicySpec")); err != nil {
		return 0, err
	}

	for _, v := range m.GetApplyToDestinations() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *LoadBalancerPolicyStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("trafficcontrol.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/trafficcontrol.policy.gloo.solo.io/v2.LoadBalancerPolicyStatus")); err != nil {
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

	err = binary.Write(hasher, binary.LittleEndian, m.GetNumSelectedDestinationPorts())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *LoadBalancerPolicyReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("trafficcontrol.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/trafficcontrol.policy.gloo.solo.io/v2.LoadBalancerPolicyReport")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetWorkspaces() {
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

	for _, v := range m.GetSelectedDestinationPorts() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *LoadBalancerPolicySpec_Config) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("trafficcontrol.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/trafficcontrol.policy.gloo.solo.io/v2.LoadBalancerPolicySpec_Config")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetWarmupDurationSecs()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("WarmupDurationSecs")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetWarmupDurationSecs(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("WarmupDurationSecs")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.LbPolicy.(type) {

	case *LoadBalancerPolicySpec_Config_Simple:

		err = binary.Write(hasher, binary.LittleEndian, m.GetSimple())
		if err != nil {
			return 0, err
		}

	case *LoadBalancerPolicySpec_Config_ConsistentHash:

		if h, ok := interface{}(m.GetConsistentHash()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ConsistentHash")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetConsistentHash(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ConsistentHash")); err != nil {
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
func (m *LoadBalancerPolicySpec_Config_ConsistentHashLB) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("trafficcontrol.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/trafficcontrol.policy.gloo.solo.io/v2.LoadBalancerPolicySpec_Config_ConsistentHashLB")); err != nil {
		return 0, err
	}

	switch m.HashKey.(type) {

	case *LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpHeaderName:

		if _, err = hasher.Write([]byte(m.GetHttpHeaderName())); err != nil {
			return 0, err
		}

	case *LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpCookie:

		if h, ok := interface{}(m.GetHttpCookie()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("HttpCookie")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHttpCookie(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("HttpCookie")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *LoadBalancerPolicySpec_Config_ConsistentHashLB_UseSourceIp:

		err = binary.Write(hasher, binary.LittleEndian, m.GetUseSourceIp())
		if err != nil {
			return 0, err
		}

	case *LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpQueryParameterName:

		if _, err = hasher.Write([]byte(m.GetHttpQueryParameterName())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *LoadBalancerPolicySpec_Config_ConsistentHashLB_HTTPCookie) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("trafficcontrol.policy.gloo.solo.io.github.com/solo-io/solo-apis/client-go/trafficcontrol.policy.gloo.solo.io/v2.LoadBalancerPolicySpec_Config_ConsistentHashLB_HTTPCookie")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPath())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTtl()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Ttl")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTtl(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Ttl")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
