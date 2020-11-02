// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/skv2-enterprise/v1alpha1/multicluster.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MultiClusterRoleSpec_Rule_Action int32

const (
	MultiClusterRoleSpec_Rule_ANY    MultiClusterRoleSpec_Rule_Action = 0
	MultiClusterRoleSpec_Rule_CREATE MultiClusterRoleSpec_Rule_Action = 1
	MultiClusterRoleSpec_Rule_UPDATE MultiClusterRoleSpec_Rule_Action = 2
	MultiClusterRoleSpec_Rule_DELETE MultiClusterRoleSpec_Rule_Action = 3
)

var MultiClusterRoleSpec_Rule_Action_name = map[int32]string{
	0: "ANY",
	1: "CREATE",
	2: "UPDATE",
	3: "DELETE",
}

var MultiClusterRoleSpec_Rule_Action_value = map[string]int32{
	"ANY":    0,
	"CREATE": 1,
	"UPDATE": 2,
	"DELETE": 3,
}

func (x MultiClusterRoleSpec_Rule_Action) String() string {
	return proto.EnumName(MultiClusterRoleSpec_Rule_Action_name, int32(x))
}

func (MultiClusterRoleSpec_Rule_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_538363f2c905c978, []int{1, 0, 0}
}

//
//Object representing the clusters and namespaces on which resources are created/updated/deleted,
//computed as the cartesian product of all declared namespace and cluster values.
type Placement struct {
	//
	//List of namespaces within each placement cluster in which to create/update/delete resources.
	//Wildcard ("*") represents any namespace.
	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	//
	//List of clusters (represented by a string) in which to create/update/delete resources.
	//Wildcard ("*") represents any cluster.
	Clusters             []string `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Placement) Reset()         { *m = Placement{} }
func (m *Placement) String() string { return proto.CompactTextString(m) }
func (*Placement) ProtoMessage()    {}
func (*Placement) Descriptor() ([]byte, []int) {
	return fileDescriptor_538363f2c905c978, []int{0}
}
func (m *Placement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Placement.Unmarshal(m, b)
}
func (m *Placement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Placement.Marshal(b, m, deterministic)
}
func (m *Placement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Placement.Merge(m, src)
}
func (m *Placement) XXX_Size() int {
	return xxx_messageInfo_Placement.Size(m)
}
func (m *Placement) XXX_DiscardUnknown() {
	xxx_messageInfo_Placement.DiscardUnknown(m)
}

var xxx_messageInfo_Placement proto.InternalMessageInfo

func (m *Placement) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *Placement) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type MultiClusterRoleSpec struct {
	Rules                []*MultiClusterRoleSpec_Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MultiClusterRoleSpec) Reset()         { *m = MultiClusterRoleSpec{} }
func (m *MultiClusterRoleSpec) String() string { return proto.CompactTextString(m) }
func (*MultiClusterRoleSpec) ProtoMessage()    {}
func (*MultiClusterRoleSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_538363f2c905c978, []int{1}
}
func (m *MultiClusterRoleSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiClusterRoleSpec.Unmarshal(m, b)
}
func (m *MultiClusterRoleSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiClusterRoleSpec.Marshal(b, m, deterministic)
}
func (m *MultiClusterRoleSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiClusterRoleSpec.Merge(m, src)
}
func (m *MultiClusterRoleSpec) XXX_Size() int {
	return xxx_messageInfo_MultiClusterRoleSpec.Size(m)
}
func (m *MultiClusterRoleSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiClusterRoleSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MultiClusterRoleSpec proto.InternalMessageInfo

func (m *MultiClusterRoleSpec) GetRules() []*MultiClusterRoleSpec_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type MultiClusterRoleSpec_Rule struct {
	ApiGroup string `protobuf:"bytes,1,opt,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	// The kind of the object to apply to, if left empty will apply to all kinds in a group.
	Kind   *types.StringValue               `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Action MultiClusterRoleSpec_Rule_Action `protobuf:"varint,3,opt,name=action,proto3,enum=multicluster.solo.io.MultiClusterRoleSpec_Rule_Action" json:"action,omitempty"`
	// The clusters and namespaces this rule applies to.
	Placements           []*Placement `protobuf:"bytes,4,rep,name=placements,proto3" json:"placements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MultiClusterRoleSpec_Rule) Reset()         { *m = MultiClusterRoleSpec_Rule{} }
func (m *MultiClusterRoleSpec_Rule) String() string { return proto.CompactTextString(m) }
func (*MultiClusterRoleSpec_Rule) ProtoMessage()    {}
func (*MultiClusterRoleSpec_Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_538363f2c905c978, []int{1, 0}
}
func (m *MultiClusterRoleSpec_Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiClusterRoleSpec_Rule.Unmarshal(m, b)
}
func (m *MultiClusterRoleSpec_Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiClusterRoleSpec_Rule.Marshal(b, m, deterministic)
}
func (m *MultiClusterRoleSpec_Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiClusterRoleSpec_Rule.Merge(m, src)
}
func (m *MultiClusterRoleSpec_Rule) XXX_Size() int {
	return xxx_messageInfo_MultiClusterRoleSpec_Rule.Size(m)
}
func (m *MultiClusterRoleSpec_Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiClusterRoleSpec_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_MultiClusterRoleSpec_Rule proto.InternalMessageInfo

func (m *MultiClusterRoleSpec_Rule) GetApiGroup() string {
	if m != nil {
		return m.ApiGroup
	}
	return ""
}

func (m *MultiClusterRoleSpec_Rule) GetKind() *types.StringValue {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *MultiClusterRoleSpec_Rule) GetAction() MultiClusterRoleSpec_Rule_Action {
	if m != nil {
		return m.Action
	}
	return MultiClusterRoleSpec_Rule_ANY
}

func (m *MultiClusterRoleSpec_Rule) GetPlacements() []*Placement {
	if m != nil {
		return m.Placements
	}
	return nil
}

type MultiClusterRoleStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiClusterRoleStatus) Reset()         { *m = MultiClusterRoleStatus{} }
func (m *MultiClusterRoleStatus) String() string { return proto.CompactTextString(m) }
func (*MultiClusterRoleStatus) ProtoMessage()    {}
func (*MultiClusterRoleStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_538363f2c905c978, []int{2}
}
func (m *MultiClusterRoleStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiClusterRoleStatus.Unmarshal(m, b)
}
func (m *MultiClusterRoleStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiClusterRoleStatus.Marshal(b, m, deterministic)
}
func (m *MultiClusterRoleStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiClusterRoleStatus.Merge(m, src)
}
func (m *MultiClusterRoleStatus) XXX_Size() int {
	return xxx_messageInfo_MultiClusterRoleStatus.Size(m)
}
func (m *MultiClusterRoleStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiClusterRoleStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MultiClusterRoleStatus proto.InternalMessageInfo

type MultiClusterRoleBindingSpec struct {
	// reference to users or groups to apply the MultiClusterRole to
	Subjects []*v1.TypedObjectRef `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty"`
	// reference to a MultiClusterRole
	RoleRef              *v1.ObjectRef `protobuf:"bytes,2,opt,name=role_ref,json=roleRef,proto3" json:"role_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MultiClusterRoleBindingSpec) Reset()         { *m = MultiClusterRoleBindingSpec{} }
func (m *MultiClusterRoleBindingSpec) String() string { return proto.CompactTextString(m) }
func (*MultiClusterRoleBindingSpec) ProtoMessage()    {}
func (*MultiClusterRoleBindingSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_538363f2c905c978, []int{3}
}
func (m *MultiClusterRoleBindingSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiClusterRoleBindingSpec.Unmarshal(m, b)
}
func (m *MultiClusterRoleBindingSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiClusterRoleBindingSpec.Marshal(b, m, deterministic)
}
func (m *MultiClusterRoleBindingSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiClusterRoleBindingSpec.Merge(m, src)
}
func (m *MultiClusterRoleBindingSpec) XXX_Size() int {
	return xxx_messageInfo_MultiClusterRoleBindingSpec.Size(m)
}
func (m *MultiClusterRoleBindingSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiClusterRoleBindingSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MultiClusterRoleBindingSpec proto.InternalMessageInfo

func (m *MultiClusterRoleBindingSpec) GetSubjects() []*v1.TypedObjectRef {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *MultiClusterRoleBindingSpec) GetRoleRef() *v1.ObjectRef {
	if m != nil {
		return m.RoleRef
	}
	return nil
}

type MultiClusterRoleBindingStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiClusterRoleBindingStatus) Reset()         { *m = MultiClusterRoleBindingStatus{} }
func (m *MultiClusterRoleBindingStatus) String() string { return proto.CompactTextString(m) }
func (*MultiClusterRoleBindingStatus) ProtoMessage()    {}
func (*MultiClusterRoleBindingStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_538363f2c905c978, []int{4}
}
func (m *MultiClusterRoleBindingStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiClusterRoleBindingStatus.Unmarshal(m, b)
}
func (m *MultiClusterRoleBindingStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiClusterRoleBindingStatus.Marshal(b, m, deterministic)
}
func (m *MultiClusterRoleBindingStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiClusterRoleBindingStatus.Merge(m, src)
}
func (m *MultiClusterRoleBindingStatus) XXX_Size() int {
	return xxx_messageInfo_MultiClusterRoleBindingStatus.Size(m)
}
func (m *MultiClusterRoleBindingStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiClusterRoleBindingStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MultiClusterRoleBindingStatus proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("multicluster.solo.io.MultiClusterRoleSpec_Rule_Action", MultiClusterRoleSpec_Rule_Action_name, MultiClusterRoleSpec_Rule_Action_value)
	proto.RegisterType((*Placement)(nil), "multicluster.solo.io.Placement")
	proto.RegisterType((*MultiClusterRoleSpec)(nil), "multicluster.solo.io.MultiClusterRoleSpec")
	proto.RegisterType((*MultiClusterRoleSpec_Rule)(nil), "multicluster.solo.io.MultiClusterRoleSpec.Rule")
	proto.RegisterType((*MultiClusterRoleStatus)(nil), "multicluster.solo.io.MultiClusterRoleStatus")
	proto.RegisterType((*MultiClusterRoleBindingSpec)(nil), "multicluster.solo.io.MultiClusterRoleBindingSpec")
	proto.RegisterType((*MultiClusterRoleBindingStatus)(nil), "multicluster.solo.io.MultiClusterRoleBindingStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/skv2-enterprise/v1alpha1/multicluster.proto", fileDescriptor_538363f2c905c978)
}

var fileDescriptor_538363f2c905c978 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x26, 0x6d, 0xe9, 0xda, 0x77, 0x12, 0xaa, 0xac, 0x0a, 0x45, 0x2d, 0x74, 0x25, 0xa7, 0x5e,
	0x96, 0xb0, 0x22, 0xe0, 0x84, 0x50, 0xbb, 0x45, 0xbb, 0x94, 0x31, 0x65, 0x05, 0x09, 0x2e, 0x93,
	0x9b, 0xbe, 0xcd, 0x4c, 0xd3, 0xd8, 0xb2, 0x9d, 0x31, 0x7e, 0x08, 0x47, 0xee, 0xfc, 0x04, 0x7e,
	0x0f, 0xff, 0x81, 0x23, 0x12, 0xb2, 0xd3, 0x86, 0xc1, 0x0a, 0xd2, 0x4e, 0x7e, 0xbf, 0x9e, 0x47,
	0x8f, 0x1f, 0xbf, 0x86, 0x49, 0xc2, 0xf4, 0x45, 0x3e, 0xf3, 0x63, 0xbe, 0x0a, 0x14, 0x4f, 0xf9,
	0x3e, 0xe3, 0xc5, 0x49, 0x05, 0x53, 0x01, 0x15, 0x2c, 0x50, 0xcb, 0xcb, 0xe1, 0x3e, 0x66, 0x1a,
	0xa5, 0x90, 0x4c, 0x61, 0x70, 0x79, 0x40, 0x53, 0x71, 0x41, 0x0f, 0x82, 0x55, 0x9e, 0x6a, 0x16,
	0xa7, 0xb9, 0xd2, 0x28, 0x7d, 0x21, 0xb9, 0xe6, 0xa4, 0xfd, 0x47, 0xcd, 0xf0, 0xf8, 0x8c, 0x77,
	0xda, 0x09, 0x4f, 0xb8, 0x1d, 0x08, 0x4c, 0x54, 0xcc, 0x76, 0x08, 0x5e, 0xe9, 0xa2, 0x88, 0x57,
	0x7a, 0x5d, 0xeb, 0x25, 0x9c, 0x27, 0x29, 0x06, 0x36, 0x9b, 0xe5, 0x8b, 0xe0, 0xa3, 0xa4, 0x42,
	0xa0, 0x54, 0xeb, 0x7e, 0xd7, 0x88, 0xb1, 0xaa, 0x62, 0x2e, 0x8d, 0x14, 0x7b, 0x16, 0x4d, 0xef,
	0x18, 0x9a, 0xa7, 0x29, 0x8d, 0x71, 0x85, 0x99, 0x26, 0x3d, 0x80, 0x8c, 0xae, 0x50, 0x09, 0x1a,
	0xa3, 0x72, 0x9d, 0x7e, 0x75, 0xd0, 0x8c, 0xae, 0x55, 0x48, 0x07, 0x1a, 0x6b, 0x99, 0xca, 0xad,
	0xd8, 0x6e, 0x99, 0x7b, 0x3f, 0x2b, 0xd0, 0x7e, 0x65, 0x2e, 0x72, 0x58, 0x54, 0x22, 0x9e, 0xe2,
	0x99, 0xc0, 0x98, 0x84, 0x70, 0x57, 0xe6, 0xe9, 0x9a, 0x6f, 0x77, 0x18, 0xf8, 0xdb, 0xae, 0xeb,
	0x6f, 0x83, 0xfa, 0x51, 0x9e, 0x62, 0x54, 0xa0, 0x3b, 0x5f, 0x2a, 0x50, 0x33, 0x39, 0xe9, 0x42,
	0x93, 0x0a, 0x76, 0x9e, 0x48, 0x9e, 0x0b, 0xd7, 0xe9, 0x3b, 0x46, 0x05, 0x15, 0xec, 0xd8, 0xe4,
	0xe4, 0x31, 0xd4, 0x96, 0x2c, 0x9b, 0xbb, 0x95, 0xbe, 0x33, 0xd8, 0x1d, 0x3e, 0xf0, 0x0b, 0x6b,
	0xfc, 0x8d, 0x35, 0xfe, 0x99, 0x96, 0x2c, 0x4b, 0xde, 0xd2, 0x34, 0xc7, 0xc8, 0x4e, 0x92, 0x13,
	0xa8, 0xd3, 0x58, 0x33, 0x9e, 0xb9, 0xd5, 0xbe, 0x33, 0xb8, 0x37, 0x7c, 0x76, 0x4b, 0x7d, 0xfe,
	0xc8, 0xa2, 0xa3, 0x35, 0x0b, 0x79, 0x09, 0x20, 0x36, 0x86, 0x2a, 0xb7, 0x66, 0xef, 0xbc, 0xb7,
	0x9d, 0xb3, 0x34, 0x3e, 0xba, 0x06, 0xf1, 0x9e, 0x42, 0xbd, 0xa0, 0x24, 0x3b, 0x50, 0x1d, 0x9d,
	0xbc, 0x6b, 0xdd, 0x21, 0x00, 0xf5, 0xc3, 0x28, 0x1c, 0x4d, 0xc3, 0x96, 0x63, 0xe2, 0x37, 0xa7,
	0x47, 0x26, 0xae, 0x98, 0xf8, 0x28, 0x9c, 0x84, 0xd3, 0xb0, 0x55, 0xf5, 0x5c, 0xb8, 0x7f, 0x43,
	0xa3, 0xa6, 0x3a, 0x57, 0xde, 0x67, 0x07, 0xba, 0x7f, 0xb7, 0xc6, 0x2c, 0x9b, 0xb3, 0x2c, 0xb1,
	0x0f, 0xf4, 0x02, 0x1a, 0x2a, 0x9f, 0x7d, 0xc0, 0x58, 0x6f, 0xde, 0xe8, 0x91, 0x6f, 0x37, 0xc4,
	0xec, 0x4d, 0x29, 0x76, 0xfa, 0x49, 0xe0, 0xfc, 0xb5, 0x1d, 0x8b, 0x70, 0x11, 0x95, 0x10, 0xf2,
	0x1c, 0x1a, 0x92, 0xa7, 0x78, 0x2e, 0x71, 0x51, 0xda, 0x7e, 0x13, 0xfe, 0x1b, 0xb9, 0x63, 0xa6,
	0x23, 0x5c, 0x78, 0x7b, 0xf0, 0xf0, 0x5f, 0xb2, 0xac, 0xf0, 0xf1, 0xe4, 0xdb, 0x8f, 0x9a, 0xf3,
	0xf5, 0x7b, 0xcf, 0x79, 0x3f, 0xfe, 0xef, 0x87, 0x13, 0xcb, 0xc4, 0xae, 0xf7, 0x36, 0x93, 0xcb,
	0x9f, 0x37, 0xab, 0xdb, 0x25, 0x78, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x4e, 0x92, 0x1b,
	0xbd, 0x03, 0x00, 0x00,
}

func (this *Placement) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Placement)
	if !ok {
		that2, ok := that.(Placement)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if this.Clusters[i] != that1.Clusters[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MultiClusterRoleSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MultiClusterRoleSpec)
	if !ok {
		that2, ok := that.(MultiClusterRoleSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Rules) != len(that1.Rules) {
		return false
	}
	for i := range this.Rules {
		if !this.Rules[i].Equal(that1.Rules[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MultiClusterRoleSpec_Rule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MultiClusterRoleSpec_Rule)
	if !ok {
		that2, ok := that.(MultiClusterRoleSpec_Rule)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ApiGroup != that1.ApiGroup {
		return false
	}
	if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if this.Action != that1.Action {
		return false
	}
	if len(this.Placements) != len(that1.Placements) {
		return false
	}
	for i := range this.Placements {
		if !this.Placements[i].Equal(that1.Placements[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MultiClusterRoleStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MultiClusterRoleStatus)
	if !ok {
		that2, ok := that.(MultiClusterRoleStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MultiClusterRoleBindingSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MultiClusterRoleBindingSpec)
	if !ok {
		that2, ok := that.(MultiClusterRoleBindingSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Subjects) != len(that1.Subjects) {
		return false
	}
	for i := range this.Subjects {
		if !this.Subjects[i].Equal(that1.Subjects[i]) {
			return false
		}
	}
	if !this.RoleRef.Equal(that1.RoleRef) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MultiClusterRoleBindingStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MultiClusterRoleBindingStatus)
	if !ok {
		that2, ok := that.(MultiClusterRoleBindingStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
