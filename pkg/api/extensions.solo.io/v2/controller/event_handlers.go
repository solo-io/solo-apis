// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	extensions_solo_io_v2 "github.com/solo-io/solo-apis/pkg/api/extensions.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the WasmDeploymentPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type WasmDeploymentPolicyEventHandler interface {
	CreateWasmDeploymentPolicy(obj *extensions_solo_io_v2.WasmDeploymentPolicy) error
	UpdateWasmDeploymentPolicy(old, new *extensions_solo_io_v2.WasmDeploymentPolicy) error
	DeleteWasmDeploymentPolicy(obj *extensions_solo_io_v2.WasmDeploymentPolicy) error
	GenericWasmDeploymentPolicy(obj *extensions_solo_io_v2.WasmDeploymentPolicy) error
}

type WasmDeploymentPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *extensions_solo_io_v2.WasmDeploymentPolicy) error
	OnUpdate  func(old, new *extensions_solo_io_v2.WasmDeploymentPolicy) error
	OnDelete  func(obj *extensions_solo_io_v2.WasmDeploymentPolicy) error
	OnGeneric func(obj *extensions_solo_io_v2.WasmDeploymentPolicy) error
}

func (f *WasmDeploymentPolicyEventHandlerFuncs) CreateWasmDeploymentPolicy(obj *extensions_solo_io_v2.WasmDeploymentPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *WasmDeploymentPolicyEventHandlerFuncs) DeleteWasmDeploymentPolicy(obj *extensions_solo_io_v2.WasmDeploymentPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *WasmDeploymentPolicyEventHandlerFuncs) UpdateWasmDeploymentPolicy(objOld, objNew *extensions_solo_io_v2.WasmDeploymentPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *WasmDeploymentPolicyEventHandlerFuncs) GenericWasmDeploymentPolicy(obj *extensions_solo_io_v2.WasmDeploymentPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type WasmDeploymentPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h WasmDeploymentPolicyEventHandler, predicates ...predicate.Predicate) error
}

type wasmDeploymentPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewWasmDeploymentPolicyEventWatcher(name string, mgr manager.Manager) WasmDeploymentPolicyEventWatcher {
	return &wasmDeploymentPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &extensions_solo_io_v2.WasmDeploymentPolicy{}),
	}
}

func (c *wasmDeploymentPolicyEventWatcher) AddEventHandler(ctx context.Context, h WasmDeploymentPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericWasmDeploymentPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericWasmDeploymentPolicyHandler implements a generic events.EventHandler
type genericWasmDeploymentPolicyHandler struct {
	handler WasmDeploymentPolicyEventHandler
}

func (h genericWasmDeploymentPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*extensions_solo_io_v2.WasmDeploymentPolicy)
	if !ok {
		return errors.Errorf("internal error: WasmDeploymentPolicy handler received event for %T", object)
	}
	return h.handler.CreateWasmDeploymentPolicy(obj)
}

func (h genericWasmDeploymentPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*extensions_solo_io_v2.WasmDeploymentPolicy)
	if !ok {
		return errors.Errorf("internal error: WasmDeploymentPolicy handler received event for %T", object)
	}
	return h.handler.DeleteWasmDeploymentPolicy(obj)
}

func (h genericWasmDeploymentPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*extensions_solo_io_v2.WasmDeploymentPolicy)
	if !ok {
		return errors.Errorf("internal error: WasmDeploymentPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*extensions_solo_io_v2.WasmDeploymentPolicy)
	if !ok {
		return errors.Errorf("internal error: WasmDeploymentPolicy handler received event for %T", new)
	}
	return h.handler.UpdateWasmDeploymentPolicy(objOld, objNew)
}

func (h genericWasmDeploymentPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*extensions_solo_io_v2.WasmDeploymentPolicy)
	if !ok {
		return errors.Errorf("internal error: WasmDeploymentPolicy handler received event for %T", object)
	}
	return h.handler.GenericWasmDeploymentPolicy(obj)
}
