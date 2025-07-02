// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlefirebaseapphostingtraffic


type GoogleFirebaseAppHostingTrafficTarget struct {
	// splits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.42.0/docs/resources/google_firebase_app_hosting_traffic#splits GoogleFirebaseAppHostingTraffic#splits}
	Splits interface{} `field:"required" json:"splits" yaml:"splits"`
}

