package main

import (
	"context"
	"fmt"
	"time"

	"istio.io/client-go/pkg/apis/networking/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func setupIstio() error {
	err := v1beta1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	return nil
}

func cleanupIstio() {
	c := getK8sConfig()
	kc, err := client.New(c, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	kc.DeleteAllOf(
		ctx,
		&v1beta1.VirtualService{}, client.InNamespace("default"))

	if err != nil {
		fmt.Println("Error removing Istio VirtualService", err)
	}
}

func iExpectIstioNamedToHaveBeenCreated(count int, crd, name string) error {
	c := getK8sConfig()
	kc, err := client.New(c, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		panic(err)
	}

	return waitForComplete(
		30*time.Second,
		func() error {
			switch crd {
			case "VirtualService":
				vs := &v1beta1.VirtualService{}
				return kc.Get(
					context.Background(),
					types.NamespacedName{
						Namespace: "default",
						Name:      name,
					},
					vs,
				)
			}

			return fmt.Errorf("Type %s, is not configured", crd)
		},
	)
}
