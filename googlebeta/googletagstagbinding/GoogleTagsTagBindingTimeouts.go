// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googletagstagbinding


type GoogleTagsTagBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.11.2/docs/resources/google_tags_tag_binding#create GoogleTagsTagBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.11.2/docs/resources/google_tags_tag_binding#delete GoogleTagsTagBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

