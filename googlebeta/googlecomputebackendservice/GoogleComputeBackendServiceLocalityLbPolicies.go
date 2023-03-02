package googlecomputebackendservice


type GoogleComputeBackendServiceLocalityLbPolicies struct {
	// custom_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_service#custom_policy GoogleComputeBackendService#custom_policy}
	CustomPolicy *GoogleComputeBackendServiceLocalityLbPoliciesCustomPolicy `field:"optional" json:"customPolicy" yaml:"customPolicy"`
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_service#policy GoogleComputeBackendService#policy}
	Policy *GoogleComputeBackendServiceLocalityLbPoliciesPolicy `field:"optional" json:"policy" yaml:"policy"`
}

