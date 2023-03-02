package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig struct {
	// Resource name of an existing Dataproc Cluster to act as a Spark History Server for the workload.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster#dataproc_cluster GoogleDataprocCluster#dataproc_cluster}
	DataprocCluster *string `field:"optional" json:"dataprocCluster" yaml:"dataprocCluster"`
}

