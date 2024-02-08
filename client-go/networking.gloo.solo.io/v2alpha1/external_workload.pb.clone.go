// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/networking/v2alpha1/external_workload.proto

package v2alpha1

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
func (m *ExternalWorkloadSpec) Clone() proto.Message {
	var target *ExternalWorkloadSpec
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec{}

	if m.GetPorts() != nil {
		target.Ports = make([]*ExternalWorkloadSpec_Port, len(m.GetPorts()))
		for idx, v := range m.GetPorts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Ports[idx] = h.Clone().(*ExternalWorkloadSpec_Port)
			} else {
				target.Ports[idx] = proto.Clone(v).(*ExternalWorkloadSpec_Port)
			}

		}
	}

	if h, ok := interface{}(m.GetIdentitySelector()).(clone.Cloner); ok {
		target.IdentitySelector = h.Clone().(*ExternalWorkloadSpec_IdentitySelector)
	} else {
		target.IdentitySelector = proto.Clone(m.GetIdentitySelector()).(*ExternalWorkloadSpec_IdentitySelector)
	}

	if m.GetConnectedClusters() != nil {
		target.ConnectedClusters = make(map[string]string, len(m.GetConnectedClusters()))
		for k, v := range m.GetConnectedClusters() {

			target.ConnectedClusters[k] = v

		}
	}

	if h, ok := interface{}(m.GetReadinessProbe()).(clone.Cloner); ok {
		target.ReadinessProbe = h.Clone().(*ExternalWorkloadSpec_Probe)
	} else {
		target.ReadinessProbe = proto.Clone(m.GetReadinessProbe()).(*ExternalWorkloadSpec_Probe)
	}

	return target
}

// Clone function
func (m *ExternalWorkloadStatus) Clone() proto.Message {
	var target *ExternalWorkloadStatus
	if m == nil {
		return target
	}
	target = &ExternalWorkloadStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	if m.GetNumAppliedPolicies() != nil {
		target.NumAppliedPolicies = make(map[string]uint32, len(m.GetNumAppliedPolicies()))
		for k, v := range m.GetNumAppliedPolicies() {

			target.NumAppliedPolicies[k] = v

		}
	}

	target.OwnedByWorkspace = m.GetOwnedByWorkspace()

	return target
}

// Clone function
func (m *ExternalWorkloadReport) Clone() proto.Message {
	var target *ExternalWorkloadReport
	if m == nil {
		return target
	}
	target = &ExternalWorkloadReport{}

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

	if m.GetAppliedDestinationPolicies() != nil {
		target.AppliedDestinationPolicies = make(map[string]*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.AppliedDestinationPortPolicies, len(m.GetAppliedDestinationPolicies()))
		for k, v := range m.GetAppliedDestinationPolicies() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AppliedDestinationPolicies[k] = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.AppliedDestinationPortPolicies)
			} else {
				target.AppliedDestinationPolicies[k] = proto.Clone(v).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.AppliedDestinationPortPolicies)
			}

		}
	}

	target.OwnerWorkspace = m.GetOwnerWorkspace()

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_Port) Clone() proto.Message {
	var target *ExternalWorkloadSpec_Port
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_Port{}

	target.Name = m.GetName()

	target.Protocol = m.GetProtocol()

	target.Number = m.GetNumber()

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_IdentitySelector) Clone() proto.Message {
	var target *ExternalWorkloadSpec_IdentitySelector
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_IdentitySelector{}

	if m.GetAws() != nil {
		target.Aws = make([]*ExternalWorkloadSpec_IdentitySelector_AWS, len(m.GetAws()))
		for idx, v := range m.GetAws() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Aws[idx] = h.Clone().(*ExternalWorkloadSpec_IdentitySelector_AWS)
			} else {
				target.Aws[idx] = proto.Clone(v).(*ExternalWorkloadSpec_IdentitySelector_AWS)
			}

		}
	}

	if m.GetGcp() != nil {
		target.Gcp = make([]*ExternalWorkloadSpec_IdentitySelector_GCP, len(m.GetGcp()))
		for idx, v := range m.GetGcp() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Gcp[idx] = h.Clone().(*ExternalWorkloadSpec_IdentitySelector_GCP)
			} else {
				target.Gcp[idx] = proto.Clone(v).(*ExternalWorkloadSpec_IdentitySelector_GCP)
			}

		}
	}

	if m.GetAzure() != nil {
		target.Azure = make([]*ExternalWorkloadSpec_IdentitySelector_Azure, len(m.GetAzure()))
		for idx, v := range m.GetAzure() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Azure[idx] = h.Clone().(*ExternalWorkloadSpec_IdentitySelector_Azure)
			} else {
				target.Azure[idx] = proto.Clone(v).(*ExternalWorkloadSpec_IdentitySelector_Azure)
			}

		}
	}

	if h, ok := interface{}(m.GetJoinToken()).(clone.Cloner); ok {
		target.JoinToken = h.Clone().(*ExternalWorkloadSpec_IdentitySelector_JoinToken)
	} else {
		target.JoinToken = proto.Clone(m.GetJoinToken()).(*ExternalWorkloadSpec_IdentitySelector_JoinToken)
	}

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_Probe) Clone() proto.Message {
	var target *ExternalWorkloadSpec_Probe
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_Probe{}

	if h, ok := interface{}(m.GetInitialDelaySeconds()).(clone.Cloner); ok {
		target.InitialDelaySeconds = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.InitialDelaySeconds = proto.Clone(m.GetInitialDelaySeconds()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetTimeoutSeconds()).(clone.Cloner); ok {
		target.TimeoutSeconds = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.TimeoutSeconds = proto.Clone(m.GetTimeoutSeconds()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetPeriodSeconds()).(clone.Cloner); ok {
		target.PeriodSeconds = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.PeriodSeconds = proto.Clone(m.GetPeriodSeconds()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetSuccessThreshold()).(clone.Cloner); ok {
		target.SuccessThreshold = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.SuccessThreshold = proto.Clone(m.GetSuccessThreshold()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetFailureThreshold()).(clone.Cloner); ok {
		target.FailureThreshold = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.FailureThreshold = proto.Clone(m.GetFailureThreshold()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	switch m.Handler.(type) {

	case *ExternalWorkloadSpec_Probe_HttpGet:

		if h, ok := interface{}(m.GetHttpGet()).(clone.Cloner); ok {
			target.Handler = &ExternalWorkloadSpec_Probe_HttpGet{
				HttpGet: h.Clone().(*ExternalWorkloadSpec_Probe_HTTPGetConfig),
			}
		} else {
			target.Handler = &ExternalWorkloadSpec_Probe_HttpGet{
				HttpGet: proto.Clone(m.GetHttpGet()).(*ExternalWorkloadSpec_Probe_HTTPGetConfig),
			}
		}

	case *ExternalWorkloadSpec_Probe_TcpSocket:

		if h, ok := interface{}(m.GetTcpSocket()).(clone.Cloner); ok {
			target.Handler = &ExternalWorkloadSpec_Probe_TcpSocket{
				TcpSocket: h.Clone().(*ExternalWorkloadSpec_Probe_TCPSocketConfig),
			}
		} else {
			target.Handler = &ExternalWorkloadSpec_Probe_TcpSocket{
				TcpSocket: proto.Clone(m.GetTcpSocket()).(*ExternalWorkloadSpec_Probe_TCPSocketConfig),
			}
		}

	case *ExternalWorkloadSpec_Probe_Exec:

		if h, ok := interface{}(m.GetExec()).(clone.Cloner); ok {
			target.Handler = &ExternalWorkloadSpec_Probe_Exec{
				Exec: h.Clone().(*ExternalWorkloadSpec_Probe_ExecConfig),
			}
		} else {
			target.Handler = &ExternalWorkloadSpec_Probe_Exec{
				Exec: proto.Clone(m.GetExec()).(*ExternalWorkloadSpec_Probe_ExecConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_IdentitySelector_AWS) Clone() proto.Message {
	var target *ExternalWorkloadSpec_IdentitySelector_AWS
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_IdentitySelector_AWS{}

	target.IamRole = m.GetIamRole()

	target.SecurityGroupName = m.GetSecurityGroupName()

	target.SecurityGroupId = m.GetSecurityGroupId()

	target.ImageId = m.GetImageId()

	target.InstanceId = m.GetInstanceId()

	target.Zone = m.GetZone()

	target.Region = m.GetRegion()

	if h, ok := interface{}(m.GetTag()).(clone.Cloner); ok {
		target.Tag = h.Clone().(*ExternalWorkloadSpec_IdentitySelector_AWS_Tag)
	} else {
		target.Tag = proto.Clone(m.GetTag()).(*ExternalWorkloadSpec_IdentitySelector_AWS_Tag)
	}

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_IdentitySelector_GCP) Clone() proto.Message {
	var target *ExternalWorkloadSpec_IdentitySelector_GCP
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_IdentitySelector_GCP{}

	target.ServiceAccount = m.GetServiceAccount()

	target.Name = m.GetName()

	target.Tag = m.GetTag()

	target.ProjectId = m.GetProjectId()

	target.Zone = m.GetZone()

	if h, ok := interface{}(m.GetLabel()).(clone.Cloner); ok {
		target.Label = h.Clone().(*ExternalWorkloadSpec_IdentitySelector_GCP_Label)
	} else {
		target.Label = proto.Clone(m.GetLabel()).(*ExternalWorkloadSpec_IdentitySelector_GCP_Label)
	}

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_IdentitySelector_Azure) Clone() proto.Message {
	var target *ExternalWorkloadSpec_IdentitySelector_Azure
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_IdentitySelector_Azure{}

	target.SubscriptionId = m.GetSubscriptionId()

	target.SecurityGroup = m.GetSecurityGroup()

	target.VirtualNetwork = m.GetVirtualNetwork()

	target.Subnet = m.GetSubnet()

	target.Name = m.GetName()

	target.ResourceGroup = m.GetResourceGroup()

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_IdentitySelector_JoinToken) Clone() proto.Message {
	var target *ExternalWorkloadSpec_IdentitySelector_JoinToken
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_IdentitySelector_JoinToken{}

	target.Enable = m.GetEnable()

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_IdentitySelector_AWS_Tag) Clone() proto.Message {
	var target *ExternalWorkloadSpec_IdentitySelector_AWS_Tag
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_IdentitySelector_AWS_Tag{}

	target.Key = m.GetKey()

	target.Value = m.GetValue()

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_IdentitySelector_GCP_Label) Clone() proto.Message {
	var target *ExternalWorkloadSpec_IdentitySelector_GCP_Label
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_IdentitySelector_GCP_Label{}

	target.Key = m.GetKey()

	target.Value = m.GetValue()

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_Probe_HTTPGetConfig) Clone() proto.Message {
	var target *ExternalWorkloadSpec_Probe_HTTPGetConfig
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_Probe_HTTPGetConfig{}

	target.Port = m.GetPort()

	target.Path = m.GetPath()

	target.Scheme = m.GetScheme()

	if m.GetHttpHeaders() != nil {
		target.HttpHeaders = make([]*ExternalWorkloadSpec_Probe_HTTPHeader, len(m.GetHttpHeaders()))
		for idx, v := range m.GetHttpHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.HttpHeaders[idx] = h.Clone().(*ExternalWorkloadSpec_Probe_HTTPHeader)
			} else {
				target.HttpHeaders[idx] = proto.Clone(v).(*ExternalWorkloadSpec_Probe_HTTPHeader)
			}

		}
	}

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_Probe_HTTPHeader) Clone() proto.Message {
	var target *ExternalWorkloadSpec_Probe_HTTPHeader
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_Probe_HTTPHeader{}

	target.Name = m.GetName()

	target.Value = m.GetValue()

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_Probe_TCPSocketConfig) Clone() proto.Message {
	var target *ExternalWorkloadSpec_Probe_TCPSocketConfig
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_Probe_TCPSocketConfig{}

	target.Host = m.GetHost()

	target.Port = m.GetPort()

	return target
}

// Clone function
func (m *ExternalWorkloadSpec_Probe_ExecConfig) Clone() proto.Message {
	var target *ExternalWorkloadSpec_Probe_ExecConfig
	if m == nil {
		return target
	}
	target = &ExternalWorkloadSpec_Probe_ExecConfig{}

	if m.GetCommand() != nil {
		target.Command = make([]string, len(m.GetCommand()))
		for idx, v := range m.GetCommand() {

			target.Command[idx] = v

		}
	}

	return target
}
