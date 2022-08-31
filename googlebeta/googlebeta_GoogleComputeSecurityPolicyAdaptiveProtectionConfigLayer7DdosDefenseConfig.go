// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig struct {
	// If set to true, enables CAAP for L7 DDoS detection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#enable GoogleComputeSecurityPolicy#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Rule visibility. Supported values include: "STANDARD", "PREMIUM".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#rule_visibility GoogleComputeSecurityPolicy#rule_visibility}
	RuleVisibility *string `field:"optional" json:"ruleVisibility" yaml:"ruleVisibility"`
}

