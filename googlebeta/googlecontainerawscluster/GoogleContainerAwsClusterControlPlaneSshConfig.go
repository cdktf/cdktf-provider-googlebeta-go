// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainerawscluster


type GoogleContainerAwsClusterControlPlaneSshConfig struct {
	// The name of the EC2 key pair used to login into cluster machines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.13.0/docs/resources/google_container_aws_cluster#ec2_key_pair GoogleContainerAwsCluster#ec2_key_pair}
	Ec2KeyPair *string `field:"required" json:"ec2KeyPair" yaml:"ec2KeyPair"`
}

