// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	security_policy_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/security.policy.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the AccessPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type AccessPolicyEventHandler interface {
	CreateAccessPolicy(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error
	UpdateAccessPolicy(old, new *security_policy_gloo_solo_io_v2.AccessPolicy) error
	DeleteAccessPolicy(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error
	GenericAccessPolicy(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error
}

type AccessPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error
	OnUpdate  func(old, new *security_policy_gloo_solo_io_v2.AccessPolicy) error
	OnDelete  func(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error
	OnGeneric func(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error
}

func (f *AccessPolicyEventHandlerFuncs) CreateAccessPolicy(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *AccessPolicyEventHandlerFuncs) DeleteAccessPolicy(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *AccessPolicyEventHandlerFuncs) UpdateAccessPolicy(objOld, objNew *security_policy_gloo_solo_io_v2.AccessPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *AccessPolicyEventHandlerFuncs) GenericAccessPolicy(obj *security_policy_gloo_solo_io_v2.AccessPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type AccessPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h AccessPolicyEventHandler, predicates ...predicate.Predicate) error
}

type accessPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewAccessPolicyEventWatcher(name string, mgr manager.Manager) AccessPolicyEventWatcher {
	return &accessPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_policy_gloo_solo_io_v2.AccessPolicy{}),
	}
}

func (c *accessPolicyEventWatcher) AddEventHandler(ctx context.Context, h AccessPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericAccessPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericAccessPolicyHandler implements a generic events.EventHandler
type genericAccessPolicyHandler struct {
	handler AccessPolicyEventHandler
}

func (h genericAccessPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.AccessPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
	}
	return h.handler.CreateAccessPolicy(obj)
}

func (h genericAccessPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.AccessPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
	}
	return h.handler.DeleteAccessPolicy(obj)
}

func (h genericAccessPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_policy_gloo_solo_io_v2.AccessPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*security_policy_gloo_solo_io_v2.AccessPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessPolicy handler received event for %T", new)
	}
	return h.handler.UpdateAccessPolicy(objOld, objNew)
}

func (h genericAccessPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.AccessPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
	}
	return h.handler.GenericAccessPolicy(obj)
}

// Handle events for the CORSPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type CORSPolicyEventHandler interface {
	CreateCORSPolicy(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error
	UpdateCORSPolicy(old, new *security_policy_gloo_solo_io_v2.CORSPolicy) error
	DeleteCORSPolicy(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error
	GenericCORSPolicy(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error
}

type CORSPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error
	OnUpdate  func(old, new *security_policy_gloo_solo_io_v2.CORSPolicy) error
	OnDelete  func(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error
	OnGeneric func(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error
}

func (f *CORSPolicyEventHandlerFuncs) CreateCORSPolicy(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *CORSPolicyEventHandlerFuncs) DeleteCORSPolicy(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *CORSPolicyEventHandlerFuncs) UpdateCORSPolicy(objOld, objNew *security_policy_gloo_solo_io_v2.CORSPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *CORSPolicyEventHandlerFuncs) GenericCORSPolicy(obj *security_policy_gloo_solo_io_v2.CORSPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type CORSPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h CORSPolicyEventHandler, predicates ...predicate.Predicate) error
}

type cORSPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewCORSPolicyEventWatcher(name string, mgr manager.Manager) CORSPolicyEventWatcher {
	return &cORSPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_policy_gloo_solo_io_v2.CORSPolicy{}),
	}
}

func (c *cORSPolicyEventWatcher) AddEventHandler(ctx context.Context, h CORSPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericCORSPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericCORSPolicyHandler implements a generic events.EventHandler
type genericCORSPolicyHandler struct {
	handler CORSPolicyEventHandler
}

func (h genericCORSPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CORSPolicy)
	if !ok {
		return errors.Errorf("internal error: CORSPolicy handler received event for %T", object)
	}
	return h.handler.CreateCORSPolicy(obj)
}

func (h genericCORSPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CORSPolicy)
	if !ok {
		return errors.Errorf("internal error: CORSPolicy handler received event for %T", object)
	}
	return h.handler.DeleteCORSPolicy(obj)
}

func (h genericCORSPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_policy_gloo_solo_io_v2.CORSPolicy)
	if !ok {
		return errors.Errorf("internal error: CORSPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*security_policy_gloo_solo_io_v2.CORSPolicy)
	if !ok {
		return errors.Errorf("internal error: CORSPolicy handler received event for %T", new)
	}
	return h.handler.UpdateCORSPolicy(objOld, objNew)
}

func (h genericCORSPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CORSPolicy)
	if !ok {
		return errors.Errorf("internal error: CORSPolicy handler received event for %T", object)
	}
	return h.handler.GenericCORSPolicy(obj)
}

// Handle events for the CSRFPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type CSRFPolicyEventHandler interface {
	CreateCSRFPolicy(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error
	UpdateCSRFPolicy(old, new *security_policy_gloo_solo_io_v2.CSRFPolicy) error
	DeleteCSRFPolicy(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error
	GenericCSRFPolicy(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error
}

type CSRFPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error
	OnUpdate  func(old, new *security_policy_gloo_solo_io_v2.CSRFPolicy) error
	OnDelete  func(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error
	OnGeneric func(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error
}

func (f *CSRFPolicyEventHandlerFuncs) CreateCSRFPolicy(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *CSRFPolicyEventHandlerFuncs) DeleteCSRFPolicy(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *CSRFPolicyEventHandlerFuncs) UpdateCSRFPolicy(objOld, objNew *security_policy_gloo_solo_io_v2.CSRFPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *CSRFPolicyEventHandlerFuncs) GenericCSRFPolicy(obj *security_policy_gloo_solo_io_v2.CSRFPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type CSRFPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h CSRFPolicyEventHandler, predicates ...predicate.Predicate) error
}

type cSRFPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewCSRFPolicyEventWatcher(name string, mgr manager.Manager) CSRFPolicyEventWatcher {
	return &cSRFPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_policy_gloo_solo_io_v2.CSRFPolicy{}),
	}
}

func (c *cSRFPolicyEventWatcher) AddEventHandler(ctx context.Context, h CSRFPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericCSRFPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericCSRFPolicyHandler implements a generic events.EventHandler
type genericCSRFPolicyHandler struct {
	handler CSRFPolicyEventHandler
}

func (h genericCSRFPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CSRFPolicy)
	if !ok {
		return errors.Errorf("internal error: CSRFPolicy handler received event for %T", object)
	}
	return h.handler.CreateCSRFPolicy(obj)
}

func (h genericCSRFPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CSRFPolicy)
	if !ok {
		return errors.Errorf("internal error: CSRFPolicy handler received event for %T", object)
	}
	return h.handler.DeleteCSRFPolicy(obj)
}

func (h genericCSRFPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_policy_gloo_solo_io_v2.CSRFPolicy)
	if !ok {
		return errors.Errorf("internal error: CSRFPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*security_policy_gloo_solo_io_v2.CSRFPolicy)
	if !ok {
		return errors.Errorf("internal error: CSRFPolicy handler received event for %T", new)
	}
	return h.handler.UpdateCSRFPolicy(objOld, objNew)
}

func (h genericCSRFPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.CSRFPolicy)
	if !ok {
		return errors.Errorf("internal error: CSRFPolicy handler received event for %T", object)
	}
	return h.handler.GenericCSRFPolicy(obj)
}

// Handle events for the ExtAuthPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type ExtAuthPolicyEventHandler interface {
	CreateExtAuthPolicy(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error
	UpdateExtAuthPolicy(old, new *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error
	DeleteExtAuthPolicy(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error
	GenericExtAuthPolicy(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error
}

type ExtAuthPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error
	OnUpdate  func(old, new *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error
	OnDelete  func(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error
	OnGeneric func(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error
}

func (f *ExtAuthPolicyEventHandlerFuncs) CreateExtAuthPolicy(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ExtAuthPolicyEventHandlerFuncs) DeleteExtAuthPolicy(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ExtAuthPolicyEventHandlerFuncs) UpdateExtAuthPolicy(objOld, objNew *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ExtAuthPolicyEventHandlerFuncs) GenericExtAuthPolicy(obj *security_policy_gloo_solo_io_v2.ExtAuthPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ExtAuthPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h ExtAuthPolicyEventHandler, predicates ...predicate.Predicate) error
}

type extAuthPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewExtAuthPolicyEventWatcher(name string, mgr manager.Manager) ExtAuthPolicyEventWatcher {
	return &extAuthPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_policy_gloo_solo_io_v2.ExtAuthPolicy{}),
	}
}

func (c *extAuthPolicyEventWatcher) AddEventHandler(ctx context.Context, h ExtAuthPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericExtAuthPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericExtAuthPolicyHandler implements a generic events.EventHandler
type genericExtAuthPolicyHandler struct {
	handler ExtAuthPolicyEventHandler
}

func (h genericExtAuthPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ExtAuthPolicy)
	if !ok {
		return errors.Errorf("internal error: ExtAuthPolicy handler received event for %T", object)
	}
	return h.handler.CreateExtAuthPolicy(obj)
}

func (h genericExtAuthPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ExtAuthPolicy)
	if !ok {
		return errors.Errorf("internal error: ExtAuthPolicy handler received event for %T", object)
	}
	return h.handler.DeleteExtAuthPolicy(obj)
}

func (h genericExtAuthPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_policy_gloo_solo_io_v2.ExtAuthPolicy)
	if !ok {
		return errors.Errorf("internal error: ExtAuthPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*security_policy_gloo_solo_io_v2.ExtAuthPolicy)
	if !ok {
		return errors.Errorf("internal error: ExtAuthPolicy handler received event for %T", new)
	}
	return h.handler.UpdateExtAuthPolicy(objOld, objNew)
}

func (h genericExtAuthPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ExtAuthPolicy)
	if !ok {
		return errors.Errorf("internal error: ExtAuthPolicy handler received event for %T", object)
	}
	return h.handler.GenericExtAuthPolicy(obj)
}

// Handle events for the WAFPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type WAFPolicyEventHandler interface {
	CreateWAFPolicy(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error
	UpdateWAFPolicy(old, new *security_policy_gloo_solo_io_v2.WAFPolicy) error
	DeleteWAFPolicy(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error
	GenericWAFPolicy(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error
}

type WAFPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error
	OnUpdate  func(old, new *security_policy_gloo_solo_io_v2.WAFPolicy) error
	OnDelete  func(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error
	OnGeneric func(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error
}

func (f *WAFPolicyEventHandlerFuncs) CreateWAFPolicy(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *WAFPolicyEventHandlerFuncs) DeleteWAFPolicy(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *WAFPolicyEventHandlerFuncs) UpdateWAFPolicy(objOld, objNew *security_policy_gloo_solo_io_v2.WAFPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *WAFPolicyEventHandlerFuncs) GenericWAFPolicy(obj *security_policy_gloo_solo_io_v2.WAFPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type WAFPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h WAFPolicyEventHandler, predicates ...predicate.Predicate) error
}

type wAFPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewWAFPolicyEventWatcher(name string, mgr manager.Manager) WAFPolicyEventWatcher {
	return &wAFPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_policy_gloo_solo_io_v2.WAFPolicy{}),
	}
}

func (c *wAFPolicyEventWatcher) AddEventHandler(ctx context.Context, h WAFPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericWAFPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericWAFPolicyHandler implements a generic events.EventHandler
type genericWAFPolicyHandler struct {
	handler WAFPolicyEventHandler
}

func (h genericWAFPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.WAFPolicy)
	if !ok {
		return errors.Errorf("internal error: WAFPolicy handler received event for %T", object)
	}
	return h.handler.CreateWAFPolicy(obj)
}

func (h genericWAFPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.WAFPolicy)
	if !ok {
		return errors.Errorf("internal error: WAFPolicy handler received event for %T", object)
	}
	return h.handler.DeleteWAFPolicy(obj)
}

func (h genericWAFPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_policy_gloo_solo_io_v2.WAFPolicy)
	if !ok {
		return errors.Errorf("internal error: WAFPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*security_policy_gloo_solo_io_v2.WAFPolicy)
	if !ok {
		return errors.Errorf("internal error: WAFPolicy handler received event for %T", new)
	}
	return h.handler.UpdateWAFPolicy(objOld, objNew)
}

func (h genericWAFPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.WAFPolicy)
	if !ok {
		return errors.Errorf("internal error: WAFPolicy handler received event for %T", object)
	}
	return h.handler.GenericWAFPolicy(obj)
}

// Handle events for the JWTPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type JWTPolicyEventHandler interface {
	CreateJWTPolicy(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error
	UpdateJWTPolicy(old, new *security_policy_gloo_solo_io_v2.JWTPolicy) error
	DeleteJWTPolicy(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error
	GenericJWTPolicy(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error
}

type JWTPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error
	OnUpdate  func(old, new *security_policy_gloo_solo_io_v2.JWTPolicy) error
	OnDelete  func(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error
	OnGeneric func(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error
}

func (f *JWTPolicyEventHandlerFuncs) CreateJWTPolicy(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *JWTPolicyEventHandlerFuncs) DeleteJWTPolicy(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *JWTPolicyEventHandlerFuncs) UpdateJWTPolicy(objOld, objNew *security_policy_gloo_solo_io_v2.JWTPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *JWTPolicyEventHandlerFuncs) GenericJWTPolicy(obj *security_policy_gloo_solo_io_v2.JWTPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type JWTPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h JWTPolicyEventHandler, predicates ...predicate.Predicate) error
}

type jWTPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewJWTPolicyEventWatcher(name string, mgr manager.Manager) JWTPolicyEventWatcher {
	return &jWTPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_policy_gloo_solo_io_v2.JWTPolicy{}),
	}
}

func (c *jWTPolicyEventWatcher) AddEventHandler(ctx context.Context, h JWTPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericJWTPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericJWTPolicyHandler implements a generic events.EventHandler
type genericJWTPolicyHandler struct {
	handler JWTPolicyEventHandler
}

func (h genericJWTPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.JWTPolicy)
	if !ok {
		return errors.Errorf("internal error: JWTPolicy handler received event for %T", object)
	}
	return h.handler.CreateJWTPolicy(obj)
}

func (h genericJWTPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.JWTPolicy)
	if !ok {
		return errors.Errorf("internal error: JWTPolicy handler received event for %T", object)
	}
	return h.handler.DeleteJWTPolicy(obj)
}

func (h genericJWTPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_policy_gloo_solo_io_v2.JWTPolicy)
	if !ok {
		return errors.Errorf("internal error: JWTPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*security_policy_gloo_solo_io_v2.JWTPolicy)
	if !ok {
		return errors.Errorf("internal error: JWTPolicy handler received event for %T", new)
	}
	return h.handler.UpdateJWTPolicy(objOld, objNew)
}

func (h genericJWTPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.JWTPolicy)
	if !ok {
		return errors.Errorf("internal error: JWTPolicy handler received event for %T", object)
	}
	return h.handler.GenericJWTPolicy(obj)
}

// Handle events for the ClientTLSPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type ClientTLSPolicyEventHandler interface {
	CreateClientTLSPolicy(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error
	UpdateClientTLSPolicy(old, new *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error
	DeleteClientTLSPolicy(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error
	GenericClientTLSPolicy(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error
}

type ClientTLSPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error
	OnUpdate  func(old, new *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error
	OnDelete  func(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error
	OnGeneric func(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error
}

func (f *ClientTLSPolicyEventHandlerFuncs) CreateClientTLSPolicy(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ClientTLSPolicyEventHandlerFuncs) DeleteClientTLSPolicy(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ClientTLSPolicyEventHandlerFuncs) UpdateClientTLSPolicy(objOld, objNew *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ClientTLSPolicyEventHandlerFuncs) GenericClientTLSPolicy(obj *security_policy_gloo_solo_io_v2.ClientTLSPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ClientTLSPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h ClientTLSPolicyEventHandler, predicates ...predicate.Predicate) error
}

type clientTLSPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewClientTLSPolicyEventWatcher(name string, mgr manager.Manager) ClientTLSPolicyEventWatcher {
	return &clientTLSPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_policy_gloo_solo_io_v2.ClientTLSPolicy{}),
	}
}

func (c *clientTLSPolicyEventWatcher) AddEventHandler(ctx context.Context, h ClientTLSPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericClientTLSPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericClientTLSPolicyHandler implements a generic events.EventHandler
type genericClientTLSPolicyHandler struct {
	handler ClientTLSPolicyEventHandler
}

func (h genericClientTLSPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ClientTLSPolicy)
	if !ok {
		return errors.Errorf("internal error: ClientTLSPolicy handler received event for %T", object)
	}
	return h.handler.CreateClientTLSPolicy(obj)
}

func (h genericClientTLSPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ClientTLSPolicy)
	if !ok {
		return errors.Errorf("internal error: ClientTLSPolicy handler received event for %T", object)
	}
	return h.handler.DeleteClientTLSPolicy(obj)
}

func (h genericClientTLSPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_policy_gloo_solo_io_v2.ClientTLSPolicy)
	if !ok {
		return errors.Errorf("internal error: ClientTLSPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*security_policy_gloo_solo_io_v2.ClientTLSPolicy)
	if !ok {
		return errors.Errorf("internal error: ClientTLSPolicy handler received event for %T", new)
	}
	return h.handler.UpdateClientTLSPolicy(objOld, objNew)
}

func (h genericClientTLSPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.ClientTLSPolicy)
	if !ok {
		return errors.Errorf("internal error: ClientTLSPolicy handler received event for %T", object)
	}
	return h.handler.GenericClientTLSPolicy(obj)
}

// Handle events for the GraphQLAllowedQueryPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type GraphQLAllowedQueryPolicyEventHandler interface {
	CreateGraphQLAllowedQueryPolicy(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error
	UpdateGraphQLAllowedQueryPolicy(old, new *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error
	DeleteGraphQLAllowedQueryPolicy(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error
	GenericGraphQLAllowedQueryPolicy(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error
}

type GraphQLAllowedQueryPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error
	OnUpdate  func(old, new *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error
	OnDelete  func(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error
	OnGeneric func(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error
}

func (f *GraphQLAllowedQueryPolicyEventHandlerFuncs) CreateGraphQLAllowedQueryPolicy(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *GraphQLAllowedQueryPolicyEventHandlerFuncs) DeleteGraphQLAllowedQueryPolicy(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *GraphQLAllowedQueryPolicyEventHandlerFuncs) UpdateGraphQLAllowedQueryPolicy(objOld, objNew *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *GraphQLAllowedQueryPolicyEventHandlerFuncs) GenericGraphQLAllowedQueryPolicy(obj *security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type GraphQLAllowedQueryPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h GraphQLAllowedQueryPolicyEventHandler, predicates ...predicate.Predicate) error
}

type graphQLAllowedQueryPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewGraphQLAllowedQueryPolicyEventWatcher(name string, mgr manager.Manager) GraphQLAllowedQueryPolicyEventWatcher {
	return &graphQLAllowedQueryPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy{}),
	}
}

func (c *graphQLAllowedQueryPolicyEventWatcher) AddEventHandler(ctx context.Context, h GraphQLAllowedQueryPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericGraphQLAllowedQueryPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericGraphQLAllowedQueryPolicyHandler implements a generic events.EventHandler
type genericGraphQLAllowedQueryPolicyHandler struct {
	handler GraphQLAllowedQueryPolicyEventHandler
}

func (h genericGraphQLAllowedQueryPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy)
	if !ok {
		return errors.Errorf("internal error: GraphQLAllowedQueryPolicy handler received event for %T", object)
	}
	return h.handler.CreateGraphQLAllowedQueryPolicy(obj)
}

func (h genericGraphQLAllowedQueryPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy)
	if !ok {
		return errors.Errorf("internal error: GraphQLAllowedQueryPolicy handler received event for %T", object)
	}
	return h.handler.DeleteGraphQLAllowedQueryPolicy(obj)
}

func (h genericGraphQLAllowedQueryPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy)
	if !ok {
		return errors.Errorf("internal error: GraphQLAllowedQueryPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy)
	if !ok {
		return errors.Errorf("internal error: GraphQLAllowedQueryPolicy handler received event for %T", new)
	}
	return h.handler.UpdateGraphQLAllowedQueryPolicy(objOld, objNew)
}

func (h genericGraphQLAllowedQueryPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_policy_gloo_solo_io_v2.GraphQLAllowedQueryPolicy)
	if !ok {
		return errors.Errorf("internal error: GraphQLAllowedQueryPolicy handler received event for %T", object)
	}
	return h.handler.GenericGraphQLAllowedQueryPolicy(obj)
}
