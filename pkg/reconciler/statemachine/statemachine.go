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

package statemachine

import (
	"context"
	statev1 "github.com/salaboy/knative-state2/pkg/apis/state/v1alpha1"
	statemachinereconciler "github.com/salaboy/knative-state2/pkg/client/injection/reconciler/state/v1alpha1/statemachine"
	"go.uber.org/zap"
	"k8s.io/client-go/dynamic"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"
)

type Reconciler struct {
	dynamicClientSet  dynamic.Interface

}

// Check that our Reconciler implements Interface
var _ statemachinereconciler.Interface = (*Reconciler)(nil)

// ReconcilerArgs are the arguments needed to create a broker.Reconciler.
type ReconcilerArgs struct {

}

func (r *Reconciler) ReconcileKind(ctx context.Context, b *statev1.StateMachine) pkgreconciler.Event {
	logging.FromContext(ctx).Infow("Reconciling", zap.Any("StateMachine", b))

	// @TODO: logic here


	return nil
}

