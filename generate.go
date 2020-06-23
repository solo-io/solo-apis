package main

import (
	"github.com/solo-io/skv2/contrib"
	"github.com/solo-io/solo-kit/pkg/code-generator/sk_anyvendor"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"

	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
)

func main() {
	log.Println("starting generate")

	skv2Cmd := codegen.Command{
		AnyVendorConfig: sk_anyvendor.CreateDefaultMatchOptions([]string{"**/*.proto"}),
		Groups: []model.Group{
			{
				GroupVersion: schema.GroupVersion{
					Group:   "ratelimit.solo.io",
					Version: "v1alpha1",
				},
				Module:  "github.com/solo-io/solo-apis",
				ApiRoot: "pkg",
				Resources: []model.Resource{
					{
						Kind: "RateLimitConfig",
						Spec: model.Field{
							Type: model.Type{
								Name:      "RateLimitConfigSpec",
								GoPackage: "github.com/solo-io/solo-apis/ratelimit.solo.io/v1alpha1/types",
							},
						},
						Status: &model.Field{
							Type: model.Type{
								Name:      "RateLimitConfigStatus",
								GoPackage: "github.com/solo-io/solo-apis/ratelimit.solo.io/v1alpha1/types",
							},
						},
					},
				},
				RenderTypes:      true,
				RenderClients:    true,
				RenderController: true,
				MockgenDirective: true,
				RenderProtos:     true,
				CustomTemplates:  contrib.AllCustomTemplates,
			},
		},
	}
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}

	log.Println("Finished generating code")
}
