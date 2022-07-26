package ratelimit

import (
	"reflect"

	"github.com/solo-io/solo-kit/pkg/utils/specutils"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources"

	skres "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/solo.io/v1"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"

	types "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"

	"github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
)

var _ resources.CustomInputResource = &RateLimitConfig{}

type RateLimitConfig v1alpha1.RateLimitConfig

func (r *RateLimitConfig) GetMetadata() *core.Metadata {
	return kubeutils.FromKubeMeta(r.ObjectMeta, true)
}

func (r *RateLimitConfig) SetMetadata(meta *core.Metadata) {
	r.ObjectMeta = kubeutils.ToKubeMeta(meta)
}

func (r *RateLimitConfig) Equal(that interface{}) bool {
	return reflect.DeepEqual(r, that)
}

func (r *RateLimitConfig) Clone() *RateLimitConfig {
	ci := v1alpha1.RateLimitConfig(*r)
	ciCopy := ci.DeepCopy()
	newCi := RateLimitConfig(*ciCopy)
	return &newCi
}

func (r *RateLimitConfig) UnmarshalSpec(spec skres.Spec) error {
	return specutils.UnmarshalSpecMapToProto(spec, &r.Spec)
}

func (r *RateLimitConfig) UnmarshalStatus(status skres.Status, unmarshaler resources.StatusUnmarshaler) {
	protoutils.UnmarshalMapToProto(status, &r.Status)
}

func (r *RateLimitConfig) MarshalSpec() (skres.Spec, error) {
	return protoutils.MarshalMapFromProto(&r.Spec)
}

func (r *RateLimitConfig) MarshalStatus() (skres.Status, error) {
	return protoutils.MarshalMapFromProto(&r.Status)
}

// Deprecated
func (r *RateLimitConfig) GetStatus() *core.Status {
	return r.convertRateLimitConfigStatusToSoloKitStatus(&r.Status)
}

// Deprecated
func (r *RateLimitConfig) SetStatus(status *core.Status) {
	if status != nil {
		r.Status = *r.convertSoloKitStatusToRateLimitConfigStatus(status)
	}
}

func (r *RateLimitConfig) GetNamespacedStatuses() *core.NamespacedStatuses {
	panic("implement me")
}

func (r *RateLimitConfig) SetNamespacedStatuses(status *core.NamespacedStatuses) {
	panic("implement me")
}

func (r *RateLimitConfig) convertSoloKitStatusToRateLimitConfigStatus(status *core.Status) *v1alpha1.RateLimitConfigStatus {
	if status == nil {
		return nil
	}

	var outputState types.RateLimitConfigStatus_State
	switch status.GetState() {
	case core.Status_Pending:
		outputState = types.RateLimitConfigStatus_PENDING
	case core.Status_Accepted:
		outputState = types.RateLimitConfigStatus_ACCEPTED
	case core.Status_Rejected:
		outputState = types.RateLimitConfigStatus_REJECTED
	case core.Status_Warning:
		// should lever happen
		panic("cannot set WARNING status on RateLimitConfig resources")
	}

	return &v1alpha1.RateLimitConfigStatus{
		State:              outputState,
		Message:            status.GetReason(),
		ObservedGeneration: r.GetGeneration(),
	}
}

func (r *RateLimitConfig) convertRateLimitConfigStatusToSoloKitStatus(status *v1alpha1.RateLimitConfigStatus) *core.Status {
	if status == nil {
		return nil
	}

	var outputState core.Status_State
	switch status.GetState() {
	case types.RateLimitConfigStatus_PENDING:
		outputState = core.Status_Pending
	case types.RateLimitConfigStatus_ACCEPTED:
		outputState = core.Status_Accepted
	case types.RateLimitConfigStatus_REJECTED:
		outputState = core.Status_Rejected
	}

	return &core.Status{
		State:  outputState,
		Reason: status.GetMessage(),
	}
}
