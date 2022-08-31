// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeUrlMapDefaultRouteActionFaultInjectionPolicyAbort struct {
	// The HTTP status code used to abort the request. The value must be between 200 and 599 inclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_url_map#http_status GoogleComputeUrlMap#http_status}
	HttpStatus *float64 `field:"optional" json:"httpStatus" yaml:"httpStatus"`
	// The percentage of traffic (connections/operations/requests) which will be aborted as part of fault injection.
	//
	// The value must be between 0.0 and 100.0 inclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_url_map#percentage GoogleComputeUrlMap#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

