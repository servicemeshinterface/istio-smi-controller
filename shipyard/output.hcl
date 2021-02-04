output "KUBECONFIG" {
  value = k8s_config("${var.istio_k8s_cluster}")
}
