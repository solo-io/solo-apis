// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/common/mutation_rules/v3/mutation_rules.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"

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
func (m *HeaderMutationRules) Clone() proto.Message {
	var target *HeaderMutationRules
	if m == nil {
		return target
	}
	target = &HeaderMutationRules{}

	if h, ok := interface{}(m.GetAllowAllRouting()).(clone.Cloner); ok {
		target.AllowAllRouting = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.AllowAllRouting = proto.Clone(m.GetAllowAllRouting()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetAllowEnvoy()).(clone.Cloner); ok {
		target.AllowEnvoy = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.AllowEnvoy = proto.Clone(m.GetAllowEnvoy()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetDisallowSystem()).(clone.Cloner); ok {
		target.DisallowSystem = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.DisallowSystem = proto.Clone(m.GetDisallowSystem()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetDisallowAll()).(clone.Cloner); ok {
		target.DisallowAll = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.DisallowAll = proto.Clone(m.GetDisallowAll()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetAllowExpression()).(clone.Cloner); ok {
		target.AllowExpression = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3.RegexMatcher)
	} else {
		target.AllowExpression = proto.Clone(m.GetAllowExpression()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3.RegexMatcher)
	}

	if h, ok := interface{}(m.GetDisallowExpression()).(clone.Cloner); ok {
		target.DisallowExpression = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3.RegexMatcher)
	} else {
		target.DisallowExpression = proto.Clone(m.GetDisallowExpression()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_type_matcher_v3.RegexMatcher)
	}

	if h, ok := interface{}(m.GetDisallowIsError()).(clone.Cloner); ok {
		target.DisallowIsError = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.DisallowIsError = proto.Clone(m.GetDisallowIsError()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *HeaderMutation) Clone() proto.Message {
	var target *HeaderMutation
	if m == nil {
		return target
	}
	target = &HeaderMutation{}

	switch m.Action.(type) {

	case *HeaderMutation_Remove:

		target.Action = &HeaderMutation_Remove{
			Remove: m.GetRemove(),
		}

	case *HeaderMutation_Append:

		if h, ok := interface{}(m.GetAppend()).(clone.Cloner); ok {
			target.Action = &HeaderMutation_Append{
				Append: h.Clone().(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.HeaderValueOption),
			}
		} else {
			target.Action = &HeaderMutation_Append{
				Append: proto.Clone(m.GetAppend()).(*github_com_solo_io_solo_apis_pkg_api_gloo_solo_io_external_envoy_config_core_v3.HeaderValueOption),
			}
		}

	}

	return target
}
