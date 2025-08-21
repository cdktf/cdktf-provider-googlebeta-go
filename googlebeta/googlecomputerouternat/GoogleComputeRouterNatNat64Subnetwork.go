// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputerouternat


type GoogleComputeRouterNatNat64Subnetwork struct {
	// Self-link of the subnetwork resource that will use NAT64.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_compute_router_nat#name GoogleComputeRouterNat#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

