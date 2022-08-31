// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRules struct {
	// exclusion_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#exclusion_rule GoogleDataLossPreventionInspectTemplate#exclusion_rule}
	ExclusionRule *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRule `field:"optional" json:"exclusionRule" yaml:"exclusionRule"`
	// hotword_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#hotword_rule GoogleDataLossPreventionInspectTemplate#hotword_rule}
	HotwordRule *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRule `field:"optional" json:"hotwordRule" yaml:"hotwordRule"`
}

