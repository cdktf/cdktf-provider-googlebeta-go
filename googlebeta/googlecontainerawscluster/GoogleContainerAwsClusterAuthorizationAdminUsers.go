// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainerawscluster


type GoogleContainerAwsClusterAuthorizationAdminUsers struct {
	// The name of the user, e.g. `my-gcp-id@gmail.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.82.0/docs/resources/google_container_aws_cluster#username GoogleContainerAwsCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

