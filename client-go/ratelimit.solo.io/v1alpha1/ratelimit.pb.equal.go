// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/vendor_any/github.com/solo-io/solo-apis/api/rate-limiter/v1alpha1/ratelimit.proto

package v1alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *RateLimitConfigSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RateLimitConfigSpec)
	if !ok {
		that2, ok := that.(RateLimitConfigSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.ConfigType.(type) {

	case *RateLimitConfigSpec_Raw_:
		if _, ok := target.ConfigType.(*RateLimitConfigSpec_Raw_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRaw()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRaw()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRaw(), target.GetRaw()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ConfigType != target.ConfigType {
			return false
		}
	}

	return true
}

// Equal function
func (m *RateLimitConfigStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RateLimitConfigStatus)
	if !ok {
		that2, ok := that.(RateLimitConfigStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if strings.Compare(m.GetMessage(), target.GetMessage()) != 0 {
		return false
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	return true
}

// Equal function
func (m *RateLimitConfigNamespacedStatuses) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RateLimitConfigNamespacedStatuses)
	if !ok {
		that2, ok := that.(RateLimitConfigNamespacedStatuses)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetStatuses()) != len(target.GetStatuses()) {
		return false
	}
	for k, v := range m.GetStatuses() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetStatuses()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetStatuses()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Descriptor) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Descriptor)
	if !ok {
		that2, ok := that.(Descriptor)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetValue(), target.GetValue()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetRateLimit()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRateLimit()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRateLimit(), target.GetRateLimit()) {
			return false
		}
	}

	if len(m.GetDescriptors()) != len(target.GetDescriptors()) {
		return false
	}
	for idx, v := range m.GetDescriptors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDescriptors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDescriptors()[idx]) {
				return false
			}
		}

	}

	if m.GetWeight() != target.GetWeight() {
		return false
	}

	if m.GetAlwaysApply() != target.GetAlwaysApply() {
		return false
	}

	return true
}

// Equal function
func (m *SetDescriptor) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SetDescriptor)
	if !ok {
		that2, ok := that.(SetDescriptor)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetSimpleDescriptors()) != len(target.GetSimpleDescriptors()) {
		return false
	}
	for idx, v := range m.GetSimpleDescriptors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSimpleDescriptors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSimpleDescriptors()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetRateLimit()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRateLimit()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRateLimit(), target.GetRateLimit()) {
			return false
		}
	}

	if m.GetAlwaysApply() != target.GetAlwaysApply() {
		return false
	}

	return true
}

// Equal function
func (m *SimpleDescriptor) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SimpleDescriptor)
	if !ok {
		that2, ok := that.(SimpleDescriptor)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetValue(), target.GetValue()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *RateLimitActions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RateLimitActions)
	if !ok {
		that2, ok := that.(RateLimitActions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetActions()) != len(target.GetActions()) {
		return false
	}
	for idx, v := range m.GetActions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetActions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetActions()[idx]) {
				return false
			}
		}

	}

	if len(m.GetSetActions()) != len(target.GetSetActions()) {
		return false
	}
	for idx, v := range m.GetSetActions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSetActions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSetActions()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *RateLimit) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RateLimit)
	if !ok {
		that2, ok := that.(RateLimit)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetUnit() != target.GetUnit() {
		return false
	}

	if m.GetRequestsPerUnit() != target.GetRequestsPerUnit() {
		return false
	}

	return true
}

// Equal function
func (m *Action) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action)
	if !ok {
		that2, ok := that.(Action)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.ActionSpecifier.(type) {

	case *Action_SourceCluster_:
		if _, ok := target.ActionSpecifier.(*Action_SourceCluster_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSourceCluster()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSourceCluster()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSourceCluster(), target.GetSourceCluster()) {
				return false
			}
		}

	case *Action_DestinationCluster_:
		if _, ok := target.ActionSpecifier.(*Action_DestinationCluster_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDestinationCluster()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDestinationCluster()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDestinationCluster(), target.GetDestinationCluster()) {
				return false
			}
		}

	case *Action_RequestHeaders_:
		if _, ok := target.ActionSpecifier.(*Action_RequestHeaders_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRequestHeaders()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestHeaders()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRequestHeaders(), target.GetRequestHeaders()) {
				return false
			}
		}

	case *Action_RemoteAddress_:
		if _, ok := target.ActionSpecifier.(*Action_RemoteAddress_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRemoteAddress()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRemoteAddress()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRemoteAddress(), target.GetRemoteAddress()) {
				return false
			}
		}

	case *Action_GenericKey_:
		if _, ok := target.ActionSpecifier.(*Action_GenericKey_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGenericKey()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGenericKey()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGenericKey(), target.GetGenericKey()) {
				return false
			}
		}

	case *Action_HeaderValueMatch_:
		if _, ok := target.ActionSpecifier.(*Action_HeaderValueMatch_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHeaderValueMatch()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeaderValueMatch()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHeaderValueMatch(), target.GetHeaderValueMatch()) {
				return false
			}
		}

	case *Action_Metadata:
		if _, ok := target.ActionSpecifier.(*Action_Metadata); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMetadata()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ActionSpecifier != target.ActionSpecifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *RateLimitConfigSpec_Raw) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RateLimitConfigSpec_Raw)
	if !ok {
		that2, ok := that.(RateLimitConfigSpec_Raw)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetDescriptors()) != len(target.GetDescriptors()) {
		return false
	}
	for idx, v := range m.GetDescriptors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDescriptors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDescriptors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetRateLimits()) != len(target.GetRateLimits()) {
		return false
	}
	for idx, v := range m.GetRateLimits() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRateLimits()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRateLimits()[idx]) {
				return false
			}
		}

	}

	if len(m.GetSetDescriptors()) != len(target.GetSetDescriptors()) {
		return false
	}
	for idx, v := range m.GetSetDescriptors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSetDescriptors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSetDescriptors()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Action_SourceCluster) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_SourceCluster)
	if !ok {
		that2, ok := that.(Action_SourceCluster)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *Action_DestinationCluster) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_DestinationCluster)
	if !ok {
		that2, ok := that.(Action_DestinationCluster)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *Action_RequestHeaders) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_RequestHeaders)
	if !ok {
		that2, ok := that.(Action_RequestHeaders)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetHeaderName(), target.GetHeaderName()) != 0 {
		return false
	}

	if strings.Compare(m.GetDescriptorKey(), target.GetDescriptorKey()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Action_RemoteAddress) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_RemoteAddress)
	if !ok {
		that2, ok := that.(Action_RemoteAddress)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *Action_GenericKey) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_GenericKey)
	if !ok {
		that2, ok := that.(Action_GenericKey)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetDescriptorValue(), target.GetDescriptorValue()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Action_HeaderValueMatch) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_HeaderValueMatch)
	if !ok {
		that2, ok := that.(Action_HeaderValueMatch)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetDescriptorValue(), target.GetDescriptorValue()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetExpectMatch()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExpectMatch()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExpectMatch(), target.GetExpectMatch()) {
			return false
		}
	}

	if len(m.GetHeaders()) != len(target.GetHeaders()) {
		return false
	}
	for idx, v := range m.GetHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHeaders()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Action_MetaData) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_MetaData)
	if !ok {
		that2, ok := that.(Action_MetaData)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetDescriptorKey(), target.GetDescriptorKey()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetMetadataKey()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadataKey()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadataKey(), target.GetMetadataKey()) {
			return false
		}
	}

	if strings.Compare(m.GetDefaultValue(), target.GetDefaultValue()) != 0 {
		return false
	}

	if m.GetSource() != target.GetSource() {
		return false
	}

	return true
}

// Equal function
func (m *Action_HeaderValueMatch_HeaderMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_HeaderValueMatch_HeaderMatcher)
	if !ok {
		that2, ok := that.(Action_HeaderValueMatch_HeaderMatcher)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if m.GetInvertMatch() != target.GetInvertMatch() {
		return false
	}

	switch m.HeaderMatchSpecifier.(type) {

	case *Action_HeaderValueMatch_HeaderMatcher_ExactMatch:
		if _, ok := target.HeaderMatchSpecifier.(*Action_HeaderValueMatch_HeaderMatcher_ExactMatch); !ok {
			return false
		}

		if strings.Compare(m.GetExactMatch(), target.GetExactMatch()) != 0 {
			return false
		}

	case *Action_HeaderValueMatch_HeaderMatcher_RegexMatch:
		if _, ok := target.HeaderMatchSpecifier.(*Action_HeaderValueMatch_HeaderMatcher_RegexMatch); !ok {
			return false
		}

		if strings.Compare(m.GetRegexMatch(), target.GetRegexMatch()) != 0 {
			return false
		}

	case *Action_HeaderValueMatch_HeaderMatcher_RangeMatch:
		if _, ok := target.HeaderMatchSpecifier.(*Action_HeaderValueMatch_HeaderMatcher_RangeMatch); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRangeMatch()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRangeMatch()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRangeMatch(), target.GetRangeMatch()) {
				return false
			}
		}

	case *Action_HeaderValueMatch_HeaderMatcher_PresentMatch:
		if _, ok := target.HeaderMatchSpecifier.(*Action_HeaderValueMatch_HeaderMatcher_PresentMatch); !ok {
			return false
		}

		if m.GetPresentMatch() != target.GetPresentMatch() {
			return false
		}

	case *Action_HeaderValueMatch_HeaderMatcher_PrefixMatch:
		if _, ok := target.HeaderMatchSpecifier.(*Action_HeaderValueMatch_HeaderMatcher_PrefixMatch); !ok {
			return false
		}

		if strings.Compare(m.GetPrefixMatch(), target.GetPrefixMatch()) != 0 {
			return false
		}

	case *Action_HeaderValueMatch_HeaderMatcher_SuffixMatch:
		if _, ok := target.HeaderMatchSpecifier.(*Action_HeaderValueMatch_HeaderMatcher_SuffixMatch); !ok {
			return false
		}

		if strings.Compare(m.GetSuffixMatch(), target.GetSuffixMatch()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.HeaderMatchSpecifier != target.HeaderMatchSpecifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *Action_HeaderValueMatch_HeaderMatcher_Int64Range) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_HeaderValueMatch_HeaderMatcher_Int64Range)
	if !ok {
		that2, ok := that.(Action_HeaderValueMatch_HeaderMatcher_Int64Range)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetStart() != target.GetStart() {
		return false
	}

	if m.GetEnd() != target.GetEnd() {
		return false
	}

	return true
}

// Equal function
func (m *Action_MetaData_MetadataKey) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_MetaData_MetadataKey)
	if !ok {
		that2, ok := that.(Action_MetaData_MetadataKey)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if len(m.GetPath()) != len(target.GetPath()) {
		return false
	}
	for idx, v := range m.GetPath() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPath()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPath()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Action_MetaData_MetadataKey_PathSegment) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action_MetaData_MetadataKey_PathSegment)
	if !ok {
		that2, ok := that.(Action_MetaData_MetadataKey_PathSegment)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Segment.(type) {

	case *Action_MetaData_MetadataKey_PathSegment_Key:
		if _, ok := target.Segment.(*Action_MetaData_MetadataKey_PathSegment_Key); !ok {
			return false
		}

		if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.Segment != target.Segment {
			return false
		}
	}

	return true
}