// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleedgecontainercluster


type GoogleEdgecontainerClusterSystemAddonsConfig struct {
	// ingress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.41.0/docs/resources/google_edgecontainer_cluster#ingress GoogleEdgecontainerCluster#ingress}
	Ingress *GoogleEdgecontainerClusterSystemAddonsConfigIngress `field:"optional" json:"ingress" yaml:"ingress"`
}

