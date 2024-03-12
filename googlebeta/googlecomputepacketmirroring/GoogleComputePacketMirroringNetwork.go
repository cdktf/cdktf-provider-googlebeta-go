// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputepacketmirroring


type GoogleComputePacketMirroringNetwork struct {
	// The full self_link URL of the network where this rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.20.0/docs/resources/google_compute_packet_mirroring#url GoogleComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

