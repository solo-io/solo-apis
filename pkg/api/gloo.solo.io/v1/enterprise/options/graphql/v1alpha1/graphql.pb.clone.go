// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/graphql/v1alpha1/graphql.proto

package v1alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
func (m *PathSegment) Clone() proto.Message {
	var target *PathSegment
	if m == nil {
		return target
	}
	target = &PathSegment{}

	switch m.Segment.(type) {

	case *PathSegment_Key:

		target.Segment = &PathSegment_Key{
			Key: m.GetKey(),
		}

	case *PathSegment_Index:

		target.Segment = &PathSegment_Index{
			Index: m.GetIndex(),
		}

	case *PathSegment_All:

		target.Segment = &PathSegment_All{
			All: m.GetAll(),
		}

	}

	return target
}

// Clone function
func (m *Path) Clone() proto.Message {
	var target *Path
	if m == nil {
		return target
	}
	target = &Path{}

	if m.GetSegments() != nil {
		target.Segments = make([]*PathSegment, len(m.GetSegments()))
		for idx, v := range m.GetSegments() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Segments[idx] = h.Clone().(*PathSegment)
			} else {
				target.Segments[idx] = proto.Clone(v).(*PathSegment)
			}

		}
	}

	return target
}

// Clone function
func (m *ValueProvider) Clone() proto.Message {
	var target *ValueProvider
	if m == nil {
		return target
	}
	target = &ValueProvider{}

	if m.GetProviders() != nil {
		target.Providers = make(map[string]*ValueProvider_Provider, len(m.GetProviders()))
		for k, v := range m.GetProviders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Providers[k] = h.Clone().(*ValueProvider_Provider)
			} else {
				target.Providers[k] = proto.Clone(v).(*ValueProvider_Provider)
			}

		}
	}

	target.ProviderTemplate = m.GetProviderTemplate()

	return target
}

// Clone function
func (m *JsonValueList) Clone() proto.Message {
	var target *JsonValueList
	if m == nil {
		return target
	}
	target = &JsonValueList{}

	if m.GetValues() != nil {
		target.Values = make([]*JsonValue, len(m.GetValues()))
		for idx, v := range m.GetValues() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Values[idx] = h.Clone().(*JsonValue)
			} else {
				target.Values[idx] = proto.Clone(v).(*JsonValue)
			}

		}
	}

	return target
}

// Clone function
func (m *JsonValue) Clone() proto.Message {
	var target *JsonValue
	if m == nil {
		return target
	}
	target = &JsonValue{}

	switch m.JsonVal.(type) {

	case *JsonValue_Node:

		if h, ok := interface{}(m.GetNode()).(clone.Cloner); ok {
			target.JsonVal = &JsonValue_Node{
				Node: h.Clone().(*JsonNode),
			}
		} else {
			target.JsonVal = &JsonValue_Node{
				Node: proto.Clone(m.GetNode()).(*JsonNode),
			}
		}

	case *JsonValue_ValueProvider:

		if h, ok := interface{}(m.GetValueProvider()).(clone.Cloner); ok {
			target.JsonVal = &JsonValue_ValueProvider{
				ValueProvider: h.Clone().(*ValueProvider),
			}
		} else {
			target.JsonVal = &JsonValue_ValueProvider{
				ValueProvider: proto.Clone(m.GetValueProvider()).(*ValueProvider),
			}
		}

	case *JsonValue_List:

		if h, ok := interface{}(m.GetList()).(clone.Cloner); ok {
			target.JsonVal = &JsonValue_List{
				List: h.Clone().(*JsonValueList),
			}
		} else {
			target.JsonVal = &JsonValue_List{
				List: proto.Clone(m.GetList()).(*JsonValueList),
			}
		}

	}

	return target
}

// Clone function
func (m *JsonKeyValue) Clone() proto.Message {
	var target *JsonKeyValue
	if m == nil {
		return target
	}
	target = &JsonKeyValue{}

	target.Key = m.GetKey()

	if h, ok := interface{}(m.GetValue()).(clone.Cloner); ok {
		target.Value = h.Clone().(*JsonValue)
	} else {
		target.Value = proto.Clone(m.GetValue()).(*JsonValue)
	}

	return target
}

// Clone function
func (m *JsonNode) Clone() proto.Message {
	var target *JsonNode
	if m == nil {
		return target
	}
	target = &JsonNode{}

	if m.GetKeyValues() != nil {
		target.KeyValues = make([]*JsonKeyValue, len(m.GetKeyValues()))
		for idx, v := range m.GetKeyValues() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.KeyValues[idx] = h.Clone().(*JsonKeyValue)
			} else {
				target.KeyValues[idx] = proto.Clone(v).(*JsonKeyValue)
			}

		}
	}

	return target
}

// Clone function
func (m *RequestTemplate) Clone() proto.Message {
	var target *RequestTemplate
	if m == nil {
		return target
	}
	target = &RequestTemplate{}

	if m.GetHeaders() != nil {
		target.Headers = make(map[string]*ValueProvider, len(m.GetHeaders()))
		for k, v := range m.GetHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Headers[k] = h.Clone().(*ValueProvider)
			} else {
				target.Headers[k] = proto.Clone(v).(*ValueProvider)
			}

		}
	}

	if m.GetQueryParams() != nil {
		target.QueryParams = make(map[string]*ValueProvider, len(m.GetQueryParams()))
		for k, v := range m.GetQueryParams() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.QueryParams[k] = h.Clone().(*ValueProvider)
			} else {
				target.QueryParams[k] = proto.Clone(v).(*ValueProvider)
			}

		}
	}

	if h, ok := interface{}(m.GetOutgoingBody()).(clone.Cloner); ok {
		target.OutgoingBody = h.Clone().(*JsonValue)
	} else {
		target.OutgoingBody = proto.Clone(m.GetOutgoingBody()).(*JsonValue)
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

	if m.GetResultRoot() != nil {
		target.ResultRoot = make([]*PathSegment, len(m.GetResultRoot()))
		for idx, v := range m.GetResultRoot() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ResultRoot[idx] = h.Clone().(*PathSegment)
			} else {
				target.ResultRoot[idx] = proto.Clone(v).(*PathSegment)
			}

		}
	}

	if m.GetSetters() != nil {
		target.Setters = make(map[string]*Path, len(m.GetSetters()))
		for k, v := range m.GetSetters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Setters[k] = h.Clone().(*Path)
			} else {
				target.Setters[k] = proto.Clone(v).(*Path)
			}

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
		target.OutgoingMessageJson = h.Clone().(*JsonValue)
	} else {
		target.OutgoingMessageJson = proto.Clone(m.GetOutgoingMessageJson()).(*JsonValue)
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

	if h, ok := interface{}(m.GetRequestTransform()).(clone.Cloner); ok {
		target.RequestTransform = h.Clone().(*RequestTemplate)
	} else {
		target.RequestTransform = proto.Clone(m.GetRequestTransform()).(*RequestTemplate)
	}

	if h, ok := interface{}(m.GetPreExecutionTransform()).(clone.Cloner); ok {
		target.PreExecutionTransform = h.Clone().(*ResponseTemplate)
	} else {
		target.PreExecutionTransform = proto.Clone(m.GetPreExecutionTransform()).(*ResponseTemplate)
	}

	target.SpanName = m.GetSpanName()

	return target
}

// Clone function
func (m *GrpcDescriptorRegistry) Clone() proto.Message {
	var target *GrpcDescriptorRegistry
	if m == nil {
		return target
	}
	target = &GrpcDescriptorRegistry{}

	target.ProtoDescriptorsBin = m.GetProtoDescriptorsBin()

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

	return target
}

// Clone function
func (m *QueryMatcher) Clone() proto.Message {
	var target *QueryMatcher
	if m == nil {
		return target
	}
	target = &QueryMatcher{}

	switch m.Match.(type) {

	case *QueryMatcher_FieldMatcher_:

		if h, ok := interface{}(m.GetFieldMatcher()).(clone.Cloner); ok {
			target.Match = &QueryMatcher_FieldMatcher_{
				FieldMatcher: h.Clone().(*QueryMatcher_FieldMatcher),
			}
		} else {
			target.Match = &QueryMatcher_FieldMatcher_{
				FieldMatcher: proto.Clone(m.GetFieldMatcher()).(*QueryMatcher_FieldMatcher),
			}
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

	if h, ok := interface{}(m.GetMatcher()).(clone.Cloner); ok {
		target.Matcher = h.Clone().(*QueryMatcher)
	} else {
		target.Matcher = proto.Clone(m.GetMatcher()).(*QueryMatcher)
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

	}

	return target
}

// Clone function
func (m *GraphQLSchema) Clone() proto.Message {
	var target *GraphQLSchema
	if m == nil {
		return target
	}
	target = &GraphQLSchema{}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	target.Schema = m.GetSchema()

	target.EnableIntrospection = m.GetEnableIntrospection()

	if m.GetResolutions() != nil {
		target.Resolutions = make([]*Resolution, len(m.GetResolutions()))
		for idx, v := range m.GetResolutions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Resolutions[idx] = h.Clone().(*Resolution)
			} else {
				target.Resolutions[idx] = proto.Clone(v).(*Resolution)
			}

		}
	}

	if h, ok := interface{}(m.GetExecutableSchema()).(clone.Cloner); ok {
		target.ExecutableSchema = h.Clone().(*ExecutableSchema)
	} else {
		target.ExecutableSchema = proto.Clone(m.GetExecutableSchema()).(*ExecutableSchema)
	}

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

	}

	return target
}

// Clone function
func (m *ValueProvider_GraphQLArgExtraction) Clone() proto.Message {
	var target *ValueProvider_GraphQLArgExtraction
	if m == nil {
		return target
	}
	target = &ValueProvider_GraphQLArgExtraction{}

	target.ArgName = m.GetArgName()

	if m.GetPath() != nil {
		target.Path = make([]*PathSegment, len(m.GetPath()))
		for idx, v := range m.GetPath() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Path[idx] = h.Clone().(*PathSegment)
			} else {
				target.Path[idx] = proto.Clone(v).(*PathSegment)
			}

		}
	}

	target.Required = m.GetRequired()

	return target
}

// Clone function
func (m *ValueProvider_GraphQLParentExtraction) Clone() proto.Message {
	var target *ValueProvider_GraphQLParentExtraction
	if m == nil {
		return target
	}
	target = &ValueProvider_GraphQLParentExtraction{}

	if m.GetPath() != nil {
		target.Path = make([]*PathSegment, len(m.GetPath()))
		for idx, v := range m.GetPath() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Path[idx] = h.Clone().(*PathSegment)
			} else {
				target.Path[idx] = proto.Clone(v).(*PathSegment)
			}

		}
	}

	return target
}

// Clone function
func (m *ValueProvider_TypedValueProvider) Clone() proto.Message {
	var target *ValueProvider_TypedValueProvider
	if m == nil {
		return target
	}
	target = &ValueProvider_TypedValueProvider{}

	target.Type = m.GetType()

	switch m.ValProvider.(type) {

	case *ValueProvider_TypedValueProvider_Header:

		target.ValProvider = &ValueProvider_TypedValueProvider_Header{
			Header: m.GetHeader(),
		}

	case *ValueProvider_TypedValueProvider_Value:

		target.ValProvider = &ValueProvider_TypedValueProvider_Value{
			Value: m.GetValue(),
		}

	}

	return target
}

// Clone function
func (m *ValueProvider_Provider) Clone() proto.Message {
	var target *ValueProvider_Provider
	if m == nil {
		return target
	}
	target = &ValueProvider_Provider{}

	switch m.Provider.(type) {

	case *ValueProvider_Provider_GraphqlArg:

		if h, ok := interface{}(m.GetGraphqlArg()).(clone.Cloner); ok {
			target.Provider = &ValueProvider_Provider_GraphqlArg{
				GraphqlArg: h.Clone().(*ValueProvider_GraphQLArgExtraction),
			}
		} else {
			target.Provider = &ValueProvider_Provider_GraphqlArg{
				GraphqlArg: proto.Clone(m.GetGraphqlArg()).(*ValueProvider_GraphQLArgExtraction),
			}
		}

	case *ValueProvider_Provider_TypedProvider:

		if h, ok := interface{}(m.GetTypedProvider()).(clone.Cloner); ok {
			target.Provider = &ValueProvider_Provider_TypedProvider{
				TypedProvider: h.Clone().(*ValueProvider_TypedValueProvider),
			}
		} else {
			target.Provider = &ValueProvider_Provider_TypedProvider{
				TypedProvider: proto.Clone(m.GetTypedProvider()).(*ValueProvider_TypedValueProvider),
			}
		}

	case *ValueProvider_Provider_GraphqlParent:

		if h, ok := interface{}(m.GetGraphqlParent()).(clone.Cloner); ok {
			target.Provider = &ValueProvider_Provider_GraphqlParent{
				GraphqlParent: h.Clone().(*ValueProvider_GraphQLParentExtraction),
			}
		} else {
			target.Provider = &ValueProvider_Provider_GraphqlParent{
				GraphqlParent: proto.Clone(m.GetGraphqlParent()).(*ValueProvider_GraphQLParentExtraction),
			}
		}

	}

	return target
}

// Clone function
func (m *QueryMatcher_FieldMatcher) Clone() proto.Message {
	var target *QueryMatcher_FieldMatcher
	if m == nil {
		return target
	}
	target = &QueryMatcher_FieldMatcher{}

	target.Type = m.GetType()

	target.Field = m.GetField()

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
		target.Resolutions = make([]*Resolution, len(m.GetResolutions()))
		for idx, v := range m.GetResolutions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Resolutions[idx] = h.Clone().(*Resolution)
			} else {
				target.Resolutions[idx] = proto.Clone(v).(*Resolution)
			}

		}
	}

	target.EnableIntrospection = m.GetEnableIntrospection()

	return target
}