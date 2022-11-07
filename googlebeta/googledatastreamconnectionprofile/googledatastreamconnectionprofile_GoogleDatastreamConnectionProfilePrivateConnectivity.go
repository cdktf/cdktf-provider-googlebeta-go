package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfilePrivateConnectivity struct {
	// A reference to a private connection resource. Format: 'projects/{project}/locations/{location}/privateConnections/{name}'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#private_connection GoogleDatastreamConnectionProfile#private_connection}
	PrivateConnection *string `field:"required" json:"privateConnection" yaml:"privateConnection"`
}

