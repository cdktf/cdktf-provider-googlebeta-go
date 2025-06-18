// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatastreamstream


type GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets struct {
	// dataset_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.40.0/docs/resources/google_datastream_stream#dataset_template GoogleDatastreamStream#dataset_template}
	DatasetTemplate *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplate `field:"required" json:"datasetTemplate" yaml:"datasetTemplate"`
}

