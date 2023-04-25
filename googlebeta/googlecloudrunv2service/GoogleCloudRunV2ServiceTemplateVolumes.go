package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.63.0/docs/resources/google_cloud_run_v2_service#name GoogleCloudRunV2Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloud_sql_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.63.0/docs/resources/google_cloud_run_v2_service#cloud_sql_instance GoogleCloudRunV2Service#cloud_sql_instance}
	CloudSqlInstance *GoogleCloudRunV2ServiceTemplateVolumesCloudSqlInstance `field:"optional" json:"cloudSqlInstance" yaml:"cloudSqlInstance"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.63.0/docs/resources/google_cloud_run_v2_service#secret GoogleCloudRunV2Service#secret}
	Secret *GoogleCloudRunV2ServiceTemplateVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

