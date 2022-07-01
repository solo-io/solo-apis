package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	groups = append(groups, GlooMeshGroups()...)
}

func GlooMeshGroups() []model.Group {
	return []model.Group{
		makeGroup("admin", "v2", []resourceToGenerate{
			{kind: "Workspace"},
			{kind: "WorkspaceSettings"},
			{kind: "KubernetesCluster"},
			{kind: "RootTrustPolicy"},
			{kind: "ExtAuthServer"},
			{kind: "RateLimitServerSettings"},
			{kind: "RateLimitServerConfig"},
			{kind: "Dashboard"},
		}),
		makeGroup("networking", "v2", []resourceToGenerate{
			{kind: "ExternalService"},
			{kind: "ExternalEndpoint"},
			{kind: "RouteTable"},
			{kind: "VirtualDestination"},
			{kind: "VirtualGateway"},
		}),
		makeGroup("extensions", "v2", []resourceToGenerate{
			{kind: "WasmDeploymentPolicy"},
		}),
		makeGroup("security", "v2", []resourceToGenerate{
			{kind: "AccessPolicy"},
			{kind: "CORSPolicy"},
			{kind: "CSRFPolicy"},
			{kind: "ExtAuthPolicy"},
			{kind: "WAFPolicy"},
			{kind: "JWTPolicy"},
		}),
		makeGroup("observability", "v2", []resourceToGenerate{
			{kind: "AccessLogPolicy"},
		}),
		makeGroup("resilience", "v2", []resourceToGenerate{
			{kind: "FailoverPolicy"},
			{kind: "OutlierDetectionPolicy"},
			{kind: "FaultInjectionPolicy"},
			{kind: "RetryTimeoutPolicy"},
		}),
		makeGroup("trafficcontrol", "v2", []resourceToGenerate{
			{kind: "MirrorPolicy"},
			{kind: "RateLimitPolicy"},
			{kind: "RateLimitClientConfig"},
			{kind: "HeaderManipulationPolicy"},
			{kind: "TransformationPolicy"},
			{kind: "ProxyProtocolPolicy"},
		}),
	}
}
