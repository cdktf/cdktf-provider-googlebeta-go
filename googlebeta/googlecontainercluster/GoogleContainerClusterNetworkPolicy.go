// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNetworkPolicy struct {
	// Whether network policy is enabled on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.9.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The selected network policy provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.9.0/docs/resources/google_container_cluster#provider GoogleContainerCluster#provider}
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
}

