// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig struct {
	// Confidence threshold of query result. This feature is only supported for types: ARTICLE_SUGGESTION, FAQ, SMART_REPLY, SMART_COMPOSE, KNOWLEDGE_SEARCH, KNOWLEDGE_ASSIST, ENTITY_EXTRACTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_dialogflow_conversation_profile#confidence_threshold GoogleDialogflowConversationProfile#confidence_threshold}
	ConfidenceThreshold *float64 `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// context_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_dialogflow_conversation_profile#context_filter_settings GoogleDialogflowConversationProfile#context_filter_settings}
	ContextFilterSettings *GoogleDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings `field:"optional" json:"contextFilterSettings" yaml:"contextFilterSettings"`
	// dialogflow_query_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_dialogflow_conversation_profile#dialogflow_query_source GoogleDialogflowConversationProfile#dialogflow_query_source}
	DialogflowQuerySource *GoogleDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource `field:"optional" json:"dialogflowQuerySource" yaml:"dialogflowQuerySource"`
	// Maximum number of results to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_dialogflow_conversation_profile#max_results GoogleDialogflowConversationProfile#max_results}
	MaxResults *float64 `field:"optional" json:"maxResults" yaml:"maxResults"`
	// sections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_dialogflow_conversation_profile#sections GoogleDialogflowConversationProfile#sections}
	Sections *GoogleDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSections `field:"optional" json:"sections" yaml:"sections"`
}

