package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfilePostgresqlProfile struct {
	// Database for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#database GoogleDatastreamConnectionProfile#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Hostname for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#hostname GoogleDatastreamConnectionProfile#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Password for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#password GoogleDatastreamConnectionProfile#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#username GoogleDatastreamConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Port for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#port GoogleDatastreamConnectionProfile#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}
