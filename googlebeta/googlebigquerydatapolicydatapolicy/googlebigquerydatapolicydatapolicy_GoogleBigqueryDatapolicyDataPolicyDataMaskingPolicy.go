package googlebigquerydatapolicydatapolicy


type GoogleBigqueryDatapolicyDataPolicyDataMaskingPolicy struct {
	// The available masking rules. Learn more here: https://cloud.google.com/bigquery/docs/column-data-masking-intro#masking_options. Possible values: ["SHA256", "ALWAYS_NULL", "DEFAULT_MASKING_VALUE"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_datapolicy_data_policy#predefined_expression GoogleBigqueryDatapolicyDataPolicy#predefined_expression}
	PredefinedExpression *string `field:"required" json:"predefinedExpression" yaml:"predefinedExpression"`
}
