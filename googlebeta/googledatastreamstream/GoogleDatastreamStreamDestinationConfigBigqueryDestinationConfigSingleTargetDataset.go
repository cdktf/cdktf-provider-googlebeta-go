package googledatastreamstream


type GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset struct {
	// Dataset ID in the format projects/{project}/datasets/{dataset_id} or {project}:{dataset_id}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#dataset_id GoogleDatastreamStream#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
}

