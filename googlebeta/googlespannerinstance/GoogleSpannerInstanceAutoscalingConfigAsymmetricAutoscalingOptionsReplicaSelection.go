// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlespannerinstance


type GoogleSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsReplicaSelection struct {
	// The location of the replica to apply asymmetric autoscaling options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.14.0/docs/resources/google_spanner_instance#location GoogleSpannerInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
}

