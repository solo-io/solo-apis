// Code generated by skv2. DO NOT EDIT.

package v1

import (
	rbac_enterprise_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/rbac.enterprise.solo.io/v1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for RoleClient from Clientset
func RoleClientFromClientsetProvider(clients rbac_enterprise_solo_io_v1.Clientset) rbac_enterprise_solo_io_v1.RoleClient {
	return clients.Roles()
}

// Provider for Role Client from Client
func RoleClientProvider(client client.Client) rbac_enterprise_solo_io_v1.RoleClient {
	return rbac_enterprise_solo_io_v1.NewRoleClient(client)
}

type RoleClientFactory func(client client.Client) rbac_enterprise_solo_io_v1.RoleClient

func RoleClientFactoryProvider() RoleClientFactory {
	return RoleClientProvider
}

type RoleClientFromConfigFactory func(cfg *rest.Config) (rbac_enterprise_solo_io_v1.RoleClient, error)

func RoleClientFromConfigFactoryProvider() RoleClientFromConfigFactory {
	return func(cfg *rest.Config) (rbac_enterprise_solo_io_v1.RoleClient, error) {
		clients, err := rbac_enterprise_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Roles(), nil
	}
}

// Provider for RoleBindingClient from Clientset
func RoleBindingClientFromClientsetProvider(clients rbac_enterprise_solo_io_v1.Clientset) rbac_enterprise_solo_io_v1.RoleBindingClient {
	return clients.RoleBindings()
}

// Provider for RoleBinding Client from Client
func RoleBindingClientProvider(client client.Client) rbac_enterprise_solo_io_v1.RoleBindingClient {
	return rbac_enterprise_solo_io_v1.NewRoleBindingClient(client)
}

type RoleBindingClientFactory func(client client.Client) rbac_enterprise_solo_io_v1.RoleBindingClient

func RoleBindingClientFactoryProvider() RoleBindingClientFactory {
	return RoleBindingClientProvider
}

type RoleBindingClientFromConfigFactory func(cfg *rest.Config) (rbac_enterprise_solo_io_v1.RoleBindingClient, error)

func RoleBindingClientFromConfigFactoryProvider() RoleBindingClientFromConfigFactory {
	return func(cfg *rest.Config) (rbac_enterprise_solo_io_v1.RoleBindingClient, error) {
		clients, err := rbac_enterprise_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.RoleBindings(), nil
	}
}
