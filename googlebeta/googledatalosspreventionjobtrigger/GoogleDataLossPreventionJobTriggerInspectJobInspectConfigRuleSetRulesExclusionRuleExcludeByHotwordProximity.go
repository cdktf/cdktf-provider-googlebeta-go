// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordProximity struct {
	// Number of characters after the finding to consider. Either this or window_before must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.4.0/docs/resources/google_data_loss_prevention_job_trigger#window_after GoogleDataLossPreventionJobTrigger#window_after}
	WindowAfter *float64 `field:"optional" json:"windowAfter" yaml:"windowAfter"`
	// Number of characters before the finding to consider. Either this or window_after must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.4.0/docs/resources/google_data_loss_prevention_job_trigger#window_before GoogleDataLossPreventionJobTrigger#window_before}
	WindowBefore *float64 `field:"optional" json:"windowBefore" yaml:"windowBefore"`
}

