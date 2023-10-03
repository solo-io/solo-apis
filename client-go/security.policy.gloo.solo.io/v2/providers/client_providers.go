// Code generated by skv2. DO NOT EDIT.

package v2

import (
	security_policy_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/security.policy.gloo.solo.io/v2"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for AccessPolicyClient from Clientset
func AccessPolicyClientFromClientsetProvider(clients security_policy_gloo_solo_io_v2.Clientset) security_policy_gloo_solo_io_v2.AccessPolicyClient {
	return clients.AccessPolicies()
}

// Provider for AccessPolicy Client from Client
func AccessPolicyClientProvider(client client.Client) security_policy_gloo_solo_io_v2.AccessPolicyClient {
	return security_policy_gloo_solo_io_v2.NewAccessPolicyClient(client)
}

type AccessPolicyClientFactory func(client client.Client) security_policy_gloo_solo_io_v2.AccessPolicyClient

func AccessPolicyClientFactoryProvider() AccessPolicyClientFactory {
	return AccessPolicyClientProvider
}

type AccessPolicyClientFromConfigFactory func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.AccessPolicyClient, error)

func AccessPolicyClientFromConfigFactoryProvider() AccessPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.AccessPolicyClient, error) {
		clients, err := security_policy_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.AccessPolicies(), nil
	}
}

// Provider for CORSPolicyClient from Clientset
func CORSPolicyClientFromClientsetProvider(clients security_policy_gloo_solo_io_v2.Clientset) security_policy_gloo_solo_io_v2.CORSPolicyClient {
	return clients.CORSPolicies()
}

// Provider for CORSPolicy Client from Client
func CORSPolicyClientProvider(client client.Client) security_policy_gloo_solo_io_v2.CORSPolicyClient {
	return security_policy_gloo_solo_io_v2.NewCORSPolicyClient(client)
}

type CORSPolicyClientFactory func(client client.Client) security_policy_gloo_solo_io_v2.CORSPolicyClient

func CORSPolicyClientFactoryProvider() CORSPolicyClientFactory {
	return CORSPolicyClientProvider
}

type CORSPolicyClientFromConfigFactory func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.CORSPolicyClient, error)

func CORSPolicyClientFromConfigFactoryProvider() CORSPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.CORSPolicyClient, error) {
		clients, err := security_policy_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.CORSPolicies(), nil
	}
}

// Provider for CSRFPolicyClient from Clientset
func CSRFPolicyClientFromClientsetProvider(clients security_policy_gloo_solo_io_v2.Clientset) security_policy_gloo_solo_io_v2.CSRFPolicyClient {
	return clients.CSRFPolicies()
}

// Provider for CSRFPolicy Client from Client
func CSRFPolicyClientProvider(client client.Client) security_policy_gloo_solo_io_v2.CSRFPolicyClient {
	return security_policy_gloo_solo_io_v2.NewCSRFPolicyClient(client)
}

type CSRFPolicyClientFactory func(client client.Client) security_policy_gloo_solo_io_v2.CSRFPolicyClient

func CSRFPolicyClientFactoryProvider() CSRFPolicyClientFactory {
	return CSRFPolicyClientProvider
}

type CSRFPolicyClientFromConfigFactory func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.CSRFPolicyClient, error)

func CSRFPolicyClientFromConfigFactoryProvider() CSRFPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.CSRFPolicyClient, error) {
		clients, err := security_policy_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.CSRFPolicies(), nil
	}
}

// Provider for ExtAuthPolicyClient from Clientset
func ExtAuthPolicyClientFromClientsetProvider(clients security_policy_gloo_solo_io_v2.Clientset) security_policy_gloo_solo_io_v2.ExtAuthPolicyClient {
	return clients.ExtAuthPolicies()
}

// Provider for ExtAuthPolicy Client from Client
func ExtAuthPolicyClientProvider(client client.Client) security_policy_gloo_solo_io_v2.ExtAuthPolicyClient {
	return security_policy_gloo_solo_io_v2.NewExtAuthPolicyClient(client)
}

type ExtAuthPolicyClientFactory func(client client.Client) security_policy_gloo_solo_io_v2.ExtAuthPolicyClient

func ExtAuthPolicyClientFactoryProvider() ExtAuthPolicyClientFactory {
	return ExtAuthPolicyClientProvider
}

type ExtAuthPolicyClientFromConfigFactory func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.ExtAuthPolicyClient, error)

func ExtAuthPolicyClientFromConfigFactoryProvider() ExtAuthPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.ExtAuthPolicyClient, error) {
		clients, err := security_policy_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ExtAuthPolicies(), nil
	}
}

// Provider for WAFPolicyClient from Clientset
func WAFPolicyClientFromClientsetProvider(clients security_policy_gloo_solo_io_v2.Clientset) security_policy_gloo_solo_io_v2.WAFPolicyClient {
	return clients.WAFPolicies()
}

// Provider for WAFPolicy Client from Client
func WAFPolicyClientProvider(client client.Client) security_policy_gloo_solo_io_v2.WAFPolicyClient {
	return security_policy_gloo_solo_io_v2.NewWAFPolicyClient(client)
}

type WAFPolicyClientFactory func(client client.Client) security_policy_gloo_solo_io_v2.WAFPolicyClient

func WAFPolicyClientFactoryProvider() WAFPolicyClientFactory {
	return WAFPolicyClientProvider
}

type WAFPolicyClientFromConfigFactory func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.WAFPolicyClient, error)

func WAFPolicyClientFromConfigFactoryProvider() WAFPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.WAFPolicyClient, error) {
		clients, err := security_policy_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.WAFPolicies(), nil
	}
}

// Provider for JWTPolicyClient from Clientset
func JWTPolicyClientFromClientsetProvider(clients security_policy_gloo_solo_io_v2.Clientset) security_policy_gloo_solo_io_v2.JWTPolicyClient {
	return clients.JWTPolicies()
}

// Provider for JWTPolicy Client from Client
func JWTPolicyClientProvider(client client.Client) security_policy_gloo_solo_io_v2.JWTPolicyClient {
	return security_policy_gloo_solo_io_v2.NewJWTPolicyClient(client)
}

type JWTPolicyClientFactory func(client client.Client) security_policy_gloo_solo_io_v2.JWTPolicyClient

func JWTPolicyClientFactoryProvider() JWTPolicyClientFactory {
	return JWTPolicyClientProvider
}

type JWTPolicyClientFromConfigFactory func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.JWTPolicyClient, error)

func JWTPolicyClientFromConfigFactoryProvider() JWTPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.JWTPolicyClient, error) {
		clients, err := security_policy_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.JWTPolicies(), nil
	}
}

// Provider for ClientTLSPolicyClient from Clientset
func ClientTLSPolicyClientFromClientsetProvider(clients security_policy_gloo_solo_io_v2.Clientset) security_policy_gloo_solo_io_v2.ClientTLSPolicyClient {
	return clients.ClientTLSPolicies()
}

// Provider for ClientTLSPolicy Client from Client
func ClientTLSPolicyClientProvider(client client.Client) security_policy_gloo_solo_io_v2.ClientTLSPolicyClient {
	return security_policy_gloo_solo_io_v2.NewClientTLSPolicyClient(client)
}

type ClientTLSPolicyClientFactory func(client client.Client) security_policy_gloo_solo_io_v2.ClientTLSPolicyClient

func ClientTLSPolicyClientFactoryProvider() ClientTLSPolicyClientFactory {
	return ClientTLSPolicyClientProvider
}

type ClientTLSPolicyClientFromConfigFactory func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.ClientTLSPolicyClient, error)

func ClientTLSPolicyClientFromConfigFactoryProvider() ClientTLSPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.ClientTLSPolicyClient, error) {
		clients, err := security_policy_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ClientTLSPolicies(), nil
	}
}

// Provider for GraphQLAllowedQueryPolicyClient from Clientset
func GraphQLAllowedQueryPolicyClientFromClientsetProvider(clients security_policy_gloo_solo_io_v2.Clientset) security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicyClient {
	return clients.GraphQLAllowedQueryPolicies()
}

// Provider for GraphQLAllowedQueryPolicy Client from Client
func GraphQLAllowedQueryPolicyClientProvider(client client.Client) security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicyClient {
	return security_policy_gloo_solo_io_v2.NewGraphQLAllowedQueryPolicyClient(client)
}

type GraphQLAllowedQueryPolicyClientFactory func(client client.Client) security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicyClient

func GraphQLAllowedQueryPolicyClientFactoryProvider() GraphQLAllowedQueryPolicyClientFactory {
	return GraphQLAllowedQueryPolicyClientProvider
}

type GraphQLAllowedQueryPolicyClientFromConfigFactory func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicyClient, error)

func GraphQLAllowedQueryPolicyClientFromConfigFactoryProvider() GraphQLAllowedQueryPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicyClient, error) {
		clients, err := security_policy_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.GraphQLAllowedQueryPolicies(), nil
	}
}

// Provider for DLPPolicyClient from Clientset
func DLPPolicyClientFromClientsetProvider(clients security_policy_gloo_solo_io_v2.Clientset) security_policy_gloo_solo_io_v2.DLPPolicyClient {
	return clients.DLPPolicies()
}

// Provider for DLPPolicy Client from Client
func DLPPolicyClientProvider(client client.Client) security_policy_gloo_solo_io_v2.DLPPolicyClient {
	return security_policy_gloo_solo_io_v2.NewDLPPolicyClient(client)
}

type DLPPolicyClientFactory func(client client.Client) security_policy_gloo_solo_io_v2.DLPPolicyClient

func DLPPolicyClientFactoryProvider() DLPPolicyClientFactory {
	return DLPPolicyClientProvider
}

type DLPPolicyClientFromConfigFactory func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.DLPPolicyClient, error)

func DLPPolicyClientFromConfigFactoryProvider() DLPPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_policy_gloo_solo_io_v2.DLPPolicyClient, error) {
		clients, err := security_policy_gloo_solo_io_v2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.DLPPolicies(), nil
	}
}
