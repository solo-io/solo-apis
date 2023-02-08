// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/internal/v2/pod_bounce_directive.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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

// Possible states in which an PodBounceDirective can exist.
type PodBounceDirectiveStatus_State int32

const (
	// The PodBounceDirective has yet to be picked up by the agent.
	PodBounceDirectiveStatus_PENDING PodBounceDirectiveStatus_State = 0
	// The agent has decided on which pods to bounce, and it's working on it.
	PodBounceDirectiveStatus_BOUNCING_PODS PodBounceDirectiveStatus_State = 1
	// Processing the pod bounce directive workflow failed.
	PodBounceDirectiveStatus_FAILED PodBounceDirectiveStatus_State = 3
	// Successfully bounced all pods
	PodBounceDirectiveStatus_FINISHED PodBounceDirectiveStatus_State = 4
)

// Enum value maps for PodBounceDirectiveStatus_State.
var (
	PodBounceDirectiveStatus_State_name = map[int32]string{
		0: "PENDING",
		1: "BOUNCING_PODS",
		3: "FAILED",
		4: "FINISHED",
	}
	PodBounceDirectiveStatus_State_value = map[string]int32{
		"PENDING":       0,
		"BOUNCING_PODS": 1,
		"FAILED":        3,
		"FINISHED":      4,
	}
)

func (x PodBounceDirectiveStatus_State) Enum() *PodBounceDirectiveStatus_State {
	p := new(PodBounceDirectiveStatus_State)
	*p = x
	return p
}

func (x PodBounceDirectiveStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PodBounceDirectiveStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_enumTypes[0].Descriptor()
}

func (PodBounceDirectiveStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_enumTypes[0]
}

func (x PodBounceDirectiveStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PodBounceDirectiveStatus_State.Descriptor instead.
func (PodBounceDirectiveStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescGZIP(), []int{1, 0}
}

// When certificates are issued, Istio-controlled pods need to be bounced (restarted) to ensure they pick up the
// new certificates due to [this issue](https://github.com/istio/istio/issues/22993).
// The certificate issuer will create a PodBounceDirective containing the namespaces and labels
// of the pods that need to be bounced in order to pick up the new certs.
type PodBounceDirectiveSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of Kubernetes pods to bounce (delete and cause a restart)
	// when the certificate is issued.
	// This will include the control plane pods as well as any Pods
	// which share a data plane with the target mesh.
	PodsToBounce []*PodBounceDirectiveSpec_PodSelector `protobuf:"bytes,1,rep,name=pods_to_bounce,json=podsToBounce,proto3" json:"pods_to_bounce,omitempty"`
	// Reference to the mesh on which this cert is being issued for
	MeshRef *v1.ObjectRef `protobuf:"bytes,2,opt,name=mesh_ref,json=meshRef,proto3" json:"mesh_ref,omitempty"`
}

func (x *PodBounceDirectiveSpec) Reset() {
	*x = PodBounceDirectiveSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodBounceDirectiveSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodBounceDirectiveSpec) ProtoMessage() {}

func (x *PodBounceDirectiveSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodBounceDirectiveSpec.ProtoReflect.Descriptor instead.
func (*PodBounceDirectiveSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescGZIP(), []int{0}
}

func (x *PodBounceDirectiveSpec) GetPodsToBounce() []*PodBounceDirectiveSpec_PodSelector {
	if x != nil {
		return x.PodsToBounce
	}
	return nil
}

func (x *PodBounceDirectiveSpec) GetMeshRef() *v1.ObjectRef {
	if x != nil {
		return x.MeshRef
	}
	return nil
}

// PodBounceDirectiveStatus reports the status for stateful Pod bounces (when bouncing pods requires waiting for readiness).
type PodBounceDirectiveStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the PodBounceDirective metadata.
	// If the `observedGeneration` does not match `metadata.generation`, the Gloo Mesh agent has not processed the most
	// recent version of this IssuedCertificate.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The current state of the IssuedCertificate workflow, reported by the agent.
	State PodBounceDirectiveStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=internal.gloo.solo.io.PodBounceDirectiveStatus_State" json:"state,omitempty"`
	Error string                         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// A list of Kubernetes pods to bounce (delete and cause a restart)
	// when the certificate is issued.
	// This will include the control plane pods as well as any Pods
	// which share a data plane with the target mesh.
	PodsBounced []*PodBounceDirectiveStatus_BouncedPodSet `protobuf:"bytes,4,rep,name=pods_bounced,json=podsBounced,proto3" json:"pods_bounced,omitempty"`
}

func (x *PodBounceDirectiveStatus) Reset() {
	*x = PodBounceDirectiveStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodBounceDirectiveStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodBounceDirectiveStatus) ProtoMessage() {}

func (x *PodBounceDirectiveStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodBounceDirectiveStatus.ProtoReflect.Descriptor instead.
func (*PodBounceDirectiveStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescGZIP(), []int{1}
}

func (x *PodBounceDirectiveStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *PodBounceDirectiveStatus) GetState() PodBounceDirectiveStatus_State {
	if x != nil {
		return x.State
	}
	return PodBounceDirectiveStatus_PENDING
}

func (x *PodBounceDirectiveStatus) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *PodBounceDirectiveStatus) GetPodsBounced() []*PodBounceDirectiveStatus_BouncedPodSet {
	if x != nil {
		return x.PodsBounced
	}
	return nil
}

// pods that will be restarted.
type PodBounceDirectiveSpec_PodSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace in which the pods live.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Any labels shared by the Pods.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Wait for this number of replacement pods to reach be fully ready before
	// deleting the next set of selected Pods.
	// This is used to ensure the control plane pods are allowed to restart
	// before sidecars and gateways are restarted.
	WaitForReplicas uint32 `protobuf:"varint,3,opt,name=wait_for_replicas,json=waitForReplicas,proto3" json:"wait_for_replicas,omitempty"`
	// Wait for the control plane to have synced all root cert configmaps in data plane namespaces before
	// bouncing these Pods.
	RootCertSync *PodBounceDirectiveSpec_PodSelector_RootCertSync `protobuf:"bytes,4,opt,name=root_cert_sync,json=rootCertSync,proto3" json:"root_cert_sync,omitempty"`
}

func (x *PodBounceDirectiveSpec_PodSelector) Reset() {
	*x = PodBounceDirectiveSpec_PodSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodBounceDirectiveSpec_PodSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodBounceDirectiveSpec_PodSelector) ProtoMessage() {}

func (x *PodBounceDirectiveSpec_PodSelector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodBounceDirectiveSpec_PodSelector.ProtoReflect.Descriptor instead.
func (*PodBounceDirectiveSpec_PodSelector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PodBounceDirectiveSpec_PodSelector) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PodBounceDirectiveSpec_PodSelector) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *PodBounceDirectiveSpec_PodSelector) GetWaitForReplicas() uint32 {
	if x != nil {
		return x.WaitForReplicas
	}
	return 0
}

func (x *PodBounceDirectiveSpec_PodSelector) GetRootCertSync() *PodBounceDirectiveSpec_PodSelector_RootCertSync {
	if x != nil {
		return x.RootCertSync
	}
	return nil
}

// RootCertSync describes values in a secret and configmap which must be equal in order for a Pod to be bounced.
type PodBounceDirectiveSpec_PodSelector_RootCertSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretRef    *v1.ObjectRef `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	SecretKey    string        `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	ConfigMapRef *v1.ObjectRef `protobuf:"bytes,3,opt,name=config_map_ref,json=configMapRef,proto3" json:"config_map_ref,omitempty"`
	ConfigMapKey string        `protobuf:"bytes,4,opt,name=config_map_key,json=configMapKey,proto3" json:"config_map_key,omitempty"`
}

func (x *PodBounceDirectiveSpec_PodSelector_RootCertSync) Reset() {
	*x = PodBounceDirectiveSpec_PodSelector_RootCertSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodBounceDirectiveSpec_PodSelector_RootCertSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodBounceDirectiveSpec_PodSelector_RootCertSync) ProtoMessage() {}

func (x *PodBounceDirectiveSpec_PodSelector_RootCertSync) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodBounceDirectiveSpec_PodSelector_RootCertSync.ProtoReflect.Descriptor instead.
func (*PodBounceDirectiveSpec_PodSelector_RootCertSync) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *PodBounceDirectiveSpec_PodSelector_RootCertSync) GetSecretRef() *v1.ObjectRef {
	if x != nil {
		return x.SecretRef
	}
	return nil
}

func (x *PodBounceDirectiveSpec_PodSelector_RootCertSync) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *PodBounceDirectiveSpec_PodSelector_RootCertSync) GetConfigMapRef() *v1.ObjectRef {
	if x != nil {
		return x.ConfigMapRef
	}
	return nil
}

func (x *PodBounceDirectiveSpec_PodSelector_RootCertSync) GetConfigMapKey() string {
	if x != nil {
		return x.ConfigMapKey
	}
	return ""
}

// A set of pods that were restarted.
type PodBounceDirectiveStatus_BouncedPodSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The names of the pods that were bounced for the corresponding selector specified in `PodBounceDirectiveSpec.PodSelector.labels`.
	BouncedPods []string `protobuf:"bytes,1,rep,name=bounced_pods,json=bouncedPods,proto3" json:"bounced_pods,omitempty"`
}

func (x *PodBounceDirectiveStatus_BouncedPodSet) Reset() {
	*x = PodBounceDirectiveStatus_BouncedPodSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodBounceDirectiveStatus_BouncedPodSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodBounceDirectiveStatus_BouncedPodSet) ProtoMessage() {}

func (x *PodBounceDirectiveStatus_BouncedPodSet) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodBounceDirectiveStatus_BouncedPodSet.ProtoReflect.Descriptor instead.
func (*PodBounceDirectiveStatus_BouncedPodSet) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PodBounceDirectiveStatus_BouncedPodSet) GetBouncedPods() []string {
	if x != nil {
		return x.BouncedPods
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6f, 0x64, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x05, 0x0a, 0x16,
	0x50, 0x6f, 0x64, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5f, 0x0a, 0x0e, 0x70, 0x6f, 0x64, 0x73, 0x5f, 0x74,
	0x6f, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x64, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x6f,
	0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0c, 0x70, 0x6f, 0x64, 0x73, 0x54,
	0x6f, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x68, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x68, 0x52, 0x65, 0x66,
	0x1a, 0xb6, 0x04, 0x0a, 0x0b, 0x50, 0x6f, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x5d,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x64, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x6f,
	0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2a, 0x0a,
	0x11, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x77, 0x61, 0x69, 0x74, 0x46, 0x6f,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x6c, 0x0a, 0x0e, 0x72, 0x6f, 0x6f,
	0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x46, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x64, 0x42, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x50, 0x6f, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x6f, 0x6f,
	0x74, 0x43, 0x65, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x0c, 0x72, 0x6f, 0x6f, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0xd4, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x65,
	0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x66,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x42, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x72, 0x65,
	0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70,
	0x52, 0x65, 0x66, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61,
	0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x22, 0x87, 0x03, 0x0a, 0x18, 0x50, 0x6f,
	0x64, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50,
	0x6f, 0x64, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x60, 0x0a, 0x0c, 0x70, 0x6f,
	0x64, 0x73, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x64, 0x42, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x50, 0x6f, 0x64, 0x53, 0x65, 0x74, 0x52,
	0x0b, 0x70, 0x6f, 0x64, 0x73, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x1a, 0x32, 0x0a, 0x0d,
	0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x50, 0x6f, 0x64, 0x53, 0x65, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x50, 0x6f, 0x64, 0x73,
	0x22, 0x41, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x4f, 0x55, 0x4e, 0x43, 0x49,
	0x4e, 0x47, 0x5f, 0x50, 0x4f, 0x44, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45,
	0x44, 0x10, 0x04, 0x42, 0x4b, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_goTypes = []interface{}{
	(PodBounceDirectiveStatus_State)(0),        // 0: internal.gloo.solo.io.PodBounceDirectiveStatus.State
	(*PodBounceDirectiveSpec)(nil),             // 1: internal.gloo.solo.io.PodBounceDirectiveSpec
	(*PodBounceDirectiveStatus)(nil),           // 2: internal.gloo.solo.io.PodBounceDirectiveStatus
	(*PodBounceDirectiveSpec_PodSelector)(nil), // 3: internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector
	nil, // 4: internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry
	(*PodBounceDirectiveSpec_PodSelector_RootCertSync)(nil), // 5: internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.RootCertSync
	(*PodBounceDirectiveStatus_BouncedPodSet)(nil),          // 6: internal.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet
	(*v1.ObjectRef)(nil), // 7: core.skv2.solo.io.ObjectRef
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_depIdxs = []int32{
	3, // 0: internal.gloo.solo.io.PodBounceDirectiveSpec.pods_to_bounce:type_name -> internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector
	7, // 1: internal.gloo.solo.io.PodBounceDirectiveSpec.mesh_ref:type_name -> core.skv2.solo.io.ObjectRef
	0, // 2: internal.gloo.solo.io.PodBounceDirectiveStatus.state:type_name -> internal.gloo.solo.io.PodBounceDirectiveStatus.State
	6, // 3: internal.gloo.solo.io.PodBounceDirectiveStatus.pods_bounced:type_name -> internal.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet
	4, // 4: internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.labels:type_name -> internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry
	5, // 5: internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.root_cert_sync:type_name -> internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.RootCertSync
	7, // 6: internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.RootCertSync.secret_ref:type_name -> core.skv2.solo.io.ObjectRef
	7, // 7: internal.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.RootCertSync.config_map_ref:type_name -> core.skv2.solo.io.ObjectRef
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodBounceDirectiveSpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodBounceDirectiveStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodBounceDirectiveSpec_PodSelector); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodBounceDirectiveSpec_PodSelector_RootCertSync); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodBounceDirectiveStatus_BouncedPodSet); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_internal_v2_pod_bounce_directive_proto_depIdxs = nil
}