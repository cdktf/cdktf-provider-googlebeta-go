// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlespannerinstance


type GoogleSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverrides struct {
	// autoscaling_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.33.0/docs/resources/google_spanner_instance#autoscaling_limits GoogleSpannerInstance#autoscaling_limits}
	AutoscalingLimits *GoogleSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimits `field:"required" json:"autoscalingLimits" yaml:"autoscalingLimits"`
}

