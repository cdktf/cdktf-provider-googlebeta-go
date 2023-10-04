// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatastreamstream


type GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects struct {
	// mysql_databases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.0.0/docs/resources/google_datastream_stream#mysql_databases GoogleDatastreamStream#mysql_databases}
	MysqlDatabases interface{} `field:"required" json:"mysqlDatabases" yaml:"mysqlDatabases"`
}

