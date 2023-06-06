// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/common/v2/references.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// References the supported types for the Workload interface in pkg/utils/workloadutils/workload_types.go
type WorkloadReference_WorkloadKind int32

const (
	// Default value for unsupported workload kind
	WorkloadReference_UNRECOGNIZED_WORKLOAD WorkloadReference_WorkloadKind = 0
	WorkloadReference_DEPLOYMENT            WorkloadReference_WorkloadKind = 1
	WorkloadReference_DAEMON_SET            WorkloadReference_WorkloadKind = 2
	WorkloadReference_STATEFUL_SET          WorkloadReference_WorkloadKind = 3
	WorkloadReference_REPLICA_SET           WorkloadReference_WorkloadKind = 4
)

// Enum value maps for WorkloadReference_WorkloadKind.
var (
	WorkloadReference_WorkloadKind_name = map[int32]string{
		0: "UNRECOGNIZED_WORKLOAD",
		1: "DEPLOYMENT",
		2: "DAEMON_SET",
		3: "STATEFUL_SET",
		4: "REPLICA_SET",
	}
	WorkloadReference_WorkloadKind_value = map[string]int32{
		"UNRECOGNIZED_WORKLOAD": 0,
		"DEPLOYMENT":            1,
		"DAEMON_SET":            2,
		"STATEFUL_SET":          3,
		"REPLICA_SET":           4,
	}
)

func (x WorkloadReference_WorkloadKind) Enum() *WorkloadReference_WorkloadKind {
	p := new(WorkloadReference_WorkloadKind)
	*p = x
	return p
}

func (x WorkloadReference_WorkloadKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkloadReference_WorkloadKind) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_enumTypes[0].Descriptor()
}

func (WorkloadReference_WorkloadKind) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_enumTypes[0]
}

func (x WorkloadReference_WorkloadKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkloadReference_WorkloadKind.Descriptor instead.
func (WorkloadReference_WorkloadKind) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescGZIP(), []int{0, 0}
}

type WorkloadReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref  *ObjectReference               `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Kind WorkloadReference_WorkloadKind `protobuf:"varint,2,opt,name=kind,proto3,enum=common.gloo.solo.io.WorkloadReference_WorkloadKind" json:"kind,omitempty"`
}

func (x *WorkloadReference) Reset() {
	*x = WorkloadReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadReference) ProtoMessage() {}

func (x *WorkloadReference) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadReference.ProtoReflect.Descriptor instead.
func (*WorkloadReference) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescGZIP(), []int{0}
}

func (x *WorkloadReference) GetRef() *ObjectReference {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *WorkloadReference) GetKind() WorkloadReference_WorkloadKind {
	if x != nil {
		return x.Kind
	}
	return WorkloadReference_UNRECOGNIZED_WORKLOAD
}

// reference to a Kubernetes API object.
// Kube API objects are referenced explicitly by the namespace and cluster containing them.
type ObjectReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the name of the object
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the namespace of the object. if the field is omitted, Gloo Mesh will use the same namespace as the parent object containing this reference.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// the cluster of the object. if the field is omitted, Gloo Mesh will use the same cluster as the parent object containing this reference.
	Cluster string `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *ObjectReference) Reset() {
	*x = ObjectReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectReference) ProtoMessage() {}

func (x *ObjectReference) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectReference.ProtoReflect.Descriptor instead.
func (*ObjectReference) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectReference) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ObjectReference) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

type ObjectReferenceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refs []*ObjectReference `protobuf:"bytes,1,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (x *ObjectReferenceList) Reset() {
	*x = ObjectReferenceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectReferenceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectReferenceList) ProtoMessage() {}

func (x *ObjectReferenceList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectReferenceList.ProtoReflect.Descriptor instead.
func (*ObjectReferenceList) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectReferenceList) GetRefs() []*ObjectReference {
	if x != nil {
		return x.Refs
	}
	return nil
}

// Destinations are pointers to routable destinations for routes. Each destination should resolve to one and only one hostname.
// Destinations can refer to a variety of object types. The behavior of the route action
// will vary depending on the type of destination selected. Defaults to the kubernetes `v1/Service`.
// Currently supported destination types:
// - Service
// - VirtualDestination (route traffic to one of the VirtualDestination's backing services, based on the locality of the request)
// - ExternalService (route traffic to a static set of service endpoints external to the mesh)
// - CloudProvider (route traffic to a cloud provider, selected using the functionCall field)
type DestinationReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reference to the destination object by its metadata
	Ref *ObjectReference `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// the kind of destination being selected. defaults to Kubernetes Service.
	Kind DestinationKind `protobuf:"varint,2,opt,name=kind,proto3,enum=common.gloo.solo.io.DestinationKind" json:"kind,omitempty"`
	// the port on the destination object being targeted. required if the object provides more than one port.
	Port *PortSelector `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	// select a subset of the destination's endpoints for routing based on their labels.
	// Not applicable for a CloudProvider destination.
	Subset map[string]string `protobuf:"bytes,4,rep,name=subset,proto3" json:"subset,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Specify the proportion of traffic to be forwarded to this destination.
	// Weights across all of the `destinations` must sum to 100.
	// Weight is only relevant when used in the context of a route with multiple destinations.
	Weight uint32 `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// Specify how the destinations should be configured, for configuring:
	//
	//	Serverless functions
	//
	// Some destinations utilize options which require or permit additional configuration on routes targeting them.
	// aws lambda functions, for example, allow you to call a serverless lambda functions on aws. If the destination
	// config is required for the destination and not provided by the user, Gloo will invalidate the destination and its
	// parent resources.
	//
	// Types that are assignable to DestinationSpec:
	//
	//	*DestinationReference_Function
	DestinationSpec isDestinationReference_DestinationSpec `protobuf_oneof:"destination_spec"`
}

func (x *DestinationReference) Reset() {
	*x = DestinationReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationReference) ProtoMessage() {}

func (x *DestinationReference) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationReference.ProtoReflect.Descriptor instead.
func (*DestinationReference) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescGZIP(), []int{3}
}

func (x *DestinationReference) GetRef() *ObjectReference {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *DestinationReference) GetKind() DestinationKind {
	if x != nil {
		return x.Kind
	}
	return DestinationKind_SERVICE
}

func (x *DestinationReference) GetPort() *PortSelector {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *DestinationReference) GetSubset() map[string]string {
	if x != nil {
		return x.Subset
	}
	return nil
}

func (x *DestinationReference) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (m *DestinationReference) GetDestinationSpec() isDestinationReference_DestinationSpec {
	if m != nil {
		return m.DestinationSpec
	}
	return nil
}

func (x *DestinationReference) GetFunction() *FunctionDestinationSpec {
	if x, ok := x.GetDestinationSpec().(*DestinationReference_Function); ok {
		return x.Function
	}
	return nil
}

type isDestinationReference_DestinationSpec interface {
	isDestinationReference_DestinationSpec()
}

type DestinationReference_Function struct {
	Function *FunctionDestinationSpec `protobuf:"bytes,6,opt,name=function,proto3,oneof"`
}

func (*DestinationReference_Function) isDestinationReference_DestinationSpec() {}

// ListenerPortReference identifies a single listener in a VirtualGateway by port number
type ListenerPortReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The gateway containing the listener.
	GatewayRef *ObjectReference `protobuf:"bytes,1,opt,name=gateway_ref,json=gatewayRef,proto3" json:"gateway_ref,omitempty"`
	// The port of the listener on the gateway.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ListenerPortReference) Reset() {
	*x = ListenerPortReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerPortReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerPortReference) ProtoMessage() {}

func (x *ListenerPortReference) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerPortReference.ProtoReflect.Descriptor instead.
func (*ListenerPortReference) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescGZIP(), []int{4}
}

func (x *ListenerPortReference) GetGatewayRef() *ObjectReference {
	if x != nil {
		return x.GatewayRef
	}
	return nil
}

func (x *ListenerPortReference) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDesc = []byte{
	0x0a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x03, 0x72, 0x65, 0x66,
	0x12, 0x47, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x6c, 0x0a, 0x0c, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x4e, 0x52,
	0x45, 0x43, 0x4f, 0x47, 0x4e, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x4c, 0x4f,
	0x41, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x41, 0x45, 0x4d, 0x4f, 0x4e, 0x5f, 0x53,
	0x45, 0x54, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x46, 0x55, 0x4c,
	0x5f, 0x53, 0x45, 0x54, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x50, 0x4c, 0x49, 0x43,
	0x41, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x04, 0x22, 0x5d, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x13, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a,
	0x04, 0x72, 0x65, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x04, 0x72, 0x65, 0x66, 0x73, 0x22, 0xc1, 0x03, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x36, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x38, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4d, 0x0a, 0x06, 0x73, 0x75, 0x62,
	0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x4a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63,
	0x48, 0x00, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b,
	0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x22, 0x72, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x42,
	0x54, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04,
	0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_goTypes = []interface{}{
	(WorkloadReference_WorkloadKind)(0), // 0: common.gloo.solo.io.WorkloadReference.WorkloadKind
	(*WorkloadReference)(nil),           // 1: common.gloo.solo.io.WorkloadReference
	(*ObjectReference)(nil),             // 2: common.gloo.solo.io.ObjectReference
	(*ObjectReferenceList)(nil),         // 3: common.gloo.solo.io.ObjectReferenceList
	(*DestinationReference)(nil),        // 4: common.gloo.solo.io.DestinationReference
	(*ListenerPortReference)(nil),       // 5: common.gloo.solo.io.ListenerPortReference
	nil,                                 // 6: common.gloo.solo.io.DestinationReference.SubsetEntry
	(DestinationKind)(0),                // 7: common.gloo.solo.io.DestinationKind
	(*PortSelector)(nil),                // 8: common.gloo.solo.io.PortSelector
	(*FunctionDestinationSpec)(nil),     // 9: common.gloo.solo.io.FunctionDestinationSpec
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_depIdxs = []int32{
	2, // 0: common.gloo.solo.io.WorkloadReference.ref:type_name -> common.gloo.solo.io.ObjectReference
	0, // 1: common.gloo.solo.io.WorkloadReference.kind:type_name -> common.gloo.solo.io.WorkloadReference.WorkloadKind
	2, // 2: common.gloo.solo.io.ObjectReferenceList.refs:type_name -> common.gloo.solo.io.ObjectReference
	2, // 3: common.gloo.solo.io.DestinationReference.ref:type_name -> common.gloo.solo.io.ObjectReference
	7, // 4: common.gloo.solo.io.DestinationReference.kind:type_name -> common.gloo.solo.io.DestinationKind
	8, // 5: common.gloo.solo.io.DestinationReference.port:type_name -> common.gloo.solo.io.PortSelector
	6, // 6: common.gloo.solo.io.DestinationReference.subset:type_name -> common.gloo.solo.io.DestinationReference.SubsetEntry
	9, // 7: common.gloo.solo.io.DestinationReference.function:type_name -> common.gloo.solo.io.FunctionDestinationSpec
	2, // 8: common.gloo.solo.io.ListenerPortReference.gateway_ref:type_name -> common.gloo.solo.io.ObjectReference
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_port_proto_init()
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_selectors_proto_init()
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_cloud_provider_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadReference); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectReference); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectReferenceList); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationReference); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenerPortReference); i {
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
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*DestinationReference_Function)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_common_v2_references_proto_depIdxs = nil
}
