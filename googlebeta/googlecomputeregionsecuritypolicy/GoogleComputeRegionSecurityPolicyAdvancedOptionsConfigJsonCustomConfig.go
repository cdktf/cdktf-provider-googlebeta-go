// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionsecuritypolicy


type GoogleComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig struct {
	// A list of custom Content-Type header values to apply the JSON parsing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_compute_region_security_policy#content_types GoogleComputeRegionSecurityPolicy#content_types}
	ContentTypes *[]*string `field:"required" json:"contentTypes" yaml:"contentTypes"`
}

