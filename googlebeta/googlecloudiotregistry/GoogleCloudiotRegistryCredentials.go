// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudiotregistry


type GoogleCloudiotRegistryCredentials struct {
	// A public key certificate format and data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.83.0/docs/resources/google_cloudiot_registry#public_key_certificate GoogleCloudiotRegistry#public_key_certificate}
	PublicKeyCertificate *map[string]*string `field:"required" json:"publicKeyCertificate" yaml:"publicKeyCertificate"`
}

