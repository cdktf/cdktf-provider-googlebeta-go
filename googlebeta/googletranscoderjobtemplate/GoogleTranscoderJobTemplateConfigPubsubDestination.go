// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigPubsubDestination struct {
	// The name of the Pub/Sub topic to publish job completion notification to. For example: projects/{project}/topics/{topic}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.11.0/docs/resources/google_transcoder_job_template#topic GoogleTranscoderJobTemplate#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

