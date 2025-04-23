// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateServiceMesh struct {
	// The Mesh resource name. For more information see https://cloud.google.com/service-mesh/docs/reference/network-services/rest/v1/projects.locations.meshes#resource:-mesh.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.31.0/docs/resources/google_cloud_run_v2_service#mesh GoogleCloudRunV2Service#mesh}
	Mesh *string `field:"optional" json:"mesh" yaml:"mesh"`
}

