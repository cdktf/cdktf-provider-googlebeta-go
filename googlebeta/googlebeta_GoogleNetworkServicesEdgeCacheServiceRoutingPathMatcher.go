// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcher struct {
	// The name to which this PathMatcher is referred by the HostRule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_services_edge_cache_service#name GoogleNetworkServicesEdgeCacheService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// route_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_services_edge_cache_service#route_rule GoogleNetworkServicesEdgeCacheService#route_rule}
	RouteRule interface{} `field:"required" json:"routeRule" yaml:"routeRule"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_services_edge_cache_service#description GoogleNetworkServicesEdgeCacheService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

