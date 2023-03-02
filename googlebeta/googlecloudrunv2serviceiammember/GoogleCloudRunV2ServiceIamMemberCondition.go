package googlecloudrunv2serviceiammember


type GoogleCloudRunV2ServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service_iam_member#expression GoogleCloudRunV2ServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service_iam_member#title GoogleCloudRunV2ServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service_iam_member#description GoogleCloudRunV2ServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

