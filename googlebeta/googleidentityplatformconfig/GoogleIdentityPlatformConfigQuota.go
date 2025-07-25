// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleidentityplatformconfig


type GoogleIdentityPlatformConfigQuota struct {
	// sign_up_quota_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_identity_platform_config#sign_up_quota_config GoogleIdentityPlatformConfig#sign_up_quota_config}
	SignUpQuotaConfig *GoogleIdentityPlatformConfigQuotaSignUpQuotaConfig `field:"optional" json:"signUpQuotaConfig" yaml:"signUpQuotaConfig"`
}

