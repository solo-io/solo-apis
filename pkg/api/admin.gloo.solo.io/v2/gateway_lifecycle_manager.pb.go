// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/admin/v2/gateway_lifecycle_manager.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/cue/encoding/protobuf/cue"
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

// The state of an Istio installation.
type GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State int32

const (
	// Waiting for resources to be installed or updated.
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_PENDING GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 0
	// Gloo Mesh server encountered a problem while attempting
	// to install Istio Gateway.
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_FAILED GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 1
	// Could not select a control plane
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_NO_CONTROL_PLANE_AVAILABLE GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 2
	// In the process of installing Istio gateway.
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_INSTALLING_GATEWAY GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 3
	// All Istio components are successfully installed and healthy.
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_HEALTHY GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 4
	// The Istio installation is no longer healthy.
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_UNHEALTHY GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 5
	// The gateway IstioOperator CR is in an 'ACTION_REQUIRED' state, please check logs of IstioOperator deployment for more info.
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_ACTION_REQUIRED GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 6
	// The gateway IstioOperator CR is in an 'UPDATING' state
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_UPDATING_GATEWAY GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 7
	// The gateway IstioOperator CR is in an 'RECONCILING' state
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_RECONCILING_GATEWAY GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 8
	// The gateway installation state could not be determined
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_UNKNOWN GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 9
)

// Enum value maps for GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State.
var (
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State_name = map[int32]string{
		0: "PENDING",
		1: "FAILED",
		2: "NO_CONTROL_PLANE_AVAILABLE",
		3: "INSTALLING_GATEWAY",
		4: "HEALTHY",
		5: "UNHEALTHY",
		6: "ACTION_REQUIRED",
		7: "UPDATING_GATEWAY",
		8: "RECONCILING_GATEWAY",
		9: "UNKNOWN",
	}
	GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State_value = map[string]int32{
		"PENDING":                    0,
		"FAILED":                     1,
		"NO_CONTROL_PLANE_AVAILABLE": 2,
		"INSTALLING_GATEWAY":         3,
		"HEALTHY":                    4,
		"UNHEALTHY":                  5,
		"ACTION_REQUIRED":            6,
		"UPDATING_GATEWAY":           7,
		"RECONCILING_GATEWAY":        8,
		"UNKNOWN":                    9,
	}
)

func (x GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) Enum() *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State {
	p := new(GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State)
	*p = x
	return p
}

func (x GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_enumTypes[0].Descriptor()
}

func (GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_enumTypes[0]
}

func (x GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State.Descriptor instead.
func (GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescGZIP(), []int{3, 1, 1, 0}
}

type GatewayLifecycleManagerSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of Istio gateway installations.
	// Any components that are NOT related to the gateway will be ignored
	// Only 1 installation is allowed per revision per cluster
	Installations []*GatewayInstallation `protobuf:"bytes,1,rep,name=installations,proto3" json:"installations,omitempty"`
}

func (x *GatewayLifecycleManagerSpec) Reset() {
	*x = GatewayLifecycleManagerSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayLifecycleManagerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayLifecycleManagerSpec) ProtoMessage() {}

func (x *GatewayLifecycleManagerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayLifecycleManagerSpec.ProtoReflect.Descriptor instead.
func (*GatewayLifecycleManagerSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayLifecycleManagerSpec) GetInstallations() []*GatewayInstallation {
	if x != nil {
		return x.Installations
	}
	return nil
}

type GatewayClusterSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the cluster that Istio should be installed in. This should
	// match the registered cluster name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional: defaults to false
	// When true this installation will be the Active gateway that primary service traffic will be routed through for the selected cluster(s)
	// If the GatewayLifecycleManagerSpec defines a service, this field will switch over service selectors based off the gatewayRevsion. Updating this can be useful for switching over Canary deployments.
	ActiveGateway bool `protobuf:"varint,2,opt,name=active_gateway,json=activeGateway,proto3" json:"active_gateway,omitempty"`
	// [Optional] The trust domain value that should be set for this cluster's
	// Istio installations. This value will be set in the installation's mesh
	// config. (See https://istio.io/latest/docs/reference/config/istio.mesh.v1alpha1)
	// Defaults to the cluster's name.
	TrustDomain string `protobuf:"bytes,5,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
}

func (x *GatewayClusterSelector) Reset() {
	*x = GatewayClusterSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayClusterSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayClusterSelector) ProtoMessage() {}

func (x *GatewayClusterSelector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayClusterSelector.ProtoReflect.Descriptor instead.
func (*GatewayClusterSelector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayClusterSelector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GatewayClusterSelector) GetActiveGateway() bool {
	if x != nil {
		return x.ActiveGateway
	}
	return false
}

func (x *GatewayClusterSelector) GetTrustDomain() string {
	if x != nil {
		return x.TrustDomain
	}
	return ""
}

type GatewayInstallation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional: default to the gatewayRevision; Selects the control plane revision to be used by the gateway. If not found no gateway will be created.
	ControlPlaneRevision string `protobuf:"bytes,1,opt,name=control_plane_revision,json=controlPlaneRevision,proto3" json:"control_plane_revision,omitempty"`
	// gatewayRevision that will be used to represent this installation.
	GatewayRevision string `protobuf:"bytes,2,opt,name=gateway_revision,json=gatewayRevision,proto3" json:"gateway_revision,omitempty"`
	// The clusters where the Istio should be installed.
	Clusters []*GatewayClusterSelector `protobuf:"bytes,3,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// Specs for the IstioOperators that should be applied.
	// See https://istio.io/latest/docs/reference/config/istio.operator.v1alpha1/
	IstioOperatorSpec *v2.IstioOperatorSpec `protobuf:"bytes,4,opt,name=istio_operator_spec,json=istioOperatorSpec,proto3" json:"istio_operator_spec,omitempty"`
}

func (x *GatewayInstallation) Reset() {
	*x = GatewayInstallation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayInstallation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayInstallation) ProtoMessage() {}

func (x *GatewayInstallation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayInstallation.ProtoReflect.Descriptor instead.
func (*GatewayInstallation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescGZIP(), []int{2}
}

func (x *GatewayInstallation) GetControlPlaneRevision() string {
	if x != nil {
		return x.ControlPlaneRevision
	}
	return ""
}

func (x *GatewayInstallation) GetGatewayRevision() string {
	if x != nil {
		return x.GatewayRevision
	}
	return ""
}

func (x *GatewayInstallation) GetClusters() []*GatewayClusterSelector {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *GatewayInstallation) GetIstioOperatorSpec() *v2.IstioOperatorSpec {
	if x != nil {
		return x.IstioOperatorSpec
	}
	return nil
}

type GatewayLifecycleManagerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of each Istio gateway installation that is being managed by Gloo Mesh, where
	// the key is the cluster name of the installation.
	Clusters map[string]*GatewayLifecycleManagerStatus_ClusterStatuses `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GatewayLifecycleManagerStatus) Reset() {
	*x = GatewayLifecycleManagerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayLifecycleManagerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayLifecycleManagerStatus) ProtoMessage() {}

func (x *GatewayLifecycleManagerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayLifecycleManagerStatus.ProtoReflect.Descriptor instead.
func (*GatewayLifecycleManagerStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescGZIP(), []int{3}
}

func (x *GatewayLifecycleManagerStatus) GetClusters() map[string]*GatewayLifecycleManagerStatus_ClusterStatuses {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type GatewayLifecycleManagerStatus_ClusterStatuses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Istio installations by revision.
	Installations map[string]*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus `protobuf:"bytes,1,rep,name=installations,proto3" json:"installations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses) Reset() {
	*x = GatewayLifecycleManagerStatus_ClusterStatuses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayLifecycleManagerStatus_ClusterStatuses) ProtoMessage() {}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayLifecycleManagerStatus_ClusterStatuses.ProtoReflect.Descriptor instead.
func (*GatewayLifecycleManagerStatus_ClusterStatuses) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescGZIP(), []int{3, 1}
}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses) GetInstallations() map[string]*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus {
	if x != nil {
		return x.Installations
	}
	return nil
}

type GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// State of the Istio installation
	State GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=admin.gloo.solo.io.GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State" json:"state,omitempty"`
	// A human readable message about the current state of the GatewayInstallationInstance.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// operator that is currently deployed for this revision
	ObservedOperator *v2.IstioOperatorSpec `protobuf:"bytes,4,opt,name=observed_operator,json=observedOperator,proto3" json:"observed_operator,omitempty"`
}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) Reset() {
	*x = GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) ProtoMessage() {}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus.ProtoReflect.Descriptor instead.
func (*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescGZIP(), []int{3, 1, 1}
}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) GetState() GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State {
	if x != nil {
		return x.State
	}
	return GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_PENDING
}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus) GetObservedOperator() *v2.IstioOperatorSpec {
	if x != nil {
		return x.ObservedOperator
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDesc = []byte{
	0x0a, 0x60, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x32, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x75, 0x65, 0x2f,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x63, 0x75, 0x65, 0x2f, 0x63, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x1b, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4d, 0x0a, 0x0d, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x76, 0x0a, 0x16, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x96, 0x02, 0x0a, 0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x29, 0x0a, 0x10, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x56, 0x0a, 0x13, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x11, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x22, 0xe5, 0x07, 0x0a, 0x1d, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5b, 0x0a, 0x08,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x7e, 0x0a, 0x0d, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x57, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xe6, 0x05, 0x0a, 0x0f, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x7a, 0x0a,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x54, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x96, 0x01, 0x0a, 0x12, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x6a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x54, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0xbd, 0x03, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x70, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x11, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x10, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xc5, 0x01, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1e, 0x0a,
	0x1a, 0x4e, 0x4f, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x4e,
	0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a,
	0x12, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x47, 0x41, 0x54, 0x45,
	0x57, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10,
	0x05, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13,
	0x52, 0x45, 0x43, 0x4f, 0x4e, 0x43, 0x49, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x47, 0x41, 0x54, 0x45,
	0x57, 0x41, 0x59, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x09, 0x42, 0x48, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32,
	0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_goTypes = []interface{}{
	(GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State)(0), // 0: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.InstallationStatus.State
	(*GatewayLifecycleManagerSpec)(nil),                                         // 1: admin.gloo.solo.io.GatewayLifecycleManagerSpec
	(*GatewayClusterSelector)(nil),                                              // 2: admin.gloo.solo.io.GatewayClusterSelector
	(*GatewayInstallation)(nil),                                                 // 3: admin.gloo.solo.io.GatewayInstallation
	(*GatewayLifecycleManagerStatus)(nil),                                       // 4: admin.gloo.solo.io.GatewayLifecycleManagerStatus
	nil,                                                                         // 5: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClustersEntry
	(*GatewayLifecycleManagerStatus_ClusterStatuses)(nil),                       // 6: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses
	nil, // 7: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.InstallationsEntry
	(*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus)(nil), // 8: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.InstallationStatus
	(*v2.IstioOperatorSpec)(nil), // 9: common.gloo.solo.io.IstioOperatorSpec
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_depIdxs = []int32{
	3, // 0: admin.gloo.solo.io.GatewayLifecycleManagerSpec.installations:type_name -> admin.gloo.solo.io.GatewayInstallation
	2, // 1: admin.gloo.solo.io.GatewayInstallation.clusters:type_name -> admin.gloo.solo.io.GatewayClusterSelector
	9, // 2: admin.gloo.solo.io.GatewayInstallation.istio_operator_spec:type_name -> common.gloo.solo.io.IstioOperatorSpec
	5, // 3: admin.gloo.solo.io.GatewayLifecycleManagerStatus.clusters:type_name -> admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClustersEntry
	6, // 4: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClustersEntry.value:type_name -> admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses
	7, // 5: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.installations:type_name -> admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.InstallationsEntry
	8, // 6: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.InstallationsEntry.value:type_name -> admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.InstallationStatus
	0, // 7: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.InstallationStatus.state:type_name -> admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.InstallationStatus.State
	9, // 8: admin.gloo.solo.io.GatewayLifecycleManagerStatus.ClusterStatuses.InstallationStatus.observed_operator:type_name -> common.gloo.solo.io.IstioOperatorSpec
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayLifecycleManagerSpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayClusterSelector); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayInstallation); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayLifecycleManagerStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayLifecycleManagerStatus_ClusterStatuses); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayLifecycleManagerStatus_ClusterStatuses_InstallationStatus); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_admin_v2_gateway_lifecycle_manager_proto_depIdxs = nil
}