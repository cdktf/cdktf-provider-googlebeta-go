// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudfunctionsfunction


type GoogleCloudfunctionsFunctionEventTriggerFailurePolicy struct {
	// Whether the function should be retried on failure. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_cloudfunctions_function#retry GoogleCloudfunctionsFunction#retry}
	Retry interface{} `field:"required" json:"retry" yaml:"retry"`
}

