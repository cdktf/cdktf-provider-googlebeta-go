package googleservicedirectoryendpoint


type GoogleServiceDirectoryEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_directory_endpoint#create GoogleServiceDirectoryEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_directory_endpoint#delete GoogleServiceDirectoryEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_directory_endpoint#update GoogleServiceDirectoryEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
