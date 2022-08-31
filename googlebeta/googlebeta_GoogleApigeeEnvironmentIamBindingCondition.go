// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleApigeeEnvironmentIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apigee_environment_iam_binding#expression GoogleApigeeEnvironmentIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apigee_environment_iam_binding#title GoogleApigeeEnvironmentIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apigee_environment_iam_binding#description GoogleApigeeEnvironmentIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

