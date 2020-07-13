package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	module = "github.com/solo-io/gloo-fed"
)

type resourceToGenerate struct {
	kind         string
	protoPackage string // set if different from group name
	noStatus     bool   // don't put a status on this resource
}

func makeGroup(
	groupPrefix, version string,
	resourcesToGenerate []resourceToGenerate,
) model.Group {
	var resources []model.Resource
	for _, resource := range resourcesToGenerate {
		res := model.Resource{
			Kind: resource.kind,
			Spec: model.Field{
				Type: model.Type{
					Name:         resource.kind + "Spec",
					ProtoPackage: resource.protoPackage,
				},
			},
		}
		if !resource.noStatus {
			res.Status = &model.Field{Type: model.Type{
				Name:         resource.kind + "Status",
				ProtoPackage: resource.protoPackage,
			}}
		}
		resources = append(resources, res)
	}

	return model.Group{
		GroupVersion: schema.GroupVersion{
			Group:   groupPrefix + "." + "solo.io",
			Version: version,
		},
		Module:           module,
		Resources:        resources,
		RenderManifests:  true,
		RenderTypes:      true,
		RenderClients:    true,
		RenderController: true,
		MockgenDirective: true,
		CustomTemplates:  contrib.AllGroupCustomTemplates,
		ApiRoot:          apiRoot,
	}
}
