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
	"github.com/prometheus/common/log"
	servingClientSet "knative.dev/serving/pkg/client/clientset/versioned"
	eventingClientSet "knative.dev/eventing/pkg/client/clientset/versioned"
	statev1 "github.com/salaboy/knative-state2/pkg/apis/state/v1alpha1"
	statemachinerunnerreconciler "github.com/salaboy/knative-state2/pkg/client/injection/reconciler/state/v1alpha1/statemachinerunner"
	listers "github.com/salaboy/knative-state2/pkg/client/listers/state/v1alpha1"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/dynamic"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"
	"os"
)

type Reconciler struct {
	servingClientSet servingClientSet.Interface
	eventingClientSet eventingClientSet.Interface

	dynamicClientSet dynamic.Interface
	statemachineLister listers.StateMachineLister

}

// Check that our Reconciler implements Interface
var _ statemachinerunnerreconciler.Interface = (*Reconciler)(nil)

// ReconcilerArgs are the arguments needed to create a broker.Reconciler.
type ReconcilerArgs struct {
}

var RUNNER_IMAGE = os.Getenv("RUNNER_IMAGE")

func (r *Reconciler) ReconcileKind(ctx context.Context, smr *statev1.StateMachineRunner) pkgreconciler.Event {
	logging.FromContext(ctx).Infow("Reconciling", zap.Any("StateMachineRunner", smr))

	logging.FromContext(ctx).Infow("StateMachineRunner with parameters: ", "sink", smr.Spec.Sink, "ref", smr.Spec.StateMachineRef)

	if smr.Spec.StateMachineRef != "" {
		stateMachine, _ := r.GetStateMachine(ctx, smr)
		logging.FromContext(ctx).Infow("StateMachine reference found: ",  zap.Any("stateMachine",  stateMachine))

		yamlStates, err := yaml.Marshal(stateMachine.Spec.StateMachineDefinition.StateMachineStates.States)
		if err != nil {
			log.Error(err, "failed to parse yaml from statemachine definition states")
			return err
		}

		if RUNNER_IMAGE == "" {
			RUNNER_IMAGE = "kind.local/knative-statemachine-runner-7a3c815d2bf3ebf9af9650f7624a29c9:93b9adcd6af50be3ba7f7b4848c79da214c0b4dcca39709c98c28905eb91b6a0"
		}

	}

	return nil
}

func (r *Reconciler) GetStateMachine(ctx context.Context,smr *statev1.StateMachineRunner ) (*statev1.StateMachine, pkgreconciler.Event){

	stateMachine, err := r.statemachineLister.StateMachines(smr.Namespace).Get(smr.Spec.StateMachineRef)
	if err != nil {
		return nil, err
	}


	return stateMachine, nil
}
