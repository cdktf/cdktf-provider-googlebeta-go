package googlecontainerattachedcluster


type GoogleContainerAttachedClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.72.1/docs/resources/google_container_attached_cluster#project GoogleContainerAttachedCluster#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

