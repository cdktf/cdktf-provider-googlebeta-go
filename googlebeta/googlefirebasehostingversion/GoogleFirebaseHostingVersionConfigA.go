// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlefirebasehostingversion


type GoogleFirebaseHostingVersionConfigA struct {
	// redirects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.10.0/docs/resources/google_firebase_hosting_version#redirects GoogleFirebaseHostingVersion#redirects}
	Redirects interface{} `field:"optional" json:"redirects" yaml:"redirects"`
	// rewrites block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.10.0/docs/resources/google_firebase_hosting_version#rewrites GoogleFirebaseHostingVersion#rewrites}
	Rewrites interface{} `field:"optional" json:"rewrites" yaml:"rewrites"`
}

