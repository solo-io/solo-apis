package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	module = "github.com/solo-io/solo-apis"
)

type resourceToGenerate struct {
	kind         string
	protoPackage string // set if different from group name
	goPackage    string // Set if different than root folder
	noStatus     bool   // don't put a status on this resource
}

func makeGroup(
	groupPrefix, version string,
	resourcesToGenerate []resourceToGenerate,
	customTemplates []model.CustomTemplates,
) model.Group {
	var resources []model.Resource
	for _, resource := range resourcesToGenerate {
		res := model.Resource{
			Kind: resource.kind,
			Spec: model.Field{
				Type: model.Type{
					Name:         resource.kind + "Spec",
					ProtoPackage: resource.protoPackage,
					GoPackage:    resource.goPackage,
				},
			},
			Stored: true,
		}
		if !resource.noStatus {
			res.Status = &model.Field{Type: model.Type{
				Name:         resource.kind + "Status",
				ProtoPackage: resource.protoPackage,
				GoPackage:    resource.goPackage,
			}}
		}
		resources = append(resources, res)
	}

	customGroupTemplates := append(contrib.AllGroupCustomTemplates, customTemplates...)

	return model.Group{
		GroupVersion: schema.GroupVersion{
			Group:   groupPrefix + "." + "solo.io",
			Version: version,
		},
		Module:                    module,
		Resources:                 resources,
		RenderManifests:           true,
		RenderTypes:               true,
		RenderClients:             true,
		RenderController:          true,
		MockgenDirective:          true,
		CustomTemplates:           customGroupTemplates,
		ApiRoot:                   apiRoot,
		SkipConditionalCRDLoading: true, // we want the alpha crds always rendered
		SkipTemplatedCRDManifest:  true, // do not make a copy of crds in templates dir
	}
}
