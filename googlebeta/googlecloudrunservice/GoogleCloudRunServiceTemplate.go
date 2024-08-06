// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunservice


type GoogleCloudRunServiceTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.40.0/docs/resources/google_cloud_run_service#metadata GoogleCloudRunService#metadata}
	Metadata *GoogleCloudRunServiceTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.40.0/docs/resources/google_cloud_run_service#spec GoogleCloudRunService#spec}
	Spec *GoogleCloudRunServiceTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

