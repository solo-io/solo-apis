// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/migrate.proto

package annotations

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MigrateAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rename the message/enum/enum value in next version.
	Rename string `protobuf:"bytes,1,opt,name=rename,proto3" json:"rename,omitempty"`
}

func (x *MigrateAnnotation) Reset() {
	*x = MigrateAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrateAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrateAnnotation) ProtoMessage() {}

func (x *MigrateAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrateAnnotation.ProtoReflect.Descriptor instead.
func (*MigrateAnnotation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescGZIP(), []int{0}
}

func (x *MigrateAnnotation) GetRename() string {
	if x != nil {
		return x.Rename
	}
	return ""
}

type FieldMigrateAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rename the field in next version.
	Rename string `protobuf:"bytes,1,opt,name=rename,proto3" json:"rename,omitempty"`
	// Add the field to a named oneof in next version. If this already exists, the
	// field will join its siblings under the oneof, otherwise a new oneof will be
	// created with the given name.
	OneofPromotion string `protobuf:"bytes,2,opt,name=oneof_promotion,json=oneofPromotion,proto3" json:"oneof_promotion,omitempty"`
}

func (x *FieldMigrateAnnotation) Reset() {
	*x = FieldMigrateAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMigrateAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMigrateAnnotation) ProtoMessage() {}

func (x *FieldMigrateAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMigrateAnnotation.ProtoReflect.Descriptor instead.
func (*FieldMigrateAnnotation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescGZIP(), []int{1}
}

func (x *FieldMigrateAnnotation) GetRename() string {
	if x != nil {
		return x.Rename
	}
	return ""
}

func (x *FieldMigrateAnnotation) GetOneofPromotion() string {
	if x != nil {
		return x.OneofPromotion
	}
	return ""
}

type FileMigrateAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Move all types in the file to another package, this implies changing proto
	// file path.
	MoveToPackage string `protobuf:"bytes,2,opt,name=move_to_package,json=moveToPackage,proto3" json:"move_to_package,omitempty"`
}

func (x *FileMigrateAnnotation) Reset() {
	*x = FileMigrateAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMigrateAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMigrateAnnotation) ProtoMessage() {}

func (x *FileMigrateAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMigrateAnnotation.ProtoReflect.Descriptor instead.
func (*FileMigrateAnnotation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescGZIP(), []int{2}
}

func (x *FileMigrateAnnotation) GetMoveToPackage() string {
	if x != nil {
		return x.MoveToPackage
	}
	return ""
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*MigrateAnnotation)(nil),
		Field:         146192544,
		Name:          "solo.io.udpa.annotations.message_migrate",
		Tag:           "bytes,146192544,opt,name=message_migrate",
		Filename:      "github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/migrate.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldMigrateAnnotation)(nil),
		Field:         146192544,
		Name:          "solo.io.udpa.annotations.field_migrate",
		Tag:           "bytes,146192544,opt,name=field_migrate",
		Filename:      "github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/migrate.proto",
	},
	{
		ExtendedType:  (*descriptor.EnumOptions)(nil),
		ExtensionType: (*MigrateAnnotation)(nil),
		Field:         146192544,
		Name:          "solo.io.udpa.annotations.enum_migrate",
		Tag:           "bytes,146192544,opt,name=enum_migrate",
		Filename:      "github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/migrate.proto",
	},
	{
		ExtendedType:  (*descriptor.EnumValueOptions)(nil),
		ExtensionType: (*MigrateAnnotation)(nil),
		Field:         146192544,
		Name:          "solo.io.udpa.annotations.enum_value_migrate",
		Tag:           "bytes,146192544,opt,name=enum_value_migrate",
		Filename:      "github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/migrate.proto",
	},
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*FileMigrateAnnotation)(nil),
		Field:         146192544,
		Name:          "solo.io.udpa.annotations.file_migrate",
		Tag:           "bytes,146192544,opt,name=file_migrate",
		Filename:      "github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/migrate.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional solo.io.udpa.annotations.MigrateAnnotation message_migrate = 146192544;
	E_MessageMigrate = &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_extTypes[0]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional solo.io.udpa.annotations.FieldMigrateAnnotation field_migrate = 146192544;
	E_FieldMigrate = &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_extTypes[1]
)

// Extension fields to descriptor.EnumOptions.
var (
	// optional solo.io.udpa.annotations.MigrateAnnotation enum_migrate = 146192544;
	E_EnumMigrate = &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_extTypes[2]
)

// Extension fields to descriptor.EnumValueOptions.
var (
	// optional solo.io.udpa.annotations.MigrateAnnotation enum_value_migrate = 146192544;
	E_EnumValueMigrate = &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_extTypes[3]
)

// Extension fields to descriptor.FileOptions.
var (
	// optional solo.io.udpa.annotations.FileMigrateAnnotation file_migrate = 146192544;
	E_FileMigrate = &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_extTypes[4]
)

var File_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDesc = []byte{
	0x0a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x75, 0x64,
	0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2b, 0x0a, 0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a,
	0x16, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x6f, 0x76, 0x65,
	0x54, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x3a, 0x78, 0x0a, 0x0f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa0, 0xf1,
	0xda, 0x45, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x65, 0x3a, 0x77, 0x0a, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xa0, 0xf1, 0xda, 0x45, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x3a, 0x6f, 0x0a, 0x0c,
	0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa0, 0xf1, 0xda, 0x45, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x75, 0x64,
	0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x3a, 0x7f, 0x0a,
	0x12, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa0, 0xf1, 0xda, 0x45, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x65, 0x6e,
	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x3a, 0x73,
	0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa0, 0xf1, 0xda,
	0x45, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x75, 0x64, 0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_goTypes = []interface{}{
	(*MigrateAnnotation)(nil),           // 0: solo.io.udpa.annotations.MigrateAnnotation
	(*FieldMigrateAnnotation)(nil),      // 1: solo.io.udpa.annotations.FieldMigrateAnnotation
	(*FileMigrateAnnotation)(nil),       // 2: solo.io.udpa.annotations.FileMigrateAnnotation
	(*descriptor.MessageOptions)(nil),   // 3: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),     // 4: google.protobuf.FieldOptions
	(*descriptor.EnumOptions)(nil),      // 5: google.protobuf.EnumOptions
	(*descriptor.EnumValueOptions)(nil), // 6: google.protobuf.EnumValueOptions
	(*descriptor.FileOptions)(nil),      // 7: google.protobuf.FileOptions
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_depIdxs = []int32{
	3,  // 0: solo.io.udpa.annotations.message_migrate:extendee -> google.protobuf.MessageOptions
	4,  // 1: solo.io.udpa.annotations.field_migrate:extendee -> google.protobuf.FieldOptions
	5,  // 2: solo.io.udpa.annotations.enum_migrate:extendee -> google.protobuf.EnumOptions
	6,  // 3: solo.io.udpa.annotations.enum_value_migrate:extendee -> google.protobuf.EnumValueOptions
	7,  // 4: solo.io.udpa.annotations.file_migrate:extendee -> google.protobuf.FileOptions
	0,  // 5: solo.io.udpa.annotations.message_migrate:type_name -> solo.io.udpa.annotations.MigrateAnnotation
	1,  // 6: solo.io.udpa.annotations.field_migrate:type_name -> solo.io.udpa.annotations.FieldMigrateAnnotation
	0,  // 7: solo.io.udpa.annotations.enum_migrate:type_name -> solo.io.udpa.annotations.MigrateAnnotation
	0,  // 8: solo.io.udpa.annotations.enum_value_migrate:type_name -> solo.io.udpa.annotations.MigrateAnnotation
	2,  // 9: solo.io.udpa.annotations.file_migrate:type_name -> solo.io.udpa.annotations.FileMigrateAnnotation
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	5,  // [5:10] is the sub-list for extension type_name
	0,  // [0:5] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrateAnnotation); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMigrateAnnotation); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMigrateAnnotation); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_msgTypes,
		ExtensionInfos:    file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_extTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_udpa_annotations_migrate_proto_depIdxs = nil
}
