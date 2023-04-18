package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_template#key GoogleComputeRegionInstanceTemplate#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_template#operator GoogleComputeRegionInstanceTemplate#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_template#values GoogleComputeRegionInstanceTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

