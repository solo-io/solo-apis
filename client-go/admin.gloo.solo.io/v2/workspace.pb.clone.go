// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/workspace.proto

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
func (m *WorkspaceSpec) Clone() proto.Message {
	var target *WorkspaceSpec
	if m == nil {
		return target
	}
	target = &WorkspaceSpec{}

	if m.GetWorkloadClusters() != nil {
		target.WorkloadClusters = make([]*ClusterSelector, len(m.GetWorkloadClusters()))
		for idx, v := range m.GetWorkloadClusters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.WorkloadClusters[idx] = h.Clone().(*ClusterSelector)
			} else {
				target.WorkloadClusters[idx] = proto.Clone(v).(*ClusterSelector)
			}

		}
	}

	return target
}

// Clone function
func (m *ClusterSelector) Clone() proto.Message {
	var target *ClusterSelector
	if m == nil {
		return target
	}
	target = &ClusterSelector{}

	target.Name = m.GetName()

	if m.GetSelector() != nil {
		target.Selector = make(map[string]string, len(m.GetSelector()))
		for k, v := range m.GetSelector() {

			target.Selector[k] = v

		}
	}

	if m.GetNamespaces() != nil {
		target.Namespaces = make([]*ClusterSelector_NamespaceSelector, len(m.GetNamespaces()))
		for idx, v := range m.GetNamespaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Namespaces[idx] = h.Clone().(*ClusterSelector_NamespaceSelector)
			} else {
				target.Namespaces[idx] = proto.Clone(v).(*ClusterSelector_NamespaceSelector)
			}

		}
	}

	target.ConfigEnabled = m.GetConfigEnabled()

	return target
}

// Clone function
func (m *WorkspaceStatus) Clone() proto.Message {
	var target *WorkspaceStatus
	if m == nil {
		return target
	}
	target = &WorkspaceStatus{}

	if h, ok := interface{}(m.GetGeneric()).(clone.Cloner); ok {
		target.Generic = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.GenericContextStatus)
	} else {
		target.Generic = proto.Clone(m.GetGeneric()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.GenericContextStatus)
	}

	if h, ok := interface{}(m.GetWorkspaceSettings()).(clone.Cloner); ok {
		target.WorkspaceSettings = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	} else {
		target.WorkspaceSettings = proto.Clone(m.GetWorkspaceSettings()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
	}

	if m.GetClusters() != nil {
		target.Clusters = make([]*WorkspaceStatus_SelectedCluster, len(m.GetClusters()))
		for idx, v := range m.GetClusters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Clusters[idx] = h.Clone().(*WorkspaceStatus_SelectedCluster)
			} else {
				target.Clusters[idx] = proto.Clone(v).(*WorkspaceStatus_SelectedCluster)
			}

		}
	}

	if m.GetResources() != nil {
		target.Resources = make(map[string]*WorkspaceStatus_SourceType, len(m.GetResources()))
		for k, v := range m.GetResources() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Resources[k] = h.Clone().(*WorkspaceStatus_SourceType)
			} else {
				target.Resources[k] = proto.Clone(v).(*WorkspaceStatus_SourceType)
			}

		}
	}

	if m.GetPolicyCounts() != nil {
		target.PolicyCounts = make(map[string]int32, len(m.GetPolicyCounts()))
		for k, v := range m.GetPolicyCounts() {

			target.PolicyCounts[k] = v

		}
	}

	if m.GetDestinationCounts() != nil {
		target.DestinationCounts = make(map[string]int32, len(m.GetDestinationCounts()))
		for k, v := range m.GetDestinationCounts() {

			target.DestinationCounts[k] = v

		}
	}

	if m.GetImportedWorkspaces() != nil {
		target.ImportedWorkspaces = make([]string, len(m.GetImportedWorkspaces()))
		for idx, v := range m.GetImportedWorkspaces() {

			target.ImportedWorkspaces[idx] = v

		}
	}

	return target
}

// Clone function
func (m *ClusterSelector_NamespaceSelector) Clone() proto.Message {
	var target *ClusterSelector_NamespaceSelector
	if m == nil {
		return target
	}
	target = &ClusterSelector_NamespaceSelector{}

	target.Name = m.GetName()

	target.ConfigEnabled = m.GetConfigEnabled()

	if m.GetLabels() != nil {
		target.Labels = make(map[string]string, len(m.GetLabels()))
		for k, v := range m.GetLabels() {

			target.Labels[k] = v

		}
	}

	return target
}

// Clone function
func (m *WorkspaceStatus_SelectedCluster) Clone() proto.Message {
	var target *WorkspaceStatus_SelectedCluster
	if m == nil {
		return target
	}
	target = &WorkspaceStatus_SelectedCluster{}

	target.Name = m.GetName()

	if m.GetNamespaces() != nil {
		target.Namespaces = make([]string, len(m.GetNamespaces()))
		for idx, v := range m.GetNamespaces() {

			target.Namespaces[idx] = v

		}
	}

	return target
}

// Clone function
func (m *WorkspaceStatus_SourceType) Clone() proto.Message {
	var target *WorkspaceStatus_SourceType
	if m == nil {
		return target
	}
	target = &WorkspaceStatus_SourceType{}

	if h, ok := interface{}(m.GetImported()).(clone.Cloner); ok {
		target.Imported = h.Clone().(*WorkspaceStatus_Imported)
	} else {
		target.Imported = proto.Clone(m.GetImported()).(*WorkspaceStatus_Imported)
	}

	if h, ok := interface{}(m.GetExported()).(clone.Cloner); ok {
		target.Exported = h.Clone().(*WorkspaceStatus_Exported)
	} else {
		target.Exported = proto.Clone(m.GetExported()).(*WorkspaceStatus_Exported)
	}

	if h, ok := interface{}(m.GetPrivate()).(clone.Cloner); ok {
		target.Private = h.Clone().(*WorkspaceStatus_Private)
	} else {
		target.Private = proto.Clone(m.GetPrivate()).(*WorkspaceStatus_Private)
	}

	return target
}

// Clone function
func (m *WorkspaceStatus_Imported) Clone() proto.Message {
	var target *WorkspaceStatus_Imported
	if m == nil {
		return target
	}
	target = &WorkspaceStatus_Imported{}

	if m.GetObjectReferences() != nil {
		target.ObjectReferences = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference, len(m.GetObjectReferences()))
		for idx, v := range m.GetObjectReferences() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ObjectReferences[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.ObjectReferences[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	return target
}

// Clone function
func (m *WorkspaceStatus_Exported) Clone() proto.Message {
	var target *WorkspaceStatus_Exported
	if m == nil {
		return target
	}
	target = &WorkspaceStatus_Exported{}

	if m.GetObjectReferences() != nil {
		target.ObjectReferences = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference, len(m.GetObjectReferences()))
		for idx, v := range m.GetObjectReferences() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ObjectReferences[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.ObjectReferences[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	return target
}

// Clone function
func (m *WorkspaceStatus_Private) Clone() proto.Message {
	var target *WorkspaceStatus_Private
	if m == nil {
		return target
	}
	target = &WorkspaceStatus_Private{}

	if m.GetObjectReferences() != nil {
		target.ObjectReferences = make([]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference, len(m.GetObjectReferences()))
		for idx, v := range m.GetObjectReferences() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ObjectReferences[idx] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			} else {
				target.ObjectReferences[idx] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.ObjectReference)
			}

		}
	}

	return target
}
