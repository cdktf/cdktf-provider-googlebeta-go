package googlecomputenetworkfirewallpolicyrule


type GoogleComputeNetworkFirewallPolicyRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_network_firewall_policy_rule#create GoogleComputeNetworkFirewallPolicyRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_network_firewall_policy_rule#delete GoogleComputeNetworkFirewallPolicyRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_network_firewall_policy_rule#update GoogleComputeNetworkFirewallPolicyRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
