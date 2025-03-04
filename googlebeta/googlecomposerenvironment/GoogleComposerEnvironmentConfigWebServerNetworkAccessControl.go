// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomposerenvironment


type GoogleComposerEnvironmentConfigWebServerNetworkAccessControl struct {
	// allowed_ip_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.24.0/docs/resources/google_composer_environment#allowed_ip_range GoogleComposerEnvironment#allowed_ip_range}
	AllowedIpRange interface{} `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
}

