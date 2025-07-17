// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputestoragepooliambinding


type GoogleComputeStoragePoolIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_compute_storage_pool_iam_binding#expression GoogleComputeStoragePoolIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_compute_storage_pool_iam_binding#title GoogleComputeStoragePoolIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_compute_storage_pool_iam_binding#description GoogleComputeStoragePoolIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

