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
// source: github.com/solo-io/solo-apis/api/gloo-mesh/external/istio.io/api/common-protos/k8s.io/apimachinery/pkg/runtime/generated.proto

package runtime

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

// RawExtension is used to hold extensions in external versions.
//
// To use this, make a field which has RawExtension as its type in your external, versioned
// struct, and Object in your internal struct. You also need to register your
// various plugin types.
//
// // Internal package:
//
//	type MyAPIObject struct {
//		runtime.TypeMeta `json:",inline"`
//		MyPlugin runtime.Object `json:"myPlugin"`
//	}
//
//	type PluginA struct {
//		AOption string `json:"aOption"`
//	}
//
// // External package:
//
//	type MyAPIObject struct {
//		runtime.TypeMeta `json:",inline"`
//		MyPlugin runtime.RawExtension `json:"myPlugin"`
//	}
//
//	type PluginA struct {
//		AOption string `json:"aOption"`
//	}
//
// // On the wire, the JSON will look something like this:
//
//	{
//		"kind":"MyAPIObject",
//		"apiVersion":"v1",
//		"myPlugin": {
//			"kind":"PluginA",
//			"aOption":"foo",
//		},
//	}
//
// So what happens? Decode first uses json or yaml to unmarshal the serialized data into
// your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked.
// The next step is to copy (using pkg/conversion) into the internal struct. The runtime
// package's DefaultScheme has conversion functions installed which will unpack the
// JSON stored in RawExtension, turning it into the correct object type, and storing it
// in the Object. (TODO: In the case where the object is of an unknown type, a
// runtime.Unknown object will be created and stored.)
//
// +k8s:deepcopy-gen=true
// +protobuf=true
// +k8s:openapi-gen=true
type RawExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Raw is the underlying serialization of this object.
	//
	// TODO: Determine how to detect ContentType and ContentEncoding of 'Raw' data.
	Raw []byte `protobuf:"bytes,1,opt,name=raw" json:"raw,omitempty"`
}

func (x *RawExtension) Reset() {
	*x = RawExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawExtension) ProtoMessage() {}

func (x *RawExtension) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawExtension.ProtoReflect.Descriptor instead.
func (*RawExtension) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescGZIP(), []int{0}
}

func (x *RawExtension) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

// TypeMeta is shared by all top level objects. The proper way to use it is to inline it in your type,
// like this:
//
//	type MyAwesomeAPIObject struct {
//	     runtime.TypeMeta    `json:",inline"`
//	     ... // other fields
//	}
//
// func (obj *MyAwesomeAPIObject) SetGroupVersionKind(gvk *metav1.GroupVersionKind) { metav1.UpdateTypeMeta(obj,gvk) }; GroupVersionKind() *GroupVersionKind
//
// TypeMeta is provided here for convenience. You may use it directly from this package or define
// your own with the same fields.
//
// +k8s:deepcopy-gen=false
// +protobuf=true
// +k8s:openapi-gen=true
type TypeMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	ApiVersion *string `protobuf:"bytes,1,opt,name=apiVersion" json:"apiVersion,omitempty"`
	// +optional
	Kind *string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
}

func (x *TypeMeta) Reset() {
	*x = TypeMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeMeta) ProtoMessage() {}

func (x *TypeMeta) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeMeta.ProtoReflect.Descriptor instead.
func (*TypeMeta) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescGZIP(), []int{1}
}

func (x *TypeMeta) GetApiVersion() string {
	if x != nil && x.ApiVersion != nil {
		return *x.ApiVersion
	}
	return ""
}

func (x *TypeMeta) GetKind() string {
	if x != nil && x.Kind != nil {
		return *x.Kind
	}
	return ""
}

// Unknown allows api objects with unknown types to be passed-through. This can be used
// to deal with the API objects from a plug-in. Unknown objects still have functioning
// TypeMeta features-- kind, version, etc.
// TODO: Make this object have easy access to field based accessors and settors for
// metadata and field mutatation.
//
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +protobuf=true
// +k8s:openapi-gen=true
type Unknown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeMeta *TypeMeta `protobuf:"bytes,1,opt,name=typeMeta" json:"typeMeta,omitempty"`
	// Raw will hold the complete serialized object which couldn't be matched
	// with a registered type. Most likely, nothing should be done with this
	// except for passing it through the system.
	Raw []byte `protobuf:"bytes,2,opt,name=raw" json:"raw,omitempty"`
	// ContentEncoding is encoding used to encode 'Raw' data.
	// Unspecified means no encoding.
	ContentEncoding *string `protobuf:"bytes,3,opt,name=contentEncoding" json:"contentEncoding,omitempty"`
	// ContentType  is serialization method used to serialize 'Raw'.
	// Unspecified means ContentTypeJSON.
	ContentType *string `protobuf:"bytes,4,opt,name=contentType" json:"contentType,omitempty"`
}

func (x *Unknown) Reset() {
	*x = Unknown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unknown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unknown) ProtoMessage() {}

func (x *Unknown) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unknown.ProtoReflect.Descriptor instead.
func (*Unknown) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescGZIP(), []int{2}
}

func (x *Unknown) GetTypeMeta() *TypeMeta {
	if x != nil {
		return x.TypeMeta
	}
	return nil
}

func (x *Unknown) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *Unknown) GetContentEncoding() string {
	if x != nil && x.ContentEncoding != nil {
		return *x.ContentEncoding
	}
	return ""
}

func (x *Unknown) GetContentType() string {
	if x != nil && x.ContentType != nil {
		return *x.ContentType
	}
	return ""
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDesc = []byte{
	0x0a, 0x7e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0x20, 0x0a, 0x0c, 0x52, 0x61, 0x77, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x72, 0x61, 0x77, 0x22, 0x3e, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12,
	0x45, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_goTypes = []interface{}{
	(*RawExtension)(nil), // 0: k8s.io.apimachinery.pkg.runtime.RawExtension
	(*TypeMeta)(nil),     // 1: k8s.io.apimachinery.pkg.runtime.TypeMeta
	(*Unknown)(nil),      // 2: k8s.io.apimachinery.pkg.runtime.Unknown
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_depIdxs = []int32{
	1, // 0: k8s.io.apimachinery.pkg.runtime.Unknown.typeMeta:type_name -> k8s.io.apimachinery.pkg.runtime.TypeMeta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawExtension); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeMeta); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unknown); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_external_istio_io_api_common_protos_k8s_io_apimachinery_pkg_runtime_generated_proto_depIdxs = nil
}
