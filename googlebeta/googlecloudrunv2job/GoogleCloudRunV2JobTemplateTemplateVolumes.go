package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_job#name GoogleCloudRunV2Job#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloud_sql_instance block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_job#cloud_sql_instance GoogleCloudRunV2Job#cloud_sql_instance}
	CloudSqlInstance *GoogleCloudRunV2JobTemplateTemplateVolumesCloudSqlInstance `field:"optional" json:"cloudSqlInstance" yaml:"cloudSqlInstance"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_job#secret GoogleCloudRunV2Job#secret}
	Secret *GoogleCloudRunV2JobTemplateTemplateVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

