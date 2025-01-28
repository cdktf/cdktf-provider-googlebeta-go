// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterNetworkConfigDhcpIpConfig struct {
	// enabled is a flag to mark if DHCP IP allocation is used for VMware user clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.18.0/docs/resources/google_gkeonprem_vmware_cluster#enabled GoogleGkeonpremVmwareCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

