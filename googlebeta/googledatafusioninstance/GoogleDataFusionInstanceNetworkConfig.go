// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatafusioninstance


type GoogleDataFusionInstanceNetworkConfig struct {
	// The IP range in CIDR notation to use for the managed Data Fusion instance nodes.
	//
	// This range must not overlap with any other ranges used in the Data Fusion instance network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.22.0/docs/resources/google_data_fusion_instance#ip_allocation GoogleDataFusionInstance#ip_allocation}
	IpAllocation *string `field:"required" json:"ipAllocation" yaml:"ipAllocation"`
	// Name of the network in the project with which the tenant project will be peered for executing pipelines.
	//
	// In case of shared VPC where the network resides in another host
	// project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.22.0/docs/resources/google_data_fusion_instance#network GoogleDataFusionInstance#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

