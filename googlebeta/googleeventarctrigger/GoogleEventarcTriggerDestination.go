// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleeventarctrigger


type GoogleEventarcTriggerDestination struct {
	// cloud_run_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_eventarc_trigger#cloud_run_service GoogleEventarcTrigger#cloud_run_service}
	CloudRunService *GoogleEventarcTriggerDestinationCloudRunService `field:"optional" json:"cloudRunService" yaml:"cloudRunService"`
	// gke block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_eventarc_trigger#gke GoogleEventarcTrigger#gke}
	Gke *GoogleEventarcTriggerDestinationGke `field:"optional" json:"gke" yaml:"gke"`
	// The resource name of the Workflow whose Executions are triggered by the events.
	//
	// The Workflow resource should be deployed in the same project as the trigger. Format: `projects/{project}/locations/{location}/workflows/{workflow}`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_eventarc_trigger#workflow GoogleEventarcTrigger#workflow}
	Workflow *string `field:"optional" json:"workflow" yaml:"workflow"`
}

