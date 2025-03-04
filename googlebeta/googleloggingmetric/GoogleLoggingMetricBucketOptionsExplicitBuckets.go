// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleloggingmetric


type GoogleLoggingMetricBucketOptionsExplicitBuckets struct {
	// The values must be monotonically increasing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.24.0/docs/resources/google_logging_metric#bounds GoogleLoggingMetric#bounds}
	Bounds *[]*float64 `field:"required" json:"bounds" yaml:"bounds"`
}

