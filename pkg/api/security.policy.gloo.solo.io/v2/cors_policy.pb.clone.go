// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/security/cors_policy.proto

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
func (m *CORSPolicySpec) Clone() proto.Message {
	var target *CORSPolicySpec
	if m == nil {
		return target
	}
	target = &CORSPolicySpec{}

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
		target.Config = h.Clone().(*CORSPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*CORSPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *CORSPolicyStatus) Clone() proto.Message {
	var target *CORSPolicyStatus
	if m == nil {
		return target
	}
	target = &CORSPolicyStatus{}

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
func (m *CORSPolicySpec_Config) Clone() proto.Message {
	var target *CORSPolicySpec_Config
	if m == nil {
		return target
	}
	target = &CORSPolicySpec_Config{}

	if m.GetAllowOrigins() != nil {
		target.AllowOrigins = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.StringMatch, len(m.GetAllowOrigins()))
		for idx, v := range m.GetAllowOrigins() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AllowOrigins[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.StringMatch)
			} else {
				target.AllowOrigins[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.StringMatch)
			}

		}
	}

	if m.GetAllowMethods() != nil {
		target.AllowMethods = make([]string, len(m.GetAllowMethods()))
		for idx, v := range m.GetAllowMethods() {

			target.AllowMethods[idx] = v

		}
	}

	if m.GetAllowHeaders() != nil {
		target.AllowHeaders = make([]string, len(m.GetAllowHeaders()))
		for idx, v := range m.GetAllowHeaders() {

			target.AllowHeaders[idx] = v

		}
	}

	if m.GetExposeHeaders() != nil {
		target.ExposeHeaders = make([]string, len(m.GetExposeHeaders()))
		for idx, v := range m.GetExposeHeaders() {

			target.ExposeHeaders[idx] = v

		}
	}

	if h, ok := interface{}(m.GetMaxAge()).(clone.Cloner); ok {
		target.MaxAge = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.MaxAge = proto.Clone(m.GetMaxAge()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetAllowCredentials()).(clone.Cloner); ok {
		target.AllowCredentials = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.AllowCredentials = proto.Clone(m.GetAllowCredentials()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}
