// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/trafficcontrol/load_balancer_policy.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *LoadBalancerPolicySpec) Clone() proto.Message {
	var target *LoadBalancerPolicySpec
	if m == nil {
		return target
	}
	target = &LoadBalancerPolicySpec{}

	if m.GetApplyToDestinations() != nil {
		target.ApplyToDestinations = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationSelector, len(m.GetApplyToDestinations()))
		for idx, v := range m.GetApplyToDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ApplyToDestinations[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationSelector)
			} else {
				target.ApplyToDestinations[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*LoadBalancerPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*LoadBalancerPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *LoadBalancerPolicyStatus) Clone() proto.Message {
	var target *LoadBalancerPolicyStatus
	if m == nil {
		return target
	}
	target = &LoadBalancerPolicyStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	target.NumSelectedDestinationPorts = m.GetNumSelectedDestinationPorts()

	return target
}

// Clone function
func (m *LoadBalancerPolicyReport) Clone() proto.Message {
	var target *LoadBalancerPolicyReport
	if m == nil {
		return target
	}
	target = &LoadBalancerPolicyReport{}

	if m.GetWorkspaces() != nil {
		target.Workspaces = make(map[string]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Report, len(m.GetWorkspaces()))
		for k, v := range m.GetWorkspaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workspaces[k] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Report)
			} else {
				target.Workspaces[k] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Report)
			}

		}
	}

	if m.GetSelectedDestinationPorts() != nil {
		target.SelectedDestinationPorts = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference, len(m.GetSelectedDestinationPorts()))
		for idx, v := range m.GetSelectedDestinationPorts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedDestinationPorts[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.SelectedDestinationPorts[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	return target
}

// Clone function
func (m *LoadBalancerPolicySpec_Config) Clone() proto.Message {
	var target *LoadBalancerPolicySpec_Config
	if m == nil {
		return target
	}
	target = &LoadBalancerPolicySpec_Config{}

	if h, ok := interface{}(m.GetWarmupDurationSecs()).(clone.Cloner); ok {
		target.WarmupDurationSecs = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.WarmupDurationSecs = proto.Clone(m.GetWarmupDurationSecs()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetHealthyPanicThreshold()).(clone.Cloner); ok {
		target.HealthyPanicThreshold = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.DoubleValue)
	} else {
		target.HealthyPanicThreshold = proto.Clone(m.GetHealthyPanicThreshold()).(*github_com_golang_protobuf_ptypes_wrappers.DoubleValue)
	}

	if h, ok := interface{}(m.GetUpdateMergeWindow()).(clone.Cloner); ok {
		target.UpdateMergeWindow = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.UpdateMergeWindow = proto.Clone(m.GetUpdateMergeWindow()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	switch m.LbPolicy.(type) {

	case *LoadBalancerPolicySpec_Config_Simple:

		target.LbPolicy = &LoadBalancerPolicySpec_Config_Simple{
			Simple: m.GetSimple(),
		}

	case *LoadBalancerPolicySpec_Config_ConsistentHash:

		if h, ok := interface{}(m.GetConsistentHash()).(clone.Cloner); ok {
			target.LbPolicy = &LoadBalancerPolicySpec_Config_ConsistentHash{
				ConsistentHash: h.Clone().(*LoadBalancerPolicySpec_Config_ConsistentHashLB),
			}
		} else {
			target.LbPolicy = &LoadBalancerPolicySpec_Config_ConsistentHash{
				ConsistentHash: proto.Clone(m.GetConsistentHash()).(*LoadBalancerPolicySpec_Config_ConsistentHashLB),
			}
		}

	}

	return target
}

// Clone function
func (m *LoadBalancerPolicySpec_Config_ConsistentHashLB) Clone() proto.Message {
	var target *LoadBalancerPolicySpec_Config_ConsistentHashLB
	if m == nil {
		return target
	}
	target = &LoadBalancerPolicySpec_Config_ConsistentHashLB{}

	switch m.HashKey.(type) {

	case *LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpHeaderName:

		target.HashKey = &LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpHeaderName{
			HttpHeaderName: m.GetHttpHeaderName(),
		}

	case *LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpCookie:

		if h, ok := interface{}(m.GetHttpCookie()).(clone.Cloner); ok {
			target.HashKey = &LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpCookie{
				HttpCookie: h.Clone().(*LoadBalancerPolicySpec_Config_ConsistentHashLB_HTTPCookie),
			}
		} else {
			target.HashKey = &LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpCookie{
				HttpCookie: proto.Clone(m.GetHttpCookie()).(*LoadBalancerPolicySpec_Config_ConsistentHashLB_HTTPCookie),
			}
		}

	case *LoadBalancerPolicySpec_Config_ConsistentHashLB_UseSourceIp:

		target.HashKey = &LoadBalancerPolicySpec_Config_ConsistentHashLB_UseSourceIp{
			UseSourceIp: m.GetUseSourceIp(),
		}

	case *LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpQueryParameterName:

		target.HashKey = &LoadBalancerPolicySpec_Config_ConsistentHashLB_HttpQueryParameterName{
			HttpQueryParameterName: m.GetHttpQueryParameterName(),
		}

	}

	return target
}

// Clone function
func (m *LoadBalancerPolicySpec_Config_ConsistentHashLB_HTTPCookie) Clone() proto.Message {
	var target *LoadBalancerPolicySpec_Config_ConsistentHashLB_HTTPCookie
	if m == nil {
		return target
	}
	target = &LoadBalancerPolicySpec_Config_ConsistentHashLB_HTTPCookie{}

	target.Name = m.GetName()

	target.Path = m.GetPath()

	if h, ok := interface{}(m.GetTtl()).(clone.Cloner); ok {
		target.Ttl = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.Ttl = proto.Clone(m.GetTtl()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}