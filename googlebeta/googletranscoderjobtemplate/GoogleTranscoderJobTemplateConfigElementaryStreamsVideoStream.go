// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigElementaryStreamsVideoStream struct {
	// h264 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.8.0/docs/resources/google_transcoder_job_template#h264 GoogleTranscoderJobTemplate#h264}
	H264 *GoogleTranscoderJobTemplateConfigElementaryStreamsVideoStreamH264 `field:"optional" json:"h264" yaml:"h264"`
}

