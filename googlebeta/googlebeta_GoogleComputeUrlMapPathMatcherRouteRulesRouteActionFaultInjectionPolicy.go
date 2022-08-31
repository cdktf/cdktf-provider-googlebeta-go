// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_url_map#abort GoogleComputeUrlMap#abort}
	Abort *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_url_map#delay GoogleComputeUrlMap#delay}
	Delay *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

