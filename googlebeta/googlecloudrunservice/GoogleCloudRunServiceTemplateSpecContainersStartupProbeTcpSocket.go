package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersStartupProbeTcpSocket struct {
	// Port number to access on the container. Number must be in the range 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.62.0/docs/resources/google_cloud_run_service#port GoogleCloudRunService#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

