// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterMasterAuthClientCertificateConfig struct {
	// Whether client certificate authorization is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.39.1/docs/resources/google_container_cluster#issue_client_certificate GoogleContainerCluster#issue_client_certificate}
	IssueClientCertificate interface{} `field:"required" json:"issueClientCertificate" yaml:"issueClientCertificate"`
}

