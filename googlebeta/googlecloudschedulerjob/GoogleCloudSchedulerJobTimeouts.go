// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudschedulerjob


type GoogleCloudSchedulerJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_cloud_scheduler_job#create GoogleCloudSchedulerJob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_cloud_scheduler_job#delete GoogleCloudSchedulerJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_cloud_scheduler_job#update GoogleCloudSchedulerJob#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

