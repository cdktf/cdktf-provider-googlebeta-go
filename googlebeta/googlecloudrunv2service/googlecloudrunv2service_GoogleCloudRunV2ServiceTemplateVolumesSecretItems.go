package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateVolumesSecretItems struct {
	// Integer octal mode bits to use on this file, must be a value between 01 and 0777 (octal).
	//
	// If 0 or not set, the Volume's default mode will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#mode GoogleCloudRunV2Service#mode}
	Mode *float64 `field:"required" json:"mode" yaml:"mode"`
	// The relative path of the secret in the container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#path GoogleCloudRunV2Service#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The Cloud Secret Manager secret version.
	//
	// Can be 'latest' for the latest value or an integer for a specific version
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#version GoogleCloudRunV2Service#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}
