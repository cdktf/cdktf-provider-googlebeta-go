package googlebigquerytable


type GoogleBigqueryTableExternalDataConfigurationAvroOptions struct {
	// If sourceFormat is set to "AVRO", indicates whether to interpret logical types as the corresponding BigQuery data type (for example, TIMESTAMP), instead of using the raw type (for example, INTEGER).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table#use_avro_logical_types GoogleBigqueryTable#use_avro_logical_types}
	UseAvroLogicalTypes interface{} `field:"required" json:"useAvroLogicalTypes" yaml:"useAvroLogicalTypes"`
}
