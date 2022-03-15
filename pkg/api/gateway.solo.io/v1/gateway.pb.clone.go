// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/gateway.proto

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

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

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
func (m *GatewaySpec) Clone() proto.Message {
	var target *GatewaySpec
	if m == nil {
		return target
	}
	target = &GatewaySpec{}

	target.Ssl = m.GetSsl()

	target.BindAddress = m.GetBindAddress()

	target.BindPort = m.GetBindPort()

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.ListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.ListenerOptions)
	}

	if h, ok := interface{}(m.GetUseProxyProto()).(clone.Cloner); ok {
		target.UseProxyProto = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.UseProxyProto = proto.Clone(m.GetUseProxyProto()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if m.GetProxyNames() != nil {
		target.ProxyNames = make([]string, len(m.GetProxyNames()))
		for idx, v := range m.GetProxyNames() {

			target.ProxyNames[idx] = v

		}
	}

	if h, ok := interface{}(m.GetRouteOptions()).(clone.Cloner); ok {
		target.RouteOptions = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RouteConfigurationOptions)
	} else {
		target.RouteOptions = proto.Clone(m.GetRouteOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RouteConfigurationOptions)
	}

	switch m.GatewayType.(type) {

	case *GatewaySpec_HttpGateway:

		if h, ok := interface{}(m.GetHttpGateway()).(clone.Cloner); ok {
			target.GatewayType = &GatewaySpec_HttpGateway{
				HttpGateway: h.Clone().(*HttpGateway),
			}
		} else {
			target.GatewayType = &GatewaySpec_HttpGateway{
				HttpGateway: proto.Clone(m.GetHttpGateway()).(*HttpGateway),
			}
		}

	case *GatewaySpec_TcpGateway:

		if h, ok := interface{}(m.GetTcpGateway()).(clone.Cloner); ok {
			target.GatewayType = &GatewaySpec_TcpGateway{
				TcpGateway: h.Clone().(*TcpGateway),
			}
		} else {
			target.GatewayType = &GatewaySpec_TcpGateway{
				TcpGateway: proto.Clone(m.GetTcpGateway()).(*TcpGateway),
			}
		}

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
func (m *TcpGateway) Clone() proto.Message {
	var target *TcpGateway
	if m == nil {
		return target
	}
	target = &TcpGateway{}

	if m.GetTcpHosts() != nil {
		target.TcpHosts = make([]*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.TcpHost, len(m.GetTcpHosts()))
		for idx, v := range m.GetTcpHosts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TcpHosts[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.TcpHost)
			} else {
				target.TcpHosts[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.TcpHost)
			}

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.TcpListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.TcpListenerOptions)
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
func (m *GatewayStatus) Clone() proto.Message {
	var target *GatewayStatus
	if m == nil {
		return target
	}
	target = &GatewayStatus{}

	target.State = m.GetState()

	target.Reason = m.GetReason()

	target.ReportedBy = m.GetReportedBy()

	if m.GetSubresourceStatuses() != nil {
		target.SubresourceStatuses = make(map[string]*GatewayStatus, len(m.GetSubresourceStatuses()))
		for k, v := range m.GetSubresourceStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SubresourceStatuses[k] = h.Clone().(*GatewayStatus)
			} else {
				target.SubresourceStatuses[k] = proto.Clone(v).(*GatewayStatus)
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
