// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2alpha1/waypoint_lifecycle_manager.proto

package v2alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

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
func (m *WaypointLifecycleManagerSpec) Clone() proto.Message {
	var target *WaypointLifecycleManagerSpec
	if m == nil {
		return target
	}
	target = &WaypointLifecycleManagerSpec{}

	if m.GetApplyToServiceAccount() != nil {
		target.ApplyToServiceAccount = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector, len(m.GetApplyToServiceAccount()))
		for idx, v := range m.GetApplyToServiceAccount() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ApplyToServiceAccount[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			} else {
				target.ApplyToServiceAccount[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetDeploymentSpec()).(clone.Cloner); ok {
		target.DeploymentSpec = h.Clone().(*github_com_golang_protobuf_ptypes_struct.Struct)
	} else {
		target.DeploymentSpec = proto.Clone(m.GetDeploymentSpec()).(*github_com_golang_protobuf_ptypes_struct.Struct)
	}

	return target
}

// Clone function
func (m *WaypointLifecycleManagerStatus) Clone() proto.Message {
	var target *WaypointLifecycleManagerStatus
	if m == nil {
		return target
	}
	target = &WaypointLifecycleManagerStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	return target
}

// Clone function
func (m *WaypointLifecycleManagerReport) Clone() proto.Message {
	var target *WaypointLifecycleManagerReport
	if m == nil {
		return target
	}
	target = &WaypointLifecycleManagerReport{}

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

	return target
}