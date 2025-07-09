// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputewiregroup


type GoogleComputeWireGroupEndpointsInterconnects struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_compute_wire_group#interconnect_name GoogleComputeWireGroup#interconnect_name}.
	InterconnectName *string `field:"required" json:"interconnectName" yaml:"interconnectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_compute_wire_group#interconnect GoogleComputeWireGroup#interconnect}.
	Interconnect *string `field:"optional" json:"interconnect" yaml:"interconnect"`
	// VLAN tags for the interconnect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_compute_wire_group#vlan_tags GoogleComputeWireGroup#vlan_tags}
	VlanTags *[]*float64 `field:"optional" json:"vlanTags" yaml:"vlanTags"`
}

