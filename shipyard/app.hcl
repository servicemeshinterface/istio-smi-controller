//k8s_config "app-config" {
//  depends_on = ["module.istio"]
//
//  cluster = "k8s_cluster.dc1"
//  paths = [
//    "./app/consul_config.yaml",
//  ]
//
//  wait_until_ready = true
//}

k8s_config "app-pods" {
  cluster = "k8s_cluster.${var.istio_k8s_cluster}"
  paths = [
    "./app/load_test.yaml",
    "./app/web.yaml",
    "./app/apiV1.yaml",
  ]

  wait_until_ready = true
}

ingress "web" {
  source {
    driver = "local"
    
    config {
      port = 18080
    }
  }
  
  destination {
    driver = "k8s"
    
    config {
      cluster = "k8s_cluster.${var.istio_k8s_cluster}"
      address = "web-service.default.svc"
      port = 9090
    }
  }
}
