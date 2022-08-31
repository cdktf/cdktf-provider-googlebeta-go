// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleSecretManagerSecretIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret_iam_member#expression GoogleSecretManagerSecretIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret_iam_member#title GoogleSecretManagerSecretIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret_iam_member#description GoogleSecretManagerSecretIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

