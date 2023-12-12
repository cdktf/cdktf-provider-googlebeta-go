// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecRulesRegexExpectation struct {
	// A regular expression the column value is expected to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.9.0/docs/resources/google_dataplex_datascan#regex GoogleDataplexDatascan#regex}
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

