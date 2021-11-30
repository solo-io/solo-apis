// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	fed_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the GlooInstance Resource.
// implemented by the user
type GlooInstanceReconciler interface {
	ReconcileGlooInstance(obj *fed_solo_io_v1.GlooInstance) (reconcile.Result, error)
}

// Reconcile deletion events for the GlooInstance Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type GlooInstanceDeletionReconciler interface {
	ReconcileGlooInstanceDeletion(req reconcile.Request) error
}

type GlooInstanceReconcilerFuncs struct {
	OnReconcileGlooInstance         func(obj *fed_solo_io_v1.GlooInstance) (reconcile.Result, error)
	OnReconcileGlooInstanceDeletion func(req reconcile.Request) error
}

func (f *GlooInstanceReconcilerFuncs) ReconcileGlooInstance(obj *fed_solo_io_v1.GlooInstance) (reconcile.Result, error) {
	if f.OnReconcileGlooInstance == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGlooInstance(obj)
}

func (f *GlooInstanceReconcilerFuncs) ReconcileGlooInstanceDeletion(req reconcile.Request) error {
	if f.OnReconcileGlooInstanceDeletion == nil {
		return nil
	}
	return f.OnReconcileGlooInstanceDeletion(req)
}

// Reconcile and finalize the GlooInstance Resource
// implemented by the user
type GlooInstanceFinalizer interface {
	GlooInstanceReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	GlooInstanceFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeGlooInstance(obj *fed_solo_io_v1.GlooInstance) error
}

type GlooInstanceReconcileLoop interface {
	RunGlooInstanceReconciler(ctx context.Context, rec GlooInstanceReconciler, predicates ...predicate.Predicate) error
}

type glooInstanceReconcileLoop struct {
	loop reconcile.Loop
}

func NewGlooInstanceReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) GlooInstanceReconcileLoop {
	return &glooInstanceReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &fed_solo_io_v1.GlooInstance{}, options),
	}
}

func (c *glooInstanceReconcileLoop) RunGlooInstanceReconciler(ctx context.Context, reconciler GlooInstanceReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericGlooInstanceReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(GlooInstanceFinalizer); ok {
		reconcilerWrapper = genericGlooInstanceFinalizer{
			genericGlooInstanceReconciler: genericReconciler,
			finalizingReconciler:          finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericGlooInstanceHandler implements a generic reconcile.Reconciler
type genericGlooInstanceReconciler struct {
	reconciler GlooInstanceReconciler
}

func (r genericGlooInstanceReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_solo_io_v1.GlooInstance)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GlooInstance handler received event for %T", object)
	}
	return r.reconciler.ReconcileGlooInstance(obj)
}

func (r genericGlooInstanceReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(GlooInstanceDeletionReconciler); ok {
		return deletionReconciler.ReconcileGlooInstanceDeletion(request)
	}
	return nil
}

// genericGlooInstanceFinalizer implements a generic reconcile.FinalizingReconciler
type genericGlooInstanceFinalizer struct {
	genericGlooInstanceReconciler
	finalizingReconciler GlooInstanceFinalizer
}

func (r genericGlooInstanceFinalizer) FinalizerName() string {
	return r.finalizingReconciler.GlooInstanceFinalizerName()
}

func (r genericGlooInstanceFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*fed_solo_io_v1.GlooInstance)
	if !ok {
		return errors.Errorf("internal error: GlooInstance handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeGlooInstance(obj)
}

// Reconcile Upsert events for the FailoverScheme Resource.
// implemented by the user
type FailoverSchemeReconciler interface {
	ReconcileFailoverScheme(obj *fed_solo_io_v1.FailoverScheme) (reconcile.Result, error)
}

// Reconcile deletion events for the FailoverScheme Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type FailoverSchemeDeletionReconciler interface {
	ReconcileFailoverSchemeDeletion(req reconcile.Request) error
}

type FailoverSchemeReconcilerFuncs struct {
	OnReconcileFailoverScheme         func(obj *fed_solo_io_v1.FailoverScheme) (reconcile.Result, error)
	OnReconcileFailoverSchemeDeletion func(req reconcile.Request) error
}

func (f *FailoverSchemeReconcilerFuncs) ReconcileFailoverScheme(obj *fed_solo_io_v1.FailoverScheme) (reconcile.Result, error) {
	if f.OnReconcileFailoverScheme == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFailoverScheme(obj)
}

func (f *FailoverSchemeReconcilerFuncs) ReconcileFailoverSchemeDeletion(req reconcile.Request) error {
	if f.OnReconcileFailoverSchemeDeletion == nil {
		return nil
	}
	return f.OnReconcileFailoverSchemeDeletion(req)
}

// Reconcile and finalize the FailoverScheme Resource
// implemented by the user
type FailoverSchemeFinalizer interface {
	FailoverSchemeReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	FailoverSchemeFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeFailoverScheme(obj *fed_solo_io_v1.FailoverScheme) error
}

type FailoverSchemeReconcileLoop interface {
	RunFailoverSchemeReconciler(ctx context.Context, rec FailoverSchemeReconciler, predicates ...predicate.Predicate) error
}

type failoverSchemeReconcileLoop struct {
	loop reconcile.Loop
}

func NewFailoverSchemeReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) FailoverSchemeReconcileLoop {
	return &failoverSchemeReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &fed_solo_io_v1.FailoverScheme{}, options),
	}
}

func (c *failoverSchemeReconcileLoop) RunFailoverSchemeReconciler(ctx context.Context, reconciler FailoverSchemeReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericFailoverSchemeReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(FailoverSchemeFinalizer); ok {
		reconcilerWrapper = genericFailoverSchemeFinalizer{
			genericFailoverSchemeReconciler: genericReconciler,
			finalizingReconciler:            finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericFailoverSchemeHandler implements a generic reconcile.Reconciler
type genericFailoverSchemeReconciler struct {
	reconciler FailoverSchemeReconciler
}

func (r genericFailoverSchemeReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_solo_io_v1.FailoverScheme)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FailoverScheme handler received event for %T", object)
	}
	return r.reconciler.ReconcileFailoverScheme(obj)
}

func (r genericFailoverSchemeReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(FailoverSchemeDeletionReconciler); ok {
		return deletionReconciler.ReconcileFailoverSchemeDeletion(request)
	}
	return nil
}

// genericFailoverSchemeFinalizer implements a generic reconcile.FinalizingReconciler
type genericFailoverSchemeFinalizer struct {
	genericFailoverSchemeReconciler
	finalizingReconciler FailoverSchemeFinalizer
}

func (r genericFailoverSchemeFinalizer) FinalizerName() string {
	return r.finalizingReconciler.FailoverSchemeFinalizerName()
}

func (r genericFailoverSchemeFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*fed_solo_io_v1.FailoverScheme)
	if !ok {
		return errors.Errorf("internal error: FailoverScheme handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeFailoverScheme(obj)
}
