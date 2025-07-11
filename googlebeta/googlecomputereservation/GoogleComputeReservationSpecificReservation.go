// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputereservation


type GoogleComputeReservationSpecificReservation struct {
	// The number of resources that are allocated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_compute_reservation#count GoogleComputeReservation#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// instance_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_compute_reservation#instance_properties GoogleComputeReservation#instance_properties}
	InstanceProperties *GoogleComputeReservationSpecificReservationInstanceProperties `field:"required" json:"instanceProperties" yaml:"instanceProperties"`
}

