// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/waf/waf.proto

package waf

import (
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type AuditLogging_AuditLogAction int32

const (
	// Never generate audit logs.
	AuditLogging_NEVER AuditLogging_AuditLogAction = 0
	// When set to RELEVANT_ONLY, this will have similar behavior to `SecAuditEngine RelevantOnly`.
	AuditLogging_RELEVANT_ONLY AuditLogging_AuditLogAction = 1
	// Always generate an audit log entry (as long as the filter is not disabled).
	AuditLogging_ALWAYS AuditLogging_AuditLogAction = 2
)

// Enum value maps for AuditLogging_AuditLogAction.
var (
	AuditLogging_AuditLogAction_name = map[int32]string{
		0: "NEVER",
		1: "RELEVANT_ONLY",
		2: "ALWAYS",
	}
	AuditLogging_AuditLogAction_value = map[string]int32{
		"NEVER":         0,
		"RELEVANT_ONLY": 1,
		"ALWAYS":        2,
	}
)

func (x AuditLogging_AuditLogAction) Enum() *AuditLogging_AuditLogAction {
	p := new(AuditLogging_AuditLogAction)
	*p = x
	return p
}

func (x AuditLogging_AuditLogAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuditLogging_AuditLogAction) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_enumTypes[0].Descriptor()
}

func (AuditLogging_AuditLogAction) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_enumTypes[0]
}

func (x AuditLogging_AuditLogAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuditLogging_AuditLogAction.Descriptor instead.
func (AuditLogging_AuditLogAction) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescGZIP(), []int{0, 0}
}

type AuditLogging_AuditLogLocation int32

const (
	// Add the audit log to the filter state.
	// it will be under the key "io.solo.modsecurity.audit_log".
	// You can use this formatter in the access log:
	// %FILTER_STATE(io.solo.modsecurity.audit_log)%
	AuditLogging_FILTER_STATE AuditLogging_AuditLogLocation = 0
	// Add the audit log to the dynamic metadata.
	// it will be under the filter name "io.solo.filters.http.modsecurity". with "audit_log" as the
	// key. You can use this formatter in the access log:
	// %DYNAMIC_METADATA("io.solo.filters.http.modsecurity:audit_log")%
	AuditLogging_DYNAMIC_METADATA AuditLogging_AuditLogLocation = 1
)

// Enum value maps for AuditLogging_AuditLogLocation.
var (
	AuditLogging_AuditLogLocation_name = map[int32]string{
		0: "FILTER_STATE",
		1: "DYNAMIC_METADATA",
	}
	AuditLogging_AuditLogLocation_value = map[string]int32{
		"FILTER_STATE":     0,
		"DYNAMIC_METADATA": 1,
	}
)

func (x AuditLogging_AuditLogLocation) Enum() *AuditLogging_AuditLogLocation {
	p := new(AuditLogging_AuditLogLocation)
	*p = x
	return p
}

func (x AuditLogging_AuditLogLocation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuditLogging_AuditLogLocation) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_enumTypes[1].Descriptor()
}

func (AuditLogging_AuditLogLocation) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_enumTypes[1]
}

func (x AuditLogging_AuditLogLocation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuditLogging_AuditLogLocation.Descriptor instead.
func (AuditLogging_AuditLogLocation) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescGZIP(), []int{0, 1}
}

type AuditLogging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action   AuditLogging_AuditLogAction   `protobuf:"varint,1,opt,name=action,proto3,enum=envoy.config.filter.http.modsecurity.v2.AuditLogging_AuditLogAction" json:"action,omitempty"`
	Location AuditLogging_AuditLogLocation `protobuf:"varint,2,opt,name=location,proto3,enum=envoy.config.filter.http.modsecurity.v2.AuditLogging_AuditLogLocation" json:"location,omitempty"`
}

func (x *AuditLogging) Reset() {
	*x = AuditLogging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLogging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLogging) ProtoMessage() {}

func (x *AuditLogging) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLogging.ProtoReflect.Descriptor instead.
func (*AuditLogging) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescGZIP(), []int{0}
}

func (x *AuditLogging) GetAction() AuditLogging_AuditLogAction {
	if x != nil {
		return x.Action
	}
	return AuditLogging_NEVER
}

func (x *AuditLogging) GetLocation() AuditLogging_AuditLogLocation {
	if x != nil {
		return x.Location
	}
	return AuditLogging_FILTER_STATE
}

type ModSecurity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Disable all rules on the current route
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Global rule sets for the current http connection manager
	RuleSets []*RuleSet `protobuf:"bytes,2,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	// Custom message to display when an intervention occurs
	CustomInterventionMessage string `protobuf:"bytes,3,opt,name=custom_intervention_message,json=customInterventionMessage,proto3" json:"custom_intervention_message,omitempty"`
	// This instructs the filter what to do with the transaction's audit log.
	AuditLogging *AuditLogging `protobuf:"bytes,5,opt,name=audit_logging,json=auditLogging,proto3" json:"audit_logging,omitempty"`
	// If set, the body will not be buffered and fed to ModSecurity. Only the headers will.
	// This can help improve performance.
	RequestHeadersOnly  bool `protobuf:"varint,6,opt,name=request_headers_only,json=requestHeadersOnly,proto3" json:"request_headers_only,omitempty"`
	ResponseHeadersOnly bool `protobuf:"varint,7,opt,name=response_headers_only,json=responseHeadersOnly,proto3" json:"response_headers_only,omitempty"`
	// log in a format suited for the OWASP regression tests.
	// this format is a multiline log format, so it is disabled for regular use.
	// do not enable this in production!
	RegressionLogs bool `protobuf:"varint,4,opt,name=regression_logs,json=regressionLogs,proto3" json:"regression_logs,omitempty"`
}

func (x *ModSecurity) Reset() {
	*x = ModSecurity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModSecurity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModSecurity) ProtoMessage() {}

func (x *ModSecurity) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModSecurity.ProtoReflect.Descriptor instead.
func (*ModSecurity) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescGZIP(), []int{1}
}

func (x *ModSecurity) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *ModSecurity) GetRuleSets() []*RuleSet {
	if x != nil {
		return x.RuleSets
	}
	return nil
}

func (x *ModSecurity) GetCustomInterventionMessage() string {
	if x != nil {
		return x.CustomInterventionMessage
	}
	return ""
}

func (x *ModSecurity) GetAuditLogging() *AuditLogging {
	if x != nil {
		return x.AuditLogging
	}
	return nil
}

func (x *ModSecurity) GetRequestHeadersOnly() bool {
	if x != nil {
		return x.RequestHeadersOnly
	}
	return false
}

func (x *ModSecurity) GetResponseHeadersOnly() bool {
	if x != nil {
		return x.ResponseHeadersOnly
	}
	return false
}

func (x *ModSecurity) GetRegressionLogs() bool {
	if x != nil {
		return x.RegressionLogs
	}
	return false
}

type RuleSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String of rules which are added directly
	RuleStr string `protobuf:"bytes,1,opt,name=rule_str,json=ruleStr,proto3" json:"rule_str,omitempty"`
	// Array of files to include
	Files []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
	// A directory to include. all *.conf files in this directory will be
	// included. sub directories will NOT be checked.
	Directory string `protobuf:"bytes,4,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *RuleSet) Reset() {
	*x = RuleSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleSet) ProtoMessage() {}

func (x *RuleSet) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleSet.ProtoReflect.Descriptor instead.
func (*RuleSet) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescGZIP(), []int{2}
}

func (x *RuleSet) GetRuleStr() string {
	if x != nil {
		return x.RuleStr
	}
	return ""
}

func (x *RuleSet) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *RuleSet) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type ModSecurityPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Disable all rules on the current route
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Overwrite the global rules on this route
	RuleSets []*RuleSet `protobuf:"bytes,2,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	// Custom message to display when an intervention occurs
	CustomInterventionMessage string `protobuf:"bytes,3,opt,name=custom_intervention_message,json=customInterventionMessage,proto3" json:"custom_intervention_message,omitempty"`
	// This instructs the filter what to do with the transaction's audit log.
	AuditLogging *AuditLogging `protobuf:"bytes,5,opt,name=audit_logging,json=auditLogging,proto3" json:"audit_logging,omitempty"`
	// If set, the body will not be buffered and fed to ModSecurity. Only the headers will.
	// This can help improve performance.
	RequestHeadersOnly  bool `protobuf:"varint,6,opt,name=request_headers_only,json=requestHeadersOnly,proto3" json:"request_headers_only,omitempty"`
	ResponseHeadersOnly bool `protobuf:"varint,7,opt,name=response_headers_only,json=responseHeadersOnly,proto3" json:"response_headers_only,omitempty"`
}

func (x *ModSecurityPerRoute) Reset() {
	*x = ModSecurityPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModSecurityPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModSecurityPerRoute) ProtoMessage() {}

func (x *ModSecurityPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModSecurityPerRoute.ProtoReflect.Descriptor instead.
func (*ModSecurityPerRoute) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescGZIP(), []int{3}
}

func (x *ModSecurityPerRoute) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *ModSecurityPerRoute) GetRuleSets() []*RuleSet {
	if x != nil {
		return x.RuleSets
	}
	return nil
}

func (x *ModSecurityPerRoute) GetCustomInterventionMessage() string {
	if x != nil {
		return x.CustomInterventionMessage
	}
	return ""
}

func (x *ModSecurityPerRoute) GetAuditLogging() *AuditLogging {
	if x != nil {
		return x.AuditLogging
	}
	return nil
}

func (x *ModSecurityPerRoute) GetRequestHeadersOnly() bool {
	if x != nil {
		return x.RequestHeadersOnly
	}
	return false
}

func (x *ModSecurityPerRoute) GetResponseHeadersOnly() bool {
	if x != nil {
		return x.ResponseHeadersOnly
	}
	return false
}

var File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDesc = []byte{
	0x0a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x61, 0x66, 0x2f, 0x77, 0x61, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6d,
	0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x12, 0x5c, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x44, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c,
	0x6f, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x62, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x46, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c,
	0x6f, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x0e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x45, 0x56, 0x45, 0x52, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x4c, 0x45, 0x56, 0x41, 0x4e, 0x54, 0x5f, 0x4f, 0x4e,
	0x4c, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4c, 0x57, 0x41, 0x59, 0x53, 0x10, 0x02,
	0x22, 0x3a, 0x0a, 0x10, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49,
	0x43, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x22, 0xa3, 0x03, 0x0a,
	0x0b, 0x4d, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x52, 0x08, 0x72,
	0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x1b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x67, 0x73, 0x22, 0x58, 0x0a, 0x07, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x82, 0x03, 0x0a,
	0x13, 0x4d, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x65, 0x72, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x4d, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6d,
	0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x73, 0x12,
	0x3e, 0x0a, 0x1b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x5a, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x14, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x32, 0x0a,
	0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c,
	0x79, 0x42, 0x59, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x77, 0x61, 0x66, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_goTypes = []interface{}{
	(AuditLogging_AuditLogAction)(0),   // 0: envoy.config.filter.http.modsecurity.v2.AuditLogging.AuditLogAction
	(AuditLogging_AuditLogLocation)(0), // 1: envoy.config.filter.http.modsecurity.v2.AuditLogging.AuditLogLocation
	(*AuditLogging)(nil),               // 2: envoy.config.filter.http.modsecurity.v2.AuditLogging
	(*ModSecurity)(nil),                // 3: envoy.config.filter.http.modsecurity.v2.ModSecurity
	(*RuleSet)(nil),                    // 4: envoy.config.filter.http.modsecurity.v2.RuleSet
	(*ModSecurityPerRoute)(nil),        // 5: envoy.config.filter.http.modsecurity.v2.ModSecurityPerRoute
}
var file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_depIdxs = []int32{
	0, // 0: envoy.config.filter.http.modsecurity.v2.AuditLogging.action:type_name -> envoy.config.filter.http.modsecurity.v2.AuditLogging.AuditLogAction
	1, // 1: envoy.config.filter.http.modsecurity.v2.AuditLogging.location:type_name -> envoy.config.filter.http.modsecurity.v2.AuditLogging.AuditLogLocation
	4, // 2: envoy.config.filter.http.modsecurity.v2.ModSecurity.rule_sets:type_name -> envoy.config.filter.http.modsecurity.v2.RuleSet
	2, // 3: envoy.config.filter.http.modsecurity.v2.ModSecurity.audit_logging:type_name -> envoy.config.filter.http.modsecurity.v2.AuditLogging
	4, // 4: envoy.config.filter.http.modsecurity.v2.ModSecurityPerRoute.rule_sets:type_name -> envoy.config.filter.http.modsecurity.v2.RuleSet
	2, // 5: envoy.config.filter.http.modsecurity.v2.ModSecurityPerRoute.audit_logging:type_name -> envoy.config.filter.http.modsecurity.v2.AuditLogging
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLogging); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModSecurity); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleSet); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModSecurityPerRoute); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_gloo_external_envoy_extensions_waf_waf_proto_depIdxs = nil
}
