// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/internal/v2alpha1/spire.proto

package v2alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
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
func (m *SpireRegistrationEntrySpec) Clone() proto.Message {
	var target *SpireRegistrationEntrySpec
	if m == nil {
		return target
	}
	target = &SpireRegistrationEntrySpec{}

	if m.GetSelectors() != nil {
		target.Selectors = make([]*SpireRegistrationEntrySpec_Selector, len(m.GetSelectors()))
		for idx, v := range m.GetSelectors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Selectors[idx] = h.Clone().(*SpireRegistrationEntrySpec_Selector)
			} else {
				target.Selectors[idx] = proto.Clone(v).(*SpireRegistrationEntrySpec_Selector)
			}

		}
	}

	target.ParentId = m.GetParentId()

	target.SpiffeId = m.GetSpiffeId()

	if h, ok := interface{}(m.GetX509SvidTtl()).(clone.Cloner); ok {
		target.X509SvidTtl = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.X509SvidTtl = proto.Clone(m.GetX509SvidTtl()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if m.GetDnsNames() != nil {
		target.DnsNames = make([]string, len(m.GetDnsNames()))
		for idx, v := range m.GetDnsNames() {

			target.DnsNames[idx] = v

		}
	}

	return target
}

// Clone function
func (m *SpireRegistrationEntryStatus) Clone() proto.Message {
	var target *SpireRegistrationEntryStatus
	if m == nil {
		return target
	}
	target = &SpireRegistrationEntryStatus{}

	if h, ok := interface{}(m.GetCommon()).(clone.Cloner); ok {
		target.Common = h.Clone().(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	} else {
		target.Common = proto.Clone(m.GetCommon()).(*github_com_solo_io_gloo_mesh_solo_apis_client_go_common_gloo_solo_io_v2.Status)
	}

	return target
}

// Clone function
func (m *SpireRegistrationEntrySpec_Selector) Clone() proto.Message {
	var target *SpireRegistrationEntrySpec_Selector
	if m == nil {
		return target
	}
	target = &SpireRegistrationEntrySpec_Selector{}

	target.Type = m.GetType()

	target.Value = m.GetValue()

	return target
}
