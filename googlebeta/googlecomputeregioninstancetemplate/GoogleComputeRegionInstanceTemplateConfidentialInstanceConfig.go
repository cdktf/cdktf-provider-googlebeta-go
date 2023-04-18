package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateConfidentialInstanceConfig struct {
	// Defines whether the instance should have confidential compute enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_template#enable_confidential_compute GoogleComputeRegionInstanceTemplate#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"required" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

