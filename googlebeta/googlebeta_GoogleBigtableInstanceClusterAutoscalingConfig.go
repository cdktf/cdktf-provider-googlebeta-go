// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleBigtableInstanceClusterAutoscalingConfig struct {
	// The target CPU utilization for autoscaling. Value must be between 10 and 80.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_instance#cpu_target GoogleBigtableInstance#cpu_target}
	CpuTarget *float64 `field:"required" json:"cpuTarget" yaml:"cpuTarget"`
	// The maximum number of nodes for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_instance#max_nodes GoogleBigtableInstance#max_nodes}
	MaxNodes *float64 `field:"required" json:"maxNodes" yaml:"maxNodes"`
	// The minimum number of nodes for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_instance#min_nodes GoogleBigtableInstance#min_nodes}
	MinNodes *float64 `field:"required" json:"minNodes" yaml:"minNodes"`
}

