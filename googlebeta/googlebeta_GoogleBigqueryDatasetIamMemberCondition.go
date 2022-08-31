// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleBigqueryDatasetIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_dataset_iam_member#expression GoogleBigqueryDatasetIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_dataset_iam_member#title GoogleBigqueryDatasetIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_dataset_iam_member#description GoogleBigqueryDatasetIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

