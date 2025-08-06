// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputepreviewfeature


type GoogleComputePreviewFeatureRolloutOperationRolloutInput struct {
	// Predefined rollout plans. Possible values: ["ROLLOUT_PLAN_FAST_ROLLOUT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_compute_preview_feature#predefined_rollout_plan GoogleComputePreviewFeature#predefined_rollout_plan}
	PredefinedRolloutPlan *string `field:"required" json:"predefinedRolloutPlan" yaml:"predefinedRolloutPlan"`
}

