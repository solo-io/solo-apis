package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	groups = append(groups, GlooFedGroups()...)
}

func GlooFedGroups() []model.Group {
	return []model.Group{
		makeFedGroup(),
		makeFedGatewayGroup(),
		makeFedGlooGroup(),
		makeFedEnterpriseGlooGroup(),
		makeFedRateLimitGroup(),
	}
}

func makeFedGroup() model.Group {
	return makeGroup("fed", "v1", []resourceToGenerate{
		{kind: "GlooInstance"},
		{kind: "FailoverScheme"},
	}, []model.CustomTemplates{})
}

func makeFedGatewayGroup() model.Group {
	return makeGroup("fed.gateway", "v1", []resourceToGenerate{
		{kind: "FederatedGateway"},
		{kind: "FederatedRouteTable"},
		{kind: "FederatedVirtualService"},
	}, []model.CustomTemplates{})
}

func makeFedGlooGroup() model.Group {
	return makeGroup("fed.gloo", "v1", []resourceToGenerate{
		{kind: "FederatedSettings"},
		{kind: "FederatedUpstream"},
		{kind: "FederatedUpstreamGroup"},
	}, []model.CustomTemplates{})
}

func makeFedEnterpriseGlooGroup() model.Group {
	return makeGroup("fed.enterprise.gloo", "v1", []resourceToGenerate{
		{kind: "FederatedAuthConfig"},
	}, []model.CustomTemplates{})
}

func makeFedRateLimitGroup() model.Group {
	return makeGroup("fed.ratelimit", "v1alpha1", []resourceToGenerate{
		{kind: "FederatedRateLimitConfig"},
	}, []model.CustomTemplates{})
}
