module "smi-controller" {
  source = "./modules/smi-controller"
}

module "monitoring" {
  source = "github.com/nicholasjackson/hashicorp-shipyard-modules/modules/monitoring"
}

module "istio" {
  source = "./modules/istio"
  depends_on = ["module.monitoring"]
}
