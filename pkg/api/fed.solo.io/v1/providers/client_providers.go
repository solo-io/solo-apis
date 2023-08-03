// Code generated by skv2. DO NOT EDIT.

package v1



import (
    fed_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1"

    "k8s.io/client-go/rest"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for GlooInstanceClient from Clientset
func GlooInstanceClientFromClientsetProvider(clients fed_solo_io_v1.Clientset) fed_solo_io_v1.GlooInstanceClient {
    return clients.GlooInstances()
}

// Provider for GlooInstance Client from Client
func GlooInstanceClientProvider(client client.Client) fed_solo_io_v1.GlooInstanceClient {
    return fed_solo_io_v1.NewGlooInstanceClient(client)
}

type GlooInstanceClientFactory func(client client.Client) fed_solo_io_v1.GlooInstanceClient

func GlooInstanceClientFactoryProvider() GlooInstanceClientFactory {
    return GlooInstanceClientProvider
}

type GlooInstanceClientFromConfigFactory func(cfg *rest.Config) (fed_solo_io_v1.GlooInstanceClient, error)

func GlooInstanceClientFromConfigFactoryProvider() GlooInstanceClientFromConfigFactory {
    return func(cfg *rest.Config) (fed_solo_io_v1.GlooInstanceClient, error) {
        clients, err := fed_solo_io_v1.NewClientsetFromConfig(cfg)
        if err != nil {
            return nil, err
        }
        return clients.GlooInstances(), nil
    }
}

// Provider for FailoverSchemeClient from Clientset
func FailoverSchemeClientFromClientsetProvider(clients fed_solo_io_v1.Clientset) fed_solo_io_v1.FailoverSchemeClient {
    return clients.FailoverSchemes()
}

// Provider for FailoverScheme Client from Client
func FailoverSchemeClientProvider(client client.Client) fed_solo_io_v1.FailoverSchemeClient {
    return fed_solo_io_v1.NewFailoverSchemeClient(client)
}

type FailoverSchemeClientFactory func(client client.Client) fed_solo_io_v1.FailoverSchemeClient

func FailoverSchemeClientFactoryProvider() FailoverSchemeClientFactory {
    return FailoverSchemeClientProvider
}

type FailoverSchemeClientFromConfigFactory func(cfg *rest.Config) (fed_solo_io_v1.FailoverSchemeClient, error)

func FailoverSchemeClientFromConfigFactoryProvider() FailoverSchemeClientFromConfigFactory {
    return func(cfg *rest.Config) (fed_solo_io_v1.FailoverSchemeClient, error) {
        clients, err := fed_solo_io_v1.NewClientsetFromConfig(cfg)
        if err != nil {
            return nil, err
        }
        return clients.FailoverSchemes(), nil
    }
}