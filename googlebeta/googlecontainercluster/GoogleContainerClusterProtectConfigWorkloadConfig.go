package googlecontainercluster


type GoogleContainerClusterProtectConfigWorkloadConfig struct {
	// Sets which mode of auditing should be used for the cluster's workloads. Accepted values are DISABLED, BASIC.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#audit_mode GoogleContainerCluster#audit_mode}
	AuditMode *string `field:"required" json:"auditMode" yaml:"auditMode"`
}

