// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleSqlDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_database#create GoogleSqlDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_database#delete GoogleSqlDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_database#update GoogleSqlDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

