// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1

import (
    proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)
// DeepCopyInto for the AuthConfig.Spec
func (in *AuthConfigSpec) DeepCopyInto(out *AuthConfigSpec) {
    var p *AuthConfigSpec
    if h, ok := interface{}(in).(clone.Cloner); ok {
        p = h.Clone().(*AuthConfigSpec)
    } else {
        p = proto.Clone(in).(*AuthConfigSpec)
    }
    *out = *p
}
// DeepCopyInto for the AuthConfig.Status
func (in *AuthConfigStatus) DeepCopyInto(out *AuthConfigStatus) {
    var p *AuthConfigStatus
    if h, ok := interface{}(in).(clone.Cloner); ok {
        p = h.Clone().(*AuthConfigStatus)
    } else {
        p = proto.Clone(in).(*AuthConfigStatus)
    }
    *out = *p
}
