// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceSubsetting struct {
	// The algorithm used for subsetting. Possible values: ["CONSISTENT_HASH_SUBSETTING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.43.1/docs/resources/google_compute_region_backend_service#policy GoogleComputeRegionBackendService#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
}

