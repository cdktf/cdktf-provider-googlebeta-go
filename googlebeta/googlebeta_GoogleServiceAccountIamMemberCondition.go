// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleServiceAccountIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_account_iam_member#expression GoogleServiceAccountIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_account_iam_member#title GoogleServiceAccountIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_account_iam_member#description GoogleServiceAccountIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

