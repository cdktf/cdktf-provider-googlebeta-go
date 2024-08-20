// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeinterconnect


type GoogleComputeInterconnectMacsec struct {
	// pre_shared_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.42.0/docs/resources/google_compute_interconnect#pre_shared_keys GoogleComputeInterconnect#pre_shared_keys}
	PreSharedKeys interface{} `field:"required" json:"preSharedKeys" yaml:"preSharedKeys"`
}

