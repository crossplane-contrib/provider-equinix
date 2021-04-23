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

package fabric

// import (
// 	"context"

// 	ecx "github.com/equinix/ecx-go"
// 	oauth2 "github.com/equinix/oauth2-go"
// 	"github.com/pkg/errors"
// 	"k8s.io/apimachinery/pkg/types"
// 	"k8s.io/client-go/util/workqueue"
// 	ctrl "sigs.k8s.io/controller-runtime"
// 	"sigs.k8s.io/controller-runtime/pkg/client"
// 	"sigs.k8s.io/controller-runtime/pkg/controller"

// 	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
// 	"github.com/crossplane/crossplane-runtime/pkg/event"
// 	"github.com/crossplane/crossplane-runtime/pkg/logging"
// 	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
// 	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
// 	"github.com/crossplane/crossplane-runtime/pkg/resource"

// 	"github.com/crossplane-contrib/provider-equinix/v1beta1"
// )

// const (
// 	errNotConnection         = "managed resource is not a Connection custom resource"
// 	errManagedUpdateFailed = "cannot update Connection custom resource"

// 	errNewClient        = "cannot create new Sqladmin Service"
// 	errCreateFailed     = "cannot create new Connection instance"
// 	errNameInUse        = "cannot create new Connection instance, resource name is unavailable because it is in use or was used recently"
// 	errDeleteFailed     = "cannot delete the Connection instance"
// 	errUpdateFailed     = "cannot update the Connection instance"
// 	errGetFailed        = "cannot get the Connection instance"
// 	errGeneratePassword = "cannot generate root password"
// 	errCheckUpToDate    = "cannot determine if Connection instance is up to date"
// )

// // SetupConnection adds a controller that reconciles
// // Connection managed resources.
// func SetupConnection(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
// 	name := managed.ControllerName(v1beta1.ConnectionGroupKind)

// 	r := managed.NewReconciler(mgr,
// 		resource.ManagedKind(v1beta1.ConnectionGroupVersionKind),
// 		managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
// 		managed.WithInitializers(managed.NewDefaultProviderConfig(mgr.GetClient()), managed.NewNameAsExternalName(mgr.GetClient())),
// 		managed.WithReferenceResolver(managed.NewAPISimpleReferenceResolver(mgr.GetClient())),
// 		managed.WithLogger(l.WithValues("controller", name)),
// 		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

// 	return ctrl.NewControllerManagedBy(mgr).
// 		Named(name).
// 		WithOptions(controller.Options{
// 			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
// 		}).
// 		For(&v1beta1.Connection{}).
// 		Complete(r)
// }

// type connector struct {
// 	kube client.Client
// }

// func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
// 	pc := &v1beta1.ProviderConfig{}
// 	t := resource.NewProviderConfigUsageTracker(c.kube, &v1beta1.ProviderConfigUsage{})
// 	if err := t.Track(ctx, mg); err != nil {
// 		return nil, err
// 	}
// 	if err := c.kube.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
// 		return nil, err
// 	}
// 	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c, pc.Spec.Credentials.CommonCredentialSelectors)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "cannot get credentials")
// 	}
// 	authConfig := oauth2.Config{
// 		ClientID:     "someClientId",
// 		ClientSecret: "someSecret",
// 		BaseURL:      "baseURL",
// 	}
// 	authClient := authConfig.New(ctx)
// 	return &connectionExternal{kube: c.kube, ecx: ecx.NewClient(ctx, "", authClient)}, nil
// }

// type connectionExternal struct {
// 	kube client.Client
// 	ecx  ecx.Client
// }

// func (c *connectionExternal) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
// 	cr, ok := mg.(*v1beta1.Connection)
// 	if !ok {
// 		return managed.ExternalObservation{}, errors.New(errNotConnection)
// 	}
// 	return managed.ExternalObservation{
// 		ResourceExists:   true,
// 		ResourceUpToDate: true,
// 	}, nil
// }

// func (c *connectionExternal) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
// 	cr, ok := mg.(*v1beta1.Connection)
// 	if !ok {
// 		return managed.ExternalCreation{}, errors.New(errNotConnection)
// 	}
// 	cr.SetConditions(xpv1.Creating())
// 	return managed.ExternalCreation{}, nil
// }

// func (c *connectionExternal) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
// 	_, ok := mg.(*v1beta1.Connection)
// 	if !ok {
// 		return managed.ExternalUpdate{}, errors.New(errNotConnection)
// 	}
// 	return managed.ExternalUpdate{}, errors.Wrap(err, errUpdateFailed)
// }

// func (c *connectionExternal) Delete(ctx context.Context, mg resource.Managed) error {
// 	cr, ok := mg.(*v1beta1.Connection)
// 	if !ok {
// 		return errors.New(errNotConnection)
// 	}
// 	cr.SetConditions(xpv1.Deleting())
// 	return errors.Wrap(nil, errDeleteFailed)
// }
