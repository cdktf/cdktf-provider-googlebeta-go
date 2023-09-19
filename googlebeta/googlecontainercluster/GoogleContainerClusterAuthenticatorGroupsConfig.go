// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterAuthenticatorGroupsConfig struct {
	// The name of the RBAC security group for use with Google security groups in Kubernetes RBAC.
	//
	// Group name must be in format gke-security-groups@yourdomain.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.83.0/docs/resources/google_container_cluster#security_group GoogleContainerCluster#security_group}
	SecurityGroup *string `field:"required" json:"securityGroup" yaml:"securityGroup"`
}

