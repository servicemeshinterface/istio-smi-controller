package istio

import (
	"context"

	"github.com/go-logr/logr"
	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/* UpsertTrafficSplit converts a Service Mesh Interface Traffic split into an Istio Virtual Service
---
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: tcp-echo-destination
spec:
  host: tcp-echo
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: tcp-echo
spec:
  hosts:
  - "*"
  gateways:
  - tcp-echo-gateway
  tcp:
  - match:
    - port: 31400
    route:
    - destination:
        host: tcp-echo
        port:
          number: 9000
        subset: v1
*/
func (l *API) UpsertTrafficSplit(
	ctx context.Context,
	r client.Client,
	log logr.Logger,
	ts *splitv1alpha4.TrafficSplit) (ctrl.Result, error) {

	log.Info("UpdateTrafficSplit called", "api", "v1alpha4", "target", ts)

	err := l.client.CreateVirtualService(ctx, r, ts)
	if err != nil {
		log.Error(err, "Unable to create Istio VirtualService")

		return ctrl.Result{Requeue: true}, err
	}

	return ctrl.Result{}, nil
}

func (l *API) DeleteTrafficSplit(
	ctx context.Context,
	r client.Client,
	log logr.Logger,
	ts *splitv1alpha4.TrafficSplit) (ctrl.Result, error) {

	log.Info("DeleteTrafficSplit called", "api", "v1alpha4", "target", ts)

	err := l.client.DeleteVirtualService(ctx, r, ts)
	if err != nil {
		log.Error(err, "Unable to delete Istio VirtualService")

		return ctrl.Result{Requeue: true}, err
	}

	return ctrl.Result{}, nil
}
