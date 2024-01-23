// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersEnvFromSecretRef struct {
	// local_object_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.13.0/docs/resources/google_cloud_run_service#local_object_reference GoogleCloudRunService#local_object_reference}
	LocalObjectReference *GoogleCloudRunServiceTemplateSpecContainersEnvFromSecretRefLocalObjectReference `field:"optional" json:"localObjectReference" yaml:"localObjectReference"`
	// Specify whether the Secret must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.13.0/docs/resources/google_cloud_run_service#optional GoogleCloudRunService#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

