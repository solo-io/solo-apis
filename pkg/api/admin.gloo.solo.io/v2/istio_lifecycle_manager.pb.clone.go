// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/admin/v2/istio_lifecycle_manager.proto

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
func (m *IstioLifecycleManagerSpec) Clone() proto.Message {
	var target *IstioLifecycleManagerSpec
	if m == nil {
		return target
	}
	target = &IstioLifecycleManagerSpec{}

	if m.GetInstallations() != nil {
		target.Installations = make([]*IstioInstallation, len(m.GetInstallations()))
		for idx, v := range m.GetInstallations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Installations[idx] = h.Clone().(*IstioInstallation)
			} else {
				target.Installations[idx] = proto.Clone(v).(*IstioInstallation)
			}

		}
	}

	return target
}

// Clone function
func (m *IstioClusterSelector) Clone() proto.Message {
	var target *IstioClusterSelector
	if m == nil {
		return target
	}
	target = &IstioClusterSelector{}

	target.Name = m.GetName()

	target.DefaultRevision = m.GetDefaultRevision()

	target.TrustDomain = m.GetTrustDomain()

	return target
}

// Clone function
func (m *IstioInstallation) Clone() proto.Message {
	var target *IstioInstallation
	if m == nil {
		return target
	}
	target = &IstioInstallation{}

	target.Revision = m.GetRevision()

	if m.GetClusters() != nil {
		target.Clusters = make([]*IstioClusterSelector, len(m.GetClusters()))
		for idx, v := range m.GetClusters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Clusters[idx] = h.Clone().(*IstioClusterSelector)
			} else {
				target.Clusters[idx] = proto.Clone(v).(*IstioClusterSelector)
			}

		}
	}

	if h, ok := interface{}(m.GetIstioOperatorSpec()).(clone.Cloner); ok {
		target.IstioOperatorSpec = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.IstioOperatorSpec)
	} else {
		target.IstioOperatorSpec = proto.Clone(m.GetIstioOperatorSpec()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.IstioOperatorSpec)
	}

	return target
}

// Clone function
func (m *IstioLifecycleManagerStatus) Clone() proto.Message {
	var target *IstioLifecycleManagerStatus
	if m == nil {
		return target
	}
	target = &IstioLifecycleManagerStatus{}

	if m.GetClusters() != nil {
		target.Clusters = make(map[string]*IstioLifecycleManagerStatus_ClusterStatuses, len(m.GetClusters()))
		for k, v := range m.GetClusters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Clusters[k] = h.Clone().(*IstioLifecycleManagerStatus_ClusterStatuses)
			} else {
				target.Clusters[k] = proto.Clone(v).(*IstioLifecycleManagerStatus_ClusterStatuses)
			}

		}
	}

	return target
}

// Clone function
func (m *IstioLifecycleManagerNewStatus) Clone() proto.Message {
	var target *IstioLifecycleManagerNewStatus
	if m == nil {
		return target
	}
	target = &IstioLifecycleManagerNewStatus{}

	return target
}

// Clone function
func (m *IstioLifecycleManagerReport) Clone() proto.Message {
	var target *IstioLifecycleManagerReport
	if m == nil {
		return target
	}
	target = &IstioLifecycleManagerReport{}

	return target
}

// Clone function
func (m *IstioLifecycleManagerStatus_ClusterStatuses) Clone() proto.Message {
	var target *IstioLifecycleManagerStatus_ClusterStatuses
	if m == nil {
		return target
	}
	target = &IstioLifecycleManagerStatus_ClusterStatuses{}

	if m.GetInstallations() != nil {
		target.Installations = make(map[string]*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus, len(m.GetInstallations()))
		for k, v := range m.GetInstallations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Installations[k] = h.Clone().(*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus)
			} else {
				target.Installations[k] = proto.Clone(v).(*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus)
			}

		}
	}

	return target
}

// Clone function
func (m *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) Clone() proto.Message {
	var target *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus
	if m == nil {
		return target
	}
	target = &IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus{}

	target.State = m.GetState()

	target.Message = m.GetMessage()

	target.ObservedRevision = m.GetObservedRevision()

	if h, ok := interface{}(m.GetObservedOperator()).(clone.Cloner); ok {
		target.ObservedOperator = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.IstioOperatorSpec)
	} else {
		target.ObservedOperator = proto.Clone(m.GetObservedOperator()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.IstioOperatorSpec)
	}

	return target
}
