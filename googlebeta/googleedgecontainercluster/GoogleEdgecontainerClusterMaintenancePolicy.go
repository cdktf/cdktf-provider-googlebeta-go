// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleedgecontainercluster


type GoogleEdgecontainerClusterMaintenancePolicy struct {
	// window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.5.0/docs/resources/google_edgecontainer_cluster#window GoogleEdgecontainerCluster#window}
	Window *GoogleEdgecontainerClusterMaintenancePolicyWindow `field:"required" json:"window" yaml:"window"`
}

