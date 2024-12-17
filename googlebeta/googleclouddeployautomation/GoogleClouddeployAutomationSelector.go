// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleclouddeployautomation


type GoogleClouddeployAutomationSelector struct {
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.14.0/docs/resources/google_clouddeploy_automation#targets GoogleClouddeployAutomation#targets}
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
}

