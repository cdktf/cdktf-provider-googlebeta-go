// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapigeesecurityaction


type GoogleApigeeSecurityActionFlagHeaders struct {
	// The header name to be sent to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_apigee_security_action#name GoogleApigeeSecurityAction#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header value to be sent to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_apigee_security_action#value GoogleApigeeSecurityAction#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

