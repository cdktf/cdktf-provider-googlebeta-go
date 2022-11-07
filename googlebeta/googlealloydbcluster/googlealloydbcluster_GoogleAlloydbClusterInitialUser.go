package googlealloydbcluster


type GoogleAlloydbClusterInitialUser struct {
	// The initial password for the user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_alloydb_cluster#password GoogleAlloydbCluster#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The database username.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_alloydb_cluster#user GoogleAlloydbCluster#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

