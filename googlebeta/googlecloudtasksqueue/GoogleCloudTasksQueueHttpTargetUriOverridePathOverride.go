// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudtasksqueue


type GoogleCloudTasksQueueHttpTargetUriOverridePathOverride struct {
	// The URI path (e.g., /users/1234). Default is an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.18.1/docs/resources/google_cloud_tasks_queue#path GoogleCloudTasksQueue#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

