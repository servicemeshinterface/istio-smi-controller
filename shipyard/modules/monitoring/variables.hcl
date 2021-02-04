variable "monitoring_helm_values_grafana" {
   default = "./helm/grafana-values.yaml" 
}

variable "monitoring_helm_values_loki" {
   default = "./helm/loki-values.yaml" 
}

variable "monitoring_helm_values_promtail" {
   default = "./helm/promtail-values.yaml" 
}

variable "monitoring_helm_values_prometheus" {
   default = "./helm/prometheus-values.yaml" 
}
