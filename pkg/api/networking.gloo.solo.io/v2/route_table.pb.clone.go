// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/networking/v2/route_table.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2"
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
		target.VirtualGateways = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference, len(m.GetVirtualGateways()))
		for idx, v := range m.GetVirtualGateways() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.VirtualGateways[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.VirtualGateways[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	if m.GetWorkloadSelectors() != nil {
		target.WorkloadSelectors = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkloadSelector, len(m.GetWorkloadSelectors()))
		for idx, v := range m.GetWorkloadSelectors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.WorkloadSelectors[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkloadSelector)
			} else {
				target.WorkloadSelectors[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkloadSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetDefaultDestination()).(clone.Cloner); ok {
		target.DefaultDestination = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
	} else {
		target.DefaultDestination = proto.Clone(m.GetDefaultDestination()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
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

	target.Weight = m.GetWeight()

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
		target.Matchers = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.HTTPRequestMatcher, len(m.GetMatchers()))
		for idx, v := range m.GetMatchers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Matchers[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.HTTPRequestMatcher)
			} else {
				target.Matchers[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.HTTPRequestMatcher)
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
		target.Destinations = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference, len(m.GetDestinations()))
		for idx, v := range m.GetDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Destinations[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.Destinations[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	target.PathRewrite = m.GetPathRewrite()

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
		target.RouteTables = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectSelector, len(m.GetRouteTables()))
		for idx, v := range m.GetRouteTables() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.RouteTables[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectSelector)
			} else {
				target.RouteTables[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectSelector)
			}

		}
	}

	target.SortMethod = m.GetSortMethod()

	return target
}

// Clone function
func (m *RouteTableStatus) Clone() proto.Message {
	var target *RouteTableStatus
	if m == nil {
		return target
	}
	target = &RouteTableStatus{}

	if h, ok := interface{}(m.GetGlobal()).(clone.Cloner); ok {
		target.Global = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericGlobalStatus)
	} else {
		target.Global = proto.Clone(m.GetGlobal()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericGlobalStatus)
	}

	if m.GetWorkspaces() != nil {
		target.Workspaces = make(map[string]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus, len(m.GetWorkspaces()))
		for k, v := range m.GetWorkspaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workspaces[k] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus)
			} else {
				target.Workspaces[k] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus)
			}

		}
	}

	if m.GetAppliedRoutePolicies() != nil {
		target.AppliedRoutePolicies = make(map[string]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.AppliedRoutePolicies, len(m.GetAppliedRoutePolicies()))
		for k, v := range m.GetAppliedRoutePolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AppliedRoutePolicies[k] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.AppliedRoutePolicies)
			} else {
				target.AppliedRoutePolicies[k] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.AppliedRoutePolicies)
			}

		}
	}

	if m.GetParentRouteTables() != nil {
		target.ParentRouteTables = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference, len(m.GetParentRouteTables()))
		for idx, v := range m.GetParentRouteTables() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ParentRouteTables[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.ParentRouteTables[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	if h, ok := interface{}(m.GetOwnerWorkspace()).(clone.Cloner); ok {
		target.OwnerWorkspace = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.OwnerWorkspace)
	} else {
		target.OwnerWorkspace = proto.Clone(m.GetOwnerWorkspace()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.OwnerWorkspace)
	}

	if m.GetAllowedVirtualGateways() != nil {
		target.AllowedVirtualGateways = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference, len(m.GetAllowedVirtualGateways()))
		for idx, v := range m.GetAllowedVirtualGateways() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AllowedVirtualGateways[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.AllowedVirtualGateways[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	return target
}
