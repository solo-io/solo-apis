// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/policy/v2/security/client_tls_policy.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2 "github.com/solo-io/solo-apis/pkg/api/common.gloo.solo.io/v2"
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
func (m *ClientTLSPolicySpec) Clone() proto.Message {
	var target *ClientTLSPolicySpec
	if m == nil {
		return target
	}
	target = &ClientTLSPolicySpec{}

	if m.GetApplyToDestinations() != nil {
		target.ApplyToDestinations = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationSelector, len(m.GetApplyToDestinations()))
		for idx, v := range m.GetApplyToDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ApplyToDestinations[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationSelector)
			} else {
				target.ApplyToDestinations[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationSelector)
			}

		}
	}

	switch m.Mode.(type) {

	case *ClientTLSPolicySpec_Disable_:

		if h, ok := interface{}(m.GetDisable()).(clone.Cloner); ok {
			target.Mode = &ClientTLSPolicySpec_Disable_{
				Disable: h.Clone().(*ClientTLSPolicySpec_Disable),
			}
		} else {
			target.Mode = &ClientTLSPolicySpec_Disable_{
				Disable: proto.Clone(m.GetDisable()).(*ClientTLSPolicySpec_Disable),
			}
		}

	case *ClientTLSPolicySpec_Simple_:

		if h, ok := interface{}(m.GetSimple()).(clone.Cloner); ok {
			target.Mode = &ClientTLSPolicySpec_Simple_{
				Simple: h.Clone().(*ClientTLSPolicySpec_Simple),
			}
		} else {
			target.Mode = &ClientTLSPolicySpec_Simple_{
				Simple: proto.Clone(m.GetSimple()).(*ClientTLSPolicySpec_Simple),
			}
		}

	case *ClientTLSPolicySpec_Mutual_:

		if h, ok := interface{}(m.GetMutual()).(clone.Cloner); ok {
			target.Mode = &ClientTLSPolicySpec_Mutual_{
				Mutual: h.Clone().(*ClientTLSPolicySpec_Mutual),
			}
		} else {
			target.Mode = &ClientTLSPolicySpec_Mutual_{
				Mutual: proto.Clone(m.GetMutual()).(*ClientTLSPolicySpec_Mutual),
			}
		}

	case *ClientTLSPolicySpec_IstioMutual_:

		if h, ok := interface{}(m.GetIstioMutual()).(clone.Cloner); ok {
			target.Mode = &ClientTLSPolicySpec_IstioMutual_{
				IstioMutual: h.Clone().(*ClientTLSPolicySpec_IstioMutual),
			}
		} else {
			target.Mode = &ClientTLSPolicySpec_IstioMutual_{
				IstioMutual: proto.Clone(m.GetIstioMutual()).(*ClientTLSPolicySpec_IstioMutual),
			}
		}

	}

	return target
}

// Clone function
func (m *TLSConfig) Clone() proto.Message {
	var target *TLSConfig
	if m == nil {
		return target
	}
	target = &TLSConfig{}

	if h, ok := interface{}(m.GetSni()).(clone.Cloner); ok {
		target.Sni = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.Sni = proto.Clone(m.GetSni()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	target.CredentialName = m.GetCredentialName()

	return target
}

// Clone function
func (m *ClientTLSPolicyStatus) Clone() proto.Message {
	var target *ClientTLSPolicyStatus
	if m == nil {
		return target
	}
	target = &ClientTLSPolicyStatus{}

	if h, ok := interface{}(m.GetGlobal()).(clone.Cloner); ok {
		target.Global = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericGlobalStatus)
	} else {
		target.Global = proto.Clone(m.GetGlobal()).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.GenericGlobalStatus)
	}

	if m.GetWorkspaces() != nil {
		target.Workspaces = make(map[string]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus, len(m.GetWorkspaces()))
		for k, v := range m.GetWorkspaces() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Workspaces[k] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus)
			} else {
				target.Workspaces[k] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.WorkspaceStatus)
			}

		}
	}

	if m.GetSelectedDestinationPorts() != nil {
		target.SelectedDestinationPorts = make([]*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference, len(m.GetSelectedDestinationPorts()))
		for idx, v := range m.GetSelectedDestinationPorts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SelectedDestinationPorts[idx] = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
			} else {
				target.SelectedDestinationPorts[idx] = proto.Clone(v).(*github_com_solo_io_solo_apis_pkg_api_common_gloo_solo_io_v2.DestinationReference)
			}

		}
	}

	return target
}

// Clone function
func (m *ClientTLSPolicySpec_Disable) Clone() proto.Message {
	var target *ClientTLSPolicySpec_Disable
	if m == nil {
		return target
	}
	target = &ClientTLSPolicySpec_Disable{}

	return target
}

// Clone function
func (m *ClientTLSPolicySpec_IstioMutual) Clone() proto.Message {
	var target *ClientTLSPolicySpec_IstioMutual
	if m == nil {
		return target
	}
	target = &ClientTLSPolicySpec_IstioMutual{}

	return target
}

// Clone function
func (m *ClientTLSPolicySpec_Mutual) Clone() proto.Message {
	var target *ClientTLSPolicySpec_Mutual
	if m == nil {
		return target
	}
	target = &ClientTLSPolicySpec_Mutual{}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*TLSConfig)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*TLSConfig)
	}

	return target
}

// Clone function
func (m *ClientTLSPolicySpec_Simple) Clone() proto.Message {
	var target *ClientTLSPolicySpec_Simple
	if m == nil {
		return target
	}
	target = &ClientTLSPolicySpec_Simple{}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*TLSConfig)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*TLSConfig)
	}

	return target
}
