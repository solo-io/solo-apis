package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	groups = append(groups, GlooMeshGroups()...)
}

func GlooMeshGroups() []model.Group {
	return []model.Group{
		makeGroup("settings", "v1", []resourceToGenerate{
			{kind: "Settings"},
			{kind: "Dashboard"},
		}),
		makeGroup("discovery", "v1", []resourceToGenerate{
			{kind: "Destination"},
			{kind: "Workload"},
			{kind: "Mesh"},
		}),
		makeGroup("networking", "v1", []resourceToGenerate{
			{kind: "TrafficPolicy"},
			{kind: "AccessPolicy"},
			{kind: "VirtualMesh"},
		}),
		makeGroup("networking.enterprise", "v1beta1", []resourceToGenerate{
			{kind: "WasmDeployment"},
			{kind: "RateLimitClientConfig"},
			{kind: "RateLimitServerConfig"},
			{kind: "VirtualDestination"},
			{kind: "VirtualGateway"},
			{kind: "VirtualHost"},
			{kind: "RouteTable"},
			{kind: "ServiceDependency"},
			{kind: "CertificateVerification"},
		}),
		makeGroup("observability.enterprise", "v1", []resourceToGenerate{
			{kind: "AccessLogRecord"},
		}),
		makeGroup("rbac.enterprise", "v1", []resourceToGenerate{
			{kind: "Role"},
			{kind: "RoleBinding"},
		}),
		makeGroup("admin.enterprise", "v1alpha1", []resourceToGenerate{
			{kind: "IstioInstallation"},
		}),
		makeGroup("certificates", "v1", []resourceToGenerate{
			{kind: "IssuedCertificate"},
			{kind: "CertificateRequest"},
			{kind: "PodBounceDirective"},
		}),
		makeGroup("xds.agent.enterprise", "v1beta1", []resourceToGenerate{
			{kind: "XdsConfig"},
		}),
	}
}
