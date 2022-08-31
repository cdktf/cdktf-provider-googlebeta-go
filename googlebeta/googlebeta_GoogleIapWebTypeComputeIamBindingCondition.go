// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleIapWebTypeComputeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_type_compute_iam_binding#expression GoogleIapWebTypeComputeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_type_compute_iam_binding#title GoogleIapWebTypeComputeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_type_compute_iam_binding#description GoogleIapWebTypeComputeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

