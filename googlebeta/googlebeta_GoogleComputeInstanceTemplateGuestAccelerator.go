// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeInstanceTemplateGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#count GoogleComputeInstanceTemplate#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource to expose to this instance. E.g. nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#type GoogleComputeInstanceTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

