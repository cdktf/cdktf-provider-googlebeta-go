package googledatastreamstream


type GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjects struct {
	// postgresql_schemas block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#postgresql_schemas GoogleDatastreamStream#postgresql_schemas}
	PostgresqlSchemas interface{} `field:"required" json:"postgresqlSchemas" yaml:"postgresqlSchemas"`
}

