// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for ratelimit.solo.io/v1alpha1 resources

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for RateLimitConfig

func (in *RateLimitConfig) DeepCopyInto(out *RateLimitConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *RateLimitConfig) DeepCopy() *RateLimitConfig {
	if in == nil {
		return nil
	}
	out := new(RateLimitConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *RateLimitConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *RateLimitConfigList) DeepCopyInto(out *RateLimitConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimitConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *RateLimitConfigList) DeepCopy() *RateLimitConfigList {
	if in == nil {
		return nil
	}
	out := new(RateLimitConfigList)
	in.DeepCopyInto(out)
	return out
}

func (in *RateLimitConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
