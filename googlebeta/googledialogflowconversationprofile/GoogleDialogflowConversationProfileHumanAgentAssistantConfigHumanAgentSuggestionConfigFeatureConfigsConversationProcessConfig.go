// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfig struct {
	// Number of recent non-small-talk sentences to use as context for article and FAQ suggestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_dialogflow_conversation_profile#recent_sentences_count GoogleDialogflowConversationProfile#recent_sentences_count}
	RecentSentencesCount *float64 `field:"optional" json:"recentSentencesCount" yaml:"recentSentencesCount"`
}

