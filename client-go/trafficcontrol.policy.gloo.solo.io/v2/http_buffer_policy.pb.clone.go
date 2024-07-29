// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/trafficcontrol/http_buffer_policy.proto

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

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *HTTPBufferPolicySpec) Clone() proto.Message {
	var target *HTTPBufferPolicySpec
	if m == nil {
		return target
	}
	target = &HTTPBufferPolicySpec{}

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
		target.Config = h.Clone().(*HTTPBufferPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*HTTPBufferPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *HTTPBufferPolicyStatus) Clone() proto.Message {
	var target *HTTPBufferPolicyStatus
	if m == nil {
		return target
	}
	target = &HTTPBufferPolicyStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	target.NumSelectedRoutes = m.GetNumSelectedRoutes()

	return target
}

// Clone function
func (m *HTTPBufferPolicyReport) Clone() proto.Message {
	var target *HTTPBufferPolicyReport
	if m == nil {
		return target
	}
	target = &HTTPBufferPolicyReport{}

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
func (m *HTTPBufferPolicySpec_Config) Clone() proto.Message {
	var target *HTTPBufferPolicySpec_Config
	if m == nil {
		return target
	}
	target = &HTTPBufferPolicySpec_Config{}

	if h, ok := interface{}(m.GetMaxRequestBytes()).(clone.Cloner); ok {
		target.MaxRequestBytes = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxRequestBytes = proto.Clone(m.GetMaxRequestBytes()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}