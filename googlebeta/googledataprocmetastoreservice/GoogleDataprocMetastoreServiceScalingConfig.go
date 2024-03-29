// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceScalingConfig struct {
	// Metastore instance sizes. Possible values: ["EXTRA_SMALL", "SMALL", "MEDIUM", "LARGE", "EXTRA_LARGE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.22.0/docs/resources/google_dataproc_metastore_service#instance_size GoogleDataprocMetastoreService#instance_size}
	InstanceSize *string `field:"optional" json:"instanceSize" yaml:"instanceSize"`
	// Scaling factor, in increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.22.0/docs/resources/google_dataproc_metastore_service#scaling_factor GoogleDataprocMetastoreService#scaling_factor}
	ScalingFactor *float64 `field:"optional" json:"scalingFactor" yaml:"scalingFactor"`
}

