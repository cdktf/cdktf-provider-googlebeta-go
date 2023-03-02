package googledatastreamstream


type GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemas struct {
	// Schema name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#schema GoogleDatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// oracle_tables block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#oracle_tables GoogleDatastreamStream#oracle_tables}
	OracleTables interface{} `field:"optional" json:"oracleTables" yaml:"oracleTables"`
}

