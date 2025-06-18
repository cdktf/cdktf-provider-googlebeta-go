// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainerazurenodepool


type GoogleContainerAzureNodePoolConfigSshConfig struct {
	// The SSH public key data for VMs managed by Anthos.
	//
	// This accepts the authorized_keys file format used in OpenSSH according to the sshd(8) manual page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.40.0/docs/resources/google_container_azure_node_pool#authorized_key GoogleContainerAzureNodePool#authorized_key}
	AuthorizedKey *string `field:"required" json:"authorizedKey" yaml:"authorizedKey"`
}

