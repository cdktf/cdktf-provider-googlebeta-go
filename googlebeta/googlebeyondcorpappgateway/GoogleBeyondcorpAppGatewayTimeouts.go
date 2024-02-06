// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlebeyondcorpappgateway


type GoogleBeyondcorpAppGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.15.0/docs/resources/google_beyondcorp_app_gateway#create GoogleBeyondcorpAppGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.15.0/docs/resources/google_beyondcorp_app_gateway#delete GoogleBeyondcorpAppGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.15.0/docs/resources/google_beyondcorp_app_gateway#update GoogleBeyondcorpAppGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

