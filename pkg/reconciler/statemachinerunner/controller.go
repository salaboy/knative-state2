/*
Copyright 2021 The Knative Authors
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

package statemachinerunner

import (
	"context"
	statemachineinformer "github.com/salaboy/knative-state2/pkg/client/injection/informers/state/v1alpha1/statemachine"
	statemachinerunnerinformer "github.com/salaboy/knative-state2/pkg/client/injection/informers/state/v1alpha1/statemachinerunner"
	statemachinerunnerreconciler "github.com/salaboy/knative-state2/pkg/client/injection/reconciler/state/v1alpha1/statemachinerunner"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection/clients/dynamicclient"
	"knative.dev/pkg/logging"
	eventingclient "knative.dev/eventing/pkg/client/injection/client"
	servingclient "knative.dev/serving/pkg/client/injection/client"
)

// NewController initializes the controller and is called by the generated code
// Registers event handlers to enqueue events
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)
	stateMachineRunnerInformer := statemachinerunnerinformer.Get(ctx)
	stateMachineInformer := statemachineinformer.Get(ctx)

	r := &Reconciler{
		dynamicClientSet:   dynamicclient.Get(ctx),
		statemachineLister: stateMachineInformer.Lister(),
		eventingClientSet:  eventingclient.Get(ctx),
		servingClientSet:  servingclient.Get(ctx),
	}
	impl := statemachinerunnerreconciler.NewImpl(ctx, r)

	logger.Info("Setting up event handlers for StateMachineRunners")

	stateMachineRunnerInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))


	return impl
}
