// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterMeshCertificates struct {
	// When enabled the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.25.0/docs/resources/google_container_cluster#enable_certificates GoogleContainerCluster#enable_certificates}
	EnableCertificates interface{} `field:"required" json:"enableCertificates" yaml:"enableCertificates"`
}

