package codegen

import (
	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/solo-kit/pkg/code-generator/sk_anyvendor"
)

const (
	apiRoot = "pkg/api"
)

var groups []model.Group

func Command() codegen.Command {
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
