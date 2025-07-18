// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlebigquerytable


type GoogleBigqueryTableSchemaForeignTypeInfo struct {
	// Specifies the system which defines the foreign data type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_bigquery_table#type_system GoogleBigqueryTable#type_system}
	TypeSystem *string `field:"required" json:"typeSystem" yaml:"typeSystem"`
}

