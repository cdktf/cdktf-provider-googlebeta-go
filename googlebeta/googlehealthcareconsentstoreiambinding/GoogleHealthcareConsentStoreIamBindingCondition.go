// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlehealthcareconsentstoreiambinding


type GoogleHealthcareConsentStoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.18.0/docs/resources/google_healthcare_consent_store_iam_binding#expression GoogleHealthcareConsentStoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.18.0/docs/resources/google_healthcare_consent_store_iam_binding#title GoogleHealthcareConsentStoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.18.0/docs/resources/google_healthcare_consent_store_iam_binding#description GoogleHealthcareConsentStoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

