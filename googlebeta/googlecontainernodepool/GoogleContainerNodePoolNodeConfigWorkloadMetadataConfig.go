// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigWorkloadMetadataConfig struct {
	// Mode is the configuration for how to expose metadata to workloads running on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.13.0/docs/resources/google_container_node_pool#mode GoogleContainerNodePool#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

