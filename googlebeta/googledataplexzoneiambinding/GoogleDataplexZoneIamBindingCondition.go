package googledataplexzoneiambinding


type GoogleDataplexZoneIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_zone_iam_binding#expression GoogleDataplexZoneIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_zone_iam_binding#title GoogleDataplexZoneIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_zone_iam_binding#description GoogleDataplexZoneIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

