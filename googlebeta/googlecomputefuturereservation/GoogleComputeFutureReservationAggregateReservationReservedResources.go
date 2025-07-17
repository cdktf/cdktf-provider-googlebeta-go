// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputefuturereservation


type GoogleComputeFutureReservationAggregateReservationReservedResources struct {
	// accelerator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_compute_future_reservation#accelerator GoogleComputeFutureReservation#accelerator}
	Accelerator *GoogleComputeFutureReservationAggregateReservationReservedResourcesAccelerator `field:"optional" json:"accelerator" yaml:"accelerator"`
}

