// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v2

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the networking.gloo.solo.io/v2 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the networking.gloo.solo.io/v2 APIs
type Clientset interface {
	// clienset for the networking.gloo.solo.io/v2/v2 APIs
	ExternalServices() ExternalServiceClient
	// clienset for the networking.gloo.solo.io/v2/v2 APIs
	ExternalEndpoints() ExternalEndpointClient
	// clienset for the networking.gloo.solo.io/v2/v2 APIs
	RouteTables() RouteTableClient
	// clienset for the networking.gloo.solo.io/v2/v2 APIs
	VirtualDestinations() VirtualDestinationClient
	// clienset for the networking.gloo.solo.io/v2/v2 APIs
	VirtualGateways() VirtualGatewayClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the networking.gloo.solo.io/v2/v2 APIs
func (c *clientSet) ExternalServices() ExternalServiceClient {
	return NewExternalServiceClient(c.client)
}

// clienset for the networking.gloo.solo.io/v2/v2 APIs
func (c *clientSet) ExternalEndpoints() ExternalEndpointClient {
	return NewExternalEndpointClient(c.client)
}

// clienset for the networking.gloo.solo.io/v2/v2 APIs
func (c *clientSet) RouteTables() RouteTableClient {
	return NewRouteTableClient(c.client)
}

// clienset for the networking.gloo.solo.io/v2/v2 APIs
func (c *clientSet) VirtualDestinations() VirtualDestinationClient {
	return NewVirtualDestinationClient(c.client)
}

// clienset for the networking.gloo.solo.io/v2/v2 APIs
func (c *clientSet) VirtualGateways() VirtualGatewayClient {
	return NewVirtualGatewayClient(c.client)
}

// Reader knows how to read and list ExternalServices.
type ExternalServiceReader interface {
	// Get retrieves a ExternalService for the given object key
	GetExternalService(ctx context.Context, key client.ObjectKey) (*ExternalService, error)

	// List retrieves list of ExternalServices for a given namespace and list options.
	ListExternalService(ctx context.Context, opts ...client.ListOption) (*ExternalServiceList, error)
}

// ExternalServiceTransitionFunction instructs the ExternalServiceWriter how to transition between an existing
// ExternalService object and a desired on an Upsert
type ExternalServiceTransitionFunction func(existing, desired *ExternalService) error

// Writer knows how to create, delete, and update ExternalServices.
type ExternalServiceWriter interface {
	// Create saves the ExternalService object.
	CreateExternalService(ctx context.Context, obj *ExternalService, opts ...client.CreateOption) error

	// Delete deletes the ExternalService object.
	DeleteExternalService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ExternalService object.
	UpdateExternalService(ctx context.Context, obj *ExternalService, opts ...client.UpdateOption) error

	// Patch patches the given ExternalService object.
	PatchExternalService(ctx context.Context, obj *ExternalService, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ExternalService objects matching the given options.
	DeleteAllOfExternalService(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ExternalService object.
	UpsertExternalService(ctx context.Context, obj *ExternalService, transitionFuncs ...ExternalServiceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ExternalService object.
type ExternalServiceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ExternalService object.
	UpdateExternalServiceStatus(ctx context.Context, obj *ExternalService, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given ExternalService object's subresource.
	PatchExternalServiceStatus(ctx context.Context, obj *ExternalService, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on ExternalServices.
type ExternalServiceClient interface {
	ExternalServiceReader
	ExternalServiceWriter
	ExternalServiceStatusWriter
}

type externalServiceClient struct {
	client client.Client
}

func NewExternalServiceClient(client client.Client) *externalServiceClient {
	return &externalServiceClient{client: client}
}

func (c *externalServiceClient) GetExternalService(ctx context.Context, key client.ObjectKey) (*ExternalService, error) {
	obj := &ExternalService{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *externalServiceClient) ListExternalService(ctx context.Context, opts ...client.ListOption) (*ExternalServiceList, error) {
	list := &ExternalServiceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *externalServiceClient) CreateExternalService(ctx context.Context, obj *ExternalService, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *externalServiceClient) DeleteExternalService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &ExternalService{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *externalServiceClient) UpdateExternalService(ctx context.Context, obj *ExternalService, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *externalServiceClient) PatchExternalService(ctx context.Context, obj *ExternalService, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *externalServiceClient) DeleteAllOfExternalService(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ExternalService{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *externalServiceClient) UpsertExternalService(ctx context.Context, obj *ExternalService, transitionFuncs ...ExternalServiceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*ExternalService), desired.(*ExternalService)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *externalServiceClient) UpdateExternalServiceStatus(ctx context.Context, obj *ExternalService, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *externalServiceClient) PatchExternalServiceStatus(ctx context.Context, obj *ExternalService, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides ExternalServiceClients for multiple clusters.
type MulticlusterExternalServiceClient interface {
	// Cluster returns a ExternalServiceClient for the given cluster
	Cluster(cluster string) (ExternalServiceClient, error)
}

type multiclusterExternalServiceClient struct {
	client multicluster.Client
}

func NewMulticlusterExternalServiceClient(client multicluster.Client) MulticlusterExternalServiceClient {
	return &multiclusterExternalServiceClient{client: client}
}

func (m *multiclusterExternalServiceClient) Cluster(cluster string) (ExternalServiceClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewExternalServiceClient(client), nil
}

// Reader knows how to read and list ExternalEndpoints.
type ExternalEndpointReader interface {
	// Get retrieves a ExternalEndpoint for the given object key
	GetExternalEndpoint(ctx context.Context, key client.ObjectKey) (*ExternalEndpoint, error)

	// List retrieves list of ExternalEndpoints for a given namespace and list options.
	ListExternalEndpoint(ctx context.Context, opts ...client.ListOption) (*ExternalEndpointList, error)
}

// ExternalEndpointTransitionFunction instructs the ExternalEndpointWriter how to transition between an existing
// ExternalEndpoint object and a desired on an Upsert
type ExternalEndpointTransitionFunction func(existing, desired *ExternalEndpoint) error

// Writer knows how to create, delete, and update ExternalEndpoints.
type ExternalEndpointWriter interface {
	// Create saves the ExternalEndpoint object.
	CreateExternalEndpoint(ctx context.Context, obj *ExternalEndpoint, opts ...client.CreateOption) error

	// Delete deletes the ExternalEndpoint object.
	DeleteExternalEndpoint(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ExternalEndpoint object.
	UpdateExternalEndpoint(ctx context.Context, obj *ExternalEndpoint, opts ...client.UpdateOption) error

	// Patch patches the given ExternalEndpoint object.
	PatchExternalEndpoint(ctx context.Context, obj *ExternalEndpoint, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ExternalEndpoint objects matching the given options.
	DeleteAllOfExternalEndpoint(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ExternalEndpoint object.
	UpsertExternalEndpoint(ctx context.Context, obj *ExternalEndpoint, transitionFuncs ...ExternalEndpointTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ExternalEndpoint object.
type ExternalEndpointStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ExternalEndpoint object.
	UpdateExternalEndpointStatus(ctx context.Context, obj *ExternalEndpoint, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given ExternalEndpoint object's subresource.
	PatchExternalEndpointStatus(ctx context.Context, obj *ExternalEndpoint, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on ExternalEndpoints.
type ExternalEndpointClient interface {
	ExternalEndpointReader
	ExternalEndpointWriter
	ExternalEndpointStatusWriter
}

type externalEndpointClient struct {
	client client.Client
}

func NewExternalEndpointClient(client client.Client) *externalEndpointClient {
	return &externalEndpointClient{client: client}
}

func (c *externalEndpointClient) GetExternalEndpoint(ctx context.Context, key client.ObjectKey) (*ExternalEndpoint, error) {
	obj := &ExternalEndpoint{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *externalEndpointClient) ListExternalEndpoint(ctx context.Context, opts ...client.ListOption) (*ExternalEndpointList, error) {
	list := &ExternalEndpointList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *externalEndpointClient) CreateExternalEndpoint(ctx context.Context, obj *ExternalEndpoint, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *externalEndpointClient) DeleteExternalEndpoint(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &ExternalEndpoint{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *externalEndpointClient) UpdateExternalEndpoint(ctx context.Context, obj *ExternalEndpoint, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *externalEndpointClient) PatchExternalEndpoint(ctx context.Context, obj *ExternalEndpoint, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *externalEndpointClient) DeleteAllOfExternalEndpoint(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ExternalEndpoint{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *externalEndpointClient) UpsertExternalEndpoint(ctx context.Context, obj *ExternalEndpoint, transitionFuncs ...ExternalEndpointTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*ExternalEndpoint), desired.(*ExternalEndpoint)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *externalEndpointClient) UpdateExternalEndpointStatus(ctx context.Context, obj *ExternalEndpoint, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *externalEndpointClient) PatchExternalEndpointStatus(ctx context.Context, obj *ExternalEndpoint, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides ExternalEndpointClients for multiple clusters.
type MulticlusterExternalEndpointClient interface {
	// Cluster returns a ExternalEndpointClient for the given cluster
	Cluster(cluster string) (ExternalEndpointClient, error)
}

type multiclusterExternalEndpointClient struct {
	client multicluster.Client
}

func NewMulticlusterExternalEndpointClient(client multicluster.Client) MulticlusterExternalEndpointClient {
	return &multiclusterExternalEndpointClient{client: client}
}

func (m *multiclusterExternalEndpointClient) Cluster(cluster string) (ExternalEndpointClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewExternalEndpointClient(client), nil
}

// Reader knows how to read and list RouteTables.
type RouteTableReader interface {
	// Get retrieves a RouteTable for the given object key
	GetRouteTable(ctx context.Context, key client.ObjectKey) (*RouteTable, error)

	// List retrieves list of RouteTables for a given namespace and list options.
	ListRouteTable(ctx context.Context, opts ...client.ListOption) (*RouteTableList, error)
}

// RouteTableTransitionFunction instructs the RouteTableWriter how to transition between an existing
// RouteTable object and a desired on an Upsert
type RouteTableTransitionFunction func(existing, desired *RouteTable) error

// Writer knows how to create, delete, and update RouteTables.
type RouteTableWriter interface {
	// Create saves the RouteTable object.
	CreateRouteTable(ctx context.Context, obj *RouteTable, opts ...client.CreateOption) error

	// Delete deletes the RouteTable object.
	DeleteRouteTable(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given RouteTable object.
	UpdateRouteTable(ctx context.Context, obj *RouteTable, opts ...client.UpdateOption) error

	// Patch patches the given RouteTable object.
	PatchRouteTable(ctx context.Context, obj *RouteTable, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all RouteTable objects matching the given options.
	DeleteAllOfRouteTable(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the RouteTable object.
	UpsertRouteTable(ctx context.Context, obj *RouteTable, transitionFuncs ...RouteTableTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a RouteTable object.
type RouteTableStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given RouteTable object.
	UpdateRouteTableStatus(ctx context.Context, obj *RouteTable, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given RouteTable object's subresource.
	PatchRouteTableStatus(ctx context.Context, obj *RouteTable, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on RouteTables.
type RouteTableClient interface {
	RouteTableReader
	RouteTableWriter
	RouteTableStatusWriter
}

type routeTableClient struct {
	client client.Client
}

func NewRouteTableClient(client client.Client) *routeTableClient {
	return &routeTableClient{client: client}
}

func (c *routeTableClient) GetRouteTable(ctx context.Context, key client.ObjectKey) (*RouteTable, error) {
	obj := &RouteTable{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *routeTableClient) ListRouteTable(ctx context.Context, opts ...client.ListOption) (*RouteTableList, error) {
	list := &RouteTableList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *routeTableClient) CreateRouteTable(ctx context.Context, obj *RouteTable, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *routeTableClient) DeleteRouteTable(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &RouteTable{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *routeTableClient) UpdateRouteTable(ctx context.Context, obj *RouteTable, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *routeTableClient) PatchRouteTable(ctx context.Context, obj *RouteTable, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *routeTableClient) DeleteAllOfRouteTable(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &RouteTable{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *routeTableClient) UpsertRouteTable(ctx context.Context, obj *RouteTable, transitionFuncs ...RouteTableTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*RouteTable), desired.(*RouteTable)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *routeTableClient) UpdateRouteTableStatus(ctx context.Context, obj *RouteTable, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *routeTableClient) PatchRouteTableStatus(ctx context.Context, obj *RouteTable, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides RouteTableClients for multiple clusters.
type MulticlusterRouteTableClient interface {
	// Cluster returns a RouteTableClient for the given cluster
	Cluster(cluster string) (RouteTableClient, error)
}

type multiclusterRouteTableClient struct {
	client multicluster.Client
}

func NewMulticlusterRouteTableClient(client multicluster.Client) MulticlusterRouteTableClient {
	return &multiclusterRouteTableClient{client: client}
}

func (m *multiclusterRouteTableClient) Cluster(cluster string) (RouteTableClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewRouteTableClient(client), nil
}

// Reader knows how to read and list VirtualDestinations.
type VirtualDestinationReader interface {
	// Get retrieves a VirtualDestination for the given object key
	GetVirtualDestination(ctx context.Context, key client.ObjectKey) (*VirtualDestination, error)

	// List retrieves list of VirtualDestinations for a given namespace and list options.
	ListVirtualDestination(ctx context.Context, opts ...client.ListOption) (*VirtualDestinationList, error)
}

// VirtualDestinationTransitionFunction instructs the VirtualDestinationWriter how to transition between an existing
// VirtualDestination object and a desired on an Upsert
type VirtualDestinationTransitionFunction func(existing, desired *VirtualDestination) error

// Writer knows how to create, delete, and update VirtualDestinations.
type VirtualDestinationWriter interface {
	// Create saves the VirtualDestination object.
	CreateVirtualDestination(ctx context.Context, obj *VirtualDestination, opts ...client.CreateOption) error

	// Delete deletes the VirtualDestination object.
	DeleteVirtualDestination(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualDestination object.
	UpdateVirtualDestination(ctx context.Context, obj *VirtualDestination, opts ...client.UpdateOption) error

	// Patch patches the given VirtualDestination object.
	PatchVirtualDestination(ctx context.Context, obj *VirtualDestination, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualDestination objects matching the given options.
	DeleteAllOfVirtualDestination(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the VirtualDestination object.
	UpsertVirtualDestination(ctx context.Context, obj *VirtualDestination, transitionFuncs ...VirtualDestinationTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a VirtualDestination object.
type VirtualDestinationStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualDestination object.
	UpdateVirtualDestinationStatus(ctx context.Context, obj *VirtualDestination, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given VirtualDestination object's subresource.
	PatchVirtualDestinationStatus(ctx context.Context, obj *VirtualDestination, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on VirtualDestinations.
type VirtualDestinationClient interface {
	VirtualDestinationReader
	VirtualDestinationWriter
	VirtualDestinationStatusWriter
}

type virtualDestinationClient struct {
	client client.Client
}

func NewVirtualDestinationClient(client client.Client) *virtualDestinationClient {
	return &virtualDestinationClient{client: client}
}

func (c *virtualDestinationClient) GetVirtualDestination(ctx context.Context, key client.ObjectKey) (*VirtualDestination, error) {
	obj := &VirtualDestination{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualDestinationClient) ListVirtualDestination(ctx context.Context, opts ...client.ListOption) (*VirtualDestinationList, error) {
	list := &VirtualDestinationList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualDestinationClient) CreateVirtualDestination(ctx context.Context, obj *VirtualDestination, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualDestinationClient) DeleteVirtualDestination(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &VirtualDestination{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualDestinationClient) UpdateVirtualDestination(ctx context.Context, obj *VirtualDestination, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualDestinationClient) PatchVirtualDestination(ctx context.Context, obj *VirtualDestination, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualDestinationClient) DeleteAllOfVirtualDestination(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &VirtualDestination{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualDestinationClient) UpsertVirtualDestination(ctx context.Context, obj *VirtualDestination, transitionFuncs ...VirtualDestinationTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*VirtualDestination), desired.(*VirtualDestination)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *virtualDestinationClient) UpdateVirtualDestinationStatus(ctx context.Context, obj *VirtualDestination, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualDestinationClient) PatchVirtualDestinationStatus(ctx context.Context, obj *VirtualDestination, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides VirtualDestinationClients for multiple clusters.
type MulticlusterVirtualDestinationClient interface {
	// Cluster returns a VirtualDestinationClient for the given cluster
	Cluster(cluster string) (VirtualDestinationClient, error)
}

type multiclusterVirtualDestinationClient struct {
	client multicluster.Client
}

func NewMulticlusterVirtualDestinationClient(client multicluster.Client) MulticlusterVirtualDestinationClient {
	return &multiclusterVirtualDestinationClient{client: client}
}

func (m *multiclusterVirtualDestinationClient) Cluster(cluster string) (VirtualDestinationClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewVirtualDestinationClient(client), nil
}

// Reader knows how to read and list VirtualGateways.
type VirtualGatewayReader interface {
	// Get retrieves a VirtualGateway for the given object key
	GetVirtualGateway(ctx context.Context, key client.ObjectKey) (*VirtualGateway, error)

	// List retrieves list of VirtualGateways for a given namespace and list options.
	ListVirtualGateway(ctx context.Context, opts ...client.ListOption) (*VirtualGatewayList, error)
}

// VirtualGatewayTransitionFunction instructs the VirtualGatewayWriter how to transition between an existing
// VirtualGateway object and a desired on an Upsert
type VirtualGatewayTransitionFunction func(existing, desired *VirtualGateway) error

// Writer knows how to create, delete, and update VirtualGateways.
type VirtualGatewayWriter interface {
	// Create saves the VirtualGateway object.
	CreateVirtualGateway(ctx context.Context, obj *VirtualGateway, opts ...client.CreateOption) error

	// Delete deletes the VirtualGateway object.
	DeleteVirtualGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualGateway object.
	UpdateVirtualGateway(ctx context.Context, obj *VirtualGateway, opts ...client.UpdateOption) error

	// Patch patches the given VirtualGateway object.
	PatchVirtualGateway(ctx context.Context, obj *VirtualGateway, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualGateway objects matching the given options.
	DeleteAllOfVirtualGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the VirtualGateway object.
	UpsertVirtualGateway(ctx context.Context, obj *VirtualGateway, transitionFuncs ...VirtualGatewayTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a VirtualGateway object.
type VirtualGatewayStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualGateway object.
	UpdateVirtualGatewayStatus(ctx context.Context, obj *VirtualGateway, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given VirtualGateway object's subresource.
	PatchVirtualGatewayStatus(ctx context.Context, obj *VirtualGateway, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on VirtualGateways.
type VirtualGatewayClient interface {
	VirtualGatewayReader
	VirtualGatewayWriter
	VirtualGatewayStatusWriter
}

type virtualGatewayClient struct {
	client client.Client
}

func NewVirtualGatewayClient(client client.Client) *virtualGatewayClient {
	return &virtualGatewayClient{client: client}
}

func (c *virtualGatewayClient) GetVirtualGateway(ctx context.Context, key client.ObjectKey) (*VirtualGateway, error) {
	obj := &VirtualGateway{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualGatewayClient) ListVirtualGateway(ctx context.Context, opts ...client.ListOption) (*VirtualGatewayList, error) {
	list := &VirtualGatewayList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualGatewayClient) CreateVirtualGateway(ctx context.Context, obj *VirtualGateway, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualGatewayClient) DeleteVirtualGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &VirtualGateway{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualGatewayClient) UpdateVirtualGateway(ctx context.Context, obj *VirtualGateway, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualGatewayClient) PatchVirtualGateway(ctx context.Context, obj *VirtualGateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualGatewayClient) DeleteAllOfVirtualGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &VirtualGateway{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualGatewayClient) UpsertVirtualGateway(ctx context.Context, obj *VirtualGateway, transitionFuncs ...VirtualGatewayTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*VirtualGateway), desired.(*VirtualGateway)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *virtualGatewayClient) UpdateVirtualGatewayStatus(ctx context.Context, obj *VirtualGateway, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualGatewayClient) PatchVirtualGatewayStatus(ctx context.Context, obj *VirtualGateway, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides VirtualGatewayClients for multiple clusters.
type MulticlusterVirtualGatewayClient interface {
	// Cluster returns a VirtualGatewayClient for the given cluster
	Cluster(cluster string) (VirtualGatewayClient, error)
}

type multiclusterVirtualGatewayClient struct {
	client multicluster.Client
}

func NewMulticlusterVirtualGatewayClient(client multicluster.Client) MulticlusterVirtualGatewayClient {
	return &multiclusterVirtualGatewayClient{client: client}
}

func (m *multiclusterVirtualGatewayClient) Cluster(cluster string) (VirtualGatewayClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewVirtualGatewayClient(client), nil
}
