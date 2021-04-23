/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package network

import (
	"context"
	"encoding/json"

	ne "github.com/equinix/ne-go"
	oauth2 "github.com/equinix/oauth2-go"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	netv1alpha1 "github.com/crossplane-contrib/provider-equinix/apis/network/v1alpha1"
	"github.com/crossplane-contrib/provider-equinix/apis/v1beta1"
)

var _ ne.Device = ne.Device{}

type creds struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

const (
	errNotDevice = "managed resource is not a Device custom resource"

	errDeleteFailed = "cannot delete the Device"
	errUpdateFailed = "cannot update the Device"
)

// SetupDevice adds a controller that reconciles
// Device managed resources.
func SetupDevice(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(netv1alpha1.DeviceGroupKind)

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(netv1alpha1.DeviceGroupVersionKind),
		managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
		managed.WithInitializers(managed.NewDefaultProviderConfig(mgr.GetClient()), managed.NewNameAsExternalName(mgr.GetClient())),
		managed.WithReferenceResolver(managed.NewAPISimpleReferenceResolver(mgr.GetClient())),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&netv1alpha1.Device{}).
		Complete(r)
}

type connector struct {
	kube client.Client
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	pc := &v1beta1.ProviderConfig{}
	t := resource.NewProviderConfigUsageTracker(c.kube, &v1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, err
	}
	if err := c.kube.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, err
	}
	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c.kube, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get credentials")
	}
	creds := &creds{}
	if err := json.Unmarshal(data, creds); err != nil {
		return nil, err
	}
	authConfig := oauth2.Config{
		ClientID:     creds.ClientID,
		ClientSecret: creds.ClientSecret,
		BaseURL:      "https://api.equinix.com",
	}
	return &deviceExternal{kube: c.kube, ne: ne.NewClient(ctx, "https://api.equinix.com", authConfig.New(ctx))}, nil
}

type deviceExternal struct {
	kube client.Client
	ne   ne.Client
}

func (c *deviceExternal) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*netv1alpha1.Device)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotDevice)
	}
	dev, err := c.ne.GetDevice(meta.GetExternalName(cr))
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	cr.Spec.ForProvider.Name = dev.Name
	if err := c.kube.Update(ctx, cr); err != nil {
		return managed.ExternalObservation{}, err
	}
	// dev, err ne.GetDevice(meta.GetExternalName(cr))
	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

func (c *deviceExternal) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*netv1alpha1.Device)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotDevice)
	}
	cr.SetConditions(xpv1.Creating())
	return managed.ExternalCreation{}, nil
}

func (c *deviceExternal) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	_, ok := mg.(*netv1alpha1.Device)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotDevice)
	}
	err := errors.New("not implemented")
	return managed.ExternalUpdate{}, errors.Wrap(err, errUpdateFailed)
}

func (c *deviceExternal) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*netv1alpha1.Device)
	if !ok {
		return errors.New(errNotDevice)
	}
	cr.SetConditions(xpv1.Deleting())
	return errors.Wrap(nil, errDeleteFailed)
}
