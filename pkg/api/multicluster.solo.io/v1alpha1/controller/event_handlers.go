// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	multicluster_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the MultiClusterRole Resource
// DEPRECATED: Prefer reconciler pattern.
type MultiClusterRoleEventHandler interface {
	CreateMultiClusterRole(obj *multicluster_solo_io_v1alpha1.MultiClusterRole) error
	UpdateMultiClusterRole(old, new *multicluster_solo_io_v1alpha1.MultiClusterRole) error
	DeleteMultiClusterRole(obj *multicluster_solo_io_v1alpha1.MultiClusterRole) error
	GenericMultiClusterRole(obj *multicluster_solo_io_v1alpha1.MultiClusterRole) error
}

type MultiClusterRoleEventHandlerFuncs struct {
	OnCreate  func(obj *multicluster_solo_io_v1alpha1.MultiClusterRole) error
	OnUpdate  func(old, new *multicluster_solo_io_v1alpha1.MultiClusterRole) error
	OnDelete  func(obj *multicluster_solo_io_v1alpha1.MultiClusterRole) error
	OnGeneric func(obj *multicluster_solo_io_v1alpha1.MultiClusterRole) error
}

func (f *MultiClusterRoleEventHandlerFuncs) CreateMultiClusterRole(obj *multicluster_solo_io_v1alpha1.MultiClusterRole) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *MultiClusterRoleEventHandlerFuncs) DeleteMultiClusterRole(obj *multicluster_solo_io_v1alpha1.MultiClusterRole) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *MultiClusterRoleEventHandlerFuncs) UpdateMultiClusterRole(objOld, objNew *multicluster_solo_io_v1alpha1.MultiClusterRole) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *MultiClusterRoleEventHandlerFuncs) GenericMultiClusterRole(obj *multicluster_solo_io_v1alpha1.MultiClusterRole) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type MultiClusterRoleEventWatcher interface {
	AddEventHandler(ctx context.Context, h MultiClusterRoleEventHandler, predicates ...predicate.Predicate) error
}

type multiClusterRoleEventWatcher struct {
	watcher events.EventWatcher
}

func NewMultiClusterRoleEventWatcher(name string, mgr manager.Manager) MultiClusterRoleEventWatcher {
	return &multiClusterRoleEventWatcher{
		watcher: events.NewWatcher(name, mgr, &multicluster_solo_io_v1alpha1.MultiClusterRole{}),
	}
}

func (c *multiClusterRoleEventWatcher) AddEventHandler(ctx context.Context, h MultiClusterRoleEventHandler, predicates ...predicate.Predicate) error {
	handler := genericMultiClusterRoleHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericMultiClusterRoleHandler implements a generic events.EventHandler
type genericMultiClusterRoleHandler struct {
	handler MultiClusterRoleEventHandler
}

func (h genericMultiClusterRoleHandler) Create(object runtime.Object) error {
	obj, ok := object.(*multicluster_solo_io_v1alpha1.MultiClusterRole)
	if !ok {
		return errors.Errorf("internal error: MultiClusterRole handler received event for %T", object)
	}
	return h.handler.CreateMultiClusterRole(obj)
}

func (h genericMultiClusterRoleHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*multicluster_solo_io_v1alpha1.MultiClusterRole)
	if !ok {
		return errors.Errorf("internal error: MultiClusterRole handler received event for %T", object)
	}
	return h.handler.DeleteMultiClusterRole(obj)
}

func (h genericMultiClusterRoleHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*multicluster_solo_io_v1alpha1.MultiClusterRole)
	if !ok {
		return errors.Errorf("internal error: MultiClusterRole handler received event for %T", old)
	}
	objNew, ok := new.(*multicluster_solo_io_v1alpha1.MultiClusterRole)
	if !ok {
		return errors.Errorf("internal error: MultiClusterRole handler received event for %T", new)
	}
	return h.handler.UpdateMultiClusterRole(objOld, objNew)
}

func (h genericMultiClusterRoleHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*multicluster_solo_io_v1alpha1.MultiClusterRole)
	if !ok {
		return errors.Errorf("internal error: MultiClusterRole handler received event for %T", object)
	}
	return h.handler.GenericMultiClusterRole(obj)
}
