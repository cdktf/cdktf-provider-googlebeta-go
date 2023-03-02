package googlecloudbuildv2connectioniammember


type GoogleCloudbuildv2ConnectionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuildv2_connection_iam_member#expression GoogleCloudbuildv2ConnectionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuildv2_connection_iam_member#title GoogleCloudbuildv2ConnectionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuildv2_connection_iam_member#description GoogleCloudbuildv2ConnectionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

