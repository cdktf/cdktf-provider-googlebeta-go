// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.10.0/docs/resources/google_compute_instance_template#key GoogleComputeInstanceTemplate#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.10.0/docs/resources/google_compute_instance_template#operator GoogleComputeInstanceTemplate#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.10.0/docs/resources/google_compute_instance_template#values GoogleComputeInstanceTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

