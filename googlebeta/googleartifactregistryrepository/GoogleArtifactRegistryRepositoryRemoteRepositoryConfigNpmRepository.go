// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository struct {
	// Address of the remote repository. Default value: "NPMJS" Possible values: ["NPMJS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.22.0/docs/resources/google_artifact_registry_repository#public_repository GoogleArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

