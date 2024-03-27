// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller



import (
	"context"

	gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Settings Resource across clusters.
// implemented by the user
type MulticlusterSettingsReconciler interface {
	ReconcileSettings(clusterName string, obj *gloo_solo_io_v1.Settings) (reconcile.Result, error)
}

// Reconcile deletion events for the Settings Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterSettingsDeletionReconciler interface {
	ReconcileSettingsDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterSettingsReconcilerFuncs struct {
	OnReconcileSettings         func(clusterName string, obj *gloo_solo_io_v1.Settings) (reconcile.Result, error)
	OnReconcileSettingsDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterSettingsReconcilerFuncs) ReconcileSettings(clusterName string, obj *gloo_solo_io_v1.Settings) (reconcile.Result, error) {
	if f.OnReconcileSettings == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileSettings(clusterName, obj)
}

func (f *MulticlusterSettingsReconcilerFuncs) ReconcileSettingsDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileSettingsDeletion == nil {
		return nil
	}
	return f.OnReconcileSettingsDeletion(clusterName, req)
}

type MulticlusterSettingsReconcileLoop interface {
	// AddMulticlusterSettingsReconciler adds a MulticlusterSettingsReconciler to the MulticlusterSettingsReconcileLoop.
	AddMulticlusterSettingsReconciler(ctx context.Context, rec MulticlusterSettingsReconciler, predicates ...predicate.Predicate)
}

type multiclusterSettingsReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterSettingsReconcileLoop) AddMulticlusterSettingsReconciler(ctx context.Context, rec MulticlusterSettingsReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericSettingsMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterSettingsReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterSettingsReconcileLoop {
	return &multiclusterSettingsReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gloo_solo_io_v1.Settings{}, options)}
}

type genericSettingsMulticlusterReconciler struct {
	reconciler MulticlusterSettingsReconciler
}

func (g genericSettingsMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterSettingsDeletionReconciler); ok {
		return deletionReconciler.ReconcileSettingsDeletion(cluster, req)
	}
	return nil
}

func (g genericSettingsMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gloo_solo_io_v1.Settings)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Settings handler received event for %T", object)
	}
	return g.reconciler.ReconcileSettings(cluster, obj)
}

// Reconcile Upsert events for the Upstream Resource across clusters.
// implemented by the user
type MulticlusterUpstreamReconciler interface {
	ReconcileUpstream(clusterName string, obj *gloo_solo_io_v1.Upstream) (reconcile.Result, error)
}

// Reconcile deletion events for the Upstream Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterUpstreamDeletionReconciler interface {
	ReconcileUpstreamDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterUpstreamReconcilerFuncs struct {
	OnReconcileUpstream         func(clusterName string, obj *gloo_solo_io_v1.Upstream) (reconcile.Result, error)
	OnReconcileUpstreamDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterUpstreamReconcilerFuncs) ReconcileUpstream(clusterName string, obj *gloo_solo_io_v1.Upstream) (reconcile.Result, error) {
	if f.OnReconcileUpstream == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileUpstream(clusterName, obj)
}

func (f *MulticlusterUpstreamReconcilerFuncs) ReconcileUpstreamDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileUpstreamDeletion == nil {
		return nil
	}
	return f.OnReconcileUpstreamDeletion(clusterName, req)
}

type MulticlusterUpstreamReconcileLoop interface {
	// AddMulticlusterUpstreamReconciler adds a MulticlusterUpstreamReconciler to the MulticlusterUpstreamReconcileLoop.
	AddMulticlusterUpstreamReconciler(ctx context.Context, rec MulticlusterUpstreamReconciler, predicates ...predicate.Predicate)
}

type multiclusterUpstreamReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterUpstreamReconcileLoop) AddMulticlusterUpstreamReconciler(ctx context.Context, rec MulticlusterUpstreamReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericUpstreamMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterUpstreamReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterUpstreamReconcileLoop {
	return &multiclusterUpstreamReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gloo_solo_io_v1.Upstream{}, options)}
}

type genericUpstreamMulticlusterReconciler struct {
	reconciler MulticlusterUpstreamReconciler
}

func (g genericUpstreamMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterUpstreamDeletionReconciler); ok {
		return deletionReconciler.ReconcileUpstreamDeletion(cluster, req)
	}
	return nil
}

func (g genericUpstreamMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gloo_solo_io_v1.Upstream)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Upstream handler received event for %T", object)
	}
	return g.reconciler.ReconcileUpstream(cluster, obj)
}

// Reconcile Upsert events for the UpstreamGroup Resource across clusters.
// implemented by the user
type MulticlusterUpstreamGroupReconciler interface {
	ReconcileUpstreamGroup(clusterName string, obj *gloo_solo_io_v1.UpstreamGroup) (reconcile.Result, error)
}

// Reconcile deletion events for the UpstreamGroup Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterUpstreamGroupDeletionReconciler interface {
	ReconcileUpstreamGroupDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterUpstreamGroupReconcilerFuncs struct {
	OnReconcileUpstreamGroup         func(clusterName string, obj *gloo_solo_io_v1.UpstreamGroup) (reconcile.Result, error)
	OnReconcileUpstreamGroupDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterUpstreamGroupReconcilerFuncs) ReconcileUpstreamGroup(clusterName string, obj *gloo_solo_io_v1.UpstreamGroup) (reconcile.Result, error) {
	if f.OnReconcileUpstreamGroup == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileUpstreamGroup(clusterName, obj)
}

func (f *MulticlusterUpstreamGroupReconcilerFuncs) ReconcileUpstreamGroupDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileUpstreamGroupDeletion == nil {
		return nil
	}
	return f.OnReconcileUpstreamGroupDeletion(clusterName, req)
}

type MulticlusterUpstreamGroupReconcileLoop interface {
	// AddMulticlusterUpstreamGroupReconciler adds a MulticlusterUpstreamGroupReconciler to the MulticlusterUpstreamGroupReconcileLoop.
	AddMulticlusterUpstreamGroupReconciler(ctx context.Context, rec MulticlusterUpstreamGroupReconciler, predicates ...predicate.Predicate)
}

type multiclusterUpstreamGroupReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterUpstreamGroupReconcileLoop) AddMulticlusterUpstreamGroupReconciler(ctx context.Context, rec MulticlusterUpstreamGroupReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericUpstreamGroupMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterUpstreamGroupReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterUpstreamGroupReconcileLoop {
	return &multiclusterUpstreamGroupReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gloo_solo_io_v1.UpstreamGroup{}, options)}
}

type genericUpstreamGroupMulticlusterReconciler struct {
	reconciler MulticlusterUpstreamGroupReconciler
}

func (g genericUpstreamGroupMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterUpstreamGroupDeletionReconciler); ok {
		return deletionReconciler.ReconcileUpstreamGroupDeletion(cluster, req)
	}
	return nil
}

func (g genericUpstreamGroupMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gloo_solo_io_v1.UpstreamGroup)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: UpstreamGroup handler received event for %T", object)
	}
	return g.reconciler.ReconcileUpstreamGroup(cluster, obj)
}

// Reconcile Upsert events for the Proxy Resource across clusters.
// implemented by the user
type MulticlusterProxyReconciler interface {
	ReconcileProxy(clusterName string, obj *gloo_solo_io_v1.Proxy) (reconcile.Result, error)
}

// Reconcile deletion events for the Proxy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterProxyDeletionReconciler interface {
	ReconcileProxyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterProxyReconcilerFuncs struct {
	OnReconcileProxy         func(clusterName string, obj *gloo_solo_io_v1.Proxy) (reconcile.Result, error)
	OnReconcileProxyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterProxyReconcilerFuncs) ReconcileProxy(clusterName string, obj *gloo_solo_io_v1.Proxy) (reconcile.Result, error) {
	if f.OnReconcileProxy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileProxy(clusterName, obj)
}

func (f *MulticlusterProxyReconcilerFuncs) ReconcileProxyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileProxyDeletion == nil {
		return nil
	}
	return f.OnReconcileProxyDeletion(clusterName, req)
}

type MulticlusterProxyReconcileLoop interface {
	// AddMulticlusterProxyReconciler adds a MulticlusterProxyReconciler to the MulticlusterProxyReconcileLoop.
	AddMulticlusterProxyReconciler(ctx context.Context, rec MulticlusterProxyReconciler, predicates ...predicate.Predicate)
}

type multiclusterProxyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterProxyReconcileLoop) AddMulticlusterProxyReconciler(ctx context.Context, rec MulticlusterProxyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericProxyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterProxyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterProxyReconcileLoop {
	return &multiclusterProxyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gloo_solo_io_v1.Proxy{}, options)}
}

type genericProxyMulticlusterReconciler struct {
	reconciler MulticlusterProxyReconciler
}

func (g genericProxyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterProxyDeletionReconciler); ok {
		return deletionReconciler.ReconcileProxyDeletion(cluster, req)
	}
	return nil
}

func (g genericProxyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gloo_solo_io_v1.Proxy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Proxy handler received event for %T", object)
	}
	return g.reconciler.ReconcileProxy(cluster, obj)
}
