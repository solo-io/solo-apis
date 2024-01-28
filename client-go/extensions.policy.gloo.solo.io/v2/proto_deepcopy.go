// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v2

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)

// DeepCopyInto for the WasmDeploymentPolicy.Spec
func (in *WasmDeploymentPolicySpec) DeepCopyInto(out *WasmDeploymentPolicySpec) {
	var p *WasmDeploymentPolicySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*WasmDeploymentPolicySpec)
	} else {
		p = proto.Clone(in).(*WasmDeploymentPolicySpec)
	}
	*out = *p
}

// DeepCopyInto for the WasmDeploymentPolicy.Status
func (in *WasmDeploymentPolicyStatus) DeepCopyInto(out *WasmDeploymentPolicyStatus) {
	var p *WasmDeploymentPolicyStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*WasmDeploymentPolicyStatus)
	} else {
		p = proto.Clone(in).(*WasmDeploymentPolicyStatus)
	}
	*out = *p
}