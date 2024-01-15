// Code generated by skv2. DO NOT EDIT.

// Policy methods for Gloo Mesh policy types.
package v2

import (
	commonv2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

// IsPolicy implements Policy interface for MirrorPolicy
func (o *MirrorPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the MirrorPolicy policy
func (o *MirrorPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// IsPolicy implements Policy interface for RateLimitPolicy
func (o *RateLimitPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the RateLimitPolicy policy
func (o *RateLimitPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// GetDestinationSelectors returns the destination selectors of the RateLimitPolicy policy
func (o *RateLimitPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for RateLimitClientConfig
func (o *RateLimitClientConfig) IsPolicy() {}

// IsPolicy implements Policy interface for HeaderManipulationPolicy
func (o *HeaderManipulationPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the HeaderManipulationPolicy policy
func (o *HeaderManipulationPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// GetRouteDestinationSelectors returns the route destination selectors of the HeaderManipulationPolicy policy
func (o *HeaderManipulationPolicy) GetRouteDestinationSelectors() []*commonv2.RouteDestinationSelector {
	return o.Spec.ApplyToRouteDestinations
}

// IsPolicy implements Policy interface for TransformationPolicy
func (o *TransformationPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the TransformationPolicy policy
func (o *TransformationPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// IsPolicy implements Policy interface for LoadBalancerPolicy
func (o *LoadBalancerPolicy) IsPolicy() {}

// GetDestinationSelectors returns the destination selectors of the LoadBalancerPolicy policy
func (o *LoadBalancerPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for ProxyProtocolPolicy
func (o *ProxyProtocolPolicy) IsPolicy() {}

// GetApplyToListeners returns the listener selectors of the ProxyProtocolPolicy policy
func (o *ProxyProtocolPolicy) GetApplyToListeners() []*commonv2.ListenerSelector {
	return o.Spec.ApplyToListeners
}

// IsPolicy implements Policy interface for HTTPBufferPolicy
func (o *HTTPBufferPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the HTTPBufferPolicy policy
func (o *HTTPBufferPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}
