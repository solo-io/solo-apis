// Code generated by skv2. DO NOT EDIT.

// Policy methods for Gloo Mesh policy types.
package v2

import (
	commonv2 "github.com/solo-io/solo-apis/client-go/common.gloo.solo.io/v2"
)

// IsPolicy implements Policy interface for GraphQLPersistedQueryCachePolicy
func (o *GraphQLPersistedQueryCachePolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the GraphQLPersistedQueryCachePolicy policy
func (o *GraphQLPersistedQueryCachePolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// IsPolicy implements Policy interface for FailoverPolicy
func (o *FailoverPolicy) IsPolicy() {}

// GetDestinationSelectors returns the destination selectors of the FailoverPolicy policy
func (o *FailoverPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for OutlierDetectionPolicy
func (o *OutlierDetectionPolicy) IsPolicy() {}

// GetDestinationSelectors returns the destination selectors of the OutlierDetectionPolicy policy
func (o *OutlierDetectionPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for AdaptiveRequestConcurrencyPolicy
func (o *AdaptiveRequestConcurrencyPolicy) IsPolicy() {}

// GetDestinationSelectors returns the destination selectors of the AdaptiveRequestConcurrencyPolicy policy
func (o *AdaptiveRequestConcurrencyPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for FaultInjectionPolicy
func (o *FaultInjectionPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the FaultInjectionPolicy policy
func (o *FaultInjectionPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// IsPolicy implements Policy interface for RetryTimeoutPolicy
func (o *RetryTimeoutPolicy) IsPolicy() {}

// GetRouteSelectors returns the route selectors of the RetryTimeoutPolicy policy
func (o *RetryTimeoutPolicy) GetRouteSelectors() []*commonv2.RouteSelector {
	return o.Spec.ApplyToRoutes
}

// IsPolicy implements Policy interface for ConnectionPolicy
func (o *ConnectionPolicy) IsPolicy() {}

// GetDestinationSelectors returns the destination selectors of the ConnectionPolicy policy
func (o *ConnectionPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for TrimProxyConfigPolicy
func (o *TrimProxyConfigPolicy) IsPolicy() {}

// GetWorkloadSelectors returns the workload selectors of the TrimProxyConfigPolicy policy
func (o *TrimProxyConfigPolicy) GetWorkloadSelectors() []*commonv2.WorkloadSelector {
	return o.Spec.ApplyToWorkloads
}

// IsPolicy implements Policy interface for ActiveHealthCheckPolicy
func (o *ActiveHealthCheckPolicy) IsPolicy() {}

// GetDestinationSelectors returns the destination selectors of the ActiveHealthCheckPolicy policy
func (o *ActiveHealthCheckPolicy) GetDestinationSelectors() []*commonv2.DestinationSelector {
	return o.Spec.ApplyToDestinations
}

// IsPolicy implements Policy interface for ListenerConnectionPolicy
func (o *ListenerConnectionPolicy) IsPolicy() {}

// GetApplyToListeners returns the listener selectors of the ListenerConnectionPolicy policy
func (o *ListenerConnectionPolicy) GetApplyToListeners() []*commonv2.ListenerSelector {
	return o.Spec.ApplyToListeners
}
