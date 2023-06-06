// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/security/ext_auth_policy.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-apis/client-go/enterprise.gloo.solo.io/v1"
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
func (m *ExtAuthPolicySpec) Clone() proto.Message {
	var target *ExtAuthPolicySpec
	if m == nil {
		return target
	}
	target = &ExtAuthPolicySpec{}

	if m.GetApplyToRoutes() != nil {
		target.ApplyToRoutes = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.RouteSelector, len(m.GetApplyToRoutes()))
		for idx, v := range m.GetApplyToRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ApplyToRoutes[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.RouteSelector)
			} else {
				target.ApplyToRoutes[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.RouteSelector)
			}

		}
	}

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
		target.Config = h.Clone().(*ExtAuthPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*ExtAuthPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *ExtAuthPolicyStatus) Clone() proto.Message {
	var target *ExtAuthPolicyStatus
	if m == nil {
		return target
	}
	target = &ExtAuthPolicyStatus{}

	if h, ok := interface{}(m.GetGlobal()).(clone.Cloner); ok {
		target.Global = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.GenericGlobalStatus)
	} else {
		target.Global = proto.Clone(m.GetGlobal()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.GenericGlobalStatus)
	}

	if m.GetWorkspaces() != nil {
		target.Workspaces = make(map[string]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.WorkspaceStatus, len(m.GetWorkspaces()))
		for k, v := range m.GetWorkspaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workspaces[k] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.WorkspaceStatus)
			} else {
				target.Workspaces[k] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.WorkspaceStatus)
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

	if m.GetSelectedRoutes() != nil {
		target.SelectedRoutes = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.RouteReference, len(m.GetSelectedRoutes()))
		for idx, v := range m.GetSelectedRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedRoutes[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.RouteReference)
			} else {
				target.SelectedRoutes[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.RouteReference)
			}

		}
	}

	return target
}

// Clone function
func (m *ExtAuthPolicySpec_Config) Clone() proto.Message {
	var target *ExtAuthPolicySpec_Config
	if m == nil {
		return target
	}
	target = &ExtAuthPolicySpec_Config{}

	if h, ok := interface{}(m.GetServer()).(clone.Cloner); ok {
		target.Server = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.Server = proto.Clone(m.GetServer()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	}

	switch m.AuthType.(type) {

	case *ExtAuthPolicySpec_Config_Disable:

		target.AuthType = &ExtAuthPolicySpec_Config_Disable{
			Disable: m.GetDisable(),
		}

	case *ExtAuthPolicySpec_Config_GlooAuth:

		if h, ok := interface{}(m.GetGlooAuth()).(clone.Cloner); ok {
			target.AuthType = &ExtAuthPolicySpec_Config_GlooAuth{
				GlooAuth: h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_enterprise_gloo_solo_io_v1.AuthConfigSpec),
			}
		} else {
			target.AuthType = &ExtAuthPolicySpec_Config_GlooAuth{
				GlooAuth: proto.Clone(m.GetGlooAuth()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_enterprise_gloo_solo_io_v1.AuthConfigSpec),
			}
		}

	case *ExtAuthPolicySpec_Config_CustomAuth_:

		if h, ok := interface{}(m.GetCustomAuth()).(clone.Cloner); ok {
			target.AuthType = &ExtAuthPolicySpec_Config_CustomAuth_{
				CustomAuth: h.Clone().(*ExtAuthPolicySpec_Config_CustomAuth),
			}
		} else {
			target.AuthType = &ExtAuthPolicySpec_Config_CustomAuth_{
				CustomAuth: proto.Clone(m.GetCustomAuth()).(*ExtAuthPolicySpec_Config_CustomAuth),
			}
		}

	}

	return target
}

// Clone function
func (m *ExtAuthPolicySpec_Config_CustomAuth) Clone() proto.Message {
	var target *ExtAuthPolicySpec_Config_CustomAuth
	if m == nil {
		return target
	}
	target = &ExtAuthPolicySpec_Config_CustomAuth{}

	if m.GetContextExtensions() != nil {
		target.ContextExtensions = make(map[string]string, len(m.GetContextExtensions()))
		for k, v := range m.GetContextExtensions() {

			target.ContextExtensions[k] = v

		}
	}

	return target
}
