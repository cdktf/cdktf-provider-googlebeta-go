// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleclouddeploytarget


type GoogleClouddeployTargetMultiTarget struct {
	// Required. The target_ids of this multiTarget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.5.0/docs/resources/google_clouddeploy_target#target_ids GoogleClouddeployTarget#target_ids}
	TargetIds *[]*string `field:"required" json:"targetIds" yaml:"targetIds"`
}

