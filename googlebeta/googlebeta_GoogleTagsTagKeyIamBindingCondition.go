// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleTagsTagKeyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_key_iam_binding#expression GoogleTagsTagKeyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_key_iam_binding#title GoogleTagsTagKeyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_key_iam_binding#description GoogleTagsTagKeyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

