// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleApiGatewayApiConfigIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config_iam_binding#expression GoogleApiGatewayApiConfigIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config_iam_binding#title GoogleApiGatewayApiConfigIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config_iam_binding#description GoogleApiGatewayApiConfigIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

