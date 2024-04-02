// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesendpointpolicy


type GoogleNetworkServicesEndpointPolicyEndpointMatcher struct {
	// metadata_label_matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.23.0/docs/resources/google_network_services_endpoint_policy#metadata_label_matcher GoogleNetworkServicesEndpointPolicy#metadata_label_matcher}
	MetadataLabelMatcher *GoogleNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcher `field:"required" json:"metadataLabelMatcher" yaml:"metadataLabelMatcher"`
}

