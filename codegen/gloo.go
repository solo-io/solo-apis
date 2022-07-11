package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	groups = append(groups, GlooGroups()...)
}

func GlooGroups() []model.Group {
	return []model.Group{
		makeGroup("enterprise", "v1", []resourceToGenerate{
			{
				kind:         "AuthConfig",
				protoPackage: "enterprise.gloo.solo.io",
				goPackage:    "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1",
			},
		}),
	}
}
