// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for apimanagement.gloo.solo.io/v2 resources

package v2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for GraphQLStitchedSchema

func (in *GraphQLStitchedSchema) DeepCopyInto(out *GraphQLStitchedSchema) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *GraphQLStitchedSchema) DeepCopy() *GraphQLStitchedSchema {
	if in == nil {
		return nil
	}
	out := new(GraphQLStitchedSchema)
	in.DeepCopyInto(out)
	return out
}

func (in *GraphQLStitchedSchema) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *GraphQLStitchedSchemaList) DeepCopyInto(out *GraphQLStitchedSchemaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GraphQLStitchedSchema, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *GraphQLStitchedSchemaList) DeepCopy() *GraphQLStitchedSchemaList {
	if in == nil {
		return nil
	}
	out := new(GraphQLStitchedSchemaList)
	in.DeepCopyInto(out)
	return out
}

func (in *GraphQLStitchedSchemaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for GraphQLResolverMap

func (in *GraphQLResolverMap) DeepCopyInto(out *GraphQLResolverMap) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *GraphQLResolverMap) DeepCopy() *GraphQLResolverMap {
	if in == nil {
		return nil
	}
	out := new(GraphQLResolverMap)
	in.DeepCopyInto(out)
	return out
}

func (in *GraphQLResolverMap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *GraphQLResolverMapList) DeepCopyInto(out *GraphQLResolverMapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GraphQLResolverMap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *GraphQLResolverMapList) DeepCopy() *GraphQLResolverMapList {
	if in == nil {
		return nil
	}
	out := new(GraphQLResolverMapList)
	in.DeepCopyInto(out)
	return out
}

func (in *GraphQLResolverMapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for GraphQLSchema

func (in *GraphQLSchema) DeepCopyInto(out *GraphQLSchema) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *GraphQLSchema) DeepCopy() *GraphQLSchema {
	if in == nil {
		return nil
	}
	out := new(GraphQLSchema)
	in.DeepCopyInto(out)
	return out
}

func (in *GraphQLSchema) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *GraphQLSchemaList) DeepCopyInto(out *GraphQLSchemaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GraphQLSchema, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *GraphQLSchemaList) DeepCopy() *GraphQLSchemaList {
	if in == nil {
		return nil
	}
	out := new(GraphQLSchemaList)
	in.DeepCopyInto(out)
	return out
}

func (in *GraphQLSchemaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for ApiDoc

func (in *ApiDoc) DeepCopyInto(out *ApiDoc) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *ApiDoc) DeepCopy() *ApiDoc {
	if in == nil {
		return nil
	}
	out := new(ApiDoc)
	in.DeepCopyInto(out)
	return out
}

func (in *ApiDoc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *ApiDocList) DeepCopyInto(out *ApiDocList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiDoc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *ApiDocList) DeepCopy() *ApiDocList {
	if in == nil {
		return nil
	}
	out := new(ApiDocList)
	in.DeepCopyInto(out)
	return out
}

func (in *ApiDocList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for Portal

func (in *Portal) DeepCopyInto(out *Portal) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *Portal) DeepCopy() *Portal {
	if in == nil {
		return nil
	}
	out := new(Portal)
	in.DeepCopyInto(out)
	return out
}

func (in *Portal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *PortalList) DeepCopyInto(out *PortalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Portal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *PortalList) DeepCopy() *PortalList {
	if in == nil {
		return nil
	}
	out := new(PortalList)
	in.DeepCopyInto(out)
	return out
}

func (in *PortalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for PortalGroup

func (in *PortalGroup) DeepCopyInto(out *PortalGroup) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *PortalGroup) DeepCopy() *PortalGroup {
	if in == nil {
		return nil
	}
	out := new(PortalGroup)
	in.DeepCopyInto(out)
	return out
}

func (in *PortalGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *PortalGroupList) DeepCopyInto(out *PortalGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PortalGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *PortalGroupList) DeepCopy() *PortalGroupList {
	if in == nil {
		return nil
	}
	out := new(PortalGroupList)
	in.DeepCopyInto(out)
	return out
}

func (in *PortalGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for ApiSchemaDiscovery

func (in *ApiSchemaDiscovery) DeepCopyInto(out *ApiSchemaDiscovery) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *ApiSchemaDiscovery) DeepCopy() *ApiSchemaDiscovery {
	if in == nil {
		return nil
	}
	out := new(ApiSchemaDiscovery)
	in.DeepCopyInto(out)
	return out
}

func (in *ApiSchemaDiscovery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *ApiSchemaDiscoveryList) DeepCopyInto(out *ApiSchemaDiscoveryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiSchemaDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *ApiSchemaDiscoveryList) DeepCopy() *ApiSchemaDiscoveryList {
	if in == nil {
		return nil
	}
	out := new(ApiSchemaDiscoveryList)
	in.DeepCopyInto(out)
	return out
}

func (in *ApiSchemaDiscoveryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}