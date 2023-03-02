package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig struct {
	// Required. The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuildv2_connection#service GoogleCloudbuildv2Connection#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

