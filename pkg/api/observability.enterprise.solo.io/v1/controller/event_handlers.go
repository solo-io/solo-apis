// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	observability_enterprise_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/observability.enterprise.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the AccessLogRecord Resource
// DEPRECATED: Prefer reconciler pattern.
type AccessLogRecordEventHandler interface {
	CreateAccessLogRecord(obj *observability_enterprise_solo_io_v1.AccessLogRecord) error
	UpdateAccessLogRecord(old, new *observability_enterprise_solo_io_v1.AccessLogRecord) error
	DeleteAccessLogRecord(obj *observability_enterprise_solo_io_v1.AccessLogRecord) error
	GenericAccessLogRecord(obj *observability_enterprise_solo_io_v1.AccessLogRecord) error
}

type AccessLogRecordEventHandlerFuncs struct {
	OnCreate  func(obj *observability_enterprise_solo_io_v1.AccessLogRecord) error
	OnUpdate  func(old, new *observability_enterprise_solo_io_v1.AccessLogRecord) error
	OnDelete  func(obj *observability_enterprise_solo_io_v1.AccessLogRecord) error
	OnGeneric func(obj *observability_enterprise_solo_io_v1.AccessLogRecord) error
}

func (f *AccessLogRecordEventHandlerFuncs) CreateAccessLogRecord(obj *observability_enterprise_solo_io_v1.AccessLogRecord) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *AccessLogRecordEventHandlerFuncs) DeleteAccessLogRecord(obj *observability_enterprise_solo_io_v1.AccessLogRecord) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *AccessLogRecordEventHandlerFuncs) UpdateAccessLogRecord(objOld, objNew *observability_enterprise_solo_io_v1.AccessLogRecord) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *AccessLogRecordEventHandlerFuncs) GenericAccessLogRecord(obj *observability_enterprise_solo_io_v1.AccessLogRecord) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type AccessLogRecordEventWatcher interface {
	AddEventHandler(ctx context.Context, h AccessLogRecordEventHandler, predicates ...predicate.Predicate) error
}

type accessLogRecordEventWatcher struct {
	watcher events.EventWatcher
}

func NewAccessLogRecordEventWatcher(name string, mgr manager.Manager) AccessLogRecordEventWatcher {
	return &accessLogRecordEventWatcher{
		watcher: events.NewWatcher(name, mgr, &observability_enterprise_solo_io_v1.AccessLogRecord{}),
	}
}

func (c *accessLogRecordEventWatcher) AddEventHandler(ctx context.Context, h AccessLogRecordEventHandler, predicates ...predicate.Predicate) error {
	handler := genericAccessLogRecordHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericAccessLogRecordHandler implements a generic events.EventHandler
type genericAccessLogRecordHandler struct {
	handler AccessLogRecordEventHandler
}

func (h genericAccessLogRecordHandler) Create(object client.Object) error {
	obj, ok := object.(*observability_enterprise_solo_io_v1.AccessLogRecord)
	if !ok {
		return errors.Errorf("internal error: AccessLogRecord handler received event for %T", object)
	}
	return h.handler.CreateAccessLogRecord(obj)
}

func (h genericAccessLogRecordHandler) Delete(object client.Object) error {
	obj, ok := object.(*observability_enterprise_solo_io_v1.AccessLogRecord)
	if !ok {
		return errors.Errorf("internal error: AccessLogRecord handler received event for %T", object)
	}
	return h.handler.DeleteAccessLogRecord(obj)
}

func (h genericAccessLogRecordHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*observability_enterprise_solo_io_v1.AccessLogRecord)
	if !ok {
		return errors.Errorf("internal error: AccessLogRecord handler received event for %T", old)
	}
	objNew, ok := new.(*observability_enterprise_solo_io_v1.AccessLogRecord)
	if !ok {
		return errors.Errorf("internal error: AccessLogRecord handler received event for %T", new)
	}
	return h.handler.UpdateAccessLogRecord(objOld, objNew)
}

func (h genericAccessLogRecordHandler) Generic(object client.Object) error {
	obj, ok := object.(*observability_enterprise_solo_io_v1.AccessLogRecord)
	if !ok {
		return errors.Errorf("internal error: AccessLogRecord handler received event for %T", object)
	}
	return h.handler.GenericAccessLogRecord(obj)
}
