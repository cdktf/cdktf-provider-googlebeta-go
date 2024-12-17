// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googletpunode


type GoogleTpuNodeSchedulingConfig struct {
	// Defines whether the TPU instance is preemptible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.14.0/docs/resources/google_tpu_node#preemptible GoogleTpuNode#preemptible}
	Preemptible interface{} `field:"required" json:"preemptible" yaml:"preemptible"`
}

