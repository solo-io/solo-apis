// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/caching/caching.proto

package caching

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/type/matcher/v3"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

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
func (m *Settings) Clone() proto.Message {
	var target *Settings
	if m == nil {
		return target
	}
	target = &Settings{}

	if h, ok := interface{}(m.GetCachingServiceRef()).(clone.Cloner); ok {
		target.CachingServiceRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.CachingServiceRef = proto.Clone(m.GetCachingServiceRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if m.GetAllowedVaryHeaders() != nil {
		target.AllowedVaryHeaders = make([]*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3.StringMatcher, len(m.GetAllowedVaryHeaders()))
		for idx, v := range m.GetAllowedVaryHeaders() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AllowedVaryHeaders[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3.StringMatcher)
			} else {
				target.AllowedVaryHeaders[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3.StringMatcher)
			}

		}
	}

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetMaxPayloadSize()).(clone.Cloner); ok {
		target.MaxPayloadSize = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxPayloadSize = proto.Clone(m.GetMaxPayloadSize()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}
