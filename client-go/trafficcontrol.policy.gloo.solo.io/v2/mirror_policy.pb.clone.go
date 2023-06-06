// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/trafficcontrol/mirror_policy.proto

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
func (m *MirrorPolicySpec) Clone() proto.Message {
	var target *MirrorPolicySpec
	if m == nil {
		return target
	}
	target = &MirrorPolicySpec{}

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

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*MirrorPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*MirrorPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *MirrorPolicyStatus) Clone() proto.Message {
	var target *MirrorPolicyStatus
	if m == nil {
		return target
	}
	target = &MirrorPolicyStatus{}

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
func (m *MirrorPolicySpec_Config) Clone() proto.Message {
	var target *MirrorPolicySpec_Config
	if m == nil {
		return target
	}
	target = &MirrorPolicySpec_Config{}

	if h, ok := interface{}(m.GetDestination()).(clone.Cloner); ok {
		target.Destination = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
	} else {
		target.Destination = proto.Clone(m.GetDestination()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
	}

	if h, ok := interface{}(m.GetPercentage()).(clone.Cloner); ok {
		target.Percentage = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.DoubleValue)
	} else {
		target.Percentage = proto.Clone(m.GetPercentage()).(*github_com_golang_protobuf_ptypes_wrappers.DoubleValue)
	}

	return target
}
