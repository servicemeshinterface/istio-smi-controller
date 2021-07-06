variable "istio_k8s_cluster" {
  default = "dc1"
}

variable "istio_k8s_network" {
  default = "dc1"
}

module "istio" {
  source = "github.com/shipyard-run/blueprints/modules//kubernetes-istio"
}