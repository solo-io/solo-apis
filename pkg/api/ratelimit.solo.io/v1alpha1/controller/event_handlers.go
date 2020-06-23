// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	ratelimit_solo_io_v1alpha1 "/pkg/api/ratelimit.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the RateLimitConfig Resource
// DEPRECATED: Prefer reconciler pattern.
type RateLimitConfigEventHandler interface {
	CreateRateLimitConfig(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error
	UpdateRateLimitConfig(old, new *ratelimit_solo_io_v1alpha1.RateLimitConfig) error
	DeleteRateLimitConfig(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error
	GenericRateLimitConfig(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error
}

type RateLimitConfigEventHandlerFuncs struct {
	OnCreate  func(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error
	OnUpdate  func(old, new *ratelimit_solo_io_v1alpha1.RateLimitConfig) error
	OnDelete  func(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error
	OnGeneric func(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error
}

func (f *RateLimitConfigEventHandlerFuncs) CreateRateLimitConfig(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RateLimitConfigEventHandlerFuncs) DeleteRateLimitConfig(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RateLimitConfigEventHandlerFuncs) UpdateRateLimitConfig(objOld, objNew *ratelimit_solo_io_v1alpha1.RateLimitConfig) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RateLimitConfigEventHandlerFuncs) GenericRateLimitConfig(obj *ratelimit_solo_io_v1alpha1.RateLimitConfig) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RateLimitConfigEventWatcher interface {
	AddEventHandler(ctx context.Context, h RateLimitConfigEventHandler, predicates ...predicate.Predicate) error
}

type rateLimitConfigEventWatcher struct {
	watcher events.EventWatcher
}

func NewRateLimitConfigEventWatcher(name string, mgr manager.Manager) RateLimitConfigEventWatcher {
	return &rateLimitConfigEventWatcher{
		watcher: events.NewWatcher(name, mgr, &ratelimit_solo_io_v1alpha1.RateLimitConfig{}),
	}
}

func (c *rateLimitConfigEventWatcher) AddEventHandler(ctx context.Context, h RateLimitConfigEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRateLimitConfigHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRateLimitConfigHandler implements a generic events.EventHandler
type genericRateLimitConfigHandler struct {
	handler RateLimitConfigEventHandler
}

func (h genericRateLimitConfigHandler) Create(object runtime.Object) error {
	obj, ok := object.(*ratelimit_solo_io_v1alpha1.RateLimitConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitConfig handler received event for %T", object)
	}
	return h.handler.CreateRateLimitConfig(obj)
}

func (h genericRateLimitConfigHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*ratelimit_solo_io_v1alpha1.RateLimitConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitConfig handler received event for %T", object)
	}
	return h.handler.DeleteRateLimitConfig(obj)
}

func (h genericRateLimitConfigHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*ratelimit_solo_io_v1alpha1.RateLimitConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitConfig handler received event for %T", old)
	}
	objNew, ok := new.(*ratelimit_solo_io_v1alpha1.RateLimitConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitConfig handler received event for %T", new)
	}
	return h.handler.UpdateRateLimitConfig(objOld, objNew)
}

func (h genericRateLimitConfigHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*ratelimit_solo_io_v1alpha1.RateLimitConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitConfig handler received event for %T", object)
	}
	return h.handler.GenericRateLimitConfig(obj)
}
