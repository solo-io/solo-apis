// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the AuthConfig Resource.
// implemented by the user
type AuthConfigReconciler interface {
	ReconcileAuthConfig(obj *enterprise_gloo_solo_io_v1.AuthConfig) (reconcile.Result, error)
}

// Reconcile deletion events for the AuthConfig Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type AuthConfigDeletionReconciler interface {
	ReconcileAuthConfigDeletion(req reconcile.Request) error
}

type AuthConfigReconcilerFuncs struct {
	OnReconcileAuthConfig         func(obj *enterprise_gloo_solo_io_v1.AuthConfig) (reconcile.Result, error)
	OnReconcileAuthConfigDeletion func(req reconcile.Request) error
}

func (f *AuthConfigReconcilerFuncs) ReconcileAuthConfig(obj *enterprise_gloo_solo_io_v1.AuthConfig) (reconcile.Result, error) {
	if f.OnReconcileAuthConfig == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAuthConfig(obj)
}

func (f *AuthConfigReconcilerFuncs) ReconcileAuthConfigDeletion(req reconcile.Request) error {
	if f.OnReconcileAuthConfigDeletion == nil {
		return nil
	}
	return f.OnReconcileAuthConfigDeletion(req)
}

// Reconcile and finalize the AuthConfig Resource
// implemented by the user
type AuthConfigFinalizer interface {
	AuthConfigReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	AuthConfigFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeAuthConfig(obj *enterprise_gloo_solo_io_v1.AuthConfig) error
}

type AuthConfigReconcileLoop interface {
	RunAuthConfigReconciler(ctx context.Context, rec AuthConfigReconciler, predicates ...predicate.Predicate) error
}

type authConfigReconcileLoop struct {
	loop reconcile.Loop
}

func NewAuthConfigReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) AuthConfigReconcileLoop {
	return &authConfigReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &enterprise_gloo_solo_io_v1.AuthConfig{}, options),
	}
}

func (c *authConfigReconcileLoop) RunAuthConfigReconciler(ctx context.Context, reconciler AuthConfigReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericAuthConfigReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(AuthConfigFinalizer); ok {
		reconcilerWrapper = genericAuthConfigFinalizer{
			genericAuthConfigReconciler: genericReconciler,
			finalizingReconciler:        finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericAuthConfigHandler implements a generic reconcile.Reconciler
type genericAuthConfigReconciler struct {
	reconciler AuthConfigReconciler
}

func (r genericAuthConfigReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*enterprise_gloo_solo_io_v1.AuthConfig)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AuthConfig handler received event for %T", object)
	}
	return r.reconciler.ReconcileAuthConfig(obj)
}

func (r genericAuthConfigReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(AuthConfigDeletionReconciler); ok {
		return deletionReconciler.ReconcileAuthConfigDeletion(request)
	}
	return nil
}

// genericAuthConfigFinalizer implements a generic reconcile.FinalizingReconciler
type genericAuthConfigFinalizer struct {
	genericAuthConfigReconciler
	finalizingReconciler AuthConfigFinalizer
}

func (r genericAuthConfigFinalizer) FinalizerName() string {
	return r.finalizingReconciler.AuthConfigFinalizerName()
}

func (r genericAuthConfigFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*enterprise_gloo_solo_io_v1.AuthConfig)
	if !ok {
		return errors.Errorf("internal error: AuthConfig handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeAuthConfig(obj)
}
