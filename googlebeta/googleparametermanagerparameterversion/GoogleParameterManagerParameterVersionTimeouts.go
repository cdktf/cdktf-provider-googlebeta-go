// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleparametermanagerparameterversion


type GoogleParameterManagerParameterVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.31.0/docs/resources/google_parameter_manager_parameter_version#create GoogleParameterManagerParameterVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.31.0/docs/resources/google_parameter_manager_parameter_version#delete GoogleParameterManagerParameterVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.31.0/docs/resources/google_parameter_manager_parameter_version#update GoogleParameterManagerParameterVersion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

