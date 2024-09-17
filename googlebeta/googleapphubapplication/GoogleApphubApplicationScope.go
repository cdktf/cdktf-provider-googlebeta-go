// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapphubapplication


type GoogleApphubApplicationScope struct {
	// Required. Scope Type.   Possible values: REGIONAL Possible values: ["REGIONAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.3.0/docs/resources/google_apphub_application#type GoogleApphubApplication#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

