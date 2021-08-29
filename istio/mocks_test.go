package istio

import (
	"context"

	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"github.com/stretchr/testify/mock"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MockClient struct {
	mock.Mock
}

func (mc *MockClient) CreateVirtualService(ctx context.Context, r client.Writer, ts *splitv1alpha4.TrafficSplit) error {
	args := mc.Called(ctx, r, ts)
	return args.Error(0)
}

func (mc *MockClient) DeleteVirtualService(ctx context.Context, r client.Writer, ts *splitv1alpha4.TrafficSplit) error {
	args := mc.Called(ctx, r, ts)
	return args.Error(0)
}

// MockWriter is a mock implementation of the contoller Writer interface
type MockControllerClient struct {
	mock.Mock
}

func (mcc *MockControllerClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object) error {
	args := mcc.Called(ctx, key, obj)
	return args.Error(0)
}

func (mcc *MockControllerClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	args := mcc.Called(ctx, list, opts)
	return args.Error(0)
}

func (mcc *MockControllerClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	args := mcc.Called(ctx, obj, opts)
	return args.Error(0)
}

func (mcc *MockControllerClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	args := mcc.Called(ctx, obj, opts)
	return args.Error(0)
}

func (mcc *MockControllerClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	args := mcc.Called(ctx, obj, opts)
	return args.Error(0)
}

func (mcc *MockControllerClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	args := mcc.Called(ctx, obj, patch, opts)
	return args.Error(0)
}

func (mcc *MockControllerClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	args := mcc.Called(ctx, obj, opts)
	return args.Error(0)
}

func (mcc *MockControllerClient) Status() client.StatusWriter {
	args := mcc.Called()
	return args.Get(0).(client.StatusWriter)
}

func (mcc *MockControllerClient) Scheme() *runtime.Scheme {
	args := mcc.Called()
	return args.Get(0).(*runtime.Scheme)
}

func (mcc *MockControllerClient) RESTMapper() meta.RESTMapper {
	args := mcc.Called()
	return args.Get(0).(meta.RESTMapper)
}

func getCalls(m *mock.Mock, method string) []mock.Call {
	rc := make([]mock.Call, 0)
	for _, c := range m.Calls {
		if c.Method == method {
			rc = append(rc, c)
		}
	}

	return rc
}
