// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleIapAppEngineVersionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_app_engine_version_iam_binding#expression GoogleIapAppEngineVersionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_app_engine_version_iam_binding#title GoogleIapAppEngineVersionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_app_engine_version_iam_binding#description GoogleIapAppEngineVersionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

