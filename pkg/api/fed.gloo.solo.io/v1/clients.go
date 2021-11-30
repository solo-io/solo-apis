// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

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

// MulticlusterClientset for the fed.gloo.solo.io/v1 APIs
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

// clienset for the fed.gloo.solo.io/v1 APIs
type Clientset interface {
	// clienset for the fed.gloo.solo.io/v1/v1 APIs
	FederatedSettings() FederatedSettingsClient
	// clienset for the fed.gloo.solo.io/v1/v1 APIs
	FederatedUpstreams() FederatedUpstreamClient
	// clienset for the fed.gloo.solo.io/v1/v1 APIs
	FederatedUpstreamGroups() FederatedUpstreamGroupClient
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

// clienset for the fed.gloo.solo.io/v1/v1 APIs
func (c *clientSet) FederatedSettings() FederatedSettingsClient {
	return NewFederatedSettingsClient(c.client)
}

// clienset for the fed.gloo.solo.io/v1/v1 APIs
func (c *clientSet) FederatedUpstreams() FederatedUpstreamClient {
	return NewFederatedUpstreamClient(c.client)
}

// clienset for the fed.gloo.solo.io/v1/v1 APIs
func (c *clientSet) FederatedUpstreamGroups() FederatedUpstreamGroupClient {
	return NewFederatedUpstreamGroupClient(c.client)
}

// Reader knows how to read and list FederatedSettingss.
type FederatedSettingsReader interface {
	// Get retrieves a FederatedSettings for the given object key
	GetFederatedSettings(ctx context.Context, key client.ObjectKey) (*FederatedSettings, error)

	// List retrieves list of FederatedSettingss for a given namespace and list options.
	ListFederatedSettings(ctx context.Context, opts ...client.ListOption) (*FederatedSettingsList, error)
}

// FederatedSettingsTransitionFunction instructs the FederatedSettingsWriter how to transition between an existing
// FederatedSettings object and a desired on an Upsert
type FederatedSettingsTransitionFunction func(existing, desired *FederatedSettings) error

// Writer knows how to create, delete, and update FederatedSettingss.
type FederatedSettingsWriter interface {
	// Create saves the FederatedSettings object.
	CreateFederatedSettings(ctx context.Context, obj *FederatedSettings, opts ...client.CreateOption) error

	// Delete deletes the FederatedSettings object.
	DeleteFederatedSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given FederatedSettings object.
	UpdateFederatedSettings(ctx context.Context, obj *FederatedSettings, opts ...client.UpdateOption) error

	// Patch patches the given FederatedSettings object.
	PatchFederatedSettings(ctx context.Context, obj *FederatedSettings, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all FederatedSettings objects matching the given options.
	DeleteAllOfFederatedSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the FederatedSettings object.
	UpsertFederatedSettings(ctx context.Context, obj *FederatedSettings, transitionFuncs ...FederatedSettingsTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a FederatedSettings object.
type FederatedSettingsStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given FederatedSettings object.
	UpdateFederatedSettingsStatus(ctx context.Context, obj *FederatedSettings, opts ...client.UpdateOption) error

	// Patch patches the given FederatedSettings object's subresource.
	PatchFederatedSettingsStatus(ctx context.Context, obj *FederatedSettings, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on FederatedSettingss.
type FederatedSettingsClient interface {
	FederatedSettingsReader
	FederatedSettingsWriter
	FederatedSettingsStatusWriter
}

type federatedSettingsClient struct {
	client client.Client
}

func NewFederatedSettingsClient(client client.Client) *federatedSettingsClient {
	return &federatedSettingsClient{client: client}
}

func (c *federatedSettingsClient) GetFederatedSettings(ctx context.Context, key client.ObjectKey) (*FederatedSettings, error) {
	obj := &FederatedSettings{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *federatedSettingsClient) ListFederatedSettings(ctx context.Context, opts ...client.ListOption) (*FederatedSettingsList, error) {
	list := &FederatedSettingsList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *federatedSettingsClient) CreateFederatedSettings(ctx context.Context, obj *FederatedSettings, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *federatedSettingsClient) DeleteFederatedSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &FederatedSettings{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *federatedSettingsClient) UpdateFederatedSettings(ctx context.Context, obj *FederatedSettings, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *federatedSettingsClient) PatchFederatedSettings(ctx context.Context, obj *FederatedSettings, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *federatedSettingsClient) DeleteAllOfFederatedSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &FederatedSettings{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *federatedSettingsClient) UpsertFederatedSettings(ctx context.Context, obj *FederatedSettings, transitionFuncs ...FederatedSettingsTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*FederatedSettings), desired.(*FederatedSettings)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *federatedSettingsClient) UpdateFederatedSettingsStatus(ctx context.Context, obj *FederatedSettings, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *federatedSettingsClient) PatchFederatedSettingsStatus(ctx context.Context, obj *FederatedSettings, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides FederatedSettingsClients for multiple clusters.
type MulticlusterFederatedSettingsClient interface {
	// Cluster returns a FederatedSettingsClient for the given cluster
	Cluster(cluster string) (FederatedSettingsClient, error)
}

type multiclusterFederatedSettingsClient struct {
	client multicluster.Client
}

func NewMulticlusterFederatedSettingsClient(client multicluster.Client) MulticlusterFederatedSettingsClient {
	return &multiclusterFederatedSettingsClient{client: client}
}

func (m *multiclusterFederatedSettingsClient) Cluster(cluster string) (FederatedSettingsClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewFederatedSettingsClient(client), nil
}

// Reader knows how to read and list FederatedUpstreams.
type FederatedUpstreamReader interface {
	// Get retrieves a FederatedUpstream for the given object key
	GetFederatedUpstream(ctx context.Context, key client.ObjectKey) (*FederatedUpstream, error)

	// List retrieves list of FederatedUpstreams for a given namespace and list options.
	ListFederatedUpstream(ctx context.Context, opts ...client.ListOption) (*FederatedUpstreamList, error)
}

// FederatedUpstreamTransitionFunction instructs the FederatedUpstreamWriter how to transition between an existing
// FederatedUpstream object and a desired on an Upsert
type FederatedUpstreamTransitionFunction func(existing, desired *FederatedUpstream) error

// Writer knows how to create, delete, and update FederatedUpstreams.
type FederatedUpstreamWriter interface {
	// Create saves the FederatedUpstream object.
	CreateFederatedUpstream(ctx context.Context, obj *FederatedUpstream, opts ...client.CreateOption) error

	// Delete deletes the FederatedUpstream object.
	DeleteFederatedUpstream(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given FederatedUpstream object.
	UpdateFederatedUpstream(ctx context.Context, obj *FederatedUpstream, opts ...client.UpdateOption) error

	// Patch patches the given FederatedUpstream object.
	PatchFederatedUpstream(ctx context.Context, obj *FederatedUpstream, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all FederatedUpstream objects matching the given options.
	DeleteAllOfFederatedUpstream(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the FederatedUpstream object.
	UpsertFederatedUpstream(ctx context.Context, obj *FederatedUpstream, transitionFuncs ...FederatedUpstreamTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a FederatedUpstream object.
type FederatedUpstreamStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given FederatedUpstream object.
	UpdateFederatedUpstreamStatus(ctx context.Context, obj *FederatedUpstream, opts ...client.UpdateOption) error

	// Patch patches the given FederatedUpstream object's subresource.
	PatchFederatedUpstreamStatus(ctx context.Context, obj *FederatedUpstream, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on FederatedUpstreams.
type FederatedUpstreamClient interface {
	FederatedUpstreamReader
	FederatedUpstreamWriter
	FederatedUpstreamStatusWriter
}

type federatedUpstreamClient struct {
	client client.Client
}

func NewFederatedUpstreamClient(client client.Client) *federatedUpstreamClient {
	return &federatedUpstreamClient{client: client}
}

func (c *federatedUpstreamClient) GetFederatedUpstream(ctx context.Context, key client.ObjectKey) (*FederatedUpstream, error) {
	obj := &FederatedUpstream{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *federatedUpstreamClient) ListFederatedUpstream(ctx context.Context, opts ...client.ListOption) (*FederatedUpstreamList, error) {
	list := &FederatedUpstreamList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *federatedUpstreamClient) CreateFederatedUpstream(ctx context.Context, obj *FederatedUpstream, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *federatedUpstreamClient) DeleteFederatedUpstream(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &FederatedUpstream{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *federatedUpstreamClient) UpdateFederatedUpstream(ctx context.Context, obj *FederatedUpstream, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *federatedUpstreamClient) PatchFederatedUpstream(ctx context.Context, obj *FederatedUpstream, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *federatedUpstreamClient) DeleteAllOfFederatedUpstream(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &FederatedUpstream{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *federatedUpstreamClient) UpsertFederatedUpstream(ctx context.Context, obj *FederatedUpstream, transitionFuncs ...FederatedUpstreamTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*FederatedUpstream), desired.(*FederatedUpstream)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *federatedUpstreamClient) UpdateFederatedUpstreamStatus(ctx context.Context, obj *FederatedUpstream, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *federatedUpstreamClient) PatchFederatedUpstreamStatus(ctx context.Context, obj *FederatedUpstream, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides FederatedUpstreamClients for multiple clusters.
type MulticlusterFederatedUpstreamClient interface {
	// Cluster returns a FederatedUpstreamClient for the given cluster
	Cluster(cluster string) (FederatedUpstreamClient, error)
}

type multiclusterFederatedUpstreamClient struct {
	client multicluster.Client
}

func NewMulticlusterFederatedUpstreamClient(client multicluster.Client) MulticlusterFederatedUpstreamClient {
	return &multiclusterFederatedUpstreamClient{client: client}
}

func (m *multiclusterFederatedUpstreamClient) Cluster(cluster string) (FederatedUpstreamClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewFederatedUpstreamClient(client), nil
}

// Reader knows how to read and list FederatedUpstreamGroups.
type FederatedUpstreamGroupReader interface {
	// Get retrieves a FederatedUpstreamGroup for the given object key
	GetFederatedUpstreamGroup(ctx context.Context, key client.ObjectKey) (*FederatedUpstreamGroup, error)

	// List retrieves list of FederatedUpstreamGroups for a given namespace and list options.
	ListFederatedUpstreamGroup(ctx context.Context, opts ...client.ListOption) (*FederatedUpstreamGroupList, error)
}

// FederatedUpstreamGroupTransitionFunction instructs the FederatedUpstreamGroupWriter how to transition between an existing
// FederatedUpstreamGroup object and a desired on an Upsert
type FederatedUpstreamGroupTransitionFunction func(existing, desired *FederatedUpstreamGroup) error

// Writer knows how to create, delete, and update FederatedUpstreamGroups.
type FederatedUpstreamGroupWriter interface {
	// Create saves the FederatedUpstreamGroup object.
	CreateFederatedUpstreamGroup(ctx context.Context, obj *FederatedUpstreamGroup, opts ...client.CreateOption) error

	// Delete deletes the FederatedUpstreamGroup object.
	DeleteFederatedUpstreamGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given FederatedUpstreamGroup object.
	UpdateFederatedUpstreamGroup(ctx context.Context, obj *FederatedUpstreamGroup, opts ...client.UpdateOption) error

	// Patch patches the given FederatedUpstreamGroup object.
	PatchFederatedUpstreamGroup(ctx context.Context, obj *FederatedUpstreamGroup, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all FederatedUpstreamGroup objects matching the given options.
	DeleteAllOfFederatedUpstreamGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the FederatedUpstreamGroup object.
	UpsertFederatedUpstreamGroup(ctx context.Context, obj *FederatedUpstreamGroup, transitionFuncs ...FederatedUpstreamGroupTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a FederatedUpstreamGroup object.
type FederatedUpstreamGroupStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given FederatedUpstreamGroup object.
	UpdateFederatedUpstreamGroupStatus(ctx context.Context, obj *FederatedUpstreamGroup, opts ...client.UpdateOption) error

	// Patch patches the given FederatedUpstreamGroup object's subresource.
	PatchFederatedUpstreamGroupStatus(ctx context.Context, obj *FederatedUpstreamGroup, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on FederatedUpstreamGroups.
type FederatedUpstreamGroupClient interface {
	FederatedUpstreamGroupReader
	FederatedUpstreamGroupWriter
	FederatedUpstreamGroupStatusWriter
}

type federatedUpstreamGroupClient struct {
	client client.Client
}

func NewFederatedUpstreamGroupClient(client client.Client) *federatedUpstreamGroupClient {
	return &federatedUpstreamGroupClient{client: client}
}

func (c *federatedUpstreamGroupClient) GetFederatedUpstreamGroup(ctx context.Context, key client.ObjectKey) (*FederatedUpstreamGroup, error) {
	obj := &FederatedUpstreamGroup{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *federatedUpstreamGroupClient) ListFederatedUpstreamGroup(ctx context.Context, opts ...client.ListOption) (*FederatedUpstreamGroupList, error) {
	list := &FederatedUpstreamGroupList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *federatedUpstreamGroupClient) CreateFederatedUpstreamGroup(ctx context.Context, obj *FederatedUpstreamGroup, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *federatedUpstreamGroupClient) DeleteFederatedUpstreamGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &FederatedUpstreamGroup{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *federatedUpstreamGroupClient) UpdateFederatedUpstreamGroup(ctx context.Context, obj *FederatedUpstreamGroup, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *federatedUpstreamGroupClient) PatchFederatedUpstreamGroup(ctx context.Context, obj *FederatedUpstreamGroup, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *federatedUpstreamGroupClient) DeleteAllOfFederatedUpstreamGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &FederatedUpstreamGroup{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *federatedUpstreamGroupClient) UpsertFederatedUpstreamGroup(ctx context.Context, obj *FederatedUpstreamGroup, transitionFuncs ...FederatedUpstreamGroupTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*FederatedUpstreamGroup), desired.(*FederatedUpstreamGroup)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *federatedUpstreamGroupClient) UpdateFederatedUpstreamGroupStatus(ctx context.Context, obj *FederatedUpstreamGroup, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *federatedUpstreamGroupClient) PatchFederatedUpstreamGroupStatus(ctx context.Context, obj *FederatedUpstreamGroup, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides FederatedUpstreamGroupClients for multiple clusters.
type MulticlusterFederatedUpstreamGroupClient interface {
	// Cluster returns a FederatedUpstreamGroupClient for the given cluster
	Cluster(cluster string) (FederatedUpstreamGroupClient, error)
}

type multiclusterFederatedUpstreamGroupClient struct {
	client multicluster.Client
}

func NewMulticlusterFederatedUpstreamGroupClient(client multicluster.Client) MulticlusterFederatedUpstreamGroupClient {
	return &multiclusterFederatedUpstreamGroupClient{client: client}
}

func (m *multiclusterFederatedUpstreamGroupClient) Cluster(cluster string) (FederatedUpstreamGroupClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewFederatedUpstreamGroupClient(client), nil
}
