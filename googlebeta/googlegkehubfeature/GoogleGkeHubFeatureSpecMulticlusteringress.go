// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkehubfeature


type GoogleGkeHubFeatureSpecMulticlusteringress struct {
	// Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: 'projects/foo-proj/locations/global/memberships/bar'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.39.1/docs/resources/google_gke_hub_feature#config_membership GoogleGkeHubFeature#config_membership}
	ConfigMembership *string `field:"required" json:"configMembership" yaml:"configMembership"`
}

