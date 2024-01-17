// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	security_policy_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/security.policy.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the AccessPolicy Resource.
// implemented by the user
type AccessPolicyReconciler interface {
	ReconcileAccessPolicy(obj *security_policy_gloo_solo_io_v2.AccessPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the AccessPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type AccessPolicyDeletionReconciler interface {
	ReconcileAccessPolicyDeletion(req reconcile.Request) error
}

type AccessPolicyReconcilerFuncs struct {
	OnReconcileAccessPolicy         func(obj *security_policy_gloo_solo_io_v2.AccessPolicy) (reconcile.Result, error)
	OnReconcileAccessPolicyDeletion func(req reconcile.Request) error
}

func (f *AccessPolicyReconcilerFuncs) ReconcileAccessPolicy(obj *security_policy_gloo_solo_io_v2.AccessPolicy) (reconcile.Result, error) {
	if f.OnReconcileAccessPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAccessPolicy(obj)
}

func (f *AccessPolicyReconcilerFuncs) ReconcileAccessPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileAccessPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileAccessPolicyDeletion(req)
}

// Reconcile and finalize the AccessPolicy Resource
// implemented by the user
type AccessPolicyFinalizer interface {
	AccessPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	AccessPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeAccessPolicy(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error
}

type AccessPolicyReconcileLoop interface {
	RunAccessPolicyReconciler(ctx context.Context, rec AccessPolicyReconciler, predicates ...predicate.Predicate) error
}

type accessPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewAccessPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) AccessPolicyReconcileLoop {
	return &accessPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_policy_gloo_solo_io_v2.AccessPolicy{}, options),
	}
}

func (c *accessPolicyReconcileLoop) RunAccessPolicyReconciler(ctx context.Context, reconciler AccessPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericAccessPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(AccessPolicyFinalizer); ok {
		reconcilerWrapper = genericAccessPolicyFinalizer{
			genericAccessPolicyReconciler: genericReconciler,
			finalizingReconciler:          finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericAccessPolicyHandler implements a generic reconcile.Reconciler
type genericAccessPolicyReconciler struct {
	reconciler AccessPolicyReconciler
}

func (r genericAccessPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.AccessPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileAccessPolicy(obj)
}

func (r genericAccessPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(AccessPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileAccessPolicyDeletion(request)
	}
	return nil
}

// genericAccessPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericAccessPolicyFinalizer struct {
	genericAccessPolicyReconciler
	finalizingReconciler AccessPolicyFinalizer
}

func (r genericAccessPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.AccessPolicyFinalizerName()
}

func (r genericAccessPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.AccessPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeAccessPolicy(obj)
}

// Reconcile Upsert events for the CORSPolicy Resource.
// implemented by the user
type CORSPolicyReconciler interface {
	ReconcileCORSPolicy(obj *security_policy_gloo_solo_io_v2.CORSPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the CORSPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type CORSPolicyDeletionReconciler interface {
	ReconcileCORSPolicyDeletion(req reconcile.Request) error
}

type CORSPolicyReconcilerFuncs struct {
	OnReconcileCORSPolicy         func(obj *security_policy_gloo_solo_io_v2.CORSPolicy) (reconcile.Result, error)
	OnReconcileCORSPolicyDeletion func(req reconcile.Request) error
}

func (f *CORSPolicyReconcilerFuncs) ReconcileCORSPolicy(obj *security_policy_gloo_solo_io_v2.CORSPolicy) (reconcile.Result, error) {
	if f.OnReconcileCORSPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCORSPolicy(obj)
}

func (f *CORSPolicyReconcilerFuncs) ReconcileCORSPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileCORSPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileCORSPolicyDeletion(req)
}

// Reconcile and finalize the CORSPolicy Resource
// implemented by the user
type CORSPolicyFinalizer interface {
	CORSPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	CORSPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeCORSPolicy(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error
}

type CORSPolicyReconcileLoop interface {
	RunCORSPolicyReconciler(ctx context.Context, rec CORSPolicyReconciler, predicates ...predicate.Predicate) error
}

type cORSPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewCORSPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) CORSPolicyReconcileLoop {
	return &cORSPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_policy_gloo_solo_io_v2.CORSPolicy{}, options),
	}
}

func (c *cORSPolicyReconcileLoop) RunCORSPolicyReconciler(ctx context.Context, reconciler CORSPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericCORSPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(CORSPolicyFinalizer); ok {
		reconcilerWrapper = genericCORSPolicyFinalizer{
			genericCORSPolicyReconciler: genericReconciler,
			finalizingReconciler:        finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericCORSPolicyHandler implements a generic reconcile.Reconciler
type genericCORSPolicyReconciler struct {
	reconciler CORSPolicyReconciler
}

func (r genericCORSPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CORSPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CORSPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileCORSPolicy(obj)
}

func (r genericCORSPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(CORSPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileCORSPolicyDeletion(request)
	}
	return nil
}

// genericCORSPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericCORSPolicyFinalizer struct {
	genericCORSPolicyReconciler
	finalizingReconciler CORSPolicyFinalizer
}

func (r genericCORSPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.CORSPolicyFinalizerName()
}

func (r genericCORSPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CORSPolicy)
	if !ok {
		return errors.Errorf("internal error: CORSPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeCORSPolicy(obj)
}

// Reconcile Upsert events for the CSRFPolicy Resource.
// implemented by the user
type CSRFPolicyReconciler interface {
	ReconcileCSRFPolicy(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the CSRFPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type CSRFPolicyDeletionReconciler interface {
	ReconcileCSRFPolicyDeletion(req reconcile.Request) error
}

type CSRFPolicyReconcilerFuncs struct {
	OnReconcileCSRFPolicy         func(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) (reconcile.Result, error)
	OnReconcileCSRFPolicyDeletion func(req reconcile.Request) error
}

func (f *CSRFPolicyReconcilerFuncs) ReconcileCSRFPolicy(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) (reconcile.Result, error) {
	if f.OnReconcileCSRFPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCSRFPolicy(obj)
}

func (f *CSRFPolicyReconcilerFuncs) ReconcileCSRFPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileCSRFPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileCSRFPolicyDeletion(req)
}

// Reconcile and finalize the CSRFPolicy Resource
// implemented by the user
type CSRFPolicyFinalizer interface {
	CSRFPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	CSRFPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeCSRFPolicy(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error
}

type CSRFPolicyReconcileLoop interface {
	RunCSRFPolicyReconciler(ctx context.Context, rec CSRFPolicyReconciler, predicates ...predicate.Predicate) error
}

type cSRFPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewCSRFPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) CSRFPolicyReconcileLoop {
	return &cSRFPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_policy_gloo_solo_io_v2.CSRFPolicy{}, options),
	}
}

func (c *cSRFPolicyReconcileLoop) RunCSRFPolicyReconciler(ctx context.Context, reconciler CSRFPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericCSRFPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(CSRFPolicyFinalizer); ok {
		reconcilerWrapper = genericCSRFPolicyFinalizer{
			genericCSRFPolicyReconciler: genericReconciler,
			finalizingReconciler:        finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericCSRFPolicyHandler implements a generic reconcile.Reconciler
type genericCSRFPolicyReconciler struct {
	reconciler CSRFPolicyReconciler
}

func (r genericCSRFPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CSRFPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CSRFPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileCSRFPolicy(obj)
}

func (r genericCSRFPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(CSRFPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileCSRFPolicyDeletion(request)
	}
	return nil
}

// genericCSRFPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericCSRFPolicyFinalizer struct {
	genericCSRFPolicyReconciler
	finalizingReconciler CSRFPolicyFinalizer
}

func (r genericCSRFPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.CSRFPolicyFinalizerName()
}

func (r genericCSRFPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CSRFPolicy)
	if !ok {
		return errors.Errorf("internal error: CSRFPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeCSRFPolicy(obj)
}

// Reconcile Upsert events for the ExtAuthPolicy Resource.
// implemented by the user
type ExtAuthPolicyReconciler interface {
	ReconcileExtAuthPolicy(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the ExtAuthPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type ExtAuthPolicyDeletionReconciler interface {
	ReconcileExtAuthPolicyDeletion(req reconcile.Request) error
}

type ExtAuthPolicyReconcilerFuncs struct {
	OnReconcileExtAuthPolicy         func(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) (reconcile.Result, error)
	OnReconcileExtAuthPolicyDeletion func(req reconcile.Request) error
}

func (f *ExtAuthPolicyReconcilerFuncs) ReconcileExtAuthPolicy(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) (reconcile.Result, error) {
	if f.OnReconcileExtAuthPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileExtAuthPolicy(obj)
}

func (f *ExtAuthPolicyReconcilerFuncs) ReconcileExtAuthPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileExtAuthPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileExtAuthPolicyDeletion(req)
}

// Reconcile and finalize the ExtAuthPolicy Resource
// implemented by the user
type ExtAuthPolicyFinalizer interface {
	ExtAuthPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	ExtAuthPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeExtAuthPolicy(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error
}

type ExtAuthPolicyReconcileLoop interface {
	RunExtAuthPolicyReconciler(ctx context.Context, rec ExtAuthPolicyReconciler, predicates ...predicate.Predicate) error
}

type extAuthPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewExtAuthPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) ExtAuthPolicyReconcileLoop {
	return &extAuthPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_policy_gloo_solo_io_v2.ExtAuthPolicy{}, options),
	}
}

func (c *extAuthPolicyReconcileLoop) RunExtAuthPolicyReconciler(ctx context.Context, reconciler ExtAuthPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericExtAuthPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(ExtAuthPolicyFinalizer); ok {
		reconcilerWrapper = genericExtAuthPolicyFinalizer{
			genericExtAuthPolicyReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericExtAuthPolicyHandler implements a generic reconcile.Reconciler
type genericExtAuthPolicyReconciler struct {
	reconciler ExtAuthPolicyReconciler
}

func (r genericExtAuthPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ExtAuthPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ExtAuthPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileExtAuthPolicy(obj)
}

func (r genericExtAuthPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(ExtAuthPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileExtAuthPolicyDeletion(request)
	}
	return nil
}

// genericExtAuthPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericExtAuthPolicyFinalizer struct {
	genericExtAuthPolicyReconciler
	finalizingReconciler ExtAuthPolicyFinalizer
}

func (r genericExtAuthPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.ExtAuthPolicyFinalizerName()
}

func (r genericExtAuthPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ExtAuthPolicy)
	if !ok {
		return errors.Errorf("internal error: ExtAuthPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeExtAuthPolicy(obj)
}

// Reconcile Upsert events for the WAFPolicy Resource.
// implemented by the user
type WAFPolicyReconciler interface {
	ReconcileWAFPolicy(obj *security_policy_gloo_solo_io_v2.WAFPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the WAFPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type WAFPolicyDeletionReconciler interface {
	ReconcileWAFPolicyDeletion(req reconcile.Request) error
}

type WAFPolicyReconcilerFuncs struct {
	OnReconcileWAFPolicy         func(obj *security_policy_gloo_solo_io_v2.WAFPolicy) (reconcile.Result, error)
	OnReconcileWAFPolicyDeletion func(req reconcile.Request) error
}

func (f *WAFPolicyReconcilerFuncs) ReconcileWAFPolicy(obj *security_policy_gloo_solo_io_v2.WAFPolicy) (reconcile.Result, error) {
	if f.OnReconcileWAFPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWAFPolicy(obj)
}

func (f *WAFPolicyReconcilerFuncs) ReconcileWAFPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileWAFPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileWAFPolicyDeletion(req)
}

// Reconcile and finalize the WAFPolicy Resource
// implemented by the user
type WAFPolicyFinalizer interface {
	WAFPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	WAFPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeWAFPolicy(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error
}

type WAFPolicyReconcileLoop interface {
	RunWAFPolicyReconciler(ctx context.Context, rec WAFPolicyReconciler, predicates ...predicate.Predicate) error
}

type wAFPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewWAFPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) WAFPolicyReconcileLoop {
	return &wAFPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_policy_gloo_solo_io_v2.WAFPolicy{}, options),
	}
}

func (c *wAFPolicyReconcileLoop) RunWAFPolicyReconciler(ctx context.Context, reconciler WAFPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericWAFPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(WAFPolicyFinalizer); ok {
		reconcilerWrapper = genericWAFPolicyFinalizer{
			genericWAFPolicyReconciler: genericReconciler,
			finalizingReconciler:       finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericWAFPolicyHandler implements a generic reconcile.Reconciler
type genericWAFPolicyReconciler struct {
	reconciler WAFPolicyReconciler
}

func (r genericWAFPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.WAFPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WAFPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileWAFPolicy(obj)
}

func (r genericWAFPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(WAFPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileWAFPolicyDeletion(request)
	}
	return nil
}

// genericWAFPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericWAFPolicyFinalizer struct {
	genericWAFPolicyReconciler
	finalizingReconciler WAFPolicyFinalizer
}

func (r genericWAFPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.WAFPolicyFinalizerName()
}

func (r genericWAFPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.WAFPolicy)
	if !ok {
		return errors.Errorf("internal error: WAFPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeWAFPolicy(obj)
}

// Reconcile Upsert events for the JWTPolicy Resource.
// implemented by the user
type JWTPolicyReconciler interface {
	ReconcileJWTPolicy(obj *security_policy_gloo_solo_io_v2.JWTPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the JWTPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type JWTPolicyDeletionReconciler interface {
	ReconcileJWTPolicyDeletion(req reconcile.Request) error
}

type JWTPolicyReconcilerFuncs struct {
	OnReconcileJWTPolicy         func(obj *security_policy_gloo_solo_io_v2.JWTPolicy) (reconcile.Result, error)
	OnReconcileJWTPolicyDeletion func(req reconcile.Request) error
}

func (f *JWTPolicyReconcilerFuncs) ReconcileJWTPolicy(obj *security_policy_gloo_solo_io_v2.JWTPolicy) (reconcile.Result, error) {
	if f.OnReconcileJWTPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileJWTPolicy(obj)
}

func (f *JWTPolicyReconcilerFuncs) ReconcileJWTPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileJWTPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileJWTPolicyDeletion(req)
}

// Reconcile and finalize the JWTPolicy Resource
// implemented by the user
type JWTPolicyFinalizer interface {
	JWTPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	JWTPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeJWTPolicy(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error
}

type JWTPolicyReconcileLoop interface {
	RunJWTPolicyReconciler(ctx context.Context, rec JWTPolicyReconciler, predicates ...predicate.Predicate) error
}

type jWTPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewJWTPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) JWTPolicyReconcileLoop {
	return &jWTPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_policy_gloo_solo_io_v2.JWTPolicy{}, options),
	}
}

func (c *jWTPolicyReconcileLoop) RunJWTPolicyReconciler(ctx context.Context, reconciler JWTPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericJWTPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(JWTPolicyFinalizer); ok {
		reconcilerWrapper = genericJWTPolicyFinalizer{
			genericJWTPolicyReconciler: genericReconciler,
			finalizingReconciler:       finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericJWTPolicyHandler implements a generic reconcile.Reconciler
type genericJWTPolicyReconciler struct {
	reconciler JWTPolicyReconciler
}

func (r genericJWTPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.JWTPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: JWTPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileJWTPolicy(obj)
}

func (r genericJWTPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(JWTPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileJWTPolicyDeletion(request)
	}
	return nil
}

// genericJWTPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericJWTPolicyFinalizer struct {
	genericJWTPolicyReconciler
	finalizingReconciler JWTPolicyFinalizer
}

func (r genericJWTPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.JWTPolicyFinalizerName()
}

func (r genericJWTPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.JWTPolicy)
	if !ok {
		return errors.Errorf("internal error: JWTPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeJWTPolicy(obj)
}

// Reconcile Upsert events for the ClientTLSPolicy Resource.
// implemented by the user
type ClientTLSPolicyReconciler interface {
	ReconcileClientTLSPolicy(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the ClientTLSPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type ClientTLSPolicyDeletionReconciler interface {
	ReconcileClientTLSPolicyDeletion(req reconcile.Request) error
}

type ClientTLSPolicyReconcilerFuncs struct {
	OnReconcileClientTLSPolicy         func(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) (reconcile.Result, error)
	OnReconcileClientTLSPolicyDeletion func(req reconcile.Request) error
}

func (f *ClientTLSPolicyReconcilerFuncs) ReconcileClientTLSPolicy(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) (reconcile.Result, error) {
	if f.OnReconcileClientTLSPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileClientTLSPolicy(obj)
}

func (f *ClientTLSPolicyReconcilerFuncs) ReconcileClientTLSPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileClientTLSPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileClientTLSPolicyDeletion(req)
}

// Reconcile and finalize the ClientTLSPolicy Resource
// implemented by the user
type ClientTLSPolicyFinalizer interface {
	ClientTLSPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	ClientTLSPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeClientTLSPolicy(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error
}

type ClientTLSPolicyReconcileLoop interface {
	RunClientTLSPolicyReconciler(ctx context.Context, rec ClientTLSPolicyReconciler, predicates ...predicate.Predicate) error
}

type clientTLSPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewClientTLSPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) ClientTLSPolicyReconcileLoop {
	return &clientTLSPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_policy_gloo_solo_io_v2.ClientTLSPolicy{}, options),
	}
}

func (c *clientTLSPolicyReconcileLoop) RunClientTLSPolicyReconciler(ctx context.Context, reconciler ClientTLSPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericClientTLSPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(ClientTLSPolicyFinalizer); ok {
		reconcilerWrapper = genericClientTLSPolicyFinalizer{
			genericClientTLSPolicyReconciler: genericReconciler,
			finalizingReconciler:             finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericClientTLSPolicyHandler implements a generic reconcile.Reconciler
type genericClientTLSPolicyReconciler struct {
	reconciler ClientTLSPolicyReconciler
}

func (r genericClientTLSPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ClientTLSPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ClientTLSPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileClientTLSPolicy(obj)
}

func (r genericClientTLSPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(ClientTLSPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileClientTLSPolicyDeletion(request)
	}
	return nil
}

// genericClientTLSPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericClientTLSPolicyFinalizer struct {
	genericClientTLSPolicyReconciler
	finalizingReconciler ClientTLSPolicyFinalizer
}

func (r genericClientTLSPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.ClientTLSPolicyFinalizerName()
}

func (r genericClientTLSPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ClientTLSPolicy)
	if !ok {
		return errors.Errorf("internal error: ClientTLSPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeClientTLSPolicy(obj)
}

// Reconcile Upsert events for the GraphQLAllowedQueryPolicy Resource.
// implemented by the user
type GraphQLAllowedQueryPolicyReconciler interface {
	ReconcileGraphQLAllowedQueryPolicy(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the GraphQLAllowedQueryPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type GraphQLAllowedQueryPolicyDeletionReconciler interface {
	ReconcileGraphQLAllowedQueryPolicyDeletion(req reconcile.Request) error
}

type GraphQLAllowedQueryPolicyReconcilerFuncs struct {
	OnReconcileGraphQLAllowedQueryPolicy         func(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) (reconcile.Result, error)
	OnReconcileGraphQLAllowedQueryPolicyDeletion func(req reconcile.Request) error
}

func (f *GraphQLAllowedQueryPolicyReconcilerFuncs) ReconcileGraphQLAllowedQueryPolicy(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) (reconcile.Result, error) {
	if f.OnReconcileGraphQLAllowedQueryPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGraphQLAllowedQueryPolicy(obj)
}

func (f *GraphQLAllowedQueryPolicyReconcilerFuncs) ReconcileGraphQLAllowedQueryPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileGraphQLAllowedQueryPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileGraphQLAllowedQueryPolicyDeletion(req)
}

// Reconcile and finalize the GraphQLAllowedQueryPolicy Resource
// implemented by the user
type GraphQLAllowedQueryPolicyFinalizer interface {
	GraphQLAllowedQueryPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	GraphQLAllowedQueryPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeGraphQLAllowedQueryPolicy(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error
}

type GraphQLAllowedQueryPolicyReconcileLoop interface {
	RunGraphQLAllowedQueryPolicyReconciler(ctx context.Context, rec GraphQLAllowedQueryPolicyReconciler, predicates ...predicate.Predicate) error
}

type graphQLAllowedQueryPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewGraphQLAllowedQueryPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) GraphQLAllowedQueryPolicyReconcileLoop {
	return &graphQLAllowedQueryPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy{}, options),
	}
}

func (c *graphQLAllowedQueryPolicyReconcileLoop) RunGraphQLAllowedQueryPolicyReconciler(ctx context.Context, reconciler GraphQLAllowedQueryPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericGraphQLAllowedQueryPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(GraphQLAllowedQueryPolicyFinalizer); ok {
		reconcilerWrapper = genericGraphQLAllowedQueryPolicyFinalizer{
			genericGraphQLAllowedQueryPolicyReconciler: genericReconciler,
			finalizingReconciler:                       finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericGraphQLAllowedQueryPolicyHandler implements a generic reconcile.Reconciler
type genericGraphQLAllowedQueryPolicyReconciler struct {
	reconciler GraphQLAllowedQueryPolicyReconciler
}

func (r genericGraphQLAllowedQueryPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GraphQLAllowedQueryPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileGraphQLAllowedQueryPolicy(obj)
}

func (r genericGraphQLAllowedQueryPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(GraphQLAllowedQueryPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileGraphQLAllowedQueryPolicyDeletion(request)
	}
	return nil
}

// genericGraphQLAllowedQueryPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericGraphQLAllowedQueryPolicyFinalizer struct {
	genericGraphQLAllowedQueryPolicyReconciler
	finalizingReconciler GraphQLAllowedQueryPolicyFinalizer
}

func (r genericGraphQLAllowedQueryPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.GraphQLAllowedQueryPolicyFinalizerName()
}

func (r genericGraphQLAllowedQueryPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy)
	if !ok {
		return errors.Errorf("internal error: GraphQLAllowedQueryPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeGraphQLAllowedQueryPolicy(obj)
}

// Reconcile Upsert events for the DLPPolicy Resource.
// implemented by the user
type DLPPolicyReconciler interface {
	ReconcileDLPPolicy(obj *security_policy_gloo_solo_io_v2.DLPPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the DLPPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type DLPPolicyDeletionReconciler interface {
	ReconcileDLPPolicyDeletion(req reconcile.Request) error
}

type DLPPolicyReconcilerFuncs struct {
	OnReconcileDLPPolicy         func(obj *security_policy_gloo_solo_io_v2.DLPPolicy) (reconcile.Result, error)
	OnReconcileDLPPolicyDeletion func(req reconcile.Request) error
}

func (f *DLPPolicyReconcilerFuncs) ReconcileDLPPolicy(obj *security_policy_gloo_solo_io_v2.DLPPolicy) (reconcile.Result, error) {
	if f.OnReconcileDLPPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileDLPPolicy(obj)
}

func (f *DLPPolicyReconcilerFuncs) ReconcileDLPPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileDLPPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileDLPPolicyDeletion(req)
}

// Reconcile and finalize the DLPPolicy Resource
// implemented by the user
type DLPPolicyFinalizer interface {
	DLPPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	DLPPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeDLPPolicy(obj *security_policy_gloo_solo_io_v2.DLPPolicy) error
}

type DLPPolicyReconcileLoop interface {
	RunDLPPolicyReconciler(ctx context.Context, rec DLPPolicyReconciler, predicates ...predicate.Predicate) error
}

type dLPPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewDLPPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) DLPPolicyReconcileLoop {
	return &dLPPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_policy_gloo_solo_io_v2.DLPPolicy{}, options),
	}
}

func (c *dLPPolicyReconcileLoop) RunDLPPolicyReconciler(ctx context.Context, reconciler DLPPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericDLPPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(DLPPolicyFinalizer); ok {
		reconcilerWrapper = genericDLPPolicyFinalizer{
			genericDLPPolicyReconciler: genericReconciler,
			finalizingReconciler:       finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericDLPPolicyHandler implements a generic reconcile.Reconciler
type genericDLPPolicyReconciler struct {
	reconciler DLPPolicyReconciler
}

func (r genericDLPPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.DLPPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: DLPPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileDLPPolicy(obj)
}

func (r genericDLPPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(DLPPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileDLPPolicyDeletion(request)
	}
	return nil
}

// genericDLPPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericDLPPolicyFinalizer struct {
	genericDLPPolicyReconciler
	finalizingReconciler DLPPolicyFinalizer
}

func (r genericDLPPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.DLPPolicyFinalizerName()
}

func (r genericDLPPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.DLPPolicy)
	if !ok {
		return errors.Errorf("internal error: DLPPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeDLPPolicy(obj)
}
