// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization struct {
	// Fraction of backend capacity utilization (set in HTTP(s) load balancing configuration) that autoscaler should maintain.
	//
	// Must
	// be a positive float value. If not defined, the default is 0.8.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_autoscaler#target GoogleComputeAutoscaler#target}
	Target *float64 `field:"required" json:"target" yaml:"target"`
}

