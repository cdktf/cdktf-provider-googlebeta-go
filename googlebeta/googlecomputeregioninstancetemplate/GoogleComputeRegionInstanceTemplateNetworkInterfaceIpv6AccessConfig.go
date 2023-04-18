package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateNetworkInterfaceIpv6AccessConfig struct {
	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet.
	//
	// Only PREMIUM tier is valid for IPv6
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_template#network_tier GoogleComputeRegionInstanceTemplate#network_tier}
	NetworkTier *string `field:"required" json:"networkTier" yaml:"networkTier"`
}

