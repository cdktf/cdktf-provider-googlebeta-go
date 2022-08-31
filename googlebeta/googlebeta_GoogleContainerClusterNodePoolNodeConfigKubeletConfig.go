// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleContainerClusterNodePoolNodeConfigKubeletConfig struct {
	// Control the CPU management policy on the node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#cpu_manager_policy GoogleContainerCluster#cpu_manager_policy}
	CpuManagerPolicy *string `field:"required" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Enable CPU CFS quota enforcement for containers that specify CPU limits.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#cpu_cfs_quota GoogleContainerCluster#cpu_cfs_quota}
	CpuCfsQuota interface{} `field:"optional" json:"cpuCfsQuota" yaml:"cpuCfsQuota"`
	// Set the CPU CFS quota period value 'cpu.cfs_period_us'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#cpu_cfs_quota_period GoogleContainerCluster#cpu_cfs_quota_period}
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
}

