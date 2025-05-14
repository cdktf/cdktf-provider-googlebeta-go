// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleintegrationsauthconfig


type GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPassword struct {
	// Password to be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.35.0/docs/resources/google_integrations_auth_config#password GoogleIntegrationsAuthConfig#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Username to be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.35.0/docs/resources/google_integrations_auth_config#username GoogleIntegrationsAuthConfig#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

