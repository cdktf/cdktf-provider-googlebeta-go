// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainerazurecluster


type GoogleContainerAzureClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.3.0/docs/resources/google_container_azure_cluster#admin_users GoogleContainerAzureCluster#admin_users}
	AdminUsers interface{} `field:"required" json:"adminUsers" yaml:"adminUsers"`
}

