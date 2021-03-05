// Install Istio using the Helm chart.

helm "istio" {
  cluster = "k8s_cluster.${var.istio_k8s_cluster}"

  // chart = "github.com/hashicorp/consul-helm?ref=crd-controller-base"
  chart = "./helm/istio/istio-operator"
  values = "./helm/istio-values.yaml"

  health_check {
    timeout = "60s"
    pods = ["name=istio-operator"]
  }
}

k8s_config "istio" {
  depends_on = ["helm.istio"]

  cluster = "k8s_cluster.dc1"
  paths = [
    "./istio_config.yaml",
  ]
  
  health_check {
    timeout = "120s"
    pods = ["istio.io/rev=default"]
  }
  wait_until_ready = true
}

//ingress "consul" {
//  source {
//    driver = "local"
//    
//    config {
//      port = 8500
//    }
//  }
//  
//  destination {
//    driver = "k8s"
//    
//    config {
//      cluster = "k8s_cluster.${var.istio_k8s_cluster}"
//      address = "consul-server.default.svc"
//      port = 8500
//    }
//  }
//}

