// Code generated by skv2. DO NOT EDIT.

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the fed.solo.io/v1 APIs
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

// clienset for the fed.solo.io/v1 APIs
type Clientset interface {
	// clienset for the fed.solo.io/v1/v1 APIs
	GlooInstances() GlooInstanceClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
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

// clienset for the fed.solo.io/v1/v1 APIs
func (c *clientSet) GlooInstances() GlooInstanceClient {
	return NewGlooInstanceClient(c.client)
}

// Reader knows how to read and list GlooInstances.
type GlooInstanceReader interface {
	// Get retrieves a GlooInstance for the given object key
	GetGlooInstance(ctx context.Context, key client.ObjectKey) (*GlooInstance, error)

	// List retrieves list of GlooInstances for a given namespace and list options.
	ListGlooInstance(ctx context.Context, opts ...client.ListOption) (*GlooInstanceList, error)
}

// GlooInstanceTransitionFunction instructs the GlooInstanceWriter how to transition between an existing
// GlooInstance object and a desired on an Upsert
type GlooInstanceTransitionFunction func(existing, desired *GlooInstance) error

// Writer knows how to create, delete, and update GlooInstances.
type GlooInstanceWriter interface {
	// Create saves the GlooInstance object.
	CreateGlooInstance(ctx context.Context, obj *GlooInstance, opts ...client.CreateOption) error

	// Delete deletes the GlooInstance object.
	DeleteGlooInstance(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given GlooInstance object.
	UpdateGlooInstance(ctx context.Context, obj *GlooInstance, opts ...client.UpdateOption) error

	// Patch patches the given GlooInstance object.
	PatchGlooInstance(ctx context.Context, obj *GlooInstance, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all GlooInstance objects matching the given options.
	DeleteAllOfGlooInstance(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the GlooInstance object.
	UpsertGlooInstance(ctx context.Context, obj *GlooInstance, transitionFuncs ...GlooInstanceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a GlooInstance object.
type GlooInstanceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given GlooInstance object.
	UpdateGlooInstanceStatus(ctx context.Context, obj *GlooInstance, opts ...client.UpdateOption) error

	// Patch patches the given GlooInstance object's subresource.
	PatchGlooInstanceStatus(ctx context.Context, obj *GlooInstance, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on GlooInstances.
type GlooInstanceClient interface {
	GlooInstanceReader
	GlooInstanceWriter
	GlooInstanceStatusWriter
}

type glooInstanceClient struct {
	client client.Client
}

func NewGlooInstanceClient(client client.Client) *glooInstanceClient {
	return &glooInstanceClient{client: client}
}

func (c *glooInstanceClient) GetGlooInstance(ctx context.Context, key client.ObjectKey) (*GlooInstance, error) {
	obj := &GlooInstance{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *glooInstanceClient) ListGlooInstance(ctx context.Context, opts ...client.ListOption) (*GlooInstanceList, error) {
	list := &GlooInstanceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *glooInstanceClient) CreateGlooInstance(ctx context.Context, obj *GlooInstance, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *glooInstanceClient) DeleteGlooInstance(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &GlooInstance{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *glooInstanceClient) UpdateGlooInstance(ctx context.Context, obj *GlooInstance, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *glooInstanceClient) PatchGlooInstance(ctx context.Context, obj *GlooInstance, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *glooInstanceClient) DeleteAllOfGlooInstance(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &GlooInstance{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *glooInstanceClient) UpsertGlooInstance(ctx context.Context, obj *GlooInstance, transitionFuncs ...GlooInstanceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*GlooInstance), desired.(*GlooInstance)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *glooInstanceClient) UpdateGlooInstanceStatus(ctx context.Context, obj *GlooInstance, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *glooInstanceClient) PatchGlooInstanceStatus(ctx context.Context, obj *GlooInstance, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides GlooInstanceClients for multiple clusters.
type MulticlusterGlooInstanceClient interface {
	// Cluster returns a GlooInstanceClient for the given cluster
	Cluster(cluster string) (GlooInstanceClient, error)
}

type multiclusterGlooInstanceClient struct {
	client multicluster.Client
}

func NewMulticlusterGlooInstanceClient(client multicluster.Client) MulticlusterGlooInstanceClient {
	return &multiclusterGlooInstanceClient{client: client}
}

func (m *multiclusterGlooInstanceClient) Cluster(cluster string) (GlooInstanceClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewGlooInstanceClient(client), nil
}
