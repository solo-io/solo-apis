// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	infrastructure_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/infrastructure.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the CloudProvider Resource across clusters.
// implemented by the user
type MulticlusterCloudProviderReconciler interface {
	ReconcileCloudProvider(clusterName string, obj *infrastructure_gloo_solo_io_v2.CloudProvider) (reconcile.Result, error)
}

// Reconcile deletion events for the CloudProvider Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterCloudProviderDeletionReconciler interface {
	ReconcileCloudProviderDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterCloudProviderReconcilerFuncs struct {
	OnReconcileCloudProvider         func(clusterName string, obj *infrastructure_gloo_solo_io_v2.CloudProvider) (reconcile.Result, error)
	OnReconcileCloudProviderDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterCloudProviderReconcilerFuncs) ReconcileCloudProvider(clusterName string, obj *infrastructure_gloo_solo_io_v2.CloudProvider) (reconcile.Result, error) {
	if f.OnReconcileCloudProvider == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCloudProvider(clusterName, obj)
}

func (f *MulticlusterCloudProviderReconcilerFuncs) ReconcileCloudProviderDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileCloudProviderDeletion == nil {
		return nil
	}
	return f.OnReconcileCloudProviderDeletion(clusterName, req)
}

type MulticlusterCloudProviderReconcileLoop interface {
	// AddMulticlusterCloudProviderReconciler adds a MulticlusterCloudProviderReconciler to the MulticlusterCloudProviderReconcileLoop.
	AddMulticlusterCloudProviderReconciler(ctx context.Context, rec MulticlusterCloudProviderReconciler, predicates ...predicate.Predicate)
}

type multiclusterCloudProviderReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterCloudProviderReconcileLoop) AddMulticlusterCloudProviderReconciler(ctx context.Context, rec MulticlusterCloudProviderReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericCloudProviderMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterCloudProviderReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterCloudProviderReconcileLoop {
	return &multiclusterCloudProviderReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &infrastructure_gloo_solo_io_v2.CloudProvider{}, options)}
}

type genericCloudProviderMulticlusterReconciler struct {
	reconciler MulticlusterCloudProviderReconciler
}

func (g genericCloudProviderMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterCloudProviderDeletionReconciler); ok {
		return deletionReconciler.ReconcileCloudProviderDeletion(cluster, req)
	}
	return nil
}

func (g genericCloudProviderMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*infrastructure_gloo_solo_io_v2.CloudProvider)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CloudProvider handler received event for %T", object)
	}
	return g.reconciler.ReconcileCloudProvider(cluster, obj)
}

// Reconcile Upsert events for the CloudResources Resource across clusters.
// implemented by the user
type MulticlusterCloudResourcesReconciler interface {
	ReconcileCloudResources(clusterName string, obj *infrastructure_gloo_solo_io_v2.CloudResources) (reconcile.Result, error)
}

// Reconcile deletion events for the CloudResources Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterCloudResourcesDeletionReconciler interface {
	ReconcileCloudResourcesDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterCloudResourcesReconcilerFuncs struct {
	OnReconcileCloudResources         func(clusterName string, obj *infrastructure_gloo_solo_io_v2.CloudResources) (reconcile.Result, error)
	OnReconcileCloudResourcesDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterCloudResourcesReconcilerFuncs) ReconcileCloudResources(clusterName string, obj *infrastructure_gloo_solo_io_v2.CloudResources) (reconcile.Result, error) {
	if f.OnReconcileCloudResources == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCloudResources(clusterName, obj)
}

func (f *MulticlusterCloudResourcesReconcilerFuncs) ReconcileCloudResourcesDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileCloudResourcesDeletion == nil {
		return nil
	}
	return f.OnReconcileCloudResourcesDeletion(clusterName, req)
}

type MulticlusterCloudResourcesReconcileLoop interface {
	// AddMulticlusterCloudResourcesReconciler adds a MulticlusterCloudResourcesReconciler to the MulticlusterCloudResourcesReconcileLoop.
	AddMulticlusterCloudResourcesReconciler(ctx context.Context, rec MulticlusterCloudResourcesReconciler, predicates ...predicate.Predicate)
}

type multiclusterCloudResourcesReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterCloudResourcesReconcileLoop) AddMulticlusterCloudResourcesReconciler(ctx context.Context, rec MulticlusterCloudResourcesReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericCloudResourcesMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterCloudResourcesReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterCloudResourcesReconcileLoop {
	return &multiclusterCloudResourcesReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &infrastructure_gloo_solo_io_v2.CloudResources{}, options)}
}

type genericCloudResourcesMulticlusterReconciler struct {
	reconciler MulticlusterCloudResourcesReconciler
}

func (g genericCloudResourcesMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterCloudResourcesDeletionReconciler); ok {
		return deletionReconciler.ReconcileCloudResourcesDeletion(cluster, req)
	}
	return nil
}

func (g genericCloudResourcesMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*infrastructure_gloo_solo_io_v2.CloudResources)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CloudResources handler received event for %T", object)
	}
	return g.reconciler.ReconcileCloudResources(cluster, obj)
}
