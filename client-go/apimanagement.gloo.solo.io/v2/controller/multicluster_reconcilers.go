// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	apimanagement_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/apimanagement.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the GraphQLStitchedSchema Resource across clusters.
// implemented by the user
type MulticlusterGraphQLStitchedSchemaReconciler interface {
	ReconcileGraphQLStitchedSchema(clusterName string, obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) (reconcile.Result, error)
}

// Reconcile deletion events for the GraphQLStitchedSchema Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGraphQLStitchedSchemaDeletionReconciler interface {
	ReconcileGraphQLStitchedSchemaDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGraphQLStitchedSchemaReconcilerFuncs struct {
	OnReconcileGraphQLStitchedSchema         func(clusterName string, obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) (reconcile.Result, error)
	OnReconcileGraphQLStitchedSchemaDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGraphQLStitchedSchemaReconcilerFuncs) ReconcileGraphQLStitchedSchema(clusterName string, obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) (reconcile.Result, error) {
	if f.OnReconcileGraphQLStitchedSchema == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGraphQLStitchedSchema(clusterName, obj)
}

func (f *MulticlusterGraphQLStitchedSchemaReconcilerFuncs) ReconcileGraphQLStitchedSchemaDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGraphQLStitchedSchemaDeletion == nil {
		return nil
	}
	return f.OnReconcileGraphQLStitchedSchemaDeletion(clusterName, req)
}

type MulticlusterGraphQLStitchedSchemaReconcileLoop interface {
	// AddMulticlusterGraphQLStitchedSchemaReconciler adds a MulticlusterGraphQLStitchedSchemaReconciler to the MulticlusterGraphQLStitchedSchemaReconcileLoop.
	AddMulticlusterGraphQLStitchedSchemaReconciler(ctx context.Context, rec MulticlusterGraphQLStitchedSchemaReconciler, predicates ...predicate.Predicate)
}

type multiclusterGraphQLStitchedSchemaReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGraphQLStitchedSchemaReconcileLoop) AddMulticlusterGraphQLStitchedSchemaReconciler(ctx context.Context, rec MulticlusterGraphQLStitchedSchemaReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGraphQLStitchedSchemaMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGraphQLStitchedSchemaReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterGraphQLStitchedSchemaReconcileLoop {
	return &multiclusterGraphQLStitchedSchemaReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema{}, options)}
}

type genericGraphQLStitchedSchemaMulticlusterReconciler struct {
	reconciler MulticlusterGraphQLStitchedSchemaReconciler
}

func (g genericGraphQLStitchedSchemaMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGraphQLStitchedSchemaDeletionReconciler); ok {
		return deletionReconciler.ReconcileGraphQLStitchedSchemaDeletion(cluster, req)
	}
	return nil
}

func (g genericGraphQLStitchedSchemaMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GraphQLStitchedSchema handler received event for %T", object)
	}
	return g.reconciler.ReconcileGraphQLStitchedSchema(cluster, obj)
}

// Reconcile Upsert events for the GraphQLResolverMap Resource across clusters.
// implemented by the user
type MulticlusterGraphQLResolverMapReconciler interface {
	ReconcileGraphQLResolverMap(clusterName string, obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) (reconcile.Result, error)
}

// Reconcile deletion events for the GraphQLResolverMap Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGraphQLResolverMapDeletionReconciler interface {
	ReconcileGraphQLResolverMapDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGraphQLResolverMapReconcilerFuncs struct {
	OnReconcileGraphQLResolverMap         func(clusterName string, obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) (reconcile.Result, error)
	OnReconcileGraphQLResolverMapDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGraphQLResolverMapReconcilerFuncs) ReconcileGraphQLResolverMap(clusterName string, obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) (reconcile.Result, error) {
	if f.OnReconcileGraphQLResolverMap == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGraphQLResolverMap(clusterName, obj)
}

func (f *MulticlusterGraphQLResolverMapReconcilerFuncs) ReconcileGraphQLResolverMapDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGraphQLResolverMapDeletion == nil {
		return nil
	}
	return f.OnReconcileGraphQLResolverMapDeletion(clusterName, req)
}

type MulticlusterGraphQLResolverMapReconcileLoop interface {
	// AddMulticlusterGraphQLResolverMapReconciler adds a MulticlusterGraphQLResolverMapReconciler to the MulticlusterGraphQLResolverMapReconcileLoop.
	AddMulticlusterGraphQLResolverMapReconciler(ctx context.Context, rec MulticlusterGraphQLResolverMapReconciler, predicates ...predicate.Predicate)
}

type multiclusterGraphQLResolverMapReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGraphQLResolverMapReconcileLoop) AddMulticlusterGraphQLResolverMapReconciler(ctx context.Context, rec MulticlusterGraphQLResolverMapReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGraphQLResolverMapMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGraphQLResolverMapReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterGraphQLResolverMapReconcileLoop {
	return &multiclusterGraphQLResolverMapReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &apimanagement_gloo_solo_io_v2.GraphQLResolverMap{}, options)}
}

type genericGraphQLResolverMapMulticlusterReconciler struct {
	reconciler MulticlusterGraphQLResolverMapReconciler
}

func (g genericGraphQLResolverMapMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGraphQLResolverMapDeletionReconciler); ok {
		return deletionReconciler.ReconcileGraphQLResolverMapDeletion(cluster, req)
	}
	return nil
}

func (g genericGraphQLResolverMapMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLResolverMap)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GraphQLResolverMap handler received event for %T", object)
	}
	return g.reconciler.ReconcileGraphQLResolverMap(cluster, obj)
}

// Reconcile Upsert events for the GraphQLSchema Resource across clusters.
// implemented by the user
type MulticlusterGraphQLSchemaReconciler interface {
	ReconcileGraphQLSchema(clusterName string, obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) (reconcile.Result, error)
}

// Reconcile deletion events for the GraphQLSchema Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGraphQLSchemaDeletionReconciler interface {
	ReconcileGraphQLSchemaDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGraphQLSchemaReconcilerFuncs struct {
	OnReconcileGraphQLSchema         func(clusterName string, obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) (reconcile.Result, error)
	OnReconcileGraphQLSchemaDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGraphQLSchemaReconcilerFuncs) ReconcileGraphQLSchema(clusterName string, obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) (reconcile.Result, error) {
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
	return &multiclusterGraphQLSchemaReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &apimanagement_gloo_solo_io_v2.GraphQLSchema{}, options)}
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
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLSchema)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GraphQLSchema handler received event for %T", object)
	}
	return g.reconciler.ReconcileGraphQLSchema(cluster, obj)
}

// Reconcile Upsert events for the ApiDoc Resource across clusters.
// implemented by the user
type MulticlusterApiDocReconciler interface {
	ReconcileApiDoc(clusterName string, obj *apimanagement_gloo_solo_io_v2.ApiDoc) (reconcile.Result, error)
}

// Reconcile deletion events for the ApiDoc Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterApiDocDeletionReconciler interface {
	ReconcileApiDocDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterApiDocReconcilerFuncs struct {
	OnReconcileApiDoc         func(clusterName string, obj *apimanagement_gloo_solo_io_v2.ApiDoc) (reconcile.Result, error)
	OnReconcileApiDocDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterApiDocReconcilerFuncs) ReconcileApiDoc(clusterName string, obj *apimanagement_gloo_solo_io_v2.ApiDoc) (reconcile.Result, error) {
	if f.OnReconcileApiDoc == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileApiDoc(clusterName, obj)
}

func (f *MulticlusterApiDocReconcilerFuncs) ReconcileApiDocDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileApiDocDeletion == nil {
		return nil
	}
	return f.OnReconcileApiDocDeletion(clusterName, req)
}

type MulticlusterApiDocReconcileLoop interface {
	// AddMulticlusterApiDocReconciler adds a MulticlusterApiDocReconciler to the MulticlusterApiDocReconcileLoop.
	AddMulticlusterApiDocReconciler(ctx context.Context, rec MulticlusterApiDocReconciler, predicates ...predicate.Predicate)
}

type multiclusterApiDocReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterApiDocReconcileLoop) AddMulticlusterApiDocReconciler(ctx context.Context, rec MulticlusterApiDocReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericApiDocMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterApiDocReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterApiDocReconcileLoop {
	return &multiclusterApiDocReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &apimanagement_gloo_solo_io_v2.ApiDoc{}, options)}
}

type genericApiDocMulticlusterReconciler struct {
	reconciler MulticlusterApiDocReconciler
}

func (g genericApiDocMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterApiDocDeletionReconciler); ok {
		return deletionReconciler.ReconcileApiDocDeletion(cluster, req)
	}
	return nil
}

func (g genericApiDocMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.ApiDoc)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ApiDoc handler received event for %T", object)
	}
	return g.reconciler.ReconcileApiDoc(cluster, obj)
}
