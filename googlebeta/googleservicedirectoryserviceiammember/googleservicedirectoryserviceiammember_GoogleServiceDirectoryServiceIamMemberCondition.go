package googleservicedirectoryserviceiammember


type GoogleServiceDirectoryServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_directory_service_iam_member#expression GoogleServiceDirectoryServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_directory_service_iam_member#title GoogleServiceDirectoryServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_directory_service_iam_member#description GoogleServiceDirectoryServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}
