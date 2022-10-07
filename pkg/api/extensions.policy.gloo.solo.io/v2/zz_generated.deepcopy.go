// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for extensions.policy.gloo.solo.io/v2 resources

package v2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for WasmDeploymentPolicy

func (in *WasmDeploymentPolicy) DeepCopyInto(out *WasmDeploymentPolicy) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *WasmDeploymentPolicy) DeepCopy() *WasmDeploymentPolicy {
	if in == nil {
		return nil
	}
	out := new(WasmDeploymentPolicy)
	in.DeepCopyInto(out)
	return out
}

func (in *WasmDeploymentPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *WasmDeploymentPolicyList) DeepCopyInto(out *WasmDeploymentPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WasmDeploymentPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *WasmDeploymentPolicyList) DeepCopy() *WasmDeploymentPolicyList {
	if in == nil {
		return nil
	}
	out := new(WasmDeploymentPolicyList)
	in.DeepCopyInto(out)
	return out
}

func (in *WasmDeploymentPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}