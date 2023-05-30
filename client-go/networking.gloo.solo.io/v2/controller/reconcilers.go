// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	networking_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the ExternalService Resource.
// implemented by the user
type ExternalServiceReconciler interface {
	ReconcileExternalService(obj *networking_gloo_solo_io_v2.ExternalService) (reconcile.Result, error)
}

// Reconcile deletion events for the ExternalService Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type ExternalServiceDeletionReconciler interface {
	ReconcileExternalServiceDeletion(req reconcile.Request) error
}

type ExternalServiceReconcilerFuncs struct {
	OnReconcileExternalService         func(obj *networking_gloo_solo_io_v2.ExternalService) (reconcile.Result, error)
	OnReconcileExternalServiceDeletion func(req reconcile.Request) error
}

func (f *ExternalServiceReconcilerFuncs) ReconcileExternalService(obj *networking_gloo_solo_io_v2.ExternalService) (reconcile.Result, error) {
	if f.OnReconcileExternalService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileExternalService(obj)
}

func (f *ExternalServiceReconcilerFuncs) ReconcileExternalServiceDeletion(req reconcile.Request) error {
	if f.OnReconcileExternalServiceDeletion == nil {
		return nil
	}
	return f.OnReconcileExternalServiceDeletion(req)
}

// Reconcile and finalize the ExternalService Resource
// implemented by the user
type ExternalServiceFinalizer interface {
	ExternalServiceReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	ExternalServiceFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeExternalService(obj *networking_gloo_solo_io_v2.ExternalService) error
}

type ExternalServiceReconcileLoop interface {
	RunExternalServiceReconciler(ctx context.Context, rec ExternalServiceReconciler, predicates ...predicate.Predicate) error
}

type externalServiceReconcileLoop struct {
	loop reconcile.Loop
}

func NewExternalServiceReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) ExternalServiceReconcileLoop {
	return &externalServiceReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_gloo_solo_io_v2.ExternalService{}, options),
	}
}

func (c *externalServiceReconcileLoop) RunExternalServiceReconciler(ctx context.Context, reconciler ExternalServiceReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericExternalServiceReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(ExternalServiceFinalizer); ok {
		reconcilerWrapper = genericExternalServiceFinalizer{
			genericExternalServiceReconciler: genericReconciler,
			finalizingReconciler:             finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericExternalServiceHandler implements a generic reconcile.Reconciler
type genericExternalServiceReconciler struct {
	reconciler ExternalServiceReconciler
}

func (r genericExternalServiceReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ExternalService handler received event for %T", object)
	}
	return r.reconciler.ReconcileExternalService(obj)
}

func (r genericExternalServiceReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(ExternalServiceDeletionReconciler); ok {
		return deletionReconciler.ReconcileExternalServiceDeletion(request)
	}
	return nil
}

// genericExternalServiceFinalizer implements a generic reconcile.FinalizingReconciler
type genericExternalServiceFinalizer struct {
	genericExternalServiceReconciler
	finalizingReconciler ExternalServiceFinalizer
}

func (r genericExternalServiceFinalizer) FinalizerName() string {
	return r.finalizingReconciler.ExternalServiceFinalizerName()
}

func (r genericExternalServiceFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalService)
	if !ok {
		return errors.Errorf("internal error: ExternalService handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeExternalService(obj)
}

// Reconcile Upsert events for the ExternalEndpoint Resource.
// implemented by the user
type ExternalEndpointReconciler interface {
	ReconcileExternalEndpoint(obj *networking_gloo_solo_io_v2.ExternalEndpoint) (reconcile.Result, error)
}

// Reconcile deletion events for the ExternalEndpoint Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type ExternalEndpointDeletionReconciler interface {
	ReconcileExternalEndpointDeletion(req reconcile.Request) error
}

type ExternalEndpointReconcilerFuncs struct {
	OnReconcileExternalEndpoint         func(obj *networking_gloo_solo_io_v2.ExternalEndpoint) (reconcile.Result, error)
	OnReconcileExternalEndpointDeletion func(req reconcile.Request) error
}

func (f *ExternalEndpointReconcilerFuncs) ReconcileExternalEndpoint(obj *networking_gloo_solo_io_v2.ExternalEndpoint) (reconcile.Result, error) {
	if f.OnReconcileExternalEndpoint == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileExternalEndpoint(obj)
}

func (f *ExternalEndpointReconcilerFuncs) ReconcileExternalEndpointDeletion(req reconcile.Request) error {
	if f.OnReconcileExternalEndpointDeletion == nil {
		return nil
	}
	return f.OnReconcileExternalEndpointDeletion(req)
}

// Reconcile and finalize the ExternalEndpoint Resource
// implemented by the user
type ExternalEndpointFinalizer interface {
	ExternalEndpointReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	ExternalEndpointFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeExternalEndpoint(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error
}

type ExternalEndpointReconcileLoop interface {
	RunExternalEndpointReconciler(ctx context.Context, rec ExternalEndpointReconciler, predicates ...predicate.Predicate) error
}

type externalEndpointReconcileLoop struct {
	loop reconcile.Loop
}

func NewExternalEndpointReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) ExternalEndpointReconcileLoop {
	return &externalEndpointReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_gloo_solo_io_v2.ExternalEndpoint{}, options),
	}
}

func (c *externalEndpointReconcileLoop) RunExternalEndpointReconciler(ctx context.Context, reconciler ExternalEndpointReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericExternalEndpointReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(ExternalEndpointFinalizer); ok {
		reconcilerWrapper = genericExternalEndpointFinalizer{
			genericExternalEndpointReconciler: genericReconciler,
			finalizingReconciler:              finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericExternalEndpointHandler implements a generic reconcile.Reconciler
type genericExternalEndpointReconciler struct {
	reconciler ExternalEndpointReconciler
}

func (r genericExternalEndpointReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalEndpoint)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ExternalEndpoint handler received event for %T", object)
	}
	return r.reconciler.ReconcileExternalEndpoint(obj)
}

func (r genericExternalEndpointReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(ExternalEndpointDeletionReconciler); ok {
		return deletionReconciler.ReconcileExternalEndpointDeletion(request)
	}
	return nil
}

// genericExternalEndpointFinalizer implements a generic reconcile.FinalizingReconciler
type genericExternalEndpointFinalizer struct {
	genericExternalEndpointReconciler
	finalizingReconciler ExternalEndpointFinalizer
}

func (r genericExternalEndpointFinalizer) FinalizerName() string {
	return r.finalizingReconciler.ExternalEndpointFinalizerName()
}

func (r genericExternalEndpointFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalEndpoint)
	if !ok {
		return errors.Errorf("internal error: ExternalEndpoint handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeExternalEndpoint(obj)
}

// Reconcile Upsert events for the RouteTable Resource.
// implemented by the user
type RouteTableReconciler interface {
	ReconcileRouteTable(obj *networking_gloo_solo_io_v2.RouteTable) (reconcile.Result, error)
}

// Reconcile deletion events for the RouteTable Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type RouteTableDeletionReconciler interface {
	ReconcileRouteTableDeletion(req reconcile.Request) error
}

type RouteTableReconcilerFuncs struct {
	OnReconcileRouteTable         func(obj *networking_gloo_solo_io_v2.RouteTable) (reconcile.Result, error)
	OnReconcileRouteTableDeletion func(req reconcile.Request) error
}

func (f *RouteTableReconcilerFuncs) ReconcileRouteTable(obj *networking_gloo_solo_io_v2.RouteTable) (reconcile.Result, error) {
	if f.OnReconcileRouteTable == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileRouteTable(obj)
}

func (f *RouteTableReconcilerFuncs) ReconcileRouteTableDeletion(req reconcile.Request) error {
	if f.OnReconcileRouteTableDeletion == nil {
		return nil
	}
	return f.OnReconcileRouteTableDeletion(req)
}

// Reconcile and finalize the RouteTable Resource
// implemented by the user
type RouteTableFinalizer interface {
	RouteTableReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	RouteTableFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeRouteTable(obj *networking_gloo_solo_io_v2.RouteTable) error
}

type RouteTableReconcileLoop interface {
	RunRouteTableReconciler(ctx context.Context, rec RouteTableReconciler, predicates ...predicate.Predicate) error
}

type routeTableReconcileLoop struct {
	loop reconcile.Loop
}

func NewRouteTableReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) RouteTableReconcileLoop {
	return &routeTableReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_gloo_solo_io_v2.RouteTable{}, options),
	}
}

func (c *routeTableReconcileLoop) RunRouteTableReconciler(ctx context.Context, reconciler RouteTableReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericRouteTableReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(RouteTableFinalizer); ok {
		reconcilerWrapper = genericRouteTableFinalizer{
			genericRouteTableReconciler: genericReconciler,
			finalizingReconciler:        finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericRouteTableHandler implements a generic reconcile.Reconciler
type genericRouteTableReconciler struct {
	reconciler RouteTableReconciler
}

func (r genericRouteTableReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_gloo_solo_io_v2.RouteTable)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return r.reconciler.ReconcileRouteTable(obj)
}

func (r genericRouteTableReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(RouteTableDeletionReconciler); ok {
		return deletionReconciler.ReconcileRouteTableDeletion(request)
	}
	return nil
}

// genericRouteTableFinalizer implements a generic reconcile.FinalizingReconciler
type genericRouteTableFinalizer struct {
	genericRouteTableReconciler
	finalizingReconciler RouteTableFinalizer
}

func (r genericRouteTableFinalizer) FinalizerName() string {
	return r.finalizingReconciler.RouteTableFinalizerName()
}

func (r genericRouteTableFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeRouteTable(obj)
}

// Reconcile Upsert events for the VirtualDestination Resource.
// implemented by the user
type VirtualDestinationReconciler interface {
	ReconcileVirtualDestination(obj *networking_gloo_solo_io_v2.VirtualDestination) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualDestination Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type VirtualDestinationDeletionReconciler interface {
	ReconcileVirtualDestinationDeletion(req reconcile.Request) error
}

type VirtualDestinationReconcilerFuncs struct {
	OnReconcileVirtualDestination         func(obj *networking_gloo_solo_io_v2.VirtualDestination) (reconcile.Result, error)
	OnReconcileVirtualDestinationDeletion func(req reconcile.Request) error
}

func (f *VirtualDestinationReconcilerFuncs) ReconcileVirtualDestination(obj *networking_gloo_solo_io_v2.VirtualDestination) (reconcile.Result, error) {
	if f.OnReconcileVirtualDestination == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualDestination(obj)
}

func (f *VirtualDestinationReconcilerFuncs) ReconcileVirtualDestinationDeletion(req reconcile.Request) error {
	if f.OnReconcileVirtualDestinationDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualDestinationDeletion(req)
}

// Reconcile and finalize the VirtualDestination Resource
// implemented by the user
type VirtualDestinationFinalizer interface {
	VirtualDestinationReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	VirtualDestinationFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeVirtualDestination(obj *networking_gloo_solo_io_v2.VirtualDestination) error
}

type VirtualDestinationReconcileLoop interface {
	RunVirtualDestinationReconciler(ctx context.Context, rec VirtualDestinationReconciler, predicates ...predicate.Predicate) error
}

type virtualDestinationReconcileLoop struct {
	loop reconcile.Loop
}

func NewVirtualDestinationReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) VirtualDestinationReconcileLoop {
	return &virtualDestinationReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_gloo_solo_io_v2.VirtualDestination{}, options),
	}
}

func (c *virtualDestinationReconcileLoop) RunVirtualDestinationReconciler(ctx context.Context, reconciler VirtualDestinationReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericVirtualDestinationReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(VirtualDestinationFinalizer); ok {
		reconcilerWrapper = genericVirtualDestinationFinalizer{
			genericVirtualDestinationReconciler: genericReconciler,
			finalizingReconciler:                finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericVirtualDestinationHandler implements a generic reconcile.Reconciler
type genericVirtualDestinationReconciler struct {
	reconciler VirtualDestinationReconciler
}

func (r genericVirtualDestinationReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualDestination)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualDestination handler received event for %T", object)
	}
	return r.reconciler.ReconcileVirtualDestination(obj)
}

func (r genericVirtualDestinationReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(VirtualDestinationDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualDestinationDeletion(request)
	}
	return nil
}

// genericVirtualDestinationFinalizer implements a generic reconcile.FinalizingReconciler
type genericVirtualDestinationFinalizer struct {
	genericVirtualDestinationReconciler
	finalizingReconciler VirtualDestinationFinalizer
}

func (r genericVirtualDestinationFinalizer) FinalizerName() string {
	return r.finalizingReconciler.VirtualDestinationFinalizerName()
}

func (r genericVirtualDestinationFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualDestination)
	if !ok {
		return errors.Errorf("internal error: VirtualDestination handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeVirtualDestination(obj)
}

// Reconcile Upsert events for the VirtualGateway Resource.
// implemented by the user
type VirtualGatewayReconciler interface {
	ReconcileVirtualGateway(obj *networking_gloo_solo_io_v2.VirtualGateway) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualGateway Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type VirtualGatewayDeletionReconciler interface {
	ReconcileVirtualGatewayDeletion(req reconcile.Request) error
}

type VirtualGatewayReconcilerFuncs struct {
	OnReconcileVirtualGateway         func(obj *networking_gloo_solo_io_v2.VirtualGateway) (reconcile.Result, error)
	OnReconcileVirtualGatewayDeletion func(req reconcile.Request) error
}

func (f *VirtualGatewayReconcilerFuncs) ReconcileVirtualGateway(obj *networking_gloo_solo_io_v2.VirtualGateway) (reconcile.Result, error) {
	if f.OnReconcileVirtualGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualGateway(obj)
}

func (f *VirtualGatewayReconcilerFuncs) ReconcileVirtualGatewayDeletion(req reconcile.Request) error {
	if f.OnReconcileVirtualGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualGatewayDeletion(req)
}

// Reconcile and finalize the VirtualGateway Resource
// implemented by the user
type VirtualGatewayFinalizer interface {
	VirtualGatewayReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	VirtualGatewayFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeVirtualGateway(obj *networking_gloo_solo_io_v2.VirtualGateway) error
}

type VirtualGatewayReconcileLoop interface {
	RunVirtualGatewayReconciler(ctx context.Context, rec VirtualGatewayReconciler, predicates ...predicate.Predicate) error
}

type virtualGatewayReconcileLoop struct {
	loop reconcile.Loop
}

func NewVirtualGatewayReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) VirtualGatewayReconcileLoop {
	return &virtualGatewayReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_gloo_solo_io_v2.VirtualGateway{}, options),
	}
}

func (c *virtualGatewayReconcileLoop) RunVirtualGatewayReconciler(ctx context.Context, reconciler VirtualGatewayReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericVirtualGatewayReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(VirtualGatewayFinalizer); ok {
		reconcilerWrapper = genericVirtualGatewayFinalizer{
			genericVirtualGatewayReconciler: genericReconciler,
			finalizingReconciler:            finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericVirtualGatewayHandler implements a generic reconcile.Reconciler
type genericVirtualGatewayReconciler struct {
	reconciler VirtualGatewayReconciler
}

func (r genericVirtualGatewayReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualGateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return r.reconciler.ReconcileVirtualGateway(obj)
}

func (r genericVirtualGatewayReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(VirtualGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualGatewayDeletion(request)
	}
	return nil
}

// genericVirtualGatewayFinalizer implements a generic reconcile.FinalizingReconciler
type genericVirtualGatewayFinalizer struct {
	genericVirtualGatewayReconciler
	finalizingReconciler VirtualGatewayFinalizer
}

func (r genericVirtualGatewayFinalizer) FinalizerName() string {
	return r.finalizingReconciler.VirtualGatewayFinalizerName()
}

func (r genericVirtualGatewayFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeVirtualGateway(obj)
}

// Reconcile Upsert events for the ExternalWorkload Resource.
// implemented by the user
type ExternalWorkloadReconciler interface {
	ReconcileExternalWorkload(obj *networking_gloo_solo_io_v2.ExternalWorkload) (reconcile.Result, error)
}

// Reconcile deletion events for the ExternalWorkload Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type ExternalWorkloadDeletionReconciler interface {
	ReconcileExternalWorkloadDeletion(req reconcile.Request) error
}

type ExternalWorkloadReconcilerFuncs struct {
	OnReconcileExternalWorkload         func(obj *networking_gloo_solo_io_v2.ExternalWorkload) (reconcile.Result, error)
	OnReconcileExternalWorkloadDeletion func(req reconcile.Request) error
}

func (f *ExternalWorkloadReconcilerFuncs) ReconcileExternalWorkload(obj *networking_gloo_solo_io_v2.ExternalWorkload) (reconcile.Result, error) {
	if f.OnReconcileExternalWorkload == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileExternalWorkload(obj)
}

func (f *ExternalWorkloadReconcilerFuncs) ReconcileExternalWorkloadDeletion(req reconcile.Request) error {
	if f.OnReconcileExternalWorkloadDeletion == nil {
		return nil
	}
	return f.OnReconcileExternalWorkloadDeletion(req)
}

// Reconcile and finalize the ExternalWorkload Resource
// implemented by the user
type ExternalWorkloadFinalizer interface {
	ExternalWorkloadReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	ExternalWorkloadFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeExternalWorkload(obj *networking_gloo_solo_io_v2.ExternalWorkload) error
}

type ExternalWorkloadReconcileLoop interface {
	RunExternalWorkloadReconciler(ctx context.Context, rec ExternalWorkloadReconciler, predicates ...predicate.Predicate) error
}

type externalWorkloadReconcileLoop struct {
	loop reconcile.Loop
}

func NewExternalWorkloadReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) ExternalWorkloadReconcileLoop {
	return &externalWorkloadReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_gloo_solo_io_v2.ExternalWorkload{}, options),
	}
}

func (c *externalWorkloadReconcileLoop) RunExternalWorkloadReconciler(ctx context.Context, reconciler ExternalWorkloadReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericExternalWorkloadReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(ExternalWorkloadFinalizer); ok {
		reconcilerWrapper = genericExternalWorkloadFinalizer{
			genericExternalWorkloadReconciler: genericReconciler,
			finalizingReconciler:              finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericExternalWorkloadHandler implements a generic reconcile.Reconciler
type genericExternalWorkloadReconciler struct {
	reconciler ExternalWorkloadReconciler
}

func (r genericExternalWorkloadReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalWorkload)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ExternalWorkload handler received event for %T", object)
	}
	return r.reconciler.ReconcileExternalWorkload(obj)
}

func (r genericExternalWorkloadReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(ExternalWorkloadDeletionReconciler); ok {
		return deletionReconciler.ReconcileExternalWorkloadDeletion(request)
	}
	return nil
}

// genericExternalWorkloadFinalizer implements a generic reconcile.FinalizingReconciler
type genericExternalWorkloadFinalizer struct {
	genericExternalWorkloadReconciler
	finalizingReconciler ExternalWorkloadFinalizer
}

func (r genericExternalWorkloadFinalizer) FinalizerName() string {
	return r.finalizingReconciler.ExternalWorkloadFinalizerName()
}

func (r genericExternalWorkloadFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalWorkload)
	if !ok {
		return errors.Errorf("internal error: ExternalWorkload handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeExternalWorkload(obj)
}
