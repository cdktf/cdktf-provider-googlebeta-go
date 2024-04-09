// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2Config struct {
	// control_plane_ip_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.24.0/docs/resources/google_gkeonprem_vmware_cluster#control_plane_ip_block GoogleGkeonpremVmwareCluster#control_plane_ip_block}
	ControlPlaneIpBlock *GoogleGkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock `field:"optional" json:"controlPlaneIpBlock" yaml:"controlPlaneIpBlock"`
}

