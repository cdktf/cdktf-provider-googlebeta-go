// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleedgecontainervpnconnection


type GoogleEdgecontainerVpnConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.31.0/docs/resources/google_edgecontainer_vpn_connection#create GoogleEdgecontainerVpnConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.31.0/docs/resources/google_edgecontainer_vpn_connection#delete GoogleEdgecontainerVpnConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.31.0/docs/resources/google_edgecontainer_vpn_connection#update GoogleEdgecontainerVpnConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

