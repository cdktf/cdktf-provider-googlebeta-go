// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlesecretmanagersecret


type GoogleSecretManagerSecretReplicationUserManaged struct {
	// replicas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.17.0/docs/resources/google_secret_manager_secret#replicas GoogleSecretManagerSecret#replicas}
	Replicas interface{} `field:"required" json:"replicas" yaml:"replicas"`
}

