// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapihubplugin


type GoogleApihubPluginTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.42.0/docs/resources/google_apihub_plugin#create GoogleApihubPlugin#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.42.0/docs/resources/google_apihub_plugin#delete GoogleApihubPlugin#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

