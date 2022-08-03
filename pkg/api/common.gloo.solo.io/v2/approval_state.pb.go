// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/common/v2/approval_state.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// State of a Policy resource reflected in the status by Gloo Mesh while processing a resource.
type ApprovalState int32

const (
	// Resources are in a Pending state before they have been processed by Gloo Mesh.
	ApprovalState_PENDING ApprovalState = 0
	// Resources are in a Accepted state when they are valid and have been applied successfully to
	// the Gloo Mesh configuration.
	ApprovalState_ACCEPTED ApprovalState = 1
	// Resources are in an Invalid state when they contain incorrect configuration parameters,
	// such as missing required values or invalid resource references.
	// An invalid state can also result when a resource's configuration is valid
	// but conflicts with another resource which was accepted in an earlier point in time.
	ApprovalState_INVALID ApprovalState = 2
	// Resources are in a Warning state when they contain partially incorrect configuration parameters,
	// but has still been processed and applied by the system.
	ApprovalState_WARNING ApprovalState = 3
	// Resources are in a Failed state when they contain correct configuration parameters,
	// but the server encountered an error trying to synchronize the system to
	// the desired state.
	ApprovalState_FAILED ApprovalState = 4
	// The license key(s) in the application do not allow you to use this resource.
	ApprovalState_UNLICENSED ApprovalState = 5
)

// Enum value maps for ApprovalState.
var (
	ApprovalState_name = map[int32]string{
		0: "PENDING",
		1: "ACCEPTED",
		2: "INVALID",
		3: "WARNING",
		4: "FAILED",
		5: "UNLICENSED",
	}
	ApprovalState_value = map[string]int32{
		"PENDING":    0,
		"ACCEPTED":   1,
		"INVALID":    2,
		"WARNING":    3,
		"FAILED":     4,
		"UNLICENSED": 5,
	}
)

func (x ApprovalState) Enum() *ApprovalState {
	p := new(ApprovalState)
	*p = x
	return p
}

func (x ApprovalState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApprovalState) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_enumTypes[0].Descriptor()
}

func (ApprovalState) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_enumTypes[0]
}

func (x ApprovalState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApprovalState.Descriptor instead.
func (ApprovalState) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDescGZIP(), []int{0}
}

// the state of a resource across clusters
type ClusterState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the approval state of the resource in each cluster
	States map[string]ApprovalState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=common.gloo.solo.io.ApprovalState"`
}

func (x *ClusterState) Reset() {
	*x = ClusterState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterState) ProtoMessage() {}

func (x *ClusterState) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterState.ProtoReflect.Descriptor instead.
func (*ClusterState) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterState) GetStates() map[string]ApprovalState {
	if x != nil {
		return x.States
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDesc = []byte{
	0x0a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x5d, 0x0a, 0x0b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x60, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e,
	0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x44, 0x10, 0x05, 0x42, 0x49, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04,
	0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_goTypes = []interface{}{
	(ApprovalState)(0),   // 0: common.gloo.solo.io.ApprovalState
	(*ClusterState)(nil), // 1: common.gloo.solo.io.ClusterState
	nil,                  // 2: common.gloo.solo.io.ClusterState.StatesEntry
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_depIdxs = []int32{
	2, // 0: common.gloo.solo.io.ClusterState.states:type_name -> common.gloo.solo.io.ClusterState.StatesEntry
	0, // 1: common.gloo.solo.io.ClusterState.StatesEntry.value:type_name -> common.gloo.solo.io.ApprovalState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterState); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_common_v2_approval_state_proto_depIdxs = nil
}
