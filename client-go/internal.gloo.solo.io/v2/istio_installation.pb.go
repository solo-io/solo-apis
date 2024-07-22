// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/internal/v2/istio_installation.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	_ "k8s.io/api/core/v1"

	v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The current state of the installation.
type ClusterIstioInstallationStatus_State int32

const (
	// Not acted upon yet
	ClusterIstioInstallationStatus_PENDING ClusterIstioInstallationStatus_State = 0
	// Waiting for resources to be installed.
	ClusterIstioInstallationStatus_INSTALLING ClusterIstioInstallationStatus_State = 1
	// Waiting for upgrade to complete.
	ClusterIstioInstallationStatus_UPGRADING ClusterIstioInstallationStatus_State = 2
	// Waiting for resources to be uninstalled.
	ClusterIstioInstallationStatus_UNINSTALLING ClusterIstioInstallationStatus_State = 3
	// The installation or upgrade process is finished.
	ClusterIstioInstallationStatus_FINISHED ClusterIstioInstallationStatus_State = 4
	// The installation or upgrade process has failed.
	ClusterIstioInstallationStatus_FAILED ClusterIstioInstallationStatus_State = 5
)

// Enum value maps for ClusterIstioInstallationStatus_State.
var (
	ClusterIstioInstallationStatus_State_name = map[int32]string{
		0: "PENDING",
		1: "INSTALLING",
		2: "UPGRADING",
		3: "UNINSTALLING",
		4: "FINISHED",
		5: "FAILED",
	}
	ClusterIstioInstallationStatus_State_value = map[string]int32{
		"PENDING":      0,
		"INSTALLING":   1,
		"UPGRADING":    2,
		"UNINSTALLING": 3,
		"FINISHED":     4,
		"FAILED":       5,
	}
)

func (x ClusterIstioInstallationStatus_State) Enum() *ClusterIstioInstallationStatus_State {
	p := new(ClusterIstioInstallationStatus_State)
	*p = x
	return p
}

func (x ClusterIstioInstallationStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterIstioInstallationStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_enumTypes[0].Descriptor()
}

func (ClusterIstioInstallationStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_enumTypes[0]
}

func (x ClusterIstioInstallationStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterIstioInstallationStatus_State.Descriptor instead.
func (ClusterIstioInstallationStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescGZIP(), []int{1, 0}
}

// $hide_from_docs
type ClusterIstioInstallationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IstioOperator specification for the control plane.
	// For more info, see the [Istio documentation](https://istio.io/latest/docs/reference/config/istio.operator.v1alpha1/).
	IstioOperatorSpec *v2.IstioOperatorSpec `protobuf:"bytes,1,opt,name=istio_operator_spec,json=istioOperatorSpec,proto3" json:"istio_operator_spec,omitempty"`
	// Optional global configuration applicable to all installations
	HelmGlobal *v2.IstioLifecycleHelmGlobals `protobuf:"bytes,2,opt,name=helm_global,json=helmGlobal,proto3" json:"helm_global,omitempty"`
}

func (x *ClusterIstioInstallationSpec) Reset() {
	*x = ClusterIstioInstallationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterIstioInstallationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterIstioInstallationSpec) ProtoMessage() {}

func (x *ClusterIstioInstallationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterIstioInstallationSpec.ProtoReflect.Descriptor instead.
func (*ClusterIstioInstallationSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterIstioInstallationSpec) GetIstioOperatorSpec() *v2.IstioOperatorSpec {
	if x != nil {
		return x.IstioOperatorSpec
	}
	return nil
}

func (x *ClusterIstioInstallationSpec) GetHelmGlobal() *v2.IstioLifecycleHelmGlobals {
	if x != nil {
		return x.HelmGlobal
	}
	return nil
}

// $hide_from_docs
type ClusterIstioInstallationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current state of the Istio installation.
	State ClusterIstioInstallationStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=internal.gloo.solo.io.ClusterIstioInstallationStatus_State" json:"state,omitempty"`
	// A human readable message about the current state of the installation.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The observed revision of the Istio installation.
	ObservedRevision string `protobuf:"bytes,3,opt,name=observed_revision,json=observedRevision,proto3" json:"observed_revision,omitempty"`
	// The IstioOperator spec that is currently deployed for this revision.
	// TODO(birkland) may want helm values instead,
	ObservedOperator *v2.IstioOperatorSpec `protobuf:"bytes,4,opt,name=observed_operator,json=observedOperator,proto3" json:"observed_operator,omitempty"`
	// For helm-based installations, report the content and status of each
	Helm *ClusterIstioInstallationStatus_HelmStatus `protobuf:"bytes,5,opt,name=helm,proto3" json:"helm,omitempty"`
}

func (x *ClusterIstioInstallationStatus) Reset() {
	*x = ClusterIstioInstallationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterIstioInstallationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterIstioInstallationStatus) ProtoMessage() {}

func (x *ClusterIstioInstallationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterIstioInstallationStatus.ProtoReflect.Descriptor instead.
func (*ClusterIstioInstallationStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterIstioInstallationStatus) GetState() ClusterIstioInstallationStatus_State {
	if x != nil {
		return x.State
	}
	return ClusterIstioInstallationStatus_PENDING
}

func (x *ClusterIstioInstallationStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ClusterIstioInstallationStatus) GetObservedRevision() string {
	if x != nil {
		return x.ObservedRevision
	}
	return ""
}

func (x *ClusterIstioInstallationStatus) GetObservedOperator() *v2.IstioOperatorSpec {
	if x != nil {
		return x.ObservedOperator
	}
	return nil
}

func (x *ClusterIstioInstallationStatus) GetHelm() *ClusterIstioInstallationStatus_HelmStatus {
	if x != nil {
		return x.Helm
	}
	return nil
}

// $hide_from_docs
type ClusterIstioInstallationStatus_HelmStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Releases created or updated by this Istio installation
	UpgradeInstall []*ClusterIstioInstallationStatus_HelmReleaseGrouping `protobuf:"bytes,1,rep,name=upgrade_install,json=upgradeInstall,proto3" json:"upgrade_install,omitempty"`
	Uninstall      []*ClusterIstioInstallationStatus_HelmReleaseGrouping `protobuf:"bytes,2,rep,name=uninstall,proto3" json:"uninstall,omitempty"`
}

func (x *ClusterIstioInstallationStatus_HelmStatus) Reset() {
	*x = ClusterIstioInstallationStatus_HelmStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterIstioInstallationStatus_HelmStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterIstioInstallationStatus_HelmStatus) ProtoMessage() {}

func (x *ClusterIstioInstallationStatus_HelmStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterIstioInstallationStatus_HelmStatus.ProtoReflect.Descriptor instead.
func (*ClusterIstioInstallationStatus_HelmStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ClusterIstioInstallationStatus_HelmStatus) GetUpgradeInstall() []*ClusterIstioInstallationStatus_HelmReleaseGrouping {
	if x != nil {
		return x.UpgradeInstall
	}
	return nil
}

func (x *ClusterIstioInstallationStatus_HelmStatus) GetUninstall() []*ClusterIstioInstallationStatus_HelmReleaseGrouping {
	if x != nil {
		return x.Uninstall
	}
	return nil
}

// $hide_from_docs
type ClusterIstioInstallationStatus_HelmReleaseGrouping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Helm releases that share common values
	Releases []*ClusterIstioInstallationStatus_HelmRelease `protobuf:"bytes,1,rep,name=releases,proto3" json:"releases,omitempty"`
	// Values used for all releases in this batch
	//
	// +kubebuilder:validation:Type=object
	Values *structpb.Struct `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *ClusterIstioInstallationStatus_HelmReleaseGrouping) Reset() {
	*x = ClusterIstioInstallationStatus_HelmReleaseGrouping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterIstioInstallationStatus_HelmReleaseGrouping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterIstioInstallationStatus_HelmReleaseGrouping) ProtoMessage() {}

func (x *ClusterIstioInstallationStatus_HelmReleaseGrouping) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterIstioInstallationStatus_HelmReleaseGrouping.ProtoReflect.Descriptor instead.
func (*ClusterIstioInstallationStatus_HelmReleaseGrouping) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ClusterIstioInstallationStatus_HelmReleaseGrouping) GetReleases() []*ClusterIstioInstallationStatus_HelmRelease {
	if x != nil {
		return x.Releases
	}
	return nil
}

func (x *ClusterIstioInstallationStatus_HelmReleaseGrouping) GetValues() *structpb.Struct {
	if x != nil {
		return x.Values
	}
	return nil
}

// $hide_from_docs
type ClusterIstioInstallationStatus_HelmRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Helm release name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Namespace where helm release resides
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Helm release revision once applying this istio instalation is finished
	HelmRevision int32 `protobuf:"varint,3,opt,name=helm_revision,json=helmRevision,proto3" json:"helm_revision,omitempty"`
	// URL to the chart used for this release (typically an OCI repo)
	ChartUrl string `protobuf:"bytes,4,opt,name=chart_url,json=chartUrl,proto3" json:"chart_url,omitempty"`
	// Helm chart version for this release
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	// State of an individual helm release
	State ClusterIstioInstallationStatus_State `protobuf:"varint,6,opt,name=state,proto3,enum=internal.gloo.solo.io.ClusterIstioInstallationStatus_State" json:"state,omitempty"`
}

func (x *ClusterIstioInstallationStatus_HelmRelease) Reset() {
	*x = ClusterIstioInstallationStatus_HelmRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterIstioInstallationStatus_HelmRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterIstioInstallationStatus_HelmRelease) ProtoMessage() {}

func (x *ClusterIstioInstallationStatus_HelmRelease) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterIstioInstallationStatus_HelmRelease.ProtoReflect.Descriptor instead.
func (*ClusterIstioInstallationStatus_HelmRelease) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescGZIP(), []int{1, 2}
}

func (x *ClusterIstioInstallationStatus_HelmRelease) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterIstioInstallationStatus_HelmRelease) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ClusterIstioInstallationStatus_HelmRelease) GetHelmRevision() int32 {
	if x != nil {
		return x.HelmRevision
	}
	return 0
}

func (x *ClusterIstioInstallationStatus_HelmRelease) GetChartUrl() string {
	if x != nil {
		return x.ChartUrl
	}
	return ""
}

func (x *ClusterIstioInstallationStatus_HelmRelease) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ClusterIstioInstallationStatus_HelmRelease) GetState() ClusterIstioInstallationStatus_State {
	if x != nil {
		return x.State
	}
	return ClusterIstioInstallationStatus_PENDING
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDesc = []byte{
	0x0a, 0x60, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6b, 0x38, 0x73,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x56, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x32, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x68, 0x65, 0x6c, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x1c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x73, 0x74, 0x69, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x56, 0x0a, 0x13, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x11, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4f, 0x0a, 0x0b,
	0x68, 0x65, 0x6c, 0x6d, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x65, 0x6c, 0x6d, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x73, 0x52, 0x0a, 0x68, 0x65, 0x6c, 0x6d, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x22, 0xcb, 0x08,
	0x0a, 0x1e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x51, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x3b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x73, 0x74, 0x69, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x11, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69,
	0x6f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x10, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x54, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x73, 0x74,
	0x69, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x04, 0x68, 0x65, 0x6c, 0x6d, 0x1a, 0xe9, 0x01, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x72, 0x0a, 0x0f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x73, 0x74,
	0x69, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x67, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x73, 0x74, 0x69, 0x6f,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x1a, 0xa5, 0x01, 0x0a, 0x13, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x5d, 0x0a, 0x08, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x73, 0x74, 0x69, 0x6f,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0xee, 0x01, 0x0a, 0x0b, 0x48, 0x65,
	0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x68,
	0x65, 0x6c, 0x6d, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x68, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5f, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10,
	0x03, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x42, 0x59, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5,
	0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_goTypes = []interface{}{
	(ClusterIstioInstallationStatus_State)(0),                  // 0: internal.gloo.solo.io.ClusterIstioInstallationStatus.State
	(*ClusterIstioInstallationSpec)(nil),                       // 1: internal.gloo.solo.io.ClusterIstioInstallationSpec
	(*ClusterIstioInstallationStatus)(nil),                     // 2: internal.gloo.solo.io.ClusterIstioInstallationStatus
	(*ClusterIstioInstallationStatus_HelmStatus)(nil),          // 3: internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmStatus
	(*ClusterIstioInstallationStatus_HelmReleaseGrouping)(nil), // 4: internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmReleaseGrouping
	(*ClusterIstioInstallationStatus_HelmRelease)(nil),         // 5: internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmRelease
	(*v2.IstioOperatorSpec)(nil),                               // 6: common.gloo.solo.io.IstioOperatorSpec
	(*v2.IstioLifecycleHelmGlobals)(nil),                       // 7: common.gloo.solo.io.IstioLifecycleHelmGlobals
	(*structpb.Struct)(nil),                                    // 8: google.protobuf.Struct
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_depIdxs = []int32{
	6,  // 0: internal.gloo.solo.io.ClusterIstioInstallationSpec.istio_operator_spec:type_name -> common.gloo.solo.io.IstioOperatorSpec
	7,  // 1: internal.gloo.solo.io.ClusterIstioInstallationSpec.helm_global:type_name -> common.gloo.solo.io.IstioLifecycleHelmGlobals
	0,  // 2: internal.gloo.solo.io.ClusterIstioInstallationStatus.state:type_name -> internal.gloo.solo.io.ClusterIstioInstallationStatus.State
	6,  // 3: internal.gloo.solo.io.ClusterIstioInstallationStatus.observed_operator:type_name -> common.gloo.solo.io.IstioOperatorSpec
	3,  // 4: internal.gloo.solo.io.ClusterIstioInstallationStatus.helm:type_name -> internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmStatus
	4,  // 5: internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmStatus.upgrade_install:type_name -> internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmReleaseGrouping
	4,  // 6: internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmStatus.uninstall:type_name -> internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmReleaseGrouping
	5,  // 7: internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmReleaseGrouping.releases:type_name -> internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmRelease
	8,  // 8: internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmReleaseGrouping.values:type_name -> google.protobuf.Struct
	0,  // 9: internal.gloo.solo.io.ClusterIstioInstallationStatus.HelmRelease.state:type_name -> internal.gloo.solo.io.ClusterIstioInstallationStatus.State
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterIstioInstallationSpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterIstioInstallationStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterIstioInstallationStatus_HelmStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterIstioInstallationStatus_HelmReleaseGrouping); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterIstioInstallationStatus_HelmRelease); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_internal_v2_istio_installation_proto_depIdxs = nil
}
