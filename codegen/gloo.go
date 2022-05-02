package codegen

import (
	"io/ioutil"

	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/util"
)

func init() {
	groups = append(groups, GlooGroups()...)
}

func GlooGroups() []model.Group {
	return []model.Group{
		makeGroup("gloo", "v1", []resourceToGenerate{
			{kind: "Settings"},
			{kind: "Upstream"},
			{kind: "UpstreamGroup"},
			{kind: "Proxy"},
		}, GlooCustomTemplates),
		makeGroup("gateway", "v1", []resourceToGenerate{
			{kind: "Gateway"},
			{kind: "HttpGateway"},
			{kind: "RouteTable"},
			{kind: "VirtualService"},
			{kind: "VirtualHostOption"},
			{kind: "RouteOption"},
		}, GlooCustomTemplates),
		makeGroup("enterprise.gloo", "v1", []resourceToGenerate{
			{
				kind:         "AuthConfig",
				protoPackage: "enterprise.gloo.solo.io",
				goPackage:    "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1",
			},
		}, GlooCustomTemplates),
		makeGroup("graphql.gloo", "v1beta1", []resourceToGenerate{
			{
				kind:         "GraphQLApi",
				protoPackage: "graphql.gloo.solo.io",
				goPackage:    "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/graphql/v1beta1",
			},
		}, GlooCustomTemplates),
	}
}

// Gloo resources, backed by solo-kit, support reporting statuses for multiple controllers (1 per namespace)
// Gloo-Fed resources, backed by skv2, do not yet. To reduce the complexity of the change, we
// chose to not introduce namespaced statuses support for gloo-fed resources as part of this PR.
// In order to handle these two cases simultaneously, we define custom status unmarshallers for Gloo-Fed resources.

var GlooCustomTemplates []model.CustomTemplates

var templatesDir = util.MustGetThisDir() + "/templates/"

const (
	GlooJsonOutputFilename     = "gloo_json.gen.go"
	GlooJsonCustomTemplatePath = "gloo_json.gen.gotmpl"
)

/*
JsonStatuses custom template
*/
var JsonStatuses = func() model.CustomTemplates {
	templateContentsBytes, err := ioutil.ReadFile(templatesDir + GlooJsonCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	templateContents := string(templateContentsBytes)
	jsonStatusTemplates := model.CustomTemplates{
		Templates: map[string]string{GlooJsonOutputFilename: templateContents},
	}
	// register json statuses
	GlooCustomTemplates = append(GlooCustomTemplates, jsonStatusTemplates)

	return jsonStatusTemplates
}()
