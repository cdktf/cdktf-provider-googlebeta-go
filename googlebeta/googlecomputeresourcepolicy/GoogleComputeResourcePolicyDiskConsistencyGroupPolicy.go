// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeresourcepolicy


type GoogleComputeResourcePolicyDiskConsistencyGroupPolicy struct {
	// Enable disk consistency on the resource policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.80.0/docs/resources/google_compute_resource_policy#enabled GoogleComputeResourcePolicy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

