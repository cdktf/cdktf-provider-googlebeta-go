// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeresourcepolicy


type GoogleComputeResourcePolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.36.0/docs/resources/google_compute_resource_policy#create GoogleComputeResourcePolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.36.0/docs/resources/google_compute_resource_policy#delete GoogleComputeResourcePolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

