package googlecloudrunv2jobiambinding


type GoogleCloudRunV2JobIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_job_iam_binding#expression GoogleCloudRunV2JobIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_job_iam_binding#title GoogleCloudRunV2JobIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_job_iam_binding#description GoogleCloudRunV2JobIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

