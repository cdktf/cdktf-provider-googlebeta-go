// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlednspolicy


type GoogleDnsPolicyDns64Config struct {
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_dns_policy#scope GoogleDnsPolicy#scope}
	Scope *GoogleDnsPolicyDns64ConfigScope `field:"required" json:"scope" yaml:"scope"`
}

