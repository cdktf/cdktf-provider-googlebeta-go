// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlednsmanagedzone


type GoogleDnsManagedZoneServiceDirectoryConfig struct {
	// namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.21.0/docs/resources/google_dns_managed_zone#namespace GoogleDnsManagedZone#namespace}
	Namespace *GoogleDnsManagedZoneServiceDirectoryConfigNamespace `field:"required" json:"namespace" yaml:"namespace"`
}

