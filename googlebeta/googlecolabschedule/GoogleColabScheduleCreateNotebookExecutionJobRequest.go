// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecolabschedule


type GoogleColabScheduleCreateNotebookExecutionJobRequest struct {
	// notebook_execution_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_colab_schedule#notebook_execution_job GoogleColabSchedule#notebook_execution_job}
	NotebookExecutionJob *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob `field:"required" json:"notebookExecutionJob" yaml:"notebookExecutionJob"`
}

