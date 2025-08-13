// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlebigtableschemabundle


type GoogleBigtableSchemaBundleProtoSchema struct {
	// Base64 encoded content of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.48.0/docs/resources/google_bigtable_schema_bundle#proto_descriptors GoogleBigtableSchemaBundle#proto_descriptors}
	ProtoDescriptors *string `field:"required" json:"protoDescriptors" yaml:"protoDescriptors"`
}

