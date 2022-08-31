// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeRegionBackendServiceConsistentHashHttpCookie struct {
	// Name of the cookie.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_backend_service#name GoogleComputeRegionBackendService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Path to set for the cookie.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_backend_service#path GoogleComputeRegionBackendService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_backend_service#ttl GoogleComputeRegionBackendService#ttl}
	Ttl *GoogleComputeRegionBackendServiceConsistentHashHttpCookieTtl `field:"optional" json:"ttl" yaml:"ttl"`
}

