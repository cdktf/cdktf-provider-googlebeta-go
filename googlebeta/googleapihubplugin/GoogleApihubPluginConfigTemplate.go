// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapihubplugin


type GoogleApihubPluginConfigTemplate struct {
	// additional_config_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_apihub_plugin#additional_config_template GoogleApihubPlugin#additional_config_template}
	AdditionalConfigTemplate interface{} `field:"optional" json:"additionalConfigTemplate" yaml:"additionalConfigTemplate"`
	// auth_config_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_apihub_plugin#auth_config_template GoogleApihubPlugin#auth_config_template}
	AuthConfigTemplate *GoogleApihubPluginConfigTemplateAuthConfigTemplate `field:"optional" json:"authConfigTemplate" yaml:"authConfigTemplate"`
}

