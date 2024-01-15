// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	admin_gloo_solo_io_v2alpha1 "github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the WaypointLifecycleManager Resource.
// implemented by the user
type WaypointLifecycleManagerReconciler interface {
	ReconcileWaypointLifecycleManager(obj *admin_gloo_solo_io_v2alpha1.WaypointLifecycleManager) (reconcile.Result, error)
}

// Reconcile deletion events for the WaypointLifecycleManager Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type WaypointLifecycleManagerDeletionReconciler interface {
	ReconcileWaypointLifecycleManagerDeletion(req reconcile.Request) error
}

type WaypointLifecycleManagerReconcilerFuncs struct {
	OnReconcileWaypointLifecycleManager         func(obj *admin_gloo_solo_io_v2alpha1.WaypointLifecycleManager) (reconcile.Result, error)
	OnReconcileWaypointLifecycleManagerDeletion func(req reconcile.Request) error
}

func (f *WaypointLifecycleManagerReconcilerFuncs) ReconcileWaypointLifecycleManager(obj *admin_gloo_solo_io_v2alpha1.WaypointLifecycleManager) (reconcile.Result, error) {
	if f.OnReconcileWaypointLifecycleManager == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWaypointLifecycleManager(obj)
}

func (f *WaypointLifecycleManagerReconcilerFuncs) ReconcileWaypointLifecycleManagerDeletion(req reconcile.Request) error {
	if f.OnReconcileWaypointLifecycleManagerDeletion == nil {
		return nil
	}
	return f.OnReconcileWaypointLifecycleManagerDeletion(req)
}

// Reconcile and finalize the WaypointLifecycleManager Resource
// implemented by the user
type WaypointLifecycleManagerFinalizer interface {
	WaypointLifecycleManagerReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	WaypointLifecycleManagerFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeWaypointLifecycleManager(obj *admin_gloo_solo_io_v2alpha1.WaypointLifecycleManager) error
}

type WaypointLifecycleManagerReconcileLoop interface {
	RunWaypointLifecycleManagerReconciler(ctx context.Context, rec WaypointLifecycleManagerReconciler, predicates ...predicate.Predicate) error
}

type waypointLifecycleManagerReconcileLoop struct {
	loop reconcile.Loop
}

func NewWaypointLifecycleManagerReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) WaypointLifecycleManagerReconcileLoop {
	return &waypointLifecycleManagerReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &admin_gloo_solo_io_v2alpha1.WaypointLifecycleManager{}, options),
	}
}

func (c *waypointLifecycleManagerReconcileLoop) RunWaypointLifecycleManagerReconciler(ctx context.Context, reconciler WaypointLifecycleManagerReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericWaypointLifecycleManagerReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(WaypointLifecycleManagerFinalizer); ok {
		reconcilerWrapper = genericWaypointLifecycleManagerFinalizer{
			genericWaypointLifecycleManagerReconciler: genericReconciler,
			finalizingReconciler:                      finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericWaypointLifecycleManagerHandler implements a generic reconcile.Reconciler
type genericWaypointLifecycleManagerReconciler struct {
	reconciler WaypointLifecycleManagerReconciler
}

func (r genericWaypointLifecycleManagerReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*admin_gloo_solo_io_v2alpha1.WaypointLifecycleManager)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WaypointLifecycleManager handler received event for %T", object)
	}
	return r.reconciler.ReconcileWaypointLifecycleManager(obj)
}

func (r genericWaypointLifecycleManagerReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(WaypointLifecycleManagerDeletionReconciler); ok {
		return deletionReconciler.ReconcileWaypointLifecycleManagerDeletion(request)
	}
	return nil
}

// genericWaypointLifecycleManagerFinalizer implements a generic reconcile.FinalizingReconciler
type genericWaypointLifecycleManagerFinalizer struct {
	genericWaypointLifecycleManagerReconciler
	finalizingReconciler WaypointLifecycleManagerFinalizer
}

func (r genericWaypointLifecycleManagerFinalizer) FinalizerName() string {
	return r.finalizingReconciler.WaypointLifecycleManagerFinalizerName()
}

func (r genericWaypointLifecycleManagerFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2alpha1.WaypointLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: WaypointLifecycleManager handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeWaypointLifecycleManager(obj)
}

// Reconcile Upsert events for the InsightsConfig Resource.
// implemented by the user
type InsightsConfigReconciler interface {
	ReconcileInsightsConfig(obj *admin_gloo_solo_io_v2alpha1.InsightsConfig) (reconcile.Result, error)
}

// Reconcile deletion events for the InsightsConfig Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type InsightsConfigDeletionReconciler interface {
	ReconcileInsightsConfigDeletion(req reconcile.Request) error
}

type InsightsConfigReconcilerFuncs struct {
	OnReconcileInsightsConfig         func(obj *admin_gloo_solo_io_v2alpha1.InsightsConfig) (reconcile.Result, error)
	OnReconcileInsightsConfigDeletion func(req reconcile.Request) error
}

func (f *InsightsConfigReconcilerFuncs) ReconcileInsightsConfig(obj *admin_gloo_solo_io_v2alpha1.InsightsConfig) (reconcile.Result, error) {
	if f.OnReconcileInsightsConfig == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileInsightsConfig(obj)
}

func (f *InsightsConfigReconcilerFuncs) ReconcileInsightsConfigDeletion(req reconcile.Request) error {
	if f.OnReconcileInsightsConfigDeletion == nil {
		return nil
	}
	return f.OnReconcileInsightsConfigDeletion(req)
}

// Reconcile and finalize the InsightsConfig Resource
// implemented by the user
type InsightsConfigFinalizer interface {
	InsightsConfigReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	InsightsConfigFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeInsightsConfig(obj *admin_gloo_solo_io_v2alpha1.InsightsConfig) error
}

type InsightsConfigReconcileLoop interface {
	RunInsightsConfigReconciler(ctx context.Context, rec InsightsConfigReconciler, predicates ...predicate.Predicate) error
}

type insightsConfigReconcileLoop struct {
	loop reconcile.Loop
}

func NewInsightsConfigReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) InsightsConfigReconcileLoop {
	return &insightsConfigReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &admin_gloo_solo_io_v2alpha1.InsightsConfig{}, options),
	}
}

func (c *insightsConfigReconcileLoop) RunInsightsConfigReconciler(ctx context.Context, reconciler InsightsConfigReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericInsightsConfigReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(InsightsConfigFinalizer); ok {
		reconcilerWrapper = genericInsightsConfigFinalizer{
			genericInsightsConfigReconciler: genericReconciler,
			finalizingReconciler:            finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericInsightsConfigHandler implements a generic reconcile.Reconciler
type genericInsightsConfigReconciler struct {
	reconciler InsightsConfigReconciler
}

func (r genericInsightsConfigReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*admin_gloo_solo_io_v2alpha1.InsightsConfig)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: InsightsConfig handler received event for %T", object)
	}
	return r.reconciler.ReconcileInsightsConfig(obj)
}

func (r genericInsightsConfigReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(InsightsConfigDeletionReconciler); ok {
		return deletionReconciler.ReconcileInsightsConfigDeletion(request)
	}
	return nil
}

// genericInsightsConfigFinalizer implements a generic reconcile.FinalizingReconciler
type genericInsightsConfigFinalizer struct {
	genericInsightsConfigReconciler
	finalizingReconciler InsightsConfigFinalizer
}

func (r genericInsightsConfigFinalizer) FinalizerName() string {
	return r.finalizingReconciler.InsightsConfigFinalizerName()
}

func (r genericInsightsConfigFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2alpha1.InsightsConfig)
	if !ok {
		return errors.Errorf("internal error: InsightsConfig handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeInsightsConfig(obj)
}
