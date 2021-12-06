// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	fed_gateway_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/fed.gateway.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the FederatedGateway Resource
// DEPRECATED: Prefer reconciler pattern.
type FederatedGatewayEventHandler interface {
	CreateFederatedGateway(obj *fed_gateway_solo_io_v1.FederatedGateway) error
	UpdateFederatedGateway(old, new *fed_gateway_solo_io_v1.FederatedGateway) error
	DeleteFederatedGateway(obj *fed_gateway_solo_io_v1.FederatedGateway) error
	GenericFederatedGateway(obj *fed_gateway_solo_io_v1.FederatedGateway) error
}

type FederatedGatewayEventHandlerFuncs struct {
	OnCreate  func(obj *fed_gateway_solo_io_v1.FederatedGateway) error
	OnUpdate  func(old, new *fed_gateway_solo_io_v1.FederatedGateway) error
	OnDelete  func(obj *fed_gateway_solo_io_v1.FederatedGateway) error
	OnGeneric func(obj *fed_gateway_solo_io_v1.FederatedGateway) error
}

func (f *FederatedGatewayEventHandlerFuncs) CreateFederatedGateway(obj *fed_gateway_solo_io_v1.FederatedGateway) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *FederatedGatewayEventHandlerFuncs) DeleteFederatedGateway(obj *fed_gateway_solo_io_v1.FederatedGateway) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *FederatedGatewayEventHandlerFuncs) UpdateFederatedGateway(objOld, objNew *fed_gateway_solo_io_v1.FederatedGateway) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *FederatedGatewayEventHandlerFuncs) GenericFederatedGateway(obj *fed_gateway_solo_io_v1.FederatedGateway) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type FederatedGatewayEventWatcher interface {
	AddEventHandler(ctx context.Context, h FederatedGatewayEventHandler, predicates ...predicate.Predicate) error
}

type federatedGatewayEventWatcher struct {
	watcher events.EventWatcher
}

func NewFederatedGatewayEventWatcher(name string, mgr manager.Manager) FederatedGatewayEventWatcher {
	return &federatedGatewayEventWatcher{
		watcher: events.NewWatcher(name, mgr, &fed_gateway_solo_io_v1.FederatedGateway{}),
	}
}

func (c *federatedGatewayEventWatcher) AddEventHandler(ctx context.Context, h FederatedGatewayEventHandler, predicates ...predicate.Predicate) error {
	handler := genericFederatedGatewayHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericFederatedGatewayHandler implements a generic events.EventHandler
type genericFederatedGatewayHandler struct {
	handler FederatedGatewayEventHandler
}

func (h genericFederatedGatewayHandler) Create(object client.Object) error {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedGateway)
	if !ok {
		return errors.Errorf("internal error: FederatedGateway handler received event for %T", object)
	}
	return h.handler.CreateFederatedGateway(obj)
}

func (h genericFederatedGatewayHandler) Delete(object client.Object) error {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedGateway)
	if !ok {
		return errors.Errorf("internal error: FederatedGateway handler received event for %T", object)
	}
	return h.handler.DeleteFederatedGateway(obj)
}

func (h genericFederatedGatewayHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*fed_gateway_solo_io_v1.FederatedGateway)
	if !ok {
		return errors.Errorf("internal error: FederatedGateway handler received event for %T", old)
	}
	objNew, ok := new.(*fed_gateway_solo_io_v1.FederatedGateway)
	if !ok {
		return errors.Errorf("internal error: FederatedGateway handler received event for %T", new)
	}
	return h.handler.UpdateFederatedGateway(objOld, objNew)
}

func (h genericFederatedGatewayHandler) Generic(object client.Object) error {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedGateway)
	if !ok {
		return errors.Errorf("internal error: FederatedGateway handler received event for %T", object)
	}
	return h.handler.GenericFederatedGateway(obj)
}

// Handle events for the FederatedRouteTable Resource
// DEPRECATED: Prefer reconciler pattern.
type FederatedRouteTableEventHandler interface {
	CreateFederatedRouteTable(obj *fed_gateway_solo_io_v1.FederatedRouteTable) error
	UpdateFederatedRouteTable(old, new *fed_gateway_solo_io_v1.FederatedRouteTable) error
	DeleteFederatedRouteTable(obj *fed_gateway_solo_io_v1.FederatedRouteTable) error
	GenericFederatedRouteTable(obj *fed_gateway_solo_io_v1.FederatedRouteTable) error
}

type FederatedRouteTableEventHandlerFuncs struct {
	OnCreate  func(obj *fed_gateway_solo_io_v1.FederatedRouteTable) error
	OnUpdate  func(old, new *fed_gateway_solo_io_v1.FederatedRouteTable) error
	OnDelete  func(obj *fed_gateway_solo_io_v1.FederatedRouteTable) error
	OnGeneric func(obj *fed_gateway_solo_io_v1.FederatedRouteTable) error
}

func (f *FederatedRouteTableEventHandlerFuncs) CreateFederatedRouteTable(obj *fed_gateway_solo_io_v1.FederatedRouteTable) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *FederatedRouteTableEventHandlerFuncs) DeleteFederatedRouteTable(obj *fed_gateway_solo_io_v1.FederatedRouteTable) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *FederatedRouteTableEventHandlerFuncs) UpdateFederatedRouteTable(objOld, objNew *fed_gateway_solo_io_v1.FederatedRouteTable) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *FederatedRouteTableEventHandlerFuncs) GenericFederatedRouteTable(obj *fed_gateway_solo_io_v1.FederatedRouteTable) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type FederatedRouteTableEventWatcher interface {
	AddEventHandler(ctx context.Context, h FederatedRouteTableEventHandler, predicates ...predicate.Predicate) error
}

type federatedRouteTableEventWatcher struct {
	watcher events.EventWatcher
}

func NewFederatedRouteTableEventWatcher(name string, mgr manager.Manager) FederatedRouteTableEventWatcher {
	return &federatedRouteTableEventWatcher{
		watcher: events.NewWatcher(name, mgr, &fed_gateway_solo_io_v1.FederatedRouteTable{}),
	}
}

func (c *federatedRouteTableEventWatcher) AddEventHandler(ctx context.Context, h FederatedRouteTableEventHandler, predicates ...predicate.Predicate) error {
	handler := genericFederatedRouteTableHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericFederatedRouteTableHandler implements a generic events.EventHandler
type genericFederatedRouteTableHandler struct {
	handler FederatedRouteTableEventHandler
}

func (h genericFederatedRouteTableHandler) Create(object client.Object) error {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedRouteTable)
	if !ok {
		return errors.Errorf("internal error: FederatedRouteTable handler received event for %T", object)
	}
	return h.handler.CreateFederatedRouteTable(obj)
}

func (h genericFederatedRouteTableHandler) Delete(object client.Object) error {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedRouteTable)
	if !ok {
		return errors.Errorf("internal error: FederatedRouteTable handler received event for %T", object)
	}
	return h.handler.DeleteFederatedRouteTable(obj)
}

func (h genericFederatedRouteTableHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*fed_gateway_solo_io_v1.FederatedRouteTable)
	if !ok {
		return errors.Errorf("internal error: FederatedRouteTable handler received event for %T", old)
	}
	objNew, ok := new.(*fed_gateway_solo_io_v1.FederatedRouteTable)
	if !ok {
		return errors.Errorf("internal error: FederatedRouteTable handler received event for %T", new)
	}
	return h.handler.UpdateFederatedRouteTable(objOld, objNew)
}

func (h genericFederatedRouteTableHandler) Generic(object client.Object) error {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedRouteTable)
	if !ok {
		return errors.Errorf("internal error: FederatedRouteTable handler received event for %T", object)
	}
	return h.handler.GenericFederatedRouteTable(obj)
}

// Handle events for the FederatedVirtualService Resource
// DEPRECATED: Prefer reconciler pattern.
type FederatedVirtualServiceEventHandler interface {
	CreateFederatedVirtualService(obj *fed_gateway_solo_io_v1.FederatedVirtualService) error
	UpdateFederatedVirtualService(old, new *fed_gateway_solo_io_v1.FederatedVirtualService) error
	DeleteFederatedVirtualService(obj *fed_gateway_solo_io_v1.FederatedVirtualService) error
	GenericFederatedVirtualService(obj *fed_gateway_solo_io_v1.FederatedVirtualService) error
}

type FederatedVirtualServiceEventHandlerFuncs struct {
	OnCreate  func(obj *fed_gateway_solo_io_v1.FederatedVirtualService) error
	OnUpdate  func(old, new *fed_gateway_solo_io_v1.FederatedVirtualService) error
	OnDelete  func(obj *fed_gateway_solo_io_v1.FederatedVirtualService) error
	OnGeneric func(obj *fed_gateway_solo_io_v1.FederatedVirtualService) error
}

func (f *FederatedVirtualServiceEventHandlerFuncs) CreateFederatedVirtualService(obj *fed_gateway_solo_io_v1.FederatedVirtualService) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *FederatedVirtualServiceEventHandlerFuncs) DeleteFederatedVirtualService(obj *fed_gateway_solo_io_v1.FederatedVirtualService) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *FederatedVirtualServiceEventHandlerFuncs) UpdateFederatedVirtualService(objOld, objNew *fed_gateway_solo_io_v1.FederatedVirtualService) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *FederatedVirtualServiceEventHandlerFuncs) GenericFederatedVirtualService(obj *fed_gateway_solo_io_v1.FederatedVirtualService) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type FederatedVirtualServiceEventWatcher interface {
	AddEventHandler(ctx context.Context, h FederatedVirtualServiceEventHandler, predicates ...predicate.Predicate) error
}

type federatedVirtualServiceEventWatcher struct {
	watcher events.EventWatcher
}

func NewFederatedVirtualServiceEventWatcher(name string, mgr manager.Manager) FederatedVirtualServiceEventWatcher {
	return &federatedVirtualServiceEventWatcher{
		watcher: events.NewWatcher(name, mgr, &fed_gateway_solo_io_v1.FederatedVirtualService{}),
	}
}

func (c *federatedVirtualServiceEventWatcher) AddEventHandler(ctx context.Context, h FederatedVirtualServiceEventHandler, predicates ...predicate.Predicate) error {
	handler := genericFederatedVirtualServiceHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericFederatedVirtualServiceHandler implements a generic events.EventHandler
type genericFederatedVirtualServiceHandler struct {
	handler FederatedVirtualServiceEventHandler
}

func (h genericFederatedVirtualServiceHandler) Create(object client.Object) error {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedVirtualService)
	if !ok {
		return errors.Errorf("internal error: FederatedVirtualService handler received event for %T", object)
	}
	return h.handler.CreateFederatedVirtualService(obj)
}

func (h genericFederatedVirtualServiceHandler) Delete(object client.Object) error {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedVirtualService)
	if !ok {
		return errors.Errorf("internal error: FederatedVirtualService handler received event for %T", object)
	}
	return h.handler.DeleteFederatedVirtualService(obj)
}

func (h genericFederatedVirtualServiceHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*fed_gateway_solo_io_v1.FederatedVirtualService)
	if !ok {
		return errors.Errorf("internal error: FederatedVirtualService handler received event for %T", old)
	}
	objNew, ok := new.(*fed_gateway_solo_io_v1.FederatedVirtualService)
	if !ok {
		return errors.Errorf("internal error: FederatedVirtualService handler received event for %T", new)
	}
	return h.handler.UpdateFederatedVirtualService(objOld, objNew)
}

func (h genericFederatedVirtualServiceHandler) Generic(object client.Object) error {
	obj, ok := object.(*fed_gateway_solo_io_v1.FederatedVirtualService)
	if !ok {
		return errors.Errorf("internal error: FederatedVirtualService handler received event for %T", object)
	}
	return h.handler.GenericFederatedVirtualService(obj)
}
