// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/transformation/transformation.proto

package transformation

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_empty "github.com/golang/protobuf/ptypes/empty"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/route/v3"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/type/matcher/v3"
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
func (m *FilterTransformations) Clone() proto.Message {
	var target *FilterTransformations
	if m == nil {
		return target
	}
	target = &FilterTransformations{}

	if m.GetTransformations() != nil {
		target.Transformations = make([]*TransformationRule, len(m.GetTransformations()))
		for idx, v := range m.GetTransformations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Transformations[idx] = h.Clone().(*TransformationRule)
			} else {
				target.Transformations[idx] = proto.Clone(v).(*TransformationRule)
			}

		}
	}

	target.Stage = m.GetStage()

	return target
}

// Clone function
func (m *TransformationRule) Clone() proto.Message {
	var target *TransformationRule
	if m == nil {
		return target
	}
	target = &TransformationRule{}

	if h, ok := interface{}(m.GetMatch()).(clone.Cloner); ok {
		target.Match = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3.RouteMatch)
	} else {
		target.Match = proto.Clone(m.GetMatch()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3.RouteMatch)
	}

	if h, ok := interface{}(m.GetRouteTransformations()).(clone.Cloner); ok {
		target.RouteTransformations = h.Clone().(*TransformationRule_Transformations)
	} else {
		target.RouteTransformations = proto.Clone(m.GetRouteTransformations()).(*TransformationRule_Transformations)
	}

	return target
}

// Clone function
func (m *RouteTransformations) Clone() proto.Message {
	var target *RouteTransformations
	if m == nil {
		return target
	}
	target = &RouteTransformations{}

	if h, ok := interface{}(m.GetRequestTransformation()).(clone.Cloner); ok {
		target.RequestTransformation = h.Clone().(*Transformation)
	} else {
		target.RequestTransformation = proto.Clone(m.GetRequestTransformation()).(*Transformation)
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(clone.Cloner); ok {
		target.ResponseTransformation = h.Clone().(*Transformation)
	} else {
		target.ResponseTransformation = proto.Clone(m.GetResponseTransformation()).(*Transformation)
	}

	target.ClearRouteCache = m.GetClearRouteCache()

	if m.GetTransformations() != nil {
		target.Transformations = make([]*RouteTransformations_RouteTransformation, len(m.GetTransformations()))
		for idx, v := range m.GetTransformations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Transformations[idx] = h.Clone().(*RouteTransformations_RouteTransformation)
			} else {
				target.Transformations[idx] = proto.Clone(v).(*RouteTransformations_RouteTransformation)
			}

		}
	}

	return target
}

// Clone function
func (m *ResponseMatcher) Clone() proto.Message {
	var target *ResponseMatcher
	if m == nil {
		return target
	}
	target = &ResponseMatcher{}

	if m.GetHeaders() != nil {
		target.Headers = make([]*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3.HeaderMatcher, len(m.GetHeaders()))
		for idx, v := range m.GetHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Headers[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3.HeaderMatcher)
			} else {
				target.Headers[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3.HeaderMatcher)
			}

		}
	}

	if h, ok := interface{}(m.GetResponseCodeDetails()).(clone.Cloner); ok {
		target.ResponseCodeDetails = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3.StringMatcher)
	} else {
		target.ResponseCodeDetails = proto.Clone(m.GetResponseCodeDetails()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3.StringMatcher)
	}

	return target
}

// Clone function
func (m *ResponseTransformationRule) Clone() proto.Message {
	var target *ResponseTransformationRule
	if m == nil {
		return target
	}
	target = &ResponseTransformationRule{}

	if h, ok := interface{}(m.GetMatch()).(clone.Cloner); ok {
		target.Match = h.Clone().(*ResponseMatcher)
	} else {
		target.Match = proto.Clone(m.GetMatch()).(*ResponseMatcher)
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(clone.Cloner); ok {
		target.ResponseTransformation = h.Clone().(*Transformation)
	} else {
		target.ResponseTransformation = proto.Clone(m.GetResponseTransformation()).(*Transformation)
	}

	return target
}

// Clone function
func (m *Transformation) Clone() proto.Message {
	var target *Transformation
	if m == nil {
		return target
	}
	target = &Transformation{}

	switch m.TransformationType.(type) {

	case *Transformation_TransformationTemplate:

		if h, ok := interface{}(m.GetTransformationTemplate()).(clone.Cloner); ok {
			target.TransformationType = &Transformation_TransformationTemplate{
				TransformationTemplate: h.Clone().(*TransformationTemplate),
			}
		} else {
			target.TransformationType = &Transformation_TransformationTemplate{
				TransformationTemplate: proto.Clone(m.GetTransformationTemplate()).(*TransformationTemplate),
			}
		}

	case *Transformation_HeaderBodyTransform:

		if h, ok := interface{}(m.GetHeaderBodyTransform()).(clone.Cloner); ok {
			target.TransformationType = &Transformation_HeaderBodyTransform{
				HeaderBodyTransform: h.Clone().(*HeaderBodyTransform),
			}
		} else {
			target.TransformationType = &Transformation_HeaderBodyTransform{
				HeaderBodyTransform: proto.Clone(m.GetHeaderBodyTransform()).(*HeaderBodyTransform),
			}
		}

	case *Transformation_TransformerConfig:

		if h, ok := interface{}(m.GetTransformerConfig()).(clone.Cloner); ok {
			target.TransformationType = &Transformation_TransformerConfig{
				TransformerConfig: h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.TypedExtensionConfig),
			}
		} else {
			target.TransformationType = &Transformation_TransformerConfig{
				TransformerConfig: proto.Clone(m.GetTransformerConfig()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.TypedExtensionConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *Extraction) Clone() proto.Message {
	var target *Extraction
	if m == nil {
		return target
	}
	target = &Extraction{}

	target.Regex = m.GetRegex()

	target.Subgroup = m.GetSubgroup()

	switch m.Source.(type) {

	case *Extraction_Header:

		target.Source = &Extraction_Header{
			Header: m.GetHeader(),
		}

	case *Extraction_Body:

		if h, ok := interface{}(m.GetBody()).(clone.Cloner); ok {
			target.Source = &Extraction_Body{
				Body: h.Clone().(*github_com_golang_protobuf_ptypes_empty.Empty),
			}
		} else {
			target.Source = &Extraction_Body{
				Body: proto.Clone(m.GetBody()).(*github_com_golang_protobuf_ptypes_empty.Empty),
			}
		}

	}

	return target
}

// Clone function
func (m *TransformationTemplate) Clone() proto.Message {
	var target *TransformationTemplate
	if m == nil {
		return target
	}
	target = &TransformationTemplate{}

	target.AdvancedTemplates = m.GetAdvancedTemplates()

	if m.GetExtractors() != nil {
		target.Extractors = make(map[string]*Extraction, len(m.GetExtractors()))
		for k, v := range m.GetExtractors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Extractors[k] = h.Clone().(*Extraction)
			} else {
				target.Extractors[k] = proto.Clone(v).(*Extraction)
			}

		}
	}

	if m.GetHeaders() != nil {
		target.Headers = make(map[string]*InjaTemplate, len(m.GetHeaders()))
		for k, v := range m.GetHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Headers[k] = h.Clone().(*InjaTemplate)
			} else {
				target.Headers[k] = proto.Clone(v).(*InjaTemplate)
			}

		}
	}

	if m.GetHeadersToAppend() != nil {
		target.HeadersToAppend = make([]*TransformationTemplate_HeaderToAppend, len(m.GetHeadersToAppend()))
		for idx, v := range m.GetHeadersToAppend() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.HeadersToAppend[idx] = h.Clone().(*TransformationTemplate_HeaderToAppend)
			} else {
				target.HeadersToAppend[idx] = proto.Clone(v).(*TransformationTemplate_HeaderToAppend)
			}

		}
	}

	target.ParseBodyBehavior = m.GetParseBodyBehavior()

	target.IgnoreErrorOnParse = m.GetIgnoreErrorOnParse()

	if m.GetDynamicMetadataValues() != nil {
		target.DynamicMetadataValues = make([]*TransformationTemplate_DynamicMetadataValue, len(m.GetDynamicMetadataValues()))
		for idx, v := range m.GetDynamicMetadataValues() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.DynamicMetadataValues[idx] = h.Clone().(*TransformationTemplate_DynamicMetadataValue)
			} else {
				target.DynamicMetadataValues[idx] = proto.Clone(v).(*TransformationTemplate_DynamicMetadataValue)
			}

		}
	}

	switch m.BodyTransformation.(type) {

	case *TransformationTemplate_Body:

		if h, ok := interface{}(m.GetBody()).(clone.Cloner); ok {
			target.BodyTransformation = &TransformationTemplate_Body{
				Body: h.Clone().(*InjaTemplate),
			}
		} else {
			target.BodyTransformation = &TransformationTemplate_Body{
				Body: proto.Clone(m.GetBody()).(*InjaTemplate),
			}
		}

	case *TransformationTemplate_Passthrough:

		if h, ok := interface{}(m.GetPassthrough()).(clone.Cloner); ok {
			target.BodyTransformation = &TransformationTemplate_Passthrough{
				Passthrough: h.Clone().(*Passthrough),
			}
		} else {
			target.BodyTransformation = &TransformationTemplate_Passthrough{
				Passthrough: proto.Clone(m.GetPassthrough()).(*Passthrough),
			}
		}

	case *TransformationTemplate_MergeExtractorsToBody:

		if h, ok := interface{}(m.GetMergeExtractorsToBody()).(clone.Cloner); ok {
			target.BodyTransformation = &TransformationTemplate_MergeExtractorsToBody{
				MergeExtractorsToBody: h.Clone().(*MergeExtractorsToBody),
			}
		} else {
			target.BodyTransformation = &TransformationTemplate_MergeExtractorsToBody{
				MergeExtractorsToBody: proto.Clone(m.GetMergeExtractorsToBody()).(*MergeExtractorsToBody),
			}
		}

	}

	return target
}

// Clone function
func (m *InjaTemplate) Clone() proto.Message {
	var target *InjaTemplate
	if m == nil {
		return target
	}
	target = &InjaTemplate{}

	target.Text = m.GetText()

	return target
}

// Clone function
func (m *Passthrough) Clone() proto.Message {
	var target *Passthrough
	if m == nil {
		return target
	}
	target = &Passthrough{}

	return target
}

// Clone function
func (m *MergeExtractorsToBody) Clone() proto.Message {
	var target *MergeExtractorsToBody
	if m == nil {
		return target
	}
	target = &MergeExtractorsToBody{}

	return target
}

// Clone function
func (m *HeaderBodyTransform) Clone() proto.Message {
	var target *HeaderBodyTransform
	if m == nil {
		return target
	}
	target = &HeaderBodyTransform{}

	target.AddRequestMetadata = m.GetAddRequestMetadata()

	return target
}

// Clone function
func (m *TransformationRule_Transformations) Clone() proto.Message {
	var target *TransformationRule_Transformations
	if m == nil {
		return target
	}
	target = &TransformationRule_Transformations{}

	if h, ok := interface{}(m.GetRequestTransformation()).(clone.Cloner); ok {
		target.RequestTransformation = h.Clone().(*Transformation)
	} else {
		target.RequestTransformation = proto.Clone(m.GetRequestTransformation()).(*Transformation)
	}

	target.ClearRouteCache = m.GetClearRouteCache()

	if h, ok := interface{}(m.GetResponseTransformation()).(clone.Cloner); ok {
		target.ResponseTransformation = h.Clone().(*Transformation)
	} else {
		target.ResponseTransformation = proto.Clone(m.GetResponseTransformation()).(*Transformation)
	}

	if h, ok := interface{}(m.GetOnStreamCompletionTransformation()).(clone.Cloner); ok {
		target.OnStreamCompletionTransformation = h.Clone().(*Transformation)
	} else {
		target.OnStreamCompletionTransformation = proto.Clone(m.GetOnStreamCompletionTransformation()).(*Transformation)
	}

	return target
}

// Clone function
func (m *RouteTransformations_RouteTransformation) Clone() proto.Message {
	var target *RouteTransformations_RouteTransformation
	if m == nil {
		return target
	}
	target = &RouteTransformations_RouteTransformation{}

	target.Stage = m.GetStage()

	switch m.Match.(type) {

	case *RouteTransformations_RouteTransformation_RequestMatch_:

		if h, ok := interface{}(m.GetRequestMatch()).(clone.Cloner); ok {
			target.Match = &RouteTransformations_RouteTransformation_RequestMatch_{
				RequestMatch: h.Clone().(*RouteTransformations_RouteTransformation_RequestMatch),
			}
		} else {
			target.Match = &RouteTransformations_RouteTransformation_RequestMatch_{
				RequestMatch: proto.Clone(m.GetRequestMatch()).(*RouteTransformations_RouteTransformation_RequestMatch),
			}
		}

	case *RouteTransformations_RouteTransformation_ResponseMatch_:

		if h, ok := interface{}(m.GetResponseMatch()).(clone.Cloner); ok {
			target.Match = &RouteTransformations_RouteTransformation_ResponseMatch_{
				ResponseMatch: h.Clone().(*RouteTransformations_RouteTransformation_ResponseMatch),
			}
		} else {
			target.Match = &RouteTransformations_RouteTransformation_ResponseMatch_{
				ResponseMatch: proto.Clone(m.GetResponseMatch()).(*RouteTransformations_RouteTransformation_ResponseMatch),
			}
		}

	}

	return target
}

// Clone function
func (m *RouteTransformations_RouteTransformation_RequestMatch) Clone() proto.Message {
	var target *RouteTransformations_RouteTransformation_RequestMatch
	if m == nil {
		return target
	}
	target = &RouteTransformations_RouteTransformation_RequestMatch{}

	if h, ok := interface{}(m.GetMatch()).(clone.Cloner); ok {
		target.Match = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3.RouteMatch)
	} else {
		target.Match = proto.Clone(m.GetMatch()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_route_v3.RouteMatch)
	}

	if h, ok := interface{}(m.GetRequestTransformation()).(clone.Cloner); ok {
		target.RequestTransformation = h.Clone().(*Transformation)
	} else {
		target.RequestTransformation = proto.Clone(m.GetRequestTransformation()).(*Transformation)
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(clone.Cloner); ok {
		target.ResponseTransformation = h.Clone().(*Transformation)
	} else {
		target.ResponseTransformation = proto.Clone(m.GetResponseTransformation()).(*Transformation)
	}

	target.ClearRouteCache = m.GetClearRouteCache()

	return target
}

// Clone function
func (m *RouteTransformations_RouteTransformation_ResponseMatch) Clone() proto.Message {
	var target *RouteTransformations_RouteTransformation_ResponseMatch
	if m == nil {
		return target
	}
	target = &RouteTransformations_RouteTransformation_ResponseMatch{}

	if h, ok := interface{}(m.GetMatch()).(clone.Cloner); ok {
		target.Match = h.Clone().(*ResponseMatcher)
	} else {
		target.Match = proto.Clone(m.GetMatch()).(*ResponseMatcher)
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(clone.Cloner); ok {
		target.ResponseTransformation = h.Clone().(*Transformation)
	} else {
		target.ResponseTransformation = proto.Clone(m.GetResponseTransformation()).(*Transformation)
	}

	return target
}

// Clone function
func (m *TransformationTemplate_HeaderToAppend) Clone() proto.Message {
	var target *TransformationTemplate_HeaderToAppend
	if m == nil {
		return target
	}
	target = &TransformationTemplate_HeaderToAppend{}

	target.Key = m.GetKey()

	if h, ok := interface{}(m.GetValue()).(clone.Cloner); ok {
		target.Value = h.Clone().(*InjaTemplate)
	} else {
		target.Value = proto.Clone(m.GetValue()).(*InjaTemplate)
	}

	return target
}

// Clone function
func (m *TransformationTemplate_DynamicMetadataValue) Clone() proto.Message {
	var target *TransformationTemplate_DynamicMetadataValue
	if m == nil {
		return target
	}
	target = &TransformationTemplate_DynamicMetadataValue{}

	target.MetadataNamespace = m.GetMetadataNamespace()

	target.Key = m.GetKey()

	if h, ok := interface{}(m.GetValue()).(clone.Cloner); ok {
		target.Value = h.Clone().(*InjaTemplate)
	} else {
		target.Value = proto.Clone(m.GetValue()).(*InjaTemplate)
	}

	return target
}
