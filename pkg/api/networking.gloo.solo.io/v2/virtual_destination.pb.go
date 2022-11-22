// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/networking/v2/virtual_destination.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v2 "github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// VirtualDestinations define groupings of backing destinations (for network traffic).
type VirtualDestinationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional: The set of custom hosts for which this virtual destination will serve traffic.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Selectors for the backing K8s services that comprise this VirtualDestination.
	// A service will be selected if it matches any of the given selectors.
	// Currently only one K8s Service can be selected per cluster. If more than one service is selected
	// within a cluster the VirtualDestination will be invalid and will not be translated.
	// When a request is routed through the VirtualDestination, it will be forwarded to one of the backing services, selected at random.
	// (To forward to the service on the local cluster only, a FailoverPolicy and/or OutlierDetectionPolicy must be configured.)
	// If a deployment is unavailable, requests will not be forwarded to that deployment.
	Services []*v2.ObjectSelector `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	// Selectors for the backing External services that comprise this VirtualDestination.
	// An external service will be selected if it matches any of the given selectors.
	// Multiple External Services can be selected.
	ExternalServices []*v2.ObjectSelector `protobuf:"bytes,3,rep,name=external_services,json=externalServices,proto3" json:"external_services,omitempty"`
	// Required: The ports on which the VirtualDestination will serve traffic. Must have at least one port.
	Ports []*VirtualDestinationSpec_PortMapping `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	// Optional: Client mode determines how the VirtualDestination will be translated.
	// If nil, the mode is inherited from the WorkspaceSettings defined by the admin.
	ClientMode *v2.ClientMode `protobuf:"bytes,5,opt,name=client_mode,json=clientMode,proto3" json:"client_mode,omitempty"`
}

func (x *VirtualDestinationSpec) Reset() {
	*x = VirtualDestinationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualDestinationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualDestinationSpec) ProtoMessage() {}

func (x *VirtualDestinationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualDestinationSpec.ProtoReflect.Descriptor instead.
func (*VirtualDestinationSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescGZIP(), []int{0}
}

func (x *VirtualDestinationSpec) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *VirtualDestinationSpec) GetServices() []*v2.ObjectSelector {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *VirtualDestinationSpec) GetExternalServices() []*v2.ObjectSelector {
	if x != nil {
		return x.ExternalServices
	}
	return nil
}

func (x *VirtualDestinationSpec) GetPorts() []*VirtualDestinationSpec_PortMapping {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *VirtualDestinationSpec) GetClientMode() *v2.ClientMode {
	if x != nil {
		return x.ClientMode
	}
	return nil
}

// reflects the status of the VirtualDestination
type VirtualDestinationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global *v2.GenericGlobalStatus `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.WorkspaceStatus `protobuf:"bytes,2,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Map of policy GVK to policy references for all policies applied on this
	// resource.
	AppliedDestinationPolicies map[string]*v2.AppliedDestinationPortPolicies `protobuf:"bytes,3,rep,name=applied_destination_policies,json=appliedDestinationPolicies,proto3" json:"applied_destination_policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Destinations that back this virtual destination
	SelectedBackingServices []*v2.DestinationReference `protobuf:"bytes,4,rep,name=selected_backing_services,json=selectedBackingServices,proto3" json:"selected_backing_services,omitempty"`
	// Name of Workspace that owns VirtualDestination
	OwnerWorkspace *v2.OwnerWorkspace `protobuf:"bytes,5,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
}

func (x *VirtualDestinationStatus) Reset() {
	*x = VirtualDestinationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualDestinationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualDestinationStatus) ProtoMessage() {}

func (x *VirtualDestinationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualDestinationStatus.ProtoReflect.Descriptor instead.
func (*VirtualDestinationStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescGZIP(), []int{1}
}

func (x *VirtualDestinationStatus) GetGlobal() *v2.GenericGlobalStatus {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *VirtualDestinationStatus) GetWorkspaces() map[string]*v2.WorkspaceStatus {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *VirtualDestinationStatus) GetAppliedDestinationPolicies() map[string]*v2.AppliedDestinationPortPolicies {
	if x != nil {
		return x.AppliedDestinationPolicies
	}
	return nil
}

func (x *VirtualDestinationStatus) GetSelectedBackingServices() []*v2.DestinationReference {
	if x != nil {
		return x.SelectedBackingServices
	}
	return nil
}

func (x *VirtualDestinationStatus) GetOwnerWorkspace() *v2.OwnerWorkspace {
	if x != nil {
		return x.OwnerWorkspace
	}
	return nil
}

// PortMapping establishes a new port that will be exposed on a VirtualDestination.
type VirtualDestinationSpec_PortMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The port number. Must be a valid, non-negative integer port number.
	Number uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// The protocol used in communication with this destination
	// MUST be one of the following: HTTP, HTTPS, GRPC, HTTP2, MONGO, TCP, TLS.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// (optional): The port number or name used to match the corresponding port on the
	// VirtualDestination's backing Services and ExternalServices.
	// All of the backing services for this VirtualDestination must contain
	// this port, matching by name or number.
	// If unspecified, will default to the value of the port number field above.
	TargetPort *v2.PortSelector `protobuf:"bytes,3,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
}

func (x *VirtualDestinationSpec_PortMapping) Reset() {
	*x = VirtualDestinationSpec_PortMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualDestinationSpec_PortMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualDestinationSpec_PortMapping) ProtoMessage() {}

func (x *VirtualDestinationSpec_PortMapping) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualDestinationSpec_PortMapping.ProtoReflect.Descriptor instead.
func (*VirtualDestinationSpec_PortMapping) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescGZIP(), []int{0, 0}
}

func (x *VirtualDestinationSpec_PortMapping) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *VirtualDestinationSpec_PortMapping) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *VirtualDestinationSpec_PortMapping) GetTargetPort() *v2.PortSelector {
	if x != nil {
		return x.TargetPort
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDesc = []byte{
	0x0a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x51, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xde, 0x03, 0x0a, 0x16, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14,
	0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x1a, 0x85, 0x01, 0x0a,
	0x0b, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x42, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x22, 0xf4, 0x05, 0x0a, 0x18, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x12, 0x61, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x93, 0x01, 0x0a, 0x1c, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x1a, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x65, 0x0a, 0x19,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x17, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x1a, 0x63, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x82, 0x01, 0x0a, 0x1f, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x49, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x4d, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5,
	0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_goTypes = []interface{}{
	(*VirtualDestinationSpec)(nil),             // 0: networking.gloo.solo.io.VirtualDestinationSpec
	(*VirtualDestinationStatus)(nil),           // 1: networking.gloo.solo.io.VirtualDestinationStatus
	(*VirtualDestinationSpec_PortMapping)(nil), // 2: networking.gloo.solo.io.VirtualDestinationSpec.PortMapping
	nil,                             // 3: networking.gloo.solo.io.VirtualDestinationStatus.WorkspacesEntry
	nil,                             // 4: networking.gloo.solo.io.VirtualDestinationStatus.AppliedDestinationPoliciesEntry
	(*v2.ObjectSelector)(nil),       // 5: common.gloo.solo.io.ObjectSelector
	(*v2.ClientMode)(nil),           // 6: common.gloo.solo.io.ClientMode
	(*v2.GenericGlobalStatus)(nil),  // 7: common.gloo.solo.io.GenericGlobalStatus
	(*v2.DestinationReference)(nil), // 8: common.gloo.solo.io.DestinationReference
	(*v2.OwnerWorkspace)(nil),       // 9: common.gloo.solo.io.OwnerWorkspace
	(*v2.PortSelector)(nil),         // 10: common.gloo.solo.io.PortSelector
	(*v2.WorkspaceStatus)(nil),      // 11: common.gloo.solo.io.WorkspaceStatus
	(*v2.AppliedDestinationPortPolicies)(nil), // 12: common.gloo.solo.io.AppliedDestinationPortPolicies
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_depIdxs = []int32{
	5,  // 0: networking.gloo.solo.io.VirtualDestinationSpec.services:type_name -> common.gloo.solo.io.ObjectSelector
	5,  // 1: networking.gloo.solo.io.VirtualDestinationSpec.external_services:type_name -> common.gloo.solo.io.ObjectSelector
	2,  // 2: networking.gloo.solo.io.VirtualDestinationSpec.ports:type_name -> networking.gloo.solo.io.VirtualDestinationSpec.PortMapping
	6,  // 3: networking.gloo.solo.io.VirtualDestinationSpec.client_mode:type_name -> common.gloo.solo.io.ClientMode
	7,  // 4: networking.gloo.solo.io.VirtualDestinationStatus.global:type_name -> common.gloo.solo.io.GenericGlobalStatus
	3,  // 5: networking.gloo.solo.io.VirtualDestinationStatus.workspaces:type_name -> networking.gloo.solo.io.VirtualDestinationStatus.WorkspacesEntry
	4,  // 6: networking.gloo.solo.io.VirtualDestinationStatus.applied_destination_policies:type_name -> networking.gloo.solo.io.VirtualDestinationStatus.AppliedDestinationPoliciesEntry
	8,  // 7: networking.gloo.solo.io.VirtualDestinationStatus.selected_backing_services:type_name -> common.gloo.solo.io.DestinationReference
	9,  // 8: networking.gloo.solo.io.VirtualDestinationStatus.owner_workspace:type_name -> common.gloo.solo.io.OwnerWorkspace
	10, // 9: networking.gloo.solo.io.VirtualDestinationSpec.PortMapping.target_port:type_name -> common.gloo.solo.io.PortSelector
	11, // 10: networking.gloo.solo.io.VirtualDestinationStatus.WorkspacesEntry.value:type_name -> common.gloo.solo.io.WorkspaceStatus
	12, // 11: networking.gloo.solo.io.VirtualDestinationStatus.AppliedDestinationPoliciesEntry.value:type_name -> common.gloo.solo.io.AppliedDestinationPortPolicies
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualDestinationSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualDestinationStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualDestinationSpec_PortMapping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_networking_v2_virtual_destination_proto_depIdxs = nil
}
