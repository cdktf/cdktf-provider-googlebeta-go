package googledataproccluster


type GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity struct {
	// The URI of a sole-tenant that the cluster will be created on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster#node_group_uri GoogleDataprocCluster#node_group_uri}
	NodeGroupUri *string `field:"required" json:"nodeGroupUri" yaml:"nodeGroupUri"`
}

