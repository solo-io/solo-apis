// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/rt_selector.proto

package v1

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v12 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	v11 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	matchers "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/core/matchers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SubRouteTableRow struct {
	// Types that are valid to be assigned to Action:
	//	*SubRouteTableRow_RouteAction
	//	*SubRouteTableRow_RedirectAction
	//	*SubRouteTableRow_DirectResponseAction
	//	*SubRouteTableRow_DelegateAction
	Action               isSubRouteTableRow_Action         `protobuf_oneof:"action"`
	Matcher              string                            `protobuf:"bytes,5,opt,name=matcher,proto3" json:"matcher,omitempty"`
	MatchType            string                            `protobuf:"bytes,6,opt,name=match_type,json=matchType,proto3" json:"match_type,omitempty"`
	Methods              []string                          `protobuf:"bytes,7,rep,name=methods,proto3" json:"methods,omitempty"`
	Headers              []*matchers.HeaderMatcher         `protobuf:"bytes,8,rep,name=headers,proto3" json:"headers,omitempty"`
	QueryParameters      []*matchers.QueryParameterMatcher `protobuf:"bytes,9,rep,name=query_parameters,json=queryParameters,proto3" json:"query_parameters,omitempty"`
	RtRoutes             []*SubRouteTableRow               `protobuf:"bytes,10,rep,name=rt_routes,json=rtRoutes,proto3" json:"rt_routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SubRouteTableRow) Reset()         { *m = SubRouteTableRow{} }
func (m *SubRouteTableRow) String() string { return proto.CompactTextString(m) }
func (*SubRouteTableRow) ProtoMessage()    {}
func (*SubRouteTableRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_18155afc6ed6dbe9, []int{0}
}
func (m *SubRouteTableRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubRouteTableRow.Unmarshal(m, b)
}
func (m *SubRouteTableRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubRouteTableRow.Marshal(b, m, deterministic)
}
func (m *SubRouteTableRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubRouteTableRow.Merge(m, src)
}
func (m *SubRouteTableRow) XXX_Size() int {
	return xxx_messageInfo_SubRouteTableRow.Size(m)
}
func (m *SubRouteTableRow) XXX_DiscardUnknown() {
	xxx_messageInfo_SubRouteTableRow.DiscardUnknown(m)
}

var xxx_messageInfo_SubRouteTableRow proto.InternalMessageInfo

type isSubRouteTableRow_Action interface {
	isSubRouteTableRow_Action()
	Equal(interface{}) bool
}

type SubRouteTableRow_RouteAction struct {
	RouteAction *v1.RouteAction `protobuf:"bytes,1,opt,name=route_action,json=routeAction,proto3,oneof" json:"route_action,omitempty"`
}
type SubRouteTableRow_RedirectAction struct {
	RedirectAction *v1.RedirectAction `protobuf:"bytes,2,opt,name=redirect_action,json=redirectAction,proto3,oneof" json:"redirect_action,omitempty"`
}
type SubRouteTableRow_DirectResponseAction struct {
	DirectResponseAction *v1.DirectResponseAction `protobuf:"bytes,3,opt,name=direct_response_action,json=directResponseAction,proto3,oneof" json:"direct_response_action,omitempty"`
}
type SubRouteTableRow_DelegateAction struct {
	DelegateAction *v11.DelegateAction `protobuf:"bytes,4,opt,name=delegate_action,json=delegateAction,proto3,oneof" json:"delegate_action,omitempty"`
}

func (*SubRouteTableRow_RouteAction) isSubRouteTableRow_Action()          {}
func (*SubRouteTableRow_RedirectAction) isSubRouteTableRow_Action()       {}
func (*SubRouteTableRow_DirectResponseAction) isSubRouteTableRow_Action() {}
func (*SubRouteTableRow_DelegateAction) isSubRouteTableRow_Action()       {}

func (m *SubRouteTableRow) GetAction() isSubRouteTableRow_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *SubRouteTableRow) GetRouteAction() *v1.RouteAction {
	if x, ok := m.GetAction().(*SubRouteTableRow_RouteAction); ok {
		return x.RouteAction
	}
	return nil
}

func (m *SubRouteTableRow) GetRedirectAction() *v1.RedirectAction {
	if x, ok := m.GetAction().(*SubRouteTableRow_RedirectAction); ok {
		return x.RedirectAction
	}
	return nil
}

func (m *SubRouteTableRow) GetDirectResponseAction() *v1.DirectResponseAction {
	if x, ok := m.GetAction().(*SubRouteTableRow_DirectResponseAction); ok {
		return x.DirectResponseAction
	}
	return nil
}

func (m *SubRouteTableRow) GetDelegateAction() *v11.DelegateAction {
	if x, ok := m.GetAction().(*SubRouteTableRow_DelegateAction); ok {
		return x.DelegateAction
	}
	return nil
}

func (m *SubRouteTableRow) GetMatcher() string {
	if m != nil {
		return m.Matcher
	}
	return ""
}

func (m *SubRouteTableRow) GetMatchType() string {
	if m != nil {
		return m.MatchType
	}
	return ""
}

func (m *SubRouteTableRow) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *SubRouteTableRow) GetHeaders() []*matchers.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *SubRouteTableRow) GetQueryParameters() []*matchers.QueryParameterMatcher {
	if m != nil {
		return m.QueryParameters
	}
	return nil
}

func (m *SubRouteTableRow) GetRtRoutes() []*SubRouteTableRow {
	if m != nil {
		return m.RtRoutes
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubRouteTableRow) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubRouteTableRow_RouteAction)(nil),
		(*SubRouteTableRow_RedirectAction)(nil),
		(*SubRouteTableRow_DirectResponseAction)(nil),
		(*SubRouteTableRow_DelegateAction)(nil),
	}
}

type GetVirtualServiceRoutesRequest struct {
	VirtualServiceRef    *v12.ClusterObjectRef `protobuf:"bytes,1,opt,name=virtual_service_ref,json=virtualServiceRef,proto3" json:"virtual_service_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetVirtualServiceRoutesRequest) Reset()         { *m = GetVirtualServiceRoutesRequest{} }
func (m *GetVirtualServiceRoutesRequest) String() string { return proto.CompactTextString(m) }
func (*GetVirtualServiceRoutesRequest) ProtoMessage()    {}
func (*GetVirtualServiceRoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18155afc6ed6dbe9, []int{1}
}
func (m *GetVirtualServiceRoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVirtualServiceRoutesRequest.Unmarshal(m, b)
}
func (m *GetVirtualServiceRoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVirtualServiceRoutesRequest.Marshal(b, m, deterministic)
}
func (m *GetVirtualServiceRoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVirtualServiceRoutesRequest.Merge(m, src)
}
func (m *GetVirtualServiceRoutesRequest) XXX_Size() int {
	return xxx_messageInfo_GetVirtualServiceRoutesRequest.Size(m)
}
func (m *GetVirtualServiceRoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVirtualServiceRoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVirtualServiceRoutesRequest proto.InternalMessageInfo

func (m *GetVirtualServiceRoutesRequest) GetVirtualServiceRef() *v12.ClusterObjectRef {
	if m != nil {
		return m.VirtualServiceRef
	}
	return nil
}

type GetVirtualServiceRoutesResponse struct {
	VsRoutes             []*SubRouteTableRow `protobuf:"bytes,1,rep,name=vs_routes,json=vsRoutes,proto3" json:"vs_routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetVirtualServiceRoutesResponse) Reset()         { *m = GetVirtualServiceRoutesResponse{} }
func (m *GetVirtualServiceRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*GetVirtualServiceRoutesResponse) ProtoMessage()    {}
func (*GetVirtualServiceRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18155afc6ed6dbe9, []int{2}
}
func (m *GetVirtualServiceRoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVirtualServiceRoutesResponse.Unmarshal(m, b)
}
func (m *GetVirtualServiceRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVirtualServiceRoutesResponse.Marshal(b, m, deterministic)
}
func (m *GetVirtualServiceRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVirtualServiceRoutesResponse.Merge(m, src)
}
func (m *GetVirtualServiceRoutesResponse) XXX_Size() int {
	return xxx_messageInfo_GetVirtualServiceRoutesResponse.Size(m)
}
func (m *GetVirtualServiceRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVirtualServiceRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVirtualServiceRoutesResponse proto.InternalMessageInfo

func (m *GetVirtualServiceRoutesResponse) GetVsRoutes() []*SubRouteTableRow {
	if m != nil {
		return m.VsRoutes
	}
	return nil
}

func init() {
	proto.RegisterType((*SubRouteTableRow)(nil), "fed.rpc.solo.io.SubRouteTableRow")
	proto.RegisterType((*GetVirtualServiceRoutesRequest)(nil), "fed.rpc.solo.io.GetVirtualServiceRoutesRequest")
	proto.RegisterType((*GetVirtualServiceRoutesResponse)(nil), "fed.rpc.solo.io.GetVirtualServiceRoutesResponse")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/rt_selector.proto", fileDescriptor_18155afc6ed6dbe9)
}

var fileDescriptor_18155afc6ed6dbe9 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x97, 0xdf, 0xb6, 0xfe, 0x71, 0x7f, 0xa2, 0xc3, 0x4c, 0x2c, 0x14, 0xd8, 0x4a, 0x91,
	0x50, 0xb9, 0x58, 0x42, 0xbb, 0x0b, 0xc4, 0xcd, 0xa4, 0x6d, 0x88, 0x4d, 0x48, 0x08, 0x48, 0x27,
	0x2e, 0x26, 0xa4, 0xc8, 0x4d, 0x4e, 0xdb, 0xb0, 0x74, 0x4e, 0x6d, 0x27, 0x6b, 0x6f, 0xb8, 0xe1,
	0x29, 0x78, 0x03, 0x1e, 0x81, 0xe7, 0xe1, 0x1d, 0xb8, 0x47, 0xb1, 0x9d, 0x54, 0xe9, 0xba, 0x6a,
	0x37, 0xad, 0x7d, 0xfc, 0xfd, 0x7e, 0x8e, 0x8f, 0x7d, 0x62, 0xf4, 0x6e, 0x18, 0x88, 0x51, 0xdc,
	0xb7, 0x3c, 0x3a, 0xb6, 0x39, 0x0d, 0xe9, 0x7e, 0x40, 0xd5, 0x3f, 0x89, 0x02, 0x6e, 0x93, 0x28,
	0xb0, 0x87, 0x21, 0xa5, 0xfb, 0x03, 0xf0, 0xed, 0x01, 0xf8, 0x16, 0x8b, 0x3c, 0x3b, 0xe9, 0xd8,
	0x4c, 0xb8, 0x1c, 0x42, 0xf0, 0x04, 0x65, 0x56, 0xc4, 0xa8, 0xa0, 0xb8, 0xae, 0x57, 0xad, 0xd4,
	0x6c, 0x05, 0xb4, 0xb1, 0x3d, 0xa4, 0x43, 0x2a, 0xd7, 0xec, 0x74, 0xa4, 0x64, 0x0d, 0x0c, 0x53,
	0xa1, 0x82, 0x30, 0x15, 0x3a, 0xf6, 0xfa, 0x66, 0x3e, 0xf5, 0x93, 0x74, 0x6c, 0x8f, 0x32, 0xb0,
	0xc7, 0x44, 0x78, 0x23, 0x60, 0x3c, 0x1f, 0x68, 0xe3, 0x8b, 0x15, 0xc6, 0x88, 0xd1, 0xe9, 0x4c,
	0xeb, 0x0e, 0x96, 0xe9, 0x88, 0x80, 0x6b, 0x32, 0x4b, 0xa5, 0x49, 0xc0, 0x44, 0x4c, 0x42, 0x97,
	0x03, 0x4b, 0x02, 0x0f, 0xb4, 0xe9, 0x31, 0xbf, 0x4c, 0xba, 0x52, 0x2f, 0xf7, 0xa0, 0xf7, 0xa2,
	0x16, 0x5b, 0x3f, 0x36, 0xd1, 0x56, 0x2f, 0xee, 0x3b, 0x34, 0x16, 0x70, 0x4e, 0xfa, 0x21, 0x38,
	0xf4, 0x1a, 0x1f, 0xa2, 0xff, 0x59, 0x1a, 0x70, 0x89, 0x27, 0x02, 0x7a, 0x65, 0x1a, 0x4d, 0xa3,
	0x5d, 0xeb, 0x3e, 0xb2, 0xd2, 0x7c, 0xd9, 0xb1, 0x58, 0xd2, 0x72, 0x24, 0x05, 0x67, 0x6b, 0x4e,
	0x8d, 0xcd, 0xa7, 0xf8, 0x14, 0xd5, 0x19, 0xf8, 0x01, 0x03, 0x4f, 0x64, 0x88, 0xff, 0x24, 0xe2,
	0xc9, 0x02, 0x42, 0x8b, 0x72, 0xca, 0x3d, 0x56, 0x88, 0xe0, 0x0b, 0xf4, 0x50, 0x63, 0x18, 0xf0,
	0x88, 0x5e, 0xf1, 0x7c, 0x4b, 0xeb, 0x92, 0xd7, 0x2a, 0xf2, 0xde, 0x4a, 0xad, 0xa3, 0xa5, 0x39,
	0x75, 0xdb, 0x5f, 0x12, 0xc7, 0xef, 0x51, 0xdd, 0x87, 0x10, 0xd2, 0xe3, 0xcb, 0xa0, 0x1b, 0x12,
	0xba, 0x67, 0xe9, 0x23, 0x9d, 0x73, 0xb5, 0x6e, 0xbe, 0x4f, 0xbf, 0x10, 0xc1, 0x26, 0x2a, 0xeb,
	0x1b, 0x35, 0x37, 0x9b, 0x46, 0xbb, 0xea, 0x64, 0x53, 0xfc, 0x14, 0x21, 0x39, 0x74, 0xc5, 0x2c,
	0x02, 0xb3, 0x24, 0x17, 0xab, 0x32, 0x72, 0x3e, 0x8b, 0x40, 0x1a, 0x41, 0x8c, 0xa8, 0xcf, 0xcd,
	0x72, 0x73, 0x5d, 0x1a, 0xd5, 0x14, 0x9f, 0xa0, 0xf2, 0x08, 0x88, 0x0f, 0x8c, 0x9b, 0x95, 0xe6,
	0x7a, 0xbb, 0xd6, 0x7d, 0x69, 0xe5, 0x4d, 0x23, 0xef, 0xaf, 0x50, 0xf9, 0x99, 0x94, 0x7e, 0x50,
	0x02, 0x27, 0x73, 0xe2, 0xaf, 0x68, 0x6b, 0x12, 0x03, 0x9b, 0xb9, 0x11, 0x61, 0x64, 0x0c, 0x22,
	0xa5, 0x55, 0x25, 0xad, 0xb3, 0x8a, 0xf6, 0x39, 0xf5, 0x7c, 0xca, 0x2c, 0x19, 0xb5, 0x3e, 0x29,
	0x84, 0x39, 0x3e, 0x44, 0x55, 0x26, 0x5c, 0x79, 0xf1, 0xdc, 0x44, 0x12, 0xfb, 0xcc, 0x5a, 0xf8,
	0x7a, 0xac, 0xc5, 0xe6, 0x72, 0x2a, 0x4c, 0xc8, 0x00, 0x3f, 0xae, 0xa0, 0x92, 0x3a, 0xf8, 0x56,
	0x8c, 0x76, 0x4f, 0x41, 0x7c, 0x51, 0xed, 0xdb, 0x53, 0xdd, 0xab, 0x44, 0x0e, 0x4c, 0x62, 0xe0,
	0x02, 0xf7, 0xd0, 0x83, 0x85, 0xee, 0x76, 0x19, 0x0c, 0x74, 0x67, 0x3e, 0x57, 0x35, 0xa4, 0x7d,
	0x9e, 0xe7, 0x3d, 0x09, 0x63, 0x2e, 0x80, 0x7d, 0xec, 0x7f, 0x93, 0x57, 0x3f, 0x70, 0xee, 0x27,
	0x45, 0x3c, 0x0c, 0x5a, 0x04, 0xed, 0xdd, 0x9a, 0x56, 0xf5, 0x4a, 0x5a, 0x63, 0xc2, 0xb3, 0x1a,
	0x8d, 0x3b, 0xd7, 0x98, 0x70, 0xc5, 0xe9, 0xfe, 0x34, 0xd0, 0xce, 0xb2, 0x04, 0x47, 0x51, 0x80,
	0xbf, 0xa3, 0x9d, 0x5b, 0xd2, 0x63, 0xfb, 0x46, 0x8e, 0xd5, 0xe7, 0xd3, 0x78, 0x75, 0x77, 0x83,
	0xaa, 0xac, 0xb5, 0x76, 0xfc, 0xe6, 0xf7, 0xdf, 0x0d, 0xe3, 0xd7, 0x9f, 0x5d, 0xe3, 0xc2, 0x5e,
	0xf9, 0x76, 0x46, 0x97, 0x43, 0xf9, 0x7c, 0xcc, 0x9f, 0xcd, 0x7e, 0x49, 0xbe, 0x1e, 0x07, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x17, 0x06, 0x6a, 0xcb, 0x75, 0x05, 0x00, 0x00,
}

func (this *SubRouteTableRow) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubRouteTableRow)
	if !ok {
		that2, ok := that.(SubRouteTableRow)
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
	if that1.Action == nil {
		if this.Action != nil {
			return false
		}
	} else if this.Action == nil {
		return false
	} else if !this.Action.Equal(that1.Action) {
		return false
	}
	if this.Matcher != that1.Matcher {
		return false
	}
	if this.MatchType != that1.MatchType {
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
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if !this.Headers[i].Equal(that1.Headers[i]) {
			return false
		}
	}
	if len(this.QueryParameters) != len(that1.QueryParameters) {
		return false
	}
	for i := range this.QueryParameters {
		if !this.QueryParameters[i].Equal(that1.QueryParameters[i]) {
			return false
		}
	}
	if len(this.RtRoutes) != len(that1.RtRoutes) {
		return false
	}
	for i := range this.RtRoutes {
		if !this.RtRoutes[i].Equal(that1.RtRoutes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SubRouteTableRow_RouteAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubRouteTableRow_RouteAction)
	if !ok {
		that2, ok := that.(SubRouteTableRow_RouteAction)
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
	if !this.RouteAction.Equal(that1.RouteAction) {
		return false
	}
	return true
}
func (this *SubRouteTableRow_RedirectAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubRouteTableRow_RedirectAction)
	if !ok {
		that2, ok := that.(SubRouteTableRow_RedirectAction)
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
	if !this.RedirectAction.Equal(that1.RedirectAction) {
		return false
	}
	return true
}
func (this *SubRouteTableRow_DirectResponseAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubRouteTableRow_DirectResponseAction)
	if !ok {
		that2, ok := that.(SubRouteTableRow_DirectResponseAction)
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
	if !this.DirectResponseAction.Equal(that1.DirectResponseAction) {
		return false
	}
	return true
}
func (this *SubRouteTableRow_DelegateAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubRouteTableRow_DelegateAction)
	if !ok {
		that2, ok := that.(SubRouteTableRow_DelegateAction)
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
	if !this.DelegateAction.Equal(that1.DelegateAction) {
		return false
	}
	return true
}
func (this *GetVirtualServiceRoutesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetVirtualServiceRoutesRequest)
	if !ok {
		that2, ok := that.(GetVirtualServiceRoutesRequest)
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
	if !this.VirtualServiceRef.Equal(that1.VirtualServiceRef) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GetVirtualServiceRoutesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetVirtualServiceRoutesResponse)
	if !ok {
		that2, ok := that.(GetVirtualServiceRoutesResponse)
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
	if len(this.VsRoutes) != len(that1.VsRoutes) {
		return false
	}
	for i := range this.VsRoutes {
		if !this.VsRoutes[i].Equal(that1.VsRoutes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualServiceRoutesApiClient is the client API for VirtualServiceRoutesApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualServiceRoutesApiClient interface {
	GetVirtualServiceRoutes(ctx context.Context, in *GetVirtualServiceRoutesRequest, opts ...grpc.CallOption) (*GetVirtualServiceRoutesResponse, error)
}

type virtualServiceRoutesApiClient struct {
	cc *grpc.ClientConn
}

func NewVirtualServiceRoutesApiClient(cc *grpc.ClientConn) VirtualServiceRoutesApiClient {
	return &virtualServiceRoutesApiClient{cc}
}

func (c *virtualServiceRoutesApiClient) GetVirtualServiceRoutes(ctx context.Context, in *GetVirtualServiceRoutesRequest, opts ...grpc.CallOption) (*GetVirtualServiceRoutesResponse, error) {
	out := new(GetVirtualServiceRoutesResponse)
	err := c.cc.Invoke(ctx, "/fed.rpc.solo.io.VirtualServiceRoutesApi/GetVirtualServiceRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualServiceRoutesApiServer is the server API for VirtualServiceRoutesApi service.
type VirtualServiceRoutesApiServer interface {
	GetVirtualServiceRoutes(context.Context, *GetVirtualServiceRoutesRequest) (*GetVirtualServiceRoutesResponse, error)
}

// UnimplementedVirtualServiceRoutesApiServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualServiceRoutesApiServer struct {
}

func (*UnimplementedVirtualServiceRoutesApiServer) GetVirtualServiceRoutes(ctx context.Context, req *GetVirtualServiceRoutesRequest) (*GetVirtualServiceRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVirtualServiceRoutes not implemented")
}

func RegisterVirtualServiceRoutesApiServer(s *grpc.Server, srv VirtualServiceRoutesApiServer) {
	s.RegisterService(&_VirtualServiceRoutesApi_serviceDesc, srv)
}

func _VirtualServiceRoutesApi_GetVirtualServiceRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVirtualServiceRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualServiceRoutesApiServer).GetVirtualServiceRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fed.rpc.solo.io.VirtualServiceRoutesApi/GetVirtualServiceRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualServiceRoutesApiServer).GetVirtualServiceRoutes(ctx, req.(*GetVirtualServiceRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualServiceRoutesApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fed.rpc.solo.io.VirtualServiceRoutesApi",
	HandlerType: (*VirtualServiceRoutesApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVirtualServiceRoutes",
			Handler:    _VirtualServiceRoutesApi_GetVirtualServiceRoutes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/rt_selector.proto",
}
