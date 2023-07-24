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

// MulticlusterClientset for the internal.gloo.solo.io/v2alpha1 APIs
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

// clienset for the internal.gloo.solo.io/v2alpha1 APIs
type Clientset interface {
	// clienset for the internal.gloo.solo.io/v2alpha1/v2alpha1 APIs
	SpireRegistrationEntries() SpireRegistrationEntryClient
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

// clienset for the internal.gloo.solo.io/v2alpha1/v2alpha1 APIs
func (c *clientSet) SpireRegistrationEntries() SpireRegistrationEntryClient {
	return NewSpireRegistrationEntryClient(c.client)
}

// Reader knows how to read and list SpireRegistrationEntrys.
type SpireRegistrationEntryReader interface {
	// Get retrieves a SpireRegistrationEntry for the given object key
	GetSpireRegistrationEntry(ctx context.Context, key client.ObjectKey) (*SpireRegistrationEntry, error)

	// List retrieves list of SpireRegistrationEntrys for a given namespace and list options.
	ListSpireRegistrationEntry(ctx context.Context, opts ...client.ListOption) (*SpireRegistrationEntryList, error)
}

// SpireRegistrationEntryTransitionFunction instructs the SpireRegistrationEntryWriter how to transition between an existing
// SpireRegistrationEntry object and a desired on an Upsert
type SpireRegistrationEntryTransitionFunction func(existing, desired *SpireRegistrationEntry) error

// Writer knows how to create, delete, and update SpireRegistrationEntrys.
type SpireRegistrationEntryWriter interface {
	// Create saves the SpireRegistrationEntry object.
	CreateSpireRegistrationEntry(ctx context.Context, obj *SpireRegistrationEntry, opts ...client.CreateOption) error

	// Delete deletes the SpireRegistrationEntry object.
	DeleteSpireRegistrationEntry(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given SpireRegistrationEntry object.
	UpdateSpireRegistrationEntry(ctx context.Context, obj *SpireRegistrationEntry, opts ...client.UpdateOption) error

	// Patch patches the given SpireRegistrationEntry object.
	PatchSpireRegistrationEntry(ctx context.Context, obj *SpireRegistrationEntry, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all SpireRegistrationEntry objects matching the given options.
	DeleteAllOfSpireRegistrationEntry(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the SpireRegistrationEntry object.
	UpsertSpireRegistrationEntry(ctx context.Context, obj *SpireRegistrationEntry, transitionFuncs ...SpireRegistrationEntryTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a SpireRegistrationEntry object.
type SpireRegistrationEntryStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given SpireRegistrationEntry object.
	UpdateSpireRegistrationEntryStatus(ctx context.Context, obj *SpireRegistrationEntry, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given SpireRegistrationEntry object's subresource.
	PatchSpireRegistrationEntryStatus(ctx context.Context, obj *SpireRegistrationEntry, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on SpireRegistrationEntrys.
type SpireRegistrationEntryClient interface {
	SpireRegistrationEntryReader
	SpireRegistrationEntryWriter
	SpireRegistrationEntryStatusWriter
}

type spireRegistrationEntryClient struct {
	client client.Client
}

func NewSpireRegistrationEntryClient(client client.Client) *spireRegistrationEntryClient {
	return &spireRegistrationEntryClient{client: client}
}

func (c *spireRegistrationEntryClient) GetSpireRegistrationEntry(ctx context.Context, key client.ObjectKey) (*SpireRegistrationEntry, error) {
	obj := &SpireRegistrationEntry{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *spireRegistrationEntryClient) ListSpireRegistrationEntry(ctx context.Context, opts ...client.ListOption) (*SpireRegistrationEntryList, error) {
	list := &SpireRegistrationEntryList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *spireRegistrationEntryClient) CreateSpireRegistrationEntry(ctx context.Context, obj *SpireRegistrationEntry, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *spireRegistrationEntryClient) DeleteSpireRegistrationEntry(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &SpireRegistrationEntry{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *spireRegistrationEntryClient) UpdateSpireRegistrationEntry(ctx context.Context, obj *SpireRegistrationEntry, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *spireRegistrationEntryClient) PatchSpireRegistrationEntry(ctx context.Context, obj *SpireRegistrationEntry, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *spireRegistrationEntryClient) DeleteAllOfSpireRegistrationEntry(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &SpireRegistrationEntry{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *spireRegistrationEntryClient) UpsertSpireRegistrationEntry(ctx context.Context, obj *SpireRegistrationEntry, transitionFuncs ...SpireRegistrationEntryTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*SpireRegistrationEntry), desired.(*SpireRegistrationEntry)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *spireRegistrationEntryClient) UpdateSpireRegistrationEntryStatus(ctx context.Context, obj *SpireRegistrationEntry, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *spireRegistrationEntryClient) PatchSpireRegistrationEntryStatus(ctx context.Context, obj *SpireRegistrationEntry, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides SpireRegistrationEntryClients for multiple clusters.
type MulticlusterSpireRegistrationEntryClient interface {
	// Cluster returns a SpireRegistrationEntryClient for the given cluster
	Cluster(cluster string) (SpireRegistrationEntryClient, error)
}

type multiclusterSpireRegistrationEntryClient struct {
	client multicluster.Client
}

func NewMulticlusterSpireRegistrationEntryClient(client multicluster.Client) MulticlusterSpireRegistrationEntryClient {
	return &multiclusterSpireRegistrationEntryClient{client: client}
}

func (m *multiclusterSpireRegistrationEntryClient) Cluster(cluster string) (SpireRegistrationEntryClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewSpireRegistrationEntryClient(client), nil
}
