// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlelookerinstance


type GoogleLookerInstanceCustomDomain struct {
	// Domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.22.0/docs/resources/google_looker_instance#domain GoogleLookerInstance#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

