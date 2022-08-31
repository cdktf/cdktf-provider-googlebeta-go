// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleApiGatewayGatewayIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_gateway_iam_binding#expression GoogleApiGatewayGatewayIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_gateway_iam_binding#title GoogleApiGatewayGatewayIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_gateway_iam_binding#description GoogleApiGatewayGatewayIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

