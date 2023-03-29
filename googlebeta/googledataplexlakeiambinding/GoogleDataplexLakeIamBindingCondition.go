package googledataplexlakeiambinding


type GoogleDataplexLakeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_lake_iam_binding#expression GoogleDataplexLakeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_lake_iam_binding#title GoogleDataplexLakeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_lake_iam_binding#description GoogleDataplexLakeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

