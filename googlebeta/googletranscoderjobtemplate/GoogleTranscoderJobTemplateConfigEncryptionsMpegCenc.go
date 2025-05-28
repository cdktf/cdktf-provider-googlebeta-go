// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigEncryptionsMpegCenc struct {
	// Specify the encryption scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_transcoder_job_template#scheme GoogleTranscoderJobTemplate#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
}

