// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfig struct {
	// Version of current baseline model.
	//
	// It will be ignored if model is set. Valid versions are: Article Suggestion baseline model: - 0.9 - 1.0 (default) Summarization baseline model: - 1.0
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_dialogflow_conversation_profile#baseline_model_version GoogleDialogflowConversationProfile#baseline_model_version}
	BaselineModelVersion *string `field:"optional" json:"baselineModelVersion" yaml:"baselineModelVersion"`
	// Conversation model resource name. Format: projects/<Project ID>/conversationModels/<Model ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_dialogflow_conversation_profile#model GoogleDialogflowConversationProfile#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
}

