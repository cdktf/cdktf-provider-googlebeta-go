// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToRemove struct {
	// The name of the header to remove.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.40.0/docs/resources/google_network_services_edge_cache_service#header_name GoogleNetworkServicesEdgeCacheService#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

