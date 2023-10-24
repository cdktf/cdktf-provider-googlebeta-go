// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterDatabaseEncryption struct {
	// ENCRYPTED or DECRYPTED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.3.0/docs/resources/google_container_cluster#state GoogleContainerCluster#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// The key to use to encrypt/decrypt secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.3.0/docs/resources/google_container_cluster#key_name GoogleContainerCluster#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
}

