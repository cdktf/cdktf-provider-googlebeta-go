// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GooglePrivatecaCaPoolIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool_iam_binding#expression GooglePrivatecaCaPoolIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool_iam_binding#title GooglePrivatecaCaPoolIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool_iam_binding#description GooglePrivatecaCaPoolIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

