// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/protocol_upgrade/protocol_upgrade.proto

package protocol_upgrade

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
func (m *ProtocolUpgradeConfig) Clone() proto.Message {
	var target *ProtocolUpgradeConfig
	if m == nil {
		return target
	}
	target = &ProtocolUpgradeConfig{}

	switch m.UpgradeType.(type) {

	case *ProtocolUpgradeConfig_Websocket:

		if h, ok := interface{}(m.GetWebsocket()).(clone.Cloner); ok {
			target.UpgradeType = &ProtocolUpgradeConfig_Websocket{
				Websocket: h.Clone().(*ProtocolUpgradeConfig_ProtocolUpgradeSpec),
			}
		} else {
			target.UpgradeType = &ProtocolUpgradeConfig_Websocket{
				Websocket: proto.Clone(m.GetWebsocket()).(*ProtocolUpgradeConfig_ProtocolUpgradeSpec),
			}
		}

	}

	return target
}

// Clone function
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) Clone() proto.Message {
	var target *ProtocolUpgradeConfig_ProtocolUpgradeSpec
	if m == nil {
		return target
	}
	target = &ProtocolUpgradeConfig_ProtocolUpgradeSpec{}

	if h, ok := interface{}(m.GetEnabled()).(clone.Cloner); ok {
		target.Enabled = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Enabled = proto.Clone(m.GetEnabled()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}