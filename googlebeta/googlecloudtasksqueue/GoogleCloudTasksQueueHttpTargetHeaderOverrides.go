// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudtasksqueue


type GoogleCloudTasksQueueHttpTargetHeaderOverrides struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.14.0/docs/resources/google_cloud_tasks_queue#header GoogleCloudTasksQueue#header}
	Header *GoogleCloudTasksQueueHttpTargetHeaderOverridesHeader `field:"required" json:"header" yaml:"header"`
}

