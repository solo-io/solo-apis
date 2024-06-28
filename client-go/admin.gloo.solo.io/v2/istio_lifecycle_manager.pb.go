// Use Gloo Platform to install Istio control planes in your workload clusters,
// as part of the Istio lifecycle management.
// In your `IstioLifecycleManager` resource, you provide `istiod` settings in an `IstioOperator` configuration.
// When you create the `IstioLifecycleManager` in your management cluster, Gloo translates the configuration
// into `istiod` control planes in your registered workload clusters for you.
//
// For more information, see the [Install Istio by using the Istio Lifecycle Manager]({{% link path="/setup/install/gloo_mesh_managed/" %}}) guide.
//
// ## Example
// This example creates an `istiod` control plane in the `istio-system` namespace of two workload clusters
// (`$REMOTE_CLUSTER1` and `$REMOTE_CLUSTER2`). You supply the repo key for the Solo distribution of Istio (`hub: $REPO`), image tag (`tag: $ISTIO_IMAGE`),
// and revision (`revision: $REVISION`).
// ```yaml
// apiVersion: admin.gloo.solo.io/v2
// kind: IstioLifecycleManager
// metadata:
//   name: istiod-control-plane
//   namespace: gloo-mesh
// spec:
//   installations:
//   # The revision for this installation
//   - revision: $REVISION
//     # List all workload clusters to install Istio into
//     clusters:
//     - name: $REMOTE_CLUSTER1
//       # If set to true, the spec for this revision is applied in the cluster
//       defaultRevision: true
//     - name: $REMOTE_CLUSTER2
//       defaultRevision: true
//     istioOperatorSpec:
//       # Only the control plane components are installed
//       # (https://istio.io/latest/docs/setup/additional-setup/config-profiles/)
//       profile: minimal
//       # Solo.io Istio distribution repository; required for the Solo distribution of Istio.
//       # You get the repo key from your Solo Account Representative.
//       hub: $REPO
//       # Any tag for the Solo distribution of Istio
//       tag: $ISTIO_IMAGE
//       namespace: istio-system
//       # Mesh configuration
//       meshConfig:
//         # Enable access logging only if using.
//         accessLogFile: /dev/stdout
//         # Encoding for the proxy access log (TEXT or JSON). Default value is TEXT.
//         accessLogEncoding: JSON
//         # Enable span tracing only if using.
//         enableTracing: true
//         defaultConfig:
//           # Wait for the istio-proxy to start before starting application pods
//           holdApplicationUntilProxyStarts: true
//           proxyMetadata:
//             # For known hosts, enable the Istio agent to handle DNS requests for any custom ServiceEntry, such as non-Kubernetes services.
//             # Unknown hosts are automatically resolved using upstream DNS servers in resolv.conf (for proxy-dns)
//             ISTIO_META_DNS_CAPTURE: "true"
//         # Set the default behavior of the sidecar for handling outbound traffic
//         # from the application
//         outboundTrafficPolicy:
//           mode: ALLOW_ANY
//         # The administrative root namespace for Istio configuration
//         rootNamespace: istio-system
//         # Set to the cluster name by default.
//         trustDomain: ${CLUSTER_NAME}
//       # Traffic management
//       components:
//         pilot:
//           k8s:
//             env:
//             # Disable selecting workload entries for local service routing, so that Kubernetes
//             # will not automatically match services to workload entries with matching selector labels.
//             # Required for Gloo Mesh VirtualDestination functionality.
//             # For more info, see https://istio.io/latest/docs/reference/commands/pilot-discovery/
//             - name: PILOT_ENABLE_K8S_SELECT_WORKLOAD_ENTRIES
//               value: "false"
//             # Skip the validation step for mTLS within the cluster.
//             # This approach is not recommended if you integrated Istio with your own CA,
//             # but is useful for PoCs or demos in which you use self-signed certificates.
//             - name: PILOT_SKIP_VALIDATE_TRUST_DOMAIN
//               value: "true"
//       # Solo.io Istio distribution repository; required for Solo distributions of Istio.
//       # You get the repo key from your Solo Account Representative.
//       hub: $REPO
//       meshConfig:
//         # Enable access logging.
//         accessLogFile: /dev/stdout
//         defaultConfig:
//           # Wait for the istio-proxy to start before starting application pods
//           holdApplicationUntilProxyStarts: true
//           proxyMetadata:
//             # For known hosts, enable the Istio agent to handle DNS requests
//             # for any custom ServiceEntry, such as non-Kubernetes services.
//             # Unknown hosts are automatically resolved using upstream DNS
//             # servers in resolv.conf (for proxy-dns)
//             ISTIO_META_DNS_CAPTURE: "true"
//         # Set the default behavior of the sidecar for handling outbound
//         # traffic from the application
//         outboundTrafficPolicy:
//           mode: ALLOW_ANY
//       namespace: istio-system
//       # Only the control plane components are installed
//       # (https://istio.io/latest/docs/setup/additional-setup/config-profiles/)
//       profile: minimal
//       # The tag of a Solo distribution of Istio
//       tag: $ISTIO_IMAGE
//     # The revision for this installation, such as {{< reuse "conrefs/versions/istio_revision.md" >}}
//     revision: $REVISION
// ```

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/solo-apis/api/gloo.solo.io/admin/v2/istio_lifecycle_manager.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/cue/encoding/protobuf/cue"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "k8s.io/api/core/v1"

	v2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The current state of the Istio installation.
type IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State int32

const (
	// Waiting for resources to be installed or updated.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_PENDING IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 0
	// The Gloo management server encountered a problem while attempting
	// to install Istio.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_FAILED IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 1
	// The controller is currently being installed.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_INSTALLING_CONTROLLER IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 2
	// The controller failed to install.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_CONTROLLER_INSTALL_FAILED IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 3
	// The Istio control plane is currently being installed.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_INSTALLING_CONTROL_PLANE IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 4
	// The Istio control plane failed to install.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_CONTROL_PLANE_INSTALL_FAILED IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 5
	// All Istio components are successfully installed and healthy.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_HEALTHY IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 6
	// The Istio installation is no longer healthy.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_UNHEALTHY IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 7
	// The control plane IstioOperator resource is in an 'ACTION_REQUIRED' state. Check the logs of the IstioOperator deployment for more info.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_ACTION_REQUIRED IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 8
	// The control plane IstioOperator resource is in an 'UPDATING' state.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_UPDATING_CONTROL_PLANE IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 9
	// The control plane IstioOperator resource is in a 'RECONCILING' state.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_RECONCILING_CONTROL_PLANE IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 10
	// The control plane installation state could not be determined.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_UNKNOWN IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 11
	// The Istio control plane is currently being uninstalled.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_UNINSTALLING_CONTROL_PLANE IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 12
	// The Istio control plane is uninstalled.
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_UNINSTALLED_CONTROL_PLANE IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State = 13
)

// Enum value maps for IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State.
var (
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State_name = map[int32]string{
		0:  "PENDING",
		1:  "FAILED",
		2:  "INSTALLING_CONTROLLER",
		3:  "CONTROLLER_INSTALL_FAILED",
		4:  "INSTALLING_CONTROL_PLANE",
		5:  "CONTROL_PLANE_INSTALL_FAILED",
		6:  "HEALTHY",
		7:  "UNHEALTHY",
		8:  "ACTION_REQUIRED",
		9:  "UPDATING_CONTROL_PLANE",
		10: "RECONCILING_CONTROL_PLANE",
		11: "UNKNOWN",
		12: "UNINSTALLING_CONTROL_PLANE",
		13: "UNINSTALLED_CONTROL_PLANE",
	}
	IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State_value = map[string]int32{
		"PENDING":                      0,
		"FAILED":                       1,
		"INSTALLING_CONTROLLER":        2,
		"CONTROLLER_INSTALL_FAILED":    3,
		"INSTALLING_CONTROL_PLANE":     4,
		"CONTROL_PLANE_INSTALL_FAILED": 5,
		"HEALTHY":                      6,
		"UNHEALTHY":                    7,
		"ACTION_REQUIRED":              8,
		"UPDATING_CONTROL_PLANE":       9,
		"RECONCILING_CONTROL_PLANE":    10,
		"UNKNOWN":                      11,
		"UNINSTALLING_CONTROL_PLANE":   12,
		"UNINSTALLED_CONTROL_PLANE":    13,
	}
)

func (x IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) Enum() *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State {
	p := new(IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State)
	*p = x
	return p
}

func (x IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_enumTypes[0].Descriptor()
}

func (IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_enumTypes[0]
}

func (x IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State.Descriptor instead.
func (IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{4, 1, 1, 0}
}

// Specifications for the `IstioLifecycleManager` resource.
type IstioLifecycleManagerSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of Istio control plane installations.
	Installations []*IstioInstallation `protobuf:"bytes,1,rep,name=installations,proto3" json:"installations,omitempty"`
}

func (x *IstioLifecycleManagerSpec) Reset() {
	*x = IstioLifecycleManagerSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioLifecycleManagerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioLifecycleManagerSpec) ProtoMessage() {}

func (x *IstioLifecycleManagerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioLifecycleManagerSpec.ProtoReflect.Descriptor instead.
func (*IstioLifecycleManagerSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{0}
}

func (x *IstioLifecycleManagerSpec) GetInstallations() []*IstioInstallation {
	if x != nil {
		return x.Installations
	}
	return nil
}

// Clusters to install the Istio control planes in.
type IstioClusterSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the cluster to install Istio into.
	// Must match the name of the cluster that you used when you registered the cluster with Gloo.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional: Defaults to false.
	// When set to true, the installation for this revision is applied as the active Istio installation in the cluster.
	// Resources with the `istio-injection=true` label entry use this revision. You might change this setting for Istio installations
	// during a canary upgrade. For more info, see the [upgrade docs](https://docs.solo.io/gloo-mesh-enterprise/latest/istio/mesh/ilm-upgrade/).
	DefaultRevision bool `protobuf:"varint,2,opt,name=default_revision,json=defaultRevision,proto3" json:"default_revision,omitempty"`
	// Optional: By default, the `trustDomain` value in the `meshConfig` section of the operator spec is automatically set by the Gloo to the name of each workload cluster.
	// To override the `trustDomain` for each cluster, you can instead specify the override value by using this `trustDomain` field,
	// and include the value in the list of cluster names. For example, if you specify `meshConfig.trustDomain: cluster1-trust-override` in your operator spec,
	// you then specify both the cluster name (`name: cluster1`) and the trust domain (`trustDomain: cluster1-trust-override`) in this `installations.clusters` section.
	// Additionally, because Gloo requires multiple trust domains for east-west routing, the `PILOT_SKIP_VALIDATE_TRUST_DOMAIN` field is set to `"true"` by default.
	// For more info, see the [Istio documentation](https://istio.io/latest/docs/reference/config/istio.mesh.v1alpha1).
	TrustDomain string `protobuf:"bytes,5,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
}

func (x *IstioClusterSelector) Reset() {
	*x = IstioClusterSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioClusterSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioClusterSelector) ProtoMessage() {}

func (x *IstioClusterSelector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioClusterSelector.ProtoReflect.Descriptor instead.
func (*IstioClusterSelector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{1}
}

func (x *IstioClusterSelector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IstioClusterSelector) GetDefaultRevision() bool {
	if x != nil {
		return x.DefaultRevision
	}
	return false
}

func (x *IstioClusterSelector) GetTrustDomain() string {
	if x != nil {
		return x.TrustDomain
	}
	return ""
}

type IstioController struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of time to wait for resources in a component to become ready before giving up. Configured using a duration string.
	// A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix,
	// such as "300ms" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	WaitForResourcesTimeout string `protobuf:"bytes,1,opt,name=wait_for_resources_timeout,json=waitForResourcesTimeout,proto3" json:"wait_for_resources_timeout,omitempty"`
	// Names of image pull secrets to use to deploy the Istio controller.
	// For more info, see the [Kubernetes docs](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/#create-a-pod-that-uses-your-secret).
	ImagePullSecrets []*v1.LocalObjectReference `protobuf:"bytes,2,rep,name=image_pull_secrets,json=imagePullSecrets,proto3" json:"image_pull_secrets,omitempty"`
	// Override for resources allocated to the Istio controller deployment.
	// For more info, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#example-1).
	Resources *v2.Resources `protobuf:"bytes,3,opt,name=resources,proto3" json:"resources,omitempty"`
	// Override for the pod's security context. For more info, see the
	// [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#securitycontext-v1-core).
	SecurityContext *v1.SecurityContext `protobuf:"bytes,4,opt,name=security_context,json=securityContext,proto3" json:"security_context,omitempty"`
	// Kubernetes pod/deployment/service labels.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Kubernetes pod/deployment/service annotations.
	Annotations map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Deployment environment variables. For more info, see the
	// [Kubernetes docs](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/).
	EnvVars []*v1.EnvVar `protobuf:"bytes,7,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty"`
}

func (x *IstioController) Reset() {
	*x = IstioController{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioController) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioController) ProtoMessage() {}

func (x *IstioController) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioController.ProtoReflect.Descriptor instead.
func (*IstioController) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{2}
}

func (x *IstioController) GetWaitForResourcesTimeout() string {
	if x != nil {
		return x.WaitForResourcesTimeout
	}
	return ""
}

func (x *IstioController) GetImagePullSecrets() []*v1.LocalObjectReference {
	if x != nil {
		return x.ImagePullSecrets
	}
	return nil
}

func (x *IstioController) GetResources() *v2.Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *IstioController) GetSecurityContext() *v1.SecurityContext {
	if x != nil {
		return x.SecurityContext
	}
	return nil
}

func (x *IstioController) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *IstioController) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *IstioController) GetEnvVars() []*v1.EnvVar {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

// List of Istio control plane installations.
// Any components that are NOT related to the control plane are ignored.
type IstioInstallation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Istio revision for this installation.
	// Label workload resources with 'istio.io/rev=$REVISION' to use this installation.
	// When set to `auto`, Gloo installs the control plane with the default supported version of the Solo distribution of Istio.
	Revision string `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// Clusters to install the Istio control planes in.
	Clusters []*IstioClusterSelector `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// IstioOperator specification for the control plane.
	// For more info, see the [Istio documentation](https://istio.io/latest/docs/reference/config/istio.operator.v1alpha1/).
	IstioOperatorSpec *v2.IstioOperatorSpec `protobuf:"bytes,3,opt,name=istio_operator_spec,json=istioOperatorSpec,proto3" json:"istio_operator_spec,omitempty"`
	// Optional configuration to tune the deployment of the IstioOperator controller deployed to each workload cluster.
	IstioController *IstioController `protobuf:"bytes,4,opt,name=istio_controller,json=istioController,proto3" json:"istio_controller,omitempty"`
	// When set to true, the lifecycle manager allows you to perform
	// in-place upgrades by skipping checks that are required for canary upgrades.
	// In production environments, canary upgrades are recommended for
	// updating the minor version. To update the patch version or make
	// configuration changes within the same version, you can use in-place upgrades.
	// Be sure to test in-place upgrades in development or staging environments first.
	SkipUpgradeValidation bool `protobuf:"varint,5,opt,name=skip_upgrade_validation,json=skipUpgradeValidation,proto3" json:"skip_upgrade_validation,omitempty"`
}

func (x *IstioInstallation) Reset() {
	*x = IstioInstallation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioInstallation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioInstallation) ProtoMessage() {}

func (x *IstioInstallation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioInstallation.ProtoReflect.Descriptor instead.
func (*IstioInstallation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{3}
}

func (x *IstioInstallation) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *IstioInstallation) GetClusters() []*IstioClusterSelector {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *IstioInstallation) GetIstioOperatorSpec() *v2.IstioOperatorSpec {
	if x != nil {
		return x.IstioOperatorSpec
	}
	return nil
}

func (x *IstioInstallation) GetIstioController() *IstioController {
	if x != nil {
		return x.IstioController
	}
	return nil
}

func (x *IstioInstallation) GetSkipUpgradeValidation() bool {
	if x != nil {
		return x.SkipUpgradeValidation
	}
	return false
}

// The status of the `IstioLifecycleManager` resource after you apply it to your Gloo environment.
type IstioLifecycleManagerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of clusters where Gloo manages Istio installations.
	Clusters map[string]*IstioLifecycleManagerStatus_ClusterStatuses `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IstioLifecycleManagerStatus) Reset() {
	*x = IstioLifecycleManagerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioLifecycleManagerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioLifecycleManagerStatus) ProtoMessage() {}

func (x *IstioLifecycleManagerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioLifecycleManagerStatus.ProtoReflect.Descriptor instead.
func (*IstioLifecycleManagerStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{4}
}

func (x *IstioLifecycleManagerStatus) GetClusters() map[string]*IstioLifecycleManagerStatus_ClusterStatuses {
	if x != nil {
		return x.Clusters
	}
	return nil
}

// $hide_from_docs
type IstioLifecycleManagerNewStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IstioLifecycleManagerNewStatus) Reset() {
	*x = IstioLifecycleManagerNewStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioLifecycleManagerNewStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioLifecycleManagerNewStatus) ProtoMessage() {}

func (x *IstioLifecycleManagerNewStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioLifecycleManagerNewStatus.ProtoReflect.Descriptor instead.
func (*IstioLifecycleManagerNewStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{5}
}

// $hide_from_docs
type IstioLifecycleManagerReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IstioLifecycleManagerReport) Reset() {
	*x = IstioLifecycleManagerReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioLifecycleManagerReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioLifecycleManagerReport) ProtoMessage() {}

func (x *IstioLifecycleManagerReport) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioLifecycleManagerReport.ProtoReflect.Descriptor instead.
func (*IstioLifecycleManagerReport) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{6}
}

// The list of clusters where Gloo manages Istio installations.
type IstioLifecycleManagerStatus_ClusterStatuses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Istio installations in the cluster, listed by revision.
	Installations map[string]*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus `protobuf:"bytes,1,rep,name=installations,proto3" json:"installations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IstioLifecycleManagerStatus_ClusterStatuses) Reset() {
	*x = IstioLifecycleManagerStatus_ClusterStatuses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioLifecycleManagerStatus_ClusterStatuses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioLifecycleManagerStatus_ClusterStatuses) ProtoMessage() {}

func (x *IstioLifecycleManagerStatus_ClusterStatuses) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioLifecycleManagerStatus_ClusterStatuses.ProtoReflect.Descriptor instead.
func (*IstioLifecycleManagerStatus_ClusterStatuses) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{4, 1}
}

func (x *IstioLifecycleManagerStatus_ClusterStatuses) GetInstallations() map[string]*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus {
	if x != nil {
		return x.Installations
	}
	return nil
}

// The status of the installation.
type IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current state of the Istio installation.
	State IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=admin.gloo.solo.io.IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State" json:"state,omitempty"`
	// A human readable message about the current state of the installation.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The observed revision of the Istio installation.
	ObservedRevision string `protobuf:"bytes,5,opt,name=observed_revision,json=observedRevision,proto3" json:"observed_revision,omitempty"`
	// The IstioOperator spec that is currently deployed for this revision.
	ObservedOperator *v2.IstioOperatorSpec `protobuf:"bytes,4,opt,name=observed_operator,json=observedOperator,proto3" json:"observed_operator,omitempty"`
}

func (x *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) Reset() {
	*x = IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) ProtoMessage() {}

func (x *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus.ProtoReflect.Descriptor instead.
func (*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP(), []int{4, 1, 1}
}

func (x *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) GetState() IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State {
	if x != nil {
		return x.State
	}
	return IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_PENDING
}

func (x *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) GetObservedRevision() string {
	if x != nil {
		return x.ObservedRevision
	}
	return ""
}

func (x *IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus) GetObservedOperator() *v2.IstioOperatorSpec {
	if x != nil {
		return x.ObservedOperator
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDesc = []byte{
	0x0a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x75, 0x65, 0x2f, 0x63, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6b, 0x38, 0x73, 0x2e,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x57,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x32, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x19, 0x49, 0x73, 0x74, 0x69, 0x6f,
	0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x4b, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x78, 0x0a, 0x14, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x87, 0x05, 0x0a, 0x0f,
	0x49, 0x73, 0x74, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x3b, 0x0a, 0x1a, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x77, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x56, 0x0a, 0x12,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x4e, 0x0a, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x47, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x56, 0x0a, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61,
	0x72, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd5, 0x02, 0x0a, 0x11, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49,
	0x73, 0x74, 0x69, 0x6f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x56, 0x0a,
	0x13, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x11, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4e, 0x0a, 0x10, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x75, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x73, 0x6b, 0x69, 0x70, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x93, 0x09,
	0x0a, 0x1b, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x59, 0x0a,
	0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x7c, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x55, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x49, 0x73, 0x74, 0x69, 0x6f, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x9a, 0x07, 0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x78, 0x0a, 0x0d, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x52, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x94, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x68, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xf5, 0x04, 0x0a, 0x12,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x6e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x58, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x11, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x10, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xd2,
	0x02, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19,
	0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41,
	0x4c, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x49,
	0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f,
	0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x45, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e,
	0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41,
	0x4c, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x48,
	0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x48, 0x45,
	0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x08, 0x12, 0x1a, 0x0a, 0x16,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c,
	0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x45, 0x10, 0x09, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x43, 0x4f,
	0x4e, 0x43, 0x49, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f,
	0x50, 0x4c, 0x41, 0x4e, 0x45, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x0b, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x4e, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x50, 0x4c, 0x41,
	0x4e, 0x45, 0x10, 0x0c, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x4e, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x4e,
	0x45, 0x10, 0x0d, 0x22, 0x20, 0x0a, 0x1e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1d, 0x0a, 0x1b, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0x53, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d,
	0x6d, 0x65, 0x73, 0x68, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x32, 0xc0, 0xf5, 0x04,
	0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescData = file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_goTypes = []interface{}{
	(IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus_State)(0), // 0: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.InstallationStatus.State
	(*IstioLifecycleManagerSpec)(nil),                                         // 1: admin.gloo.solo.io.IstioLifecycleManagerSpec
	(*IstioClusterSelector)(nil),                                              // 2: admin.gloo.solo.io.IstioClusterSelector
	(*IstioController)(nil),                                                   // 3: admin.gloo.solo.io.IstioController
	(*IstioInstallation)(nil),                                                 // 4: admin.gloo.solo.io.IstioInstallation
	(*IstioLifecycleManagerStatus)(nil),                                       // 5: admin.gloo.solo.io.IstioLifecycleManagerStatus
	(*IstioLifecycleManagerNewStatus)(nil),                                    // 6: admin.gloo.solo.io.IstioLifecycleManagerNewStatus
	(*IstioLifecycleManagerReport)(nil),                                       // 7: admin.gloo.solo.io.IstioLifecycleManagerReport
	nil,                                                                       // 8: admin.gloo.solo.io.IstioController.LabelsEntry
	nil,                                                                       // 9: admin.gloo.solo.io.IstioController.AnnotationsEntry
	nil,                                                                       // 10: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClustersEntry
	(*IstioLifecycleManagerStatus_ClusterStatuses)(nil), // 11: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses
	nil, // 12: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.InstallationsEntry
	(*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus)(nil), // 13: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.InstallationStatus
	(*v1.LocalObjectReference)(nil),                                        // 14: k8s.io.api.core.v1.LocalObjectReference
	(*v2.Resources)(nil),                                                   // 15: common.gloo.solo.io.Resources
	(*v1.SecurityContext)(nil),                                             // 16: k8s.io.api.core.v1.SecurityContext
	(*v1.EnvVar)(nil),                                                      // 17: k8s.io.api.core.v1.EnvVar
	(*v2.IstioOperatorSpec)(nil),                                           // 18: common.gloo.solo.io.IstioOperatorSpec
}
var file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_depIdxs = []int32{
	4,  // 0: admin.gloo.solo.io.IstioLifecycleManagerSpec.installations:type_name -> admin.gloo.solo.io.IstioInstallation
	14, // 1: admin.gloo.solo.io.IstioController.image_pull_secrets:type_name -> k8s.io.api.core.v1.LocalObjectReference
	15, // 2: admin.gloo.solo.io.IstioController.resources:type_name -> common.gloo.solo.io.Resources
	16, // 3: admin.gloo.solo.io.IstioController.security_context:type_name -> k8s.io.api.core.v1.SecurityContext
	8,  // 4: admin.gloo.solo.io.IstioController.labels:type_name -> admin.gloo.solo.io.IstioController.LabelsEntry
	9,  // 5: admin.gloo.solo.io.IstioController.annotations:type_name -> admin.gloo.solo.io.IstioController.AnnotationsEntry
	17, // 6: admin.gloo.solo.io.IstioController.env_vars:type_name -> k8s.io.api.core.v1.EnvVar
	2,  // 7: admin.gloo.solo.io.IstioInstallation.clusters:type_name -> admin.gloo.solo.io.IstioClusterSelector
	18, // 8: admin.gloo.solo.io.IstioInstallation.istio_operator_spec:type_name -> common.gloo.solo.io.IstioOperatorSpec
	3,  // 9: admin.gloo.solo.io.IstioInstallation.istio_controller:type_name -> admin.gloo.solo.io.IstioController
	10, // 10: admin.gloo.solo.io.IstioLifecycleManagerStatus.clusters:type_name -> admin.gloo.solo.io.IstioLifecycleManagerStatus.ClustersEntry
	11, // 11: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClustersEntry.value:type_name -> admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses
	12, // 12: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.installations:type_name -> admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.InstallationsEntry
	13, // 13: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.InstallationsEntry.value:type_name -> admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.InstallationStatus
	0,  // 14: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.InstallationStatus.state:type_name -> admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.InstallationStatus.State
	18, // 15: admin.gloo.solo.io.IstioLifecycleManagerStatus.ClusterStatuses.InstallationStatus.observed_operator:type_name -> common.gloo.solo.io.IstioOperatorSpec
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_init()
}
func file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_init() {
	if File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioLifecycleManagerSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioClusterSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioController); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioInstallation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioLifecycleManagerStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioLifecycleManagerNewStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioLifecycleManagerReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioLifecycleManagerStatus_ClusterStatuses); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioLifecycleManagerStatus_ClusterStatuses_InstallationStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto = out.File
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_solo_apis_api_gloo_solo_io_admin_v2_istio_lifecycle_manager_proto_depIdxs = nil
}
