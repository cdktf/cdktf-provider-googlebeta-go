package googlecontainercluster


type GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy struct {
	// Number of blue nodes to drain in a batch.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#batch_node_count GoogleContainerCluster#batch_node_count}
	BatchNodeCount *float64 `field:"optional" json:"batchNodeCount" yaml:"batchNodeCount"`
	// Percentage of the blue pool nodes to drain in a batch.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#batch_percentage GoogleContainerCluster#batch_percentage}
	BatchPercentage *float64 `field:"optional" json:"batchPercentage" yaml:"batchPercentage"`
	// Soak time after each batch gets drained.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#batch_soak_duration GoogleContainerCluster#batch_soak_duration}
	BatchSoakDuration *string `field:"optional" json:"batchSoakDuration" yaml:"batchSoakDuration"`
}
