package main

import (
	"github.com/solo-io/skv2/contrib"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"

	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
)

func main() {
	log.Println("starting generate")

	skv2Cmd := codegen.Command{
		Groups: []model.Group{
			{
				GroupVersion: schema.GroupVersion{
					Group:   "ratelimit.solo.io",
					Version: "v1alpha1",
				},
				//Module: "github.com/solo-io/rate-limiter",
				Resources: []model.Resource{
					{
						Kind: "RateLimitConfig",
						Spec: model.Field{
							Type: model.Type{
								Name:      "RateLimitConfigSpec",
								GoPackage: "github.com/solo-io/solo-apis/pkg/ratelimit.solo.io/v1alpha1/types",
							},
						},
						Status: &model.Field{
							Type: model.Type{
								Name:      "RateLimitConfigStatus",
								GoPackage: "github.com/solo-io/solo-apis/pkg/ratelimit.solo.io/v1alpha1/types",
							},
						},
					},
				},
				RenderTypes:      true,
				RenderClients:    true,
				RenderController: true,
				MockgenDirective: true,
				//RenderProtos:     true,
				CustomTemplates: contrib.AllCustomTemplates,
				ApiRoot:         "pkg/api",
			},
		},
	}
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}

	log.Println("Finished generating code")
}
