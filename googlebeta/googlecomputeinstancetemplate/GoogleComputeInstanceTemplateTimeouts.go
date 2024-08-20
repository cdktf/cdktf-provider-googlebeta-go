// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.42.0/docs/resources/google_compute_instance_template#create GoogleComputeInstanceTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.42.0/docs/resources/google_compute_instance_template#delete GoogleComputeInstanceTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

