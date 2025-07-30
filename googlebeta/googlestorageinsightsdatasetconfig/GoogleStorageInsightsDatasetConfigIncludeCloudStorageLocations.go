// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlestorageinsightsdatasetconfig


type GoogleStorageInsightsDatasetConfigIncludeCloudStorageLocations struct {
	// The list of cloud storage locations to include in the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.46.0/docs/resources/google_storage_insights_dataset_config#locations GoogleStorageInsightsDatasetConfig#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
}

