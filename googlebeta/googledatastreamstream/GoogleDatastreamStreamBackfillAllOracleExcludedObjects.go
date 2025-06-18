// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatastreamstream


type GoogleDatastreamStreamBackfillAllOracleExcludedObjects struct {
	// oracle_schemas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.40.0/docs/resources/google_datastream_stream#oracle_schemas GoogleDatastreamStream#oracle_schemas}
	OracleSchemas interface{} `field:"required" json:"oracleSchemas" yaml:"oracleSchemas"`
}

