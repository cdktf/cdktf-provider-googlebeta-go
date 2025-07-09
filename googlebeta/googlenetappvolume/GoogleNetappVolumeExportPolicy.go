// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetappvolume


type GoogleNetappVolumeExportPolicy struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_netapp_volume#rules GoogleNetappVolume#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

