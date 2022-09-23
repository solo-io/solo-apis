// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/route_table.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RouteTableStatus_State int32

const (
	// Pending status indicates the resource has not yet been validated
	RouteTableStatus_Pending RouteTableStatus_State = 0
	// Accepted indicates the resource has been validated
	RouteTableStatus_Accepted RouteTableStatus_State = 1
	// Rejected indicates an invalid configuration by the user
	// Rejected resources may be propagated to the xDS server depending on their severity
	RouteTableStatus_Rejected RouteTableStatus_State = 2
	// Warning indicates a partially invalid configuration by the user
	// Resources with Warnings may be partially accepted by a controller, depending on the implementation
	RouteTableStatus_Warning RouteTableStatus_State = 3
)

// Enum value maps for RouteTableStatus_State.
var (
	RouteTableStatus_State_name = map[int32]string{
		0: "Pending",
		1: "Accepted",
		2: "Rejected",
		3: "Warning",
	}
	RouteTableStatus_State_value = map[string]int32{
		"Pending":  0,
		"Accepted": 1,
		"Rejected": 2,
		"Warning":  3,
	}
)

func (x RouteTableStatus_State) Enum() *RouteTableStatus_State {
	p := new(RouteTableStatus_State)
	*p = x
	return p
}

func (x RouteTableStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteTableStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_enumTypes[0].Descriptor()
}

func (RouteTableStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_enumTypes[0]
}

func (x RouteTableStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteTableStatus_State.Descriptor instead.
func (RouteTableStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescGZIP(), []int{1, 0}
}

// The **RouteTable** is a child routing object for the Gloo Gateway.
//
// A **RouteTable** gets built into the complete routing configuration when it is referenced by a `delegateAction`,
// either in a parent VirtualService or another RouteTable.
//
// Routes specified in a RouteTable must have their paths start with the prefix provided in the parent's matcher.
//
// For example, the following configuration:
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /1
//
// ```
//
// would *not be valid*, while
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /a/1
//
// ```
//
// *would* be valid.
//
// A complete configuration might look as follows:
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: VirtualService
// metadata:
//
//	name: 'any'
//	namespace: 'any'
//
// spec:
//
//	virtualHost:
//	  domains:
//	  - 'any.com'
//	  routes:
//	  - matchers:
//	    - prefix: '/a' # delegate ownership of routes for `any.com/a`
//	    delegateAction:
//	      ref:
//	        name: 'a-routes'
//	        namespace: 'a'
//	  - matchers:
//	    - prefix: '/b' # delegate ownership of routes for `any.com/b`
//	    delegateAction:
//	      ref:
//	        name: 'b-routes'
//	        namespace: 'b'
//
// ```
//
// * A root-level **VirtualService** which delegates routing to to the `a-routes` and `b-routes` **RouteTables**.
// * Routes with `delegateActions` can only use a `prefix` matcher.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//
//	name: 'a-routes'
//	namespace: 'a'
//
// spec:
//
//	routes:
//	  - matchers:
//	    # the path matchers in this RouteTable must begin with the prefix `/a/`
//	    - prefix: '/a/1'
//	    routeAction:
//	      single:
//	        upstream:
//	          name: 'foo-upstream'
//
//	  - matchers:
//	    - prefix: '/a/2'
//	    routeAction:
//	      single:
//	        upstream:
//	          name: 'bar-upstream'
//
// ```
//
// * A **RouteTable** which defines two routes.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//
//	name: 'b-routes'
//	namespace: 'b'
//
// spec:
//
//	routes:
//	  - matchers:
//	    # the path matchers in this RouteTable must begin with the prefix `/b/`
//	    - regex: '/b/3'
//	    routeAction:
//	      single:
//	        upstream:
//	          name: 'bar-upstream'
//	  - matchers:
//	    - prefix: '/b/c/'
//	    # routes in the RouteTable can perform any action, including a delegateAction
//	    delegateAction:
//	      ref:
//	        name: 'c-routes'
//	        namespace: 'c'
//
// ```
//
// * A **RouteTable** which both *defines a route* and *delegates to* another **RouteTable**.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//
//	name: 'c-routes'
//	namespace: 'c'
//
// spec:
//
//	routes:
//	  - matchers:
//	    - exact: '/b/c/4'
//	    routeAction:
//	      single:
//	        upstream:
//	          name: 'qux-upstream'
//
// ```
//
// * A RouteTable which is a child of another route table.
//
// Would produce the following route config for `mydomain.com`:
//
// ```
// /a/1 -> foo-upstream
// /a/2 -> bar-upstream
// /b/3 -> baz-upstream
// /b/c/4 -> qux-upstream
// ```
type RouteTableSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of routes for the route table
	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// When a delegated route defines a `RouteTableSelector` that matches multiple route tables, Gloo will inspect this
	// field to determine the order in which the route tables are to be evaluated. This determines the order in which
	// the routes will appear on the final `Proxy` resource. The field is optional; if no value is specified, the weight
	// defaults to 0 (zero).
	//
	// Gloo will process the route tables matched by a selector in ascending order by weight and collect the routes of
	// each route table in the order they are defined. If multiple route tables define the same weight, Gloo will sort the
	// routes which belong to those tables to avoid short-circuiting (e.g. making sure `/foo/bar` comes before `/foo`).
	// In this scenario, Gloo will also alert the user by adding a warning to the status of the parent resource
	// (the one that specifies the `RouteTableSelector`).
	Weight *wrappers.Int32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *RouteTableSpec) Reset() {
	*x = RouteTableSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableSpec) ProtoMessage() {}

func (x *RouteTableSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableSpec.ProtoReflect.Descriptor instead.
func (*RouteTableSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescGZIP(), []int{0}
}

func (x *RouteTableSpec) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *RouteTableSpec) GetWeight() *wrappers.Int32Value {
	if x != nil {
		return x.Weight
	}
	return nil
}

type RouteTableStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// State is the enum indicating the state of the resource
	State RouteTableStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=gateway.solo.io.RouteTableStatus_State" json:"state,omitempty"`
	// Reason is a description of the error for Rejected resources. If the resource is pending or accepted, this field will be empty
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	// Reference to the reporter who wrote this status
	ReportedBy string `protobuf:"bytes,3,opt,name=reported_by,json=reportedBy,proto3" json:"reported_by,omitempty"`
	// Reference to statuses (by resource-ref string: "Kind.Namespace.Name") of subresources of the parent resource
	SubresourceStatuses map[string]*RouteTableStatus `protobuf:"bytes,4,rep,name=subresource_statuses,json=subresourceStatuses,proto3" json:"subresource_statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Opaque details about status results
	Details *_struct.Struct `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *RouteTableStatus) Reset() {
	*x = RouteTableStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableStatus) ProtoMessage() {}

func (x *RouteTableStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableStatus.ProtoReflect.Descriptor instead.
func (*RouteTableStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescGZIP(), []int{1}
}

func (x *RouteTableStatus) GetState() RouteTableStatus_State {
	if x != nil {
		return x.State
	}
	return RouteTableStatus_Pending
}

func (x *RouteTableStatus) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *RouteTableStatus) GetReportedBy() string {
	if x != nil {
		return x.ReportedBy
	}
	return ""
}

func (x *RouteTableStatus) GetSubresourceStatuses() map[string]*RouteTableStatus {
	if x != nil {
		return x.SubresourceStatuses
	}
	return nil
}

func (x *RouteTableStatus) GetDetails() *_struct.Struct {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xd6, 0x03, 0x0a, 0x10, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x6d, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x13, 0x73, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x69, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x10, 0x03, 0x42, 0x41, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04,
	0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_goTypes = []interface{}{
	(RouteTableStatus_State)(0), // 0: gateway.solo.io.RouteTableStatus.State
	(*RouteTableSpec)(nil),      // 1: gateway.solo.io.RouteTableSpec
	(*RouteTableStatus)(nil),    // 2: gateway.solo.io.RouteTableStatus
	nil,                         // 3: gateway.solo.io.RouteTableStatus.SubresourceStatusesEntry
	(*Route)(nil),               // 4: gateway.solo.io.Route
	(*wrappers.Int32Value)(nil), // 5: google.protobuf.Int32Value
	(*_struct.Struct)(nil),      // 6: google.protobuf.Struct
}
var file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_depIdxs = []int32{
	4, // 0: gateway.solo.io.RouteTableSpec.routes:type_name -> gateway.solo.io.Route
	5, // 1: gateway.solo.io.RouteTableSpec.weight:type_name -> google.protobuf.Int32Value
	0, // 2: gateway.solo.io.RouteTableStatus.state:type_name -> gateway.solo.io.RouteTableStatus.State
	3, // 3: gateway.solo.io.RouteTableStatus.subresource_statuses:type_name -> gateway.solo.io.RouteTableStatus.SubresourceStatusesEntry
	6, // 4: gateway.solo.io.RouteTableStatus.details:type_name -> google.protobuf.Struct
	2, // 5: gateway.solo.io.RouteTableStatus.SubresourceStatusesEntry.value:type_name -> gateway.solo.io.RouteTableStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_init() }
func file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto != nil {
		return
	}
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_virtual_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableSpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableStatus); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gateway_v1_route_table_proto_depIdxs = nil
}
