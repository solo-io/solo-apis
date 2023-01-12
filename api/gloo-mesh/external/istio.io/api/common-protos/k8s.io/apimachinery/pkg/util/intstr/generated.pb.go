//
//Copyright The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file was autogenerated by go-to-protobuf. Do not edit it manually!

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/external/istio.io/api/common-protos/k8s.io/apimachinery/pkg/util/intstr/generated.proto

package intstr

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// IntOrString is a type that can hold an int32 or a string.  When used in
// JSON or YAML marshalling and unmarshalling, it produces or consumes the
// inner type.  This allows you to have, for example, a JSON field that can
// accept a name or number.
// TODO: Rename to Int32OrString
//
// +protobuf=true
// +protobuf.options.(gogoproto.goproto_stringer)=false
// +k8s:openapi-gen=true
type IntOrString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   *int64  `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	IntVal *int32  `protobuf:"varint,2,opt,name=intVal" json:"intVal,omitempty"`
	StrVal *string `protobuf:"bytes,3,opt,name=strVal" json:"strVal,omitempty"`
}

func (x *IntOrString) Reset() {
	*x = IntOrString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntOrString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntOrString) ProtoMessage() {}

func (x *IntOrString) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntOrString.ProtoReflect.Descriptor instead.
func (*IntOrString) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDescGZIP(), []int{0}
}

func (x *IntOrString) GetType() int64 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *IntOrString) GetIntVal() int32 {
	if x != nil && x.IntVal != nil {
		return *x.IntVal
	}
	return 0
}

func (x *IntOrString) GetStrVal() string {
	if x != nil && x.StrVal != nil {
		return *x.StrVal
	}
	return ""
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDesc = []byte{
	0x0a, 0x82, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x69,
	0x6e, 0x74, 0x73, 0x74, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x75,
	0x74, 0x69, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x22, 0x51, 0x0a, 0x0b, 0x49, 0x6e,
	0x74, 0x4f, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x42, 0x08, 0x5a,
	0x06, 0x69, 0x6e, 0x74, 0x73, 0x74, 0x72,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_goTypes = []interface{}{
	(*IntOrString)(nil), // 0: k8s.io.apimachinery.pkg.util.intstr.IntOrString
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntOrString); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_util_intstr_generated_proto_depIdxs = nil
}
