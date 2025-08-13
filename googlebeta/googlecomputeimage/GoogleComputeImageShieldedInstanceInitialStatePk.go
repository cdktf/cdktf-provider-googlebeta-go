// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeimage


type GoogleComputeImageShieldedInstanceInitialStatePk struct {
	// The raw content in the secure keys file.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.48.0/docs/resources/google_compute_image#content GoogleComputeImage#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The file type of source file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.48.0/docs/resources/google_compute_image#file_type GoogleComputeImage#file_type}
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

