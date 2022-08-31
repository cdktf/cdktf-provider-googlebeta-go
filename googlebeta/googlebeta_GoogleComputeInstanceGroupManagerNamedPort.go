// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeInstanceGroupManagerNamedPort struct {
	// The name of the port.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_group_manager#name GoogleComputeInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_group_manager#port GoogleComputeInstanceGroupManager#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

