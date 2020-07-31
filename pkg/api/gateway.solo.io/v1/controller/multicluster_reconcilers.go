// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	gateway_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Gateway Resource across clusters.
// implemented by the user
type MulticlusterGatewayReconciler interface {
	ReconcileGateway(clusterName string, obj *gateway_solo_io_v1.Gateway) (reconcile.Result, error)
}

// Reconcile deletion events for the Gateway Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGatewayDeletionReconciler interface {
	ReconcileGatewayDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGatewayReconcilerFuncs struct {
	OnReconcileGateway         func(clusterName string, obj *gateway_solo_io_v1.Gateway) (reconcile.Result, error)
	OnReconcileGatewayDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGatewayReconcilerFuncs) ReconcileGateway(clusterName string, obj *gateway_solo_io_v1.Gateway) (reconcile.Result, error) {
	if f.OnReconcileGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGateway(clusterName, obj)
}

func (f *MulticlusterGatewayReconcilerFuncs) ReconcileGatewayDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayDeletion(clusterName, req)
}

type MulticlusterGatewayReconcileLoop interface {
	// AddMulticlusterGatewayReconciler adds a MulticlusterGatewayReconciler to the MulticlusterGatewayReconcileLoop.
	AddMulticlusterGatewayReconciler(ctx context.Context, rec MulticlusterGatewayReconciler, predicates ...predicate.Predicate)
}

type multiclusterGatewayReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGatewayReconcileLoop) AddMulticlusterGatewayReconciler(ctx context.Context, rec MulticlusterGatewayReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGatewayMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGatewayReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterGatewayReconcileLoop {
	return &multiclusterGatewayReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gateway_solo_io_v1.Gateway{})}
}

type genericGatewayMulticlusterReconciler struct {
	reconciler MulticlusterGatewayReconciler
}

func (g genericGatewayMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayDeletion(cluster, req)
	}
	return nil
}

func (g genericGatewayMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gateway_solo_io_v1.Gateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return g.reconciler.ReconcileGateway(cluster, obj)
}

// Reconcile Upsert events for the RouteTable Resource across clusters.
// implemented by the user
type MulticlusterRouteTableReconciler interface {
	ReconcileRouteTable(clusterName string, obj *gateway_solo_io_v1.RouteTable) (reconcile.Result, error)
}

// Reconcile deletion events for the RouteTable Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterRouteTableDeletionReconciler interface {
	ReconcileRouteTableDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterRouteTableReconcilerFuncs struct {
	OnReconcileRouteTable         func(clusterName string, obj *gateway_solo_io_v1.RouteTable) (reconcile.Result, error)
	OnReconcileRouteTableDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterRouteTableReconcilerFuncs) ReconcileRouteTable(clusterName string, obj *gateway_solo_io_v1.RouteTable) (reconcile.Result, error) {
	if f.OnReconcileRouteTable == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileRouteTable(clusterName, obj)
}

func (f *MulticlusterRouteTableReconcilerFuncs) ReconcileRouteTableDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileRouteTableDeletion == nil {
		return nil
	}
	return f.OnReconcileRouteTableDeletion(clusterName, req)
}

type MulticlusterRouteTableReconcileLoop interface {
	// AddMulticlusterRouteTableReconciler adds a MulticlusterRouteTableReconciler to the MulticlusterRouteTableReconcileLoop.
	AddMulticlusterRouteTableReconciler(ctx context.Context, rec MulticlusterRouteTableReconciler, predicates ...predicate.Predicate)
}

type multiclusterRouteTableReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterRouteTableReconcileLoop) AddMulticlusterRouteTableReconciler(ctx context.Context, rec MulticlusterRouteTableReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericRouteTableMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterRouteTableReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterRouteTableReconcileLoop {
	return &multiclusterRouteTableReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gateway_solo_io_v1.RouteTable{})}
}

type genericRouteTableMulticlusterReconciler struct {
	reconciler MulticlusterRouteTableReconciler
}

func (g genericRouteTableMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterRouteTableDeletionReconciler); ok {
		return deletionReconciler.ReconcileRouteTableDeletion(cluster, req)
	}
	return nil
}

func (g genericRouteTableMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gateway_solo_io_v1.RouteTable)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return g.reconciler.ReconcileRouteTable(cluster, obj)
}

// Reconcile Upsert events for the VirtualService Resource across clusters.
// implemented by the user
type MulticlusterVirtualServiceReconciler interface {
	ReconcileVirtualService(clusterName string, obj *gateway_solo_io_v1.VirtualService) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualService Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualServiceDeletionReconciler interface {
	ReconcileVirtualServiceDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterVirtualServiceReconcilerFuncs struct {
	OnReconcileVirtualService         func(clusterName string, obj *gateway_solo_io_v1.VirtualService) (reconcile.Result, error)
	OnReconcileVirtualServiceDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterVirtualServiceReconcilerFuncs) ReconcileVirtualService(clusterName string, obj *gateway_solo_io_v1.VirtualService) (reconcile.Result, error) {
	if f.OnReconcileVirtualService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualService(clusterName, obj)
}

func (f *MulticlusterVirtualServiceReconcilerFuncs) ReconcileVirtualServiceDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileVirtualServiceDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualServiceDeletion(clusterName, req)
}

type MulticlusterVirtualServiceReconcileLoop interface {
	// AddMulticlusterVirtualServiceReconciler adds a MulticlusterVirtualServiceReconciler to the MulticlusterVirtualServiceReconcileLoop.
	AddMulticlusterVirtualServiceReconciler(ctx context.Context, rec MulticlusterVirtualServiceReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualServiceReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualServiceReconcileLoop) AddMulticlusterVirtualServiceReconciler(ctx context.Context, rec MulticlusterVirtualServiceReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualServiceMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualServiceReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterVirtualServiceReconcileLoop {
	return &multiclusterVirtualServiceReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gateway_solo_io_v1.VirtualService{})}
}

type genericVirtualServiceMulticlusterReconciler struct {
	reconciler MulticlusterVirtualServiceReconciler
}

func (g genericVirtualServiceMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualServiceDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualServiceDeletion(cluster, req)
	}
	return nil
}

func (g genericVirtualServiceMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gateway_solo_io_v1.VirtualService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualService(cluster, obj)
}
