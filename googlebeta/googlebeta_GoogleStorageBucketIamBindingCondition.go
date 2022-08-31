// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleStorageBucketIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket_iam_binding#expression GoogleStorageBucketIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket_iam_binding#title GoogleStorageBucketIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket_iam_binding#description GoogleStorageBucketIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

