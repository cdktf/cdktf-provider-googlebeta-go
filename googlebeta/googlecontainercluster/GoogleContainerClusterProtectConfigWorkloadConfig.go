package googlecontainercluster


type GoogleContainerClusterProtectConfigWorkloadConfig struct {
	// Mode defines how to audit the workload configs. Accepted values are MODE_UNSPECIFIED, DISABLED, BASIC.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#audit_mode GoogleContainerCluster#audit_mode}
	AuditMode *string `field:"required" json:"auditMode" yaml:"auditMode"`
}

