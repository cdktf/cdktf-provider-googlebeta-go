// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleBigqueryDatasetAccessDatasetDataset struct {
	// The ID of the dataset containing this table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_dataset#dataset_id GoogleBigqueryDataset#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_dataset#project_id GoogleBigqueryDataset#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

