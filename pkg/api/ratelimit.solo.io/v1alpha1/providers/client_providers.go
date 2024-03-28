// Code generated by skv2. DO NOT EDIT.

package v1alpha1

import (
	ratelimit_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for RateLimitConfigClient from Clientset
func RateLimitConfigClientFromClientsetProvider(clients ratelimit_solo_io_v1alpha1.Clientset) ratelimit_solo_io_v1alpha1.RateLimitConfigClient {
	return clients.RateLimitConfigs()
}

// Provider for RateLimitConfig Client from Client
func RateLimitConfigClientProvider(client client.Client) ratelimit_solo_io_v1alpha1.RateLimitConfigClient {
	return ratelimit_solo_io_v1alpha1.NewRateLimitConfigClient(client)
}

type RateLimitConfigClientFactory func(client client.Client) ratelimit_solo_io_v1alpha1.RateLimitConfigClient

func RateLimitConfigClientFactoryProvider() RateLimitConfigClientFactory {
	return RateLimitConfigClientProvider
}

type RateLimitConfigClientFromConfigFactory func(cfg *rest.Config) (ratelimit_solo_io_v1alpha1.RateLimitConfigClient, error)

func RateLimitConfigClientFromConfigFactoryProvider() RateLimitConfigClientFromConfigFactory {
	return func(cfg *rest.Config) (ratelimit_solo_io_v1alpha1.RateLimitConfigClient, error) {
		clients, err := ratelimit_solo_io_v1alpha1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.RateLimitConfigs(), nil
	}
}
