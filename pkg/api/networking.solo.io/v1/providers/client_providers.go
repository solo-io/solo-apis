// Code generated by skv2. DO NOT EDIT.

package v1

import (
	networking_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/networking.solo.io/v1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for TrafficPolicyClient from Clientset
func TrafficPolicyClientFromClientsetProvider(clients networking_solo_io_v1.Clientset) networking_solo_io_v1.TrafficPolicyClient {
	return clients.TrafficPolicies()
}

// Provider for TrafficPolicy Client from Client
func TrafficPolicyClientProvider(client client.Client) networking_solo_io_v1.TrafficPolicyClient {
	return networking_solo_io_v1.NewTrafficPolicyClient(client)
}

type TrafficPolicyClientFactory func(client client.Client) networking_solo_io_v1.TrafficPolicyClient

func TrafficPolicyClientFactoryProvider() TrafficPolicyClientFactory {
	return TrafficPolicyClientProvider
}

type TrafficPolicyClientFromConfigFactory func(cfg *rest.Config) (networking_solo_io_v1.TrafficPolicyClient, error)

func TrafficPolicyClientFromConfigFactoryProvider() TrafficPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_solo_io_v1.TrafficPolicyClient, error) {
		clients, err := networking_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.TrafficPolicies(), nil
	}
}

// Provider for AccessPolicyClient from Clientset
func AccessPolicyClientFromClientsetProvider(clients networking_solo_io_v1.Clientset) networking_solo_io_v1.AccessPolicyClient {
	return clients.AccessPolicies()
}

// Provider for AccessPolicy Client from Client
func AccessPolicyClientProvider(client client.Client) networking_solo_io_v1.AccessPolicyClient {
	return networking_solo_io_v1.NewAccessPolicyClient(client)
}

type AccessPolicyClientFactory func(client client.Client) networking_solo_io_v1.AccessPolicyClient

func AccessPolicyClientFactoryProvider() AccessPolicyClientFactory {
	return AccessPolicyClientProvider
}

type AccessPolicyClientFromConfigFactory func(cfg *rest.Config) (networking_solo_io_v1.AccessPolicyClient, error)

func AccessPolicyClientFromConfigFactoryProvider() AccessPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_solo_io_v1.AccessPolicyClient, error) {
		clients, err := networking_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.AccessPolicies(), nil
	}
}

// Provider for VirtualMeshClient from Clientset
func VirtualMeshClientFromClientsetProvider(clients networking_solo_io_v1.Clientset) networking_solo_io_v1.VirtualMeshClient {
	return clients.VirtualMeshes()
}

// Provider for VirtualMesh Client from Client
func VirtualMeshClientProvider(client client.Client) networking_solo_io_v1.VirtualMeshClient {
	return networking_solo_io_v1.NewVirtualMeshClient(client)
}

type VirtualMeshClientFactory func(client client.Client) networking_solo_io_v1.VirtualMeshClient

func VirtualMeshClientFactoryProvider() VirtualMeshClientFactory {
	return VirtualMeshClientProvider
}

type VirtualMeshClientFromConfigFactory func(cfg *rest.Config) (networking_solo_io_v1.VirtualMeshClient, error)

func VirtualMeshClientFromConfigFactoryProvider() VirtualMeshClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_solo_io_v1.VirtualMeshClient, error) {
		clients, err := networking_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualMeshes(), nil
	}
}
