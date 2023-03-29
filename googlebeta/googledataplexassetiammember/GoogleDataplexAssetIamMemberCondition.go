package googledataplexassetiammember


type GoogleDataplexAssetIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_asset_iam_member#expression GoogleDataplexAssetIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_asset_iam_member#title GoogleDataplexAssetIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_asset_iam_member#description GoogleDataplexAssetIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

