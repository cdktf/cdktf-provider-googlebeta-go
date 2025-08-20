// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigBootDisk struct {
	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_container_cluster#disk_type GoogleContainerCluster#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Configured IOPs provisioning. Only valid with disk type hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_container_cluster#provisioned_iops GoogleContainerCluster#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Configured throughput provisioning. Only valid with disk type hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_container_cluster#provisioned_throughput GoogleContainerCluster#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_container_cluster#size_gb GoogleContainerCluster#size_gb}
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

