// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapihubplugininstance


type GoogleApihubPluginInstanceActionsCurationConfigCustomCuration struct {
	// The unique name of the curation resource. This will be the name of the curation resource in the format: 'projects/{project}/locations/{location}/curations/{curation}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_apihub_plugin_instance#curation GoogleApihubPluginInstance#curation}
	Curation *string `field:"required" json:"curation" yaml:"curation"`
}

