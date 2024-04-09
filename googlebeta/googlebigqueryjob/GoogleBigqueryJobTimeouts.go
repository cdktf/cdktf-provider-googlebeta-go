// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryjob


type GoogleBigqueryJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.24.0/docs/resources/google_bigquery_job#create GoogleBigqueryJob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.24.0/docs/resources/google_bigquery_job#delete GoogleBigqueryJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.24.0/docs/resources/google_bigquery_job#update GoogleBigqueryJob#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

