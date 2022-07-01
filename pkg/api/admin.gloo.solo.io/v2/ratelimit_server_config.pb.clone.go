// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/admin/v2/ratelimit_server_config.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2"

	github_com_solo_io_solo_apis_pkg_api_ratelimit_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
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
func (m *RateLimitServerConfigSpec) Clone() proto.Message {
	var target *RateLimitServerConfigSpec
	if m == nil {
		return target
	}
	target = &RateLimitServerConfigSpec{}

	if m.GetDestinationServers() != nil {
		target.DestinationServers = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference, len(m.GetDestinationServers()))
		for idx, v := range m.GetDestinationServers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.DestinationServers[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.DestinationServers[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	switch m.ConfigType.(type) {

	case *RateLimitServerConfigSpec_Raw_:

		if h, ok := interface{}(m.GetRaw()).(clone.Cloner); ok {
			target.ConfigType = &RateLimitServerConfigSpec_Raw_{
				Raw: h.Clone().(*RateLimitServerConfigSpec_Raw),
			}
		} else {
			target.ConfigType = &RateLimitServerConfigSpec_Raw_{
				Raw: proto.Clone(m.GetRaw()).(*RateLimitServerConfigSpec_Raw),
			}
		}

	}

	return target
}

// Clone function
func (m *RateLimitServerConfigStatus) Clone() proto.Message {
	var target *RateLimitServerConfigStatus
	if m == nil {
		return target
	}
	target = &RateLimitServerConfigStatus{}

	target.ObservedGeneration = m.GetObservedGeneration()

	target.State = m.GetState()

	return target
}

// Clone function
func (m *RateLimitServerConfigSpec_Raw) Clone() proto.Message {
	var target *RateLimitServerConfigSpec_Raw
	if m == nil {
		return target
	}
	target = &RateLimitServerConfigSpec_Raw{}

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
