// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed/v1/instance.proto

package types

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
func (m *GlooInstanceSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCluster())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetIsEnterprise())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetControlPlane()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ControlPlane")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetControlPlane(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ControlPlane")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetProxies() {

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

	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAdmin()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Admin")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAdmin(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Admin")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCheck()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Check")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCheck(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Check")); err != nil {
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
func (m *GlooInstanceStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceStatus")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GlooInstanceSpec_ControlPlane) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceSpec_ControlPlane")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetVersion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetNamespace())); err != nil {
		return 0, err
	}

	for _, v := range m.GetWatchedNamespaces() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GlooInstanceSpec_Proxy) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceSpec_Proxy")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetReplicas())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetAvailableReplicas())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetReadyReplicas())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetWasmEnabled())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetReadConfigMulticlusterEnabled())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetVersion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetNamespace())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetWorkloadControllerType())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetZones() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetIngressEndpoints() {

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
func (m *GlooInstanceSpec_Admin) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceSpec_Admin")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetWriteNamespace())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetProxyId()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ProxyId")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetProxyId(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ProxyId")); err != nil {
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
func (m *GlooInstanceSpec_Check) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceSpec_Check")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetGateways()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Gateways")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGateways(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Gateways")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetVirtualServices()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("VirtualServices")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetVirtualServices(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("VirtualServices")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRouteTables()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RouteTables")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRouteTables(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RouteTables")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAuthConfigs()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AuthConfigs")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAuthConfigs(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AuthConfigs")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSettings()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Settings")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSettings(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Settings")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUpstreams()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Upstreams")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpstreams(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Upstreams")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUpstreamGroups()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UpstreamGroups")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpstreamGroups(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UpstreamGroups")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetProxies()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Proxies")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetProxies(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Proxies")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRateLimitConfigs()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RateLimitConfigs")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRateLimitConfigs(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RateLimitConfigs")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMatchableHttpGateways()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MatchableHttpGateways")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatchableHttpGateways(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MatchableHttpGateways")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMatchableTcpGateways()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MatchableTcpGateways")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatchableTcpGateways(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MatchableTcpGateways")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDeployments()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Deployments")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDeployments(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Deployments")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPods()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Pods")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPods(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Pods")); err != nil {
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
func (m *GlooInstanceSpec_Proxy_IngressEndpoint) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceSpec_Proxy_IngressEndpoint")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetAddress())); err != nil {
		return 0, err
	}

	for _, v := range m.GetPorts() {

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

	if _, err = hasher.Write([]byte(m.GetServiceName())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GlooInstanceSpec_Proxy_IngressEndpoint_Port) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceSpec_Proxy_IngressEndpoint_Port")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPort())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GlooInstanceSpec_Check_Summary) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceSpec_Check_Summary")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTotal())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetErrors() {

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

	for _, v := range m.GetWarnings() {

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
func (m *GlooInstanceSpec_Check_Summary_ResourceReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types.GlooInstanceSpec_Check_Summary_ResourceReport")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Ref")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Ref")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetMessage())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
