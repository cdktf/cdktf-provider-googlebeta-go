// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeresourcepolicy


type GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule struct {
	// day_of_weeks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.27.0/docs/resources/google_compute_resource_policy#day_of_weeks GoogleComputeResourcePolicy#day_of_weeks}
	DayOfWeeks interface{} `field:"required" json:"dayOfWeeks" yaml:"dayOfWeeks"`
}

