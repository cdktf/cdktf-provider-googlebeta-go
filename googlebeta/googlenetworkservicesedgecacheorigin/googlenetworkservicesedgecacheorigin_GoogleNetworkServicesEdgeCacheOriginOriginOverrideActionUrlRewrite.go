package googlenetworkservicesedgecacheorigin


type GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite struct {
	// Prior to forwarding the request to the selected origin, the request's host header is replaced with contents of the hostRewrite.
	//
	// This value must be between 1 and 255 characters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_services_edge_cache_origin#host_rewrite GoogleNetworkServicesEdgeCacheOrigin#host_rewrite}
	HostRewrite *string `field:"optional" json:"hostRewrite" yaml:"hostRewrite"`
}
