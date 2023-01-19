package googledatastreamstream


type GoogleDatastreamStreamBackfillAll struct {
	// mysql_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#mysql_excluded_objects GoogleDatastreamStream#mysql_excluded_objects}
	MysqlExcludedObjects *GoogleDatastreamStreamBackfillAllMysqlExcludedObjects `field:"optional" json:"mysqlExcludedObjects" yaml:"mysqlExcludedObjects"`
}

