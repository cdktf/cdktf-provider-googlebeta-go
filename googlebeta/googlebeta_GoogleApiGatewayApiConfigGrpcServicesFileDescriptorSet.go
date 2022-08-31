// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleApiGatewayApiConfigGrpcServicesFileDescriptorSet struct {
	// Base64 encoded content of the file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config#contents GoogleApiGatewayApiConfigA#contents}
	Contents *string `field:"required" json:"contents" yaml:"contents"`
	// The file path (full or relative path). This is typically the path of the file when it is uploaded.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config#path GoogleApiGatewayApiConfigA#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

