package googlebeyondcorpappconnector


type GoogleBeyondcorpAppConnectorPrincipalInfo struct {
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_beyondcorp_app_connector#service_account GoogleBeyondcorpAppConnector#service_account}
	ServiceAccount *GoogleBeyondcorpAppConnectorPrincipalInfoServiceAccount `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

