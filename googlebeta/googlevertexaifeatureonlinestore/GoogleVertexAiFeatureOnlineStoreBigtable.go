// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlevertexaifeatureonlinestore


type GoogleVertexAiFeatureOnlineStoreBigtable struct {
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.42.0/docs/resources/google_vertex_ai_feature_online_store#auto_scaling GoogleVertexAiFeatureOnlineStore#auto_scaling}
	AutoScaling *GoogleVertexAiFeatureOnlineStoreBigtableAutoScaling `field:"required" json:"autoScaling" yaml:"autoScaling"`
}

