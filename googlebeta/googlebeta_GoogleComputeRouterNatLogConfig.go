// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeRouterNatLogConfig struct {
	// Indicates whether or not to export logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_nat#enable GoogleComputeRouterNat#enable}
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Specifies the desired filtering of logs on this NAT. Possible values: ["ERRORS_ONLY", "TRANSLATIONS_ONLY", "ALL"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_nat#filter GoogleComputeRouterNat#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
}

