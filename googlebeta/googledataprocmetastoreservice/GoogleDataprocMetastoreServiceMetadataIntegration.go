// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceMetadataIntegration struct {
	// data_catalog_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.17.0/docs/resources/google_dataproc_metastore_service#data_catalog_config GoogleDataprocMetastoreService#data_catalog_config}
	DataCatalogConfig *GoogleDataprocMetastoreServiceMetadataIntegrationDataCatalogConfig `field:"required" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
}

