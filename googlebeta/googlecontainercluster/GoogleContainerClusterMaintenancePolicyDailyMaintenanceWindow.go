// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterMaintenancePolicyDailyMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_container_cluster#start_time GoogleContainerCluster#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

