// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	fed_enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/fed.enterprise.gloo.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the FederatedAuthConfig Resource.
// implemented by the user
type FederatedAuthConfigReconciler interface {
	ReconcileFederatedAuthConfig(obj *fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) (reconcile.Result, error)
}

// Reconcile deletion events for the FederatedAuthConfig Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type FederatedAuthConfigDeletionReconciler interface {
	ReconcileFederatedAuthConfigDeletion(req reconcile.Request) error
}

type FederatedAuthConfigReconcilerFuncs struct {
	OnReconcileFederatedAuthConfig         func(obj *fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) (reconcile.Result, error)
	OnReconcileFederatedAuthConfigDeletion func(req reconcile.Request) error
}

func (f *FederatedAuthConfigReconcilerFuncs) ReconcileFederatedAuthConfig(obj *fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) (reconcile.Result, error) {
	if f.OnReconcileFederatedAuthConfig == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFederatedAuthConfig(obj)
}

func (f *FederatedAuthConfigReconcilerFuncs) ReconcileFederatedAuthConfigDeletion(req reconcile.Request) error {
	if f.OnReconcileFederatedAuthConfigDeletion == nil {
		return nil
	}
	return f.OnReconcileFederatedAuthConfigDeletion(req)
}

// Reconcile and finalize the FederatedAuthConfig Resource
// implemented by the user
type FederatedAuthConfigFinalizer interface {
	FederatedAuthConfigReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	FederatedAuthConfigFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeFederatedAuthConfig(obj *fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig) error
}

type FederatedAuthConfigReconcileLoop interface {
	RunFederatedAuthConfigReconciler(ctx context.Context, rec FederatedAuthConfigReconciler, predicates ...predicate.Predicate) error
}

type federatedAuthConfigReconcileLoop struct {
	loop reconcile.Loop
}

func NewFederatedAuthConfigReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) FederatedAuthConfigReconcileLoop {
	return &federatedAuthConfigReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig{}, options),
	}
}

func (c *federatedAuthConfigReconcileLoop) RunFederatedAuthConfigReconciler(ctx context.Context, reconciler FederatedAuthConfigReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericFederatedAuthConfigReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(FederatedAuthConfigFinalizer); ok {
		reconcilerWrapper = genericFederatedAuthConfigFinalizer{
			genericFederatedAuthConfigReconciler: genericReconciler,
			finalizingReconciler:                 finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericFederatedAuthConfigHandler implements a generic reconcile.Reconciler
type genericFederatedAuthConfigReconciler struct {
	reconciler FederatedAuthConfigReconciler
}

func (r genericFederatedAuthConfigReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FederatedAuthConfig handler received event for %T", object)
	}
	return r.reconciler.ReconcileFederatedAuthConfig(obj)
}

func (r genericFederatedAuthConfigReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(FederatedAuthConfigDeletionReconciler); ok {
		return deletionReconciler.ReconcileFederatedAuthConfigDeletion(request)
	}
	return nil
}

// genericFederatedAuthConfigFinalizer implements a generic reconcile.FinalizingReconciler
type genericFederatedAuthConfigFinalizer struct {
	genericFederatedAuthConfigReconciler
	finalizingReconciler FederatedAuthConfigFinalizer
}

func (r genericFederatedAuthConfigFinalizer) FinalizerName() string {
	return r.finalizingReconciler.FederatedAuthConfigFinalizerName()
}

func (r genericFederatedAuthConfigFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig)
	if !ok {
		return errors.Errorf("internal error: FederatedAuthConfig handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeFederatedAuthConfig(obj)
}
