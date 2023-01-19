package googledatastreamstream


type GoogleDatastreamStreamSourceConfig struct {
	// mysql_source_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#mysql_source_config GoogleDatastreamStream#mysql_source_config}
	MysqlSourceConfig *GoogleDatastreamStreamSourceConfigMysqlSourceConfig `field:"required" json:"mysqlSourceConfig" yaml:"mysqlSourceConfig"`
	// Source connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#source_connection_profile GoogleDatastreamStream#source_connection_profile}
	SourceConnectionProfile *string `field:"required" json:"sourceConnectionProfile" yaml:"sourceConnectionProfile"`
}

