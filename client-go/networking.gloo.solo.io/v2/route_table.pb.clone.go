// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/networking/v2/route_table.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"

	github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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
func (m *RouteTableSpec) Clone() proto.Message {
	var target *RouteTableSpec
	if m == nil {
		return target
	}
	target = &RouteTableSpec{}

	if m.GetHosts() != nil {
		target.Hosts = make([]string, len(m.GetHosts()))
		for idx, v := range m.GetHosts() {

			target.Hosts[idx] = v

		}
	}

	if m.GetVirtualGateways() != nil {
		target.VirtualGateways = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference, len(m.GetVirtualGateways()))
		for idx, v := range m.GetVirtualGateways() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.VirtualGateways[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.VirtualGateways[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	if m.GetWorkloadSelectors() != nil {
		target.WorkloadSelectors = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.WorkloadSelector, len(m.GetWorkloadSelectors()))
		for idx, v := range m.GetWorkloadSelectors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.WorkloadSelectors[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.WorkloadSelector)
			} else {
				target.WorkloadSelectors[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.WorkloadSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetDefaultDestination()).(clone.Cloner); ok {
		target.DefaultDestination = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
	} else {
		target.DefaultDestination = proto.Clone(m.GetDefaultDestination()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
	}

	if m.GetHttp() != nil {
		target.Http = make([]*HTTPRoute, len(m.GetHttp()))
		for idx, v := range m.GetHttp() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Http[idx] = h.Clone().(*HTTPRoute)
			} else {
				target.Http[idx] = proto.Clone(v).(*HTTPRoute)
			}

		}
	}

	if m.GetTcp() != nil {
		target.Tcp = make([]*TCPRoute, len(m.GetTcp()))
		for idx, v := range m.GetTcp() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Tcp[idx] = h.Clone().(*TCPRoute)
			} else {
				target.Tcp[idx] = proto.Clone(v).(*TCPRoute)
			}

		}
	}

	if m.GetTls() != nil {
		target.Tls = make([]*TLSRoute, len(m.GetTls()))
		for idx, v := range m.GetTls() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Tls[idx] = h.Clone().(*TLSRoute)
			} else {
				target.Tls[idx] = proto.Clone(v).(*TLSRoute)
			}

		}
	}

	target.Weight = m.GetWeight()

	if h, ok := interface{}(m.GetPortalMetadata()).(clone.Cloner); ok {
		target.PortalMetadata = h.Clone().(*PortalMetadata)
	} else {
		target.PortalMetadata = proto.Clone(m.GetPortalMetadata()).(*PortalMetadata)
	}

	return target
}

// Clone function
func (m *HTTPRoute) Clone() proto.Message {
	var target *HTTPRoute
	if m == nil {
		return target
	}
	target = &HTTPRoute{}

	target.Name = m.GetName()

	if m.GetLabels() != nil {
		target.Labels = make(map[string]string, len(m.GetLabels()))
		for k, v := range m.GetLabels() {

			target.Labels[k] = v

		}
	}

	if m.GetMatchers() != nil {
		target.Matchers = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.HTTPRequestMatcher, len(m.GetMatchers()))
		for idx, v := range m.GetMatchers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Matchers[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.HTTPRequestMatcher)
			} else {
				target.Matchers[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.HTTPRequestMatcher)
			}

		}
	}

	switch m.ActionType.(type) {

	case *HTTPRoute_ForwardTo:

		if h, ok := interface{}(m.GetForwardTo()).(clone.Cloner); ok {
			target.ActionType = &HTTPRoute_ForwardTo{
				ForwardTo: h.Clone().(*ForwardToAction),
			}
		} else {
			target.ActionType = &HTTPRoute_ForwardTo{
				ForwardTo: proto.Clone(m.GetForwardTo()).(*ForwardToAction),
			}
		}

	case *HTTPRoute_Delegate:

		if h, ok := interface{}(m.GetDelegate()).(clone.Cloner); ok {
			target.ActionType = &HTTPRoute_Delegate{
				Delegate: h.Clone().(*DelegateAction),
			}
		} else {
			target.ActionType = &HTTPRoute_Delegate{
				Delegate: proto.Clone(m.GetDelegate()).(*DelegateAction),
			}
		}

	case *HTTPRoute_Redirect:

		if h, ok := interface{}(m.GetRedirect()).(clone.Cloner); ok {
			target.ActionType = &HTTPRoute_Redirect{
				Redirect: h.Clone().(*RedirectAction),
			}
		} else {
			target.ActionType = &HTTPRoute_Redirect{
				Redirect: proto.Clone(m.GetRedirect()).(*RedirectAction),
			}
		}

	case *HTTPRoute_DirectResponse:

		if h, ok := interface{}(m.GetDirectResponse()).(clone.Cloner); ok {
			target.ActionType = &HTTPRoute_DirectResponse{
				DirectResponse: h.Clone().(*DirectResponseAction),
			}
		} else {
			target.ActionType = &HTTPRoute_DirectResponse{
				DirectResponse: proto.Clone(m.GetDirectResponse()).(*DirectResponseAction),
			}
		}

	case *HTTPRoute_Graphql:

		if h, ok := interface{}(m.GetGraphql()).(clone.Cloner); ok {
			target.ActionType = &HTTPRoute_Graphql{
				Graphql: h.Clone().(*GraphQLAction),
			}
		} else {
			target.ActionType = &HTTPRoute_Graphql{
				Graphql: proto.Clone(m.GetGraphql()).(*GraphQLAction),
			}
		}

	}

	return target
}

// Clone function
func (m *TCPRoute) Clone() proto.Message {
	var target *TCPRoute
	if m == nil {
		return target
	}
	target = &TCPRoute{}

	if m.GetMatchers() != nil {
		target.Matchers = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.TCPRequestMatcher, len(m.GetMatchers()))
		for idx, v := range m.GetMatchers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Matchers[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.TCPRequestMatcher)
			} else {
				target.Matchers[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.TCPRequestMatcher)
			}

		}
	}

	switch m.ActionType.(type) {

	case *TCPRoute_ForwardTo:

		if h, ok := interface{}(m.GetForwardTo()).(clone.Cloner); ok {
			target.ActionType = &TCPRoute_ForwardTo{
				ForwardTo: h.Clone().(*ForwardToAction),
			}
		} else {
			target.ActionType = &TCPRoute_ForwardTo{
				ForwardTo: proto.Clone(m.GetForwardTo()).(*ForwardToAction),
			}
		}

	}

	return target
}

// Clone function
func (m *TLSRoute) Clone() proto.Message {
	var target *TLSRoute
	if m == nil {
		return target
	}
	target = &TLSRoute{}

	if m.GetMatchers() != nil {
		target.Matchers = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.TLSRequestMatcher, len(m.GetMatchers()))
		for idx, v := range m.GetMatchers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Matchers[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.TLSRequestMatcher)
			} else {
				target.Matchers[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.TLSRequestMatcher)
			}

		}
	}

	switch m.ActionType.(type) {

	case *TLSRoute_ForwardTo:

		if h, ok := interface{}(m.GetForwardTo()).(clone.Cloner); ok {
			target.ActionType = &TLSRoute_ForwardTo{
				ForwardTo: h.Clone().(*TLSRoute_TLSForwardToAction),
			}
		} else {
			target.ActionType = &TLSRoute_ForwardTo{
				ForwardTo: proto.Clone(m.GetForwardTo()).(*TLSRoute_TLSForwardToAction),
			}
		}

	}

	return target
}

// Clone function
func (m *GraphQLAction) Clone() proto.Message {
	var target *GraphQLAction
	if m == nil {
		return target
	}
	target = &GraphQLAction{}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*GraphQLAction_Options)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*GraphQLAction_Options)
	}

	switch m.GraphqlSchema.(type) {

	case *GraphQLAction_Schema:

		if h, ok := interface{}(m.GetSchema()).(clone.Cloner); ok {
			target.GraphqlSchema = &GraphQLAction_Schema{
				Schema: h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ClusterObjectRef),
			}
		} else {
			target.GraphqlSchema = &GraphQLAction_Schema{
				Schema: proto.Clone(m.GetSchema()).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ClusterObjectRef),
			}
		}

	case *GraphQLAction_StitchedSchema:

		if h, ok := interface{}(m.GetStitchedSchema()).(clone.Cloner); ok {
			target.GraphqlSchema = &GraphQLAction_StitchedSchema{
				StitchedSchema: h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ClusterObjectRef),
			}
		} else {
			target.GraphqlSchema = &GraphQLAction_StitchedSchema{
				StitchedSchema: proto.Clone(m.GetStitchedSchema()).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ClusterObjectRef),
			}
		}

	}

	return target
}

// Clone function
func (m *ForwardToAction) Clone() proto.Message {
	var target *ForwardToAction
	if m == nil {
		return target
	}
	target = &ForwardToAction{}

	if m.GetDestinations() != nil {
		target.Destinations = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference, len(m.GetDestinations()))
		for idx, v := range m.GetDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Destinations[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.Destinations[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	target.PathRewrite = m.GetPathRewrite()

	target.HostRewrite = m.GetHostRewrite()

	return target
}

// Clone function
func (m *RedirectAction) Clone() proto.Message {
	var target *RedirectAction
	if m == nil {
		return target
	}
	target = &RedirectAction{}

	target.HostRedirect = m.GetHostRedirect()

	target.ResponseCode = m.GetResponseCode()

	switch m.PathRewriteSpecifier.(type) {

	case *RedirectAction_PathRedirect:

		target.PathRewriteSpecifier = &RedirectAction_PathRedirect{
			PathRedirect: m.GetPathRedirect(),
		}

	}

	return target
}

// Clone function
func (m *DirectResponseAction) Clone() proto.Message {
	var target *DirectResponseAction
	if m == nil {
		return target
	}
	target = &DirectResponseAction{}

	target.Status = m.GetStatus()

	target.Body = m.GetBody()

	return target
}

// Clone function
func (m *DelegateAction) Clone() proto.Message {
	var target *DelegateAction
	if m == nil {
		return target
	}
	target = &DelegateAction{}

	if m.GetRouteTables() != nil {
		target.RouteTables = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector, len(m.GetRouteTables()))
		for idx, v := range m.GetRouteTables() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.RouteTables[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			} else {
				target.RouteTables[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			}

		}
	}

	if m.GetAllowedRoutes() != nil {
		target.AllowedRoutes = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.RouteFilter, len(m.GetAllowedRoutes()))
		for idx, v := range m.GetAllowedRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AllowedRoutes[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.RouteFilter)
			} else {
				target.AllowedRoutes[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.RouteFilter)
			}

		}
	}

	target.SortMethod = m.GetSortMethod()

	return target
}

// Clone function
func (m *PortalMetadata) Clone() proto.Message {
	var target *PortalMetadata
	if m == nil {
		return target
	}
	target = &PortalMetadata{}

	target.Title = m.GetTitle()

	target.Description = m.GetDescription()

	target.TermsOfService = m.GetTermsOfService()

	target.Contact = m.GetContact()

	target.License = m.GetLicense()

	if m.GetCustomMetadata() != nil {
		target.CustomMetadata = make(map[string]string, len(m.GetCustomMetadata()))
		for k, v := range m.GetCustomMetadata() {

			target.CustomMetadata[k] = v

		}
	}

	return target
}

// Clone function
func (m *RouteTableStatus) Clone() proto.Message {
	var target *RouteTableStatus
	if m == nil {
		return target
	}
	target = &RouteTableStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	if m.GetNumAppliedRoutePolicies() != nil {
		target.NumAppliedRoutePolicies = make(map[string]uint32, len(m.GetNumAppliedRoutePolicies()))
		for k, v := range m.GetNumAppliedRoutePolicies() {

			target.NumAppliedRoutePolicies[k] = v

		}
	}

	target.NumParentRouteTables = m.GetNumParentRouteTables()

	target.OwnedByWorkspace = m.GetOwnedByWorkspace()

	target.NumAllowedVirtualGateways = m.GetNumAllowedVirtualGateways()

	return target
}

// Clone function
func (m *RouteTableReport) Clone() proto.Message {
	var target *RouteTableReport
	if m == nil {
		return target
	}
	target = &RouteTableReport{}

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

	if m.GetAppliedRoutePolicies() != nil {
		target.AppliedRoutePolicies = make(map[string]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.AppliedRoutePolicies, len(m.GetAppliedRoutePolicies()))
		for k, v := range m.GetAppliedRoutePolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AppliedRoutePolicies[k] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.AppliedRoutePolicies)
			} else {
				target.AppliedRoutePolicies[k] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.AppliedRoutePolicies)
			}

		}
	}

	if m.GetParentRouteTables() != nil {
		target.ParentRouteTables = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference, len(m.GetParentRouteTables()))
		for idx, v := range m.GetParentRouteTables() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ParentRouteTables[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.ParentRouteTables[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	target.OwnerWorkspace = m.GetOwnerWorkspace()

	if m.GetAllowedVirtualGateways() != nil {
		target.AllowedVirtualGateways = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference, len(m.GetAllowedVirtualGateways()))
		for idx, v := range m.GetAllowedVirtualGateways() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AllowedVirtualGateways[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.AllowedVirtualGateways[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	if m.GetDelegatedToRouteTables() != nil {
		target.DelegatedToRouteTables = make([]*RouteTableReport_DelegatedRouteTableReference, len(m.GetDelegatedToRouteTables()))
		for idx, v := range m.GetDelegatedToRouteTables() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.DelegatedToRouteTables[idx] = h.Clone().(*RouteTableReport_DelegatedRouteTableReference)
			} else {
				target.DelegatedToRouteTables[idx] = proto.Clone(v).(*RouteTableReport_DelegatedRouteTableReference)
			}

		}
	}

	return target
}

// Clone function
func (m *TLSRoute_TLSForwardToAction) Clone() proto.Message {
	var target *TLSRoute_TLSForwardToAction
	if m == nil {
		return target
	}
	target = &TLSRoute_TLSForwardToAction{}

	if m.GetDestinations() != nil {
		target.Destinations = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference, len(m.GetDestinations()))
		for idx, v := range m.GetDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Destinations[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.Destinations[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	return target
}

// Clone function
func (m *GraphQLAction_Options) Clone() proto.Message {
	var target *GraphQLAction_Options
	if m == nil {
		return target
	}
	target = &GraphQLAction_Options{}

	if h, ok := interface{}(m.GetLogSensitiveInfo()).(clone.Cloner); ok {
		target.LogSensitiveInfo = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.LogSensitiveInfo = proto.Clone(m.GetLogSensitiveInfo()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *RouteTableReport_DelegatedRouteTableReference) Clone() proto.Message {
	var target *RouteTableReport_DelegatedRouteTableReference
	if m == nil {
		return target
	}
	target = &RouteTableReport_DelegatedRouteTableReference{}

	target.RouteIndex = m.GetRouteIndex()

	if h, ok := interface{}(m.GetRouteTable()).(clone.Cloner); ok {
		target.RouteTable = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.RouteTable = proto.Clone(m.GetRouteTable()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	}

	return target
}
