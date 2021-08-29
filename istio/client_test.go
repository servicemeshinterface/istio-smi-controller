package istio

import (
	"context"
	"testing"

	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

func setupClientTests(t *testing.T) (*IstioClient, *MockControllerClient) {
	mcc := &MockControllerClient{}
	mcc.On("Scheme").Return(runtime.NewScheme())

	return &IstioClient{}, mcc
}

func TestCreateVirtualServiceCallsCreateWithValidObject(t *testing.T) {
	is, mcc := setupClientTests(t)
	mcc.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ts := &splitv1alpha4.TrafficSplit{}

	is.CreateVirtualService(context.Background(), mcc, ts)

	vs := getCalls(&mcc.Mock, "Create")[0].Arguments[1].(*v1beta1.VirtualService)

	require.Equal(t, vs.ObjectMeta.Name, ts.ObjectMeta.Name)
	require.Equal(t, vs.ObjectMeta.Namespace, ts.ObjectMeta.Namespace)

	require.Equal(t, vs.Spec.Hosts[0], ts.Spec.Service)

	// check the backends are converted into a HTTPRouteDestination
	require.Len(t, vs.Spec.Http[0].Route, len(ts.Spec.Backends))

	for i, be := range ts.Spec.Backends {
		require.Equal(t, ts.Spec.Service, vs.Spec.Http[0].Route[i].Destination.Host)
		require.Equal(t, be.Weight, vs.Spec.Http[0].Route[i].Weight)
		require.Equal(t, be.Service, vs.Spec.Http[0].Route[i].Destination.Subset)
	}
}

func TestDeleteVirtualServiceCallsCreateWithValidObject(t *testing.T) {
	is, mcc := setupClientTests(t)
	mcc.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ts := &splitv1alpha4.TrafficSplit{}

	is.DeleteVirtualService(context.Background(), mcc, ts)

	vs := getCalls(&mcc.Mock, "Delete")[0].Arguments[1].(*v1beta1.VirtualService)

	require.Equal(t, vs.ObjectMeta.Name, ts.ObjectMeta.Name)
	require.Equal(t, vs.ObjectMeta.Namespace, ts.ObjectMeta.Namespace)

}
