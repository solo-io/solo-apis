// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	extensions_solo_io_v2 "github.com/solo-io/solo-apis/pkg/api/extensions.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the WasmDeploymentPolicy Resource across clusters.
// implemented by the user
type MulticlusterWasmDeploymentPolicyReconciler interface {
	ReconcileWasmDeploymentPolicy(clusterName string, obj *extensions_solo_io_v2.WasmDeploymentPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the WasmDeploymentPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterWasmDeploymentPolicyDeletionReconciler interface {
	ReconcileWasmDeploymentPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterWasmDeploymentPolicyReconcilerFuncs struct {
	OnReconcileWasmDeploymentPolicy         func(clusterName string, obj *extensions_solo_io_v2.WasmDeploymentPolicy) (reconcile.Result, error)
	OnReconcileWasmDeploymentPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterWasmDeploymentPolicyReconcilerFuncs) ReconcileWasmDeploymentPolicy(clusterName string, obj *extensions_solo_io_v2.WasmDeploymentPolicy) (reconcile.Result, error) {
	if f.OnReconcileWasmDeploymentPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWasmDeploymentPolicy(clusterName, obj)
}

func (f *MulticlusterWasmDeploymentPolicyReconcilerFuncs) ReconcileWasmDeploymentPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileWasmDeploymentPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileWasmDeploymentPolicyDeletion(clusterName, req)
}

type MulticlusterWasmDeploymentPolicyReconcileLoop interface {
	// AddMulticlusterWasmDeploymentPolicyReconciler adds a MulticlusterWasmDeploymentPolicyReconciler to the MulticlusterWasmDeploymentPolicyReconcileLoop.
	AddMulticlusterWasmDeploymentPolicyReconciler(ctx context.Context, rec MulticlusterWasmDeploymentPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterWasmDeploymentPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterWasmDeploymentPolicyReconcileLoop) AddMulticlusterWasmDeploymentPolicyReconciler(ctx context.Context, rec MulticlusterWasmDeploymentPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericWasmDeploymentPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterWasmDeploymentPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterWasmDeploymentPolicyReconcileLoop {
	return &multiclusterWasmDeploymentPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &extensions_solo_io_v2.WasmDeploymentPolicy{}, options)}
}

type genericWasmDeploymentPolicyMulticlusterReconciler struct {
	reconciler MulticlusterWasmDeploymentPolicyReconciler
}

func (g genericWasmDeploymentPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterWasmDeploymentPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileWasmDeploymentPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericWasmDeploymentPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*extensions_solo_io_v2.WasmDeploymentPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WasmDeploymentPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileWasmDeploymentPolicy(cluster, obj)
}
