package istio

import (
	"testing"
)

func setupTests(t *testing.T) (*API, *MockClient, *MockControllerClient) {
	mc := &MockClient{}
	mcc := &MockControllerClient{}
	api := New(mc)

	return api, mc, mcc
}
