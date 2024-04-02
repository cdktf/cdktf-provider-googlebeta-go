// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworkserviceshttproute


type GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicy struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.23.0/docs/resources/google_network_services_http_route#destination GoogleNetworkServicesHttpRoute#destination}
	Destination *GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyDestination `field:"optional" json:"destination" yaml:"destination"`
}

