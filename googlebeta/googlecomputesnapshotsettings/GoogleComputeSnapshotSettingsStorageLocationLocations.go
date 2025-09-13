// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputesnapshotsettings


type GoogleComputeSnapshotSettingsStorageLocationLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_compute_snapshot_settings#location GoogleComputeSnapshotSettings#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the location.
	//
	// It should be one of the Cloud Storage buckets.
	// Only one location can be specified. (should match location)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_compute_snapshot_settings#name GoogleComputeSnapshotSettings#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

