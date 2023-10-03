// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v2alpha1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the admin.gloo.solo.io/v2alpha1 APIs
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

// clienset for the admin.gloo.solo.io/v2alpha1 APIs
type Clientset interface {
	// clienset for the admin.gloo.solo.io/v2alpha1/v2alpha1 APIs
	WaypointLifecycleManagers() WaypointLifecycleManagerClient
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

// clienset for the admin.gloo.solo.io/v2alpha1/v2alpha1 APIs
func (c *clientSet) WaypointLifecycleManagers() WaypointLifecycleManagerClient {
	return NewWaypointLifecycleManagerClient(c.client)
}

// Reader knows how to read and list WaypointLifecycleManagers.
type WaypointLifecycleManagerReader interface {
	// Get retrieves a WaypointLifecycleManager for the given object key
	GetWaypointLifecycleManager(ctx context.Context, key client.ObjectKey) (*WaypointLifecycleManager, error)

	// List retrieves list of WaypointLifecycleManagers for a given namespace and list options.
	ListWaypointLifecycleManager(ctx context.Context, opts ...client.ListOption) (*WaypointLifecycleManagerList, error)
}

// WaypointLifecycleManagerTransitionFunction instructs the WaypointLifecycleManagerWriter how to transition between an existing
// WaypointLifecycleManager object and a desired on an Upsert
type WaypointLifecycleManagerTransitionFunction func(existing, desired *WaypointLifecycleManager) error

// Writer knows how to create, delete, and update WaypointLifecycleManagers.
type WaypointLifecycleManagerWriter interface {
	// Create saves the WaypointLifecycleManager object.
	CreateWaypointLifecycleManager(ctx context.Context, obj *WaypointLifecycleManager, opts ...client.CreateOption) error

	// Delete deletes the WaypointLifecycleManager object.
	DeleteWaypointLifecycleManager(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given WaypointLifecycleManager object.
	UpdateWaypointLifecycleManager(ctx context.Context, obj *WaypointLifecycleManager, opts ...client.UpdateOption) error

	// Patch patches the given WaypointLifecycleManager object.
	PatchWaypointLifecycleManager(ctx context.Context, obj *WaypointLifecycleManager, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all WaypointLifecycleManager objects matching the given options.
	DeleteAllOfWaypointLifecycleManager(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the WaypointLifecycleManager object.
	UpsertWaypointLifecycleManager(ctx context.Context, obj *WaypointLifecycleManager, transitionFuncs ...WaypointLifecycleManagerTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a WaypointLifecycleManager object.
type WaypointLifecycleManagerStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given WaypointLifecycleManager object.
	UpdateWaypointLifecycleManagerStatus(ctx context.Context, obj *WaypointLifecycleManager, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given WaypointLifecycleManager object's subresource.
	PatchWaypointLifecycleManagerStatus(ctx context.Context, obj *WaypointLifecycleManager, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on WaypointLifecycleManagers.
type WaypointLifecycleManagerClient interface {
	WaypointLifecycleManagerReader
	WaypointLifecycleManagerWriter
	WaypointLifecycleManagerStatusWriter
}

type waypointLifecycleManagerClient struct {
	client client.Client
}

func NewWaypointLifecycleManagerClient(client client.Client) *waypointLifecycleManagerClient {
	return &waypointLifecycleManagerClient{client: client}
}

func (c *waypointLifecycleManagerClient) GetWaypointLifecycleManager(ctx context.Context, key client.ObjectKey) (*WaypointLifecycleManager, error) {
	obj := &WaypointLifecycleManager{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *waypointLifecycleManagerClient) ListWaypointLifecycleManager(ctx context.Context, opts ...client.ListOption) (*WaypointLifecycleManagerList, error) {
	list := &WaypointLifecycleManagerList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *waypointLifecycleManagerClient) CreateWaypointLifecycleManager(ctx context.Context, obj *WaypointLifecycleManager, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *waypointLifecycleManagerClient) DeleteWaypointLifecycleManager(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &WaypointLifecycleManager{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *waypointLifecycleManagerClient) UpdateWaypointLifecycleManager(ctx context.Context, obj *WaypointLifecycleManager, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *waypointLifecycleManagerClient) PatchWaypointLifecycleManager(ctx context.Context, obj *WaypointLifecycleManager, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *waypointLifecycleManagerClient) DeleteAllOfWaypointLifecycleManager(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &WaypointLifecycleManager{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *waypointLifecycleManagerClient) UpsertWaypointLifecycleManager(ctx context.Context, obj *WaypointLifecycleManager, transitionFuncs ...WaypointLifecycleManagerTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*WaypointLifecycleManager), desired.(*WaypointLifecycleManager)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *waypointLifecycleManagerClient) UpdateWaypointLifecycleManagerStatus(ctx context.Context, obj *WaypointLifecycleManager, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *waypointLifecycleManagerClient) PatchWaypointLifecycleManagerStatus(ctx context.Context, obj *WaypointLifecycleManager, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides WaypointLifecycleManagerClients for multiple clusters.
type MulticlusterWaypointLifecycleManagerClient interface {
	// Cluster returns a WaypointLifecycleManagerClient for the given cluster
	Cluster(cluster string) (WaypointLifecycleManagerClient, error)
}

type multiclusterWaypointLifecycleManagerClient struct {
	client multicluster.Client
}

func NewMulticlusterWaypointLifecycleManagerClient(client multicluster.Client) MulticlusterWaypointLifecycleManagerClient {
	return &multiclusterWaypointLifecycleManagerClient{client: client}
}

func (m *multiclusterWaypointLifecycleManagerClient) Cluster(cluster string) (WaypointLifecycleManagerClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewWaypointLifecycleManagerClient(client), nil
}
