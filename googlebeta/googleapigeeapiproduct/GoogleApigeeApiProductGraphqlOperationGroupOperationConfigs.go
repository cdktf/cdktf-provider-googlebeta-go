// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapigeeapiproduct


type GoogleApigeeApiProductGraphqlOperationGroupOperationConfigs struct {
	// Required. Name of the API proxy endpoint or remote service with which the GraphQL operation and quota are associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.46.0/docs/resources/google_apigee_api_product#api_source GoogleApigeeApiProduct#api_source}
	ApiSource *string `field:"optional" json:"apiSource" yaml:"apiSource"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.46.0/docs/resources/google_apigee_api_product#attributes GoogleApigeeApiProduct#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.46.0/docs/resources/google_apigee_api_product#operations GoogleApigeeApiProduct#operations}
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.46.0/docs/resources/google_apigee_api_product#quota GoogleApigeeApiProduct#quota}
	Quota *GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsQuota `field:"optional" json:"quota" yaml:"quota"`
}

