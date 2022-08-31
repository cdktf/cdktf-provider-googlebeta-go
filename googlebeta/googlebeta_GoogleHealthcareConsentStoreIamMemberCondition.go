// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleHealthcareConsentStoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_consent_store_iam_member#expression GoogleHealthcareConsentStoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_consent_store_iam_member#title GoogleHealthcareConsentStoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_consent_store_iam_member#description GoogleHealthcareConsentStoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

