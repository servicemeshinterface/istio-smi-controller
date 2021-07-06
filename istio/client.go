package istio

import (
	"context"

	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Client allows the creation of Istio objects from their SMI counterparts
type Client interface {
	CreateVirtualService(ctx context.Context, r client.Writer, ts *splitv1alpha4.TrafficSplit) error
}

// IstioClient is a concrete implementation of Client
type IstioClient struct {
}

// CreateVirtualService creates an Istio VirtualService from the given TrafficSplit
func (ic *IstioClient) CreateVirtualService(ctx context.Context, r client.Writer, ts *splitv1alpha4.TrafficSplit) error {
	vs := &v1alpha3.VirtualService{}
	return r.Create(ctx, vs, &client.CreateOptions{})
}
