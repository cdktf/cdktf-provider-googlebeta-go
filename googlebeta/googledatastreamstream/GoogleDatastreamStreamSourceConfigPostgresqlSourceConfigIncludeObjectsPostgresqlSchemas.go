package googledatastreamstream


type GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemas struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#schema GoogleDatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// postgresql_tables block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#postgresql_tables GoogleDatastreamStream#postgresql_tables}
	PostgresqlTables interface{} `field:"optional" json:"postgresqlTables" yaml:"postgresqlTables"`
}

