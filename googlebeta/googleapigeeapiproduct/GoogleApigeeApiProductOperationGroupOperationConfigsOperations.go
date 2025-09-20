// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapigeeapiproduct


type GoogleApigeeApiProductOperationGroupOperationConfigsOperations struct {
	// Methods refers to the REST verbs, when none specified, all verb types are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.50.0/docs/resources/google_apigee_api_product#methods GoogleApigeeApiProduct#methods}
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// Required. REST resource path associated with the API proxy or remote service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.50.0/docs/resources/google_apigee_api_product#resource GoogleApigeeApiProduct#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

