// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/http_gateway.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
func (m *MatchableHttpGatewaySpec) Clone() proto.Message {
	var target *MatchableHttpGatewaySpec
	if m == nil {
		return target
	}
	target = &MatchableHttpGatewaySpec{}

	if h, ok := interface{}(m.GetMatcher()).(clone.Cloner); ok {
		target.Matcher = h.Clone().(*MatchableHttpGatewaySpec_Matcher)
	} else {
		target.Matcher = proto.Clone(m.GetMatcher()).(*MatchableHttpGatewaySpec_Matcher)
	}

	if h, ok := interface{}(m.GetHttpGateway()).(clone.Cloner); ok {
		target.HttpGateway = h.Clone().(*HttpGateway)
	} else {
		target.HttpGateway = proto.Clone(m.GetHttpGateway()).(*HttpGateway)
	}

	return target
}

// Clone function
func (m *HttpGateway) Clone() proto.Message {
	var target *HttpGateway
	if m == nil {
		return target
	}
	target = &HttpGateway{}

	if m.GetVirtualServices() != nil {
		target.VirtualServices = make([]*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef, len(m.GetVirtualServices()))
		for idx, v := range m.GetVirtualServices() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.VirtualServices[idx] = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			} else {
				target.VirtualServices[idx] = proto.Clone(v).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			}

		}
	}

	if m.GetVirtualServiceSelector() != nil {
		target.VirtualServiceSelector = make(map[string]string, len(m.GetVirtualServiceSelector()))
		for k, v := range m.GetVirtualServiceSelector() {

			target.VirtualServiceSelector[k] = v

		}
	}

	if h, ok := interface{}(m.GetVirtualServiceExpressions()).(clone.Cloner); ok {
		target.VirtualServiceExpressions = h.Clone().(*VirtualServiceSelectorExpressions)
	} else {
		target.VirtualServiceExpressions = proto.Clone(m.GetVirtualServiceExpressions()).(*VirtualServiceSelectorExpressions)
	}

	if m.GetVirtualServiceNamespaces() != nil {
		target.VirtualServiceNamespaces = make([]string, len(m.GetVirtualServiceNamespaces()))
		for idx, v := range m.GetVirtualServiceNamespaces() {

			target.VirtualServiceNamespaces[idx] = v

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.HttpListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.HttpListenerOptions)
	}

	return target
}

// Clone function
func (m *VirtualServiceSelectorExpressions) Clone() proto.Message {
	var target *VirtualServiceSelectorExpressions
	if m == nil {
		return target
	}
	target = &VirtualServiceSelectorExpressions{}

	if m.GetExpressions() != nil {
		target.Expressions = make([]*VirtualServiceSelectorExpressions_Expression, len(m.GetExpressions()))
		for idx, v := range m.GetExpressions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Expressions[idx] = h.Clone().(*VirtualServiceSelectorExpressions_Expression)
			} else {
				target.Expressions[idx] = proto.Clone(v).(*VirtualServiceSelectorExpressions_Expression)
			}

		}
	}

	return target
}

// Clone function
func (m *MatchableHttpGatewayStatus) Clone() proto.Message {
	var target *MatchableHttpGatewayStatus
	if m == nil {
		return target
	}
	target = &MatchableHttpGatewayStatus{}

	target.State = m.GetState()

	target.Reason = m.GetReason()

	target.ReportedBy = m.GetReportedBy()

	if m.GetSubresourceStatuses() != nil {
		target.SubresourceStatuses = make(map[string]*MatchableHttpGatewayStatus, len(m.GetSubresourceStatuses()))
		for k, v := range m.GetSubresourceStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SubresourceStatuses[k] = h.Clone().(*MatchableHttpGatewayStatus)
			} else {
				target.SubresourceStatuses[k] = proto.Clone(v).(*MatchableHttpGatewayStatus)
			}

		}
	}

	if h, ok := interface{}(m.GetDetails()).(clone.Cloner); ok {
		target.Details = h.Clone().(*github_com_golang_protobuf_ptypes_struct.Struct)
	} else {
		target.Details = proto.Clone(m.GetDetails()).(*github_com_golang_protobuf_ptypes_struct.Struct)
	}

	return target
}

// Clone function
func (m *MatchableHttpGatewayNamespacedStatuses) Clone() proto.Message {
	var target *MatchableHttpGatewayNamespacedStatuses
	if m == nil {
		return target
	}
	target = &MatchableHttpGatewayNamespacedStatuses{}

	if m.GetStatuses() != nil {
		target.Statuses = make(map[string]*MatchableHttpGatewayStatus, len(m.GetStatuses()))
		for k, v := range m.GetStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Statuses[k] = h.Clone().(*MatchableHttpGatewayStatus)
			} else {
				target.Statuses[k] = proto.Clone(v).(*MatchableHttpGatewayStatus)
			}

		}
	}

	return target
}

// Clone function
func (m *MatchableHttpGatewaySpec_Matcher) Clone() proto.Message {
	var target *MatchableHttpGatewaySpec_Matcher
	if m == nil {
		return target
	}
	target = &MatchableHttpGatewaySpec_Matcher{}

	if m.GetSourcePrefixRanges() != nil {
		target.SourcePrefixRanges = make([]*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.CidrRange, len(m.GetSourcePrefixRanges()))
		for idx, v := range m.GetSourcePrefixRanges() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SourcePrefixRanges[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.CidrRange)
			} else {
				target.SourcePrefixRanges[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.CidrRange)
			}

		}
	}

	if h, ok := interface{}(m.GetSslConfig()).(clone.Cloner); ok {
		target.SslConfig = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.SslConfig)
	} else {
		target.SslConfig = proto.Clone(m.GetSslConfig()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.SslConfig)
	}

	return target
}

// Clone function
func (m *VirtualServiceSelectorExpressions_Expression) Clone() proto.Message {
	var target *VirtualServiceSelectorExpressions_Expression
	if m == nil {
		return target
	}
	target = &VirtualServiceSelectorExpressions_Expression{}

	target.Key = m.GetKey()

	target.Operator = m.GetOperator()

	if m.GetValues() != nil {
		target.Values = make([]string, len(m.GetValues()))
		for idx, v := range m.GetValues() {

			target.Values[idx] = v

		}
	}

	return target
}
