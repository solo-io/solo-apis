// Code generated by skv2. DO NOT EDIT.

package v2

import (
	apimanagement_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/apimanagement.gloo.solo.io/v2"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for GraphQLStitchedSchemaClient from Clientset
func GraphQLStitchedSchemaClientFromClientsetProvider(clients apimanagement_gloo_solo_io_v2.Clientset) apimanagement_gloo_solo_io_v2.GraphQLStitchedSchemaClient {
	return clients.GraphQLStitchedSchemas()
}

// Provider for GraphQLStitchedSchema Client from Client
func GraphQLStitchedSchemaClientProvider(client client.Client) apimanagement_gloo_solo_io_v2.GraphQLStitchedSchemaClient {
	return apimanagement_gloo_solo_io_v2.NewGraphQLStitchedSchemaClient(client)
}

type GraphQLStitchedSchemaClientFactory func(client client.Client) apimanagement_gloo_solo_io_v2.GraphQLStitchedSchemaClient

func GraphQLStitchedSchemaClientFactoryProvider() GraphQLStitchedSchemaClientFactory {
	return GraphQLStitchedSchemaClientProvider
}

type GraphQLStitchedSchemaClientFromConfigFactory func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.GraphQLStitchedSchemaClient, error)

func GraphQLStitchedSchemaClientFromConfigFactoryProvider() GraphQLStitchedSchemaClientFromConfigFactory {
	return func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.GraphQLStitchedSchemaClient, error) {
		clients, err := apimanagement_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.GraphQLStitchedSchemas(), nil
	}
}

// Provider for GraphQLResolverMapClient from Clientset
func GraphQLResolverMapClientFromClientsetProvider(clients apimanagement_gloo_solo_io_v2.Clientset) apimanagement_gloo_solo_io_v2.GraphQLResolverMapClient {
	return clients.GraphQLResolverMaps()
}

// Provider for GraphQLResolverMap Client from Client
func GraphQLResolverMapClientProvider(client client.Client) apimanagement_gloo_solo_io_v2.GraphQLResolverMapClient {
	return apimanagement_gloo_solo_io_v2.NewGraphQLResolverMapClient(client)
}

type GraphQLResolverMapClientFactory func(client client.Client) apimanagement_gloo_solo_io_v2.GraphQLResolverMapClient

func GraphQLResolverMapClientFactoryProvider() GraphQLResolverMapClientFactory {
	return GraphQLResolverMapClientProvider
}

type GraphQLResolverMapClientFromConfigFactory func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.GraphQLResolverMapClient, error)

func GraphQLResolverMapClientFromConfigFactoryProvider() GraphQLResolverMapClientFromConfigFactory {
	return func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.GraphQLResolverMapClient, error) {
		clients, err := apimanagement_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.GraphQLResolverMaps(), nil
	}
}

// Provider for GraphQLSchemaClient from Clientset
func GraphQLSchemaClientFromClientsetProvider(clients apimanagement_gloo_solo_io_v2.Clientset) apimanagement_gloo_solo_io_v2.GraphQLSchemaClient {
	return clients.GraphQLSchemas()
}

// Provider for GraphQLSchema Client from Client
func GraphQLSchemaClientProvider(client client.Client) apimanagement_gloo_solo_io_v2.GraphQLSchemaClient {
	return apimanagement_gloo_solo_io_v2.NewGraphQLSchemaClient(client)
}

type GraphQLSchemaClientFactory func(client client.Client) apimanagement_gloo_solo_io_v2.GraphQLSchemaClient

func GraphQLSchemaClientFactoryProvider() GraphQLSchemaClientFactory {
	return GraphQLSchemaClientProvider
}

type GraphQLSchemaClientFromConfigFactory func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.GraphQLSchemaClient, error)

func GraphQLSchemaClientFromConfigFactoryProvider() GraphQLSchemaClientFromConfigFactory {
	return func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.GraphQLSchemaClient, error) {
		clients, err := apimanagement_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.GraphQLSchemas(), nil
	}
}

// Provider for ApiDocClient from Clientset
func ApiDocClientFromClientsetProvider(clients apimanagement_gloo_solo_io_v2.Clientset) apimanagement_gloo_solo_io_v2.ApiDocClient {
	return clients.ApiDocs()
}

// Provider for ApiDoc Client from Client
func ApiDocClientProvider(client client.Client) apimanagement_gloo_solo_io_v2.ApiDocClient {
	return apimanagement_gloo_solo_io_v2.NewApiDocClient(client)
}

type ApiDocClientFactory func(client client.Client) apimanagement_gloo_solo_io_v2.ApiDocClient

func ApiDocClientFactoryProvider() ApiDocClientFactory {
	return ApiDocClientProvider
}

type ApiDocClientFromConfigFactory func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.ApiDocClient, error)

func ApiDocClientFromConfigFactoryProvider() ApiDocClientFromConfigFactory {
	return func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.ApiDocClient, error) {
		clients, err := apimanagement_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ApiDocs(), nil
	}
}

// Provider for PortalClient from Clientset
func PortalClientFromClientsetProvider(clients apimanagement_gloo_solo_io_v2.Clientset) apimanagement_gloo_solo_io_v2.PortalClient {
	return clients.Portals()
}

// Provider for Portal Client from Client
func PortalClientProvider(client client.Client) apimanagement_gloo_solo_io_v2.PortalClient {
	return apimanagement_gloo_solo_io_v2.NewPortalClient(client)
}

type PortalClientFactory func(client client.Client) apimanagement_gloo_solo_io_v2.PortalClient

func PortalClientFactoryProvider() PortalClientFactory {
	return PortalClientProvider
}

type PortalClientFromConfigFactory func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.PortalClient, error)

func PortalClientFromConfigFactoryProvider() PortalClientFromConfigFactory {
	return func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.PortalClient, error) {
		clients, err := apimanagement_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Portals(), nil
	}
}

// Provider for ApiProductClient from Clientset
func ApiProductClientFromClientsetProvider(clients apimanagement_gloo_solo_io_v2.Clientset) apimanagement_gloo_solo_io_v2.ApiProductClient {
	return clients.ApiProducts()
}

// Provider for ApiProduct Client from Client
func ApiProductClientProvider(client client.Client) apimanagement_gloo_solo_io_v2.ApiProductClient {
	return apimanagement_gloo_solo_io_v2.NewApiProductClient(client)
}

type ApiProductClientFactory func(client client.Client) apimanagement_gloo_solo_io_v2.ApiProductClient

func ApiProductClientFactoryProvider() ApiProductClientFactory {
	return ApiProductClientProvider
}

type ApiProductClientFromConfigFactory func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.ApiProductClient, error)

func ApiProductClientFromConfigFactoryProvider() ApiProductClientFromConfigFactory {
	return func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.ApiProductClient, error) {
		clients, err := apimanagement_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ApiProducts(), nil
	}
}

// Provider for PortalGroupClient from Clientset
func PortalGroupClientFromClientsetProvider(clients apimanagement_gloo_solo_io_v2.Clientset) apimanagement_gloo_solo_io_v2.PortalGroupClient {
	return clients.PortalGroups()
}

// Provider for PortalGroup Client from Client
func PortalGroupClientProvider(client client.Client) apimanagement_gloo_solo_io_v2.PortalGroupClient {
	return apimanagement_gloo_solo_io_v2.NewPortalGroupClient(client)
}

type PortalGroupClientFactory func(client client.Client) apimanagement_gloo_solo_io_v2.PortalGroupClient

func PortalGroupClientFactoryProvider() PortalGroupClientFactory {
	return PortalGroupClientProvider
}

type PortalGroupClientFromConfigFactory func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.PortalGroupClient, error)

func PortalGroupClientFromConfigFactoryProvider() PortalGroupClientFromConfigFactory {
	return func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.PortalGroupClient, error) {
		clients, err := apimanagement_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.PortalGroups(), nil
	}
}

// Provider for ApiSchemaDiscoveryClient from Clientset
func ApiSchemaDiscoveryClientFromClientsetProvider(clients apimanagement_gloo_solo_io_v2.Clientset) apimanagement_gloo_solo_io_v2.ApiSchemaDiscoveryClient {
	return clients.ApiSchemaDiscoveries()
}

// Provider for ApiSchemaDiscovery Client from Client
func ApiSchemaDiscoveryClientProvider(client client.Client) apimanagement_gloo_solo_io_v2.ApiSchemaDiscoveryClient {
	return apimanagement_gloo_solo_io_v2.NewApiSchemaDiscoveryClient(client)
}

type ApiSchemaDiscoveryClientFactory func(client client.Client) apimanagement_gloo_solo_io_v2.ApiSchemaDiscoveryClient

func ApiSchemaDiscoveryClientFactoryProvider() ApiSchemaDiscoveryClientFactory {
	return ApiSchemaDiscoveryClientProvider
}

type ApiSchemaDiscoveryClientFromConfigFactory func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.ApiSchemaDiscoveryClient, error)

func ApiSchemaDiscoveryClientFromConfigFactoryProvider() ApiSchemaDiscoveryClientFromConfigFactory {
	return func(cfg *rest.Config) (apimanagement_gloo_solo_io_v2.ApiSchemaDiscoveryClient, error) {
		clients, err := apimanagement_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ApiSchemaDiscoveries(), nil
	}
}
