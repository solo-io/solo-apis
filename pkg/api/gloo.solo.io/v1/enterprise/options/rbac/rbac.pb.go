// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/rbac/rbac.proto

package rbac

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Global RBAC settings
type Settings struct {
	// Require RBAC for all virtual hosts. A vhost without an RBAC policy set will fallback to a deny-all policy.
	RequireRbac          bool     `protobuf:"varint,1,opt,name=require_rbac,json=requireRbac,proto3" json:"require_rbac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_623e06e48384f149, []int{0}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetRequireRbac() bool {
	if m != nil {
		return m.RequireRbac
	}
	return false
}

// RBAC settings for Virtual Hosts and Routes
type ExtensionSettings struct {
	// Disable RBAC checks on this resource (default false). This is useful to allow access to static resources/login page without RBAC checks.
	// If provided on a route, all route settings override any vhost settings
	Disable bool `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
	// Named policies to apply.
	Policies             map[string]*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExtensionSettings) Reset()         { *m = ExtensionSettings{} }
func (m *ExtensionSettings) String() string { return proto.CompactTextString(m) }
func (*ExtensionSettings) ProtoMessage()    {}
func (*ExtensionSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_623e06e48384f149, []int{1}
}
func (m *ExtensionSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionSettings.Unmarshal(m, b)
}
func (m *ExtensionSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionSettings.Marshal(b, m, deterministic)
}
func (m *ExtensionSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionSettings.Merge(m, src)
}
func (m *ExtensionSettings) XXX_Size() int {
	return xxx_messageInfo_ExtensionSettings.Size(m)
}
func (m *ExtensionSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionSettings proto.InternalMessageInfo

func (m *ExtensionSettings) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func (m *ExtensionSettings) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type Policy struct {
	// Principals in this policy.
	Principals []*Principal `protobuf:"bytes,1,rep,name=principals,proto3" json:"principals,omitempty"`
	// Permissions granted to the principals.
	Permissions          *Permissions `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_623e06e48384f149, []int{2}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *Policy) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// An RBAC principal - the identity entity (usually a user or a service account).
type Principal struct {
	JwtPrincipal         *JWTPrincipal `protobuf:"bytes,1,opt,name=jwt_principal,json=jwtPrincipal,proto3" json:"jwt_principal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_623e06e48384f149, []int{3}
}
func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

func (m *Principal) GetJwtPrincipal() *JWTPrincipal {
	if m != nil {
		return m.JwtPrincipal
	}
	return nil
}

// A JWT principal. To use this, JWT option MUST be enabled.
type JWTPrincipal struct {
	// Set of claims that make up this principal. Commonly, the 'iss' and 'sub' or 'email' claims are used.
	// all claims must be present on the JWT.
	Claims map[string]string `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Verify that the JWT came from a specific provider. This usually can be left empty
	// and a provider will be chosen automatically.
	Provider             string   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTPrincipal) Reset()         { *m = JWTPrincipal{} }
func (m *JWTPrincipal) String() string { return proto.CompactTextString(m) }
func (*JWTPrincipal) ProtoMessage()    {}
func (*JWTPrincipal) Descriptor() ([]byte, []int) {
	return fileDescriptor_623e06e48384f149, []int{4}
}
func (m *JWTPrincipal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWTPrincipal.Unmarshal(m, b)
}
func (m *JWTPrincipal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWTPrincipal.Marshal(b, m, deterministic)
}
func (m *JWTPrincipal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTPrincipal.Merge(m, src)
}
func (m *JWTPrincipal) XXX_Size() int {
	return xxx_messageInfo_JWTPrincipal.Size(m)
}
func (m *JWTPrincipal) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTPrincipal.DiscardUnknown(m)
}

var xxx_messageInfo_JWTPrincipal proto.InternalMessageInfo

func (m *JWTPrincipal) GetClaims() map[string]string {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *JWTPrincipal) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

// What permissions should be granted. An empty field means allow-all.
// If more than one field is added, all of them need to match.
type Permissions struct {
	// Paths that have this prefix will be allowed.
	PathPrefix string `protobuf:"bytes,1,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	// What http methods (GET, POST, ...) are allowed.
	Methods              []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_623e06e48384f149, []int{5}
}
func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *Permissions) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func init() {
	proto.RegisterType((*Settings)(nil), "rbac.options.gloo.solo.io.Settings")
	proto.RegisterType((*ExtensionSettings)(nil), "rbac.options.gloo.solo.io.ExtensionSettings")
	proto.RegisterMapType((map[string]*Policy)(nil), "rbac.options.gloo.solo.io.ExtensionSettings.PoliciesEntry")
	proto.RegisterType((*Policy)(nil), "rbac.options.gloo.solo.io.Policy")
	proto.RegisterType((*Principal)(nil), "rbac.options.gloo.solo.io.Principal")
	proto.RegisterType((*JWTPrincipal)(nil), "rbac.options.gloo.solo.io.JWTPrincipal")
	proto.RegisterMapType((map[string]string)(nil), "rbac.options.gloo.solo.io.JWTPrincipal.ClaimsEntry")
	proto.RegisterType((*Permissions)(nil), "rbac.options.gloo.solo.io.Permissions")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/rbac/rbac.proto", fileDescriptor_623e06e48384f149)
}

var fileDescriptor_623e06e48384f149 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0xa6, 0x10, 0xe2, 0x71, 0x2a, 0xc1, 0xaa, 0x07, 0xe3, 0x03, 0xa4, 0x16, 0x82, 0x5c,
	0x6a, 0x8b, 0xf4, 0x00, 0xf4, 0x08, 0x54, 0xaa, 0xa0, 0x87, 0xc8, 0xa0, 0x22, 0x38, 0x10, 0xd9,
	0xce, 0xe2, 0x4c, 0xeb, 0x78, 0x97, 0xdd, 0x4d, 0x9a, 0xbc, 0x09, 0x8f, 0xc0, 0x99, 0x13, 0xcf,
	0x83, 0xc4, 0x23, 0x70, 0x47, 0x5e, 0xff, 0xd4, 0x08, 0xb0, 0x7a, 0xb1, 0x77, 0xbe, 0x99, 0xef,
	0x9b, 0x99, 0x9d, 0x1d, 0x98, 0xa6, 0xa8, 0x17, 0xab, 0xd8, 0x4f, 0xf8, 0x32, 0x50, 0x3c, 0xe3,
	0x07, 0xc8, 0xcb, 0x7f, 0x24, 0x50, 0x05, 0x91, 0xc0, 0x20, 0xcd, 0x38, 0x2f, 0x3f, 0xeb, 0xc7,
	0x01, 0xcb, 0x35, 0x93, 0x42, 0xa2, 0x62, 0x01, 0x17, 0x1a, 0x79, 0xae, 0x02, 0x19, 0x47, 0x89,
	0xf9, 0xf8, 0x42, 0x72, 0xcd, 0xe9, 0x5d, 0x73, 0xae, 0xbc, 0x7e, 0x41, 0xf4, 0x0b, 0x41, 0x1f,
	0xb9, 0xbb, 0x97, 0xf2, 0x94, 0x9b, 0xa8, 0xa0, 0x38, 0x95, 0x04, 0x97, 0xb2, 0x8d, 0x2e, 0x41,
	0xb6, 0xd1, 0x25, 0xe6, 0x1d, 0xc0, 0xe0, 0x0d, 0xd3, 0x1a, 0xf3, 0x54, 0xd1, 0x7d, 0x18, 0x4a,
	0xf6, 0x79, 0x85, 0x92, 0xcd, 0x0a, 0x69, 0x87, 0x8c, 0xc8, 0x78, 0x10, 0xda, 0x15, 0x16, 0xc6,
	0x51, 0xe2, 0xfd, 0x24, 0x70, 0xe7, 0x78, 0xa3, 0x59, 0xae, 0x90, 0xe7, 0x0d, 0xd1, 0x81, 0x5b,
	0x73, 0x54, 0x51, 0x9c, 0xb1, 0x8a, 0x53, 0x9b, 0xf4, 0x0c, 0x06, 0x82, 0x67, 0x98, 0x20, 0x53,
	0x4e, 0x6f, 0xb4, 0x33, 0xb6, 0x27, 0x47, 0xfe, 0x7f, 0xcb, 0xf6, 0xff, 0x52, 0xf6, 0xa7, 0x15,
	0xf9, 0x38, 0xd7, 0x72, 0x1b, 0x36, 0x5a, 0xee, 0x47, 0xd8, 0xfd, 0xc3, 0x45, 0x6f, 0xc3, 0xce,
	0x05, 0xdb, 0x9a, 0xf4, 0x56, 0x58, 0x1c, 0xe9, 0x13, 0xb8, 0xb9, 0x8e, 0xb2, 0x15, 0x73, 0x7a,
	0x23, 0x32, 0xb6, 0x27, 0xfb, 0x1d, 0x79, 0x8d, 0xd4, 0x36, 0x2c, 0xe3, 0x8f, 0x7a, 0x4f, 0x89,
	0xf7, 0x85, 0x40, 0xbf, 0x44, 0xe9, 0x4b, 0x00, 0x21, 0x31, 0x4f, 0x50, 0x44, 0x99, 0x72, 0x88,
	0x69, 0xe2, 0x41, 0x97, 0x58, 0x1d, 0x1c, 0xb6, 0x78, 0xf4, 0x04, 0x6c, 0xc1, 0xe4, 0x12, 0x55,
	0xd1, 0x9e, 0xaa, 0x6a, 0x7a, 0xd8, 0x25, 0x73, 0x15, 0x1d, 0xb6, 0xa9, 0xde, 0x7b, 0xb0, 0x9a,
	0x14, 0xf4, 0x14, 0x76, 0xcf, 0x2f, 0xf5, 0xac, 0x49, 0x64, 0x2e, 0xc0, 0x9e, 0x3c, 0xea, 0x10,
	0x7e, 0xf5, 0xee, 0xed, 0x55, 0x89, 0xc3, 0xf3, 0x4b, 0xdd, 0x58, 0xde, 0x37, 0x02, 0xc3, 0xb6,
	0x9b, 0xbe, 0x86, 0x7e, 0x92, 0x45, 0xb8, 0xac, 0xfb, 0x3e, 0xbc, 0xa6, 0xae, 0xff, 0xc2, 0xb0,
	0xca, 0xa9, 0x55, 0x12, 0xd4, 0x85, 0x81, 0x90, 0x7c, 0x8d, 0x73, 0x26, 0x4d, 0xff, 0x56, 0xd8,
	0xd8, 0xee, 0x33, 0xb0, 0x5b, 0x94, 0x7f, 0x4c, 0x73, 0xaf, 0x3d, 0x4d, 0xab, 0x3d, 0xaa, 0x13,
	0xb0, 0x5b, 0x77, 0x45, 0xef, 0x83, 0x2d, 0x22, 0xbd, 0x98, 0x09, 0xc9, 0x3e, 0xe1, 0xa6, 0x92,
	0x80, 0x02, 0x9a, 0x1a, 0xa4, 0x78, 0xac, 0x4b, 0xa6, 0x17, 0x7c, 0x5e, 0xbe, 0x48, 0x2b, 0xac,
	0xcd, 0xe7, 0x67, 0xdf, 0x7f, 0xdd, 0x20, 0x5f, 0x7f, 0xdc, 0x23, 0x1f, 0x4e, 0x3b, 0x97, 0x55,
	0x5c, 0xa4, 0xcd, 0xc2, 0xd6, 0xed, 0x77, 0xec, 0x6c, 0xdc, 0x37, 0xab, 0x76, 0xf8, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0xce, 0x4e, 0x23, 0x5b, 0x03, 0x04, 0x00, 0x00,
}

func (this *Settings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
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
	if this.RequireRbac != that1.RequireRbac {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ExtensionSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ExtensionSettings)
	if !ok {
		that2, ok := that.(ExtensionSettings)
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
	if this.Disable != that1.Disable {
		return false
	}
	if len(this.Policies) != len(that1.Policies) {
		return false
	}
	for i := range this.Policies {
		if !this.Policies[i].Equal(that1.Policies[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Policy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Policy)
	if !ok {
		that2, ok := that.(Policy)
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
	if len(this.Principals) != len(that1.Principals) {
		return false
	}
	for i := range this.Principals {
		if !this.Principals[i].Equal(that1.Principals[i]) {
			return false
		}
	}
	if !this.Permissions.Equal(that1.Permissions) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Principal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Principal)
	if !ok {
		that2, ok := that.(Principal)
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
	if !this.JwtPrincipal.Equal(that1.JwtPrincipal) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *JWTPrincipal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JWTPrincipal)
	if !ok {
		that2, ok := that.(JWTPrincipal)
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
	if len(this.Claims) != len(that1.Claims) {
		return false
	}
	for i := range this.Claims {
		if this.Claims[i] != that1.Claims[i] {
			return false
		}
	}
	if this.Provider != that1.Provider {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Permissions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Permissions)
	if !ok {
		that2, ok := that.(Permissions)
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
	if this.PathPrefix != that1.PathPrefix {
		return false
	}
	if len(this.Methods) != len(that1.Methods) {
		return false
	}
	for i := range this.Methods {
		if this.Methods[i] != that1.Methods[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
