// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for networking.gloo.solo.io/v2alpha1 resources

package v2alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for ExternalWorkload

func (in *ExternalWorkload) DeepCopyInto(out *ExternalWorkload) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *ExternalWorkload) DeepCopy() *ExternalWorkload {
	if in == nil {
		return nil
	}
	out := new(ExternalWorkload)
	in.DeepCopyInto(out)
	return out
}

func (in *ExternalWorkload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *ExternalWorkloadList) DeepCopyInto(out *ExternalWorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalWorkload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *ExternalWorkloadList) DeepCopy() *ExternalWorkloadList {
	if in == nil {
		return nil
	}
	out := new(ExternalWorkloadList)
	in.DeepCopyInto(out)
	return out
}

func (in *ExternalWorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}