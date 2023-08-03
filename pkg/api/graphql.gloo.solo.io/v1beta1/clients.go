// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1beta1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the graphql.gloo.solo.io/v1beta1 APIs
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

// clienset for the graphql.gloo.solo.io/v1beta1 APIs
type Clientset interface {
	// clienset for the graphql.gloo.solo.io/v1beta1/v1beta1 APIs
	GraphQLApis() GraphQLApiClient
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

// clienset for the graphql.gloo.solo.io/v1beta1/v1beta1 APIs
func (c *clientSet) GraphQLApis() GraphQLApiClient {
	return NewGraphQLApiClient(c.client)
}

// Reader knows how to read and list GraphQLApis.
type GraphQLApiReader interface {
	// Get retrieves a GraphQLApi for the given object key
	GetGraphQLApi(ctx context.Context, key client.ObjectKey) (*GraphQLApi, error)

	// List retrieves list of GraphQLApis for a given namespace and list options.
	ListGraphQLApi(ctx context.Context, opts ...client.ListOption) (*GraphQLApiList, error)
}

// GraphQLApiTransitionFunction instructs the GraphQLApiWriter how to transition between an existing
// GraphQLApi object and a desired on an Upsert
type GraphQLApiTransitionFunction func(existing, desired *GraphQLApi) error

// Writer knows how to create, delete, and update GraphQLApis.
type GraphQLApiWriter interface {
	// Create saves the GraphQLApi object.
	CreateGraphQLApi(ctx context.Context, obj *GraphQLApi, opts ...client.CreateOption) error

	// Delete deletes the GraphQLApi object.
	DeleteGraphQLApi(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given GraphQLApi object.
	UpdateGraphQLApi(ctx context.Context, obj *GraphQLApi, opts ...client.UpdateOption) error

	// Patch patches the given GraphQLApi object.
	PatchGraphQLApi(ctx context.Context, obj *GraphQLApi, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all GraphQLApi objects matching the given options.
	DeleteAllOfGraphQLApi(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the GraphQLApi object.
	UpsertGraphQLApi(ctx context.Context, obj *GraphQLApi, transitionFuncs ...GraphQLApiTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a GraphQLApi object.
type GraphQLApiStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given GraphQLApi object.
	UpdateGraphQLApiStatus(ctx context.Context, obj *GraphQLApi, opts ...client.UpdateOption) error

	// Patch patches the given GraphQLApi object's subresource.
	PatchGraphQLApiStatus(ctx context.Context, obj *GraphQLApi, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on GraphQLApis.
type GraphQLApiClient interface {
	GraphQLApiReader
	GraphQLApiWriter
	GraphQLApiStatusWriter
}

type graphQLApiClient struct {
	client client.Client
}

func NewGraphQLApiClient(client client.Client) *graphQLApiClient {
	return &graphQLApiClient{client: client}
}

func (c *graphQLApiClient) GetGraphQLApi(ctx context.Context, key client.ObjectKey) (*GraphQLApi, error) {
	obj := &GraphQLApi{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *graphQLApiClient) ListGraphQLApi(ctx context.Context, opts ...client.ListOption) (*GraphQLApiList, error) {
	list := &GraphQLApiList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *graphQLApiClient) CreateGraphQLApi(ctx context.Context, obj *GraphQLApi, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *graphQLApiClient) DeleteGraphQLApi(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &GraphQLApi{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *graphQLApiClient) UpdateGraphQLApi(ctx context.Context, obj *GraphQLApi, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *graphQLApiClient) PatchGraphQLApi(ctx context.Context, obj *GraphQLApi, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *graphQLApiClient) DeleteAllOfGraphQLApi(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &GraphQLApi{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *graphQLApiClient) UpsertGraphQLApi(ctx context.Context, obj *GraphQLApi, transitionFuncs ...GraphQLApiTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*GraphQLApi), desired.(*GraphQLApi)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *graphQLApiClient) UpdateGraphQLApiStatus(ctx context.Context, obj *GraphQLApi, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *graphQLApiClient) PatchGraphQLApiStatus(ctx context.Context, obj *GraphQLApi, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides GraphQLApiClients for multiple clusters.
type MulticlusterGraphQLApiClient interface {
	// Cluster returns a GraphQLApiClient for the given cluster
	Cluster(cluster string) (GraphQLApiClient, error)
}

type multiclusterGraphQLApiClient struct {
	client multicluster.Client
}

func NewMulticlusterGraphQLApiClient(client multicluster.Client) MulticlusterGraphQLApiClient {
	return &multiclusterGraphQLApiClient{client: client}
}

func (m *multiclusterGraphQLApiClient) Cluster(cluster string) (GraphQLApiClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewGraphQLApiClient(client), nil
}
