// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/trafficcontrol/transformation_policy.proto

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

	github_com_solo_io_solo_apis_pkg_api_envoy_gloo_api_envoy_config_filter_http_transformation_v2 "github.com/solo-io/solo-apis/pkg/api/envoy-gloo/api/envoy/config/filter/http/transformation/v2"
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
func (m *TransformationPolicySpec) Clone() proto.Message {
	var target *TransformationPolicySpec
	if m == nil {
		return target
	}
	target = &TransformationPolicySpec{}

	if m.GetApplyToRoutes() != nil {
		target.ApplyToRoutes = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteSelector, len(m.GetApplyToRoutes()))
		for idx, v := range m.GetApplyToRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ApplyToRoutes[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteSelector)
			} else {
				target.ApplyToRoutes[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*TransformationPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*TransformationPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *TransformationPolicyStatus) Clone() proto.Message {
	var target *TransformationPolicyStatus
	if m == nil {
		return target
	}
	target = &TransformationPolicyStatus{}

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

	if m.GetSelectedRoutes() != nil {
		target.SelectedRoutes = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteReference, len(m.GetSelectedRoutes()))
		for idx, v := range m.GetSelectedRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedRoutes[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteReference)
			} else {
				target.SelectedRoutes[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.RouteReference)
			}

		}
	}

	return target
}

// Clone function
func (m *TransformationPolicySpec_Config) Clone() proto.Message {
	var target *TransformationPolicySpec_Config
	if m == nil {
		return target
	}
	target = &TransformationPolicySpec_Config{}

	if h, ok := interface{}(m.GetPhase()).(clone.Cloner); ok {
		target.Phase = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PrioritizedPhase)
	} else {
		target.Phase = proto.Clone(m.GetPhase()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.PrioritizedPhase)
	}

	if h, ok := interface{}(m.GetRequest()).(clone.Cloner); ok {
		target.Request = h.Clone().(*TransformationPolicySpec_Config_RequestTransformation)
	} else {
		target.Request = proto.Clone(m.GetRequest()).(*TransformationPolicySpec_Config_RequestTransformation)
	}

	if h, ok := interface{}(m.GetResponse()).(clone.Cloner); ok {
		target.Response = h.Clone().(*TransformationPolicySpec_Config_ResponseTransformation)
	} else {
		target.Response = proto.Clone(m.GetResponse()).(*TransformationPolicySpec_Config_ResponseTransformation)
	}

	return target
}

// Clone function
func (m *TransformationPolicySpec_Config_RequestTransformation) Clone() proto.Message {
	var target *TransformationPolicySpec_Config_RequestTransformation
	if m == nil {
		return target
	}
	target = &TransformationPolicySpec_Config_RequestTransformation{}

	target.RecalculateRoutingDestination = m.GetRecalculateRoutingDestination()

	if h, ok := interface{}(m.GetInjaTemplate()).(clone.Cloner); ok {
		target.InjaTemplate = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_envoy_gloo_api_envoy_config_filter_http_transformation_v2.TransformationTemplate)
	} else {
		target.InjaTemplate = proto.Clone(m.GetInjaTemplate()).(*github_com_solo_io_solo_apis_pkg_api_envoy_gloo_api_envoy_config_filter_http_transformation_v2.TransformationTemplate)
	}

	return target
}

// Clone function
func (m *TransformationPolicySpec_Config_ResponseTransformation) Clone() proto.Message {
	var target *TransformationPolicySpec_Config_ResponseTransformation
	if m == nil {
		return target
	}
	target = &TransformationPolicySpec_Config_ResponseTransformation{}

	if h, ok := interface{}(m.GetInjaTemplate()).(clone.Cloner); ok {
		target.InjaTemplate = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_envoy_gloo_api_envoy_config_filter_http_transformation_v2.TransformationTemplate)
	} else {
		target.InjaTemplate = proto.Clone(m.GetInjaTemplate()).(*github_com_solo_io_solo_apis_pkg_api_envoy_gloo_api_envoy_config_filter_http_transformation_v2.TransformationTemplate)
	}

	return target
}
