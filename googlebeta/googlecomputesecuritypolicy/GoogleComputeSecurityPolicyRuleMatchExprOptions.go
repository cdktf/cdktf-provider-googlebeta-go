// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyRuleMatchExprOptions struct {
	// recaptcha_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.14.1/docs/resources/google_compute_security_policy#recaptcha_options GoogleComputeSecurityPolicy#recaptcha_options}
	RecaptchaOptions *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptions `field:"required" json:"recaptchaOptions" yaml:"recaptchaOptions"`
}

