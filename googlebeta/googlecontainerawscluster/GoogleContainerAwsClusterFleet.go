// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainerawscluster


type GoogleContainerAwsClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.41.0/docs/resources/google_container_aws_cluster#project GoogleContainerAwsCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

