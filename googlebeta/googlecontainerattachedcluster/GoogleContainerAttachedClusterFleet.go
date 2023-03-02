package googlecontainerattachedcluster


type GoogleContainerAttachedClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_attached_cluster#project GoogleContainerAttachedCluster#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

