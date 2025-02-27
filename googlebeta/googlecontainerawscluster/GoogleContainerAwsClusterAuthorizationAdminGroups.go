// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainerawscluster


type GoogleContainerAwsClusterAuthorizationAdminGroups struct {
	// The name of the group, e.g. `my-group@domain.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.23.0/docs/resources/google_container_aws_cluster#group GoogleContainerAwsCluster#group}
	Group *string `field:"required" json:"group" yaml:"group"`
}

