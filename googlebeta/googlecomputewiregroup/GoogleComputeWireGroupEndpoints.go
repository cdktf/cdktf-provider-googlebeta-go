// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputewiregroup


type GoogleComputeWireGroupEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_compute_wire_group#endpoint GoogleComputeWireGroup#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// interconnects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_compute_wire_group#interconnects GoogleComputeWireGroup#interconnects}
	Interconnects interface{} `field:"optional" json:"interconnects" yaml:"interconnects"`
}

