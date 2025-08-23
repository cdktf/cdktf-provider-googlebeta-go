// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlestorageinsightsdatasetconfig


type GoogleStorageInsightsDatasetConfigSourceFolders struct {
	// The list of folder numbers to include in the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.2/docs/resources/google_storage_insights_dataset_config#folder_numbers GoogleStorageInsightsDatasetConfig#folder_numbers}
	FolderNumbers *[]*string `field:"optional" json:"folderNumbers" yaml:"folderNumbers"`
}

