// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/ratelimit/ratelimit.proto

package ratelimit

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"

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
func (m *IngressRateLimit) Clone() proto.Message {
	var target *IngressRateLimit
	if m == nil {
		return target
	}
	target = &IngressRateLimit{}

	if h, ok := interface{}(m.GetAuthorizedLimits()).(clone.Cloner); ok {
		target.AuthorizedLimits = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimit)
	} else {
		target.AuthorizedLimits = proto.Clone(m.GetAuthorizedLimits()).(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimit)
	}

	if h, ok := interface{}(m.GetAnonymousLimits()).(clone.Cloner); ok {
		target.AnonymousLimits = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimit)
	} else {
		target.AnonymousLimits = proto.Clone(m.GetAnonymousLimits()).(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimit)
	}

	return target
}

// Clone function
func (m *Settings) Clone() proto.Message {
	var target *Settings
	if m == nil {
		return target
	}
	target = &Settings{}

	if h, ok := interface{}(m.GetRatelimitServerRef()).(clone.Cloner); ok {
		target.RatelimitServerRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.RatelimitServerRef = proto.Clone(m.GetRatelimitServerRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(clone.Cloner); ok {
		target.RequestTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.RequestTimeout = proto.Clone(m.GetRequestTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	target.DenyOnFail = m.GetDenyOnFail()

	target.EnableXRatelimitHeaders = m.GetEnableXRatelimitHeaders()

	target.RateLimitBeforeAuth = m.GetRateLimitBeforeAuth()

	if h, ok := interface{}(m.GetGrpcService()).(clone.Cloner); ok {
		target.GrpcService = h.Clone().(*GrpcService)
	} else {
		target.GrpcService = proto.Clone(m.GetGrpcService()).(*GrpcService)
	}

	return target
}

// Clone function
func (m *GrpcService) Clone() proto.Message {
	var target *GrpcService
	if m == nil {
		return target
	}
	target = &GrpcService{}

	target.Authority = m.GetAuthority()

	return target
}

// Clone function
func (m *ServiceSettings) Clone() proto.Message {
	var target *ServiceSettings
	if m == nil {
		return target
	}
	target = &ServiceSettings{}

	if m.GetDescriptors() != nil {
		target.Descriptors = make([]*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.Descriptor, len(m.GetDescriptors()))
		for idx, v := range m.GetDescriptors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Descriptors[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.Descriptor)
			} else {
				target.Descriptors[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.Descriptor)
			}

		}
	}

	if m.GetSetDescriptors() != nil {
		target.SetDescriptors = make([]*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.SetDescriptor, len(m.GetSetDescriptors()))
		for idx, v := range m.GetSetDescriptors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SetDescriptors[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.SetDescriptor)
			} else {
				target.SetDescriptors[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.SetDescriptor)
			}

		}
	}

	return target
}

// Clone function
func (m *RateLimitConfigRefs) Clone() proto.Message {
	var target *RateLimitConfigRefs
	if m == nil {
		return target
	}
	target = &RateLimitConfigRefs{}

	if m.GetRefs() != nil {
		target.Refs = make([]*RateLimitConfigRef, len(m.GetRefs()))
		for idx, v := range m.GetRefs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Refs[idx] = h.Clone().(*RateLimitConfigRef)
			} else {
				target.Refs[idx] = proto.Clone(v).(*RateLimitConfigRef)
			}

		}
	}

	return target
}

// Clone function
func (m *RateLimitConfigRef) Clone() proto.Message {
	var target *RateLimitConfigRef
	if m == nil {
		return target
	}
	target = &RateLimitConfigRef{}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	return target
}

// Clone function
func (m *RateLimitVhostExtension) Clone() proto.Message {
	var target *RateLimitVhostExtension
	if m == nil {
		return target
	}
	target = &RateLimitVhostExtension{}

	if m.GetRateLimits() != nil {
		target.RateLimits = make([]*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimitActions, len(m.GetRateLimits()))
		for idx, v := range m.GetRateLimits() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.RateLimits[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimitActions)
			} else {
				target.RateLimits[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimitActions)
			}

		}
	}

	return target
}

// Clone function
func (m *RateLimitRouteExtension) Clone() proto.Message {
	var target *RateLimitRouteExtension
	if m == nil {
		return target
	}
	target = &RateLimitRouteExtension{}

	target.IncludeVhRateLimits = m.GetIncludeVhRateLimits()

	if m.GetRateLimits() != nil {
		target.RateLimits = make([]*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimitActions, len(m.GetRateLimits()))
		for idx, v := range m.GetRateLimits() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.RateLimits[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimitActions)
			} else {
				target.RateLimits[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1.RateLimitActions)
			}

		}
	}

	return target
}
