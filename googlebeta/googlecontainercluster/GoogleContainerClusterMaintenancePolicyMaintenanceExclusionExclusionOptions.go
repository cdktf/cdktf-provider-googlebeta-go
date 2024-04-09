// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterMaintenancePolicyMaintenanceExclusionExclusionOptions struct {
	// The scope of automatic upgrades to restrict in the exclusion window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.24.0/docs/resources/google_container_cluster#scope GoogleContainerCluster#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

