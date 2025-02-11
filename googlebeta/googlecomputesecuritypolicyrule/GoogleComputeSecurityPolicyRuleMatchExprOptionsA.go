// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputesecuritypolicyrule


type GoogleComputeSecurityPolicyRuleMatchExprOptionsA struct {
	// recaptcha_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.20.0/docs/resources/google_compute_security_policy_rule#recaptcha_options GoogleComputeSecurityPolicyRuleA#recaptcha_options}
	RecaptchaOptions *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsA `field:"required" json:"recaptchaOptions" yaml:"recaptchaOptions"`
}

