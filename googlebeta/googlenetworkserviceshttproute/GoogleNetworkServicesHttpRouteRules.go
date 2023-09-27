// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworkserviceshttproute


type GoogleNetworkServicesHttpRouteRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.84.0/docs/resources/google_network_services_http_route#action GoogleNetworkServicesHttpRoute#action}
	Action *GoogleNetworkServicesHttpRouteRulesAction `field:"optional" json:"action" yaml:"action"`
	// matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.84.0/docs/resources/google_network_services_http_route#matches GoogleNetworkServicesHttpRoute#matches}
	Matches interface{} `field:"optional" json:"matches" yaml:"matches"`
}

