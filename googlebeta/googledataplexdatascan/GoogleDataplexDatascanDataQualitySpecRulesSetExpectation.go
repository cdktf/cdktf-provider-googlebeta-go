// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecRulesSetExpectation struct {
	// Expected values for the column value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.17.0/docs/resources/google_dataplex_datascan#values GoogleDataplexDatascan#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

