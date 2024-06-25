// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlealloydbinstance


type GoogleAlloydbInstanceNetworkConfigAuthorizedExternalNetworks struct {
	// CIDR range for one authorized network of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.35.0/docs/resources/google_alloydb_instance#cidr_range GoogleAlloydbInstance#cidr_range}
	CidrRange *string `field:"optional" json:"cidrRange" yaml:"cidrRange"`
}

