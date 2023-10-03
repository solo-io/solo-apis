// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/resilience/retry_timeout_policy.proto

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
func (m *RetryTimeoutPolicySpec) Clone() proto.Message {
	var target *RetryTimeoutPolicySpec
	if m == nil {
		return target
	}
	target = &RetryTimeoutPolicySpec{}

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
		target.Config = h.Clone().(*RetryTimeoutPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*RetryTimeoutPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *RetryTimeoutPolicyStatus) Clone() proto.Message {
	var target *RetryTimeoutPolicyStatus
	if m == nil {
		return target
	}
	target = &RetryTimeoutPolicyStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	target.NumSelectedRoutes = m.GetNumSelectedRoutes()

	return target
}

// Clone function
func (m *RetryTimeoutPolicyReport) Clone() proto.Message {
	var target *RetryTimeoutPolicyReport
	if m == nil {
		return target
	}
	target = &RetryTimeoutPolicyReport{}

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
func (m *RetryTimeoutPolicySpec_Config) Clone() proto.Message {
	var target *RetryTimeoutPolicySpec_Config
	if m == nil {
		return target
	}
	target = &RetryTimeoutPolicySpec_Config{}

	if h, ok := interface{}(m.GetRetries()).(clone.Cloner); ok {
		target.Retries = h.Clone().(*RetryTimeoutPolicySpec_Config_RetryPolicy)
	} else {
		target.Retries = proto.Clone(m.GetRetries()).(*RetryTimeoutPolicySpec_Config_RetryPolicy)
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(clone.Cloner); ok {
		target.RequestTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.RequestTimeout = proto.Clone(m.GetRequestTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}

// Clone function
func (m *RetryTimeoutPolicySpec_Config_RetryPolicy) Clone() proto.Message {
	var target *RetryTimeoutPolicySpec_Config_RetryPolicy
	if m == nil {
		return target
	}
	target = &RetryTimeoutPolicySpec_Config_RetryPolicy{}

	if h, ok := interface{}(m.GetAttempts()).(clone.Cloner); ok {
		target.Attempts = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	} else {
		target.Attempts = proto.Clone(m.GetAttempts()).(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	}

	if h, ok := interface{}(m.GetPerTryTimeout()).(clone.Cloner); ok {
		target.PerTryTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.PerTryTimeout = proto.Clone(m.GetPerTryTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	target.RetryOn = m.GetRetryOn()

	if h, ok := interface{}(m.GetRetryRemoteLocalities()).(clone.Cloner); ok {
		target.RetryRemoteLocalities = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.RetryRemoteLocalities = proto.Clone(m.GetRetryRemoteLocalities()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}
