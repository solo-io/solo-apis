// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	fed_gateway_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/fed.gateway.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the FederatedGateway Resource across clusters.
// implemented by the user
type MulticlusterFederatedGatewayReconciler interface {
	ReconcileFederatedGateway(clusterName string, obj *fed_gateway_solo_io_v1.FederatedGateway) (reconcile.Result, error)
}

// Reconcile deletion events for the FederatedGateway Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFederatedGatewayDeletionReconciler interface {
	ReconcileFederatedGatewayDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterFederatedGatewayReconcilerFuncs struct {
	OnReconcileFederatedGateway         func(clusterName string, obj *fed_gateway_solo_io_v1.FederatedGateway) (reconcile.Result, error)
	OnReconcileFederatedGatewayDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterFederatedGatewayReconcilerFuncs) ReconcileFederatedGateway(clusterName string, obj *fed_gateway_solo_io_v1.FederatedGateway) (reconcile.Result, error) {
	if f.OnReconcileFederatedGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFederatedGateway(clusterName, obj)
}

func (f *MulticlusterFederatedGatewayReconcilerFuncs) ReconcileFederatedGatewayDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileFederatedGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileFederatedGatewayDeletion(clusterName, req)
}

type MulticlusterFederatedGatewayReconcileLoop interface {
	// AddMulticlusterFederatedGatewayReconciler adds a MulticlusterFederatedGatewayReconciler to the MulticlusterFederatedGatewayReconcileLoop.
	AddMulticlusterFederatedGatewayReconciler(ctx context.Context, rec MulticlusterFederatedGatewayReconciler, predicates ...predicate.Predicate)
}

type multiclusterFederatedGatewayReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFederatedGatewayReconcileLoop) AddMulticlusterFederatedGatewayReconciler(ctx context.Context, rec MulticlusterFederatedGatewayReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFederatedGatewayMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFederatedGatewayReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterFederatedGatewayReconcileLoop {
	return &multiclusterFederatedGatewayReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &fed_gateway_solo_io_v1.FederatedGateway{}, options)}
}

type genericFederatedGatewayMulticlusterReconciler struct {
	reconciler MulticlusterFederatedGatewayReconciler
}

func (g genericFederatedGatewayMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFederatedGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileFederatedGatewayDeletion(cluster, req)
	}
	return nil
}

func (g genericFederatedGatewayMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedGateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FederatedGateway handler received event for %T", object)
	}
	return g.reconciler.ReconcileFederatedGateway(cluster, obj)
}

// Reconcile Upsert events for the FederatedMatchableHttpGateway Resource across clusters.
// implemented by the user
type MulticlusterFederatedMatchableHttpGatewayReconciler interface {
	ReconcileFederatedMatchableHttpGateway(clusterName string, obj *fed_gateway_solo_io_v1.FederatedMatchableHttpGateway) (reconcile.Result, error)
}

// Reconcile deletion events for the FederatedMatchableHttpGateway Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFederatedMatchableHttpGatewayDeletionReconciler interface {
	ReconcileFederatedMatchableHttpGatewayDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterFederatedMatchableHttpGatewayReconcilerFuncs struct {
	OnReconcileFederatedMatchableHttpGateway         func(clusterName string, obj *fed_gateway_solo_io_v1.FederatedMatchableHttpGateway) (reconcile.Result, error)
	OnReconcileFederatedMatchableHttpGatewayDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterFederatedMatchableHttpGatewayReconcilerFuncs) ReconcileFederatedMatchableHttpGateway(clusterName string, obj *fed_gateway_solo_io_v1.FederatedMatchableHttpGateway) (reconcile.Result, error) {
	if f.OnReconcileFederatedMatchableHttpGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFederatedMatchableHttpGateway(clusterName, obj)
}

func (f *MulticlusterFederatedMatchableHttpGatewayReconcilerFuncs) ReconcileFederatedMatchableHttpGatewayDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileFederatedMatchableHttpGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileFederatedMatchableHttpGatewayDeletion(clusterName, req)
}

type MulticlusterFederatedMatchableHttpGatewayReconcileLoop interface {
	// AddMulticlusterFederatedMatchableHttpGatewayReconciler adds a MulticlusterFederatedMatchableHttpGatewayReconciler to the MulticlusterFederatedMatchableHttpGatewayReconcileLoop.
	AddMulticlusterFederatedMatchableHttpGatewayReconciler(ctx context.Context, rec MulticlusterFederatedMatchableHttpGatewayReconciler, predicates ...predicate.Predicate)
}

type multiclusterFederatedMatchableHttpGatewayReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFederatedMatchableHttpGatewayReconcileLoop) AddMulticlusterFederatedMatchableHttpGatewayReconciler(ctx context.Context, rec MulticlusterFederatedMatchableHttpGatewayReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFederatedMatchableHttpGatewayMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFederatedMatchableHttpGatewayReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterFederatedMatchableHttpGatewayReconcileLoop {
	return &multiclusterFederatedMatchableHttpGatewayReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &fed_gateway_solo_io_v1.FederatedMatchableHttpGateway{}, options)}
}

type genericFederatedMatchableHttpGatewayMulticlusterReconciler struct {
	reconciler MulticlusterFederatedMatchableHttpGatewayReconciler
}

func (g genericFederatedMatchableHttpGatewayMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFederatedMatchableHttpGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileFederatedMatchableHttpGatewayDeletion(cluster, req)
	}
	return nil
}

func (g genericFederatedMatchableHttpGatewayMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedMatchableHttpGateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FederatedMatchableHttpGateway handler received event for %T", object)
	}
	return g.reconciler.ReconcileFederatedMatchableHttpGateway(cluster, obj)
}

// Reconcile Upsert events for the FederatedMatchableTcpGateway Resource across clusters.
// implemented by the user
type MulticlusterFederatedMatchableTcpGatewayReconciler interface {
	ReconcileFederatedMatchableTcpGateway(clusterName string, obj *fed_gateway_solo_io_v1.FederatedMatchableTcpGateway) (reconcile.Result, error)
}

// Reconcile deletion events for the FederatedMatchableTcpGateway Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFederatedMatchableTcpGatewayDeletionReconciler interface {
	ReconcileFederatedMatchableTcpGatewayDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterFederatedMatchableTcpGatewayReconcilerFuncs struct {
	OnReconcileFederatedMatchableTcpGateway         func(clusterName string, obj *fed_gateway_solo_io_v1.FederatedMatchableTcpGateway) (reconcile.Result, error)
	OnReconcileFederatedMatchableTcpGatewayDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterFederatedMatchableTcpGatewayReconcilerFuncs) ReconcileFederatedMatchableTcpGateway(clusterName string, obj *fed_gateway_solo_io_v1.FederatedMatchableTcpGateway) (reconcile.Result, error) {
	if f.OnReconcileFederatedMatchableTcpGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFederatedMatchableTcpGateway(clusterName, obj)
}

func (f *MulticlusterFederatedMatchableTcpGatewayReconcilerFuncs) ReconcileFederatedMatchableTcpGatewayDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileFederatedMatchableTcpGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileFederatedMatchableTcpGatewayDeletion(clusterName, req)
}

type MulticlusterFederatedMatchableTcpGatewayReconcileLoop interface {
	// AddMulticlusterFederatedMatchableTcpGatewayReconciler adds a MulticlusterFederatedMatchableTcpGatewayReconciler to the MulticlusterFederatedMatchableTcpGatewayReconcileLoop.
	AddMulticlusterFederatedMatchableTcpGatewayReconciler(ctx context.Context, rec MulticlusterFederatedMatchableTcpGatewayReconciler, predicates ...predicate.Predicate)
}

type multiclusterFederatedMatchableTcpGatewayReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFederatedMatchableTcpGatewayReconcileLoop) AddMulticlusterFederatedMatchableTcpGatewayReconciler(ctx context.Context, rec MulticlusterFederatedMatchableTcpGatewayReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFederatedMatchableTcpGatewayMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFederatedMatchableTcpGatewayReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterFederatedMatchableTcpGatewayReconcileLoop {
	return &multiclusterFederatedMatchableTcpGatewayReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &fed_gateway_solo_io_v1.FederatedMatchableTcpGateway{}, options)}
}

type genericFederatedMatchableTcpGatewayMulticlusterReconciler struct {
	reconciler MulticlusterFederatedMatchableTcpGatewayReconciler
}

func (g genericFederatedMatchableTcpGatewayMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFederatedMatchableTcpGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileFederatedMatchableTcpGatewayDeletion(cluster, req)
	}
	return nil
}

func (g genericFederatedMatchableTcpGatewayMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedMatchableTcpGateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FederatedMatchableTcpGateway handler received event for %T", object)
	}
	return g.reconciler.ReconcileFederatedMatchableTcpGateway(cluster, obj)
}

// Reconcile Upsert events for the FederatedRouteTable Resource across clusters.
// implemented by the user
type MulticlusterFederatedRouteTableReconciler interface {
	ReconcileFederatedRouteTable(clusterName string, obj *fed_gateway_solo_io_v1.FederatedRouteTable) (reconcile.Result, error)
}

// Reconcile deletion events for the FederatedRouteTable Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFederatedRouteTableDeletionReconciler interface {
	ReconcileFederatedRouteTableDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterFederatedRouteTableReconcilerFuncs struct {
	OnReconcileFederatedRouteTable         func(clusterName string, obj *fed_gateway_solo_io_v1.FederatedRouteTable) (reconcile.Result, error)
	OnReconcileFederatedRouteTableDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterFederatedRouteTableReconcilerFuncs) ReconcileFederatedRouteTable(clusterName string, obj *fed_gateway_solo_io_v1.FederatedRouteTable) (reconcile.Result, error) {
	if f.OnReconcileFederatedRouteTable == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFederatedRouteTable(clusterName, obj)
}

func (f *MulticlusterFederatedRouteTableReconcilerFuncs) ReconcileFederatedRouteTableDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileFederatedRouteTableDeletion == nil {
		return nil
	}
	return f.OnReconcileFederatedRouteTableDeletion(clusterName, req)
}

type MulticlusterFederatedRouteTableReconcileLoop interface {
	// AddMulticlusterFederatedRouteTableReconciler adds a MulticlusterFederatedRouteTableReconciler to the MulticlusterFederatedRouteTableReconcileLoop.
	AddMulticlusterFederatedRouteTableReconciler(ctx context.Context, rec MulticlusterFederatedRouteTableReconciler, predicates ...predicate.Predicate)
}

type multiclusterFederatedRouteTableReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFederatedRouteTableReconcileLoop) AddMulticlusterFederatedRouteTableReconciler(ctx context.Context, rec MulticlusterFederatedRouteTableReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFederatedRouteTableMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFederatedRouteTableReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterFederatedRouteTableReconcileLoop {
	return &multiclusterFederatedRouteTableReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &fed_gateway_solo_io_v1.FederatedRouteTable{}, options)}
}

type genericFederatedRouteTableMulticlusterReconciler struct {
	reconciler MulticlusterFederatedRouteTableReconciler
}

func (g genericFederatedRouteTableMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFederatedRouteTableDeletionReconciler); ok {
		return deletionReconciler.ReconcileFederatedRouteTableDeletion(cluster, req)
	}
	return nil
}

func (g genericFederatedRouteTableMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedRouteTable)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FederatedRouteTable handler received event for %T", object)
	}
	return g.reconciler.ReconcileFederatedRouteTable(cluster, obj)
}

// Reconcile Upsert events for the FederatedVirtualService Resource across clusters.
// implemented by the user
type MulticlusterFederatedVirtualServiceReconciler interface {
	ReconcileFederatedVirtualService(clusterName string, obj *fed_gateway_solo_io_v1.FederatedVirtualService) (reconcile.Result, error)
}

// Reconcile deletion events for the FederatedVirtualService Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFederatedVirtualServiceDeletionReconciler interface {
	ReconcileFederatedVirtualServiceDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterFederatedVirtualServiceReconcilerFuncs struct {
	OnReconcileFederatedVirtualService         func(clusterName string, obj *fed_gateway_solo_io_v1.FederatedVirtualService) (reconcile.Result, error)
	OnReconcileFederatedVirtualServiceDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterFederatedVirtualServiceReconcilerFuncs) ReconcileFederatedVirtualService(clusterName string, obj *fed_gateway_solo_io_v1.FederatedVirtualService) (reconcile.Result, error) {
	if f.OnReconcileFederatedVirtualService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFederatedVirtualService(clusterName, obj)
}

func (f *MulticlusterFederatedVirtualServiceReconcilerFuncs) ReconcileFederatedVirtualServiceDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileFederatedVirtualServiceDeletion == nil {
		return nil
	}
	return f.OnReconcileFederatedVirtualServiceDeletion(clusterName, req)
}

type MulticlusterFederatedVirtualServiceReconcileLoop interface {
	// AddMulticlusterFederatedVirtualServiceReconciler adds a MulticlusterFederatedVirtualServiceReconciler to the MulticlusterFederatedVirtualServiceReconcileLoop.
	AddMulticlusterFederatedVirtualServiceReconciler(ctx context.Context, rec MulticlusterFederatedVirtualServiceReconciler, predicates ...predicate.Predicate)
}

type multiclusterFederatedVirtualServiceReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFederatedVirtualServiceReconcileLoop) AddMulticlusterFederatedVirtualServiceReconciler(ctx context.Context, rec MulticlusterFederatedVirtualServiceReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFederatedVirtualServiceMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFederatedVirtualServiceReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterFederatedVirtualServiceReconcileLoop {
	return &multiclusterFederatedVirtualServiceReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &fed_gateway_solo_io_v1.FederatedVirtualService{}, options)}
}

type genericFederatedVirtualServiceMulticlusterReconciler struct {
	reconciler MulticlusterFederatedVirtualServiceReconciler
}

func (g genericFederatedVirtualServiceMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFederatedVirtualServiceDeletionReconciler); ok {
		return deletionReconciler.ReconcileFederatedVirtualServiceDeletion(cluster, req)
	}
	return nil
}

func (g genericFederatedVirtualServiceMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedVirtualService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FederatedVirtualService handler received event for %T", object)
	}
	return g.reconciler.ReconcileFederatedVirtualService(cluster, obj)
}
