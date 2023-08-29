// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfig struct {
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.80.0/docs/resources/google_container_node_pool#sysctls GoogleContainerNodePool#sysctls}
	Sysctls *map[string]*string `field:"required" json:"sysctls" yaml:"sysctls"`
}

