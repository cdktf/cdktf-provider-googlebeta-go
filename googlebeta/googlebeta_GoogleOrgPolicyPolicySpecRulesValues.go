// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleOrgPolicyPolicySpecRulesValues struct {
	// List of values allowed at this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_org_policy_policy#allowed_values GoogleOrgPolicyPolicy#allowed_values}
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// List of values denied at this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_org_policy_policy#denied_values GoogleOrgPolicyPolicy#denied_values}
	DeniedValues *[]*string `field:"optional" json:"deniedValues" yaml:"deniedValues"`
}

