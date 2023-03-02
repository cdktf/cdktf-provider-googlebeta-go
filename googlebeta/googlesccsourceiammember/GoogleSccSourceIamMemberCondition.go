package googlesccsourceiammember


type GoogleSccSourceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_scc_source_iam_member#expression GoogleSccSourceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_scc_source_iam_member#title GoogleSccSourceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_scc_source_iam_member#description GoogleSccSourceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

