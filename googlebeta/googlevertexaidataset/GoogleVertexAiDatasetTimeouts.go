// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlevertexaidataset


type GoogleVertexAiDatasetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.4.0/docs/resources/google_vertex_ai_dataset#create GoogleVertexAiDataset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.4.0/docs/resources/google_vertex_ai_dataset#delete GoogleVertexAiDataset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.4.0/docs/resources/google_vertex_ai_dataset#update GoogleVertexAiDataset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

