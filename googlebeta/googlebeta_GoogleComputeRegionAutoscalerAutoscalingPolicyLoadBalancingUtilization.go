// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization struct {
	// Fraction of backend capacity utilization (set in HTTP(s) load balancing configuration) that autoscaler should maintain.
	//
	// Must
	// be a positive float value. If not defined, the default is 0.8.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_autoscaler#target GoogleComputeRegionAutoscaler#target}
	Target *float64 `field:"required" json:"target" yaml:"target"`
}

