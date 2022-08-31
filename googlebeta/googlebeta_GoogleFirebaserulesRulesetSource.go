// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleFirebaserulesRulesetSource struct {
	// files block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebaserules_ruleset#files GoogleFirebaserulesRuleset#files}
	Files interface{} `field:"required" json:"files" yaml:"files"`
	// `Language` of the `Source` bundle. If unspecified, the language will default to `FIREBASE_RULES`. Possible values: LANGUAGE_UNSPECIFIED, FIREBASE_RULES, EVENT_FLOW_TRIGGERS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebaserules_ruleset#language GoogleFirebaserulesRuleset#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
}

