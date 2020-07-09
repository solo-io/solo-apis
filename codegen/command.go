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

	return codegen.Command{
		AnyVendorConfig: sk_anyvendor.CreateDefaultMatchOptions([]string{"**/*.proto"}),
		RenderProtos:    true,
		Groups:          groups,
	}
}
