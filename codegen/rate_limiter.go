package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	groups = append(groups, rateLimiterGroups()...)
}

func rateLimiterGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: schema.GroupVersion{
				Group:   "ratelimit.solo.io",
				Version: "v1alpha1",
			},
			Module: "github.com/solo-io/solo-apis",
			Resources: []model.Resource{
				{
					Kind: "RateLimitConfig",
					Spec: model.Field{
						Type: model.Type{
							Name:      "RateLimitConfigSpec",
							GoPackage: "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1/types",
						},
					},
					Status: &model.Field{
						Type: model.Type{
							Name:      "RateLimitConfigStatus",
							GoPackage: "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1/types",
						},
					},
				},
			},
			RenderTypes:      true,
			RenderClients:    true,
			RenderController: true,
			MockgenDirective: true,
			CustomTemplates:  contrib.AllGroupCustomTemplates,
		},
	}
}
