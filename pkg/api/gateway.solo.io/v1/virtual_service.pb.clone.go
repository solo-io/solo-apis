// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/virtual_service.proto

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

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/core/matchers"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_ssl "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/ssl"

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
func (m *VirtualServiceSpec) Clone() proto.Message {
	var target *VirtualServiceSpec
	if m == nil {
		return target
	}
	target = &VirtualServiceSpec{}

	if h, ok := interface{}(m.GetVirtualHost()).(clone.Cloner); ok {
		target.VirtualHost = h.Clone().(*VirtualHost)
	} else {
		target.VirtualHost = proto.Clone(m.GetVirtualHost()).(*VirtualHost)
	}

	if h, ok := interface{}(m.GetSslConfig()).(clone.Cloner); ok {
		target.SslConfig = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_ssl.SslConfig)
	} else {
		target.SslConfig = proto.Clone(m.GetSslConfig()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_ssl.SslConfig)
	}

	target.DisplayName = m.GetDisplayName()

	return target
}

// Clone function
func (m *VirtualHost) Clone() proto.Message {
	var target *VirtualHost
	if m == nil {
		return target
	}
	target = &VirtualHost{}

	if m.GetDomains() != nil {
		target.Domains = make([]string, len(m.GetDomains()))
		for idx, v := range m.GetDomains() {

			target.Domains[idx] = v

		}
	}

	if m.GetRoutes() != nil {
		target.Routes = make([]*Route, len(m.GetRoutes()))
		for idx, v := range m.GetRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Routes[idx] = h.Clone().(*Route)
			} else {
				target.Routes[idx] = proto.Clone(v).(*Route)
			}

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.VirtualHostOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.VirtualHostOptions)
	}

	switch m.ExternalOptionsConfig.(type) {

	case *VirtualHost_OptionsConfigRefs:

		if h, ok := interface{}(m.GetOptionsConfigRefs()).(clone.Cloner); ok {
			target.ExternalOptionsConfig = &VirtualHost_OptionsConfigRefs{
				OptionsConfigRefs: h.Clone().(*DelegateOptionsRefs),
			}
		} else {
			target.ExternalOptionsConfig = &VirtualHost_OptionsConfigRefs{
				OptionsConfigRefs: proto.Clone(m.GetOptionsConfigRefs()).(*DelegateOptionsRefs),
			}
		}

	}

	return target
}

// Clone function
func (m *Route) Clone() proto.Message {
	var target *Route
	if m == nil {
		return target
	}
	target = &Route{}

	if m.GetMatchers() != nil {
		target.Matchers = make([]*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers.Matcher, len(m.GetMatchers()))
		for idx, v := range m.GetMatchers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Matchers[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers.Matcher)
			} else {
				target.Matchers[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers.Matcher)
			}

		}
	}

	if h, ok := interface{}(m.GetInheritableMatchers()).(clone.Cloner); ok {
		target.InheritableMatchers = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.InheritableMatchers = proto.Clone(m.GetInheritableMatchers()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetInheritablePathMatchers()).(clone.Cloner); ok {
		target.InheritablePathMatchers = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.InheritablePathMatchers = proto.Clone(m.GetInheritablePathMatchers()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RouteOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RouteOptions)
	}

	target.Name = m.GetName()

	switch m.Action.(type) {

	case *Route_RouteAction:

		if h, ok := interface{}(m.GetRouteAction()).(clone.Cloner); ok {
			target.Action = &Route_RouteAction{
				RouteAction: h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RouteAction),
			}
		} else {
			target.Action = &Route_RouteAction{
				RouteAction: proto.Clone(m.GetRouteAction()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RouteAction),
			}
		}

	case *Route_RedirectAction:

		if h, ok := interface{}(m.GetRedirectAction()).(clone.Cloner); ok {
			target.Action = &Route_RedirectAction{
				RedirectAction: h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RedirectAction),
			}
		} else {
			target.Action = &Route_RedirectAction{
				RedirectAction: proto.Clone(m.GetRedirectAction()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.RedirectAction),
			}
		}

	case *Route_DirectResponseAction:

		if h, ok := interface{}(m.GetDirectResponseAction()).(clone.Cloner); ok {
			target.Action = &Route_DirectResponseAction{
				DirectResponseAction: h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.DirectResponseAction),
			}
		} else {
			target.Action = &Route_DirectResponseAction{
				DirectResponseAction: proto.Clone(m.GetDirectResponseAction()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1.DirectResponseAction),
			}
		}

	case *Route_DelegateAction:

		if h, ok := interface{}(m.GetDelegateAction()).(clone.Cloner); ok {
			target.Action = &Route_DelegateAction{
				DelegateAction: h.Clone().(*DelegateAction),
			}
		} else {
			target.Action = &Route_DelegateAction{
				DelegateAction: proto.Clone(m.GetDelegateAction()).(*DelegateAction),
			}
		}

	case *Route_GraphqlApiRef:

		if h, ok := interface{}(m.GetGraphqlApiRef()).(clone.Cloner); ok {
			target.Action = &Route_GraphqlApiRef{
				GraphqlApiRef: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.Action = &Route_GraphqlApiRef{
				GraphqlApiRef: proto.Clone(m.GetGraphqlApiRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	}

	switch m.ExternalOptionsConfig.(type) {

	case *Route_OptionsConfigRefs:

		if h, ok := interface{}(m.GetOptionsConfigRefs()).(clone.Cloner); ok {
			target.ExternalOptionsConfig = &Route_OptionsConfigRefs{
				OptionsConfigRefs: h.Clone().(*DelegateOptionsRefs),
			}
		} else {
			target.ExternalOptionsConfig = &Route_OptionsConfigRefs{
				OptionsConfigRefs: proto.Clone(m.GetOptionsConfigRefs()).(*DelegateOptionsRefs),
			}
		}

	}

	return target
}

// Clone function
func (m *DelegateOptionsRefs) Clone() proto.Message {
	var target *DelegateOptionsRefs
	if m == nil {
		return target
	}
	target = &DelegateOptionsRefs{}

	if m.GetDelegateOptions() != nil {
		target.DelegateOptions = make([]*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef, len(m.GetDelegateOptions()))
		for idx, v := range m.GetDelegateOptions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.DelegateOptions[idx] = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			} else {
				target.DelegateOptions[idx] = proto.Clone(v).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			}

		}
	}

	return target
}

// Clone function
func (m *DelegateAction) Clone() proto.Message {
	var target *DelegateAction
	if m == nil {
		return target
	}
	target = &DelegateAction{}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	switch m.DelegationType.(type) {

	case *DelegateAction_Ref:

		if h, ok := interface{}(m.GetRef()).(clone.Cloner); ok {
			target.DelegationType = &DelegateAction_Ref{
				Ref: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.DelegationType = &DelegateAction_Ref{
				Ref: proto.Clone(m.GetRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *DelegateAction_Selector:

		if h, ok := interface{}(m.GetSelector()).(clone.Cloner); ok {
			target.DelegationType = &DelegateAction_Selector{
				Selector: h.Clone().(*RouteTableSelector),
			}
		} else {
			target.DelegationType = &DelegateAction_Selector{
				Selector: proto.Clone(m.GetSelector()).(*RouteTableSelector),
			}
		}

	}

	return target
}

// Clone function
func (m *RouteTableSelector) Clone() proto.Message {
	var target *RouteTableSelector
	if m == nil {
		return target
	}
	target = &RouteTableSelector{}

	if m.GetNamespaces() != nil {
		target.Namespaces = make([]string, len(m.GetNamespaces()))
		for idx, v := range m.GetNamespaces() {

			target.Namespaces[idx] = v

		}
	}

	if m.GetLabels() != nil {
		target.Labels = make(map[string]string, len(m.GetLabels()))
		for k, v := range m.GetLabels() {

			target.Labels[k] = v

		}
	}

	if m.GetExpressions() != nil {
		target.Expressions = make([]*RouteTableSelector_Expression, len(m.GetExpressions()))
		for idx, v := range m.GetExpressions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Expressions[idx] = h.Clone().(*RouteTableSelector_Expression)
			} else {
				target.Expressions[idx] = proto.Clone(v).(*RouteTableSelector_Expression)
			}

		}
	}

	return target
}

// Clone function
func (m *VirtualServiceStatus) Clone() proto.Message {
	var target *VirtualServiceStatus
	if m == nil {
		return target
	}
	target = &VirtualServiceStatus{}

	target.State = m.GetState()

	target.Reason = m.GetReason()

	target.ReportedBy = m.GetReportedBy()

	if m.GetSubresourceStatuses() != nil {
		target.SubresourceStatuses = make(map[string]*VirtualServiceStatus, len(m.GetSubresourceStatuses()))
		for k, v := range m.GetSubresourceStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SubresourceStatuses[k] = h.Clone().(*VirtualServiceStatus)
			} else {
				target.SubresourceStatuses[k] = proto.Clone(v).(*VirtualServiceStatus)
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
func (m *VirtualServiceNamespacedStatuses) Clone() proto.Message {
	var target *VirtualServiceNamespacedStatuses
	if m == nil {
		return target
	}
	target = &VirtualServiceNamespacedStatuses{}

	if m.GetStatuses() != nil {
		target.Statuses = make(map[string]*VirtualServiceStatus, len(m.GetStatuses()))
		for k, v := range m.GetStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Statuses[k] = h.Clone().(*VirtualServiceStatus)
			} else {
				target.Statuses[k] = proto.Clone(v).(*VirtualServiceStatus)
			}

		}
	}

	return target
}

// Clone function
func (m *RouteTableSelector_Expression) Clone() proto.Message {
	var target *RouteTableSelector_Expression
	if m == nil {
		return target
	}
	target = &RouteTableSelector_Expression{}

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
