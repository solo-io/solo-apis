// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	ratelimit_solo_io_v1alpha1 "github.com/solo-io/gloo-fed/pkg/api/ratelimit.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the RateLimitConfig Resource across clusters.
// implemented by the user
type MulticlusterRateLimitConfigReconciler interface {
	ReconcileRateLimitConfig(clusterName string, obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) (reconcile.Result, error)
}

// Reconcile deletion events for the RateLimitConfig Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterRateLimitConfigDeletionReconciler interface {
	ReconcileRateLimitConfigDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterRateLimitConfigReconcilerFuncs struct {
	OnReconcileRateLimitConfig         func(clusterName string, obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) (reconcile.Result, error)
	OnReconcileRateLimitConfigDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterRateLimitConfigReconcilerFuncs) ReconcileRateLimitConfig(clusterName string, obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) (reconcile.Result, error) {
	if f.OnReconcileRateLimitConfig == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileRateLimitConfig(clusterName, obj)
}

func (f *MulticlusterRateLimitConfigReconcilerFuncs) ReconcileRateLimitConfigDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileRateLimitConfigDeletion == nil {
		return nil
	}
	return f.OnReconcileRateLimitConfigDeletion(clusterName, req)
}

type MulticlusterRateLimitConfigReconcileLoop interface {
	// AddMulticlusterRateLimitConfigReconciler adds a MulticlusterRateLimitConfigReconciler to the MulticlusterRateLimitConfigReconcileLoop.
	AddMulticlusterRateLimitConfigReconciler(ctx context.Context, rec MulticlusterRateLimitConfigReconciler, predicates ...predicate.Predicate)
}

type multiclusterRateLimitConfigReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterRateLimitConfigReconcileLoop) AddMulticlusterRateLimitConfigReconciler(ctx context.Context, rec MulticlusterRateLimitConfigReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericRateLimitConfigMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterRateLimitConfigReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterRateLimitConfigReconcileLoop {
	return &multiclusterRateLimitConfigReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &ratelimit_solo_io_v1alpha1.RateLimitConfig{})}
}

type genericRateLimitConfigMulticlusterReconciler struct {
	reconciler MulticlusterRateLimitConfigReconciler
}

func (g genericRateLimitConfigMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterRateLimitConfigDeletionReconciler); ok {
		return deletionReconciler.ReconcileRateLimitConfigDeletion(cluster, req)
	}
	return nil
}

func (g genericRateLimitConfigMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*ratelimit_solo_io_v1alpha1.RateLimitConfig)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: RateLimitConfig handler received event for %T", object)
	}
	return g.reconciler.ReconcileRateLimitConfig(cluster, obj)
}
