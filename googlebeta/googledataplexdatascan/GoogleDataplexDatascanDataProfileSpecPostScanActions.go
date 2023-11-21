// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataplexdatascan


type GoogleDataplexDatascanDataProfileSpecPostScanActions struct {
	// bigquery_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dataplex_datascan#bigquery_export GoogleDataplexDatascan#bigquery_export}
	BigqueryExport *GoogleDataplexDatascanDataProfileSpecPostScanActionsBigqueryExport `field:"optional" json:"bigqueryExport" yaml:"bigqueryExport"`
}

