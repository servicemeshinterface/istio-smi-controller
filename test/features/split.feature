Feature: split.smi-spec.io
  In order to test the TrafficSplit
  As a developer
  I need to ensure the specification is accepted by the server

  @split
  Scenario: Apply alpha1 TrafficSplit
    Given the server is running
    When I create the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha1
      kind: TrafficSplit
      metadata:
        name: trafficsplit-sample
      spec:
        service: foo
        backends:
          - service: bar
            weight: 50m
          - service: baz
            weight: 50m
    ```
    Then I expect 1 Istio "VirtualService" named "trafficsplit-sample" to have been created
  
  @split
  Scenario: Apply alpha2 TrafficSplit
    Given the server is running
    When I create the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha2
      kind: TrafficSplit
      metadata:
        name: trafficsplit-sample
      spec:
        service: foo
        backends:
          - service: bar
            weight: 50
          - service: baz
            weight: 50
    ```
    Then I expect 1 Istio "VirtualService" named "trafficsplit-sample" to have been created
  
  @split @alpha3
  Scenario: Apply alpha3 TrafficSplit
    Given the server is running
    When I create the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha3
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
    And I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha3
      kind: HTTPRouteGroup
      metadata:
        name: ab-test
      spec:
        matches:
        - name: metrics
          pathRegex: "/metrics"
          methods:
          - GET
        - name: health
          pathRegex: "/ping"
          methods: ["*"]
    ```
    Then I expect 1 Istio "VirtualService" named "ab-test" to have been created
  
  @split/alpha4
  Scenario: Apply alpha4 TrafficSplit
    Given the server is running
    When I create the following resource
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
    And I create the following resource
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
    Then I expect 1 Istio "VirtualService" named "ab-test" to have been created

  @split
  Scenario: Delete alpha1 TrafficSplit
    Given the server is running
    And the following resource exists
    ```
      apiVersion: split.smi-spec.io/v1alpha1
      kind: TrafficSplit
      metadata:
        name: trafficsplit-sample
      spec:
        service: foo
        backends:
          - service: bar
            weight: 50m
          - service: baz
            weight: 50m
    ```
    And I verify that 1 Istio "VirtualService" named "trafficsplit-sample" exists
    When I delete the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha1
      kind: TrafficSplit
      metadata:
        name: trafficsplit-sample
      spec:
        service: foo
        backends:
          - service: bar
            weight: 50m
          - service: baz
            weight: 50m
    ```
    Then I expect no Istio "VirtualService" named "trafficsplit-sample" to exist

  @split
  Scenario: Delete alpha2 TrafficSplit
    Given the server is running
    And the following resource exists
    ```
      apiVersion: split.smi-spec.io/v1alpha2
      kind: TrafficSplit
      metadata:
        name: trafficsplit-sample
      spec:
        service: foo
        backends:
          - service: bar
            weight: 50
          - service: baz
            weight: 50
    ```
    And I verify that 1 Istio "VirtualService" named "trafficsplit-sample" exists
    When I delete the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha2
      kind: TrafficSplit
      metadata:
        name: trafficsplit-sample
      spec:
        service: foo
        backends:
          - service: bar
            weight: 50
          - service: baz
            weight: 50
    ```
    Then I expect no Istio "VirtualService" named "trafficsplit-sample" to exist
  
  @split @alpha3
  Scenario: Delete alpha3 TrafficSplit
    Given the server is running
    And the following resource exists
    ```
      apiVersion: split.smi-spec.io/v1alpha3
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
    And I verify that 1 Istio "VirtualService" named "ab-test" exists
    When I delete the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha3
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
    Then I expect no Istio "VirtualService" named "ab-test" to exist
  
  @split @alpha4
  Scenario: Delete alpha4 TrafficSplit
    Given the server is running
    And the following resource exists
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
    And I verify that 1 Istio "VirtualService" named "ab-test" exists
    When I delete the following resource
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
    Then I expect no Istio "VirtualService" named "ab-test" to exist