// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/ratelimit_server_settings.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"
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
func (m *RateLimitServerSettingsSpec) Clone() proto.Message {
	var target *RateLimitServerSettingsSpec
	if m == nil {
		return target
	}
	target = &RateLimitServerSettingsSpec{}

	if h, ok := interface{}(m.GetDestinationServer()).(clone.Cloner); ok {
		target.DestinationServer = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
	} else {
		target.DestinationServer = proto.Clone(m.GetDestinationServer()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.DestinationReference)
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(clone.Cloner); ok {
		target.RequestTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.RequestTimeout = proto.Clone(m.GetRequestTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	target.DenyOnFail = m.GetDenyOnFail()

	return target
}

// Clone function
func (m *RateLimitServerSettingsStatus) Clone() proto.Message {
	var target *RateLimitServerSettingsStatus
	if m == nil {
		return target
	}
	target = &RateLimitServerSettingsStatus{}

	target.ObservedGeneration = m.GetObservedGeneration()

	target.State = m.GetState()

	return target
}