// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputebackendservice


type GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicy struct {
	// zonal_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.48.0/docs/resources/google_compute_backend_service#zonal_affinity GoogleComputeBackendService#zonal_affinity}
	ZonalAffinity *GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity `field:"optional" json:"zonalAffinity" yaml:"zonalAffinity"`
}

