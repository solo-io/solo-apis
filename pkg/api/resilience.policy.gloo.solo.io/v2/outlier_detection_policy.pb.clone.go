// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/resilience/outlier_detection_policy.proto

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
func (m *OutlierDetectionPolicySpec) Clone() proto.Message {
	var target *OutlierDetectionPolicySpec
	if m == nil {
		return target
	}
	target = &OutlierDetectionPolicySpec{}

	if m.GetApplyToDestinations() != nil {
		target.ApplyToDestinations = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationSelector, len(m.GetApplyToDestinations()))
		for idx, v := range m.GetApplyToDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ApplyToDestinations[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationSelector)
			} else {
				target.ApplyToDestinations[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*OutlierDetectionPolicySpec_Config)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*OutlierDetectionPolicySpec_Config)
	}

	return target
}

// Clone function
func (m *OutlierDetectionPolicyStatus) Clone() proto.Message {
	var target *OutlierDetectionPolicyStatus
	if m == nil {
		return target
	}
	target = &OutlierDetectionPolicyStatus{}

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

	if m.GetSelectedDestinationPorts() != nil {
		target.SelectedDestinationPorts = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference, len(m.GetSelectedDestinationPorts()))
		for idx, v := range m.GetSelectedDestinationPorts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedDestinationPorts[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.SelectedDestinationPorts[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	return target
}

// Clone function
func (m *OutlierDetectionPolicySpec_Config) Clone() proto.Message {
	var target *OutlierDetectionPolicySpec_Config
	if m == nil {
		return target
	}
	target = &OutlierDetectionPolicySpec_Config{}

	target.ConsecutiveErrors = m.GetConsecutiveErrors()

	if h, ok := interface{}(m.GetInterval()).(clone.Cloner); ok {
		target.Interval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.Interval = proto.Clone(m.GetInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetBaseEjectionTime()).(clone.Cloner); ok {
		target.BaseEjectionTime = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.BaseEjectionTime = proto.Clone(m.GetBaseEjectionTime()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	target.MaxEjectionPercent = m.GetMaxEjectionPercent()

	return target
}
