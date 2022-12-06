package googlecontainercluster


type GoogleContainerClusterMasterAuthorizedNetworksConfig struct {
	// cidr_blocks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#cidr_blocks GoogleContainerCluster#cidr_blocks}
	CidrBlocks interface{} `field:"optional" json:"cidrBlocks" yaml:"cidrBlocks"`
	// Whether master is accessbile via Google Compute Engine Public IP addresses.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#gcp_public_cidrs_access_enabled GoogleContainerCluster#gcp_public_cidrs_access_enabled}
	GcpPublicCidrsAccessEnabled interface{} `field:"optional" json:"gcpPublicCidrsAccessEnabled" yaml:"gcpPublicCidrsAccessEnabled"`
}

