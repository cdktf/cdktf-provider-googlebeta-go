package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateReservationAffinity struct {
	// The type of reservation from which this instance can consume resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_template#type GoogleComputeRegionInstanceTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// specific_reservation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_template#specific_reservation GoogleComputeRegionInstanceTemplate#specific_reservation}
	SpecificReservation *GoogleComputeRegionInstanceTemplateReservationAffinitySpecificReservation `field:"optional" json:"specificReservation" yaml:"specificReservation"`
}

