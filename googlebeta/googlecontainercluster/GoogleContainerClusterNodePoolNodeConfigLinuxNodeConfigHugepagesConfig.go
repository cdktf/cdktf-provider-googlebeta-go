// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfig struct {
	// Amount of 1G hugepages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.24.0/docs/resources/google_container_cluster#hugepage_size_1g GoogleContainerCluster#hugepage_size_1g}
	HugepageSize1G *float64 `field:"optional" json:"hugepageSize1G" yaml:"hugepageSize1G"`
	// Amount of 2M hugepages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.24.0/docs/resources/google_container_cluster#hugepage_size_2m GoogleContainerCluster#hugepage_size_2m}
	HugepageSize2M *float64 `field:"optional" json:"hugepageSize2M" yaml:"hugepageSize2M"`
}

