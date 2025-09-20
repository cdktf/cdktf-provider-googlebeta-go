// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryVirtualRepositoryConfig struct {
	// upstream_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.50.0/docs/resources/google_artifact_registry_repository#upstream_policies GoogleArtifactRegistryRepository#upstream_policies}
	UpstreamPolicies interface{} `field:"optional" json:"upstreamPolicies" yaml:"upstreamPolicies"`
}

