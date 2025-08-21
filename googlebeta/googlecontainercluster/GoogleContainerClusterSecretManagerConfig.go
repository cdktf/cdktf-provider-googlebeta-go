// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterSecretManagerConfig struct {
	// Enable the Secret manager csi component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// rotation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_container_cluster#rotation_config GoogleContainerCluster#rotation_config}
	RotationConfig *GoogleContainerClusterSecretManagerConfigRotationConfig `field:"optional" json:"rotationConfig" yaml:"rotationConfig"`
}

