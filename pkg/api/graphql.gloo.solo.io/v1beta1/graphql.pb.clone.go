// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/graphql.gloo/v1beta1/graphql.proto

package v1beta1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

	google_golang_org_protobuf_types_known_structpb "google.golang.org/protobuf/types/known/structpb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *RequestTemplate) Clone() proto.Message {
	var target *RequestTemplate
	if m == nil {
		return target
	}
	target = &RequestTemplate{}

	if m.GetHeaders() != nil {
		target.Headers = make(map[string]string, len(m.GetHeaders()))
		for k, v := range m.GetHeaders() {

			target.Headers[k] = v

		}
	}

	if m.GetQueryParams() != nil {
		target.QueryParams = make(map[string]string, len(m.GetQueryParams()))
		for k, v := range m.GetQueryParams() {

			target.QueryParams[k] = v

		}
	}

	if h, ok := interface{}(m.GetBody()).(clone.Cloner); ok {
		target.Body = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Value)
	} else {
		target.Body = proto.Clone(m.GetBody()).(*google_golang_org_protobuf_types_known_structpb.Value)
	}

	return target
}

// Clone function
func (m *ResponseTemplate) Clone() proto.Message {
	var target *ResponseTemplate
	if m == nil {
		return target
	}
	target = &ResponseTemplate{}

	target.ResultRoot = m.GetResultRoot()

	if m.GetSetters() != nil {
		target.Setters = make(map[string]string, len(m.GetSetters()))
		for k, v := range m.GetSetters() {

			target.Setters[k] = v

		}
	}

	return target
}

// Clone function
func (m *GrpcRequestTemplate) Clone() proto.Message {
	var target *GrpcRequestTemplate
	if m == nil {
		return target
	}
	target = &GrpcRequestTemplate{}

	if h, ok := interface{}(m.GetOutgoingMessageJson()).(clone.Cloner); ok {
		target.OutgoingMessageJson = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Value)
	} else {
		target.OutgoingMessageJson = proto.Clone(m.GetOutgoingMessageJson()).(*google_golang_org_protobuf_types_known_structpb.Value)
	}

	target.ServiceName = m.GetServiceName()

	target.MethodName = m.GetMethodName()

	if m.GetRequestMetadata() != nil {
		target.RequestMetadata = make(map[string]string, len(m.GetRequestMetadata()))
		for k, v := range m.GetRequestMetadata() {

			target.RequestMetadata[k] = v

		}
	}

	return target
}

// Clone function
func (m *RESTResolver) Clone() proto.Message {
	var target *RESTResolver
	if m == nil {
		return target
	}
	target = &RESTResolver{}

	if h, ok := interface{}(m.GetUpstreamRef()).(clone.Cloner); ok {
		target.UpstreamRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.UpstreamRef = proto.Clone(m.GetUpstreamRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if h, ok := interface{}(m.GetRequest()).(clone.Cloner); ok {
		target.Request = h.Clone().(*RequestTemplate)
	} else {
		target.Request = proto.Clone(m.GetRequest()).(*RequestTemplate)
	}

	if h, ok := interface{}(m.GetResponse()).(clone.Cloner); ok {
		target.Response = h.Clone().(*ResponseTemplate)
	} else {
		target.Response = proto.Clone(m.GetResponse()).(*ResponseTemplate)
	}

	target.SpanName = m.GetSpanName()

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	return target
}

// Clone function
func (m *GrpcDescriptorRegistry) Clone() proto.Message {
	var target *GrpcDescriptorRegistry
	if m == nil {
		return target
	}
	target = &GrpcDescriptorRegistry{}

	switch m.DescriptorSet.(type) {

	case *GrpcDescriptorRegistry_ProtoDescriptor:

		target.DescriptorSet = &GrpcDescriptorRegistry_ProtoDescriptor{
			ProtoDescriptor: m.GetProtoDescriptor(),
		}

	case *GrpcDescriptorRegistry_ProtoDescriptorBin:

		if m.GetProtoDescriptorBin() != nil {
			newArr := make([]byte, len(m.GetProtoDescriptorBin()))
			copy(newArr, m.GetProtoDescriptorBin())
			target.DescriptorSet = &GrpcDescriptorRegistry_ProtoDescriptorBin{
				ProtoDescriptorBin: newArr,
			}
		} else {
			target.DescriptorSet = &GrpcDescriptorRegistry_ProtoDescriptorBin{
				ProtoDescriptorBin: nil,
			}
		}

	case *GrpcDescriptorRegistry_ProtoRefsList:

		if h, ok := interface{}(m.GetProtoRefsList()).(clone.Cloner); ok {
			target.DescriptorSet = &GrpcDescriptorRegistry_ProtoRefsList{
				ProtoRefsList: h.Clone().(*GrpcDescriptorRegistry_ProtoRefs),
			}
		} else {
			target.DescriptorSet = &GrpcDescriptorRegistry_ProtoRefsList{
				ProtoRefsList: proto.Clone(m.GetProtoRefsList()).(*GrpcDescriptorRegistry_ProtoRefs),
			}
		}

	}

	return target
}

// Clone function
func (m *GrpcResolver) Clone() proto.Message {
	var target *GrpcResolver
	if m == nil {
		return target
	}
	target = &GrpcResolver{}

	if h, ok := interface{}(m.GetUpstreamRef()).(clone.Cloner); ok {
		target.UpstreamRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.UpstreamRef = proto.Clone(m.GetUpstreamRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if h, ok := interface{}(m.GetRequestTransform()).(clone.Cloner); ok {
		target.RequestTransform = h.Clone().(*GrpcRequestTemplate)
	} else {
		target.RequestTransform = proto.Clone(m.GetRequestTransform()).(*GrpcRequestTemplate)
	}

	target.SpanName = m.GetSpanName()

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	return target
}

// Clone function
func (m *StitchedSchema) Clone() proto.Message {
	var target *StitchedSchema
	if m == nil {
		return target
	}
	target = &StitchedSchema{}

	if m.GetSubschemas() != nil {
		target.Subschemas = make([]*StitchedSchema_SubschemaConfig, len(m.GetSubschemas()))
		for idx, v := range m.GetSubschemas() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Subschemas[idx] = h.Clone().(*StitchedSchema_SubschemaConfig)
			} else {
				target.Subschemas[idx] = proto.Clone(v).(*StitchedSchema_SubschemaConfig)
			}

		}
	}

	return target
}

// Clone function
func (m *MockResolver) Clone() proto.Message {
	var target *MockResolver
	if m == nil {
		return target
	}
	target = &MockResolver{}

	switch m.Response.(type) {

	case *MockResolver_SyncResponse:

		if h, ok := interface{}(m.GetSyncResponse()).(clone.Cloner); ok {
			target.Response = &MockResolver_SyncResponse{
				SyncResponse: h.Clone().(*google_golang_org_protobuf_types_known_structpb.Value),
			}
		} else {
			target.Response = &MockResolver_SyncResponse{
				SyncResponse: proto.Clone(m.GetSyncResponse()).(*google_golang_org_protobuf_types_known_structpb.Value),
			}
		}

	case *MockResolver_AsyncResponse_:

		if h, ok := interface{}(m.GetAsyncResponse()).(clone.Cloner); ok {
			target.Response = &MockResolver_AsyncResponse_{
				AsyncResponse: h.Clone().(*MockResolver_AsyncResponse),
			}
		} else {
			target.Response = &MockResolver_AsyncResponse_{
				AsyncResponse: proto.Clone(m.GetAsyncResponse()).(*MockResolver_AsyncResponse),
			}
		}

	case *MockResolver_ErrorResponse:

		target.Response = &MockResolver_ErrorResponse{
			ErrorResponse: m.GetErrorResponse(),
		}

	}

	return target
}

// Clone function
func (m *Resolution) Clone() proto.Message {
	var target *Resolution
	if m == nil {
		return target
	}
	target = &Resolution{}

	if h, ok := interface{}(m.GetStatPrefix()).(clone.Cloner); ok {
		target.StatPrefix = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.StatPrefix = proto.Clone(m.GetStatPrefix()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	switch m.Resolver.(type) {

	case *Resolution_RestResolver:

		if h, ok := interface{}(m.GetRestResolver()).(clone.Cloner); ok {
			target.Resolver = &Resolution_RestResolver{
				RestResolver: h.Clone().(*RESTResolver),
			}
		} else {
			target.Resolver = &Resolution_RestResolver{
				RestResolver: proto.Clone(m.GetRestResolver()).(*RESTResolver),
			}
		}

	case *Resolution_GrpcResolver:

		if h, ok := interface{}(m.GetGrpcResolver()).(clone.Cloner); ok {
			target.Resolver = &Resolution_GrpcResolver{
				GrpcResolver: h.Clone().(*GrpcResolver),
			}
		} else {
			target.Resolver = &Resolution_GrpcResolver{
				GrpcResolver: proto.Clone(m.GetGrpcResolver()).(*GrpcResolver),
			}
		}

	case *Resolution_MockResolver:

		if h, ok := interface{}(m.GetMockResolver()).(clone.Cloner); ok {
			target.Resolver = &Resolution_MockResolver{
				MockResolver: h.Clone().(*MockResolver),
			}
		} else {
			target.Resolver = &Resolution_MockResolver{
				MockResolver: proto.Clone(m.GetMockResolver()).(*MockResolver),
			}
		}

	}

	return target
}

// Clone function
func (m *GraphQLApiSpec) Clone() proto.Message {
	var target *GraphQLApiSpec
	if m == nil {
		return target
	}
	target = &GraphQLApiSpec{}

	if h, ok := interface{}(m.GetStatPrefix()).(clone.Cloner); ok {
		target.StatPrefix = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.StatPrefix = proto.Clone(m.GetStatPrefix()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetPersistedQueryCacheConfig()).(clone.Cloner); ok {
		target.PersistedQueryCacheConfig = h.Clone().(*PersistedQueryCacheConfig)
	} else {
		target.PersistedQueryCacheConfig = proto.Clone(m.GetPersistedQueryCacheConfig()).(*PersistedQueryCacheConfig)
	}

	if m.GetAllowedQueryHashes() != nil {
		target.AllowedQueryHashes = make([]string, len(m.GetAllowedQueryHashes()))
		for idx, v := range m.GetAllowedQueryHashes() {

			target.AllowedQueryHashes[idx] = v

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*GraphQLApiSpec_GraphQLApiOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*GraphQLApiSpec_GraphQLApiOptions)
	}

	switch m.Schema.(type) {

	case *GraphQLApiSpec_ExecutableSchema:

		if h, ok := interface{}(m.GetExecutableSchema()).(clone.Cloner); ok {
			target.Schema = &GraphQLApiSpec_ExecutableSchema{
				ExecutableSchema: h.Clone().(*ExecutableSchema),
			}
		} else {
			target.Schema = &GraphQLApiSpec_ExecutableSchema{
				ExecutableSchema: proto.Clone(m.GetExecutableSchema()).(*ExecutableSchema),
			}
		}

	case *GraphQLApiSpec_StitchedSchema:

		if h, ok := interface{}(m.GetStitchedSchema()).(clone.Cloner); ok {
			target.Schema = &GraphQLApiSpec_StitchedSchema{
				StitchedSchema: h.Clone().(*StitchedSchema),
			}
		} else {
			target.Schema = &GraphQLApiSpec_StitchedSchema{
				StitchedSchema: proto.Clone(m.GetStitchedSchema()).(*StitchedSchema),
			}
		}

	}

	return target
}

// Clone function
func (m *PersistedQueryCacheConfig) Clone() proto.Message {
	var target *PersistedQueryCacheConfig
	if m == nil {
		return target
	}
	target = &PersistedQueryCacheConfig{}

	target.CacheSize = m.GetCacheSize()

	return target
}

// Clone function
func (m *ExecutableSchema) Clone() proto.Message {
	var target *ExecutableSchema
	if m == nil {
		return target
	}
	target = &ExecutableSchema{}

	target.SchemaDefinition = m.GetSchemaDefinition()

	if h, ok := interface{}(m.GetExecutor()).(clone.Cloner); ok {
		target.Executor = h.Clone().(*Executor)
	} else {
		target.Executor = proto.Clone(m.GetExecutor()).(*Executor)
	}

	if h, ok := interface{}(m.GetGrpcDescriptorRegistry()).(clone.Cloner); ok {
		target.GrpcDescriptorRegistry = h.Clone().(*GrpcDescriptorRegistry)
	} else {
		target.GrpcDescriptorRegistry = proto.Clone(m.GetGrpcDescriptorRegistry()).(*GrpcDescriptorRegistry)
	}

	return target
}

// Clone function
func (m *Executor) Clone() proto.Message {
	var target *Executor
	if m == nil {
		return target
	}
	target = &Executor{}

	switch m.Executor.(type) {

	case *Executor_Local_:

		if h, ok := interface{}(m.GetLocal()).(clone.Cloner); ok {
			target.Executor = &Executor_Local_{
				Local: h.Clone().(*Executor_Local),
			}
		} else {
			target.Executor = &Executor_Local_{
				Local: proto.Clone(m.GetLocal()).(*Executor_Local),
			}
		}

	case *Executor_Remote_:

		if h, ok := interface{}(m.GetRemote()).(clone.Cloner); ok {
			target.Executor = &Executor_Remote_{
				Remote: h.Clone().(*Executor_Remote),
			}
		} else {
			target.Executor = &Executor_Remote_{
				Remote: proto.Clone(m.GetRemote()).(*Executor_Remote),
			}
		}

	}

	return target
}

// Clone function
func (m *GraphQLApiStatus) Clone() proto.Message {
	var target *GraphQLApiStatus
	if m == nil {
		return target
	}
	target = &GraphQLApiStatus{}

	target.State = m.GetState()

	target.Reason = m.GetReason()

	target.ReportedBy = m.GetReportedBy()

	if m.GetSubresourceStatuses() != nil {
		target.SubresourceStatuses = make(map[string]*GraphQLApiStatus, len(m.GetSubresourceStatuses()))
		for k, v := range m.GetSubresourceStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SubresourceStatuses[k] = h.Clone().(*GraphQLApiStatus)
			} else {
				target.SubresourceStatuses[k] = proto.Clone(v).(*GraphQLApiStatus)
			}

		}
	}

	if h, ok := interface{}(m.GetDetails()).(clone.Cloner); ok {
		target.Details = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Struct)
	} else {
		target.Details = proto.Clone(m.GetDetails()).(*google_golang_org_protobuf_types_known_structpb.Struct)
	}

	return target
}

// Clone function
func (m *GraphQLApiNamespacedStatuses) Clone() proto.Message {
	var target *GraphQLApiNamespacedStatuses
	if m == nil {
		return target
	}
	target = &GraphQLApiNamespacedStatuses{}

	if m.GetStatuses() != nil {
		target.Statuses = make(map[string]*GraphQLApiStatus, len(m.GetStatuses()))
		for k, v := range m.GetStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Statuses[k] = h.Clone().(*GraphQLApiStatus)
			} else {
				target.Statuses[k] = proto.Clone(v).(*GraphQLApiStatus)
			}

		}
	}

	return target
}

// Clone function
func (m *GrpcDescriptorRegistry_ProtoRefs) Clone() proto.Message {
	var target *GrpcDescriptorRegistry_ProtoRefs
	if m == nil {
		return target
	}
	target = &GrpcDescriptorRegistry_ProtoRefs{}

	if m.GetConfigMapRefs() != nil {
		target.ConfigMapRefs = make([]*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef, len(m.GetConfigMapRefs()))
		for idx, v := range m.GetConfigMapRefs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ConfigMapRefs[idx] = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			} else {
				target.ConfigMapRefs[idx] = proto.Clone(v).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			}

		}
	}

	return target
}

// Clone function
func (m *StitchedSchema_SubschemaConfig) Clone() proto.Message {
	var target *StitchedSchema_SubschemaConfig
	if m == nil {
		return target
	}
	target = &StitchedSchema_SubschemaConfig{}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	if m.GetTypeMerge() != nil {
		target.TypeMerge = make(map[string]*StitchedSchema_SubschemaConfig_TypeMergeConfig, len(m.GetTypeMerge()))
		for k, v := range m.GetTypeMerge() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TypeMerge[k] = h.Clone().(*StitchedSchema_SubschemaConfig_TypeMergeConfig)
			} else {
				target.TypeMerge[k] = proto.Clone(v).(*StitchedSchema_SubschemaConfig_TypeMergeConfig)
			}

		}
	}

	return target
}

// Clone function
func (m *StitchedSchema_SubschemaConfig_TypeMergeConfig) Clone() proto.Message {
	var target *StitchedSchema_SubschemaConfig_TypeMergeConfig
	if m == nil {
		return target
	}
	target = &StitchedSchema_SubschemaConfig_TypeMergeConfig{}

	target.SelectionSet = m.GetSelectionSet()

	target.QueryName = m.GetQueryName()

	if m.GetArgs() != nil {
		target.Args = make(map[string]string, len(m.GetArgs()))
		for k, v := range m.GetArgs() {

			target.Args[k] = v

		}
	}

	return target
}

// Clone function
func (m *MockResolver_AsyncResponse) Clone() proto.Message {
	var target *MockResolver_AsyncResponse
	if m == nil {
		return target
	}
	target = &MockResolver_AsyncResponse{}

	if h, ok := interface{}(m.GetResponse()).(clone.Cloner); ok {
		target.Response = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Value)
	} else {
		target.Response = proto.Clone(m.GetResponse()).(*google_golang_org_protobuf_types_known_structpb.Value)
	}

	if h, ok := interface{}(m.GetDelay()).(clone.Cloner); ok {
		target.Delay = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Delay = proto.Clone(m.GetDelay()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	return target
}

// Clone function
func (m *GraphQLApiSpec_GraphQLApiOptions) Clone() proto.Message {
	var target *GraphQLApiSpec_GraphQLApiOptions
	if m == nil {
		return target
	}
	target = &GraphQLApiSpec_GraphQLApiOptions{}

	target.LogSensitiveInfo = m.GetLogSensitiveInfo()

	return target
}

// Clone function
func (m *Executor_Local) Clone() proto.Message {
	var target *Executor_Local
	if m == nil {
		return target
	}
	target = &Executor_Local{}

	if m.GetResolutions() != nil {
		target.Resolutions = make(map[string]*Resolution, len(m.GetResolutions()))
		for k, v := range m.GetResolutions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Resolutions[k] = h.Clone().(*Resolution)
			} else {
				target.Resolutions[k] = proto.Clone(v).(*Resolution)
			}

		}
	}

	target.EnableIntrospection = m.GetEnableIntrospection()

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*Executor_Local_LocalExecutorOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*Executor_Local_LocalExecutorOptions)
	}

	return target
}

// Clone function
func (m *Executor_Remote) Clone() proto.Message {
	var target *Executor_Remote
	if m == nil {
		return target
	}
	target = &Executor_Remote{}

	if h, ok := interface{}(m.GetUpstreamRef()).(clone.Cloner); ok {
		target.UpstreamRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.UpstreamRef = proto.Clone(m.GetUpstreamRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if m.GetHeaders() != nil {
		target.Headers = make(map[string]string, len(m.GetHeaders()))
		for k, v := range m.GetHeaders() {

			target.Headers[k] = v

		}
	}

	if m.GetQueryParams() != nil {
		target.QueryParams = make(map[string]string, len(m.GetQueryParams()))
		for k, v := range m.GetQueryParams() {

			target.QueryParams[k] = v

		}
	}

	target.SpanName = m.GetSpanName()

	return target
}

// Clone function
func (m *Executor_Local_LocalExecutorOptions) Clone() proto.Message {
	var target *Executor_Local_LocalExecutorOptions
	if m == nil {
		return target
	}
	target = &Executor_Local_LocalExecutorOptions{}

	if h, ok := interface{}(m.GetMaxDepth()).(clone.Cloner); ok {
		target.MaxDepth = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxDepth = proto.Clone(m.GetMaxDepth()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}
