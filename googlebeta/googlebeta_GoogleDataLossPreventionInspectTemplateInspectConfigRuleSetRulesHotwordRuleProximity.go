// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity struct {
	// Number of characters after the finding to consider. Either this or window_before must be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#window_after GoogleDataLossPreventionInspectTemplate#window_after}
	WindowAfter *float64 `field:"optional" json:"windowAfter" yaml:"windowAfter"`
	// Number of characters before the finding to consider. Either this or window_after must be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#window_before GoogleDataLossPreventionInspectTemplate#window_before}
	WindowBefore *float64 `field:"optional" json:"windowBefore" yaml:"windowBefore"`
}

