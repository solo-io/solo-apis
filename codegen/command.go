package codegen

import (
	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/solo-kit/pkg/code-generator/sk_anyvendor"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	apiRoot = "pkg/api"
)

var groups []model.Group

func Command() codegen.Command {
	groups = []model.Group{
		{
			GroupVersion: schema.GroupVersion{
				Group:   "fed.solo.io",
				Version: "v1",
			},
			Module: "github.com/solo-io/solo-apis",
			Resources: []model.Resource{
				{
					Kind: "GlooInstance",
					Spec: model.Field{
						Type: model.Type{
							Name:      "GlooInstanceSpec",
							GoPackage: "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types",
						},
					},
					Status: &model.Field{
						Type: model.Type{
							Name:      "GlooInstanceStatus",
							GoPackage: "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types",
						}},
				},
			},
			RenderTypes:   true,
			RenderClients: true,
			ApiRoot:       "pkg/api",
		},
	}
	for i, group := range groups {
		group.ApiRoot = apiRoot
		groups[i] = group
	}
	skv2Imports := sk_anyvendor.CreateDefaultMatchOptions([]string{
		"**/*.proto",
	})
	skv2Imports.External["github.com/solo-io/skv2"] = []string{
		"api/**/*.proto",
	}

	return codegen.Command{
		AnyVendorConfig: skv2Imports,
		RenderProtos:    true,
		Groups:          groups,
	}
}
