package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersEnvValueSource struct {
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#secret_key_ref GoogleCloudRunV2Service#secret_key_ref}
	SecretKeyRef *GoogleCloudRunV2ServiceTemplateContainersEnvValueSourceSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}
