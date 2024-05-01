// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleclouddeploydeliverypipelineiammember


type GoogleClouddeployDeliveryPipelineIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.27.0/docs/resources/google_clouddeploy_delivery_pipeline_iam_member#expression GoogleClouddeployDeliveryPipelineIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.27.0/docs/resources/google_clouddeploy_delivery_pipeline_iam_member#title GoogleClouddeployDeliveryPipelineIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.27.0/docs/resources/google_clouddeploy_delivery_pipeline_iam_member#description GoogleClouddeployDeliveryPipelineIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

