package istio

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"github.com/stretchr/testify/mock"
)

func TestCreateVirtualService(t *testing.T) {
	api, mc, mcc := setupTests(t)
	mc.On("CreateVirtualService", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx := context.Background()
	ts := &splitv1alpha4.TrafficSplit{}
	lgr := &logr.DiscardLogger{}

	api.UpsertTrafficSplit(ctx, mcc, lgr, ts)

	mc.AssertCalled(t, "CreateVirtualService", ctx, mcc, ts)
}
