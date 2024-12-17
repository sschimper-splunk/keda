/*
Copyright 2023 The KEDA Authors

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
//
//nolint:dupl
package eventing

import (
	"context"
	"sync"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	eventingv1alpha1 "github.com/kedacore/keda/v2/apis/eventing/v1alpha1"
	"github.com/kedacore/keda/v2/pkg/eventemitter"
	"github.com/kedacore/keda/v2/pkg/metricscollector"
	"github.com/kedacore/keda/v2/pkg/util"
)

// CloudEventSourceReconciler reconciles a EventSource object
type CloudEventSourceReconciler struct {
	client.Client
	eventEmitter eventemitter.EventHandler

	cloudEventSourceGenerations *sync.Map
	eventSourcePromMetricsMap   map[string]string
	eventSourcePromMetricsLock  *sync.Mutex
}

// NewCloudEventSourceReconciler creates a new CloudEventSourceReconciler
func NewCloudEventSourceReconciler(c client.Client, e eventemitter.EventHandler) *CloudEventSourceReconciler {
	return &CloudEventSourceReconciler{
		Client:                      c,
		eventEmitter:                e,
		cloudEventSourceGenerations: &sync.Map{},
		eventSourcePromMetricsMap:   make(map[string]string),
		eventSourcePromMetricsLock:  &sync.Mutex{},
	}
}

// +kubebuilder:rbac:groups=eventing.keda.sh,resources=cloudeventsources;cloudeventsources/status,verbs=get;list;watch;update;patch

// Reconcile performs reconciliation on the identified EventSource resource based on the request information passed, returns the result and an error (if any).

func (r *CloudEventSourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := log.FromContext(ctx)
	cloudEventSource := &eventingv1alpha1.CloudEventSource{}
	return Reconcile(ctx, reqLogger, r, req, cloudEventSource)
}

// SetupWithManager sets up the controller with the Manager.
func (r *CloudEventSourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&eventingv1alpha1.CloudEventSource{}, builder.WithPredicates(predicate.GenerationChangedPredicate{})).
		WithEventFilter(util.IgnoreOtherNamespaces()).
		Complete(r)
}

func (r *CloudEventSourceReconciler) GetClient() client.Client {
	return r.Client
}

func (r *CloudEventSourceReconciler) GetEventEmitter() eventemitter.EventHandler {
	return r.eventEmitter
}

func (r *CloudEventSourceReconciler) GetCloudEventSourceGeneration() *sync.Map {
	return r.cloudEventSourceGenerations
}

func (r *CloudEventSourceReconciler) UpdatePromMetrics(eventSource eventingv1alpha1.CloudEventSourceInterface, namespacedName string) {
	r.eventSourcePromMetricsLock.Lock()
	defer r.eventSourcePromMetricsLock.Unlock()

	if ns, ok := r.eventSourcePromMetricsMap[namespacedName]; ok {
		metricscollector.DecrementCRDTotal(metricscollector.CloudEventSourceResource, ns)
	}

	metricscollector.IncrementCRDTotal(metricscollector.CloudEventSourceResource, eventSource.GetNamespace())
	r.eventSourcePromMetricsMap[namespacedName] = eventSource.GetNamespace()
}

// UpdatePromMetricsOnDelete is idempotent, so it can be called multiple times without side-effects
func (r *CloudEventSourceReconciler) UpdatePromMetricsOnDelete(namespacedName string) {
	r.eventSourcePromMetricsLock.Lock()
	defer r.eventSourcePromMetricsLock.Unlock()

	if ns, ok := r.eventSourcePromMetricsMap[namespacedName]; ok {
		metricscollector.DecrementCRDTotal(metricscollector.CloudEventSourceResource, ns)
	}

	delete(r.eventSourcePromMetricsMap, namespacedName)
}
