// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintPolicyRules struct {
	// Setting this to true means that all values are allowed.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_securityposture_posture#allow_all GoogleSecurityposturePosture#allow_all}
	AllowAll interface{} `field:"optional" json:"allowAll" yaml:"allowAll"`
	// Setting this to true means that all values are denied.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_securityposture_posture#deny_all GoogleSecurityposturePosture#deny_all}
	DenyAll interface{} `field:"optional" json:"denyAll" yaml:"denyAll"`
	// If 'true', then the policy is enforced.
	//
	// If 'false', then any configuration is acceptable.
	// This field can be set only in policies for boolean constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_securityposture_posture#enforce GoogleSecurityposturePosture#enforce}
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_securityposture_posture#expr GoogleSecurityposturePosture#expr}
	Expr *GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintPolicyRulesExpr `field:"optional" json:"expr" yaml:"expr"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_securityposture_posture#values GoogleSecurityposturePosture#values}
	Values *GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintPolicyRulesValues `field:"optional" json:"values" yaml:"values"`
}

