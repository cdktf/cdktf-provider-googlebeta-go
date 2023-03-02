package googlecontainercluster


type GoogleContainerClusterGatewayApiConfig struct {
	// The Gateway API release channel to use for Gateway API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#channel GoogleContainerCluster#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
}

