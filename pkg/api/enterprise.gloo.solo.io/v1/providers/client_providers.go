// Code generated by skv2. DO NOT EDIT.

package v1

import (
	enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for AuthConfigClient from Clientset
func AuthConfigClientFromClientsetProvider(clients enterprise_gloo_solo_io_v1.Clientset) enterprise_gloo_solo_io_v1.AuthConfigClient {
	return clients.AuthConfigs()
}

// Provider for AuthConfig Client from Client
func AuthConfigClientProvider(client client.Client) enterprise_gloo_solo_io_v1.AuthConfigClient {
	return enterprise_gloo_solo_io_v1.NewAuthConfigClient(client)
}

type AuthConfigClientFactory func(client client.Client) enterprise_gloo_solo_io_v1.AuthConfigClient

func AuthConfigClientFactoryProvider() AuthConfigClientFactory {
	return AuthConfigClientProvider
}

type AuthConfigClientFromConfigFactory func(cfg *rest.Config) (enterprise_gloo_solo_io_v1.AuthConfigClient, error)

func AuthConfigClientFromConfigFactoryProvider() AuthConfigClientFromConfigFactory {
	return func(cfg *rest.Config) (enterprise_gloo_solo_io_v1.AuthConfigClient, error) {
		clients, err := enterprise_gloo_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.AuthConfigs(), nil
	}
}