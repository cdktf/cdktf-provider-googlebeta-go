// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyHttpRulesFrom struct {
	// not_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#not_sources GoogleNetworkSecurityAuthzPolicy#not_sources}
	NotSources interface{} `field:"optional" json:"notSources" yaml:"notSources"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#sources GoogleNetworkSecurityAuthzPolicy#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

