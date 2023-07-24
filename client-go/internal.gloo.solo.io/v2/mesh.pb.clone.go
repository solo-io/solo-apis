// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/internal/v2/mesh.proto

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
func (m *MeshSpec) Clone() proto.Message {
	var target *MeshSpec
	if m == nil {
		return target
	}
	target = &MeshSpec{}

	if h, ok := interface{}(m.GetInstallation()).(clone.Cloner); ok {
		target.Installation = h.Clone().(*MeshSpec_Installation)
	} else {
		target.Installation = proto.Clone(m.GetInstallation()).(*MeshSpec_Installation)
	}

	target.TrustDomain = m.GetTrustDomain()

	target.IstiodServiceAccount = m.GetIstiodServiceAccount()

	if m.GetDiscoveryNamespaces() != nil {
		target.DiscoveryNamespaces = make([]string, len(m.GetDiscoveryNamespaces()))
		for idx, v := range m.GetDiscoveryNamespaces() {

			target.DiscoveryNamespaces[idx] = v

		}
	}

	target.SmartDnsProxyingEnabled = m.GetSmartDnsProxyingEnabled()

	target.RootNamespace = m.GetRootNamespace()

	if h, ok := interface{}(m.GetAgentInfo()).(clone.Cloner); ok {
		target.AgentInfo = h.Clone().(*MeshSpec_AgentInfo)
	} else {
		target.AgentInfo = proto.Clone(m.GetAgentInfo()).(*MeshSpec_AgentInfo)
	}

	target.Hub = m.GetHub()

	target.Tag = m.GetTag()

	target.IpFamily = m.GetIpFamily()

	target.AmbientCapable = m.GetAmbientCapable()

	return target
}

// Clone function
func (m *MeshStatus) Clone() proto.Message {
	var target *MeshStatus
	if m == nil {
		return target
	}
	target = &MeshStatus{}

	target.ObservedGeneration = m.GetObservedGeneration()

	return target
}

// Clone function
func (m *MeshSpec_Installation) Clone() proto.Message {
	var target *MeshSpec_Installation
	if m == nil {
		return target
	}
	target = &MeshSpec_Installation{}

	target.Namespace = m.GetNamespace()

	target.Cluster = m.GetCluster()

	if m.GetPodLabels() != nil {
		target.PodLabels = make(map[string]string, len(m.GetPodLabels()))
		for k, v := range m.GetPodLabels() {

			target.PodLabels[k] = v

		}
	}

	target.Version = m.GetVersion()

	target.Revision = m.GetRevision()

	return target
}

// Clone function
func (m *MeshSpec_AgentInfo) Clone() proto.Message {
	var target *MeshSpec_AgentInfo
	if m == nil {
		return target
	}
	target = &MeshSpec_AgentInfo{}

	target.Namespace = m.GetNamespace()

	target.Version = m.GetVersion()

	if h, ok := interface{}(m.GetRelayRootTlsSecret()).(clone.Cloner); ok {
		target.RelayRootTlsSecret = h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef)
	} else {
		target.RelayRootTlsSecret = proto.Clone(m.GetRelayRootTlsSecret()).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.ObjectRef)
	}

	return target
}
