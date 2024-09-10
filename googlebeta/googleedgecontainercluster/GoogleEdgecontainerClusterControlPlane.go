// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleedgecontainercluster


type GoogleEdgecontainerClusterControlPlane struct {
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.2.0/docs/resources/google_edgecontainer_cluster#local GoogleEdgecontainerCluster#local}
	Local *GoogleEdgecontainerClusterControlPlaneLocal `field:"optional" json:"local" yaml:"local"`
	// remote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.2.0/docs/resources/google_edgecontainer_cluster#remote GoogleEdgecontainerCluster#remote}
	Remote *GoogleEdgecontainerClusterControlPlaneRemote `field:"optional" json:"remote" yaml:"remote"`
}

