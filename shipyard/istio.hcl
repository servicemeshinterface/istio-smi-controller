module "smi-controller" {
  source = "./modules/smi-controller"
}

module "monitoring" {
  source = "./modules/monitoring"
}

module "istio" {
  source = "./modules/istio"
  depends_on = ["module.monitoring"]
}
