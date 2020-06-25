package codegen

import (
	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
)

const (
	apiRoot = "pkg"
)

var groups []model.Group

func Command() codegen.Command {
	for i, group := range groups {
		group.ApiRoot = apiRoot
		groups[i] = group
	}

	return codegen.Command{
		Groups: groups,
	}
}
