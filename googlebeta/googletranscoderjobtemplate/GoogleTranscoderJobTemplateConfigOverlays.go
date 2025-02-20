// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigOverlays struct {
	// animations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.21.0/docs/resources/google_transcoder_job_template#animations GoogleTranscoderJobTemplate#animations}
	Animations interface{} `field:"optional" json:"animations" yaml:"animations"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.21.0/docs/resources/google_transcoder_job_template#image GoogleTranscoderJobTemplate#image}
	Image *GoogleTranscoderJobTemplateConfigOverlaysImage `field:"optional" json:"image" yaml:"image"`
}

