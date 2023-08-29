// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlebigtablegcpolicy


type GoogleBigtableGcPolicyMaxVersion struct {
	// Number of version before applying the GC policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.80.0/docs/resources/google_bigtable_gc_policy#number GoogleBigtableGcPolicy#number}
	Number *float64 `field:"required" json:"number" yaml:"number"`
}

