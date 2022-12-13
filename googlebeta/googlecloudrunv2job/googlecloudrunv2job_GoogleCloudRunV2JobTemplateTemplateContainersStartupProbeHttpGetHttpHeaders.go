package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateContainersStartupProbeHttpGetHttpHeaders struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_job#name GoogleCloudRunV2Job#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_job#value GoogleCloudRunV2Job#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

