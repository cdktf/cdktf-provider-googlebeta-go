// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleBigqueryConnectionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_connection_iam_member#expression GoogleBigqueryConnectionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_connection_iam_member#title GoogleBigqueryConnectionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_connection_iam_member#description GoogleBigqueryConnectionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

