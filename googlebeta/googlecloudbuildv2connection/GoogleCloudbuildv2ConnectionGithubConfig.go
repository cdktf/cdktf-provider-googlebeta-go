package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionGithubConfig struct {
	// GitHub App installation id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuildv2_connection#app_installation_id GoogleCloudbuildv2Connection#app_installation_id}
	AppInstallationId *float64 `field:"optional" json:"appInstallationId" yaml:"appInstallationId"`
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuildv2_connection#authorizer_credential GoogleCloudbuildv2Connection#authorizer_credential}
	AuthorizerCredential *GoogleCloudbuildv2ConnectionGithubConfigAuthorizerCredential `field:"optional" json:"authorizerCredential" yaml:"authorizerCredential"`
}

