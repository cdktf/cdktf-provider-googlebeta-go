package googlecomputerouternat


type GoogleComputeRouterNatRulesAction struct {
	// A list of URLs of the IP resources used for this NAT rule.
	//
	// These IP addresses must be valid static external IP addresses assigned to the project.
	// This field is used for public NAT.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_nat#source_nat_active_ips GoogleComputeRouterNat#source_nat_active_ips}
	SourceNatActiveIps *[]*string `field:"optional" json:"sourceNatActiveIps" yaml:"sourceNatActiveIps"`
	// A list of URLs of the IP resources to be drained.
	//
	// These IPs must be valid static external IPs that have been assigned to the NAT.
	// These IPs should be used for updating/patching a NAT rule only.
	// This field is used for public NAT.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_nat#source_nat_drain_ips GoogleComputeRouterNat#source_nat_drain_ips}
	SourceNatDrainIps *[]*string `field:"optional" json:"sourceNatDrainIps" yaml:"sourceNatDrainIps"`
}
