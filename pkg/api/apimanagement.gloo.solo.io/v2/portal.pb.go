// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/apimanagement/v2/portal.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/solo-io/cue/encoding/protobuf/cue"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v2 "github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2"
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

// The Portal resource configures a developer portal that you can use to securely expose your APIs to end users.
// Before you create the Portal, you bundle the APIs that you want to expose into a route table.
// Then, you prepare a usage plan to control access to your APIs
// by applying rate limiting and external auth policies to the routes in the route table.
// For more information, see the [Portal docs](https://docs.solo.io/gloo-gateway/main/portal/).
//
// The following examples show a Portal resource that refers to usage plans
// that are named in the rate limit server config's descriptors.
// ```yaml
// apiVersion: portal.gloo.solo.io/v2
// kind: Portal
//
//	metadata:
//	  name: public-portal
//	  namespace: gloo-mesh
//
// spec:
//
//	portalBackendSelectors:
//	  - selector:
//	      namespace: gloo-mesh
//	      cluster: cluster-1
//	usagePlans:
//	  - name: bronze
//	    displayName: "Bronze Plan"
//	    description: "A basic usage plan"
//	  - name: silver
//	    description: "A better usage plan"
//	  - name: gold
//	    description: "The best usage plan!"
//	apis:
//	  - routeTables:
//	      name: productpage
//	      namespace: bookinfo
//	      cluster: cluster-1
//
// ```
//
// ```yaml
// apiVersion: admin.gloo.solo.io/v2
// kind: RateLimitServerConfig
//
//	metadata:
//	  name: usage-plans
//	  namespace: gloo-mesh-addons
//
// spec:
//
//	destinationServers: [ ] # omitted, server refs
//	raw:
//	  descriptors:
//	    - key: usagePlan
//	      value: bronze
//	      descriptors:
//	        - key: userId
//	          rateLimit:
//	            requestsPerUnit: 50
//	            unit: MINUTE
//	    - key: usagePlan
//	      value: silver
//	      descriptors:
//	        - key: userId
//	          rateLimit:
//	            requestsPerUnit: 200
//	            unit: MINUTE
//	     - key: usagePlan
//	       value: gold
//	       descriptors:
//	         - key: userId
//	           rateLimit:
//	             requestsPerUnit: 1000
//	             unit: MINUTE
//
// ```
type PortalSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workloads where an existing portal backend is running.
	PortalBackendSelector []*v2.WorkloadSelector `protobuf:"bytes,1,rep,name=portal_backend_selector,json=portalBackendSelector,proto3" json:"portal_backend_selector,omitempty"`
	// A list of route tables that have routes to the APIs you want to expose in the developer portal.
	Apis []*API `protobuf:"bytes,2,rep,name=apis,proto3" json:"apis,omitempty"`
	// The usage plans to control access to the APIs that the developer portal exposes.
	UsagePlans []*PortalSpec_UsagePlan `protobuf:"bytes,3,rep,name=usage_plans,json=usagePlans,proto3" json:"usage_plans,omitempty"`
	// The domains on which this Portal will be served. The Host header received by the
	// Portal Web App will be matched to one of these domains in order to determine which Portal will be served.
	//
	// If you are using the Gateway through which you are exposing the Portal
	// is listening on a port other than 80/443, you should include the port as part of the
	// domain string, e.g. "portal.solo.io:8443".
	//
	// To prevent undefined behavior, creating a Portal whose domain conflicts with
	// an existing Portal will result in the Portal resource being placed into an 'Invalid' state.
	Domains []string `protobuf:"bytes,4,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (x *PortalSpec) Reset() {
	*x = PortalSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortalSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortalSpec) ProtoMessage() {}

func (x *PortalSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortalSpec.ProtoReflect.Descriptor instead.
func (*PortalSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescGZIP(), []int{0}
}

func (x *PortalSpec) GetPortalBackendSelector() []*v2.WorkloadSelector {
	if x != nil {
		return x.PortalBackendSelector
	}
	return nil
}

func (x *PortalSpec) GetApis() []*API {
	if x != nil {
		return x.Apis
	}
	return nil
}

func (x *PortalSpec) GetUsagePlans() []*PortalSpec_UsagePlan {
	if x != nil {
		return x.UsagePlans
	}
	return nil
}

func (x *PortalSpec) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

type API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of route tables with routes to the APIs you want the developer portal to expose.
	// The route table might also have 'portalMetadata' key-value fields that you want to display in the developer portal for end users.
	RouteTable *v2.ObjectSelector `protobuf:"bytes,1,opt,name=route_table,json=routeTable,proto3" json:"route_table,omitempty"`
}

func (x *API) Reset() {
	*x = API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*API) ProtoMessage() {}

func (x *API) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use API.ProtoReflect.Descriptor instead.
func (*API) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescGZIP(), []int{1}
}

func (x *API) GetRouteTable() *v2.ObjectSelector {
	if x != nil {
		return x.RouteTable
	}
	return nil
}

// reflects the status of the Portal resource
type PortalStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global *v2.GenericGlobalStatus `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.WorkspaceStatus `protobuf:"bytes,2,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Name of workspace that owns the portal.
	OwnerWorkspace *v2.OwnerWorkspace `protobuf:"bytes,3,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
}

func (x *PortalStatus) Reset() {
	*x = PortalStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortalStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortalStatus) ProtoMessage() {}

func (x *PortalStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortalStatus.ProtoReflect.Descriptor instead.
func (*PortalStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescGZIP(), []int{2}
}

func (x *PortalStatus) GetGlobal() *v2.GenericGlobalStatus {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *PortalStatus) GetWorkspaces() map[string]*v2.WorkspaceStatus {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *PortalStatus) GetOwnerWorkspace() *v2.OwnerWorkspace {
	if x != nil {
		return x.OwnerWorkspace
	}
	return nil
}

// $hide_from_docs
type PortalNewStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common         *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	OwnerWorkspace string     `protobuf:"bytes,2,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
}

func (x *PortalNewStatus) Reset() {
	*x = PortalNewStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortalNewStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortalNewStatus) ProtoMessage() {}

func (x *PortalNewStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortalNewStatus.ProtoReflect.Descriptor instead.
func (*PortalNewStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescGZIP(), []int{3}
}

func (x *PortalNewStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *PortalNewStatus) GetOwnerWorkspace() string {
	if x != nil {
		return x.OwnerWorkspace
	}
	return ""
}

// $hide_from_docs
type PortalReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspaces     map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OwnerWorkspace string                `protobuf:"bytes,2,opt,name=owner_workspace,json=ownerWorkspace,proto3" json:"owner_workspace,omitempty"`
}

func (x *PortalReport) Reset() {
	*x = PortalReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortalReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortalReport) ProtoMessage() {}

func (x *PortalReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortalReport.ProtoReflect.Descriptor instead.
func (*PortalReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescGZIP(), []int{4}
}

func (x *PortalReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *PortalReport) GetOwnerWorkspace() string {
	if x != nil {
		return x.OwnerWorkspace
	}
	return ""
}

type PortalSpec_UsagePlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Match the names of the usage plans with the descriptors that you defined in the rate limit server config.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional display name for the usage plan to show end users in the developer portal.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional description for the usage plan to show end users in the developer portal.
	// You might include information about how to get the plan or what the plan includes and excludes.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PortalSpec_UsagePlan) Reset() {
	*x = PortalSpec_UsagePlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortalSpec_UsagePlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortalSpec_UsagePlan) ProtoMessage() {}

func (x *PortalSpec_UsagePlan) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortalSpec_UsagePlan.ProtoReflect.Descriptor instead.
func (*PortalSpec_UsagePlan) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PortalSpec_UsagePlan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PortalSpec_UsagePlan) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *PortalSpec_UsagePlan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDesc = []byte{
	0x0a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x1a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x75, 0x65, 0x2f, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63,
	0x75, 0x65, 0x2f, 0x63, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x65,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x02, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x5d, 0x0a, 0x17, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x15, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x61, 0x70, 0x69, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x04, 0x61, 0x70, 0x69, 0x73, 0x12, 0x51, 0x0a, 0x0b, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x0a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x1a, 0x64, 0x0a, 0x09, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x03, 0x41,
	0x50, 0x49, 0x12, 0x44, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xdd, 0x02, 0x0a, 0x0c, 0x50, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x58, 0x0a, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x1a, 0x63, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6f, 0x0a, 0x0f, 0x50, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xed, 0x01, 0x0a, 0x0c, 0x50, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x58, 0x0a, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x5a, 0x0a,
	0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x50, 0x5a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0,
	0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_goTypes = []interface{}{
	(*PortalSpec)(nil),             // 0: apimanagement.gloo.solo.io.PortalSpec
	(*API)(nil),                    // 1: apimanagement.gloo.solo.io.API
	(*PortalStatus)(nil),           // 2: apimanagement.gloo.solo.io.PortalStatus
	(*PortalNewStatus)(nil),        // 3: apimanagement.gloo.solo.io.PortalNewStatus
	(*PortalReport)(nil),           // 4: apimanagement.gloo.solo.io.PortalReport
	(*PortalSpec_UsagePlan)(nil),   // 5: apimanagement.gloo.solo.io.PortalSpec.UsagePlan
	nil,                            // 6: apimanagement.gloo.solo.io.PortalStatus.WorkspacesEntry
	nil,                            // 7: apimanagement.gloo.solo.io.PortalReport.WorkspacesEntry
	(*v2.WorkloadSelector)(nil),    // 8: common.gloo.solo.io.WorkloadSelector
	(*v2.ObjectSelector)(nil),      // 9: common.gloo.solo.io.ObjectSelector
	(*v2.GenericGlobalStatus)(nil), // 10: common.gloo.solo.io.GenericGlobalStatus
	(*v2.OwnerWorkspace)(nil),      // 11: common.gloo.solo.io.OwnerWorkspace
	(*v2.Status)(nil),              // 12: common.gloo.solo.io.Status
	(*v2.WorkspaceStatus)(nil),     // 13: common.gloo.solo.io.WorkspaceStatus
	(*v2.Report)(nil),              // 14: common.gloo.solo.io.Report
}
var file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_depIdxs = []int32{
	8,  // 0: apimanagement.gloo.solo.io.PortalSpec.portal_backend_selector:type_name -> common.gloo.solo.io.WorkloadSelector
	1,  // 1: apimanagement.gloo.solo.io.PortalSpec.apis:type_name -> apimanagement.gloo.solo.io.API
	5,  // 2: apimanagement.gloo.solo.io.PortalSpec.usage_plans:type_name -> apimanagement.gloo.solo.io.PortalSpec.UsagePlan
	9,  // 3: apimanagement.gloo.solo.io.API.route_table:type_name -> common.gloo.solo.io.ObjectSelector
	10, // 4: apimanagement.gloo.solo.io.PortalStatus.global:type_name -> common.gloo.solo.io.GenericGlobalStatus
	6,  // 5: apimanagement.gloo.solo.io.PortalStatus.workspaces:type_name -> apimanagement.gloo.solo.io.PortalStatus.WorkspacesEntry
	11, // 6: apimanagement.gloo.solo.io.PortalStatus.owner_workspace:type_name -> common.gloo.solo.io.OwnerWorkspace
	12, // 7: apimanagement.gloo.solo.io.PortalNewStatus.common:type_name -> common.gloo.solo.io.Status
	7,  // 8: apimanagement.gloo.solo.io.PortalReport.workspaces:type_name -> apimanagement.gloo.solo.io.PortalReport.WorkspacesEntry
	13, // 9: apimanagement.gloo.solo.io.PortalStatus.WorkspacesEntry.value:type_name -> common.gloo.solo.io.WorkspaceStatus
	14, // 10: apimanagement.gloo.solo.io.PortalReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortalSpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*API); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortalStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortalNewStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortalReport); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortalSpec_UsagePlan); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_mesh_gloo_solo_io_apimanagement_v2_portal_proto_depIdxs = nil
}
