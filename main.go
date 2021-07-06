package main

import (
	"github.com/nicholasjackson/istio-smi-controller/istio"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk/controller"
)

func main() {
	// register our lifecycle callbacks with the controller
	sdk.API().RegisterV1Alpha(&istio.API{})

	// create and start a the controller
	config := controller.DefaultConfig()
	controller.Start(config)
}
