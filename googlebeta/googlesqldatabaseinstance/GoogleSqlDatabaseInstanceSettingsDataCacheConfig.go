// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsDataCacheConfig struct {
	// Whether data cache is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.13.0/docs/resources/google_sql_database_instance#data_cache_enabled GoogleSqlDatabaseInstance#data_cache_enabled}
	DataCacheEnabled interface{} `field:"optional" json:"dataCacheEnabled" yaml:"dataCacheEnabled"`
}

