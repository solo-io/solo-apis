// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v2

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)

// DeepCopyInto for the GraphQLPersistedQueryCachePolicy.Spec
func (in *GraphQLPersistedQueryCachePolicySpec) DeepCopyInto(out *GraphQLPersistedQueryCachePolicySpec) {
	var p *GraphQLPersistedQueryCachePolicySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*GraphQLPersistedQueryCachePolicySpec)
	} else {
		p = proto.Clone(in).(*GraphQLPersistedQueryCachePolicySpec)
	}
	*out = *p
}

// DeepCopyInto for the GraphQLPersistedQueryCachePolicy.Status
func (in *GraphQLPersistedQueryCachePolicyStatus) DeepCopyInto(out *GraphQLPersistedQueryCachePolicyStatus) {
	var p *GraphQLPersistedQueryCachePolicyStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*GraphQLPersistedQueryCachePolicyStatus)
	} else {
		p = proto.Clone(in).(*GraphQLPersistedQueryCachePolicyStatus)
	}
	*out = *p
}

// DeepCopyInto for the FailoverPolicy.Spec
func (in *FailoverPolicySpec) DeepCopyInto(out *FailoverPolicySpec) {
	var p *FailoverPolicySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FailoverPolicySpec)
	} else {
		p = proto.Clone(in).(*FailoverPolicySpec)
	}
	*out = *p
}

// DeepCopyInto for the FailoverPolicy.Status
func (in *FailoverPolicyStatus) DeepCopyInto(out *FailoverPolicyStatus) {
	var p *FailoverPolicyStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FailoverPolicyStatus)
	} else {
		p = proto.Clone(in).(*FailoverPolicyStatus)
	}
	*out = *p
}

// DeepCopyInto for the OutlierDetectionPolicy.Spec
func (in *OutlierDetectionPolicySpec) DeepCopyInto(out *OutlierDetectionPolicySpec) {
	var p *OutlierDetectionPolicySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*OutlierDetectionPolicySpec)
	} else {
		p = proto.Clone(in).(*OutlierDetectionPolicySpec)
	}
	*out = *p
}

// DeepCopyInto for the OutlierDetectionPolicy.Status
func (in *OutlierDetectionPolicyStatus) DeepCopyInto(out *OutlierDetectionPolicyStatus) {
	var p *OutlierDetectionPolicyStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*OutlierDetectionPolicyStatus)
	} else {
		p = proto.Clone(in).(*OutlierDetectionPolicyStatus)
	}
	*out = *p
}

// DeepCopyInto for the FaultInjectionPolicy.Spec
func (in *FaultInjectionPolicySpec) DeepCopyInto(out *FaultInjectionPolicySpec) {
	var p *FaultInjectionPolicySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FaultInjectionPolicySpec)
	} else {
		p = proto.Clone(in).(*FaultInjectionPolicySpec)
	}
	*out = *p
}

// DeepCopyInto for the FaultInjectionPolicy.Status
func (in *FaultInjectionPolicyStatus) DeepCopyInto(out *FaultInjectionPolicyStatus) {
	var p *FaultInjectionPolicyStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FaultInjectionPolicyStatus)
	} else {
		p = proto.Clone(in).(*FaultInjectionPolicyStatus)
	}
	*out = *p
}

// DeepCopyInto for the RetryTimeoutPolicy.Spec
func (in *RetryTimeoutPolicySpec) DeepCopyInto(out *RetryTimeoutPolicySpec) {
	var p *RetryTimeoutPolicySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RetryTimeoutPolicySpec)
	} else {
		p = proto.Clone(in).(*RetryTimeoutPolicySpec)
	}
	*out = *p
}

// DeepCopyInto for the RetryTimeoutPolicy.Status
func (in *RetryTimeoutPolicyStatus) DeepCopyInto(out *RetryTimeoutPolicyStatus) {
	var p *RetryTimeoutPolicyStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*RetryTimeoutPolicyStatus)
	} else {
		p = proto.Clone(in).(*RetryTimeoutPolicyStatus)
	}
	*out = *p
}

// DeepCopyInto for the ConnectionPolicy.Spec
func (in *ConnectionPolicySpec) DeepCopyInto(out *ConnectionPolicySpec) {
	var p *ConnectionPolicySpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ConnectionPolicySpec)
	} else {
		p = proto.Clone(in).(*ConnectionPolicySpec)
	}
	*out = *p
}

// DeepCopyInto for the ConnectionPolicy.Status
func (in *ConnectionPolicyStatus) DeepCopyInto(out *ConnectionPolicyStatus) {
	var p *ConnectionPolicyStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*ConnectionPolicyStatus)
	} else {
		p = proto.Clone(in).(*ConnectionPolicyStatus)
	}
	*out = *p
}
