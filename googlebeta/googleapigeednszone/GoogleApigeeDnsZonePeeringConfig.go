// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapigeednszone


type GoogleApigeeDnsZonePeeringConfig struct {
	// The name of the producer VPC network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.2/docs/resources/google_apigee_dns_zone#target_network_id GoogleApigeeDnsZone#target_network_id}
	TargetNetworkId *string `field:"required" json:"targetNetworkId" yaml:"targetNetworkId"`
	// The ID of the project that contains the producer VPC network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.2/docs/resources/google_apigee_dns_zone#target_project_id GoogleApigeeDnsZone#target_project_id}
	TargetProjectId *string `field:"required" json:"targetProjectId" yaml:"targetProjectId"`
}

