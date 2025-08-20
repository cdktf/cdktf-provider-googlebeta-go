// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlestorageinsightsdatasetconfig


type GoogleStorageInsightsDatasetConfigExcludeCloudStorageBuckets struct {
	// cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_storage_insights_dataset_config#cloud_storage_buckets GoogleStorageInsightsDatasetConfig#cloud_storage_buckets}
	CloudStorageBuckets interface{} `field:"required" json:"cloudStorageBuckets" yaml:"cloudStorageBuckets"`
}

