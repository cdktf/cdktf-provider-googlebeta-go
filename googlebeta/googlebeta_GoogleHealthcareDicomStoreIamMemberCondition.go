// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleHealthcareDicomStoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dicom_store_iam_member#expression GoogleHealthcareDicomStoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dicom_store_iam_member#title GoogleHealthcareDicomStoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dicom_store_iam_member#description GoogleHealthcareDicomStoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

