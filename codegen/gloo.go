package codegen

import (
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/solo-io/skv2/codegen/render"

	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/util"
)

func init() {
	groups = append(groups, GlooGroups()...)
}

func GlooGroups() []model.Group {
	return []model.Group{
		makeGroup("gloo.apis", "v1", []resourceToGenerate{
			{kind: "Settings",
				protoPackage: "gloo.apis.solo.io",
			},
			{kind: "Upstream",
				protoPackage: "gloo.apis.solo.io",
			},
			{kind: "UpstreamGroup",
				protoPackage: "gloo.apis.solo.io",
			},
			{kind: "Proxy",
				protoPackage: "gloo.apis.solo.io",
			},
		}, GlooCustomTemplates),
		makeGroup("gateway.apis", "v1", []resourceToGenerate{
			{kind: "Gateway"},
			{kind: "RouteTable"},
			{kind: "VirtualService"},
			{kind: "VirtualHostOption"},
			{kind: "RouteOption"},
		}, GlooCustomTemplates),
		makeGroup("enterprise.gloo.apis", "v1", []resourceToGenerate{
			{
				kind:         "AuthConfig",
				protoPackage: "enterprise.gloo.apis.solo.io",
				goPackage:    "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/enterprise/options/extauth/v1",
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
	RegisterTemplatePath       = "register.gotmpl"
	RegisterOutputFilename     = "solo_apis_register.go"
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

var RegisterOriginalGroup = func() model.CustomTemplates {
	templateContentsBytes, err := ioutil.ReadFile(templatesDir + RegisterTemplatePath)
	if err != nil {
		panic(err)
	}
	templateContents := string(templateContentsBytes)
	originalGroupTemplates := model.CustomTemplates{
		Templates: map[string]string{RegisterOutputFilename: templateContents},
		Funcs:     GetTemplateFuncs(),
	}
	GlooCustomTemplates = append(GlooCustomTemplates, originalGroupTemplates)
	return originalGroupTemplates
}()

func GetTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"original_group": func(grp render.Group) string {
			return strings.ReplaceAll(grp.Group, ".apis", "")
		},
	}
}
