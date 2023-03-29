package googledataplexlakeiammember


type GoogleDataplexLakeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_lake_iam_member#expression GoogleDataplexLakeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_lake_iam_member#title GoogleDataplexLakeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_lake_iam_member#description GoogleDataplexLakeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

