// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googletranscoderjob


type GoogleTranscoderJobConfigOutput struct {
	// URI for the output file(s). For example, gs://my-bucket/outputs/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.31.0/docs/resources/google_transcoder_job#uri GoogleTranscoderJob#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

