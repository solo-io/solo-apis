// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for fed.solo.io/v1 resources

package v1

import (
    runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for GlooInstance

func (in *GlooInstance) DeepCopyInto(out *GlooInstance) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *GlooInstance) DeepCopy() *GlooInstance {
    if in == nil {
        return nil
    }
    out := new(GlooInstance)
    in.DeepCopyInto(out)
    return out
}

func (in *GlooInstance) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *GlooInstanceList) DeepCopyInto(out *GlooInstanceList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]GlooInstance, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *GlooInstanceList) DeepCopy() *GlooInstanceList {
    if in == nil {
        return nil
    }
    out := new(GlooInstanceList)
    in.DeepCopyInto(out)
    return out
}

func (in *GlooInstanceList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

// Generated Deepcopy methods for FailoverScheme

func (in *FailoverScheme) DeepCopyInto(out *FailoverScheme) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *FailoverScheme) DeepCopy() *FailoverScheme {
    if in == nil {
        return nil
    }
    out := new(FailoverScheme)
    in.DeepCopyInto(out)
    return out
}

func (in *FailoverScheme) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *FailoverSchemeList) DeepCopyInto(out *FailoverSchemeList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]FailoverScheme, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *FailoverSchemeList) DeepCopy() *FailoverSchemeList {
    if in == nil {
        return nil
    }
    out := new(FailoverSchemeList)
    in.DeepCopyInto(out)
    return out
}

func (in *FailoverSchemeList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

