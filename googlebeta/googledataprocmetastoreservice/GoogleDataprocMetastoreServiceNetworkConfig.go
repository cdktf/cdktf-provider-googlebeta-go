package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceNetworkConfig struct {
	// consumers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_service#consumers GoogleDataprocMetastoreService#consumers}
	Consumers interface{} `field:"required" json:"consumers" yaml:"consumers"`
}

