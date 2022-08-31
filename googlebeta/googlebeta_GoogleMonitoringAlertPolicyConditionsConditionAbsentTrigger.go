// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleMonitoringAlertPolicyConditionsConditionAbsentTrigger struct {
	// The absolute number of time series that must fail the predicate for the condition to be triggered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_alert_policy#count GoogleMonitoringAlertPolicy#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The percentage of time series that must fail the predicate for the condition to be triggered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_alert_policy#percent GoogleMonitoringAlertPolicy#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

