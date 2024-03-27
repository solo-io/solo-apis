// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for fed.gateway.solo.io/v1 resources

package v1

import (
    runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for FederatedGateway

func (in *FederatedGateway) DeepCopyInto(out *FederatedGateway) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *FederatedGateway) DeepCopy() *FederatedGateway {
    if in == nil {
        return nil
    }
    out := new(FederatedGateway)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedGateway) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *FederatedGatewayList) DeepCopyInto(out *FederatedGatewayList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]FederatedGateway, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *FederatedGatewayList) DeepCopy() *FederatedGatewayList {
    if in == nil {
        return nil
    }
    out := new(FederatedGatewayList)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedGatewayList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

// Generated Deepcopy methods for FederatedMatchableHttpGateway

func (in *FederatedMatchableHttpGateway) DeepCopyInto(out *FederatedMatchableHttpGateway) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *FederatedMatchableHttpGateway) DeepCopy() *FederatedMatchableHttpGateway {
    if in == nil {
        return nil
    }
    out := new(FederatedMatchableHttpGateway)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedMatchableHttpGateway) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *FederatedMatchableHttpGatewayList) DeepCopyInto(out *FederatedMatchableHttpGatewayList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]FederatedMatchableHttpGateway, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *FederatedMatchableHttpGatewayList) DeepCopy() *FederatedMatchableHttpGatewayList {
    if in == nil {
        return nil
    }
    out := new(FederatedMatchableHttpGatewayList)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedMatchableHttpGatewayList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

// Generated Deepcopy methods for FederatedMatchableTcpGateway

func (in *FederatedMatchableTcpGateway) DeepCopyInto(out *FederatedMatchableTcpGateway) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *FederatedMatchableTcpGateway) DeepCopy() *FederatedMatchableTcpGateway {
    if in == nil {
        return nil
    }
    out := new(FederatedMatchableTcpGateway)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedMatchableTcpGateway) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *FederatedMatchableTcpGatewayList) DeepCopyInto(out *FederatedMatchableTcpGatewayList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]FederatedMatchableTcpGateway, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *FederatedMatchableTcpGatewayList) DeepCopy() *FederatedMatchableTcpGatewayList {
    if in == nil {
        return nil
    }
    out := new(FederatedMatchableTcpGatewayList)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedMatchableTcpGatewayList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

// Generated Deepcopy methods for FederatedRouteTable

func (in *FederatedRouteTable) DeepCopyInto(out *FederatedRouteTable) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *FederatedRouteTable) DeepCopy() *FederatedRouteTable {
    if in == nil {
        return nil
    }
    out := new(FederatedRouteTable)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedRouteTable) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *FederatedRouteTableList) DeepCopyInto(out *FederatedRouteTableList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]FederatedRouteTable, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *FederatedRouteTableList) DeepCopy() *FederatedRouteTableList {
    if in == nil {
        return nil
    }
    out := new(FederatedRouteTableList)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedRouteTableList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

// Generated Deepcopy methods for FederatedVirtualService

func (in *FederatedVirtualService) DeepCopyInto(out *FederatedVirtualService) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *FederatedVirtualService) DeepCopy() *FederatedVirtualService {
    if in == nil {
        return nil
    }
    out := new(FederatedVirtualService)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedVirtualService) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *FederatedVirtualServiceList) DeepCopyInto(out *FederatedVirtualServiceList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]FederatedVirtualService, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *FederatedVirtualServiceList) DeepCopy() *FederatedVirtualServiceList {
    if in == nil {
        return nil
    }
    out := new(FederatedVirtualServiceList)
    in.DeepCopyInto(out)
    return out
}

func (in *FederatedVirtualServiceList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

