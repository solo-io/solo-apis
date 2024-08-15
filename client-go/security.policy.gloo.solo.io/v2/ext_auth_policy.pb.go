// Set up an external authentication and authorization to protect the workloads in your cluster.
// For example, you can set up basic, passthrough, API key, OAuth, OPA, or LDAP authentication.
//
// To enforce external auth, each workspace must have an external auth server.
// You can select which external auth service is used to enforce the policy by configuring an ExtAuthServer resource.
// If you don't configure this resource, then the policy uses the default ext-auth-service that was created for you at installation.
// If you did not enable the external-auth-service during installation, the policy cannot be enforced.
// For more information, see the [ExtAuthServer API reference]({{% link path="/reference/api/ext_auth_server/" %}}).
//
// The external auth policies tell the ext-auth-service how to enforce authentication and authorization of matching traffic.
// You can match traffic by selecting routes or destinations.
// If you don't select any routes or destinations, an external auth policy is applied to all destinations by default.
// You cannot apply more than one external auth policy to the same route or destination.
// Note that when you create the policy with a destination selector, only Kubernetes services can be specified in the
// `applyToDestination` section. Virtual destinations or external services are not supported.
//
// For more information, see the [external auth guides]({{% link path="/security/external-auth/" %}}).
//
// </br>**Example**: The following basic example refers to the Kubernetes service of the default ext-auth-service that you got during installation,
// but you can also use a virtual destination that points to an external auth service instead.
// ```yaml
// {{% readfile file="/static/content/examples/generated/int/ext_auth_destinations/cluster-1/ext-auth-policy_bookinfo_basic-auth.yaml" %}}
// ---
// {{% readfile file="/static/content/examples/generated/int/ext_auth_external_service/cluster-1/ext-auth-server_bookinfo_default-server.yaml" %}}
// ```
//
// For more complex examples, see the [guide]({{% link path="/security/external-auth/" %}}) for the type of external auth policy that you want to apply.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/policy/v2/security/ext_auth_policy.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
	v1 "github.com/solo-io/solo-apis/client-go/enterprise.gloo.solo.io/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specifications for the policy.
type ExtAuthPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Routes to apply the policy to.
	// If empty, the policy does not apply to any routes in the workspace.
	//
	// </br>**Configuration constraints**: Only one external auth policy can apply to a route.
	// Subsequent policies (sorted by creation time) are ignored and put into a FAILED state.
	ApplyToRoutes []*v2.RouteSelector `protobuf:"bytes,1,rep,name=apply_to_routes,json=applyToRoutes,proto3" json:"apply_to_routes,omitempty"`
	// Destinations to apply the policy to.
	//
	// </br>**Implementation notes**:
	// <li>Only Kubernetes services can be specified. Virtual destinations or external services are not supported.</li>
	// <li>If empty and `applyToRoutes` is unset, the policy applies to all destinations in the workspace.</li>
	// <li>If empty and `applyToRoutes` is set, the policy does not apply to any destinations in the workspace.</li></ul>
	//
	// </br>**Configuration constraints**: Only one external auth policy can apply to a destination.
	// Subsequent policies (sorted by creation time) are ignored and put into a FAILED state.
	ApplyToDestinations []*v2.DestinationSelector `protobuf:"bytes,2,rep,name=apply_to_destinations,json=applyToDestinations,proto3" json:"apply_to_destinations,omitempty"`
	// The details of the external auth policy to apply to the selected routes and/or destinations.
	Config *ExtAuthPolicySpec_Config `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ExtAuthPolicySpec) Reset() {
	*x = ExtAuthPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtAuthPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtAuthPolicySpec) ProtoMessage() {}

func (x *ExtAuthPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtAuthPolicySpec.ProtoReflect.Descriptor instead.
func (*ExtAuthPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescGZIP(), []int{0}
}

func (x *ExtAuthPolicySpec) GetApplyToRoutes() []*v2.RouteSelector {
	if x != nil {
		return x.ApplyToRoutes
	}
	return nil
}

func (x *ExtAuthPolicySpec) GetApplyToDestinations() []*v2.DestinationSelector {
	if x != nil {
		return x.ApplyToDestinations
	}
	return nil
}

func (x *ExtAuthPolicySpec) GetConfig() *ExtAuthPolicySpec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// The status of the policy after it is applied to your Gloo environment.
type ExtAuthPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state and workspace conditions of the applied resource.
	Common *v2.Status `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// The number of destination ports selected by the policy.
	NumSelectedDestinationPorts uint32 `protobuf:"varint,2,opt,name=num_selected_destination_ports,json=numSelectedDestinationPorts,proto3" json:"num_selected_destination_ports,omitempty"`
	// The number of routes selected by the policy.
	NumSelectedRoutes uint32 `protobuf:"varint,3,opt,name=num_selected_routes,json=numSelectedRoutes,proto3" json:"num_selected_routes,omitempty"`
}

func (x *ExtAuthPolicyStatus) Reset() {
	*x = ExtAuthPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtAuthPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtAuthPolicyStatus) ProtoMessage() {}

func (x *ExtAuthPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtAuthPolicyStatus.ProtoReflect.Descriptor instead.
func (*ExtAuthPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescGZIP(), []int{1}
}

func (x *ExtAuthPolicyStatus) GetCommon() *v2.Status {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *ExtAuthPolicyStatus) GetNumSelectedDestinationPorts() uint32 {
	if x != nil {
		return x.NumSelectedDestinationPorts
	}
	return 0
}

func (x *ExtAuthPolicyStatus) GetNumSelectedRoutes() uint32 {
	if x != nil {
		return x.NumSelectedRoutes
	}
	return 0
}

type ExtAuthPolicyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the resource in each workspace that it exists in.
	Workspaces map[string]*v2.Report `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of destination ports selected by the policy.
	SelectedDestinationPorts []*v2.DestinationReference `protobuf:"bytes,2,rep,name=selected_destination_ports,json=selectedDestinationPorts,proto3" json:"selected_destination_ports,omitempty"`
	// A list of references to all route selected by the policy.
	SelectedRoutes []*v2.RouteReference `protobuf:"bytes,3,rep,name=selected_routes,json=selectedRoutes,proto3" json:"selected_routes,omitempty"`
}

func (x *ExtAuthPolicyReport) Reset() {
	*x = ExtAuthPolicyReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtAuthPolicyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtAuthPolicyReport) ProtoMessage() {}

func (x *ExtAuthPolicyReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtAuthPolicyReport.ProtoReflect.Descriptor instead.
func (*ExtAuthPolicyReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescGZIP(), []int{2}
}

func (x *ExtAuthPolicyReport) GetWorkspaces() map[string]*v2.Report {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

func (x *ExtAuthPolicyReport) GetSelectedDestinationPorts() []*v2.DestinationReference {
	if x != nil {
		return x.SelectedDestinationPorts
	}
	return nil
}

func (x *ExtAuthPolicyReport) GetSelectedRoutes() []*v2.RouteReference {
	if x != nil {
		return x.SelectedRoutes
	}
	return nil
}

// The details of the external auth policy to apply to the selected routes or destinations.
//
// +kubebuilder:validation:XValidation:rule="!has(self.disable)",message="The disable field is currently unsupported"
type ExtAuthPolicySpec_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AuthType:
	//
	//	*ExtAuthPolicySpec_Config_Disable
	//	*ExtAuthPolicySpec_Config_GlooAuth
	//	*ExtAuthPolicySpec_Config_CustomAuth_
	AuthType isExtAuthPolicySpec_Config_AuthType `protobuf_oneof:"auth_type"`
	// Reference to the ExtAuthServer to use for this policy.
	// To create an ExtAuthServer resource, see [External auth server setup](../../../security/external-auth/server/#setup).
	//
	// </br>**Implementation notes**:<ul>
	// <li>The ExtAuthServer must be in the same workspace as this ExtAuthPolicy. Be sure to select the appropriate ExtAuthServer to use,
	// which might be in a different cluster and namespace than the policy.</li>
	// <li>If this field is omitted, Gloo generates a default ExtAuthServer configuration.
	// To review the default values for an ExtAuthServer, see the [API reference](../ext_auth_server/).</li></ul>
	//
	// </br>**Configuration constraints:**<ul>
	// <li>Routes on a single gateway must share a single ExtAuthServer.</li>
	// <li>If `glooAuth` is used, the `server` must reference an ExtAuthServer that is backed by one or more Gloo `ext-auth-service` Kubernetes services.</li>
	// <li>If `customAuth` is used, the `server` must reference an ExtAuthServer that is backed by your own custom external auth server.</li>
	// <li>For auth configurations that require a client secret from the identity provider issuer,
	// the secret must be in the same cluster as the ExtAuthServer resource.</li></ul>
	Server *v2.ObjectReference `protobuf:"bytes,4,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *ExtAuthPolicySpec_Config) Reset() {
	*x = ExtAuthPolicySpec_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtAuthPolicySpec_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtAuthPolicySpec_Config) ProtoMessage() {}

func (x *ExtAuthPolicySpec_Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtAuthPolicySpec_Config.ProtoReflect.Descriptor instead.
func (*ExtAuthPolicySpec_Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ExtAuthPolicySpec_Config) GetAuthType() isExtAuthPolicySpec_Config_AuthType {
	if m != nil {
		return m.AuthType
	}
	return nil
}

// Deprecated: Do not use.
func (x *ExtAuthPolicySpec_Config) GetDisable() bool {
	if x, ok := x.GetAuthType().(*ExtAuthPolicySpec_Config_Disable); ok {
		return x.Disable
	}
	return false
}

func (x *ExtAuthPolicySpec_Config) GetGlooAuth() *v1.AuthConfigSpec {
	if x, ok := x.GetAuthType().(*ExtAuthPolicySpec_Config_GlooAuth); ok {
		return x.GlooAuth
	}
	return nil
}

func (x *ExtAuthPolicySpec_Config) GetCustomAuth() *ExtAuthPolicySpec_Config_CustomAuth {
	if x, ok := x.GetAuthType().(*ExtAuthPolicySpec_Config_CustomAuth_); ok {
		return x.CustomAuth
	}
	return nil
}

func (x *ExtAuthPolicySpec_Config) GetServer() *v2.ObjectReference {
	if x != nil {
		return x.Server
	}
	return nil
}

type isExtAuthPolicySpec_Config_AuthType interface {
	isExtAuthPolicySpec_Config_AuthType()
}

type ExtAuthPolicySpec_Config_Disable struct {
	// <!-- Set to true to disable auth on the route.-->
	// **DEPRECATED**: This field is currently unsupported and cannot be used.
	//
	// Deprecated: Do not use.
	Disable bool `protobuf:"varint,1,opt,name=disable,proto3,oneof"`
}

type ExtAuthPolicySpec_Config_GlooAuth struct {
	// Configure the selected route or destination with auth options provided by one or more Gloo `ext-auth-service` Kubernetes services.
	GlooAuth *v1.AuthConfigSpec `protobuf:"bytes,2,opt,name=gloo_auth,json=glooAuth,proto3,oneof"`
}

type ExtAuthPolicySpec_Config_CustomAuth_ struct {
	// Configure the selected route or destination with auth options provided by your own custom external auth server.
	CustomAuth *ExtAuthPolicySpec_Config_CustomAuth `protobuf:"bytes,3,opt,name=custom_auth,json=customAuth,proto3,oneof"`
}

func (*ExtAuthPolicySpec_Config_Disable) isExtAuthPolicySpec_Config_AuthType() {}

func (*ExtAuthPolicySpec_Config_GlooAuth) isExtAuthPolicySpec_Config_AuthType() {}

func (*ExtAuthPolicySpec_Config_CustomAuth_) isExtAuthPolicySpec_Config_AuthType() {}

// Configure the selected route or destination with auth options provided by your own external auth server.
type ExtAuthPolicySpec_Config_CustomAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Add context extensions to the request that matches on the route or destination and is sent to the external auth server.
	// This allows the server to base the auth decision on metadata that you define on the source of the request.
	//
	// This attribute is analogous to Envoy's config.filter.http.ext_authz.v2.CheckSettings. For more info, see the
	// [Envoy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/ext_authz_filter.html?highlight=extauthz#config-filter-http-ext-authz-v2-checksettings).
	ContextExtensions map[string]string `protobuf:"bytes,1,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExtAuthPolicySpec_Config_CustomAuth) Reset() {
	*x = ExtAuthPolicySpec_Config_CustomAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtAuthPolicySpec_Config_CustomAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtAuthPolicySpec_Config_CustomAuth) ProtoMessage() {}

func (x *ExtAuthPolicySpec_Config_CustomAuth) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtAuthPolicySpec_Config_CustomAuth.ProtoReflect.Descriptor instead.
func (*ExtAuthPolicySpec_Config_CustomAuth) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ExtAuthPolicySpec_Config_CustomAuth) GetContextExtensions() map[string]string {
	if x != nil {
		return x.ContextExtensions
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDesc = []byte{
	0x0a, 0x64, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2f, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x90, 0x06, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4a, 0x0a, 0x0f, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x15, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74,
	0x6f, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x13,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x80, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e,
	0x0a, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x46,
	0x0a, 0x09, 0x67, 0x6c, 0x6f, 0x6f, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x08, 0x67, 0x6c,
	0x6f, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x64, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x12, 0x3c, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0xdc, 0x01, 0x0a, 0x0a, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x12, 0x87, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x58, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x44, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x1e, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x6e, 0x75, 0x6d,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x75, 0x6d, 0x5f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x8b, 0x03, 0x0a, 0x13, 0x45, 0x78, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x61, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x67, 0x0a, 0x1a, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x18, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x5a, 0x0a, 0x0f, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x60, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04, 0x01,
	0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_goTypes = []interface{}{
	(*ExtAuthPolicySpec)(nil),                   // 0: security.policy.gloo.solo.io.ExtAuthPolicySpec
	(*ExtAuthPolicyStatus)(nil),                 // 1: security.policy.gloo.solo.io.ExtAuthPolicyStatus
	(*ExtAuthPolicyReport)(nil),                 // 2: security.policy.gloo.solo.io.ExtAuthPolicyReport
	(*ExtAuthPolicySpec_Config)(nil),            // 3: security.policy.gloo.solo.io.ExtAuthPolicySpec.Config
	(*ExtAuthPolicySpec_Config_CustomAuth)(nil), // 4: security.policy.gloo.solo.io.ExtAuthPolicySpec.Config.CustomAuth
	nil,                             // 5: security.policy.gloo.solo.io.ExtAuthPolicySpec.Config.CustomAuth.ContextExtensionsEntry
	nil,                             // 6: security.policy.gloo.solo.io.ExtAuthPolicyReport.WorkspacesEntry
	(*v2.RouteSelector)(nil),        // 7: common.gloo.solo.io.RouteSelector
	(*v2.DestinationSelector)(nil),  // 8: common.gloo.solo.io.DestinationSelector
	(*v2.Status)(nil),               // 9: common.gloo.solo.io.Status
	(*v2.DestinationReference)(nil), // 10: common.gloo.solo.io.DestinationReference
	(*v2.RouteReference)(nil),       // 11: common.gloo.solo.io.RouteReference
	(*v1.AuthConfigSpec)(nil),       // 12: enterprise.gloo.solo.io.AuthConfigSpec
	(*v2.ObjectReference)(nil),      // 13: common.gloo.solo.io.ObjectReference
	(*v2.Report)(nil),               // 14: common.gloo.solo.io.Report
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_depIdxs = []int32{
	7,  // 0: security.policy.gloo.solo.io.ExtAuthPolicySpec.apply_to_routes:type_name -> common.gloo.solo.io.RouteSelector
	8,  // 1: security.policy.gloo.solo.io.ExtAuthPolicySpec.apply_to_destinations:type_name -> common.gloo.solo.io.DestinationSelector
	3,  // 2: security.policy.gloo.solo.io.ExtAuthPolicySpec.config:type_name -> security.policy.gloo.solo.io.ExtAuthPolicySpec.Config
	9,  // 3: security.policy.gloo.solo.io.ExtAuthPolicyStatus.common:type_name -> common.gloo.solo.io.Status
	6,  // 4: security.policy.gloo.solo.io.ExtAuthPolicyReport.workspaces:type_name -> security.policy.gloo.solo.io.ExtAuthPolicyReport.WorkspacesEntry
	10, // 5: security.policy.gloo.solo.io.ExtAuthPolicyReport.selected_destination_ports:type_name -> common.gloo.solo.io.DestinationReference
	11, // 6: security.policy.gloo.solo.io.ExtAuthPolicyReport.selected_routes:type_name -> common.gloo.solo.io.RouteReference
	12, // 7: security.policy.gloo.solo.io.ExtAuthPolicySpec.Config.gloo_auth:type_name -> enterprise.gloo.solo.io.AuthConfigSpec
	4,  // 8: security.policy.gloo.solo.io.ExtAuthPolicySpec.Config.custom_auth:type_name -> security.policy.gloo.solo.io.ExtAuthPolicySpec.Config.CustomAuth
	13, // 9: security.policy.gloo.solo.io.ExtAuthPolicySpec.Config.server:type_name -> common.gloo.solo.io.ObjectReference
	5,  // 10: security.policy.gloo.solo.io.ExtAuthPolicySpec.Config.CustomAuth.context_extensions:type_name -> security.policy.gloo.solo.io.ExtAuthPolicySpec.Config.CustomAuth.ContextExtensionsEntry
	14, // 11: security.policy.gloo.solo.io.ExtAuthPolicyReport.WorkspacesEntry.value:type_name -> common.gloo.solo.io.Report
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtAuthPolicySpec); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtAuthPolicyStatus); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtAuthPolicyReport); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtAuthPolicySpec_Config); i {
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
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtAuthPolicySpec_Config_CustomAuth); i {
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
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ExtAuthPolicySpec_Config_Disable)(nil),
		(*ExtAuthPolicySpec_Config_GlooAuth)(nil),
		(*ExtAuthPolicySpec_Config_CustomAuth_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_policy_v2_security_ext_auth_policy_proto_depIdxs = nil
}
