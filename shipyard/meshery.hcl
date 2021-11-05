// Optional meshery install
// helm "meshery" {
//   cluster = "k8s_cluster.dc1"

//   chart = "github.com/layer5io/meshery/install/kubernetes/helm//meshery"

//   values_string = {
//     "mesheryistio.enabled" = "true"
//     "mesheryistio.fullnameOverride" = "meshery-istio"
//     "mesherylinkerd.enabled" = "true"
//     "mesherylinkerd.fullnameOverride" = "meshery-linkerd"
//     "mesheryconsul.enabled" = "true"
//     "mesheryconsul.fullnameOverride" = "meshery-consul"
//   }

//   create_namespace = true
//   namespace = "meshery"
// }


// ingress "meshery" {

//   destination {
//     driver = "k8s"

//     config { 
//       cluster = "k8s_cluster.dc1"
//       address = "meshery.meshery.svc"
//       port = 9081
//     }
//   }
  
//   source {
//     driver = "local"

//     config {
//       port = 9081
//    }
//   }
// }

// ingress "public" {

//   destination {
//     driver = "k8s"

//     config { 
//       cluster = "k8s_cluster.dc1"
//       address = "public.default.svc"
//       port = 9090
//     }
//   }
  
//   source {
//     driver = "local"

//     config {
//       port = 19090
//    }

//   }
// }
