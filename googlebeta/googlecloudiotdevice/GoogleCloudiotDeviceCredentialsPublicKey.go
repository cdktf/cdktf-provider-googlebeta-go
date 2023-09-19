// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudiotdevice


type GoogleCloudiotDeviceCredentialsPublicKey struct {
	// The format of the key. Possible values: ["RSA_PEM", "RSA_X509_PEM", "ES256_PEM", "ES256_X509_PEM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.83.0/docs/resources/google_cloudiot_device#format GoogleCloudiotDevice#format}
	Format *string `field:"required" json:"format" yaml:"format"`
	// The key data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.83.0/docs/resources/google_cloudiot_device#key GoogleCloudiotDevice#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

