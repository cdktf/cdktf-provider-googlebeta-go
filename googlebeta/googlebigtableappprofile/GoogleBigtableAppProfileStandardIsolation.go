// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlebigtableappprofile


type GoogleBigtableAppProfileStandardIsolation struct {
	// The priority of requests sent using this app profile. Possible values: ["PRIORITY_LOW", "PRIORITY_MEDIUM", "PRIORITY_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_bigtable_app_profile#priority GoogleBigtableAppProfile#priority}
	Priority *string `field:"required" json:"priority" yaml:"priority"`
}

