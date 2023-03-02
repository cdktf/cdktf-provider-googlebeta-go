package googledatastreamstream


type GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabases struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#database GoogleDatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// mysql_tables block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#mysql_tables GoogleDatastreamStream#mysql_tables}
	MysqlTables interface{} `field:"optional" json:"mysqlTables" yaml:"mysqlTables"`
}

