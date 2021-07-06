package istio

import (
	"context"
	"testing"

	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
)

func setupClientTests(t *testing.T) (*IstioClient, *MockControllerClient) {
	return &IstioClient{}, &MockControllerClient{}

}

func TestCreateVirtualServiceCallsCreateWithValidObject(t *testing.T) {
	is, mcc := setupClientTests(t)
	mcc.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ts := &splitv1alpha4.TrafficSplit{}

	is.CreateVirtualService(context.Background(), mcc, ts)

	vs := getCalls(&mcc.Mock, "Create")[0].Arguments[1].(*v1alpha3.VirtualService)

	require.Equal(t, vs.ObjectMeta.Namespace, ts.ObjectMeta.Namespace)
}
