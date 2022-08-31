// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleEndpointsServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_endpoints_service_iam_binding#expression GoogleEndpointsServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_endpoints_service_iam_binding#title GoogleEndpointsServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_endpoints_service_iam_binding#description GoogleEndpointsServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

