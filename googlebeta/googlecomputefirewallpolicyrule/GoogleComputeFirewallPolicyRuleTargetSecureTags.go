// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputefirewallpolicyrule


type GoogleComputeFirewallPolicyRuleTargetSecureTags struct {
	// Name of the secure tag, created with TagManager's TagValue API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.46.0/docs/resources/google_compute_firewall_policy_rule#name GoogleComputeFirewallPolicyRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

