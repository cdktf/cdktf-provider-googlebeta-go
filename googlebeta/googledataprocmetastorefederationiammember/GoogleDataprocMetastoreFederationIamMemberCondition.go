// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataprocmetastorefederationiammember


type GoogleDataprocMetastoreFederationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_dataproc_metastore_federation_iam_member#expression GoogleDataprocMetastoreFederationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_dataproc_metastore_federation_iam_member#title GoogleDataprocMetastoreFederationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.16.0/docs/resources/google_dataproc_metastore_federation_iam_member#description GoogleDataprocMetastoreFederationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

