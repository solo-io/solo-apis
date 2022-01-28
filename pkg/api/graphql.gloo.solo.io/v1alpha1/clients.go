// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the graphql.gloo.solo.io/v1alpha1 APIs
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

// clienset for the graphql.gloo.solo.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the graphql.gloo.solo.io/v1alpha1/v1alpha1 APIs
	GraphQLSchemas() GraphQLSchemaClient
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

// clienset for the graphql.gloo.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) GraphQLSchemas() GraphQLSchemaClient {
	return NewGraphQLSchemaClient(c.client)
}

// Reader knows how to read and list GraphQLSchemas.
type GraphQLSchemaReader interface {
	// Get retrieves a GraphQLSchema for the given object key
	GetGraphQLSchema(ctx context.Context, key client.ObjectKey) (*GraphQLSchema, error)

	// List retrieves list of GraphQLSchemas for a given namespace and list options.
	ListGraphQLSchema(ctx context.Context, opts ...client.ListOption) (*GraphQLSchemaList, error)
}

// GraphQLSchemaTransitionFunction instructs the GraphQLSchemaWriter how to transition between an existing
// GraphQLSchema object and a desired on an Upsert
type GraphQLSchemaTransitionFunction func(existing, desired *GraphQLSchema) error

// Writer knows how to create, delete, and update GraphQLSchemas.
type GraphQLSchemaWriter interface {
	// Create saves the GraphQLSchema object.
	CreateGraphQLSchema(ctx context.Context, obj *GraphQLSchema, opts ...client.CreateOption) error

	// Delete deletes the GraphQLSchema object.
	DeleteGraphQLSchema(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given GraphQLSchema object.
	UpdateGraphQLSchema(ctx context.Context, obj *GraphQLSchema, opts ...client.UpdateOption) error

	// Patch patches the given GraphQLSchema object.
	PatchGraphQLSchema(ctx context.Context, obj *GraphQLSchema, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all GraphQLSchema objects matching the given options.
	DeleteAllOfGraphQLSchema(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the GraphQLSchema object.
	UpsertGraphQLSchema(ctx context.Context, obj *GraphQLSchema, transitionFuncs ...GraphQLSchemaTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a GraphQLSchema object.
type GraphQLSchemaStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given GraphQLSchema object.
	UpdateGraphQLSchemaStatus(ctx context.Context, obj *GraphQLSchema, opts ...client.UpdateOption) error

	// Patch patches the given GraphQLSchema object's subresource.
	PatchGraphQLSchemaStatus(ctx context.Context, obj *GraphQLSchema, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on GraphQLSchemas.
type GraphQLSchemaClient interface {
	GraphQLSchemaReader
	GraphQLSchemaWriter
	GraphQLSchemaStatusWriter
}

type graphQLSchemaClient struct {
	client client.Client
}

func NewGraphQLSchemaClient(client client.Client) *graphQLSchemaClient {
	return &graphQLSchemaClient{client: client}
}

func (c *graphQLSchemaClient) GetGraphQLSchema(ctx context.Context, key client.ObjectKey) (*GraphQLSchema, error) {
	obj := &GraphQLSchema{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *graphQLSchemaClient) ListGraphQLSchema(ctx context.Context, opts ...client.ListOption) (*GraphQLSchemaList, error) {
	list := &GraphQLSchemaList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *graphQLSchemaClient) CreateGraphQLSchema(ctx context.Context, obj *GraphQLSchema, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *graphQLSchemaClient) DeleteGraphQLSchema(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &GraphQLSchema{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *graphQLSchemaClient) UpdateGraphQLSchema(ctx context.Context, obj *GraphQLSchema, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *graphQLSchemaClient) PatchGraphQLSchema(ctx context.Context, obj *GraphQLSchema, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *graphQLSchemaClient) DeleteAllOfGraphQLSchema(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &GraphQLSchema{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *graphQLSchemaClient) UpsertGraphQLSchema(ctx context.Context, obj *GraphQLSchema, transitionFuncs ...GraphQLSchemaTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*GraphQLSchema), desired.(*GraphQLSchema)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *graphQLSchemaClient) UpdateGraphQLSchemaStatus(ctx context.Context, obj *GraphQLSchema, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *graphQLSchemaClient) PatchGraphQLSchemaStatus(ctx context.Context, obj *GraphQLSchema, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides GraphQLSchemaClients for multiple clusters.
type MulticlusterGraphQLSchemaClient interface {
	// Cluster returns a GraphQLSchemaClient for the given cluster
	Cluster(cluster string) (GraphQLSchemaClient, error)
}

type multiclusterGraphQLSchemaClient struct {
	client multicluster.Client
}

func NewMulticlusterGraphQLSchemaClient(client multicluster.Client) MulticlusterGraphQLSchemaClient {
	return &multiclusterGraphQLSchemaClient{client: client}
}

func (m *multiclusterGraphQLSchemaClient) Cluster(cluster string) (GraphQLSchemaClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewGraphQLSchemaClient(client), nil
}
