// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/wasm/wasm.proto

package wasm

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// represents the different types of WASM VMs available with which envoy can run
// the WASM filter module
type WasmFilter_VmType int32

const (
	WasmFilter_V8   WasmFilter_VmType = 0
	WasmFilter_WAVM WasmFilter_VmType = 1
)

var WasmFilter_VmType_name = map[int32]string{
	0: "V8",
	1: "WAVM",
}

var WasmFilter_VmType_value = map[string]int32{
	"V8":   0,
	"WAVM": 1,
}

func (x WasmFilter_VmType) String() string {
	return proto.EnumName(WasmFilter_VmType_name, int32(x))
}

func (WasmFilter_VmType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d6af8991406bfa4, []int{1, 0}
}

// list of filter stages which can be selected for a WASM filter
type FilterStage_Stage int32

const (
	FilterStage_FaultStage     FilterStage_Stage = 0
	FilterStage_CorsStage      FilterStage_Stage = 1
	FilterStage_WafStage       FilterStage_Stage = 2
	FilterStage_AuthNStage     FilterStage_Stage = 3
	FilterStage_AuthZStage     FilterStage_Stage = 4
	FilterStage_RateLimitStage FilterStage_Stage = 5
	FilterStage_AcceptedStage  FilterStage_Stage = 6
	FilterStage_OutAuthStage   FilterStage_Stage = 7
	FilterStage_RouteStage     FilterStage_Stage = 8
)

var FilterStage_Stage_name = map[int32]string{
	0: "FaultStage",
	1: "CorsStage",
	2: "WafStage",
	3: "AuthNStage",
	4: "AuthZStage",
	5: "RateLimitStage",
	6: "AcceptedStage",
	7: "OutAuthStage",
	8: "RouteStage",
}

var FilterStage_Stage_value = map[string]int32{
	"FaultStage":     0,
	"CorsStage":      1,
	"WafStage":       2,
	"AuthNStage":     3,
	"AuthZStage":     4,
	"RateLimitStage": 5,
	"AcceptedStage":  6,
	"OutAuthStage":   7,
	"RouteStage":     8,
}

func (x FilterStage_Stage) String() string {
	return proto.EnumName(FilterStage_Stage_name, int32(x))
}

func (FilterStage_Stage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d6af8991406bfa4, []int{2, 0}
}

// During is the 0th member so that it is the default, even though
// the reading order can be a little confusing
type FilterStage_Predicate int32

const (
	FilterStage_During FilterStage_Predicate = 0
	FilterStage_Before FilterStage_Predicate = 1
	FilterStage_After  FilterStage_Predicate = 2
)

var FilterStage_Predicate_name = map[int32]string{
	0: "During",
	1: "Before",
	2: "After",
}

var FilterStage_Predicate_value = map[string]int32{
	"During": 0,
	"Before": 1,
	"After":  2,
}

func (x FilterStage_Predicate) String() string {
	return proto.EnumName(FilterStage_Predicate_name, int32(x))
}

func (FilterStage_Predicate) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d6af8991406bfa4, []int{2, 1}
}

//
//Options config for WASM filters
type PluginSource struct {
	// list of WASM filters to be added into the filter chain
	Filters              []*WasmFilter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PluginSource) Reset()         { *m = PluginSource{} }
func (m *PluginSource) String() string { return proto.CompactTextString(m) }
func (*PluginSource) ProtoMessage()    {}
func (*PluginSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d6af8991406bfa4, []int{0}
}
func (m *PluginSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginSource.Unmarshal(m, b)
}
func (m *PluginSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginSource.Marshal(b, m, deterministic)
}
func (m *PluginSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginSource.Merge(m, src)
}
func (m *PluginSource) XXX_Size() int {
	return xxx_messageInfo_PluginSource.Size(m)
}
func (m *PluginSource) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginSource.DiscardUnknown(m)
}

var xxx_messageInfo_PluginSource proto.InternalMessageInfo

func (m *PluginSource) GetFilters() []*WasmFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

//
//This message defines a single Envoy WASM filter to be placed into the filter chain
type WasmFilter struct {
	// name of image which houses the compiled wasm filter
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// string of the config sent to the wasm filter
	// currently has to be json or will crash
	// TODO: update to proto.Any or proto.Struct? and then turn into json
	Config string `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	// the stage in the filter chain where this filter should be placed
	FilterStage *FilterStage `protobuf:"bytes,4,opt,name=filter_stage,json=filterStage,proto3" json:"filter_stage,omitempty"`
	// the name of the filter, used for logging
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// the root_id of the filter which should be run, if this value is incorrect, or
	// empty the filter will crash
	RootId string `protobuf:"bytes,6,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	// selected VM type
	VmType               WasmFilter_VmType `protobuf:"varint,7,opt,name=vm_type,json=vmType,proto3,enum=wasm.options.gloo.solo.io.WasmFilter_VmType" json:"vm_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WasmFilter) Reset()         { *m = WasmFilter{} }
func (m *WasmFilter) String() string { return proto.CompactTextString(m) }
func (*WasmFilter) ProtoMessage()    {}
func (*WasmFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d6af8991406bfa4, []int{1}
}
func (m *WasmFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmFilter.Unmarshal(m, b)
}
func (m *WasmFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmFilter.Marshal(b, m, deterministic)
}
func (m *WasmFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmFilter.Merge(m, src)
}
func (m *WasmFilter) XXX_Size() int {
	return xxx_messageInfo_WasmFilter.Size(m)
}
func (m *WasmFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmFilter.DiscardUnknown(m)
}

var xxx_messageInfo_WasmFilter proto.InternalMessageInfo

func (m *WasmFilter) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *WasmFilter) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *WasmFilter) GetFilterStage() *FilterStage {
	if m != nil {
		return m.FilterStage
	}
	return nil
}

func (m *WasmFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WasmFilter) GetRootId() string {
	if m != nil {
		return m.RootId
	}
	return ""
}

func (m *WasmFilter) GetVmType() WasmFilter_VmType {
	if m != nil {
		return m.VmType
	}
	return WasmFilter_V8
}

type FilterStage struct {
	// stage of the filter chain in which the selected filter should be added
	Stage FilterStage_Stage `protobuf:"varint,1,opt,name=stage,proto3,enum=wasm.options.gloo.solo.io.FilterStage_Stage" json:"stage,omitempty"`
	// How this filter should be placed relative to the stage
	Predicate            FilterStage_Predicate `protobuf:"varint,2,opt,name=predicate,proto3,enum=wasm.options.gloo.solo.io.FilterStage_Predicate" json:"predicate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FilterStage) Reset()         { *m = FilterStage{} }
func (m *FilterStage) String() string { return proto.CompactTextString(m) }
func (*FilterStage) ProtoMessage()    {}
func (*FilterStage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d6af8991406bfa4, []int{2}
}
func (m *FilterStage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterStage.Unmarshal(m, b)
}
func (m *FilterStage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterStage.Marshal(b, m, deterministic)
}
func (m *FilterStage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterStage.Merge(m, src)
}
func (m *FilterStage) XXX_Size() int {
	return xxx_messageInfo_FilterStage.Size(m)
}
func (m *FilterStage) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterStage.DiscardUnknown(m)
}

var xxx_messageInfo_FilterStage proto.InternalMessageInfo

func (m *FilterStage) GetStage() FilterStage_Stage {
	if m != nil {
		return m.Stage
	}
	return FilterStage_FaultStage
}

func (m *FilterStage) GetPredicate() FilterStage_Predicate {
	if m != nil {
		return m.Predicate
	}
	return FilterStage_During
}

func init() {
	proto.RegisterEnum("wasm.options.gloo.solo.io.WasmFilter_VmType", WasmFilter_VmType_name, WasmFilter_VmType_value)
	proto.RegisterEnum("wasm.options.gloo.solo.io.FilterStage_Stage", FilterStage_Stage_name, FilterStage_Stage_value)
	proto.RegisterEnum("wasm.options.gloo.solo.io.FilterStage_Predicate", FilterStage_Predicate_name, FilterStage_Predicate_value)
	proto.RegisterType((*PluginSource)(nil), "wasm.options.gloo.solo.io.PluginSource")
	proto.RegisterType((*WasmFilter)(nil), "wasm.options.gloo.solo.io.WasmFilter")
	proto.RegisterType((*FilterStage)(nil), "wasm.options.gloo.solo.io.FilterStage")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/wasm/wasm.proto", fileDescriptor_9d6af8991406bfa4)
}

var fileDescriptor_9d6af8991406bfa4 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x8e, 0x9d, 0xd8, 0x49, 0x26, 0x69, 0xe4, 0x7f, 0x54, 0xfd, 0x98, 0x1c, 0x50, 0x64, 0x09,
	0x94, 0x03, 0xd8, 0x10, 0x2e, 0xdc, 0x50, 0x02, 0xad, 0x54, 0x01, 0x6d, 0xe5, 0xa2, 0x44, 0xea,
	0x25, 0x72, 0x9d, 0xb5, 0xbb, 0x22, 0xce, 0x5a, 0xf6, 0x3a, 0xb4, 0x0f, 0xc2, 0x91, 0x3b, 0x8f,
	0xc0, 0xf3, 0xf0, 0x06, 0x1c, 0xb8, 0xa3, 0xdd, 0x49, 0x9a, 0x0a, 0x89, 0xd2, 0xcb, 0xee, 0x7c,
	0xdf, 0xce, 0xf7, 0xcd, 0xcc, 0x4a, 0x03, 0x07, 0x29, 0x97, 0x97, 0xd5, 0x85, 0x1f, 0x8b, 0x2c,
	0x28, 0xc5, 0x52, 0x3c, 0xe3, 0x82, 0xee, 0x28, 0xe7, 0x65, 0x10, 0xe5, 0x3c, 0x48, 0x97, 0x42,
	0xd0, 0xb1, 0x7e, 0x11, 0x88, 0x5c, 0x72, 0xb1, 0x2a, 0x83, 0xcf, 0x51, 0x99, 0xe9, 0xc3, 0xcf,
	0x0b, 0x21, 0x05, 0x3e, 0xd4, 0xf1, 0xe6, 0xd5, 0x57, 0xd9, 0xbe, 0x72, 0xf1, 0xb9, 0xe8, 0xef,
	0xa7, 0x22, 0x15, 0x3a, 0x2b, 0x50, 0x11, 0x09, 0xfa, 0xc8, 0xae, 0x24, 0x91, 0xec, 0x4a, 0x12,
	0xe7, 0x9d, 0x40, 0xf7, 0x74, 0x59, 0xa5, 0x7c, 0x75, 0x26, 0xaa, 0x22, 0x66, 0xf8, 0x1a, 0x9a,
	0x09, 0x5f, 0x4a, 0x56, 0x94, 0xae, 0x31, 0xa8, 0x0f, 0x3b, 0xa3, 0xc7, 0xfe, 0x5f, 0xcb, 0xf8,
	0xb3, 0xa8, 0xcc, 0x0e, 0x75, 0x76, 0xb8, 0x55, 0x79, 0x5f, 0x4c, 0x80, 0x1d, 0x8f, 0xfb, 0x60,
	0xf1, 0x2c, 0x4a, 0x99, 0x6b, 0x0e, 0x8c, 0x61, 0x3b, 0x24, 0x80, 0xff, 0x83, 0x1d, 0x8b, 0x55,
	0xc2, 0x53, 0xb7, 0xae, 0xe9, 0x0d, 0xc2, 0x23, 0xe8, 0x92, 0xcf, 0xbc, 0x94, 0x4a, 0xd4, 0x18,
	0x18, 0xc3, 0xce, 0xe8, 0xc9, 0x1d, 0x2d, 0x50, 0x99, 0x33, 0x95, 0x1d, 0x76, 0x92, 0x1d, 0x40,
	0x84, 0xc6, 0x2a, 0xca, 0x98, 0x6b, 0xe9, 0x02, 0x3a, 0xc6, 0x07, 0xd0, 0x2c, 0x84, 0x90, 0x73,
	0xbe, 0x70, 0x6d, 0xaa, 0xab, 0xe0, 0xd1, 0x02, 0x0f, 0xa0, 0xb9, 0xce, 0xe6, 0xf2, 0x3a, 0x67,
	0x6e, 0x73, 0x60, 0x0c, 0x7b, 0xa3, 0xa7, 0xf7, 0x9a, 0xda, 0x9f, 0x66, 0x1f, 0xaf, 0x73, 0x16,
	0xda, 0x6b, 0x7d, 0x7b, 0x7d, 0xb0, 0x89, 0x41, 0x1b, 0xcc, 0xe9, 0x2b, 0xa7, 0x86, 0x2d, 0x68,
	0xcc, 0xc6, 0xd3, 0x0f, 0x8e, 0xe1, 0xfd, 0x34, 0xa1, 0x73, 0xab, 0x59, 0x9c, 0x80, 0x45, 0x33,
	0x1a, 0xff, 0x2c, 0x78, 0x4b, 0xe6, 0xd3, 0xa4, 0x24, 0xc5, 0x63, 0x68, 0xe7, 0x05, 0x5b, 0xf0,
	0x38, 0x92, 0xf4, 0xc1, 0xbd, 0xd1, 0xf3, 0x7b, 0xfa, 0x9c, 0x6e, 0x75, 0xe1, 0xce, 0xc2, 0xfb,
	0x6a, 0x80, 0x45, 0xdd, 0xf5, 0x00, 0x0e, 0xa3, 0x6a, 0x29, 0x35, 0x72, 0x6a, 0xb8, 0x07, 0xed,
	0x37, 0xa2, 0x28, 0x09, 0x1a, 0xd8, 0x85, 0xd6, 0x2c, 0x4a, 0x08, 0x99, 0x2a, 0x79, 0x5c, 0xc9,
	0xcb, 0x63, 0xc2, 0xf5, 0x2d, 0x3e, 0x27, 0xdc, 0x40, 0x84, 0x5e, 0x18, 0x49, 0xf6, 0x9e, 0x67,
	0x7c, 0x63, 0x68, 0xe1, 0x7f, 0xb0, 0x37, 0x8e, 0x63, 0x96, 0x4b, 0xb6, 0x20, 0xca, 0x46, 0x07,
	0xba, 0x27, 0x95, 0x54, 0x4a, 0x62, 0x9a, 0xca, 0x28, 0x14, 0x95, 0x64, 0x84, 0x5b, 0x9e, 0x0f,
	0xed, 0x9b, 0xbe, 0x11, 0xc0, 0x7e, 0x5b, 0x15, 0x7c, 0x95, 0x3a, 0x35, 0x15, 0x4f, 0x58, 0x22,
	0x0a, 0xd5, 0x5b, 0x1b, 0xac, 0x71, 0x22, 0x59, 0xe1, 0x98, 0x93, 0x77, 0xdf, 0x7f, 0x35, 0x8c,
	0x6f, 0x3f, 0x1e, 0x19, 0xe7, 0xe3, 0x3b, 0x57, 0x2e, 0xff, 0x94, 0xde, 0xac, 0xdd, 0xf6, 0xb7,
	0xfe, 0xdc, 0xbc, 0x0b, 0x5b, 0x2f, 0xcc, 0xcb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x93, 0xe1,
	0x99, 0x8b, 0xbe, 0x03, 0x00, 0x00,
}

func (this *PluginSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PluginSource)
	if !ok {
		that2, ok := that.(PluginSource)
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
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if !this.Filters[i].Equal(that1.Filters[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WasmFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WasmFilter)
	if !ok {
		that2, ok := that.(WasmFilter)
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
	if this.Image != that1.Image {
		return false
	}
	if this.Config != that1.Config {
		return false
	}
	if !this.FilterStage.Equal(that1.FilterStage) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.RootId != that1.RootId {
		return false
	}
	if this.VmType != that1.VmType {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FilterStage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FilterStage)
	if !ok {
		that2, ok := that.(FilterStage)
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
	if this.Stage != that1.Stage {
		return false
	}
	if this.Predicate != that1.Predicate {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
