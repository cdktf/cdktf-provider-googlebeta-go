package googlecomputenetworkfirewallpolicyrule


type GoogleComputeNetworkFirewallPolicyRuleMatchSrcSecureTags struct {
	// Name of the secure tag, created with TagManager's TagValue API. @pattern tagValues/[0-9]+.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_network_firewall_policy_rule#name GoogleComputeNetworkFirewallPolicyRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}
