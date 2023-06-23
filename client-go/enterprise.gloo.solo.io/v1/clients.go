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

// MulticlusterClientset for the enterprise.gloo.solo.io/v1 APIs
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

// clienset for the enterprise.gloo.solo.io/v1 APIs
type Clientset interface {
	// clienset for the enterprise.gloo.solo.io/v1/v1 APIs
	AuthConfigs() AuthConfigClient
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

// clienset for the enterprise.gloo.solo.io/v1/v1 APIs
func (c *clientSet) AuthConfigs() AuthConfigClient {
	return NewAuthConfigClient(c.client)
}

// Reader knows how to read and list AuthConfigs.
type AuthConfigReader interface {
	// Get retrieves a AuthConfig for the given object key
	GetAuthConfig(ctx context.Context, key client.ObjectKey) (*AuthConfig, error)

	// List retrieves list of AuthConfigs for a given namespace and list options.
	ListAuthConfig(ctx context.Context, opts ...client.ListOption) (*AuthConfigList, error)
}

// AuthConfigTransitionFunction instructs the AuthConfigWriter how to transition between an existing
// AuthConfig object and a desired on an Upsert
type AuthConfigTransitionFunction func(existing, desired *AuthConfig) error

// Writer knows how to create, delete, and update AuthConfigs.
type AuthConfigWriter interface {
	// Create saves the AuthConfig object.
	CreateAuthConfig(ctx context.Context, obj *AuthConfig, opts ...client.CreateOption) error

	// Delete deletes the AuthConfig object.
	DeleteAuthConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given AuthConfig object.
	UpdateAuthConfig(ctx context.Context, obj *AuthConfig, opts ...client.UpdateOption) error

	// Patch patches the given AuthConfig object.
	PatchAuthConfig(ctx context.Context, obj *AuthConfig, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all AuthConfig objects matching the given options.
	DeleteAllOfAuthConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the AuthConfig object.
	UpsertAuthConfig(ctx context.Context, obj *AuthConfig, transitionFuncs ...AuthConfigTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a AuthConfig object.
type AuthConfigStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given AuthConfig object.
	UpdateAuthConfigStatus(ctx context.Context, obj *AuthConfig, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given AuthConfig object's subresource.
	PatchAuthConfigStatus(ctx context.Context, obj *AuthConfig, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on AuthConfigs.
type AuthConfigClient interface {
	AuthConfigReader
	AuthConfigWriter
	AuthConfigStatusWriter
}

type authConfigClient struct {
	client client.Client
}

func NewAuthConfigClient(client client.Client) *authConfigClient {
	return &authConfigClient{client: client}
}

func (c *authConfigClient) GetAuthConfig(ctx context.Context, key client.ObjectKey) (*AuthConfig, error) {
	obj := &AuthConfig{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *authConfigClient) ListAuthConfig(ctx context.Context, opts ...client.ListOption) (*AuthConfigList, error) {
	list := &AuthConfigList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *authConfigClient) CreateAuthConfig(ctx context.Context, obj *AuthConfig, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *authConfigClient) DeleteAuthConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &AuthConfig{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *authConfigClient) UpdateAuthConfig(ctx context.Context, obj *AuthConfig, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *authConfigClient) PatchAuthConfig(ctx context.Context, obj *AuthConfig, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *authConfigClient) DeleteAllOfAuthConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &AuthConfig{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *authConfigClient) UpsertAuthConfig(ctx context.Context, obj *AuthConfig, transitionFuncs ...AuthConfigTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*AuthConfig), desired.(*AuthConfig)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *authConfigClient) UpdateAuthConfigStatus(ctx context.Context, obj *AuthConfig, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *authConfigClient) PatchAuthConfigStatus(ctx context.Context, obj *AuthConfig, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides AuthConfigClients for multiple clusters.
type MulticlusterAuthConfigClient interface {
	// Cluster returns a AuthConfigClient for the given cluster
	Cluster(cluster string) (AuthConfigClient, error)
}

type multiclusterAuthConfigClient struct {
	client multicluster.Client
}

func NewMulticlusterAuthConfigClient(client multicluster.Client) MulticlusterAuthConfigClient {
	return &multiclusterAuthConfigClient{client: client}
}

func (m *multiclusterAuthConfigClient) Cluster(cluster string) (AuthConfigClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewAuthConfigClient(client), nil
}
