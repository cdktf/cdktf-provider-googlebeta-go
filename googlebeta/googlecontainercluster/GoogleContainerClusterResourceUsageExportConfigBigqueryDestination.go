package googlecontainercluster


type GoogleContainerClusterResourceUsageExportConfigBigqueryDestination struct {
	// The ID of a BigQuery Dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.63.1/docs/resources/google_container_cluster#dataset_id GoogleContainerCluster#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
}

