// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleeventarcpipeline


type GoogleEventarcPipelineDestinationsAuthenticationConfig struct {
	// google_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.48.0/docs/resources/google_eventarc_pipeline#google_oidc GoogleEventarcPipeline#google_oidc}
	GoogleOidc *GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidc `field:"optional" json:"googleOidc" yaml:"googleOidc"`
	// oauth_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.48.0/docs/resources/google_eventarc_pipeline#oauth_token GoogleEventarcPipeline#oauth_token}
	OauthToken *GoogleEventarcPipelineDestinationsAuthenticationConfigOauthToken `field:"optional" json:"oauthToken" yaml:"oauthToken"`
}

