// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/transformation/transformation.proto

package transformation

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_extensions_transformers_xslt "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/extensions/transformers/xslt"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/core/matchers"

	google_golang_org_protobuf_types_known_emptypb "google.golang.org/protobuf/types/known/emptypb"

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
func (m *ResponseMatch) Clone() proto.Message {
	var target *ResponseMatch
	if m == nil {
		return target
	}
	target = &ResponseMatch{}

	if m.GetMatchers() != nil {
		target.Matchers = make([]*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers.HeaderMatcher, len(m.GetMatchers()))
		for idx, v := range m.GetMatchers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Matchers[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers.HeaderMatcher)
			} else {
				target.Matchers[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers.HeaderMatcher)
			}

		}
	}

	target.ResponseCodeDetails = m.GetResponseCodeDetails()

	if h, ok := interface{}(m.GetResponseTransformation()).(clone.Cloner); ok {
		target.ResponseTransformation = h.Clone().(*Transformation)
	} else {
		target.ResponseTransformation = proto.Clone(m.GetResponseTransformation()).(*Transformation)
	}

	return target
}

// Clone function
func (m *RequestMatch) Clone() proto.Message {
	var target *RequestMatch
	if m == nil {
		return target
	}
	target = &RequestMatch{}

	if h, ok := interface{}(m.GetMatcher()).(clone.Cloner); ok {
		target.Matcher = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers.Matcher)
	} else {
		target.Matcher = proto.Clone(m.GetMatcher()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_v1_core_matchers.Matcher)
	}

	target.ClearRouteCache = m.GetClearRouteCache()

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

	return target
}

// Clone function
func (m *Transformations) Clone() proto.Message {
	var target *Transformations
	if m == nil {
		return target
	}
	target = &Transformations{}

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

	return target
}

// Clone function
func (m *RequestResponseTransformations) Clone() proto.Message {
	var target *RequestResponseTransformations
	if m == nil {
		return target
	}
	target = &RequestResponseTransformations{}

	if m.GetRequestTransforms() != nil {
		target.RequestTransforms = make([]*RequestMatch, len(m.GetRequestTransforms()))
		for idx, v := range m.GetRequestTransforms() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.RequestTransforms[idx] = h.Clone().(*RequestMatch)
			} else {
				target.RequestTransforms[idx] = proto.Clone(v).(*RequestMatch)
			}

		}
	}

	if m.GetResponseTransforms() != nil {
		target.ResponseTransforms = make([]*ResponseMatch, len(m.GetResponseTransforms()))
		for idx, v := range m.GetResponseTransforms() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ResponseTransforms[idx] = h.Clone().(*ResponseMatch)
			} else {
				target.ResponseTransforms[idx] = proto.Clone(v).(*ResponseMatch)
			}

		}
	}

	return target
}

// Clone function
func (m *TransformationStages) Clone() proto.Message {
	var target *TransformationStages
	if m == nil {
		return target
	}
	target = &TransformationStages{}

	if h, ok := interface{}(m.GetEarly()).(clone.Cloner); ok {
		target.Early = h.Clone().(*RequestResponseTransformations)
	} else {
		target.Early = proto.Clone(m.GetEarly()).(*RequestResponseTransformations)
	}

	if h, ok := interface{}(m.GetRegular()).(clone.Cloner); ok {
		target.Regular = h.Clone().(*RequestResponseTransformations)
	} else {
		target.Regular = proto.Clone(m.GetRegular()).(*RequestResponseTransformations)
	}

	if h, ok := interface{}(m.GetPostRouting()).(clone.Cloner); ok {
		target.PostRouting = h.Clone().(*RequestResponseTransformations)
	} else {
		target.PostRouting = proto.Clone(m.GetPostRouting()).(*RequestResponseTransformations)
	}

	target.InheritTransformation = m.GetInheritTransformation()

	if h, ok := interface{}(m.GetLogRequestResponseInfo()).(clone.Cloner); ok {
		target.LogRequestResponseInfo = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.LogRequestResponseInfo = proto.Clone(m.GetLogRequestResponseInfo()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetEscapeCharacters()).(clone.Cloner); ok {
		target.EscapeCharacters = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.EscapeCharacters = proto.Clone(m.GetEscapeCharacters()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
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

	target.LogRequestResponseInfo = m.GetLogRequestResponseInfo()

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

	case *Transformation_XsltTransformation:

		if h, ok := interface{}(m.GetXsltTransformation()).(clone.Cloner); ok {
			target.TransformationType = &Transformation_XsltTransformation{
				XsltTransformation: h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_extensions_transformers_xslt.XsltTransformation),
			}
		} else {
			target.TransformationType = &Transformation_XsltTransformation{
				XsltTransformation: proto.Clone(m.GetXsltTransformation()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_extensions_transformers_xslt.XsltTransformation),
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

	if h, ok := interface{}(m.GetReplacementText()).(clone.Cloner); ok {
		target.ReplacementText = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.ReplacementText = proto.Clone(m.GetReplacementText()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	target.Mode = m.GetMode()

	switch m.Source.(type) {

	case *Extraction_Header:

		target.Source = &Extraction_Header{
			Header: m.GetHeader(),
		}

	case *Extraction_Body:

		if h, ok := interface{}(m.GetBody()).(clone.Cloner); ok {
			target.Source = &Extraction_Body{
				Body: h.Clone().(*google_golang_org_protobuf_types_known_emptypb.Empty),
			}
		} else {
			target.Source = &Extraction_Body{
				Body: proto.Clone(m.GetBody()).(*google_golang_org_protobuf_types_known_emptypb.Empty),
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

	if m.GetHeadersToRemove() != nil {
		target.HeadersToRemove = make([]string, len(m.GetHeadersToRemove()))
		for idx, v := range m.GetHeadersToRemove() {

			target.HeadersToRemove[idx] = v

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

	if h, ok := interface{}(m.GetEscapeCharacters()).(clone.Cloner); ok {
		target.EscapeCharacters = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.EscapeCharacters = proto.Clone(m.GetEscapeCharacters()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
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

	case *TransformationTemplate_MergeJsonKeys:

		if h, ok := interface{}(m.GetMergeJsonKeys()).(clone.Cloner); ok {
			target.BodyTransformation = &TransformationTemplate_MergeJsonKeys{
				MergeJsonKeys: h.Clone().(*MergeJsonKeys),
			}
		} else {
			target.BodyTransformation = &TransformationTemplate_MergeJsonKeys{
				MergeJsonKeys: proto.Clone(m.GetMergeJsonKeys()).(*MergeJsonKeys),
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
func (m *MergeJsonKeys) Clone() proto.Message {
	var target *MergeJsonKeys
	if m == nil {
		return target
	}
	target = &MergeJsonKeys{}

	if m.GetJsonKeys() != nil {
		target.JsonKeys = make(map[string]*MergeJsonKeys_OverridableTemplate, len(m.GetJsonKeys()))
		for k, v := range m.GetJsonKeys() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.JsonKeys[k] = h.Clone().(*MergeJsonKeys_OverridableTemplate)
			} else {
				target.JsonKeys[k] = proto.Clone(v).(*MergeJsonKeys_OverridableTemplate)
			}

		}
	}

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

	target.JsonToProto = m.GetJsonToProto()

	return target
}

// Clone function
func (m *MergeJsonKeys_OverridableTemplate) Clone() proto.Message {
	var target *MergeJsonKeys_OverridableTemplate
	if m == nil {
		return target
	}
	target = &MergeJsonKeys_OverridableTemplate{}

	if h, ok := interface{}(m.GetTmpl()).(clone.Cloner); ok {
		target.Tmpl = h.Clone().(*InjaTemplate)
	} else {
		target.Tmpl = proto.Clone(m.GetTmpl()).(*InjaTemplate)
	}

	target.OverrideEmpty = m.GetOverrideEmpty()

	return target
}
