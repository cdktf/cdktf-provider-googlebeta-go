package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersStartupProbeHttpGet struct {
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.63.0/docs/resources/google_cloud_run_v2_service#http_headers GoogleCloudRunV2Service#http_headers}
	HttpHeaders interface{} `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Path to access on the HTTP server. Defaults to '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.63.0/docs/resources/google_cloud_run_v2_service#path GoogleCloudRunV2Service#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

