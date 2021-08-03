// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for multicluster.solo.io/v1alpha1 resources

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for MultiClusterRole

func (in *MultiClusterRole) DeepCopyInto(out *MultiClusterRole) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *MultiClusterRole) DeepCopy() *MultiClusterRole {
	if in == nil {
		return nil
	}
	out := new(MultiClusterRole)
	in.DeepCopyInto(out)
	return out
}

func (in *MultiClusterRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *MultiClusterRoleList) DeepCopyInto(out *MultiClusterRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *MultiClusterRoleList) DeepCopy() *MultiClusterRoleList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterRoleList)
	in.DeepCopyInto(out)
	return out
}

func (in *MultiClusterRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
