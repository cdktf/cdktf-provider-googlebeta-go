// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterUpgradePolicy struct {
	// Controls whether the upgrade applies to the control plane only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_gkeonprem_vmware_cluster#control_plane_only GoogleGkeonpremVmwareCluster#control_plane_only}
	ControlPlaneOnly interface{} `field:"optional" json:"controlPlaneOnly" yaml:"controlPlaneOnly"`
}

