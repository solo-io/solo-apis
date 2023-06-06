// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	networking_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/networking.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the ExternalService Resource
// DEPRECATED: Prefer reconciler pattern.
type ExternalServiceEventHandler interface {
	CreateExternalService(obj *networking_gloo_solo_io_v2.ExternalService) error
	UpdateExternalService(old, new *networking_gloo_solo_io_v2.ExternalService) error
	DeleteExternalService(obj *networking_gloo_solo_io_v2.ExternalService) error
	GenericExternalService(obj *networking_gloo_solo_io_v2.ExternalService) error
}

type ExternalServiceEventHandlerFuncs struct {
	OnCreate  func(obj *networking_gloo_solo_io_v2.ExternalService) error
	OnUpdate  func(old, new *networking_gloo_solo_io_v2.ExternalService) error
	OnDelete  func(obj *networking_gloo_solo_io_v2.ExternalService) error
	OnGeneric func(obj *networking_gloo_solo_io_v2.ExternalService) error
}

func (f *ExternalServiceEventHandlerFuncs) CreateExternalService(obj *networking_gloo_solo_io_v2.ExternalService) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ExternalServiceEventHandlerFuncs) DeleteExternalService(obj *networking_gloo_solo_io_v2.ExternalService) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ExternalServiceEventHandlerFuncs) UpdateExternalService(objOld, objNew *networking_gloo_solo_io_v2.ExternalService) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ExternalServiceEventHandlerFuncs) GenericExternalService(obj *networking_gloo_solo_io_v2.ExternalService) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ExternalServiceEventWatcher interface {
	AddEventHandler(ctx context.Context, h ExternalServiceEventHandler, predicates ...predicate.Predicate) error
}

type externalServiceEventWatcher struct {
	watcher events.EventWatcher
}

func NewExternalServiceEventWatcher(name string, mgr manager.Manager) ExternalServiceEventWatcher {
	return &externalServiceEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_gloo_solo_io_v2.ExternalService{}),
	}
}

func (c *externalServiceEventWatcher) AddEventHandler(ctx context.Context, h ExternalServiceEventHandler, predicates ...predicate.Predicate) error {
	handler := genericExternalServiceHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericExternalServiceHandler implements a generic events.EventHandler
type genericExternalServiceHandler struct {
	handler ExternalServiceEventHandler
}

func (h genericExternalServiceHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalService)
	if !ok {
		return errors.Errorf("internal error: ExternalService handler received event for %T", object)
	}
	return h.handler.CreateExternalService(obj)
}

func (h genericExternalServiceHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalService)
	if !ok {
		return errors.Errorf("internal error: ExternalService handler received event for %T", object)
	}
	return h.handler.DeleteExternalService(obj)
}

func (h genericExternalServiceHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_gloo_solo_io_v2.ExternalService)
	if !ok {
		return errors.Errorf("internal error: ExternalService handler received event for %T", old)
	}
	objNew, ok := new.(*networking_gloo_solo_io_v2.ExternalService)
	if !ok {
		return errors.Errorf("internal error: ExternalService handler received event for %T", new)
	}
	return h.handler.UpdateExternalService(objOld, objNew)
}

func (h genericExternalServiceHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalService)
	if !ok {
		return errors.Errorf("internal error: ExternalService handler received event for %T", object)
	}
	return h.handler.GenericExternalService(obj)
}

// Handle events for the ExternalEndpoint Resource
// DEPRECATED: Prefer reconciler pattern.
type ExternalEndpointEventHandler interface {
	CreateExternalEndpoint(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error
	UpdateExternalEndpoint(old, new *networking_gloo_solo_io_v2.ExternalEndpoint) error
	DeleteExternalEndpoint(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error
	GenericExternalEndpoint(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error
}

type ExternalEndpointEventHandlerFuncs struct {
	OnCreate  func(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error
	OnUpdate  func(old, new *networking_gloo_solo_io_v2.ExternalEndpoint) error
	OnDelete  func(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error
	OnGeneric func(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error
}

func (f *ExternalEndpointEventHandlerFuncs) CreateExternalEndpoint(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ExternalEndpointEventHandlerFuncs) DeleteExternalEndpoint(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ExternalEndpointEventHandlerFuncs) UpdateExternalEndpoint(objOld, objNew *networking_gloo_solo_io_v2.ExternalEndpoint) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ExternalEndpointEventHandlerFuncs) GenericExternalEndpoint(obj *networking_gloo_solo_io_v2.ExternalEndpoint) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ExternalEndpointEventWatcher interface {
	AddEventHandler(ctx context.Context, h ExternalEndpointEventHandler, predicates ...predicate.Predicate) error
}

type externalEndpointEventWatcher struct {
	watcher events.EventWatcher
}

func NewExternalEndpointEventWatcher(name string, mgr manager.Manager) ExternalEndpointEventWatcher {
	return &externalEndpointEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_gloo_solo_io_v2.ExternalEndpoint{}),
	}
}

func (c *externalEndpointEventWatcher) AddEventHandler(ctx context.Context, h ExternalEndpointEventHandler, predicates ...predicate.Predicate) error {
	handler := genericExternalEndpointHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericExternalEndpointHandler implements a generic events.EventHandler
type genericExternalEndpointHandler struct {
	handler ExternalEndpointEventHandler
}

func (h genericExternalEndpointHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalEndpoint)
	if !ok {
		return errors.Errorf("internal error: ExternalEndpoint handler received event for %T", object)
	}
	return h.handler.CreateExternalEndpoint(obj)
}

func (h genericExternalEndpointHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalEndpoint)
	if !ok {
		return errors.Errorf("internal error: ExternalEndpoint handler received event for %T", object)
	}
	return h.handler.DeleteExternalEndpoint(obj)
}

func (h genericExternalEndpointHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_gloo_solo_io_v2.ExternalEndpoint)
	if !ok {
		return errors.Errorf("internal error: ExternalEndpoint handler received event for %T", old)
	}
	objNew, ok := new.(*networking_gloo_solo_io_v2.ExternalEndpoint)
	if !ok {
		return errors.Errorf("internal error: ExternalEndpoint handler received event for %T", new)
	}
	return h.handler.UpdateExternalEndpoint(objOld, objNew)
}

func (h genericExternalEndpointHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.ExternalEndpoint)
	if !ok {
		return errors.Errorf("internal error: ExternalEndpoint handler received event for %T", object)
	}
	return h.handler.GenericExternalEndpoint(obj)
}

// Handle events for the RouteTable Resource
// DEPRECATED: Prefer reconciler pattern.
type RouteTableEventHandler interface {
	CreateRouteTable(obj *networking_gloo_solo_io_v2.RouteTable) error
	UpdateRouteTable(old, new *networking_gloo_solo_io_v2.RouteTable) error
	DeleteRouteTable(obj *networking_gloo_solo_io_v2.RouteTable) error
	GenericRouteTable(obj *networking_gloo_solo_io_v2.RouteTable) error
}

type RouteTableEventHandlerFuncs struct {
	OnCreate  func(obj *networking_gloo_solo_io_v2.RouteTable) error
	OnUpdate  func(old, new *networking_gloo_solo_io_v2.RouteTable) error
	OnDelete  func(obj *networking_gloo_solo_io_v2.RouteTable) error
	OnGeneric func(obj *networking_gloo_solo_io_v2.RouteTable) error
}

func (f *RouteTableEventHandlerFuncs) CreateRouteTable(obj *networking_gloo_solo_io_v2.RouteTable) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RouteTableEventHandlerFuncs) DeleteRouteTable(obj *networking_gloo_solo_io_v2.RouteTable) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RouteTableEventHandlerFuncs) UpdateRouteTable(objOld, objNew *networking_gloo_solo_io_v2.RouteTable) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RouteTableEventHandlerFuncs) GenericRouteTable(obj *networking_gloo_solo_io_v2.RouteTable) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RouteTableEventWatcher interface {
	AddEventHandler(ctx context.Context, h RouteTableEventHandler, predicates ...predicate.Predicate) error
}

type routeTableEventWatcher struct {
	watcher events.EventWatcher
}

func NewRouteTableEventWatcher(name string, mgr manager.Manager) RouteTableEventWatcher {
	return &routeTableEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_gloo_solo_io_v2.RouteTable{}),
	}
}

func (c *routeTableEventWatcher) AddEventHandler(ctx context.Context, h RouteTableEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRouteTableHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRouteTableHandler implements a generic events.EventHandler
type genericRouteTableHandler struct {
	handler RouteTableEventHandler
}

func (h genericRouteTableHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return h.handler.CreateRouteTable(obj)
}

func (h genericRouteTableHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return h.handler.DeleteRouteTable(obj)
}

func (h genericRouteTableHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_gloo_solo_io_v2.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", old)
	}
	objNew, ok := new.(*networking_gloo_solo_io_v2.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", new)
	}
	return h.handler.UpdateRouteTable(objOld, objNew)
}

func (h genericRouteTableHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.RouteTable)
	if !ok {
		return errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return h.handler.GenericRouteTable(obj)
}

// Handle events for the VirtualDestination Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualDestinationEventHandler interface {
	CreateVirtualDestination(obj *networking_gloo_solo_io_v2.VirtualDestination) error
	UpdateVirtualDestination(old, new *networking_gloo_solo_io_v2.VirtualDestination) error
	DeleteVirtualDestination(obj *networking_gloo_solo_io_v2.VirtualDestination) error
	GenericVirtualDestination(obj *networking_gloo_solo_io_v2.VirtualDestination) error
}

type VirtualDestinationEventHandlerFuncs struct {
	OnCreate  func(obj *networking_gloo_solo_io_v2.VirtualDestination) error
	OnUpdate  func(old, new *networking_gloo_solo_io_v2.VirtualDestination) error
	OnDelete  func(obj *networking_gloo_solo_io_v2.VirtualDestination) error
	OnGeneric func(obj *networking_gloo_solo_io_v2.VirtualDestination) error
}

func (f *VirtualDestinationEventHandlerFuncs) CreateVirtualDestination(obj *networking_gloo_solo_io_v2.VirtualDestination) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *VirtualDestinationEventHandlerFuncs) DeleteVirtualDestination(obj *networking_gloo_solo_io_v2.VirtualDestination) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *VirtualDestinationEventHandlerFuncs) UpdateVirtualDestination(objOld, objNew *networking_gloo_solo_io_v2.VirtualDestination) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *VirtualDestinationEventHandlerFuncs) GenericVirtualDestination(obj *networking_gloo_solo_io_v2.VirtualDestination) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type VirtualDestinationEventWatcher interface {
	AddEventHandler(ctx context.Context, h VirtualDestinationEventHandler, predicates ...predicate.Predicate) error
}

type virtualDestinationEventWatcher struct {
	watcher events.EventWatcher
}

func NewVirtualDestinationEventWatcher(name string, mgr manager.Manager) VirtualDestinationEventWatcher {
	return &virtualDestinationEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_gloo_solo_io_v2.VirtualDestination{}),
	}
}

func (c *virtualDestinationEventWatcher) AddEventHandler(ctx context.Context, h VirtualDestinationEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualDestinationHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericVirtualDestinationHandler implements a generic events.EventHandler
type genericVirtualDestinationHandler struct {
	handler VirtualDestinationEventHandler
}

func (h genericVirtualDestinationHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualDestination)
	if !ok {
		return errors.Errorf("internal error: VirtualDestination handler received event for %T", object)
	}
	return h.handler.CreateVirtualDestination(obj)
}

func (h genericVirtualDestinationHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualDestination)
	if !ok {
		return errors.Errorf("internal error: VirtualDestination handler received event for %T", object)
	}
	return h.handler.DeleteVirtualDestination(obj)
}

func (h genericVirtualDestinationHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_gloo_solo_io_v2.VirtualDestination)
	if !ok {
		return errors.Errorf("internal error: VirtualDestination handler received event for %T", old)
	}
	objNew, ok := new.(*networking_gloo_solo_io_v2.VirtualDestination)
	if !ok {
		return errors.Errorf("internal error: VirtualDestination handler received event for %T", new)
	}
	return h.handler.UpdateVirtualDestination(objOld, objNew)
}

func (h genericVirtualDestinationHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualDestination)
	if !ok {
		return errors.Errorf("internal error: VirtualDestination handler received event for %T", object)
	}
	return h.handler.GenericVirtualDestination(obj)
}

// Handle events for the VirtualGateway Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualGatewayEventHandler interface {
	CreateVirtualGateway(obj *networking_gloo_solo_io_v2.VirtualGateway) error
	UpdateVirtualGateway(old, new *networking_gloo_solo_io_v2.VirtualGateway) error
	DeleteVirtualGateway(obj *networking_gloo_solo_io_v2.VirtualGateway) error
	GenericVirtualGateway(obj *networking_gloo_solo_io_v2.VirtualGateway) error
}

type VirtualGatewayEventHandlerFuncs struct {
	OnCreate  func(obj *networking_gloo_solo_io_v2.VirtualGateway) error
	OnUpdate  func(old, new *networking_gloo_solo_io_v2.VirtualGateway) error
	OnDelete  func(obj *networking_gloo_solo_io_v2.VirtualGateway) error
	OnGeneric func(obj *networking_gloo_solo_io_v2.VirtualGateway) error
}

func (f *VirtualGatewayEventHandlerFuncs) CreateVirtualGateway(obj *networking_gloo_solo_io_v2.VirtualGateway) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *VirtualGatewayEventHandlerFuncs) DeleteVirtualGateway(obj *networking_gloo_solo_io_v2.VirtualGateway) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *VirtualGatewayEventHandlerFuncs) UpdateVirtualGateway(objOld, objNew *networking_gloo_solo_io_v2.VirtualGateway) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *VirtualGatewayEventHandlerFuncs) GenericVirtualGateway(obj *networking_gloo_solo_io_v2.VirtualGateway) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type VirtualGatewayEventWatcher interface {
	AddEventHandler(ctx context.Context, h VirtualGatewayEventHandler, predicates ...predicate.Predicate) error
}

type virtualGatewayEventWatcher struct {
	watcher events.EventWatcher
}

func NewVirtualGatewayEventWatcher(name string, mgr manager.Manager) VirtualGatewayEventWatcher {
	return &virtualGatewayEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_gloo_solo_io_v2.VirtualGateway{}),
	}
}

func (c *virtualGatewayEventWatcher) AddEventHandler(ctx context.Context, h VirtualGatewayEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualGatewayHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericVirtualGatewayHandler implements a generic events.EventHandler
type genericVirtualGatewayHandler struct {
	handler VirtualGatewayEventHandler
}

func (h genericVirtualGatewayHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return h.handler.CreateVirtualGateway(obj)
}

func (h genericVirtualGatewayHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return h.handler.DeleteVirtualGateway(obj)
}

func (h genericVirtualGatewayHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_gloo_solo_io_v2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", old)
	}
	objNew, ok := new.(*networking_gloo_solo_io_v2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", new)
	}
	return h.handler.UpdateVirtualGateway(objOld, objNew)
}

func (h genericVirtualGatewayHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_gloo_solo_io_v2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return h.handler.GenericVirtualGateway(obj)
}
