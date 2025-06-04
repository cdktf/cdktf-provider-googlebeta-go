// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceDynamicForwarding struct {
	// ip_port_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.38.0/docs/resources/google_compute_region_backend_service#ip_port_selection GoogleComputeRegionBackendService#ip_port_selection}
	IpPortSelection *GoogleComputeRegionBackendServiceDynamicForwardingIpPortSelection `field:"optional" json:"ipPortSelection" yaml:"ipPortSelection"`
}

