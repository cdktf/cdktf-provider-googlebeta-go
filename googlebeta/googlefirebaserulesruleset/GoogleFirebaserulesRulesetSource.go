// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlefirebaserulesruleset


type GoogleFirebaserulesRulesetSource struct {
	// files block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.27.0/docs/resources/google_firebaserules_ruleset#files GoogleFirebaserulesRuleset#files}
	Files interface{} `field:"required" json:"files" yaml:"files"`
	// `Language` of the `Source` bundle. If unspecified, the language will default to `FIREBASE_RULES`. Possible values: LANGUAGE_UNSPECIFIED, FIREBASE_RULES, EVENT_FLOW_TRIGGERS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.27.0/docs/resources/google_firebaserules_ruleset#language GoogleFirebaserulesRuleset#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
}

