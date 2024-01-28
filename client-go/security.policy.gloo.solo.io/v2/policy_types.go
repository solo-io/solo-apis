// Code generated by skv2. DO NOT EDIT.

// Policy methods for Gloo Mesh policy types.
package v2

import (
	commonv2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

// IsPolicy implements Policy interface for AccessPolicy
func (o *AccessPolicy) IsPolicy() {}

// GetDestinationSelectors returns the destination selectors of the AccessPolicy policy
func (o *AccessPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// GetNamespaceWorkloadSelectors returns the namespace workload selectors of the AccessPolicy policy
func (o *AccessPolicy) GetNamespaceWorkloadSelectors() []*AccessPolicySpec_NamespaceWorkloadSelector {
	return o.Spec.ApplyToWorkloads
}

// IsPolicy implements Policy interface for CORSPolicy
func (o *CORSPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the CORSPolicy policy
func (o *CORSPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// IsPolicy implements Policy interface for CSRFPolicy
func (o *CSRFPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the CSRFPolicy policy
func (o *CSRFPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// IsPolicy implements Policy interface for ExtAuthPolicy
func (o *ExtAuthPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the ExtAuthPolicy policy
func (o *ExtAuthPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// GetDestinationSelectors returns the destination selectors of the ExtAuthPolicy policy
func (o *ExtAuthPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for WAFPolicy
func (o *WAFPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the WAFPolicy policy
func (o *WAFPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// IsPolicy implements Policy interface for JWTPolicy
func (o *JWTPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the JWTPolicy policy
func (o *JWTPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// GetDestinationSelectors returns the destination selectors of the JWTPolicy policy
func (o *JWTPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for ClientTLSPolicy
func (o *ClientTLSPolicy) IsPolicy() {}

// GetDestinationSelectors returns the destination selectors of the ClientTLSPolicy policy
func (o *ClientTLSPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for GraphQLAllowedQueryPolicy
func (o *GraphQLAllowedQueryPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the GraphQLAllowedQueryPolicy policy
func (o *GraphQLAllowedQueryPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// IsPolicy implements Policy interface for DLPPolicy
func (o *DLPPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the DLPPolicy policy
func (o *DLPPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}