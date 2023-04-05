package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersEnv struct {
	// Name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#name GoogleCloudRunService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Defaults to "".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#value GoogleCloudRunService#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_from block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#value_from GoogleCloudRunService#value_from}
	ValueFrom *GoogleCloudRunServiceTemplateSpecContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

