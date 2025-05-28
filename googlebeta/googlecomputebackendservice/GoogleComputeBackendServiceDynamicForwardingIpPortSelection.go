// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputebackendservice


type GoogleComputeBackendServiceDynamicForwardingIpPortSelection struct {
	// A boolean flag enabling IP:PORT based dynamic forwarding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_backend_service#enabled GoogleComputeBackendService#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

