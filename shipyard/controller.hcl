variable "smi_controller_k8s_cluster" {
  default = "dc1"
}

variable "smi_controller_k8s_network" {
  default = "dc1"
}

variable "smi_controller_enabled" {
  default = false
}

variable "smi_controller_webhook_enabled" {
  default = false
}

variable "smi_controller_webhook_port" {
  default = 9443
}

variable "smi_controller_namespace" {
  default = "shipyard"
}

variable "smi_controller_additional_dns" {
  default = "smi-webhook.shipyard.svc"
}

variable "smi_controller_repository" {
  default = "nicholasjackson/istio-smi-controller"
}

variable "smi_controller_tag" {
  default = "0.1.0"
}

variable "install_example_app" {
  default = false
}

module "smi-controller" {
  #source = "/home/nicj/go/src/github.com/shipyard-run/blueprints/modules/kubernetes-smi-controller"
  source = "github.com/shipyard-run/blueprints/modules/kubernetes-smi-controller"
}

k8s_config "app" {
  cluster = "k8s_cluster.dc1"
  paths = [
    "${file_dir()}/../controller-rbac/rbac.yaml",
  ]

  wait_until_ready = true
}

# Create an ingress which exposes the locally running webhook from kubernetes
ingress "smi-webhook" {
  disabled = var.smi_controller_webhook_enabled

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