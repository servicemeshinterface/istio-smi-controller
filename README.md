# SMI Compliant Controller for Istio  Service Mesh

This repository implements the Service Mesh Interface SDK for Istio.

## Required software
To build the Istio controller, you will need the following software installed.

* Go [https://golang.org/doc/install](https://golang.org/doc/install)
* Docker [https://www.docker.com/get-started](https://www.docker.com/get-started)
* Shipyard [https://shipyard.run/docs/install](https://shipyard.run/docs/install)

## Running unit tests

Running the unit tests requies no special software, you can run `go test -v -race ./...`

The are no dependencies on these tests as they only test the API and not the main controller.

```
?       github.com/nicholasjackson/istio-smi-controller [no test files]
=== RUN   TestCreateVirtualServiceCallsCreateWithValidObject
--- PASS: TestCreateVirtualServiceCallsCreateWithValidObject (0.01s)
=== RUN   TestCreateVirtualService
--- PASS: TestCreateVirtualService (0.00s)
PASS
ok      github.com/nicholasjackson/istio-smi-controller/istio   0.198s
?       github.com/nicholasjackson/istio-smi-controller/test    [no test files]
```

## Running functional tests
To functionally test the controller, the following steps can be used. 

1. Start a development version of Kubernetes in Docker with Shipyard. 
```shell
➜ shipyard run ./shipyard
Running configuration from:  ./shipyard

2021-07-07T13:41:49.454+0100 [INFO]  Generating template: ref=smi_controller_config output=/home/nicj/.shipyard/data/smi_controller/smi-controller-values.yaml
2021-07-07T13:41:49.454+0100 [INFO]  Creating Network: ref=dc1
2021-07-07T13:41:49.454+0100 [INFO]  Creating Output: ref=KUBECONFIG
2021-07-07T13:41:49.479+0100 [INFO]  Creating ImageCache: ref=docker-cache

...

########################################################

Title Consul Service Mesh on Kubernetes with Monitoring
Author Nic Jackson

shipyard_version: ">= 0.2.1"

1 Service Mesh Interface Controller SDK
────────────────────────────────────────────────────────────────────────────────

This blueprint creates a Kubernetes cluster and installs the following elements:
• Cert Manager
• SMI Controller CRDs and webhook config
• Local ingress exposing port 9443 for webhook to local machine

This blueprint defines 1 output variables.

You can set output variables as environment variables for your current terminal session using the following command:

eval $(shipyard env)

To list output variables use the command:

shipyard output
```

2. Shipyard installs the components for the controler such as the CRDs and creates the certificates and webhook config needed to be run but does not install the controller. Traffic destined for the webhook service `smi-webhook` to your local environment. This allows the functional tests to be run and converions webhooks to be tested without needing to install the controller on Kubernetes.

3. Once running you can set the `KUBECONFIG` environment variable so that the controller can contact the Kubernetes API using the following command.

```
eval $(shipyard env)
```

4. You can test that the cluster is working and everything is installed:

```shell
kubectl get pods -n smi
NAME                                       READY   STATUS    RESTARTS   AGE
cert-manager-cainjector-669495854f-6ccnv   1/1     Running   0          4m57s
cert-manager-f77c469b-mqncg                1/1     Running   0          4m57s
cert-manager-webhook-74ddb4696d-cvmqw      1/1     Running   0          4m57s
```

```shell
➜ kubectl get svc -n smi
NAME                   TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
cert-manager-webhook   ClusterIP   10.43.10.42     <none>        443/TCP    5m29s
cert-manager           ClusterIP   10.43.170.237   <none>        9402/TCP   5m29s
```

```
➜ kubectl get crds -n smi
NAME                                       CREATED AT
addons.k3s.cattle.io                       2021-07-07T12:41:55Z
helmcharts.helm.cattle.io                  2021-07-07T12:41:55Z
helmchartconfigs.helm.cattle.io            2021-07-07T12:41:55Z
certificaterequests.cert-manager.io        2021-07-07T12:42:38Z
certificates.cert-manager.io               2021-07-07T12:42:38Z
orders.acme.cert-manager.io                2021-07-07T12:42:38Z
challenges.acme.cert-manager.io            2021-07-07T12:42:38Z
issuers.cert-manager.io                    2021-07-07T12:42:38Z
clusterissuers.cert-manager.io             2021-07-07T12:42:38Z
authorizationpolicies.security.istio.io    2021-07-07T12:42:39Z
destinationrules.networking.istio.io       2021-07-07T12:42:39Z
envoyfilters.networking.istio.io           2021-07-07T12:42:39Z
gateways.networking.istio.io               2021-07-07T12:42:39Z
istiooperators.install.istio.io            2021-07-07T12:42:39Z
peerauthentications.security.istio.io      2021-07-07T12:42:39Z
requestauthentications.security.istio.io   2021-07-07T12:42:39Z
serviceentries.networking.istio.io         2021-07-07T12:42:39Z
sidecars.networking.istio.io               2021-07-07T12:42:39Z
telemetries.telemetry.istio.io             2021-07-07T12:42:39Z
virtualservices.networking.istio.io        2021-07-07T12:42:39Z
workloadentries.networking.istio.io        2021-07-07T12:42:39Z
workloadgroups.networking.istio.io         2021-07-07T12:42:39Z
udproutes.specs.smi-spec.io                2021-07-07T12:42:47Z
tcproutes.specs.smi-spec.io                2021-07-07T12:42:47Z
httproutegroups.specs.smi-spec.io          2021-07-07T12:42:47Z
traffictargets.access.smi-spec.io          2021-07-07T12:42:47Z
trafficsplits.split.smi-spec.io            2021-07-07T12:42:47Z
```

5. Finally run the tests

```shell
make functional_tests
```

The project uses Cucumber for Go to run the functional tests, you can find the test setup and the feature files in the `./test` subfolder. To run the test we start a local instance of the controller and connect it to the
Docker based Kubernetes cluster.  SMI resources are then applied to the cluster using the Kubernetes API and 
assertions that the correct Istio resources have been created.

```shell
➜ make functional_test
mkdir -p /tmp/k8s-webhook-server/serving-certs/
kubectl get secret smi-controller-webhook-certificate -n shipyard -o json | \
        jq -r '.data."tls.crt"' | \
        base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.crt
kubectl get secret smi-controller-webhook-certificate -n shipyard -o json | \
        jq -r '.data."tls.key"' | \
        base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.key
cd test && go run .
Feature: access.smi-spec.io
  In order to test the TrafficTarget
  As a developer
  I need to ensure the specification is accepted by the server

  Scenario: Apply alpha4 TrafficSplit                                           # features/split.feature:81
    Given the server is running                                                 # main.go:277 -> main.theServerIsRunning
    When I create the following resource                                        # main.go:293 -> main.iCreateTheFollowingResource
      ```
        apiVersion: split.smi-spec.io/v1alpha4
        kind: TrafficSplit
        metadata:
          name: ab-test
        spec:
          service: website
          matches:
          - kind: HTTPRouteGroup
            name: ab-test
            apiGroup: specs.smi-spec.io
          backends:
          - service: website-v1
            weight: 0
          - service: website-v2
            weight: 100
      ```
    And I create the following resource                                         # main.go:293 -> main.iCreateTheFollowingResource
      ```
        apiVersion: specs.smi-spec.io/v1alpha4
        kind: HTTPRouteGroup
        metadata:
          name: ab-test
        spec:
          matches:
          - name: metrics
            pathRegex: "/metrics"
            methods:
            - GET
            headers:
              x-debug: "1"
          - name: health
            pathRegex: "/ping"
            methods: ["*"]
      ```
    Then I expect 1 Istio "VirtualService" named "ab-test" to have been created # istio.go:40 -> main.iExpectIstioNamedToHaveBeenCreated

16 scenarios (16 passed)
35 steps (35 passed)
1m24.7936919s
```

6. Once the tests have completed you can destroy the dev environment using the following command.

```
shipyard destroy
```

Note: After the functional tests have executed, any resources that are created as part of the tests are automatically removed. It should be safe to run the command `make functional_tests` multiple times without
needing to restart the Kubernetes cluster.


## Running locally

To run the controller locally you can follow the steps in the `Running functional tests` section stopping at 
step 5. Instead of the `make functional_test` command you can run the following command.

```
➜ make run_local
mkdir -p /tmp/k8s-webhook-server/serving-certs/
kubectl get secret smi-controller-webhook-certificate -n shipyard -o json | \
        jq -r '.data."tls.crt"' | \
        base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.crt
kubectl get secret smi-controller-webhook-certificate -n shipyard -o json | \
        jq -r '.data."tls.key"' | \
        base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.key
go run .
I0707 13:59:06.256853   27560 request.go:645] Throttling request took 1.0367285s, request: GET:https://127.0.0.1:32687/apis/security.istio.io/v1beta1?timeout=32s
2021-07-07T13:59:06.808+0100    INFO    controller-runtime.metrics      metrics server is starting to listen    {"addr": ":9102"}
2021-07-07T13:59:06.809+0100    INFO    controller-runtime.builder      skip registering a mutating webhook, admission.Defaulter interface is not implemented   {"GVK": "access.smi-spec.io/v1alpha1, Kind=TrafficTarget"}
2021-07-07T13:59:06.809+0100    INFO    controller-runtime.builder      skip registering a validating webhook, admission.Validator interface is not implemented {"GVK": "access.smi-spec.io/v1alpha1, Kind=TrafficTarget"}
```

## Building Docker images
TODO, please see makefile

## Installing with Helm
TODO

## Running the example app
TODO ...

1. First setup a local Kubernetes cluster and install the Istio controller

```shell
shipyard run \
  --var "smi_controller_enabled=true" \
  --var "smi_controller_webhook_enabled=true" \
  --var "smi_controller_namespace=smi" \
  --var "install_example_app=true" ./shipyard
```