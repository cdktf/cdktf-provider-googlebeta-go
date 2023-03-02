package googlealloydbcluster


type GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetention struct {
	// The retention period. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_alloydb_cluster#retention_period GoogleAlloydbCluster#retention_period}
	RetentionPeriod *string `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

