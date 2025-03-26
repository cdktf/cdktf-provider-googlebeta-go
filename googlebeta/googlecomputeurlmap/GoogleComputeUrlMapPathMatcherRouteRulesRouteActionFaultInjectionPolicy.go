// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.27.0/docs/resources/google_compute_url_map#abort GoogleComputeUrlMap#abort}
	Abort *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.27.0/docs/resources/google_compute_url_map#delay GoogleComputeUrlMap#delay}
	Delay *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

