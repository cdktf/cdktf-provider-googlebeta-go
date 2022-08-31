// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleBigtableTableIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_table_iam_binding#expression GoogleBigtableTableIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_table_iam_binding#title GoogleBigtableTableIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_table_iam_binding#description GoogleBigtableTableIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

