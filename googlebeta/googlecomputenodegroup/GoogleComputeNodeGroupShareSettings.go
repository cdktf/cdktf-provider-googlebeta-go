package googlecomputenodegroup


type GoogleComputeNodeGroupShareSettings struct {
	// Node group sharing type. Possible values: ["ORGANIZATION", "SPECIFIC_PROJECTS", "LOCAL"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_node_group#share_type GoogleComputeNodeGroup#share_type}
	ShareType *string `field:"required" json:"shareType" yaml:"shareType"`
	// project_map block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_node_group#project_map GoogleComputeNodeGroup#project_map}
	ProjectMap interface{} `field:"optional" json:"projectMap" yaml:"projectMap"`
}

