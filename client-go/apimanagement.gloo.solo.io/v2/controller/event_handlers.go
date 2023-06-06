// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	apimanagement_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/apimanagement.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the GraphQLStitchedSchema Resource
// DEPRECATED: Prefer reconciler pattern.
type GraphQLStitchedSchemaEventHandler interface {
	CreateGraphQLStitchedSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error
	UpdateGraphQLStitchedSchema(old, new *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error
	DeleteGraphQLStitchedSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error
	GenericGraphQLStitchedSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error
}

type GraphQLStitchedSchemaEventHandlerFuncs struct {
	OnCreate  func(obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error
	OnUpdate  func(old, new *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error
	OnDelete  func(obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error
	OnGeneric func(obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error
}

func (f *GraphQLStitchedSchemaEventHandlerFuncs) CreateGraphQLStitchedSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *GraphQLStitchedSchemaEventHandlerFuncs) DeleteGraphQLStitchedSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *GraphQLStitchedSchemaEventHandlerFuncs) UpdateGraphQLStitchedSchema(objOld, objNew *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *GraphQLStitchedSchemaEventHandlerFuncs) GenericGraphQLStitchedSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type GraphQLStitchedSchemaEventWatcher interface {
	AddEventHandler(ctx context.Context, h GraphQLStitchedSchemaEventHandler, predicates ...predicate.Predicate) error
}

type graphQLStitchedSchemaEventWatcher struct {
	watcher events.EventWatcher
}

func NewGraphQLStitchedSchemaEventWatcher(name string, mgr manager.Manager) GraphQLStitchedSchemaEventWatcher {
	return &graphQLStitchedSchemaEventWatcher{
		watcher: events.NewWatcher(name, mgr, &apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema{}),
	}
}

func (c *graphQLStitchedSchemaEventWatcher) AddEventHandler(ctx context.Context, h GraphQLStitchedSchemaEventHandler, predicates ...predicate.Predicate) error {
	handler := genericGraphQLStitchedSchemaHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericGraphQLStitchedSchemaHandler implements a generic events.EventHandler
type genericGraphQLStitchedSchemaHandler struct {
	handler GraphQLStitchedSchemaEventHandler
}

func (h genericGraphQLStitchedSchemaHandler) Create(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLStitchedSchema handler received event for %T", object)
	}
	return h.handler.CreateGraphQLStitchedSchema(obj)
}

func (h genericGraphQLStitchedSchemaHandler) Delete(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLStitchedSchema handler received event for %T", object)
	}
	return h.handler.DeleteGraphQLStitchedSchema(obj)
}

func (h genericGraphQLStitchedSchemaHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLStitchedSchema handler received event for %T", old)
	}
	objNew, ok := new.(*apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLStitchedSchema handler received event for %T", new)
	}
	return h.handler.UpdateGraphQLStitchedSchema(objOld, objNew)
}

func (h genericGraphQLStitchedSchemaHandler) Generic(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLStitchedSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLStitchedSchema handler received event for %T", object)
	}
	return h.handler.GenericGraphQLStitchedSchema(obj)
}

// Handle events for the GraphQLResolverMap Resource
// DEPRECATED: Prefer reconciler pattern.
type GraphQLResolverMapEventHandler interface {
	CreateGraphQLResolverMap(obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error
	UpdateGraphQLResolverMap(old, new *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error
	DeleteGraphQLResolverMap(obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error
	GenericGraphQLResolverMap(obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error
}

type GraphQLResolverMapEventHandlerFuncs struct {
	OnCreate  func(obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error
	OnUpdate  func(old, new *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error
	OnDelete  func(obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error
	OnGeneric func(obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error
}

func (f *GraphQLResolverMapEventHandlerFuncs) CreateGraphQLResolverMap(obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *GraphQLResolverMapEventHandlerFuncs) DeleteGraphQLResolverMap(obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *GraphQLResolverMapEventHandlerFuncs) UpdateGraphQLResolverMap(objOld, objNew *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *GraphQLResolverMapEventHandlerFuncs) GenericGraphQLResolverMap(obj *apimanagement_gloo_solo_io_v2.GraphQLResolverMap) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type GraphQLResolverMapEventWatcher interface {
	AddEventHandler(ctx context.Context, h GraphQLResolverMapEventHandler, predicates ...predicate.Predicate) error
}

type graphQLResolverMapEventWatcher struct {
	watcher events.EventWatcher
}

func NewGraphQLResolverMapEventWatcher(name string, mgr manager.Manager) GraphQLResolverMapEventWatcher {
	return &graphQLResolverMapEventWatcher{
		watcher: events.NewWatcher(name, mgr, &apimanagement_gloo_solo_io_v2.GraphQLResolverMap{}),
	}
}

func (c *graphQLResolverMapEventWatcher) AddEventHandler(ctx context.Context, h GraphQLResolverMapEventHandler, predicates ...predicate.Predicate) error {
	handler := genericGraphQLResolverMapHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericGraphQLResolverMapHandler implements a generic events.EventHandler
type genericGraphQLResolverMapHandler struct {
	handler GraphQLResolverMapEventHandler
}

func (h genericGraphQLResolverMapHandler) Create(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLResolverMap)
	if !ok {
		return errors.Errorf("internal error: GraphQLResolverMap handler received event for %T", object)
	}
	return h.handler.CreateGraphQLResolverMap(obj)
}

func (h genericGraphQLResolverMapHandler) Delete(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLResolverMap)
	if !ok {
		return errors.Errorf("internal error: GraphQLResolverMap handler received event for %T", object)
	}
	return h.handler.DeleteGraphQLResolverMap(obj)
}

func (h genericGraphQLResolverMapHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*apimanagement_gloo_solo_io_v2.GraphQLResolverMap)
	if !ok {
		return errors.Errorf("internal error: GraphQLResolverMap handler received event for %T", old)
	}
	objNew, ok := new.(*apimanagement_gloo_solo_io_v2.GraphQLResolverMap)
	if !ok {
		return errors.Errorf("internal error: GraphQLResolverMap handler received event for %T", new)
	}
	return h.handler.UpdateGraphQLResolverMap(objOld, objNew)
}

func (h genericGraphQLResolverMapHandler) Generic(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLResolverMap)
	if !ok {
		return errors.Errorf("internal error: GraphQLResolverMap handler received event for %T", object)
	}
	return h.handler.GenericGraphQLResolverMap(obj)
}

// Handle events for the GraphQLSchema Resource
// DEPRECATED: Prefer reconciler pattern.
type GraphQLSchemaEventHandler interface {
	CreateGraphQLSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) error
	UpdateGraphQLSchema(old, new *apimanagement_gloo_solo_io_v2.GraphQLSchema) error
	DeleteGraphQLSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) error
	GenericGraphQLSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) error
}

type GraphQLSchemaEventHandlerFuncs struct {
	OnCreate  func(obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) error
	OnUpdate  func(old, new *apimanagement_gloo_solo_io_v2.GraphQLSchema) error
	OnDelete  func(obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) error
	OnGeneric func(obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) error
}

func (f *GraphQLSchemaEventHandlerFuncs) CreateGraphQLSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *GraphQLSchemaEventHandlerFuncs) DeleteGraphQLSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *GraphQLSchemaEventHandlerFuncs) UpdateGraphQLSchema(objOld, objNew *apimanagement_gloo_solo_io_v2.GraphQLSchema) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *GraphQLSchemaEventHandlerFuncs) GenericGraphQLSchema(obj *apimanagement_gloo_solo_io_v2.GraphQLSchema) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type GraphQLSchemaEventWatcher interface {
	AddEventHandler(ctx context.Context, h GraphQLSchemaEventHandler, predicates ...predicate.Predicate) error
}

type graphQLSchemaEventWatcher struct {
	watcher events.EventWatcher
}

func NewGraphQLSchemaEventWatcher(name string, mgr manager.Manager) GraphQLSchemaEventWatcher {
	return &graphQLSchemaEventWatcher{
		watcher: events.NewWatcher(name, mgr, &apimanagement_gloo_solo_io_v2.GraphQLSchema{}),
	}
}

func (c *graphQLSchemaEventWatcher) AddEventHandler(ctx context.Context, h GraphQLSchemaEventHandler, predicates ...predicate.Predicate) error {
	handler := genericGraphQLSchemaHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericGraphQLSchemaHandler implements a generic events.EventHandler
type genericGraphQLSchemaHandler struct {
	handler GraphQLSchemaEventHandler
}

func (h genericGraphQLSchemaHandler) Create(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLSchema handler received event for %T", object)
	}
	return h.handler.CreateGraphQLSchema(obj)
}

func (h genericGraphQLSchemaHandler) Delete(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLSchema handler received event for %T", object)
	}
	return h.handler.DeleteGraphQLSchema(obj)
}

func (h genericGraphQLSchemaHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*apimanagement_gloo_solo_io_v2.GraphQLSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLSchema handler received event for %T", old)
	}
	objNew, ok := new.(*apimanagement_gloo_solo_io_v2.GraphQLSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLSchema handler received event for %T", new)
	}
	return h.handler.UpdateGraphQLSchema(objOld, objNew)
}

func (h genericGraphQLSchemaHandler) Generic(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.GraphQLSchema)
	if !ok {
		return errors.Errorf("internal error: GraphQLSchema handler received event for %T", object)
	}
	return h.handler.GenericGraphQLSchema(obj)
}

// Handle events for the ApiDoc Resource
// DEPRECATED: Prefer reconciler pattern.
type ApiDocEventHandler interface {
	CreateApiDoc(obj *apimanagement_gloo_solo_io_v2.ApiDoc) error
	UpdateApiDoc(old, new *apimanagement_gloo_solo_io_v2.ApiDoc) error
	DeleteApiDoc(obj *apimanagement_gloo_solo_io_v2.ApiDoc) error
	GenericApiDoc(obj *apimanagement_gloo_solo_io_v2.ApiDoc) error
}

type ApiDocEventHandlerFuncs struct {
	OnCreate  func(obj *apimanagement_gloo_solo_io_v2.ApiDoc) error
	OnUpdate  func(old, new *apimanagement_gloo_solo_io_v2.ApiDoc) error
	OnDelete  func(obj *apimanagement_gloo_solo_io_v2.ApiDoc) error
	OnGeneric func(obj *apimanagement_gloo_solo_io_v2.ApiDoc) error
}

func (f *ApiDocEventHandlerFuncs) CreateApiDoc(obj *apimanagement_gloo_solo_io_v2.ApiDoc) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ApiDocEventHandlerFuncs) DeleteApiDoc(obj *apimanagement_gloo_solo_io_v2.ApiDoc) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ApiDocEventHandlerFuncs) UpdateApiDoc(objOld, objNew *apimanagement_gloo_solo_io_v2.ApiDoc) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ApiDocEventHandlerFuncs) GenericApiDoc(obj *apimanagement_gloo_solo_io_v2.ApiDoc) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ApiDocEventWatcher interface {
	AddEventHandler(ctx context.Context, h ApiDocEventHandler, predicates ...predicate.Predicate) error
}

type apiDocEventWatcher struct {
	watcher events.EventWatcher
}

func NewApiDocEventWatcher(name string, mgr manager.Manager) ApiDocEventWatcher {
	return &apiDocEventWatcher{
		watcher: events.NewWatcher(name, mgr, &apimanagement_gloo_solo_io_v2.ApiDoc{}),
	}
}

func (c *apiDocEventWatcher) AddEventHandler(ctx context.Context, h ApiDocEventHandler, predicates ...predicate.Predicate) error {
	handler := genericApiDocHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericApiDocHandler implements a generic events.EventHandler
type genericApiDocHandler struct {
	handler ApiDocEventHandler
}

func (h genericApiDocHandler) Create(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.ApiDoc)
	if !ok {
		return errors.Errorf("internal error: ApiDoc handler received event for %T", object)
	}
	return h.handler.CreateApiDoc(obj)
}

func (h genericApiDocHandler) Delete(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.ApiDoc)
	if !ok {
		return errors.Errorf("internal error: ApiDoc handler received event for %T", object)
	}
	return h.handler.DeleteApiDoc(obj)
}

func (h genericApiDocHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*apimanagement_gloo_solo_io_v2.ApiDoc)
	if !ok {
		return errors.Errorf("internal error: ApiDoc handler received event for %T", old)
	}
	objNew, ok := new.(*apimanagement_gloo_solo_io_v2.ApiDoc)
	if !ok {
		return errors.Errorf("internal error: ApiDoc handler received event for %T", new)
	}
	return h.handler.UpdateApiDoc(objOld, objNew)
}

func (h genericApiDocHandler) Generic(object client.Object) error {
	obj, ok := object.(*apimanagement_gloo_solo_io_v2.ApiDoc)
	if !ok {
		return errors.Errorf("internal error: ApiDoc handler received event for %T", object)
	}
	return h.handler.GenericApiDoc(obj)
}
