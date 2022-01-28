// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	enterprise_gloo_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the GraphQLSchema Resource across clusters.
// implemented by the user
type MulticlusterGraphQLSchemaReconciler interface {
	ReconcileGraphQLSchema(clusterName string, obj *enterprise_gloo_solo_io_v1alpha1.GraphQLSchema) (reconcile.Result, error)
}

// Reconcile deletion events for the GraphQLSchema Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGraphQLSchemaDeletionReconciler interface {
	ReconcileGraphQLSchemaDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGraphQLSchemaReconcilerFuncs struct {
	OnReconcileGraphQLSchema         func(clusterName string, obj *enterprise_gloo_solo_io_v1alpha1.GraphQLSchema) (reconcile.Result, error)
	OnReconcileGraphQLSchemaDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGraphQLSchemaReconcilerFuncs) ReconcileGraphQLSchema(clusterName string, obj *enterprise_gloo_solo_io_v1alpha1.GraphQLSchema) (reconcile.Result, error) {
	if f.OnReconcileGraphQLSchema == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGraphQLSchema(clusterName, obj)
}

func (f *MulticlusterGraphQLSchemaReconcilerFuncs) ReconcileGraphQLSchemaDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGraphQLSchemaDeletion == nil {
		return nil
	}
	return f.OnReconcileGraphQLSchemaDeletion(clusterName, req)
}

type MulticlusterGraphQLSchemaReconcileLoop interface {
	// AddMulticlusterGraphQLSchemaReconciler adds a MulticlusterGraphQLSchemaReconciler to the MulticlusterGraphQLSchemaReconcileLoop.
	AddMulticlusterGraphQLSchemaReconciler(ctx context.Context, rec MulticlusterGraphQLSchemaReconciler, predicates ...predicate.Predicate)
}

type multiclusterGraphQLSchemaReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGraphQLSchemaReconcileLoop) AddMulticlusterGraphQLSchemaReconciler(ctx context.Context, rec MulticlusterGraphQLSchemaReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGraphQLSchemaMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGraphQLSchemaReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterGraphQLSchemaReconcileLoop {
	return &multiclusterGraphQLSchemaReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &enterprise_gloo_solo_io_v1alpha1.GraphQLSchema{}, options)}
}

type genericGraphQLSchemaMulticlusterReconciler struct {
	reconciler MulticlusterGraphQLSchemaReconciler
}

func (g genericGraphQLSchemaMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGraphQLSchemaDeletionReconciler); ok {
		return deletionReconciler.ReconcileGraphQLSchemaDeletion(cluster, req)
	}
	return nil
}

func (g genericGraphQLSchemaMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*enterprise_gloo_solo_io_v1alpha1.GraphQLSchema)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GraphQLSchema handler received event for %T", object)
	}
	return g.reconciler.ReconcileGraphQLSchema(cluster, obj)
}
