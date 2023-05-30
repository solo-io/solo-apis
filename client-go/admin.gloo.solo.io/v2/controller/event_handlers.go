// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	admin_gloo_solo_io_v2 "github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the Workspace Resource
// DEPRECATED: Prefer reconciler pattern.
type WorkspaceEventHandler interface {
	CreateWorkspace(obj *admin_gloo_solo_io_v2.Workspace) error
	UpdateWorkspace(old, new *admin_gloo_solo_io_v2.Workspace) error
	DeleteWorkspace(obj *admin_gloo_solo_io_v2.Workspace) error
	GenericWorkspace(obj *admin_gloo_solo_io_v2.Workspace) error
}

type WorkspaceEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.Workspace) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.Workspace) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.Workspace) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.Workspace) error
}

func (f *WorkspaceEventHandlerFuncs) CreateWorkspace(obj *admin_gloo_solo_io_v2.Workspace) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *WorkspaceEventHandlerFuncs) DeleteWorkspace(obj *admin_gloo_solo_io_v2.Workspace) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *WorkspaceEventHandlerFuncs) UpdateWorkspace(objOld, objNew *admin_gloo_solo_io_v2.Workspace) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *WorkspaceEventHandlerFuncs) GenericWorkspace(obj *admin_gloo_solo_io_v2.Workspace) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type WorkspaceEventWatcher interface {
	AddEventHandler(ctx context.Context, h WorkspaceEventHandler, predicates ...predicate.Predicate) error
}

type workspaceEventWatcher struct {
	watcher events.EventWatcher
}

func NewWorkspaceEventWatcher(name string, mgr manager.Manager) WorkspaceEventWatcher {
	return &workspaceEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.Workspace{}),
	}
}

func (c *workspaceEventWatcher) AddEventHandler(ctx context.Context, h WorkspaceEventHandler, predicates ...predicate.Predicate) error {
	handler := genericWorkspaceHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericWorkspaceHandler implements a generic events.EventHandler
type genericWorkspaceHandler struct {
	handler WorkspaceEventHandler
}

func (h genericWorkspaceHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.Workspace)
	if !ok {
		return errors.Errorf("internal error: Workspace handler received event for %T", object)
	}
	return h.handler.CreateWorkspace(obj)
}

func (h genericWorkspaceHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.Workspace)
	if !ok {
		return errors.Errorf("internal error: Workspace handler received event for %T", object)
	}
	return h.handler.DeleteWorkspace(obj)
}

func (h genericWorkspaceHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.Workspace)
	if !ok {
		return errors.Errorf("internal error: Workspace handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.Workspace)
	if !ok {
		return errors.Errorf("internal error: Workspace handler received event for %T", new)
	}
	return h.handler.UpdateWorkspace(objOld, objNew)
}

func (h genericWorkspaceHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.Workspace)
	if !ok {
		return errors.Errorf("internal error: Workspace handler received event for %T", object)
	}
	return h.handler.GenericWorkspace(obj)
}

// Handle events for the WorkspaceSettings Resource
// DEPRECATED: Prefer reconciler pattern.
type WorkspaceSettingsEventHandler interface {
	CreateWorkspaceSettings(obj *admin_gloo_solo_io_v2.WorkspaceSettings) error
	UpdateWorkspaceSettings(old, new *admin_gloo_solo_io_v2.WorkspaceSettings) error
	DeleteWorkspaceSettings(obj *admin_gloo_solo_io_v2.WorkspaceSettings) error
	GenericWorkspaceSettings(obj *admin_gloo_solo_io_v2.WorkspaceSettings) error
}

type WorkspaceSettingsEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.WorkspaceSettings) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.WorkspaceSettings) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.WorkspaceSettings) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.WorkspaceSettings) error
}

func (f *WorkspaceSettingsEventHandlerFuncs) CreateWorkspaceSettings(obj *admin_gloo_solo_io_v2.WorkspaceSettings) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *WorkspaceSettingsEventHandlerFuncs) DeleteWorkspaceSettings(obj *admin_gloo_solo_io_v2.WorkspaceSettings) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *WorkspaceSettingsEventHandlerFuncs) UpdateWorkspaceSettings(objOld, objNew *admin_gloo_solo_io_v2.WorkspaceSettings) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *WorkspaceSettingsEventHandlerFuncs) GenericWorkspaceSettings(obj *admin_gloo_solo_io_v2.WorkspaceSettings) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type WorkspaceSettingsEventWatcher interface {
	AddEventHandler(ctx context.Context, h WorkspaceSettingsEventHandler, predicates ...predicate.Predicate) error
}

type workspaceSettingsEventWatcher struct {
	watcher events.EventWatcher
}

func NewWorkspaceSettingsEventWatcher(name string, mgr manager.Manager) WorkspaceSettingsEventWatcher {
	return &workspaceSettingsEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.WorkspaceSettings{}),
	}
}

func (c *workspaceSettingsEventWatcher) AddEventHandler(ctx context.Context, h WorkspaceSettingsEventHandler, predicates ...predicate.Predicate) error {
	handler := genericWorkspaceSettingsHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericWorkspaceSettingsHandler implements a generic events.EventHandler
type genericWorkspaceSettingsHandler struct {
	handler WorkspaceSettingsEventHandler
}

func (h genericWorkspaceSettingsHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.WorkspaceSettings)
	if !ok {
		return errors.Errorf("internal error: WorkspaceSettings handler received event for %T", object)
	}
	return h.handler.CreateWorkspaceSettings(obj)
}

func (h genericWorkspaceSettingsHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.WorkspaceSettings)
	if !ok {
		return errors.Errorf("internal error: WorkspaceSettings handler received event for %T", object)
	}
	return h.handler.DeleteWorkspaceSettings(obj)
}

func (h genericWorkspaceSettingsHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.WorkspaceSettings)
	if !ok {
		return errors.Errorf("internal error: WorkspaceSettings handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.WorkspaceSettings)
	if !ok {
		return errors.Errorf("internal error: WorkspaceSettings handler received event for %T", new)
	}
	return h.handler.UpdateWorkspaceSettings(objOld, objNew)
}

func (h genericWorkspaceSettingsHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.WorkspaceSettings)
	if !ok {
		return errors.Errorf("internal error: WorkspaceSettings handler received event for %T", object)
	}
	return h.handler.GenericWorkspaceSettings(obj)
}

// Handle events for the KubernetesCluster Resource
// DEPRECATED: Prefer reconciler pattern.
type KubernetesClusterEventHandler interface {
	CreateKubernetesCluster(obj *admin_gloo_solo_io_v2.KubernetesCluster) error
	UpdateKubernetesCluster(old, new *admin_gloo_solo_io_v2.KubernetesCluster) error
	DeleteKubernetesCluster(obj *admin_gloo_solo_io_v2.KubernetesCluster) error
	GenericKubernetesCluster(obj *admin_gloo_solo_io_v2.KubernetesCluster) error
}

type KubernetesClusterEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.KubernetesCluster) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.KubernetesCluster) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.KubernetesCluster) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.KubernetesCluster) error
}

func (f *KubernetesClusterEventHandlerFuncs) CreateKubernetesCluster(obj *admin_gloo_solo_io_v2.KubernetesCluster) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *KubernetesClusterEventHandlerFuncs) DeleteKubernetesCluster(obj *admin_gloo_solo_io_v2.KubernetesCluster) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *KubernetesClusterEventHandlerFuncs) UpdateKubernetesCluster(objOld, objNew *admin_gloo_solo_io_v2.KubernetesCluster) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *KubernetesClusterEventHandlerFuncs) GenericKubernetesCluster(obj *admin_gloo_solo_io_v2.KubernetesCluster) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type KubernetesClusterEventWatcher interface {
	AddEventHandler(ctx context.Context, h KubernetesClusterEventHandler, predicates ...predicate.Predicate) error
}

type kubernetesClusterEventWatcher struct {
	watcher events.EventWatcher
}

func NewKubernetesClusterEventWatcher(name string, mgr manager.Manager) KubernetesClusterEventWatcher {
	return &kubernetesClusterEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.KubernetesCluster{}),
	}
}

func (c *kubernetesClusterEventWatcher) AddEventHandler(ctx context.Context, h KubernetesClusterEventHandler, predicates ...predicate.Predicate) error {
	handler := genericKubernetesClusterHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericKubernetesClusterHandler implements a generic events.EventHandler
type genericKubernetesClusterHandler struct {
	handler KubernetesClusterEventHandler
}

func (h genericKubernetesClusterHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.KubernetesCluster)
	if !ok {
		return errors.Errorf("internal error: KubernetesCluster handler received event for %T", object)
	}
	return h.handler.CreateKubernetesCluster(obj)
}

func (h genericKubernetesClusterHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.KubernetesCluster)
	if !ok {
		return errors.Errorf("internal error: KubernetesCluster handler received event for %T", object)
	}
	return h.handler.DeleteKubernetesCluster(obj)
}

func (h genericKubernetesClusterHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.KubernetesCluster)
	if !ok {
		return errors.Errorf("internal error: KubernetesCluster handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.KubernetesCluster)
	if !ok {
		return errors.Errorf("internal error: KubernetesCluster handler received event for %T", new)
	}
	return h.handler.UpdateKubernetesCluster(objOld, objNew)
}

func (h genericKubernetesClusterHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.KubernetesCluster)
	if !ok {
		return errors.Errorf("internal error: KubernetesCluster handler received event for %T", object)
	}
	return h.handler.GenericKubernetesCluster(obj)
}

// Handle events for the RootTrustPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type RootTrustPolicyEventHandler interface {
	CreateRootTrustPolicy(obj *admin_gloo_solo_io_v2.RootTrustPolicy) error
	UpdateRootTrustPolicy(old, new *admin_gloo_solo_io_v2.RootTrustPolicy) error
	DeleteRootTrustPolicy(obj *admin_gloo_solo_io_v2.RootTrustPolicy) error
	GenericRootTrustPolicy(obj *admin_gloo_solo_io_v2.RootTrustPolicy) error
}

type RootTrustPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.RootTrustPolicy) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.RootTrustPolicy) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.RootTrustPolicy) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.RootTrustPolicy) error
}

func (f *RootTrustPolicyEventHandlerFuncs) CreateRootTrustPolicy(obj *admin_gloo_solo_io_v2.RootTrustPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RootTrustPolicyEventHandlerFuncs) DeleteRootTrustPolicy(obj *admin_gloo_solo_io_v2.RootTrustPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RootTrustPolicyEventHandlerFuncs) UpdateRootTrustPolicy(objOld, objNew *admin_gloo_solo_io_v2.RootTrustPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RootTrustPolicyEventHandlerFuncs) GenericRootTrustPolicy(obj *admin_gloo_solo_io_v2.RootTrustPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RootTrustPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h RootTrustPolicyEventHandler, predicates ...predicate.Predicate) error
}

type rootTrustPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewRootTrustPolicyEventWatcher(name string, mgr manager.Manager) RootTrustPolicyEventWatcher {
	return &rootTrustPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.RootTrustPolicy{}),
	}
}

func (c *rootTrustPolicyEventWatcher) AddEventHandler(ctx context.Context, h RootTrustPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRootTrustPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRootTrustPolicyHandler implements a generic events.EventHandler
type genericRootTrustPolicyHandler struct {
	handler RootTrustPolicyEventHandler
}

func (h genericRootTrustPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.RootTrustPolicy)
	if !ok {
		return errors.Errorf("internal error: RootTrustPolicy handler received event for %T", object)
	}
	return h.handler.CreateRootTrustPolicy(obj)
}

func (h genericRootTrustPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.RootTrustPolicy)
	if !ok {
		return errors.Errorf("internal error: RootTrustPolicy handler received event for %T", object)
	}
	return h.handler.DeleteRootTrustPolicy(obj)
}

func (h genericRootTrustPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.RootTrustPolicy)
	if !ok {
		return errors.Errorf("internal error: RootTrustPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.RootTrustPolicy)
	if !ok {
		return errors.Errorf("internal error: RootTrustPolicy handler received event for %T", new)
	}
	return h.handler.UpdateRootTrustPolicy(objOld, objNew)
}

func (h genericRootTrustPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.RootTrustPolicy)
	if !ok {
		return errors.Errorf("internal error: RootTrustPolicy handler received event for %T", object)
	}
	return h.handler.GenericRootTrustPolicy(obj)
}

// Handle events for the ExtAuthServer Resource
// DEPRECATED: Prefer reconciler pattern.
type ExtAuthServerEventHandler interface {
	CreateExtAuthServer(obj *admin_gloo_solo_io_v2.ExtAuthServer) error
	UpdateExtAuthServer(old, new *admin_gloo_solo_io_v2.ExtAuthServer) error
	DeleteExtAuthServer(obj *admin_gloo_solo_io_v2.ExtAuthServer) error
	GenericExtAuthServer(obj *admin_gloo_solo_io_v2.ExtAuthServer) error
}

type ExtAuthServerEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.ExtAuthServer) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.ExtAuthServer) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.ExtAuthServer) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.ExtAuthServer) error
}

func (f *ExtAuthServerEventHandlerFuncs) CreateExtAuthServer(obj *admin_gloo_solo_io_v2.ExtAuthServer) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ExtAuthServerEventHandlerFuncs) DeleteExtAuthServer(obj *admin_gloo_solo_io_v2.ExtAuthServer) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ExtAuthServerEventHandlerFuncs) UpdateExtAuthServer(objOld, objNew *admin_gloo_solo_io_v2.ExtAuthServer) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ExtAuthServerEventHandlerFuncs) GenericExtAuthServer(obj *admin_gloo_solo_io_v2.ExtAuthServer) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ExtAuthServerEventWatcher interface {
	AddEventHandler(ctx context.Context, h ExtAuthServerEventHandler, predicates ...predicate.Predicate) error
}

type extAuthServerEventWatcher struct {
	watcher events.EventWatcher
}

func NewExtAuthServerEventWatcher(name string, mgr manager.Manager) ExtAuthServerEventWatcher {
	return &extAuthServerEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.ExtAuthServer{}),
	}
}

func (c *extAuthServerEventWatcher) AddEventHandler(ctx context.Context, h ExtAuthServerEventHandler, predicates ...predicate.Predicate) error {
	handler := genericExtAuthServerHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericExtAuthServerHandler implements a generic events.EventHandler
type genericExtAuthServerHandler struct {
	handler ExtAuthServerEventHandler
}

func (h genericExtAuthServerHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.ExtAuthServer)
	if !ok {
		return errors.Errorf("internal error: ExtAuthServer handler received event for %T", object)
	}
	return h.handler.CreateExtAuthServer(obj)
}

func (h genericExtAuthServerHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.ExtAuthServer)
	if !ok {
		return errors.Errorf("internal error: ExtAuthServer handler received event for %T", object)
	}
	return h.handler.DeleteExtAuthServer(obj)
}

func (h genericExtAuthServerHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.ExtAuthServer)
	if !ok {
		return errors.Errorf("internal error: ExtAuthServer handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.ExtAuthServer)
	if !ok {
		return errors.Errorf("internal error: ExtAuthServer handler received event for %T", new)
	}
	return h.handler.UpdateExtAuthServer(objOld, objNew)
}

func (h genericExtAuthServerHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.ExtAuthServer)
	if !ok {
		return errors.Errorf("internal error: ExtAuthServer handler received event for %T", object)
	}
	return h.handler.GenericExtAuthServer(obj)
}

// Handle events for the RateLimitServerSettings Resource
// DEPRECATED: Prefer reconciler pattern.
type RateLimitServerSettingsEventHandler interface {
	CreateRateLimitServerSettings(obj *admin_gloo_solo_io_v2.RateLimitServerSettings) error
	UpdateRateLimitServerSettings(old, new *admin_gloo_solo_io_v2.RateLimitServerSettings) error
	DeleteRateLimitServerSettings(obj *admin_gloo_solo_io_v2.RateLimitServerSettings) error
	GenericRateLimitServerSettings(obj *admin_gloo_solo_io_v2.RateLimitServerSettings) error
}

type RateLimitServerSettingsEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.RateLimitServerSettings) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.RateLimitServerSettings) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.RateLimitServerSettings) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.RateLimitServerSettings) error
}

func (f *RateLimitServerSettingsEventHandlerFuncs) CreateRateLimitServerSettings(obj *admin_gloo_solo_io_v2.RateLimitServerSettings) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RateLimitServerSettingsEventHandlerFuncs) DeleteRateLimitServerSettings(obj *admin_gloo_solo_io_v2.RateLimitServerSettings) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RateLimitServerSettingsEventHandlerFuncs) UpdateRateLimitServerSettings(objOld, objNew *admin_gloo_solo_io_v2.RateLimitServerSettings) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RateLimitServerSettingsEventHandlerFuncs) GenericRateLimitServerSettings(obj *admin_gloo_solo_io_v2.RateLimitServerSettings) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RateLimitServerSettingsEventWatcher interface {
	AddEventHandler(ctx context.Context, h RateLimitServerSettingsEventHandler, predicates ...predicate.Predicate) error
}

type rateLimitServerSettingsEventWatcher struct {
	watcher events.EventWatcher
}

func NewRateLimitServerSettingsEventWatcher(name string, mgr manager.Manager) RateLimitServerSettingsEventWatcher {
	return &rateLimitServerSettingsEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.RateLimitServerSettings{}),
	}
}

func (c *rateLimitServerSettingsEventWatcher) AddEventHandler(ctx context.Context, h RateLimitServerSettingsEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRateLimitServerSettingsHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRateLimitServerSettingsHandler implements a generic events.EventHandler
type genericRateLimitServerSettingsHandler struct {
	handler RateLimitServerSettingsEventHandler
}

func (h genericRateLimitServerSettingsHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.RateLimitServerSettings)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerSettings handler received event for %T", object)
	}
	return h.handler.CreateRateLimitServerSettings(obj)
}

func (h genericRateLimitServerSettingsHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.RateLimitServerSettings)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerSettings handler received event for %T", object)
	}
	return h.handler.DeleteRateLimitServerSettings(obj)
}

func (h genericRateLimitServerSettingsHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.RateLimitServerSettings)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerSettings handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.RateLimitServerSettings)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerSettings handler received event for %T", new)
	}
	return h.handler.UpdateRateLimitServerSettings(objOld, objNew)
}

func (h genericRateLimitServerSettingsHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.RateLimitServerSettings)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerSettings handler received event for %T", object)
	}
	return h.handler.GenericRateLimitServerSettings(obj)
}

// Handle events for the RateLimitServerConfig Resource
// DEPRECATED: Prefer reconciler pattern.
type RateLimitServerConfigEventHandler interface {
	CreateRateLimitServerConfig(obj *admin_gloo_solo_io_v2.RateLimitServerConfig) error
	UpdateRateLimitServerConfig(old, new *admin_gloo_solo_io_v2.RateLimitServerConfig) error
	DeleteRateLimitServerConfig(obj *admin_gloo_solo_io_v2.RateLimitServerConfig) error
	GenericRateLimitServerConfig(obj *admin_gloo_solo_io_v2.RateLimitServerConfig) error
}

type RateLimitServerConfigEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.RateLimitServerConfig) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.RateLimitServerConfig) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.RateLimitServerConfig) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.RateLimitServerConfig) error
}

func (f *RateLimitServerConfigEventHandlerFuncs) CreateRateLimitServerConfig(obj *admin_gloo_solo_io_v2.RateLimitServerConfig) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RateLimitServerConfigEventHandlerFuncs) DeleteRateLimitServerConfig(obj *admin_gloo_solo_io_v2.RateLimitServerConfig) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RateLimitServerConfigEventHandlerFuncs) UpdateRateLimitServerConfig(objOld, objNew *admin_gloo_solo_io_v2.RateLimitServerConfig) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RateLimitServerConfigEventHandlerFuncs) GenericRateLimitServerConfig(obj *admin_gloo_solo_io_v2.RateLimitServerConfig) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RateLimitServerConfigEventWatcher interface {
	AddEventHandler(ctx context.Context, h RateLimitServerConfigEventHandler, predicates ...predicate.Predicate) error
}

type rateLimitServerConfigEventWatcher struct {
	watcher events.EventWatcher
}

func NewRateLimitServerConfigEventWatcher(name string, mgr manager.Manager) RateLimitServerConfigEventWatcher {
	return &rateLimitServerConfigEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.RateLimitServerConfig{}),
	}
}

func (c *rateLimitServerConfigEventWatcher) AddEventHandler(ctx context.Context, h RateLimitServerConfigEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRateLimitServerConfigHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRateLimitServerConfigHandler implements a generic events.EventHandler
type genericRateLimitServerConfigHandler struct {
	handler RateLimitServerConfigEventHandler
}

func (h genericRateLimitServerConfigHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.RateLimitServerConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerConfig handler received event for %T", object)
	}
	return h.handler.CreateRateLimitServerConfig(obj)
}

func (h genericRateLimitServerConfigHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.RateLimitServerConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerConfig handler received event for %T", object)
	}
	return h.handler.DeleteRateLimitServerConfig(obj)
}

func (h genericRateLimitServerConfigHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.RateLimitServerConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerConfig handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.RateLimitServerConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerConfig handler received event for %T", new)
	}
	return h.handler.UpdateRateLimitServerConfig(objOld, objNew)
}

func (h genericRateLimitServerConfigHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.RateLimitServerConfig)
	if !ok {
		return errors.Errorf("internal error: RateLimitServerConfig handler received event for %T", object)
	}
	return h.handler.GenericRateLimitServerConfig(obj)
}

// Handle events for the Dashboard Resource
// DEPRECATED: Prefer reconciler pattern.
type DashboardEventHandler interface {
	CreateDashboard(obj *admin_gloo_solo_io_v2.Dashboard) error
	UpdateDashboard(old, new *admin_gloo_solo_io_v2.Dashboard) error
	DeleteDashboard(obj *admin_gloo_solo_io_v2.Dashboard) error
	GenericDashboard(obj *admin_gloo_solo_io_v2.Dashboard) error
}

type DashboardEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.Dashboard) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.Dashboard) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.Dashboard) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.Dashboard) error
}

func (f *DashboardEventHandlerFuncs) CreateDashboard(obj *admin_gloo_solo_io_v2.Dashboard) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *DashboardEventHandlerFuncs) DeleteDashboard(obj *admin_gloo_solo_io_v2.Dashboard) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *DashboardEventHandlerFuncs) UpdateDashboard(objOld, objNew *admin_gloo_solo_io_v2.Dashboard) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *DashboardEventHandlerFuncs) GenericDashboard(obj *admin_gloo_solo_io_v2.Dashboard) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type DashboardEventWatcher interface {
	AddEventHandler(ctx context.Context, h DashboardEventHandler, predicates ...predicate.Predicate) error
}

type dashboardEventWatcher struct {
	watcher events.EventWatcher
}

func NewDashboardEventWatcher(name string, mgr manager.Manager) DashboardEventWatcher {
	return &dashboardEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.Dashboard{}),
	}
}

func (c *dashboardEventWatcher) AddEventHandler(ctx context.Context, h DashboardEventHandler, predicates ...predicate.Predicate) error {
	handler := genericDashboardHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericDashboardHandler implements a generic events.EventHandler
type genericDashboardHandler struct {
	handler DashboardEventHandler
}

func (h genericDashboardHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.Dashboard)
	if !ok {
		return errors.Errorf("internal error: Dashboard handler received event for %T", object)
	}
	return h.handler.CreateDashboard(obj)
}

func (h genericDashboardHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.Dashboard)
	if !ok {
		return errors.Errorf("internal error: Dashboard handler received event for %T", object)
	}
	return h.handler.DeleteDashboard(obj)
}

func (h genericDashboardHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.Dashboard)
	if !ok {
		return errors.Errorf("internal error: Dashboard handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.Dashboard)
	if !ok {
		return errors.Errorf("internal error: Dashboard handler received event for %T", new)
	}
	return h.handler.UpdateDashboard(objOld, objNew)
}

func (h genericDashboardHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.Dashboard)
	if !ok {
		return errors.Errorf("internal error: Dashboard handler received event for %T", object)
	}
	return h.handler.GenericDashboard(obj)
}

// Handle events for the IstioLifecycleManager Resource
// DEPRECATED: Prefer reconciler pattern.
type IstioLifecycleManagerEventHandler interface {
	CreateIstioLifecycleManager(obj *admin_gloo_solo_io_v2.IstioLifecycleManager) error
	UpdateIstioLifecycleManager(old, new *admin_gloo_solo_io_v2.IstioLifecycleManager) error
	DeleteIstioLifecycleManager(obj *admin_gloo_solo_io_v2.IstioLifecycleManager) error
	GenericIstioLifecycleManager(obj *admin_gloo_solo_io_v2.IstioLifecycleManager) error
}

type IstioLifecycleManagerEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.IstioLifecycleManager) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.IstioLifecycleManager) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.IstioLifecycleManager) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.IstioLifecycleManager) error
}

func (f *IstioLifecycleManagerEventHandlerFuncs) CreateIstioLifecycleManager(obj *admin_gloo_solo_io_v2.IstioLifecycleManager) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *IstioLifecycleManagerEventHandlerFuncs) DeleteIstioLifecycleManager(obj *admin_gloo_solo_io_v2.IstioLifecycleManager) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *IstioLifecycleManagerEventHandlerFuncs) UpdateIstioLifecycleManager(objOld, objNew *admin_gloo_solo_io_v2.IstioLifecycleManager) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *IstioLifecycleManagerEventHandlerFuncs) GenericIstioLifecycleManager(obj *admin_gloo_solo_io_v2.IstioLifecycleManager) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type IstioLifecycleManagerEventWatcher interface {
	AddEventHandler(ctx context.Context, h IstioLifecycleManagerEventHandler, predicates ...predicate.Predicate) error
}

type istioLifecycleManagerEventWatcher struct {
	watcher events.EventWatcher
}

func NewIstioLifecycleManagerEventWatcher(name string, mgr manager.Manager) IstioLifecycleManagerEventWatcher {
	return &istioLifecycleManagerEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.IstioLifecycleManager{}),
	}
}

func (c *istioLifecycleManagerEventWatcher) AddEventHandler(ctx context.Context, h IstioLifecycleManagerEventHandler, predicates ...predicate.Predicate) error {
	handler := genericIstioLifecycleManagerHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericIstioLifecycleManagerHandler implements a generic events.EventHandler
type genericIstioLifecycleManagerHandler struct {
	handler IstioLifecycleManagerEventHandler
}

func (h genericIstioLifecycleManagerHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.IstioLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: IstioLifecycleManager handler received event for %T", object)
	}
	return h.handler.CreateIstioLifecycleManager(obj)
}

func (h genericIstioLifecycleManagerHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.IstioLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: IstioLifecycleManager handler received event for %T", object)
	}
	return h.handler.DeleteIstioLifecycleManager(obj)
}

func (h genericIstioLifecycleManagerHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.IstioLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: IstioLifecycleManager handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.IstioLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: IstioLifecycleManager handler received event for %T", new)
	}
	return h.handler.UpdateIstioLifecycleManager(objOld, objNew)
}

func (h genericIstioLifecycleManagerHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.IstioLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: IstioLifecycleManager handler received event for %T", object)
	}
	return h.handler.GenericIstioLifecycleManager(obj)
}

// Handle events for the GatewayLifecycleManager Resource
// DEPRECATED: Prefer reconciler pattern.
type GatewayLifecycleManagerEventHandler interface {
	CreateGatewayLifecycleManager(obj *admin_gloo_solo_io_v2.GatewayLifecycleManager) error
	UpdateGatewayLifecycleManager(old, new *admin_gloo_solo_io_v2.GatewayLifecycleManager) error
	DeleteGatewayLifecycleManager(obj *admin_gloo_solo_io_v2.GatewayLifecycleManager) error
	GenericGatewayLifecycleManager(obj *admin_gloo_solo_io_v2.GatewayLifecycleManager) error
}

type GatewayLifecycleManagerEventHandlerFuncs struct {
	OnCreate  func(obj *admin_gloo_solo_io_v2.GatewayLifecycleManager) error
	OnUpdate  func(old, new *admin_gloo_solo_io_v2.GatewayLifecycleManager) error
	OnDelete  func(obj *admin_gloo_solo_io_v2.GatewayLifecycleManager) error
	OnGeneric func(obj *admin_gloo_solo_io_v2.GatewayLifecycleManager) error
}

func (f *GatewayLifecycleManagerEventHandlerFuncs) CreateGatewayLifecycleManager(obj *admin_gloo_solo_io_v2.GatewayLifecycleManager) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *GatewayLifecycleManagerEventHandlerFuncs) DeleteGatewayLifecycleManager(obj *admin_gloo_solo_io_v2.GatewayLifecycleManager) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *GatewayLifecycleManagerEventHandlerFuncs) UpdateGatewayLifecycleManager(objOld, objNew *admin_gloo_solo_io_v2.GatewayLifecycleManager) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *GatewayLifecycleManagerEventHandlerFuncs) GenericGatewayLifecycleManager(obj *admin_gloo_solo_io_v2.GatewayLifecycleManager) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type GatewayLifecycleManagerEventWatcher interface {
	AddEventHandler(ctx context.Context, h GatewayLifecycleManagerEventHandler, predicates ...predicate.Predicate) error
}

type gatewayLifecycleManagerEventWatcher struct {
	watcher events.EventWatcher
}

func NewGatewayLifecycleManagerEventWatcher(name string, mgr manager.Manager) GatewayLifecycleManagerEventWatcher {
	return &gatewayLifecycleManagerEventWatcher{
		watcher: events.NewWatcher(name, mgr, &admin_gloo_solo_io_v2.GatewayLifecycleManager{}),
	}
}

func (c *gatewayLifecycleManagerEventWatcher) AddEventHandler(ctx context.Context, h GatewayLifecycleManagerEventHandler, predicates ...predicate.Predicate) error {
	handler := genericGatewayLifecycleManagerHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericGatewayLifecycleManagerHandler implements a generic events.EventHandler
type genericGatewayLifecycleManagerHandler struct {
	handler GatewayLifecycleManagerEventHandler
}

func (h genericGatewayLifecycleManagerHandler) Create(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.GatewayLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: GatewayLifecycleManager handler received event for %T", object)
	}
	return h.handler.CreateGatewayLifecycleManager(obj)
}

func (h genericGatewayLifecycleManagerHandler) Delete(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.GatewayLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: GatewayLifecycleManager handler received event for %T", object)
	}
	return h.handler.DeleteGatewayLifecycleManager(obj)
}

func (h genericGatewayLifecycleManagerHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*admin_gloo_solo_io_v2.GatewayLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: GatewayLifecycleManager handler received event for %T", old)
	}
	objNew, ok := new.(*admin_gloo_solo_io_v2.GatewayLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: GatewayLifecycleManager handler received event for %T", new)
	}
	return h.handler.UpdateGatewayLifecycleManager(objOld, objNew)
}

func (h genericGatewayLifecycleManagerHandler) Generic(object client.Object) error {
	obj, ok := object.(*admin_gloo_solo_io_v2.GatewayLifecycleManager)
	if !ok {
		return errors.Errorf("internal error: GatewayLifecycleManager handler received event for %T", object)
	}
	return h.handler.GenericGatewayLifecycleManager(obj)
}
