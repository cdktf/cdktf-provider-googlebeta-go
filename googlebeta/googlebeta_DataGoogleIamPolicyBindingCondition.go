// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type DataGoogleIamPolicyBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_iam_policy#expression DataGoogleIamPolicy#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_iam_policy#title DataGoogleIamPolicy#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_iam_policy#description DataGoogleIamPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

