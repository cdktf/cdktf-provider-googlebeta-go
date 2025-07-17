// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatafusioninstanceiambinding


type GoogleDataFusionInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_data_fusion_instance_iam_binding#expression GoogleDataFusionInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_data_fusion_instance_iam_binding#title GoogleDataFusionInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_data_fusion_instance_iam_binding#description GoogleDataFusionInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

