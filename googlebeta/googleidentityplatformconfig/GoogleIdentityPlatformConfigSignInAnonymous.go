// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleidentityplatformconfig


type GoogleIdentityPlatformConfigSignInAnonymous struct {
	// Whether anonymous user auth is enabled for the project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.2/docs/resources/google_identity_platform_config#enabled GoogleIdentityPlatformConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

