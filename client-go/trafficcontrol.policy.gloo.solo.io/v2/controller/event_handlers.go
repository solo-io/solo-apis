// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	trafficcontrol_policy_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/trafficcontrol.policy.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the MirrorPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type MirrorPolicyEventHandler interface {
	CreateMirrorPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error
	UpdateMirrorPolicy(old, new *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error
	DeleteMirrorPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error
	GenericMirrorPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error
}

type MirrorPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error
	OnUpdate  func(old, new *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error
	OnDelete  func(obj *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error
	OnGeneric func(obj *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error
}

func (f *MirrorPolicyEventHandlerFuncs) CreateMirrorPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *MirrorPolicyEventHandlerFuncs) DeleteMirrorPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *MirrorPolicyEventHandlerFuncs) UpdateMirrorPolicy(objOld, objNew *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *MirrorPolicyEventHandlerFuncs) GenericMirrorPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type MirrorPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h MirrorPolicyEventHandler, predicates ...predicate.Predicate) error
}

type mirrorPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewMirrorPolicyEventWatcher(name string, mgr manager.Manager) MirrorPolicyEventWatcher {
	return &mirrorPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy{}),
	}
}

func (c *mirrorPolicyEventWatcher) AddEventHandler(ctx context.Context, h MirrorPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericMirrorPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericMirrorPolicyHandler implements a generic events.EventHandler
type genericMirrorPolicyHandler struct {
	handler MirrorPolicyEventHandler
}

func (h genericMirrorPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy)
	if !ok {
		return errors.Errorf("internal error: MirrorPolicy handler received event for %T", object)
	}
	return h.handler.CreateMirrorPolicy(obj)
}

func (h genericMirrorPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy)
	if !ok {
		return errors.Errorf("internal error: MirrorPolicy handler received event for %T", object)
	}
	return h.handler.DeleteMirrorPolicy(obj)
}

func (h genericMirrorPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy)
	if !ok {
		return errors.Errorf("internal error: MirrorPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy)
	if !ok {
		return errors.Errorf("internal error: MirrorPolicy handler received event for %T", new)
	}
	return h.handler.UpdateMirrorPolicy(objOld, objNew)
}

func (h genericMirrorPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.MirrorPolicy)
	if !ok {
		return errors.Errorf("internal error: MirrorPolicy handler received event for %T", object)
	}
	return h.handler.GenericMirrorPolicy(obj)
}

// Handle events for the RateLimitPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type RateLimitPolicyEventHandler interface {
	CreateRateLimitPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error
	UpdateRateLimitPolicy(old, new *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error
	DeleteRateLimitPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error
	GenericRateLimitPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error
}

type RateLimitPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error
	OnUpdate  func(old, new *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error
	OnDelete  func(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error
	OnGeneric func(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error
}

func (f *RateLimitPolicyEventHandlerFuncs) CreateRateLimitPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RateLimitPolicyEventHandlerFuncs) DeleteRateLimitPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RateLimitPolicyEventHandlerFuncs) UpdateRateLimitPolicy(objOld, objNew *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RateLimitPolicyEventHandlerFuncs) GenericRateLimitPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RateLimitPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h RateLimitPolicyEventHandler, predicates ...predicate.Predicate) error
}

type rateLimitPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewRateLimitPolicyEventWatcher(name string, mgr manager.Manager) RateLimitPolicyEventWatcher {
	return &rateLimitPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy{}),
	}
}

func (c *rateLimitPolicyEventWatcher) AddEventHandler(ctx context.Context, h RateLimitPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRateLimitPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRateLimitPolicyHandler implements a generic events.EventHandler
type genericRateLimitPolicyHandler struct {
	handler RateLimitPolicyEventHandler
}

func (h genericRateLimitPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy)
	if !ok {
		return errors.Errorf("internal error: RateLimitPolicy handler received event for %T", object)
	}
	return h.handler.CreateRateLimitPolicy(obj)
}

func (h genericRateLimitPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy)
	if !ok {
		return errors.Errorf("internal error: RateLimitPolicy handler received event for %T", object)
	}
	return h.handler.DeleteRateLimitPolicy(obj)
}

func (h genericRateLimitPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy)
	if !ok {
		return errors.Errorf("internal error: RateLimitPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy)
	if !ok {
		return errors.Errorf("internal error: RateLimitPolicy handler received event for %T", new)
	}
	return h.handler.UpdateRateLimitPolicy(objOld, objNew)
}

func (h genericRateLimitPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitPolicy)
	if !ok {
		return errors.Errorf("internal error: RateLimitPolicy handler received event for %T", object)
	}
	return h.handler.GenericRateLimitPolicy(obj)
}

// Handle events for the RateLimitClientConfig Resource
// DEPRECATED: Prefer reconciler pattern.
type RateLimitClientConfigEventHandler interface {
	CreateRateLimitClientConfig(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error
	UpdateRateLimitClientConfig(old, new *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error
	DeleteRateLimitClientConfig(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error
	GenericRateLimitClientConfig(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error
}

type RateLimitClientConfigEventHandlerFuncs struct {
	OnCreate  func(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error
	OnUpdate  func(old, new *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error
	OnDelete  func(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error
	OnGeneric func(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error
}

func (f *RateLimitClientConfigEventHandlerFuncs) CreateRateLimitClientConfig(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RateLimitClientConfigEventHandlerFuncs) DeleteRateLimitClientConfig(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RateLimitClientConfigEventHandlerFuncs) UpdateRateLimitClientConfig(objOld, objNew *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RateLimitClientConfigEventHandlerFuncs) GenericRateLimitClientConfig(obj *trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RateLimitClientConfigEventWatcher interface {
	AddEventHandler(ctx context.Context, h RateLimitClientConfigEventHandler, predicates ...predicate.Predicate) error
}

type rateLimitClientConfigEventWatcher struct {
	watcher events.EventWatcher
}

func NewRateLimitClientConfigEventWatcher(name string, mgr manager.Manager) RateLimitClientConfigEventWatcher {
	return &rateLimitClientConfigEventWatcher{
		watcher: events.NewWatcher(name, mgr, &trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig{}),
	}
}

func (c *rateLimitClientConfigEventWatcher) AddEventHandler(ctx context.Context, h RateLimitClientConfigEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRateLimitClientConfigHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRateLimitClientConfigHandler implements a generic events.EventHandler
type genericRateLimitClientConfigHandler struct {
	handler RateLimitClientConfigEventHandler
}

func (h genericRateLimitClientConfigHandler) Create(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitClientConfig handler received event for %T", object)
	}
	return h.handler.CreateRateLimitClientConfig(obj)
}

func (h genericRateLimitClientConfigHandler) Delete(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitClientConfig handler received event for %T", object)
	}
	return h.handler.DeleteRateLimitClientConfig(obj)
}

func (h genericRateLimitClientConfigHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitClientConfig handler received event for %T", old)
	}
	objNew, ok := new.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitClientConfig handler received event for %T", new)
	}
	return h.handler.UpdateRateLimitClientConfig(objOld, objNew)
}

func (h genericRateLimitClientConfigHandler) Generic(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.RateLimitClientConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitClientConfig handler received event for %T", object)
	}
	return h.handler.GenericRateLimitClientConfig(obj)
}

// Handle events for the HeaderManipulationPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type HeaderManipulationPolicyEventHandler interface {
	CreateHeaderManipulationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error
	UpdateHeaderManipulationPolicy(old, new *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error
	DeleteHeaderManipulationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error
	GenericHeaderManipulationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error
}

type HeaderManipulationPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error
	OnUpdate  func(old, new *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error
	OnDelete  func(obj *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error
	OnGeneric func(obj *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error
}

func (f *HeaderManipulationPolicyEventHandlerFuncs) CreateHeaderManipulationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *HeaderManipulationPolicyEventHandlerFuncs) DeleteHeaderManipulationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *HeaderManipulationPolicyEventHandlerFuncs) UpdateHeaderManipulationPolicy(objOld, objNew *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *HeaderManipulationPolicyEventHandlerFuncs) GenericHeaderManipulationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type HeaderManipulationPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h HeaderManipulationPolicyEventHandler, predicates ...predicate.Predicate) error
}

type headerManipulationPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewHeaderManipulationPolicyEventWatcher(name string, mgr manager.Manager) HeaderManipulationPolicyEventWatcher {
	return &headerManipulationPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy{}),
	}
}

func (c *headerManipulationPolicyEventWatcher) AddEventHandler(ctx context.Context, h HeaderManipulationPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericHeaderManipulationPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericHeaderManipulationPolicyHandler implements a generic events.EventHandler
type genericHeaderManipulationPolicyHandler struct {
	handler HeaderManipulationPolicyEventHandler
}

func (h genericHeaderManipulationPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy)
	if !ok {
		return errors.Errorf("internal error: HeaderManipulationPolicy handler received event for %T", object)
	}
	return h.handler.CreateHeaderManipulationPolicy(obj)
}

func (h genericHeaderManipulationPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy)
	if !ok {
		return errors.Errorf("internal error: HeaderManipulationPolicy handler received event for %T", object)
	}
	return h.handler.DeleteHeaderManipulationPolicy(obj)
}

func (h genericHeaderManipulationPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy)
	if !ok {
		return errors.Errorf("internal error: HeaderManipulationPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy)
	if !ok {
		return errors.Errorf("internal error: HeaderManipulationPolicy handler received event for %T", new)
	}
	return h.handler.UpdateHeaderManipulationPolicy(objOld, objNew)
}

func (h genericHeaderManipulationPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.HeaderManipulationPolicy)
	if !ok {
		return errors.Errorf("internal error: HeaderManipulationPolicy handler received event for %T", object)
	}
	return h.handler.GenericHeaderManipulationPolicy(obj)
}

// Handle events for the TransformationPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type TransformationPolicyEventHandler interface {
	CreateTransformationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error
	UpdateTransformationPolicy(old, new *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error
	DeleteTransformationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error
	GenericTransformationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error
}

type TransformationPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error
	OnUpdate  func(old, new *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error
	OnDelete  func(obj *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error
	OnGeneric func(obj *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error
}

func (f *TransformationPolicyEventHandlerFuncs) CreateTransformationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *TransformationPolicyEventHandlerFuncs) DeleteTransformationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *TransformationPolicyEventHandlerFuncs) UpdateTransformationPolicy(objOld, objNew *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *TransformationPolicyEventHandlerFuncs) GenericTransformationPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type TransformationPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h TransformationPolicyEventHandler, predicates ...predicate.Predicate) error
}

type transformationPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewTransformationPolicyEventWatcher(name string, mgr manager.Manager) TransformationPolicyEventWatcher {
	return &transformationPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy{}),
	}
}

func (c *transformationPolicyEventWatcher) AddEventHandler(ctx context.Context, h TransformationPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericTransformationPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericTransformationPolicyHandler implements a generic events.EventHandler
type genericTransformationPolicyHandler struct {
	handler TransformationPolicyEventHandler
}

func (h genericTransformationPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy)
	if !ok {
		return errors.Errorf("internal error: TransformationPolicy handler received event for %T", object)
	}
	return h.handler.CreateTransformationPolicy(obj)
}

func (h genericTransformationPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy)
	if !ok {
		return errors.Errorf("internal error: TransformationPolicy handler received event for %T", object)
	}
	return h.handler.DeleteTransformationPolicy(obj)
}

func (h genericTransformationPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy)
	if !ok {
		return errors.Errorf("internal error: TransformationPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy)
	if !ok {
		return errors.Errorf("internal error: TransformationPolicy handler received event for %T", new)
	}
	return h.handler.UpdateTransformationPolicy(objOld, objNew)
}

func (h genericTransformationPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.TransformationPolicy)
	if !ok {
		return errors.Errorf("internal error: TransformationPolicy handler received event for %T", object)
	}
	return h.handler.GenericTransformationPolicy(obj)
}

// Handle events for the LoadBalancerPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type LoadBalancerPolicyEventHandler interface {
	CreateLoadBalancerPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error
	UpdateLoadBalancerPolicy(old, new *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error
	DeleteLoadBalancerPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error
	GenericLoadBalancerPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error
}

type LoadBalancerPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error
	OnUpdate  func(old, new *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error
	OnDelete  func(obj *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error
	OnGeneric func(obj *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error
}

func (f *LoadBalancerPolicyEventHandlerFuncs) CreateLoadBalancerPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *LoadBalancerPolicyEventHandlerFuncs) DeleteLoadBalancerPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *LoadBalancerPolicyEventHandlerFuncs) UpdateLoadBalancerPolicy(objOld, objNew *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *LoadBalancerPolicyEventHandlerFuncs) GenericLoadBalancerPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type LoadBalancerPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h LoadBalancerPolicyEventHandler, predicates ...predicate.Predicate) error
}

type loadBalancerPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewLoadBalancerPolicyEventWatcher(name string, mgr manager.Manager) LoadBalancerPolicyEventWatcher {
	return &loadBalancerPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy{}),
	}
}

func (c *loadBalancerPolicyEventWatcher) AddEventHandler(ctx context.Context, h LoadBalancerPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericLoadBalancerPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericLoadBalancerPolicyHandler implements a generic events.EventHandler
type genericLoadBalancerPolicyHandler struct {
	handler LoadBalancerPolicyEventHandler
}

func (h genericLoadBalancerPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy)
	if !ok {
		return errors.Errorf("internal error: LoadBalancerPolicy handler received event for %T", object)
	}
	return h.handler.CreateLoadBalancerPolicy(obj)
}

func (h genericLoadBalancerPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy)
	if !ok {
		return errors.Errorf("internal error: LoadBalancerPolicy handler received event for %T", object)
	}
	return h.handler.DeleteLoadBalancerPolicy(obj)
}

func (h genericLoadBalancerPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy)
	if !ok {
		return errors.Errorf("internal error: LoadBalancerPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy)
	if !ok {
		return errors.Errorf("internal error: LoadBalancerPolicy handler received event for %T", new)
	}
	return h.handler.UpdateLoadBalancerPolicy(objOld, objNew)
}

func (h genericLoadBalancerPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.LoadBalancerPolicy)
	if !ok {
		return errors.Errorf("internal error: LoadBalancerPolicy handler received event for %T", object)
	}
	return h.handler.GenericLoadBalancerPolicy(obj)
}

// Handle events for the ProxyProtocolPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type ProxyProtocolPolicyEventHandler interface {
	CreateProxyProtocolPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error
	UpdateProxyProtocolPolicy(old, new *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error
	DeleteProxyProtocolPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error
	GenericProxyProtocolPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error
}

type ProxyProtocolPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error
	OnUpdate  func(old, new *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error
	OnDelete  func(obj *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error
	OnGeneric func(obj *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error
}

func (f *ProxyProtocolPolicyEventHandlerFuncs) CreateProxyProtocolPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ProxyProtocolPolicyEventHandlerFuncs) DeleteProxyProtocolPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ProxyProtocolPolicyEventHandlerFuncs) UpdateProxyProtocolPolicy(objOld, objNew *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ProxyProtocolPolicyEventHandlerFuncs) GenericProxyProtocolPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ProxyProtocolPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h ProxyProtocolPolicyEventHandler, predicates ...predicate.Predicate) error
}

type proxyProtocolPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewProxyProtocolPolicyEventWatcher(name string, mgr manager.Manager) ProxyProtocolPolicyEventWatcher {
	return &proxyProtocolPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy{}),
	}
}

func (c *proxyProtocolPolicyEventWatcher) AddEventHandler(ctx context.Context, h ProxyProtocolPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericProxyProtocolPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericProxyProtocolPolicyHandler implements a generic events.EventHandler
type genericProxyProtocolPolicyHandler struct {
	handler ProxyProtocolPolicyEventHandler
}

func (h genericProxyProtocolPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy)
	if !ok {
		return errors.Errorf("internal error: ProxyProtocolPolicy handler received event for %T", object)
	}
	return h.handler.CreateProxyProtocolPolicy(obj)
}

func (h genericProxyProtocolPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy)
	if !ok {
		return errors.Errorf("internal error: ProxyProtocolPolicy handler received event for %T", object)
	}
	return h.handler.DeleteProxyProtocolPolicy(obj)
}

func (h genericProxyProtocolPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy)
	if !ok {
		return errors.Errorf("internal error: ProxyProtocolPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy)
	if !ok {
		return errors.Errorf("internal error: ProxyProtocolPolicy handler received event for %T", new)
	}
	return h.handler.UpdateProxyProtocolPolicy(objOld, objNew)
}

func (h genericProxyProtocolPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.ProxyProtocolPolicy)
	if !ok {
		return errors.Errorf("internal error: ProxyProtocolPolicy handler received event for %T", object)
	}
	return h.handler.GenericProxyProtocolPolicy(obj)
}

// Handle events for the HTTPBufferPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type HTTPBufferPolicyEventHandler interface {
	CreateHTTPBufferPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error
	UpdateHTTPBufferPolicy(old, new *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error
	DeleteHTTPBufferPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error
	GenericHTTPBufferPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error
}

type HTTPBufferPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error
	OnUpdate  func(old, new *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error
	OnDelete  func(obj *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error
	OnGeneric func(obj *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error
}

func (f *HTTPBufferPolicyEventHandlerFuncs) CreateHTTPBufferPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *HTTPBufferPolicyEventHandlerFuncs) DeleteHTTPBufferPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *HTTPBufferPolicyEventHandlerFuncs) UpdateHTTPBufferPolicy(objOld, objNew *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *HTTPBufferPolicyEventHandlerFuncs) GenericHTTPBufferPolicy(obj *trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type HTTPBufferPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h HTTPBufferPolicyEventHandler, predicates ...predicate.Predicate) error
}

type hTTPBufferPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewHTTPBufferPolicyEventWatcher(name string, mgr manager.Manager) HTTPBufferPolicyEventWatcher {
	return &hTTPBufferPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy{}),
	}
}

func (c *hTTPBufferPolicyEventWatcher) AddEventHandler(ctx context.Context, h HTTPBufferPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericHTTPBufferPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericHTTPBufferPolicyHandler implements a generic events.EventHandler
type genericHTTPBufferPolicyHandler struct {
	handler HTTPBufferPolicyEventHandler
}

func (h genericHTTPBufferPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy)
	if !ok {
		return errors.Errorf("internal error: HTTPBufferPolicy handler received event for %T", object)
	}
	return h.handler.CreateHTTPBufferPolicy(obj)
}

func (h genericHTTPBufferPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy)
	if !ok {
		return errors.Errorf("internal error: HTTPBufferPolicy handler received event for %T", object)
	}
	return h.handler.DeleteHTTPBufferPolicy(obj)
}

func (h genericHTTPBufferPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy)
	if !ok {
		return errors.Errorf("internal error: HTTPBufferPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy)
	if !ok {
		return errors.Errorf("internal error: HTTPBufferPolicy handler received event for %T", new)
	}
	return h.handler.UpdateHTTPBufferPolicy(objOld, objNew)
}

func (h genericHTTPBufferPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*trafficcontrol_policy_gloo_solo_io_v2.HTTPBufferPolicy)
	if !ok {
		return errors.Errorf("internal error: HTTPBufferPolicy handler received event for %T", object)
	}
	return h.handler.GenericHTTPBufferPolicy(obj)
}