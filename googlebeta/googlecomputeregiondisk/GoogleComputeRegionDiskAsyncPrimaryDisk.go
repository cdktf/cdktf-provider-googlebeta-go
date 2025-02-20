// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregiondisk


type GoogleComputeRegionDiskAsyncPrimaryDisk struct {
	// Primary disk for asynchronous disk replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.21.0/docs/resources/google_compute_region_disk#disk GoogleComputeRegionDisk#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
}

