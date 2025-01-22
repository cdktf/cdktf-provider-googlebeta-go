// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleiapsettings


type GoogleIapSettingsAccessSettingsAllowedDomainsSettings struct {
	// List of trusted domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.17.0/docs/resources/google_iap_settings#domains GoogleIapSettings#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Configuration for customers to opt in for the feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.17.0/docs/resources/google_iap_settings#enable GoogleIapSettings#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

