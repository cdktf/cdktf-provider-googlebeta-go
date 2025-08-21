// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients struct {
	// The email recipients who will receive the DataQualityScan results report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_dataplex_datascan#emails GoogleDataplexDatascan#emails}
	Emails *[]*string `field:"optional" json:"emails" yaml:"emails"`
}

