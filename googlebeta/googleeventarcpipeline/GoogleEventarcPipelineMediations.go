// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleeventarcpipeline


type GoogleEventarcPipelineMediations struct {
	// transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.41.0/docs/resources/google_eventarc_pipeline#transformation GoogleEventarcPipeline#transformation}
	Transformation *GoogleEventarcPipelineMediationsTransformation `field:"optional" json:"transformation" yaml:"transformation"`
}

