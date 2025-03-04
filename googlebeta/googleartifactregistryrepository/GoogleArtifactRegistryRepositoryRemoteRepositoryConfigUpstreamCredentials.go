// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials struct {
	// username_password_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.24.0/docs/resources/google_artifact_registry_repository#username_password_credentials GoogleArtifactRegistryRepository#username_password_credentials}
	UsernamePasswordCredentials *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentials `field:"optional" json:"usernamePasswordCredentials" yaml:"usernamePasswordCredentials"`
}

