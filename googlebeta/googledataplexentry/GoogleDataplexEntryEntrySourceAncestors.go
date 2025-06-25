// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataplexentry


type GoogleDataplexEntryEntrySourceAncestors struct {
	// The name of the ancestor resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.41.0/docs/resources/google_dataplex_entry#name GoogleDataplexEntry#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the ancestor resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.41.0/docs/resources/google_dataplex_entry#type GoogleDataplexEntry#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

