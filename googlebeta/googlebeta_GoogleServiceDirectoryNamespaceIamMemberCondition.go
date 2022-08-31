// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleServiceDirectoryNamespaceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_directory_namespace_iam_member#expression GoogleServiceDirectoryNamespaceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_directory_namespace_iam_member#title GoogleServiceDirectoryNamespaceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_directory_namespace_iam_member#description GoogleServiceDirectoryNamespaceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

