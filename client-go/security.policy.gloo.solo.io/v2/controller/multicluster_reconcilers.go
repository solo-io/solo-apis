// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	security_policy_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/security.policy.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the AccessPolicy Resource across clusters.
// implemented by the user
type MulticlusterAccessPolicyReconciler interface {
	ReconcileAccessPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.AccessPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the AccessPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterAccessPolicyDeletionReconciler interface {
	ReconcileAccessPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterAccessPolicyReconcilerFuncs struct {
	OnReconcileAccessPolicy         func(clusterName string, obj *security_policy_gloo_solo_io_v2.AccessPolicy) (reconcile.Result, error)
	OnReconcileAccessPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterAccessPolicyReconcilerFuncs) ReconcileAccessPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.AccessPolicy) (reconcile.Result, error) {
	if f.OnReconcileAccessPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAccessPolicy(clusterName, obj)
}

func (f *MulticlusterAccessPolicyReconcilerFuncs) ReconcileAccessPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileAccessPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileAccessPolicyDeletion(clusterName, req)
}

type MulticlusterAccessPolicyReconcileLoop interface {
	// AddMulticlusterAccessPolicyReconciler adds a MulticlusterAccessPolicyReconciler to the MulticlusterAccessPolicyReconcileLoop.
	AddMulticlusterAccessPolicyReconciler(ctx context.Context, rec MulticlusterAccessPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterAccessPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterAccessPolicyReconcileLoop) AddMulticlusterAccessPolicyReconciler(ctx context.Context, rec MulticlusterAccessPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericAccessPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterAccessPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterAccessPolicyReconcileLoop {
	return &multiclusterAccessPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_policy_gloo_solo_io_v2.AccessPolicy{}, options)}
}

type genericAccessPolicyMulticlusterReconciler struct {
	reconciler MulticlusterAccessPolicyReconciler
}

func (g genericAccessPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterAccessPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileAccessPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericAccessPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.AccessPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileAccessPolicy(cluster, obj)
}

// Reconcile Upsert events for the CORSPolicy Resource across clusters.
// implemented by the user
type MulticlusterCORSPolicyReconciler interface {
	ReconcileCORSPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.CORSPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the CORSPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterCORSPolicyDeletionReconciler interface {
	ReconcileCORSPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterCORSPolicyReconcilerFuncs struct {
	OnReconcileCORSPolicy         func(clusterName string, obj *security_policy_gloo_solo_io_v2.CORSPolicy) (reconcile.Result, error)
	OnReconcileCORSPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterCORSPolicyReconcilerFuncs) ReconcileCORSPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.CORSPolicy) (reconcile.Result, error) {
	if f.OnReconcileCORSPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCORSPolicy(clusterName, obj)
}

func (f *MulticlusterCORSPolicyReconcilerFuncs) ReconcileCORSPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileCORSPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileCORSPolicyDeletion(clusterName, req)
}

type MulticlusterCORSPolicyReconcileLoop interface {
	// AddMulticlusterCORSPolicyReconciler adds a MulticlusterCORSPolicyReconciler to the MulticlusterCORSPolicyReconcileLoop.
	AddMulticlusterCORSPolicyReconciler(ctx context.Context, rec MulticlusterCORSPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterCORSPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterCORSPolicyReconcileLoop) AddMulticlusterCORSPolicyReconciler(ctx context.Context, rec MulticlusterCORSPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericCORSPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterCORSPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterCORSPolicyReconcileLoop {
	return &multiclusterCORSPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_policy_gloo_solo_io_v2.CORSPolicy{}, options)}
}

type genericCORSPolicyMulticlusterReconciler struct {
	reconciler MulticlusterCORSPolicyReconciler
}

func (g genericCORSPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterCORSPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileCORSPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericCORSPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CORSPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CORSPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileCORSPolicy(cluster, obj)
}

// Reconcile Upsert events for the CSRFPolicy Resource across clusters.
// implemented by the user
type MulticlusterCSRFPolicyReconciler interface {
	ReconcileCSRFPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.CSRFPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the CSRFPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterCSRFPolicyDeletionReconciler interface {
	ReconcileCSRFPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterCSRFPolicyReconcilerFuncs struct {
	OnReconcileCSRFPolicy         func(clusterName string, obj *security_policy_gloo_solo_io_v2.CSRFPolicy) (reconcile.Result, error)
	OnReconcileCSRFPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterCSRFPolicyReconcilerFuncs) ReconcileCSRFPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.CSRFPolicy) (reconcile.Result, error) {
	if f.OnReconcileCSRFPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCSRFPolicy(clusterName, obj)
}

func (f *MulticlusterCSRFPolicyReconcilerFuncs) ReconcileCSRFPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileCSRFPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileCSRFPolicyDeletion(clusterName, req)
}

type MulticlusterCSRFPolicyReconcileLoop interface {
	// AddMulticlusterCSRFPolicyReconciler adds a MulticlusterCSRFPolicyReconciler to the MulticlusterCSRFPolicyReconcileLoop.
	AddMulticlusterCSRFPolicyReconciler(ctx context.Context, rec MulticlusterCSRFPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterCSRFPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterCSRFPolicyReconcileLoop) AddMulticlusterCSRFPolicyReconciler(ctx context.Context, rec MulticlusterCSRFPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericCSRFPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterCSRFPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterCSRFPolicyReconcileLoop {
	return &multiclusterCSRFPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_policy_gloo_solo_io_v2.CSRFPolicy{}, options)}
}

type genericCSRFPolicyMulticlusterReconciler struct {
	reconciler MulticlusterCSRFPolicyReconciler
}

func (g genericCSRFPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterCSRFPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileCSRFPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericCSRFPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CSRFPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CSRFPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileCSRFPolicy(cluster, obj)
}

// Reconcile Upsert events for the ExtAuthPolicy Resource across clusters.
// implemented by the user
type MulticlusterExtAuthPolicyReconciler interface {
	ReconcileExtAuthPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the ExtAuthPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterExtAuthPolicyDeletionReconciler interface {
	ReconcileExtAuthPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterExtAuthPolicyReconcilerFuncs struct {
	OnReconcileExtAuthPolicy         func(clusterName string, obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) (reconcile.Result, error)
	OnReconcileExtAuthPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterExtAuthPolicyReconcilerFuncs) ReconcileExtAuthPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) (reconcile.Result, error) {
	if f.OnReconcileExtAuthPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileExtAuthPolicy(clusterName, obj)
}

func (f *MulticlusterExtAuthPolicyReconcilerFuncs) ReconcileExtAuthPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileExtAuthPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileExtAuthPolicyDeletion(clusterName, req)
}

type MulticlusterExtAuthPolicyReconcileLoop interface {
	// AddMulticlusterExtAuthPolicyReconciler adds a MulticlusterExtAuthPolicyReconciler to the MulticlusterExtAuthPolicyReconcileLoop.
	AddMulticlusterExtAuthPolicyReconciler(ctx context.Context, rec MulticlusterExtAuthPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterExtAuthPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterExtAuthPolicyReconcileLoop) AddMulticlusterExtAuthPolicyReconciler(ctx context.Context, rec MulticlusterExtAuthPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericExtAuthPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterExtAuthPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterExtAuthPolicyReconcileLoop {
	return &multiclusterExtAuthPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_policy_gloo_solo_io_v2.ExtAuthPolicy{}, options)}
}

type genericExtAuthPolicyMulticlusterReconciler struct {
	reconciler MulticlusterExtAuthPolicyReconciler
}

func (g genericExtAuthPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterExtAuthPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileExtAuthPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericExtAuthPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ExtAuthPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ExtAuthPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileExtAuthPolicy(cluster, obj)
}

// Reconcile Upsert events for the WAFPolicy Resource across clusters.
// implemented by the user
type MulticlusterWAFPolicyReconciler interface {
	ReconcileWAFPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.WAFPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the WAFPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterWAFPolicyDeletionReconciler interface {
	ReconcileWAFPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterWAFPolicyReconcilerFuncs struct {
	OnReconcileWAFPolicy         func(clusterName string, obj *security_policy_gloo_solo_io_v2.WAFPolicy) (reconcile.Result, error)
	OnReconcileWAFPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterWAFPolicyReconcilerFuncs) ReconcileWAFPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.WAFPolicy) (reconcile.Result, error) {
	if f.OnReconcileWAFPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWAFPolicy(clusterName, obj)
}

func (f *MulticlusterWAFPolicyReconcilerFuncs) ReconcileWAFPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileWAFPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileWAFPolicyDeletion(clusterName, req)
}

type MulticlusterWAFPolicyReconcileLoop interface {
	// AddMulticlusterWAFPolicyReconciler adds a MulticlusterWAFPolicyReconciler to the MulticlusterWAFPolicyReconcileLoop.
	AddMulticlusterWAFPolicyReconciler(ctx context.Context, rec MulticlusterWAFPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterWAFPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterWAFPolicyReconcileLoop) AddMulticlusterWAFPolicyReconciler(ctx context.Context, rec MulticlusterWAFPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericWAFPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterWAFPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterWAFPolicyReconcileLoop {
	return &multiclusterWAFPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_policy_gloo_solo_io_v2.WAFPolicy{}, options)}
}

type genericWAFPolicyMulticlusterReconciler struct {
	reconciler MulticlusterWAFPolicyReconciler
}

func (g genericWAFPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterWAFPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileWAFPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericWAFPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.WAFPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WAFPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileWAFPolicy(cluster, obj)
}

// Reconcile Upsert events for the JWTPolicy Resource across clusters.
// implemented by the user
type MulticlusterJWTPolicyReconciler interface {
	ReconcileJWTPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.JWTPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the JWTPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterJWTPolicyDeletionReconciler interface {
	ReconcileJWTPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterJWTPolicyReconcilerFuncs struct {
	OnReconcileJWTPolicy         func(clusterName string, obj *security_policy_gloo_solo_io_v2.JWTPolicy) (reconcile.Result, error)
	OnReconcileJWTPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterJWTPolicyReconcilerFuncs) ReconcileJWTPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.JWTPolicy) (reconcile.Result, error) {
	if f.OnReconcileJWTPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileJWTPolicy(clusterName, obj)
}

func (f *MulticlusterJWTPolicyReconcilerFuncs) ReconcileJWTPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileJWTPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileJWTPolicyDeletion(clusterName, req)
}

type MulticlusterJWTPolicyReconcileLoop interface {
	// AddMulticlusterJWTPolicyReconciler adds a MulticlusterJWTPolicyReconciler to the MulticlusterJWTPolicyReconcileLoop.
	AddMulticlusterJWTPolicyReconciler(ctx context.Context, rec MulticlusterJWTPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterJWTPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterJWTPolicyReconcileLoop) AddMulticlusterJWTPolicyReconciler(ctx context.Context, rec MulticlusterJWTPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericJWTPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterJWTPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterJWTPolicyReconcileLoop {
	return &multiclusterJWTPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_policy_gloo_solo_io_v2.JWTPolicy{}, options)}
}

type genericJWTPolicyMulticlusterReconciler struct {
	reconciler MulticlusterJWTPolicyReconciler
}

func (g genericJWTPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterJWTPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileJWTPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericJWTPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.JWTPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: JWTPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileJWTPolicy(cluster, obj)
}

// Reconcile Upsert events for the ClientTLSPolicy Resource across clusters.
// implemented by the user
type MulticlusterClientTLSPolicyReconciler interface {
	ReconcileClientTLSPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the ClientTLSPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterClientTLSPolicyDeletionReconciler interface {
	ReconcileClientTLSPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterClientTLSPolicyReconcilerFuncs struct {
	OnReconcileClientTLSPolicy         func(clusterName string, obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) (reconcile.Result, error)
	OnReconcileClientTLSPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterClientTLSPolicyReconcilerFuncs) ReconcileClientTLSPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) (reconcile.Result, error) {
	if f.OnReconcileClientTLSPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileClientTLSPolicy(clusterName, obj)
}

func (f *MulticlusterClientTLSPolicyReconcilerFuncs) ReconcileClientTLSPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileClientTLSPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileClientTLSPolicyDeletion(clusterName, req)
}

type MulticlusterClientTLSPolicyReconcileLoop interface {
	// AddMulticlusterClientTLSPolicyReconciler adds a MulticlusterClientTLSPolicyReconciler to the MulticlusterClientTLSPolicyReconcileLoop.
	AddMulticlusterClientTLSPolicyReconciler(ctx context.Context, rec MulticlusterClientTLSPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterClientTLSPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterClientTLSPolicyReconcileLoop) AddMulticlusterClientTLSPolicyReconciler(ctx context.Context, rec MulticlusterClientTLSPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericClientTLSPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterClientTLSPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterClientTLSPolicyReconcileLoop {
	return &multiclusterClientTLSPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_policy_gloo_solo_io_v2.ClientTLSPolicy{}, options)}
}

type genericClientTLSPolicyMulticlusterReconciler struct {
	reconciler MulticlusterClientTLSPolicyReconciler
}

func (g genericClientTLSPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterClientTLSPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileClientTLSPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericClientTLSPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ClientTLSPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ClientTLSPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileClientTLSPolicy(cluster, obj)
}

// Reconcile Upsert events for the GraphQLAllowedQueryPolicy Resource across clusters.
// implemented by the user
type MulticlusterGraphQLAllowedQueryPolicyReconciler interface {
	ReconcileGraphQLAllowedQueryPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the GraphQLAllowedQueryPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGraphQLAllowedQueryPolicyDeletionReconciler interface {
	ReconcileGraphQLAllowedQueryPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGraphQLAllowedQueryPolicyReconcilerFuncs struct {
	OnReconcileGraphQLAllowedQueryPolicy         func(clusterName string, obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) (reconcile.Result, error)
	OnReconcileGraphQLAllowedQueryPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGraphQLAllowedQueryPolicyReconcilerFuncs) ReconcileGraphQLAllowedQueryPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) (reconcile.Result, error) {
	if f.OnReconcileGraphQLAllowedQueryPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGraphQLAllowedQueryPolicy(clusterName, obj)
}

func (f *MulticlusterGraphQLAllowedQueryPolicyReconcilerFuncs) ReconcileGraphQLAllowedQueryPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGraphQLAllowedQueryPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileGraphQLAllowedQueryPolicyDeletion(clusterName, req)
}

type MulticlusterGraphQLAllowedQueryPolicyReconcileLoop interface {
	// AddMulticlusterGraphQLAllowedQueryPolicyReconciler adds a MulticlusterGraphQLAllowedQueryPolicyReconciler to the MulticlusterGraphQLAllowedQueryPolicyReconcileLoop.
	AddMulticlusterGraphQLAllowedQueryPolicyReconciler(ctx context.Context, rec MulticlusterGraphQLAllowedQueryPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterGraphQLAllowedQueryPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGraphQLAllowedQueryPolicyReconcileLoop) AddMulticlusterGraphQLAllowedQueryPolicyReconciler(ctx context.Context, rec MulticlusterGraphQLAllowedQueryPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGraphQLAllowedQueryPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGraphQLAllowedQueryPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterGraphQLAllowedQueryPolicyReconcileLoop {
	return &multiclusterGraphQLAllowedQueryPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy{}, options)}
}

type genericGraphQLAllowedQueryPolicyMulticlusterReconciler struct {
	reconciler MulticlusterGraphQLAllowedQueryPolicyReconciler
}

func (g genericGraphQLAllowedQueryPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGraphQLAllowedQueryPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileGraphQLAllowedQueryPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericGraphQLAllowedQueryPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GraphQLAllowedQueryPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileGraphQLAllowedQueryPolicy(cluster, obj)
}

// Reconcile Upsert events for the DLPPolicy Resource across clusters.
// implemented by the user
type MulticlusterDLPPolicyReconciler interface {
	ReconcileDLPPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.DLPPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the DLPPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterDLPPolicyDeletionReconciler interface {
	ReconcileDLPPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterDLPPolicyReconcilerFuncs struct {
	OnReconcileDLPPolicy         func(clusterName string, obj *security_policy_gloo_solo_io_v2.DLPPolicy) (reconcile.Result, error)
	OnReconcileDLPPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterDLPPolicyReconcilerFuncs) ReconcileDLPPolicy(clusterName string, obj *security_policy_gloo_solo_io_v2.DLPPolicy) (reconcile.Result, error) {
	if f.OnReconcileDLPPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileDLPPolicy(clusterName, obj)
}

func (f *MulticlusterDLPPolicyReconcilerFuncs) ReconcileDLPPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileDLPPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileDLPPolicyDeletion(clusterName, req)
}

type MulticlusterDLPPolicyReconcileLoop interface {
	// AddMulticlusterDLPPolicyReconciler adds a MulticlusterDLPPolicyReconciler to the MulticlusterDLPPolicyReconcileLoop.
	AddMulticlusterDLPPolicyReconciler(ctx context.Context, rec MulticlusterDLPPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterDLPPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterDLPPolicyReconcileLoop) AddMulticlusterDLPPolicyReconciler(ctx context.Context, rec MulticlusterDLPPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericDLPPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterDLPPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterDLPPolicyReconcileLoop {
	return &multiclusterDLPPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_policy_gloo_solo_io_v2.DLPPolicy{}, options)}
}

type genericDLPPolicyMulticlusterReconciler struct {
	reconciler MulticlusterDLPPolicyReconciler
}

func (g genericDLPPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterDLPPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileDLPPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericDLPPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.DLPPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: DLPPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileDLPPolicy(cluster, obj)
}