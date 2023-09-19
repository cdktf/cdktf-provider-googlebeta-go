// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig struct {
	// Whether to enable multiple network interfaces for your pods. When set network_config.advanced_networking is automatically set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.83.0/docs/resources/google_gkeonprem_bare_metal_cluster#enabled GoogleGkeonpremBareMetalCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

