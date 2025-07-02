// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlefirebaseapphostingdomain


type GoogleFirebaseAppHostingDomainServe struct {
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.42.0/docs/resources/google_firebase_app_hosting_domain#redirect GoogleFirebaseAppHostingDomain#redirect}
	Redirect *GoogleFirebaseAppHostingDomainServeRedirect `field:"optional" json:"redirect" yaml:"redirect"`
}

