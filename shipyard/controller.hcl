ingress "smi-webhook" {
  source {
    driver = "k8s"
    
    config {
      cluster = "k8s_cluster.dc1"
      port = 9443
    }
  }
  
  destination {
    driver = "local"
    
    config { 
      address = "localhost"
      port = 9443
    }
  
  }
}
