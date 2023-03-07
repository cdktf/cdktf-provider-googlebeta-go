package googlecontainercluster


type GoogleContainerClusterProtectConfig struct {
	// workload_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#workload_config GoogleContainerCluster#workload_config}
	WorkloadConfig *GoogleContainerClusterProtectConfigWorkloadConfig `field:"optional" json:"workloadConfig" yaml:"workloadConfig"`
	// WorkloadVulnerabilityMode defines mode to perform vulnerability scanning. Accepted values are WORKLOAD_VULNERABILITY_MODE_UNSPECIFIED, DISABLED, BASIC.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#workload_vulnerability_mode GoogleContainerCluster#workload_vulnerability_mode}
	WorkloadVulnerabilityMode *string `field:"optional" json:"workloadVulnerabilityMode" yaml:"workloadVulnerabilityMode"`
}

