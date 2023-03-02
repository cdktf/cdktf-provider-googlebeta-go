package googlebigquerydatapolicydatapolicyiambinding


type GoogleBigqueryDatapolicyDataPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_datapolicy_data_policy_iam_binding#expression GoogleBigqueryDatapolicyDataPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_datapolicy_data_policy_iam_binding#title GoogleBigqueryDatapolicyDataPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_datapolicy_data_policy_iam_binding#description GoogleBigqueryDatapolicyDataPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

