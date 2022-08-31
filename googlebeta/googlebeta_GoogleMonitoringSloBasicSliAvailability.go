// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleMonitoringSloBasicSliAvailability struct {
	// Whether an availability SLI is enabled or not. Must be set to true. Defaults to 'true'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_slo#enabled GoogleMonitoringSlo#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

