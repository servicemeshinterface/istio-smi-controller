package istio

import (
	"context"
	"sync"

	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	networkingv1beta1 "istio.io/api/networking/v1beta1"
	"istio.io/client-go/pkg/apis/networking/v1beta1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Client allows the creation of Istio objects from their SMI counterparts
type Client interface {
	CreateVirtualService(ctx context.Context, r client.Writer, ts *splitv1alpha4.TrafficSplit) error
}

// IstioClient is a concrete implementation of Client
type IstioClient struct {
	once sync.Once
}

// lazyRegisterIstioTypesToScheme lazily registers the Istio types to the
// controllers client. This allows us to use the default client that is created by the
// controller instead of manually creating one.
//
// Since the client is not available until the first time the controller is called
// this function must be called from every Client interface method.
// sync.Once ensures that types are only registered once.
func (ic *IstioClient) lazyRegisterIstioTypesToScheme(c client.Client) {
	ic.once.Do(func() {
		// register types to the scheme
		err := v1beta1.AddToScheme(c.Scheme())
		if err != nil {
			panic(err)
		}
	})
}

// CreateVirtualService creates an Istio VirtualService from the given TrafficSplit
func (ic *IstioClient) CreateVirtualService(ctx context.Context, r client.Writer, ts *splitv1alpha4.TrafficSplit) error {
	ic.lazyRegisterIstioTypesToScheme(r.(client.Client))

	vs := &v1beta1.VirtualService{}

	vs.ObjectMeta.Name = ts.ObjectMeta.Name
	vs.ObjectMeta.Namespace = ts.ObjectMeta.Namespace

	vs.Spec.Hosts = []string{ts.Spec.Service}

	// if matches is not present then create as a HTTPRoute
	httpRoute := &networkingv1beta1.HTTPRoute{}
	httpRoute.Route = []*networkingv1beta1.HTTPRouteDestination{}

	for _, be := range ts.Spec.Backends {
		hrd := &networkingv1beta1.HTTPRouteDestination{}
		hrd.Destination = &networkingv1beta1.Destination{
			Host:   ts.Spec.Service,
			Subset: be.Service,
		}
		hrd.Weight = int32(be.Weight)

		httpRoute.Route = append(httpRoute.Route, hrd)
	}

	vs.Spec.Http = []*networkingv1beta1.HTTPRoute{httpRoute}

	return r.Create(ctx, vs, &client.CreateOptions{})
}
