// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/gateway_lifecycle_manager.proto

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
func (m *GatewayLifecycleManagerSpec) Clone() proto.Message {
	var target *GatewayLifecycleManagerSpec
	if m == nil {
		return target
	}
	target = &GatewayLifecycleManagerSpec{}

	if m.GetInstallations() != nil {
		target.Installations = make([]*GatewayInstallation, len(m.GetInstallations()))
		for idx, v := range m.GetInstallations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Installations[idx] = h.Clone().(*GatewayInstallation)
			} else {
				target.Installations[idx] = proto.Clone(v).(*GatewayInstallation)
			}

		}
	}

	return target
}

// Clone function
func (m *GatewayClusterSelector) Clone() proto.Message {
	var target *GatewayClusterSelector
	if m == nil {
		return target
	}
	target = &GatewayClusterSelector{}

	target.Name = m.GetName()

	target.ActiveGateway = m.GetActiveGateway()

	target.TrustDomain = m.GetTrustDomain()

	return target
}

// Clone function
func (m *GatewayInstallation) Clone() proto.Message {
	var target *GatewayInstallation
	if m == nil {
		return target
	}
	target = &GatewayInstallation{}

	target.ControlPlaneRevision = m.GetControlPlaneRevision()

	target.GatewayRevision = m.GetGatewayRevision()

	if m.GetClusters() != nil {
		target.Clusters = make([]*GatewayClusterSelector, len(m.GetClusters()))
		for idx, v := range m.GetClusters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Clusters[idx] = h.Clone().(*GatewayClusterSelector)
			} else {
				target.Clusters[idx] = proto.Clone(v).(*GatewayClusterSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetIstioOperatorSpec()).(clone.Cloner); ok {
		target.IstioOperatorSpec = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.IstioOperatorSpec)
	} else {
		target.IstioOperatorSpec = proto.Clone(m.GetIstioOperatorSpec()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.IstioOperatorSpec)
	}

	target.SkipUpgradeValidation = m.GetSkipUpgradeValidation()

	return target
}

// Clone function
func (m *GatewayLifecycleManagerStatus) Clone() proto.Message {
	var target *GatewayLifecycleManagerStatus
	if m == nil {
		return target
	}
	target = &GatewayLifecycleManagerStatus{}

	if m.GetClusters() != nil {
		target.Clusters = make(map[string]*GatewayLifecycleManagerStatus_ClusterStatuses, len(m.GetClusters()))
		for k, v := range m.GetClusters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Clusters[k] = h.Clone().(*GatewayLifecycleManagerStatus_ClusterStatuses)
			} else {
				target.Clusters[k] = proto.Clone(v).(*GatewayLifecycleManagerStatus_ClusterStatuses)
			}

		}
	}

	return target
}

// Clone function
func (m *GatewayLifecycleManagerNewStatus) Clone() proto.Message {
	var target *GatewayLifecycleManagerNewStatus
	if m == nil {
		return target
	}
	target = &GatewayLifecycleManagerNewStatus{}

	return target
}

// Clone function
func (m *GatewayLifecycleManagerReport) Clone() proto.Message {
	var target *GatewayLifecycleManagerReport
	if m == nil {
		return target
	}
	target = &GatewayLifecycleManagerReport{}

	return target
}

// Clone function
func (m *GatewayLifecycleManagerStatus_ClusterStatuses) Clone() proto.Message {
	var target *GatewayLifecycleManagerStatus_ClusterStatuses
	if m == nil {
		return target
	}
	target = &GatewayLifecycleManagerStatus_ClusterStatuses{}

	if m.GetInstallations() != nil {
		target.Installations = make(map[string]*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus, len(m.GetInstallations()))
		for k, v := range m.GetInstallations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Installations[k] = h.Clone().(*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus)
			} else {
				target.Installations[k] = proto.Clone(v).(*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus)
			}

		}
	}

	return target
}

// Clone function
func (m *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) Clone() proto.Message {
	var target *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus
	if m == nil {
		return target
	}
	target = &GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus{}

	target.State = m.GetState()

	target.Message = m.GetMessage()

	target.ObservedRevision = m.GetObservedRevision()

	if h, ok := interface{}(m.GetObservedOperator()).(clone.Cloner); ok {
		target.ObservedOperator = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.IstioOperatorSpec)
	} else {
		target.ObservedOperator = proto.Clone(m.GetObservedOperator()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.IstioOperatorSpec)
	}

	return target
}