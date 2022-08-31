// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstances struct {
	// instances block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_connectivity_spoke#instances GoogleNetworkConnectivitySpoke#instances}
	Instances interface{} `field:"required" json:"instances" yaml:"instances"`
	// A value that controls whether site-to-site data transfer is enabled for these resources.
	//
	// Note that data transfer is available only in supported locations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_connectivity_spoke#site_to_site_data_transfer GoogleNetworkConnectivitySpoke#site_to_site_data_transfer}
	SiteToSiteDataTransfer interface{} `field:"required" json:"siteToSiteDataTransfer" yaml:"siteToSiteDataTransfer"`
}

