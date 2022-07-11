// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-mesh/gloo.solo.io/internal/v2/issued_certificate.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"

	github_com_solo_io_solo_apis_pkg_api_security_policy_gloo_solo_io_v2_tls "github.com/solo-io/solo-apis/pkg/api/security.policy.gloo.solo.io/v2/tls"
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
func (m *IssuedCertificateSpec) Clone() proto.Message {
	var target *IssuedCertificateSpec
	if m == nil {
		return target
	}
	target = &IssuedCertificateSpec{}

	if m.GetHosts() != nil {
		target.Hosts = make([]string, len(m.GetHosts()))
		for idx, v := range m.GetHosts() {

			target.Hosts[idx] = v

		}
	}

	if h, ok := interface{}(m.GetIssuedCertificateSecret()).(clone.Cloner); ok {
		target.IssuedCertificateSecret = h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef)
	} else {
		target.IssuedCertificateSecret = proto.Clone(m.GetIssuedCertificateSecret()).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef)
	}

	if h, ok := interface{}(m.GetCertOptions()).(clone.Cloner); ok {
		target.CertOptions = h.Clone().(*github_com_solo_io_solo_apis_pkg_api_security_policy_gloo_solo_io_v2_tls.CommonCertOptions)
	} else {
		target.CertOptions = proto.Clone(m.GetCertOptions()).(*github_com_solo_io_solo_apis_pkg_api_security_policy_gloo_solo_io_v2_tls.CommonCertOptions)
	}

	if h, ok := interface{}(m.GetMeshRef()).(clone.Cloner); ok {
		target.MeshRef = h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef)
	} else {
		target.MeshRef = proto.Clone(m.GetMeshRef()).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef)
	}

	target.AutoRestartPods = m.GetAutoRestartPods()

	switch m.CertificateAuthority.(type) {

	case *IssuedCertificateSpec_MgmtServerCa:

		if h, ok := interface{}(m.GetMgmtServerCa()).(clone.Cloner); ok {
			target.CertificateAuthority = &IssuedCertificateSpec_MgmtServerCa{
				MgmtServerCa: h.Clone().(*MgmtServerCertificateAuthority),
			}
		} else {
			target.CertificateAuthority = &IssuedCertificateSpec_MgmtServerCa{
				MgmtServerCa: proto.Clone(m.GetMgmtServerCa()).(*MgmtServerCertificateAuthority),
			}
		}

	case *IssuedCertificateSpec_AgentCa:

		if h, ok := interface{}(m.GetAgentCa()).(clone.Cloner); ok {
			target.CertificateAuthority = &IssuedCertificateSpec_AgentCa{
				AgentCa: h.Clone().(*github_com_solo_io_solo_apis_pkg_api_security_policy_gloo_solo_io_v2_tls.AgentCertificateAuthority),
			}
		} else {
			target.CertificateAuthority = &IssuedCertificateSpec_AgentCa{
				AgentCa: proto.Clone(m.GetAgentCa()).(*github_com_solo_io_solo_apis_pkg_api_security_policy_gloo_solo_io_v2_tls.AgentCertificateAuthority),
			}
		}

	}

	return target
}

// Clone function
func (m *MgmtServerCertificateAuthority) Clone() proto.Message {
	var target *MgmtServerCertificateAuthority
	if m == nil {
		return target
	}
	target = &MgmtServerCertificateAuthority{}

	switch m.CertificateAuthority.(type) {

	case *MgmtServerCertificateAuthority_SigningCertificateSecret:

		if h, ok := interface{}(m.GetSigningCertificateSecret()).(clone.Cloner); ok {
			target.CertificateAuthority = &MgmtServerCertificateAuthority_SigningCertificateSecret{
				SigningCertificateSecret: h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef),
			}
		} else {
			target.CertificateAuthority = &MgmtServerCertificateAuthority_SigningCertificateSecret{
				SigningCertificateSecret: proto.Clone(m.GetSigningCertificateSecret()).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef),
			}
		}

	}

	return target
}

// Clone function
func (m *IssuedCertificateStatus) Clone() proto.Message {
	var target *IssuedCertificateStatus
	if m == nil {
		return target
	}
	target = &IssuedCertificateStatus{}

	target.ObservedGeneration = m.GetObservedGeneration()

	target.Error = m.GetError()

	target.State = m.GetState()

	return target
}