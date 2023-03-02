package googledatastreamstream


type GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOracleColumns struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#column GoogleDatastreamStream#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// The Oracle data type. Full data types list can be found here: https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Data-Types.html.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#data_type GoogleDatastreamStream#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
}

