package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	groups = append(groups, GlooGroups()...)
}

func GlooGroups() []model.Group {
	return []model.Group{
		makeGroup("gloo", "v1", []resourceToGenerate{
			{kind: "AuthConfig"},
			{kind: "Settings", noStatus: true},
			{kind: "Upstream"},
			{kind: "UpstreamGroup"},
			{kind: "Proxy"},
		}),
		makeGroup("gateway", "v1", []resourceToGenerate{
			{kind: "Gateway"},
			{kind: "RouteTable"},
			{kind: "VirtualService"},
		}),
	}
}
