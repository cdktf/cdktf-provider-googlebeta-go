// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapigeecontrolplaneaccess


type GoogleApigeeControlPlaneAccessTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.39.0/docs/resources/google_apigee_control_plane_access#create GoogleApigeeControlPlaneAccess#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.39.0/docs/resources/google_apigee_control_plane_access#delete GoogleApigeeControlPlaneAccess#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.39.0/docs/resources/google_apigee_control_plane_access#update GoogleApigeeControlPlaneAccess#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

