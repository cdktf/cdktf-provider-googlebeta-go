package googlenetworkservicesedgecacheorigin


type GoogleNetworkServicesEdgeCacheOriginOriginRedirect struct {
	// The set of redirect response codes that the CDN follows. Values of [RedirectConditions](https://cloud.google.com/media-cdn/docs/reference/rest/v1/projects.locations.edgeCacheOrigins#redirectconditions) are accepted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_services_edge_cache_origin#redirect_conditions GoogleNetworkServicesEdgeCacheOrigin#redirect_conditions}
	RedirectConditions *[]*string `field:"optional" json:"redirectConditions" yaml:"redirectConditions"`
}
