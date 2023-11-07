// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkehubfeature


type GoogleGkeHubFeatureSpec struct {
	// fleetobservability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.5.0/docs/resources/google_gke_hub_feature#fleetobservability GoogleGkeHubFeature#fleetobservability}
	Fleetobservability *GoogleGkeHubFeatureSpecFleetobservability `field:"optional" json:"fleetobservability" yaml:"fleetobservability"`
	// multiclusteringress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.5.0/docs/resources/google_gke_hub_feature#multiclusteringress GoogleGkeHubFeature#multiclusteringress}
	Multiclusteringress *GoogleGkeHubFeatureSpecMulticlusteringress `field:"optional" json:"multiclusteringress" yaml:"multiclusteringress"`
}

