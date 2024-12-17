// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatastreamstream


type GoogleDatastreamStreamBackfillAllSqlServerExcludedObjects struct {
	// schemas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.14.0/docs/resources/google_datastream_stream#schemas GoogleDatastreamStream#schemas}
	Schemas interface{} `field:"required" json:"schemas" yaml:"schemas"`
}

