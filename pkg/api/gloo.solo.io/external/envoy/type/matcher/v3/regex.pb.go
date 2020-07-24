// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/type/matcher/v3/regex.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
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

// A regex matcher designed for safety when used with untrusted input.
type RegexMatcher struct {
	// Types that are valid to be assigned to EngineType:
	//	*RegexMatcher_GoogleRe2
	EngineType isRegexMatcher_EngineType `protobuf_oneof:"engine_type"`
	// The regex match string. The string must be supported by the configured engine.
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegexMatcher) Reset()         { *m = RegexMatcher{} }
func (m *RegexMatcher) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher) ProtoMessage()    {}
func (*RegexMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_75e5f35cce31867d, []int{0}
}
func (m *RegexMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher.Unmarshal(m, b)
}
func (m *RegexMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher.Marshal(b, m, deterministic)
}
func (m *RegexMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher.Merge(m, src)
}
func (m *RegexMatcher) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher.Size(m)
}
func (m *RegexMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher proto.InternalMessageInfo

type isRegexMatcher_EngineType interface {
	isRegexMatcher_EngineType()
	Equal(interface{}) bool
}

type RegexMatcher_GoogleRe2 struct {
	GoogleRe2 *RegexMatcher_GoogleRE2 `protobuf:"bytes,1,opt,name=google_re2,json=googleRe2,proto3,oneof" json:"google_re2,omitempty"`
}

func (*RegexMatcher_GoogleRe2) isRegexMatcher_EngineType() {}

func (m *RegexMatcher) GetEngineType() isRegexMatcher_EngineType {
	if m != nil {
		return m.EngineType
	}
	return nil
}

func (m *RegexMatcher) GetGoogleRe2() *RegexMatcher_GoogleRE2 {
	if x, ok := m.GetEngineType().(*RegexMatcher_GoogleRe2); ok {
		return x.GoogleRe2
	}
	return nil
}

func (m *RegexMatcher) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RegexMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RegexMatcher_GoogleRe2)(nil),
	}
}

// Google's `RE2 <https://github.com/google/re2>`_ regex engine. The regex string must adhere to
// the documented `syntax <https://github.com/google/re2/wiki/Syntax>`_. The engine is designed
// to complete execution in linear time as well as limit the amount of memory used.
//
// Envoy supports program size checking via runtime. The runtime keys `re2.max_program_size.error_level`
// and `re2.max_program_size.warn_level` can be set to integers as the maximum program size or
// complexity that a compiled regex can have before an exception is thrown or a warning is
// logged, respectively. `re2.max_program_size.error_level` defaults to 100, and
// `re2.max_program_size.warn_level` has no default if unset (will not check/log a warning).
//
// Envoy emits two stats for tracking the program size of regexes: the histogram `re2.program_size`,
// which records the program size, and the counter `re2.exceeded_warn_level`, which is incremented
// each time the program size exceeds the warn level threshold.
type RegexMatcher_GoogleRE2 struct {
	// This field controls the RE2 "program size" which is a rough estimate of how complex a
	// compiled regex is to evaluate. A regex that has a program size greater than the configured
	// value will fail to compile. In this case, the configured max program size can be increased
	// or the regex can be simplified. If not specified, the default is 100.
	//
	// This field is deprecated; regexp validation should be performed on the management server
	// instead of being done by each individual client.
	MaxProgramSize       *types.UInt32Value `protobuf:"bytes,1,opt,name=max_program_size,json=maxProgramSize,proto3" json:"max_program_size,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RegexMatcher_GoogleRE2) Reset()         { *m = RegexMatcher_GoogleRE2{} }
func (m *RegexMatcher_GoogleRE2) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher_GoogleRE2) ProtoMessage()    {}
func (*RegexMatcher_GoogleRE2) Descriptor() ([]byte, []int) {
	return fileDescriptor_75e5f35cce31867d, []int{0, 0}
}
func (m *RegexMatcher_GoogleRE2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Unmarshal(m, b)
}
func (m *RegexMatcher_GoogleRE2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Marshal(b, m, deterministic)
}
func (m *RegexMatcher_GoogleRE2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher_GoogleRE2.Merge(m, src)
}
func (m *RegexMatcher_GoogleRE2) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Size(m)
}
func (m *RegexMatcher_GoogleRE2) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher_GoogleRE2.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher_GoogleRE2 proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *RegexMatcher_GoogleRE2) GetMaxProgramSize() *types.UInt32Value {
	if m != nil {
		return m.MaxProgramSize
	}
	return nil
}

// Describes how to match a string and then produce a new string using a regular
// expression and a substitution string.
type RegexMatchAndSubstitute struct {
	// The regular expression used to find portions of a string (hereafter called
	// the "subject string") that should be replaced. When a new string is
	// produced during the substitution operation, the new string is initially
	// the same as the subject string, but then all matches in the subject string
	// are replaced by the substitution string. If replacing all matches isn't
	// desired, regular expression anchors can be used to ensure a single match,
	// so as to replace just one occurrence of a pattern. Capture groups can be
	// used in the pattern to extract portions of the subject string, and then
	// referenced in the substitution string.
	Pattern *RegexMatcher `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	// The string that should be substituted into matching portions of the
	// subject string during a substitution operation to produce a new string.
	// Capture groups in the pattern can be referenced in the substitution
	// string. Note, however, that the syntax for referring to capture groups is
	// defined by the chosen regular expression engine. Google's `RE2
	// <https://github.com/google/re2>`_ regular expression engine uses a
	// backslash followed by the capture group number to denote a numbered
	// capture group. E.g., ``\1`` refers to capture group 1, and ``\2`` refers
	// to capture group 2.
	Substitution         string   `protobuf:"bytes,2,opt,name=substitution,proto3" json:"substitution,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegexMatchAndSubstitute) Reset()         { *m = RegexMatchAndSubstitute{} }
func (m *RegexMatchAndSubstitute) String() string { return proto.CompactTextString(m) }
func (*RegexMatchAndSubstitute) ProtoMessage()    {}
func (*RegexMatchAndSubstitute) Descriptor() ([]byte, []int) {
	return fileDescriptor_75e5f35cce31867d, []int{1}
}
func (m *RegexMatchAndSubstitute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatchAndSubstitute.Unmarshal(m, b)
}
func (m *RegexMatchAndSubstitute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatchAndSubstitute.Marshal(b, m, deterministic)
}
func (m *RegexMatchAndSubstitute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatchAndSubstitute.Merge(m, src)
}
func (m *RegexMatchAndSubstitute) XXX_Size() int {
	return xxx_messageInfo_RegexMatchAndSubstitute.Size(m)
}
func (m *RegexMatchAndSubstitute) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatchAndSubstitute.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatchAndSubstitute proto.InternalMessageInfo

func (m *RegexMatchAndSubstitute) GetPattern() *RegexMatcher {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *RegexMatchAndSubstitute) GetSubstitution() string {
	if m != nil {
		return m.Substitution
	}
	return ""
}

func init() {
	proto.RegisterType((*RegexMatcher)(nil), "envoy.type.matcher.v3.RegexMatcher")
	proto.RegisterType((*RegexMatcher_GoogleRE2)(nil), "envoy.type.matcher.v3.RegexMatcher.GoogleRE2")
	proto.RegisterType((*RegexMatchAndSubstitute)(nil), "envoy.type.matcher.v3.RegexMatchAndSubstitute")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/type/matcher/v3/regex.proto", fileDescriptor_75e5f35cce31867d)
}

var fileDescriptor_75e5f35cce31867d = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6a, 0x14, 0x4d,
	0x14, 0xfd, 0xaa, 0x3f, 0x93, 0x38, 0x95, 0x20, 0x43, 0xa3, 0x64, 0x18, 0xcc, 0x38, 0x99, 0x80,
	0xc4, 0x9f, 0x54, 0x69, 0xcf, 0x6e, 0x76, 0x36, 0x48, 0x54, 0x10, 0x86, 0x0e, 0x11, 0x71, 0x33,
	0xd4, 0x64, 0xca, 0x4a, 0x61, 0x77, 0xdd, 0xa2, 0xba, 0xba, 0xed, 0xc9, 0xca, 0xa5, 0xba, 0x73,
	0xeb, 0x13, 0xf8, 0x0a, 0xba, 0x17, 0xdc, 0x89, 0xaf, 0xe0, 0x53, 0x48, 0x56, 0xd2, 0x5d, 0xdd,
	0x89, 0xc1, 0x21, 0x71, 0x33, 0x73, 0xb9, 0xf7, 0x9c, 0xcb, 0x39, 0xa7, 0xfa, 0xe2, 0x7d, 0x21,
	0xed, 0x61, 0x36, 0x25, 0x07, 0x90, 0xd0, 0x14, 0x62, 0xd8, 0x91, 0xe0, 0xfe, 0x99, 0x96, 0x29,
	0x65, 0x5a, 0x52, 0x11, 0x03, 0xb8, 0x1f, 0x5e, 0x58, 0x6e, 0x14, 0x8b, 0x29, 0x57, 0x39, 0xcc,
	0xa9, 0x9d, 0x6b, 0x4e, 0x13, 0x66, 0x0f, 0x0e, 0xb9, 0xa1, 0xf9, 0x90, 0x1a, 0x2e, 0x78, 0x41,
	0xb4, 0x01, 0x0b, 0xfe, 0xb5, 0x0a, 0x42, 0x4a, 0x08, 0xa9, 0x21, 0x24, 0x1f, 0x76, 0x7b, 0x02,
	0x40, 0xc4, 0x9c, 0x56, 0xa0, 0x69, 0xf6, 0x92, 0xbe, 0x36, 0x4c, 0x6b, 0x6e, 0x52, 0x47, 0xeb,
	0x6e, 0x64, 0x33, 0xcd, 0x28, 0x53, 0x0a, 0x2c, 0xb3, 0x12, 0x54, 0x4a, 0x53, 0xcb, 0x6c, 0xd6,
	0x8c, 0x37, 0xff, 0x1a, 0xe7, 0xdc, 0xa4, 0x12, 0x94, 0x54, 0xa2, 0x86, 0xac, 0xe7, 0x2c, 0x96,
	0x33, 0x66, 0x39, 0x6d, 0x8a, 0x7a, 0x70, 0x55, 0x80, 0x80, 0xaa, 0xa4, 0x65, 0xe5, 0xba, 0x83,
	0xef, 0x1e, 0x5e, 0x8b, 0x4a, 0xdd, 0x4f, 0x9d, 0x48, 0xff, 0x39, 0xc6, 0x4e, 0xe3, 0xc4, 0xf0,
	0xa0, 0x83, 0xfa, 0x68, 0x7b, 0x35, 0xd8, 0x21, 0x0b, 0xdd, 0x90, 0x3f, 0x89, 0x64, 0xb7, 0x62,
	0x45, 0x0f, 0x83, 0xf0, 0xf2, 0x71, 0xb8, 0xf4, 0x1e, 0x79, 0x6d, 0xf4, 0xe8, 0xbf, 0xa8, 0xe5,
	0x96, 0x45, 0x3c, 0xf0, 0x37, 0xf0, 0x52, 0x95, 0x50, 0xc7, 0xeb, 0xa3, 0xed, 0x56, 0xb8, 0x72,
	0x1c, 0x5e, 0x32, 0x5e, 0x1f, 0x45, 0xae, 0xdb, 0x7d, 0x87, 0x70, 0xeb, 0x64, 0x87, 0xff, 0x04,
	0xb7, 0x13, 0x56, 0x4c, 0xb4, 0x01, 0x61, 0x58, 0x32, 0x49, 0xe5, 0x11, 0xaf, 0xc5, 0x5c, 0x27,
	0x6e, 0x25, 0x69, 0x32, 0x24, 0xfb, 0x8f, 0x95, 0x1d, 0x06, 0xcf, 0x58, 0x9c, 0xf1, 0xd0, 0xeb,
	0xa0, 0xe8, 0x4a, 0xc2, 0x8a, 0xb1, 0x23, 0xee, 0xc9, 0x23, 0x3e, 0xba, 0xf7, 0xf1, 0xeb, 0xdb,
	0xde, 0x1d, 0x7c, 0x6b, 0x81, 0x89, 0xc5, 0x0e, 0x46, 0x37, 0x4b, 0xc6, 0x26, 0xbe, 0x71, 0x01,
	0x23, 0xf4, 0xf1, 0x2a, 0x57, 0x42, 0x2a, 0x3e, 0x29, 0x31, 0xfe, 0xff, 0xbf, 0x42, 0x34, 0xf8,
	0x8c, 0xf0, 0xfa, 0x29, 0xe8, 0x81, 0x9a, 0xed, 0x65, 0xd3, 0xd4, 0x4a, 0x9b, 0x59, 0xee, 0xef,
	0xe2, 0x15, 0xcd, 0x6c, 0xf9, 0x11, 0xd5, 0x66, 0xb6, 0xfe, 0x21, 0xd9, 0xd3, 0x3c, 0xa3, 0x86,
	0xed, 0x0f, 0xf0, 0x5a, 0xda, 0xac, 0x95, 0xa0, 0x5c, 0xa4, 0xd1, 0x99, 0xde, 0xe8, 0x7e, 0x69,
	0xe2, 0x2e, 0xbe, 0x7d, 0xae, 0x89, 0x33, 0xfa, 0xc2, 0x0f, 0xe8, 0xd3, 0xcf, 0x1e, 0xfa, 0xf2,
	0xe6, 0xdb, 0x8f, 0x65, 0xaf, 0xed, 0xe1, 0x2d, 0x09, 0x4e, 0x9f, 0x36, 0x50, 0xcc, 0x17, 0x4b,
	0x0d, 0x71, 0xb5, 0x6c, 0x5c, 0x3e, 0xc6, 0x18, 0xbd, 0x18, 0x9f, 0x7b, 0x4e, 0xfa, 0x95, 0x38,
	0x39, 0x29, 0x52, 0xb6, 0x89, 0xbc, 0xe8, 0xaa, 0xa6, 0xcb, 0xd5, 0x3b, 0x0f, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xee, 0x65, 0x40, 0xb7, 0xa9, 0x03, 0x00, 0x00,
}

func (this *RegexMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegexMatcher)
	if !ok {
		that2, ok := that.(RegexMatcher)
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
	if that1.EngineType == nil {
		if this.EngineType != nil {
			return false
		}
	} else if this.EngineType == nil {
		return false
	} else if !this.EngineType.Equal(that1.EngineType) {
		return false
	}
	if this.Regex != that1.Regex {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RegexMatcher_GoogleRe2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegexMatcher_GoogleRe2)
	if !ok {
		that2, ok := that.(RegexMatcher_GoogleRe2)
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
	if !this.GoogleRe2.Equal(that1.GoogleRe2) {
		return false
	}
	return true
}
func (this *RegexMatcher_GoogleRE2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegexMatcher_GoogleRE2)
	if !ok {
		that2, ok := that.(RegexMatcher_GoogleRE2)
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
	if !this.MaxProgramSize.Equal(that1.MaxProgramSize) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RegexMatchAndSubstitute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegexMatchAndSubstitute)
	if !ok {
		that2, ok := that.(RegexMatchAndSubstitute)
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
	if !this.Pattern.Equal(that1.Pattern) {
		return false
	}
	if this.Substitution != that1.Substitution {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
