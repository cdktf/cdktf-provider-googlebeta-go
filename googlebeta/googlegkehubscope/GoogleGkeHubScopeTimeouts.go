// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkehubscope


type GoogleGkeHubScopeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.82.0/docs/resources/google_gke_hub_scope#create GoogleGkeHubScope#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.82.0/docs/resources/google_gke_hub_scope#delete GoogleGkeHubScope#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

