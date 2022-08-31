// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeReservationSpecificReservation struct {
	// The number of resources that are allocated.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_reservation#count GoogleComputeReservation#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// instance_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_reservation#instance_properties GoogleComputeReservation#instance_properties}
	InstanceProperties *GoogleComputeReservationSpecificReservationInstanceProperties `field:"required" json:"instanceProperties" yaml:"instanceProperties"`
}

