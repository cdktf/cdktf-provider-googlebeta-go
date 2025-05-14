// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlebeyondcorpappconnector


type GoogleBeyondcorpAppConnectorPrincipalInfoServiceAccount struct {
	// Email address of the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.35.0/docs/resources/google_beyondcorp_app_connector#email GoogleBeyondcorpAppConnector#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

