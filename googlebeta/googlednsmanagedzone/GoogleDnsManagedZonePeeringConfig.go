// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlednsmanagedzone


type GoogleDnsManagedZonePeeringConfig struct {
	// target_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.17.0/docs/resources/google_dns_managed_zone#target_network GoogleDnsManagedZone#target_network}
	TargetNetwork *GoogleDnsManagedZonePeeringConfigTargetNetwork `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
}

