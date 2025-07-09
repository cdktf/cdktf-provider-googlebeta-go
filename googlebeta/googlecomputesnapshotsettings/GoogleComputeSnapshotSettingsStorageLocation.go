// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputesnapshotsettings


type GoogleComputeSnapshotSettingsStorageLocation struct {
	// The chosen location policy Possible values: ["NEAREST_MULTI_REGION", "LOCAL_REGION", "SPECIFIC_LOCATIONS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_compute_snapshot_settings#policy GoogleComputeSnapshotSettings#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_compute_snapshot_settings#locations GoogleComputeSnapshotSettings#locations}
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
}

