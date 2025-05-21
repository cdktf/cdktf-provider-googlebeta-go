// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeinstance


type GoogleComputeInstanceSchedulingOnInstanceStopAction struct {
	// If true, the contents of any attached Local SSD disks will be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.0/docs/resources/google_compute_instance#discard_local_ssd GoogleComputeInstance#discard_local_ssd}
	DiscardLocalSsd interface{} `field:"optional" json:"discardLocalSsd" yaml:"discardLocalSsd"`
}

