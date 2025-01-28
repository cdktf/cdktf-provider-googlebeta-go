// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildAvailableSecrets struct {
	// secret_manager block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.18.0/docs/resources/google_cloudbuild_trigger#secret_manager GoogleCloudbuildTrigger#secret_manager}
	SecretManager interface{} `field:"required" json:"secretManager" yaml:"secretManager"`
}

