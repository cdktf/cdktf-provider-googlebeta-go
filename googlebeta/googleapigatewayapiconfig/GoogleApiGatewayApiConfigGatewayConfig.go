// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapigatewayapiconfig


type GoogleApiGatewayApiConfigGatewayConfig struct {
	// backend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.4.0/docs/resources/google_api_gateway_api_config#backend_config GoogleApiGatewayApiConfigA#backend_config}
	BackendConfig *GoogleApiGatewayApiConfigGatewayConfigBackendConfig `field:"required" json:"backendConfig" yaml:"backendConfig"`
}

