// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowcxenvironment


type GoogleDialogflowCxEnvironmentVersionConfigs struct {
	// Format: projects/{{project}}/locations/{{location}}/agents/{{agent}}/flows/{{flow}}/versions/{{version}}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.40.0/docs/resources/google_dialogflow_cx_environment#version GoogleDialogflowCxEnvironment#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

