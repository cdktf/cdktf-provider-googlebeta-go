// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlekmskeyhandle


type GoogleKmsKeyHandleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.1/docs/resources/google_kms_key_handle#create GoogleKmsKeyHandle#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.1/docs/resources/google_kms_key_handle#delete GoogleKmsKeyHandle#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

