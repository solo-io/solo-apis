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
		}, []model.CustomTemplates{}),
	}
}

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
	// register sets
	GlooCustomTemplates = append(GlooCustomTemplates, jsonStatusTemplates)

	return jsonStatusTemplates
}()
