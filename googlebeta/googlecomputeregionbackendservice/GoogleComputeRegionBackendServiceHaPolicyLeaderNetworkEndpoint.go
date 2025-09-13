// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceHaPolicyLeaderNetworkEndpoint struct {
	// The name of the VM instance of the leader network endpoint.
	//
	// The instance must
	// already be attached to the NEG specified in the haPolicy.leader.backendGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_compute_region_backend_service#instance GoogleComputeRegionBackendService#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
}

