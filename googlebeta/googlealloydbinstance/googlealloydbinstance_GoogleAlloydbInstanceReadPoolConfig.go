package googlealloydbinstance


type GoogleAlloydbInstanceReadPoolConfig struct {
	// Read capacity, i.e. number of nodes in a read pool instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_alloydb_instance#node_count GoogleAlloydbInstance#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
}

