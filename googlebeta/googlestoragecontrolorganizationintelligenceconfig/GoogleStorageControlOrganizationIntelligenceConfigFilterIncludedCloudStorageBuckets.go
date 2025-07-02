// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlestoragecontrolorganizationintelligenceconfig


type GoogleStorageControlOrganizationIntelligenceConfigFilterIncludedCloudStorageBuckets struct {
	// List of bucket id regexes to exclude in the storage intelligence plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.42.0/docs/resources/google_storage_control_organization_intelligence_config#bucket_id_regexes GoogleStorageControlOrganizationIntelligenceConfig#bucket_id_regexes}
	BucketIdRegexes *[]*string `field:"required" json:"bucketIdRegexes" yaml:"bucketIdRegexes"`
}

