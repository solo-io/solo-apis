// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for enterprise.gloo.apis.solo.io/v1 resources

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for AuthConfig

func (in *AuthConfig) DeepCopyInto(out *AuthConfig) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *AuthConfig) DeepCopy() *AuthConfig {
	if in == nil {
		return nil
	}
	out := new(AuthConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *AuthConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *AuthConfigList) DeepCopyInto(out *AuthConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *AuthConfigList) DeepCopy() *AuthConfigList {
	if in == nil {
		return nil
	}
	out := new(AuthConfigList)
	in.DeepCopyInto(out)
	return out
}

func (in *AuthConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
