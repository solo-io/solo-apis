// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/transformation_ee/transformation.proto

package transformation_ee

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	route "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/api/v2/route"
	_type "github.com/solo-io/solo-kit/pkg/api/external/envoy/type"
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

type FilterTransformations struct {
	// Specifies transformations based on the route matches. The first matched transformation will be
	// applied. If there are overlapped match conditions, please put the most specific match first.
	Transformations      []*TransformationRule `protobuf:"bytes,1,rep,name=transformations,proto3" json:"transformations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FilterTransformations) Reset()         { *m = FilterTransformations{} }
func (m *FilterTransformations) String() string { return proto.CompactTextString(m) }
func (*FilterTransformations) ProtoMessage()    {}
func (*FilterTransformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fb829ad719952e, []int{0}
}
func (m *FilterTransformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterTransformations.Unmarshal(m, b)
}
func (m *FilterTransformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterTransformations.Marshal(b, m, deterministic)
}
func (m *FilterTransformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterTransformations.Merge(m, src)
}
func (m *FilterTransformations) XXX_Size() int {
	return xxx_messageInfo_FilterTransformations.Size(m)
}
func (m *FilterTransformations) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterTransformations.DiscardUnknown(m)
}

var xxx_messageInfo_FilterTransformations proto.InternalMessageInfo

func (m *FilterTransformations) GetTransformations() []*TransformationRule {
	if m != nil {
		return m.Transformations
	}
	return nil
}

type TransformationRule struct {
	// The route matching parameter. Only when the match is satisfied, the "requires" field will
	// apply.
	//
	// For example: following match will match all requests.
	//
	// .. code-block:: yaml
	//
	//    match:
	//      prefix: /
	//
	Match *route.RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// transformation to perform
	RouteTransformations *RouteTransformations `protobuf:"bytes,2,opt,name=route_transformations,json=routeTransformations,proto3" json:"route_transformations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TransformationRule) Reset()         { *m = TransformationRule{} }
func (m *TransformationRule) String() string { return proto.CompactTextString(m) }
func (*TransformationRule) ProtoMessage()    {}
func (*TransformationRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fb829ad719952e, []int{1}
}
func (m *TransformationRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransformationRule.Unmarshal(m, b)
}
func (m *TransformationRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransformationRule.Marshal(b, m, deterministic)
}
func (m *TransformationRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransformationRule.Merge(m, src)
}
func (m *TransformationRule) XXX_Size() int {
	return xxx_messageInfo_TransformationRule.Size(m)
}
func (m *TransformationRule) XXX_DiscardUnknown() {
	xxx_messageInfo_TransformationRule.DiscardUnknown(m)
}

var xxx_messageInfo_TransformationRule proto.InternalMessageInfo

func (m *TransformationRule) GetMatch() *route.RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *TransformationRule) GetRouteTransformations() *RouteTransformations {
	if m != nil {
		return m.RouteTransformations
	}
	return nil
}

type RouteTransformations struct {
	RequestTransformation *Transformation `protobuf:"bytes,1,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	// clear the route cache if the request transformation was applied
	ClearRouteCache        bool            `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	ResponseTransformation *Transformation `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *RouteTransformations) Reset()         { *m = RouteTransformations{} }
func (m *RouteTransformations) String() string { return proto.CompactTextString(m) }
func (*RouteTransformations) ProtoMessage()    {}
func (*RouteTransformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fb829ad719952e, []int{2}
}
func (m *RouteTransformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTransformations.Unmarshal(m, b)
}
func (m *RouteTransformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTransformations.Marshal(b, m, deterministic)
}
func (m *RouteTransformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTransformations.Merge(m, src)
}
func (m *RouteTransformations) XXX_Size() int {
	return xxx_messageInfo_RouteTransformations.Size(m)
}
func (m *RouteTransformations) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTransformations.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTransformations proto.InternalMessageInfo

func (m *RouteTransformations) GetRequestTransformation() *Transformation {
	if m != nil {
		return m.RequestTransformation
	}
	return nil
}

func (m *RouteTransformations) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *RouteTransformations) GetResponseTransformation() *Transformation {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

type Transformation struct {
	// Template is in the transformed request language domain
	//
	// Types that are valid to be assigned to TransformationType:
	//	*Transformation_DlpTransformation
	TransformationType   isTransformation_TransformationType `protobuf_oneof:"transformation_type"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *Transformation) Reset()         { *m = Transformation{} }
func (m *Transformation) String() string { return proto.CompactTextString(m) }
func (*Transformation) ProtoMessage()    {}
func (*Transformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fb829ad719952e, []int{3}
}
func (m *Transformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transformation.Unmarshal(m, b)
}
func (m *Transformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transformation.Marshal(b, m, deterministic)
}
func (m *Transformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transformation.Merge(m, src)
}
func (m *Transformation) XXX_Size() int {
	return xxx_messageInfo_Transformation.Size(m)
}
func (m *Transformation) XXX_DiscardUnknown() {
	xxx_messageInfo_Transformation.DiscardUnknown(m)
}

var xxx_messageInfo_Transformation proto.InternalMessageInfo

type isTransformation_TransformationType interface {
	isTransformation_TransformationType()
}

type Transformation_DlpTransformation struct {
	DlpTransformation *DlpTransformation `protobuf:"bytes,1,opt,name=dlp_transformation,json=dlpTransformation,proto3,oneof" json:"dlp_transformation,omitempty"`
}

func (*Transformation_DlpTransformation) isTransformation_TransformationType() {}

func (m *Transformation) GetTransformationType() isTransformation_TransformationType {
	if m != nil {
		return m.TransformationType
	}
	return nil
}

func (m *Transformation) GetDlpTransformation() *DlpTransformation {
	if x, ok := m.GetTransformationType().(*Transformation_DlpTransformation); ok {
		return x.DlpTransformation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Transformation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Transformation_DlpTransformation)(nil),
	}
}

type DlpTransformation struct {
	// list of actions to apply
	Actions              []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DlpTransformation) Reset()         { *m = DlpTransformation{} }
func (m *DlpTransformation) String() string { return proto.CompactTextString(m) }
func (*DlpTransformation) ProtoMessage()    {}
func (*DlpTransformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fb829ad719952e, []int{4}
}
func (m *DlpTransformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DlpTransformation.Unmarshal(m, b)
}
func (m *DlpTransformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DlpTransformation.Marshal(b, m, deterministic)
}
func (m *DlpTransformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DlpTransformation.Merge(m, src)
}
func (m *DlpTransformation) XXX_Size() int {
	return xxx_messageInfo_DlpTransformation.Size(m)
}
func (m *DlpTransformation) XXX_DiscardUnknown() {
	xxx_messageInfo_DlpTransformation.DiscardUnknown(m)
}

var xxx_messageInfo_DlpTransformation proto.InternalMessageInfo

func (m *DlpTransformation) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type Action struct {
	// Identifier for this action.
	// Used mostly to help ID specific actions in logs.
	// If left null will default to unknown
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of regexes to apply to the response body to match data which should be masked
	// They will be applied iteratively in the order which they are specified
	Regex []string `protobuf:"bytes,2,rep,name=regex,proto3" json:"regex,omitempty"`
	// If specified, this rule will not actually be applied, but only logged.
	Shadow bool `protobuf:"varint,3,opt,name=shadow,proto3" json:"shadow,omitempty"`
	// The percent of the string which should be masked.
	// If not set, defaults to 75%
	Percent *_type.Percent `protobuf:"bytes,4,opt,name=percent,proto3" json:"percent,omitempty"`
	// The character which should overwrite the masked data
	// If left empty, defaults to "X"
	MaskChar             string   `protobuf:"bytes,5,opt,name=mask_char,json=maskChar,proto3" json:"mask_char,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fb829ad719952e, []int{5}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Action) GetRegex() []string {
	if m != nil {
		return m.Regex
	}
	return nil
}

func (m *Action) GetShadow() bool {
	if m != nil {
		return m.Shadow
	}
	return false
}

func (m *Action) GetPercent() *_type.Percent {
	if m != nil {
		return m.Percent
	}
	return nil
}

func (m *Action) GetMaskChar() string {
	if m != nil {
		return m.MaskChar
	}
	return ""
}

func init() {
	proto.RegisterType((*FilterTransformations)(nil), "envoy.config.filter.http.transformation_ee.v2.FilterTransformations")
	proto.RegisterType((*TransformationRule)(nil), "envoy.config.filter.http.transformation_ee.v2.TransformationRule")
	proto.RegisterType((*RouteTransformations)(nil), "envoy.config.filter.http.transformation_ee.v2.RouteTransformations")
	proto.RegisterType((*Transformation)(nil), "envoy.config.filter.http.transformation_ee.v2.Transformation")
	proto.RegisterType((*DlpTransformation)(nil), "envoy.config.filter.http.transformation_ee.v2.DlpTransformation")
	proto.RegisterType((*Action)(nil), "envoy.config.filter.http.transformation_ee.v2.Action")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/transformation_ee/transformation.proto", fileDescriptor_d1fb829ad719952e)
}

var fileDescriptor_d1fb829ad719952e = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0x76, 0xb6, 0x49, 0xda, 0x4c, 0xc5, 0xda, 0x69, 0xd3, 0x2e, 0x15, 0x4a, 0x58, 0x3c, 0x04,
	0xa1, 0xb3, 0x10, 0xf1, 0x24, 0x8a, 0xdd, 0x5a, 0xf1, 0x22, 0x96, 0xc1, 0x93, 0x20, 0x61, 0xba,
	0x99, 0x66, 0x87, 0x6c, 0x76, 0xa6, 0x33, 0x93, 0x98, 0xdc, 0x3d, 0x79, 0xf4, 0xe8, 0x7f, 0x10,
	0xbc, 0xfb, 0x1f, 0xfc, 0x41, 0x39, 0xc9, 0xce, 0x6c, 0xc4, 0xdd, 0x04, 0x61, 0xe9, 0x65, 0xb2,
	0xf3, 0xbd, 0xf7, 0xbe, 0xf7, 0xbd, 0x2f, 0xcc, 0x83, 0xe3, 0x11, 0x37, 0xc9, 0xf4, 0x1a, 0xc7,
	0x62, 0x12, 0x6a, 0x91, 0x8a, 0x33, 0x2e, 0xdc, 0x2f, 0x95, 0x5c, 0x87, 0x54, 0xf2, 0x70, 0x94,
	0x0a, 0xe1, 0x0e, 0x36, 0x37, 0x4c, 0x65, 0x34, 0x0d, 0x59, 0x36, 0x13, 0x0b, 0x7b, 0xcd, 0x34,
	0x17, 0x99, 0x0e, 0x8d, 0xa2, 0x99, 0xbe, 0x11, 0x6a, 0x42, 0x0d, 0x17, 0xd9, 0x80, 0xb1, 0x0a,
	0x82, 0xa5, 0x12, 0x46, 0xa0, 0x33, 0x5b, 0x88, 0x63, 0x91, 0xdd, 0xf0, 0x11, 0xbe, 0xe1, 0xa9,
	0x61, 0x0a, 0x27, 0xc6, 0x48, 0xbc, 0x46, 0x80, 0x67, 0xfd, 0x93, 0xe3, 0x19, 0x4d, 0xf9, 0x90,
	0x1a, 0x16, 0xae, 0x3e, 0x1c, 0xcf, 0xc9, 0xa9, 0x13, 0x90, 0xab, 0x9b, 0xf5, 0x43, 0x25, 0xa6,
	0x86, 0xb9, 0xb3, 0x88, 0xfb, 0x2e, 0x6e, 0x16, 0x92, 0x85, 0x92, 0xa9, 0x98, 0x65, 0xc6, 0x45,
	0x82, 0x2f, 0x00, 0x76, 0xde, 0xd8, 0xbe, 0x1f, 0x4a, 0x1d, 0x35, 0x1a, 0xc3, 0xbd, 0xb2, 0x08,
	0xed, 0x83, 0xee, 0x56, 0x6f, 0xb7, 0x7f, 0x8e, 0x6b, 0xa9, 0xc6, 0x65, 0x62, 0x32, 0x4d, 0x19,
	0xa9, 0x32, 0x07, 0xbf, 0x01, 0x44, 0xeb, 0x79, 0xe8, 0x25, 0x6c, 0x4e, 0xa8, 0x89, 0x13, 0x1f,
	0x74, 0x41, 0x6f, 0xb7, 0x7f, 0x5a, 0x74, 0xa6, 0x92, 0xe7, 0xc4, 0x6e, 0x42, 0x92, 0x9f, 0xef,
	0xf2, 0xac, 0x68, 0x67, 0x19, 0x35, 0xbf, 0x02, 0xef, 0x21, 0x20, 0xae, 0x0c, 0xcd, 0x61, 0xc7,
	0x26, 0x0d, 0xaa, 0x93, 0x78, 0x96, 0xef, 0xa2, 0xe6, 0x24, 0xb6, 0x55, 0xc5, 0x27, 0x72, 0xa8,
	0x36, 0xa0, 0xc1, 0x4f, 0x0f, 0x1e, 0x6e, 0x4a, 0x47, 0x06, 0x1e, 0x29, 0x76, 0x3b, 0x65, 0xda,
	0x54, 0x44, 0x15, 0x33, 0xbe, 0xb8, 0x9b, 0xbb, 0x9d, 0x82, 0xbc, 0x0c, 0xa3, 0x27, 0x70, 0x3f,
	0x4e, 0x19, 0x55, 0x03, 0x67, 0x47, 0x4c, 0xe3, 0x84, 0xf9, 0x5b, 0x5d, 0xd0, 0xdb, 0x21, 0x7b,
	0x36, 0x60, 0xb5, 0x5e, 0xe4, 0x30, 0x9a, 0xc1, 0x63, 0xc5, 0xb4, 0x14, 0x99, 0xae, 0xfa, 0x56,
	0xd8, 0x76, 0x47, 0x89, 0x47, 0x2b, 0xf6, 0x32, 0x1e, 0x7c, 0x07, 0xf0, 0x41, 0x45, 0xf6, 0x2d,
	0x44, 0xc3, 0x54, 0x6e, 0x36, 0xea, 0x55, 0x4d, 0x15, 0xaf, 0x53, 0x59, 0x66, 0x7f, 0x7b, 0x8f,
	0xec, 0x0f, 0xab, 0x60, 0xd4, 0x81, 0x07, 0x95, 0xf2, 0xfc, 0xd5, 0x04, 0x43, 0xb8, 0xbf, 0x46,
	0x80, 0xde, 0xc3, 0x6d, 0x1a, 0xff, 0xfb, 0x34, 0x9e, 0xd5, 0xd4, 0x74, 0x6e, 0xab, 0xc9, 0x8a,
	0x25, 0xf8, 0x01, 0x60, 0xcb, 0x61, 0x08, 0xc1, 0x46, 0x46, 0x27, 0xcc, 0x0e, 0xdb, 0x26, 0xf6,
	0x1b, 0x05, 0xb0, 0xa9, 0xd8, 0x88, 0xcd, 0x7d, 0xaf, 0xbb, 0xd5, 0x6b, 0x47, 0xf7, 0x97, 0x51,
	0xfb, 0x1b, 0x68, 0x05, 0x0d, 0xe5, 0x75, 0x01, 0x71, 0x21, 0x74, 0x04, 0x5b, 0x3a, 0xa1, 0x43,
	0xf1, 0xb9, 0xf8, 0x7b, 0x8b, 0x1b, 0x3a, 0x83, 0xdb, 0xc5, 0xcb, 0xf7, 0x1b, 0xd6, 0xbf, 0x83,
	0x42, 0x6b, 0x3e, 0x1e, 0xbe, 0x72, 0x21, 0xb2, 0xca, 0x41, 0x8f, 0x61, 0x7b, 0x42, 0xf5, 0x78,
	0x10, 0x27, 0x54, 0xf9, 0xcd, 0x5c, 0x43, 0xb4, 0xbd, 0x8c, 0x1a, 0xca, 0xeb, 0x01, 0xb2, 0x93,
	0x47, 0x2e, 0x12, 0xaa, 0xa2, 0x5f, 0x00, 0x3e, 0xe7, 0xc2, 0x11, 0x49, 0x25, 0xe6, 0x8b, 0x7a,
	0xf3, 0x47, 0x8f, 0xca, 0x86, 0x5e, 0x5e, 0xba, 0x55, 0x74, 0x95, 0xaf, 0xa6, 0x2b, 0xf0, 0xf1,
	0xd3, 0x7f, 0x77, 0xb1, 0x1c, 0x8f, 0xfe, 0xee, 0x63, 0x9c, 0xc3, 0x98, 0xd7, 0x5a, 0xc9, 0xd7,
	0x2d, 0xbb, 0x02, 0x9f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x30, 0x64, 0x8f, 0x6b, 0xf3, 0x05,
	0x00, 0x00,
}
