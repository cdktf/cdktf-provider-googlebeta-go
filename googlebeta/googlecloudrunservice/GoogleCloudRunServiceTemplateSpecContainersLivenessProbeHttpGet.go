package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersLivenessProbeHttpGet struct {
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.62.1/docs/resources/google_cloud_run_service#http_headers GoogleCloudRunService#http_headers}
	HttpHeaders interface{} `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Path to access on the HTTP server. If set, it should not be empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.62.1/docs/resources/google_cloud_run_service#path GoogleCloudRunService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

