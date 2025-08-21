// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputevpntunnel


type GoogleComputeVpnTunnelCipherSuitePhase2 struct {
	// Encryption algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_compute_vpn_tunnel#encryption GoogleComputeVpnTunnel#encryption}
	Encryption *[]*string `field:"optional" json:"encryption" yaml:"encryption"`
	// Integrity algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_compute_vpn_tunnel#integrity GoogleComputeVpnTunnel#integrity}
	Integrity *[]*string `field:"optional" json:"integrity" yaml:"integrity"`
	// Perfect forward secrecy groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_compute_vpn_tunnel#pfs GoogleComputeVpnTunnel#pfs}
	Pfs *[]*string `field:"optional" json:"pfs" yaml:"pfs"`
}

